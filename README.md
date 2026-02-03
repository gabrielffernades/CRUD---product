# CRUD Product – Go (PoC)

This project is a **proof of concept (PoC)** built to **practice and refine Go fundamentals** using a domain I already know well: **RESTful CRUD APIs**.

One of the main motivations for starting this project is to **gain a deeper understanding of Go’s core and theoretical concepts**, while also applying studies related to **project organization, code structure, incremental steps, and engineering decision-making**.

Instead of focusing on business complexity, the goal is to clearly understand how a Go backend application is structured and executed — from starting an HTTP server to routing requests, handling HTTP methods, and returning JSON responses in an idiomatic way.

So far, the project includes:

- An HTTP server built using Go’s standard `net/http` package
- Explicit route registration with method-based handling (GET / POST)
- Clear separation of concerns following a domain-oriented structure
- Domain models defined using Go structs
- A service layer responsible for orchestrating business logic
- A repository layer simulating persistence using in-memory storage
- Automatic ID generation at the persistence layer
- JSON request/response handling using typed structs
- A simple and intentional structure prepared for incremental evolution

This repository is **not intended to be production-ready**. It serves as a **controlled environment for experimentation and learning**, prioritizing clarity, correctness, and idiomatic Go over unnecessary abstractions.

---

## Available Routes

| Method | Route        | Description |
|------|-------------|-------------|
| GET  | `/health`   | Health check endpoint used to verify that the HTTP server is running and responding correctly. Returns status `200 OK`. |
| GET  | `/products` | Returns a list of products in JSON format. Data is retrieved from an in-memory repository. |
| POST | `/products` | Creates a new product. The request body is validated and converted into a domain entity, persisted in memory, and returned with a generated ID. |

These routes establish a complete read/write HTTP flow and serve as the foundation for expanding the CRUD operations.

---

## Project Structure

```text
CRUD---product/
├── database/
│   └── connection.go     # Database connection placeholder (future use)
├── handlers/
│   └── health.go         # Health check HTTP handler
├── product/
│   ├── routes.go         # HTTP method routing (GET / POST)
│   ├── handler.go        # HTTP handlers (request/response translation)
│   ├── service.go        # Domain logic and orchestration
│   ├── repository.go    # Data persistence (in-memory)
│   └── model.go          # Domain models and input contracts
├── main.go               # Application entry point and HTTP server setup
├── go.mod
└── README.md



## Running the project

```bash
go run .
