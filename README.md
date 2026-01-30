# CRUD Product – Go (PoC)

This project is a **proof of concept (PoC)** built to **practice and refine Go fundamentals** through a domain I already know well: **RESTful CRUD APIs**.

Instead of focusing on business complexity, the goal was to deeply understand how a Go backend application is structured and executed — from starting an HTTP server to handling requests and returning JSON responses in an idiomatic way.

So far, the project includes:

- An HTTP server built using Go’s standard `net/http` package
- Clear and minimal route registration
- Separation of concerns by isolating HTTP handlers by domain
- A simple and intentional structure prepared for future growth

This repository is **not intended to be production-ready**. It serves as a **controlled environment for experimentation and incremental improvement**, focusing on clarity, readability, and correct use of Go’s standard library.

---

## Available Routes

| Method | Route       | Description |
|------|------------|-------------|
| GET  | `/health`  | Health check endpoint used to verify that the HTTP server is running and responding correctly. Returns status `200 OK` with a simple response body. |
| GET  | `/products` | Returns a list of products in JSON format. At this stage, the data is mocked and used only to validate request handling, JSON encoding, and response structure. |

These routes establish the base HTTP flow of the application and serve as the foundation for the upcoming CRUD implementation.

---

## Next Steps

Planned evolutions for the project include:

- Defining domain structs for products
- Introducing a service layer for business logic
- Adding persistence with a relational database
- Expanding the API to support full CRUD operations

---

## Running the project

```bash
go run .
