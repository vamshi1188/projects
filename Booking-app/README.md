# Booking-app: Go Conference Ticket Booking System

## Description:

This project is a simple Go application that simulates a conference ticket booking system for the "go conference". Users can enter their information and book tickets, with the application tracking available tickets and validating user input.

## Features:

- Book tickets for the "go conference"
- Check remaining ticket availability
- Validate user information (name, email, number of tickets)
- Store booking information
- Simulate sending ticket confirmation emails

## Getting Started:

1. Clone the repository: `git clone https://github.com/vamshi1188/projects/tree/main/Booking-app`
2. Navigate to the project directory: `cd Booking-app`
3. Run the application: `go run main.go`
   
## Getting Started with Docker:

1. Clone the repository: `git clone https://github.com/vamshi1188/projects.git`
2. Navigate to the project directory: `cd Booking-app`
3. Build the Docker image: ` docker build -t Booking-app`
4. Run the Docker container: `docker run -it Booking-app`


## How it Works:

1. The application welcomes users and displays the conference name and available tickets.
2. Users enter their name, email, and desired number of tickets.
3. The application validates the input:
   - Names must be at least 2 characters long.
   - Emails must contain "@" and ".com" or ".in".
   - Desired number of tickets must be positive and not exceed remaining tickets.
4. If valid, the application books the tickets, updates availability, and sends a simulated confirmation email.
5. Bookings are stored for reference.
6. The application halts when all tickets are sold out.

## Dependencies:

This project only requires the standard Go libraries and no external dependencies.

## Contributing:

We welcome contributions! Feel free to submit pull requests with improvements, bug fixes, or new features.


## Additional Notes:

- This is a basic example and can be extended with features like user authentication, different ticket types, payment processing, etc.
- The "sendTicket" function currently simulates sending an email and doesn't integrate with an actual email service.

## Additional Resources:

- Code repository: [Booking-app](https://github.com/vamshi1188/projects/tree/main/Booking-app)
- Go programming language: [golang.org](https://golang.org/)

I hope this README provides a clear and informative overview of your Booking-app project! Feel free to customize it further with screenshots, documentation links, or additional information specific to your application.
