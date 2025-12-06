package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
)

const (
	radiusMeters = 20000 // 20km in meters
	channelName  = "proximity-chat"
)

var (
	pool     *redis.Pool
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	// Track active connections
	connections = make(map[string]*Client)
	connMutex   sync.RWMutex
)

type Client struct {
	UserID   string
	Username string
	Lat      float64
	Lon      float64
	Conn     *websocket.Conn
	LastSeen time.Time
}

type Message struct {
	Type     string   `json:"type"`
	UserID   string   `json:"userId,omitempty"`
	Username string   `json:"username,omitempty"`
	Text     string   `json:"text,omitempty"`
	Lat      float64  `json:"lat,omitempty"`
	Lon      float64  `json:"lon,omitempty"`
	Count    int      `json:"count,omitempty"`
	Users    []string `json:"users,omitempty"`
}

func main() {
	var addr string
	flag.StringVar(&addr, "tile38", ":9851", "Tile38 Address")
	flag.Parse()

	// Initialize Tile38 connection pool
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
		TestOnBorrow: func(conn redis.Conn, _ time.Time) error {
			if resp, _ := redis.String(conn.Do("PING")); resp != "PONG" {
				return errors.New("expected PONG")
			}
			return nil
		},
	}

	// Setup geofence subscription
	go subscribeToGeofence()

	// Cleanup inactive users
	go cleanupInactiveUsers()

	// HTTP handlers
	http.HandleFunc("/ws", handleWebSocket)
	http.Handle("/", http.FileServer(http.Dir("web")))

	port := ":8080"
	log.Printf("🌍 Proximity Chat Server running on %s", port)
	log.Printf("📍 Radius: %d meters (%.1f km)", radiusMeters, float64(radiusMeters)/1000)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	var client *Client

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			if client != nil {
				handleLeave(client)
			}
			break
		}

		switch msg.Type {
		case "join":
			client = &Client{
				UserID:   msg.UserID,
				Username: msg.Username,
				Lat:      msg.Lat,
				Lon:      msg.Lon,
				Conn:     conn,
				LastSeen: time.Now(),
			}
			handleJoin(client)

		case "message":
			if client != nil {
				client.Lat = msg.Lat
				client.Lon = msg.Lon
				client.LastSeen = time.Now()
				handleMessage(client, msg.Text)
			}

		case "update_location":
			if client != nil {
				client.Lat = msg.Lat
				client.Lon = msg.Lon
				client.LastSeen = time.Now()
				updateUserLocation(client)
			}

		case "leave":
			if client != nil {
				handleLeave(client)
			}
			return
		}
	}
}

func handleJoin(client *Client) {
	connMutex.Lock()
	connections[client.UserID] = client
	connMutex.Unlock()

	// Store user location in Tile38
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "users", client.UserID,
		"POINT", client.Lat, client.Lon,
		"EX", 300, // Expire after 5 minutes
		"FIELD", "username", client.Username)

	if err != nil {
		log.Printf("Error storing user in Tile38: %v", err)
		return
	}

	log.Printf("✅ User joined: %s (%s) at [%.4f, %.4f]",
		client.Username, client.UserID, client.Lat, client.Lon)

	// Notify user of nearby users
	notifyNearbyUsers(client)
}

func handleMessage(client *Client, text string) {
	// Get nearby users within radius
	nearbyUserIDs := getNearbyUsers(client.Lat, client.Lon)

	msg := Message{
		Type:     "message",
		UserID:   client.UserID,
		Username: client.Username,
		Text:     text,
	}

	// Broadcast to nearby users including sender
	broadcastToUsers(nearbyUserIDs, msg)

	log.Printf("💬 %s: %s (sent to %d nearby users)",
		client.Username, text, len(nearbyUserIDs))
}

func handleLeave(client *Client) {
	connMutex.Lock()
	delete(connections, client.UserID)
	connMutex.Unlock()

	// Remove from Tile38
	conn := pool.Get()
	defer conn.Close()
	conn.Do("DEL", "users", client.UserID)

	log.Printf("👋 User left: %s (%s)", client.Username, client.UserID)
}

func updateUserLocation(client *Client) {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "users", client.UserID,
		"POINT", client.Lat, client.Lon,
		"EX", 300,
		"FIELD", "username", client.Username)

	if err != nil {
		log.Printf("Error updating location: %v", err)
	}
}

func getNearbyUsers(lat, lon float64) []string {
	conn := pool.Get()
	defer conn.Close()

	// Use NEARBY to find users within radius
	reply, err := redis.Values(conn.Do("NEARBY", "users",
		"POINT", lat, lon, radiusMeters))

	if err != nil {
		log.Printf("Error getting nearby users: %v", err)
		return []string{}
	}

	var userIDs []string
	if len(reply) > 1 {
		objects, _ := redis.Values(reply[1], nil)
		for _, obj := range objects {
			data, _ := redis.Strings(obj, nil)
			if len(data) > 0 {
				userIDs = append(userIDs, data[0])
			}
		}
	}

	return userIDs
}

func notifyNearbyUsers(client *Client) {
	nearbyUserIDs := getNearbyUsers(client.Lat, client.Lon)

	msg := Message{
		Type:  "user_count",
		Count: len(nearbyUserIDs) - 1, // Exclude self
	}

	if err := client.Conn.WriteJSON(msg); err != nil {
		log.Printf("Error sending nearby count: %v", err)
	}
}

func broadcastToUsers(userIDs []string, msg Message) {
	connMutex.RLock()
	defer connMutex.RUnlock()

	for _, userID := range userIDs {
		if client, ok := connections[userID]; ok {
			if err := client.Conn.WriteJSON(msg); err != nil {
				log.Printf("Error sending to %s: %v", userID, err)
			}
		}
	}
}

func subscribeToGeofence() {
	for {
		err := func() error {
			conn := pool.Get()
			defer conn.Close()

			// Set up geofence channel for the entire users collection
			_, err := conn.Do("SETCHAN", channelName,
				"NEARBY", "users", "FENCE", "DETECT", "enter,exit",
				"POINT", 0, 0, 50000000) // Global fence

			if err != nil {
				return err
			}

			psc := redis.PubSubConn{Conn: pool.Get()}
			defer psc.Close()

			if err := psc.Subscribe(channelName); err != nil {
				return err
			}

			log.Printf("📡 Subscribed to geofence channel: %s", channelName)

			for {
				switch v := psc.Receive().(type) {
				case redis.Message:
					// Process geofence events if needed
					// For now, we're using direct queries
				case error:
					return v
				}
			}
		}()

		if err != nil {
			log.Printf("Geofence subscription error: %v", err)
		}
		time.Sleep(time.Second * 5)
	}
}

func cleanupInactiveUsers() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		connMutex.Lock()
		now := time.Now()
		for userID, client := range connections {
			if now.Sub(client.LastSeen) > 5*time.Minute {
				log.Printf("🧹 Cleaning up inactive user: %s", client.Username)
				delete(connections, userID)

				conn := pool.Get()
				conn.Do("DEL", "users", userID)
				conn.Close()
			}
		}
		connMutex.Unlock()
	}
}
