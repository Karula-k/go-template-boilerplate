# GO Boilerplate

A minimal and extensible boilerplate project for building web applications in Go. It integrates essential tools and patterns for modern Go development including SQL database integration, hot reloading, and code generation.

## Features

- ⚙️ **Go** — The powerful and efficient programming language.
- ⚡ **Fiber** — An Express-inspired web framework for Go.
- 🧩 **sqlc** — Type-safe SQL query generator for Go.
- 🔁 **Air** — Live reload for Go apps during development.

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (>= 1.20)
- [Fiber](https://github.com/gofiber/fiber)
- [sqlc](https://github.com/kyleconroy/sqlc#installation)
- [Air](https://github.com/cosmtrek/air#installation)
- A SQL database (e.g., PostgreSQL)

---

### Installation

1. **Clone the repository:**

```bash
git clone https://github.com/yourusername/go-boilerplate.git
cd go-boilerplate
```

2. **Prerequire**

```bash
go install github.com/air-verse/air@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

3. **Run the project**

```bash
make run
```

## Goals

- [x] login

- [x] register

- [x] Swagger

- [ ] permission management

- [x] env management
