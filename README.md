# CRUD Product – Go (PoC)

This project is a **proof of concept (PoC)** built to **practice and refine Go fundamentals** using a domain I already know well: **RESTful CRUD APIs**.

One of the main motivations for starting this project is to **gain a deeper understanding of Go’s core and theoretical concepts**, while also applying studies related to **project organization, code structure, incremental steps, and engineering decision-making**.

Instead of focusing on business complexity, the goal is to clearly understand how a Go backend application is structured and executed — from starting an HTTP server to registering routes, handling requests, and returning JSON responses in an idiomatic way.

So far, the project includes:

- An HTTP server built using Go’s standard `net/http` package
- Explicit route registration and request handling
- Clear separation of concerns by isolating HTTP handlers by domain
- A defined domain model using Go structs
- JSON responses encoded directly from typed domain models
- A simple and intentional structure prepared for incremental evolution

This repository is **not intended to be production-ready**. It serves as a **controlled environment for experimentation and learning**, prioritizing clarity, correctness, and idiomatic Go over unnecessary abstractions.

---

## Available Routes

| Method | Route        | Description |
|------|-------------|-------------|
| GET  | `/health`   | Health check endpoint used to verify that the HTTP server is running and responding correctly. Returns status `200 OK` with a simple response body. |
| GET  | `/products` | Returns a list of products in JSON format. The data is currently mocked and served from an in-memory slice of typed `Product` structs, validating request handling, JSON encoding, and response structure. |

These routes establish the base HTTP flow of the application and serve as the foundation for the upcoming CRUD implementation.

---

## Project Structure

- `main.go` — application entry point and HTTP server setup
- `handlers/` — HTTP handlers (e.g. health check)
- `product/`
  - `model.go` — domain definition for `Product`
  - `handler.go` — HTTP handlers related to products
  - `service.go` — placeholder for business logic
  - `repository.go` — placeholder for persistence logic

---

## Running the project

```bash
go run .
