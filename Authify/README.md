<div align="center">

# 🔐 JWT Authentication App

<p align="center">
  <strong>A secure, full-stack authentication system built with Go & JavaScript</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black" alt="JavaScript">
  <img src="https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white" alt="SQLite">
  <img src="https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white" alt="Bootstrap">
</p>

<p align="center">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://img.shields.io/badge/Version-1.0.0-blue.svg" alt="Version">
  <img src="https://img.shields.io/badge/Status-Active-success.svg" alt="Status">
</p>

</div>

---

## 📖 Overview

A modern authentication system featuring **JWT tokens**, **secure cookie sessions**, and **real-time validation**. Built with **Go Fiber** backend and **Vanilla JavaScript** frontend for maximum performance and simplicity.

---

## 🏗️ Project Architecture

<table>
<tr>
<td width="50%">

### 🛠️ **Backend (Go + Fiber)**
- 🚀 **High-performance** RESTful API
- 🔑 **JWT** token-based authentication
- 🍚 **Secure HttpOnly** cookie sessions
- 📄 **SQLite** database with GORM
- 🛡️ **bcrypt** password hashing

</td>
<td width="50%">

### 🎨 **Frontend (Vanilla JS)**
- ⚡ **Single Page Application** (SPA)
- 🧱 **Client-side routing**
- ✨ **Dynamic UI** rendering
- 📱 **Responsive** Bootstrap design
- ⚡ **Real-time** form validation

</td>
</tr>
</table>

---

## 🚀 Quick Start

### Prerequisites
- Go 1.19+ installed
- Node.js 16+ installed

### 🔧 Backend Setup

```bash
# Navigate to backend directory
cd backend

# Install dependencies
go mod tidy

# Run the server
go run main.go
```

> 🎯 **Backend runs on:** `http://localhost:8000`

### 🌐 Frontend Setup

```bash
# Navigate to frontend directory
cd frontend

# Start the file server
node server.js
```

> 🎯 **Frontend available at:** `http://localhost:3000`

---

## 🔌 API Reference

<div align="center">

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `POST` | `/api/register` | 📝 Register new user | ❌ |
| `POST` | `/api/login` | 🔐 User login | ❌ |
| `GET` | `/api/user` | 👤 Get user profile | ✅ |
| `POST` | `/api/logout` | 🚪 User logout | ✅ |

</div>

---

## ✨ Key Features

<div align="center">

| Feature | Description |
|---------|-------------|
| 🔐 **Secure Authentication** | JWT tokens with HttpOnly cookies |
| 🛡️ **Password Security** | bcrypt hashing with salt rounds |
| 🔒 **Protected Routes** | Authentication middleware |
| 📱 **Responsive Design** | Mobile-first Bootstrap UI |
| ⚡ **Real-time Validation** | Instant form feedback |
| 🔄 **Smart Redirects** | Context-aware navigation |
| 🍚 **Secure Sessions** | XSS-protected cookie storage |
| ⏰ **Token Expiry** | 24-hour automatic timeout |

</div>

---

## 🔒 Security Features

<table>
<tr>
<td align="center">🔐</td>
<td><strong>Password Hashing</strong><br>bcrypt with configurable salt rounds</td>
</tr>
<tr>
<td align="center">🍚</td>
<td><strong>HttpOnly Cookies</strong><br>XSS protection for JWT storage</td>
</tr>
<tr>
<td align="center">🌐</td>
<td><strong>CORS Configuration</strong><br>Restricted to frontend origin only</td>
</tr>
<tr>
<td align="center">⏰</td>
<td><strong>Token Expiration</strong><br>Automatic 24-hour timeout</td>
</tr>
</table>

---

## 🛠️ Technology Stack

### Backend Dependencies
```go
// Core Framework
github.com/gofiber/fiber/v2

// Database ORM
gorm.io/gorm
gorm.io/driver/sqlite

// JWT Handling
github.com/golang-jwt/jwt/v4

// Password Hashing
golang.org/x/crypto/bcrypt

// Testing
github.com/stretchr/testify
```

### Frontend Dependencies
- 🎨 **Bootstrap 5** - Modern UI framework
- ⚡ **Vanilla JavaScript** - No framework overhead
- 🌐 **Fetch API** - Native HTTP requests

---

## 📁 Project Structure

```
JWTAUTHAPI/
├── backend/
│   ├── main.go              # Server entry point
│   ├── controllers/         # Handler logic
│   ├── database/            # DB connection logic
│   ├── models/              # Data models
│   ├── routes/              # API route setup
│   ├── tests/               # Unit & Integration tests
│   └── go.mod               # Go dependencies
├── frontend/
│   ├── index.html           # Main HTML file
│   ├── script.js            # JavaScript logic
│   ├── style.css            # Custom styles
│   └── server.js            # Static file server
└── README.md
```

---

## 🎯 Usage Examples

### User Registration
```javascript
const response = await fetch('/api/register', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    name: 'John Doe',
    email: 'john@example.com',
    password: 'securePassword123'
  })
});
```

### Protected Route Access
```javascript
const userInfo = await fetch('/api/user', {
  credentials: 'include' // Include cookies
});
```

---

## 📃 Testing Overview

### Testing Frameworks Used
- ✅ `testing` - Go's built-in testing package
- ✅ `testify/assert` - for fluent assertions
- ✅ `httptest` - simulate HTTP requests
- ✅ `sqlite :memory:` - in-memory DB for integration tests

### How to Run Tests
```bash
cd backend/tests
./coverage.sh
```

### Coverage Screenshot

![Test Coverage](coverage-screenshot.png)

> ✅ **Achieved 83% code coverage** across unit, integration, and API tests.

---

## 👥 Contributing

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

---

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

---

<div align="center">

### ✨ **Star this repo if you found it helpful!** ✨

<p>
  <a href="https://github.com/yourusername/JWTAUTHAPI/issues">🐛 Report Bug</a> •
  <a href="https://github.com/yourusername/JWTAUTHAPI/issues">✨ Request Feature</a>
</p>

</div>
