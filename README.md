# 🚌 Bus API

A modern RESTful API built in Go, empowering users to create, manage, and explore bus routes with real-time mapping features via MapBox and secure JWT-based authentication.

---

## 🚀 Overview

**Bus API** enables developers and businesses to streamline route management for buses, shuttles, or similar transport services. With robust endpoints, seamless MapBox integration, and secure user authentication, it’s ready for production use or learning projects.

---

## ✨ Features

- **Route Management**: Create, update, delete, and retrieve bus routes.
- **MapBox Integration**: Visualize and manage routes on interactive maps.
- **JWT Authentication**: Secure all endpoints with JSON Web Tokens.
- **RESTful API Design**: Intuitive and standardized endpoints.
- **Scalable Architecture**: Built with Go for speed and reliability.

---

## ⚡️ Quick Start

### Prerequisites

- [Go](https://golang.org/dl/) (v1.18+ recommended)
- [MapBox API Key](https://account.mapbox.com/)
- (Optional) [Docker](https://www.docker.com/) for containerized deployment -- this is going to be added later please don't worry about this now

### Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/AdamSoftware/BusAPI.git
   cd BusAPI
   ```

2. **Set Environment Variables:**

   Create a `.env` file with:

   ```
   MAPBOX_API_KEY=your_mapbox_key
   JWT_SECRET=your_jwt_secret 'U can also bypase this by just logging in everytime when u are trying things'
   ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run the application:**

   ```bash
   go run main.go
   ```

---

## 🛣️ API Endpoints

| Method | Endpoint            | Description                   | Auth Required |
|--------|---------------------|-------------------------------|--------------|
| POST   | `/signup`           | Register a new user           | ❌           |
| POST   | `/login`            | Obtain JWT token              | ❌           |
| GET    | `/routes`           | List all routes               | ✅           |
| POST   | `/routes`           | Create a new route            | ✅           |
| GET    | `/routes/{id}`      | Get route details             | ✅           |
| PUT    | `/routes/{id}`      | Update a route                | ✅           |
| DELETE | `/routes/{id}`      | Delete a route                | ✅           |

*All protected endpoints require an `Authorization: Bearer <token>` header.*

---

## 🔒 Authentication

- Register and log in to get a JWT token.
- Include the JWT token in the `Authorization` header for all protected requests.

---

## 🗺️ MapBox Integration

- Uses MapBox APIs to visualize and manage bus routes.
- Ensure you set your MapBox API key in the environment.

---

## 🧑‍💻 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss your ideas.

---

## 📄 License

MIT

---

## 📫 Contact

Questions, suggestions, or want to collaborate?  
[Open an issue](https://github.com/AdamSoftware/BusAPI/issues) or email me at [AdamSteinberg16@gmail.com].

---

> “The journey of a thousand miles begins with a single route.” 🚍
