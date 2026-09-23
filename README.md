# Valentine Omondi Awili — Personal Developer Portfolio

An editorial-style personal portfolio website for **Valentine Omondi Awili**, Full-Stack Developer. Built with modern, clean engineering practices using **React + Vite** on the frontend, a clean **Go REST API** on the backend, and **SQLite** for lightweight, robust persistence.

---

## Architecture Overview

```
MyPortfolio/
├── frontend/             # React + Vite client application
│   ├── src/
│   │   ├── components/  # Modular UI components (Navbar, Hero, About, Projects, etc.)
│   │   ├── pages/       # Page views
│   │   ├── services/    # API communication layer
│   │   ├── assets/      # Static assets & images
│   │   ├── App.jsx
│   │   └── main.jsx
│   └── package.json
│
├── backend/              # Go REST API backend
│   ├── cmd/server/      # Server entry point
│   ├── internal/        # Handlers, services, repositories, models, database
│   └── go.mod
│
├── documentations/       # Project specifications & documentation
│   ├── PRD.md           # Product Requirements Document
│   ├── Architecture.md  # Architectural layout & technical decisions
│   ├── Design.md        # Design system & visual specifications
│   ├── Security.md      # Security principles & defense-in-depth rules
│   ├── Memory.md        # Agent memory & project guidelines
│   └── Task.md          # Comprehensive step-by-step task breakdown
│
├── Done.md               # Log of completed tasks and verification steps
├── README.md
└── .gitignore
```

---

## Tech Stack

- **Frontend:** React, Vite, JavaScript, CSS (Fluid typography with `clamp()`, CSS variables, responsive design)
- **Backend:** Go (Standard library HTTP / net/http, REST API, JSON endpoints)
- **Database:** SQLite (Relational storage for projects, skills, contact submissions)
- **Tooling:** Git, Linux, npm, Go toolchain

---

## Getting Started

### Prerequisites

- Node.js (v18+)
- Go (v1.22+)
- SQLite3

### Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

The frontend will run locally at `http://localhost:5173`.

### Backend Setup

```bash
cd backend
go run cmd/server/main.go
```

The backend server will run on port `8080` (or configured `PORT`), serving API endpoints at `http://localhost:8080/api`.

---

## Documentation

Full architectural and design specifications are maintained in the [`documentations/`](./documentations/) directory:
- [PRD.md](./documentations/PRD.md)
- [Architecture.md](./documentations/Architecture.md)
- [Design.md](./documentations/Design.md)
- [Security.md](./documentations/Security.md)
- [Task.md](./documentations/Task.md)
- [Memory.md](./documentations/Memory.md)

---

## Development Progress

Progress tracking and verification steps are maintained in [Done.md](./Done.md).
