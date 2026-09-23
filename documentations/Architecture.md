1. Project Overview

This project is a personal portfolio website built as a full-stack web application.

The portfolio will present the developer's background, technical skills, projects, and contact information through a modern, responsive interface.

The application will use:

Frontend: React + Vite
Backend: Go
API: REST API
Database: SQLite
Styling: CSS
Frontend communication: HTTP/JSON
Version control: Git

The architecture separates the user interface from backend logic so that the frontend and backend can be developed, tested, and maintained independently.

2. Application Goals

The application should:

Present a professional personal brand.
Clearly communicate the developer's technical abilities.
Showcase completed and ongoing projects.
Allow visitors to learn about the developer.
Provide a simple way for visitors to make contact.
Provide a foundation that can later be expanded without restructuring the entire application.
3. High-Level Architecture
                         USER
                           │
                           ▼
                  ┌─────────────────┐
                  │   React + Vite  │
                  │    Frontend     │
                  └────────┬────────┘
                           │
                    HTTP / JSON
                           │
                           ▼
                  ┌─────────────────┐
                  │    Go API       │
                  │    Backend      │
                  └────────┬────────┘
                           │
                           ▼
                  ┌─────────────────┐
                  │    Services     │
                  │ Business Logic  │
                  └────────┬────────┘
                           │
                           ▼
                  ┌─────────────────┐
                  │    SQLite DB    │
                  └─────────────────┘

The frontend is responsible for the user interface and user interactions.

The Go backend is responsible for:

API endpoints
Business logic
Data validation
Database operations
Contact form processing
Project data retrieval

The SQLite database stores dynamic portfolio information.

4. Frontend Architecture
Technology
React
Vite
JavaScript
CSS
React Router
Fetch API

React will be responsible for building reusable UI components.

Vite will provide the development server and production build system.

React Router will manage navigation between the portfolio pages.

5. Frontend Page Structure

The application will contain five primary pages:

/
├── Home
├── About
├── Projects
├── Skills
└── Contact

Navigation should allow the visitor to move between all five pages without a full browser reload.

6. Home Page
Purpose

The Home page is the main entry point of the portfolio.

It should immediately communicate:

Who the developer is
What the developer does
Main technical focus
A clear call to action
UI Structure
┌──────────────────────────────────────────────┐
│ LOGO / NAME             Home About Projects │
│                          Skills Contact      │
├──────────────────────────────────────────────┤
│                                              │
│              HERO SECTION                   │
│                                              │
│              Hi, I'm [Name]                │
│              Full-Stack Developer           │
│                                              │
│      Building modern web applications       │
│      with React and Go.                     │
│                                              │
│       [View Projects] [Contact Me]           │
│                                              │
├──────────────────────────────────────────────┤
│              TECH STACK                     │
│                                              │
│       React   JavaScript   Go   SQL         │
│                                              │
├──────────────────────────────────────────────┤
│              FEATURED PROJECTS              │
│                                              │
│       Project 1   Project 2   Project 3     │
│                                              │
├──────────────────────────────────────────────┤
│              SHORT INTRO                    │
│                                              │
│        Brief developer introduction          │
│                                              │
└──────────────────────────────────────────────┘
Functionality

The Home page should:

Navigate to Projects when "View Projects" is selected.
Navigate to Contact when "Contact Me" is selected.
Display selected featured projects.
Provide links to professional external profiles where applicable.
7. About Page
Purpose

The About page explains the developer's background and development journey.

UI Structure
┌──────────────────────────────────────────────┐
│                 ABOUT                        │
│                                              │
│     [Profile / Developer Image]              │
│                                              │
│     About Me                                 │
│                                              │
│     Introduction                             │
│                                              │
│     Education                                │
│                                              │
│     Development Journey                      │
│                                              │
│     What I Build                             │
│                                              │
├──────────────────────────────────────────────┤
│              EXPERIENCE / JOURNEY            │
│                                              │
│       Education → Projects → Development     │
│                                              │
└──────────────────────────────────────────────┘
Functionality

The About page will primarily contain static information.

Information should be stored as frontend content unless there is a future requirement for an admin dashboard.

8. Projects Page
Purpose

The Projects page is the main portfolio showcase.

It should allow visitors to understand what has been built, which technologies were used, and what problem each project solves.

UI Structure
┌──────────────────────────────────────────────┐
│                 PROJECTS                     │
│                                              │
│  All   Web   Backend   Full Stack             │
│                                              │
├──────────────────────────────────────────────┤
│                                              │
│  ┌────────────────┐  ┌────────────────┐      │
│  │ Project Image  │  │ Project Image  │      │
│  │                │  │                │      │
│  │ Project Name   │  │ Project Name   │      │
│  │ Description    │  │ Description    │      │
│  │ React | Go     │  │ React | Go     │      │
│  │                │  │                │      │
│  │ [View Project] │  │ [View Project] │      │
│  └────────────────┘  └────────────────┘      │
│                                              │
│  ┌────────────────┐  ┌────────────────┐      │
│  │ Project Image  │  │ Project Image  │      │
│  │                │  │                │      │
│  │ Project Name   │  │ Project Name   │      │
│  │ Description    │  │ Description    │      │
│  │ Technologies   │  │ Technologies   │      │
│  └────────────────┘  └────────────────┘      │
│                                              │
└──────────────────────────────────────────────┘
Project Data

Projects will eventually be retrieved from the Go backend.

Example:

{
  "id": 1,
  "title": "Safeguarding Reporting Platform",
  "description": "Anonymous reporting platform...",
  "technologies": ["React", "Go", "SQLite"],
  "image": "/projects/project-1.png",
  "githubUrl": "...",
  "liveUrl": "...",
  "category": "Full Stack"
}
API

The frontend will request projects through:

GET /api/projects

A single project can be retrieved using:

GET /api/projects/:id
9. Skills Page
Purpose

The Skills page presents the technologies and technical areas the developer works with.

UI Structure
┌──────────────────────────────────────────────┐
│                 SKILLS                       │
│                                              │
│  Frontend                                    │
│  ┌────────────────────────────────────────┐  │
│  │ HTML    CSS    JavaScript    React     │  │
│  └────────────────────────────────────────┘  │
│                                              │
│  Backend                                     │
│  ┌────────────────────────────────────────┐  │
│  │ Go      REST APIs      SQL              │  │
│  └────────────────────────────────────────┘  │
│                                              │
│  Tools                                       │
│  ┌────────────────────────────────────────┐  │
│  │ Git    Linux    GitHub    VS Code      │  │
│  └────────────────────────────────────────┘  │
│                                              │
│  Other                                       │
│  ┌────────────────────────────────────────┐  │
│  │ Networking    Debugging    Problem      │  │
│  │ Solving                                   │  │
│  └────────────────────────────────────────┘  │
└──────────────────────────────────────────────┘

Skills should be grouped rather than displayed as one large list.

10. Contact Page
Purpose

The Contact page allows visitors to send a message.

UI Structure
┌──────────────────────────────────────────────┐
│                 CONTACT                      │
│                                              │
│      Let's work together                    │
│                                              │
│  Name                                        │
│  ┌────────────────────────────────────────┐  │
│  │ Enter your name                        │  │
│  └────────────────────────────────────────┘  │
│                                              │
│  Email                                       │
│  ┌────────────────────────────────────────┐  │
│  │ Enter your email                       │  │
│  └────────────────────────────────────────┘  │
│                                              │
│  Subject                                     │
│  ┌────────────────────────────────────────┐  │
│  │ Subject                                │  │
│  └────────────────────────────────────────┘  │
│                                              │
│  Message                                     │
│  ┌────────────────────────────────────────┐  │
│  │ Write your message                     │  │
│  │                                        │  │
│  └────────────────────────────────────────┘  │
│                                              │
│              [ Send Message ]                │
│                                              │
│  GitHub | LinkedIn | Email                   │
└──────────────────────────────────────────────┘
Functionality

When the visitor submits the form:

React Form
     ↓
Client-side validation
     ↓
POST /api/contact
     ↓
Go API
     ↓
Validate request
     ↓
Store message in SQLite
     ↓
Return JSON response
     ↓
React displays success/error state

Example endpoint:

POST /api/contact

Example request:

{
  "name": "John Doe",
  "email": "john@example.com",
  "subject": "Project Inquiry",
  "message": "I would like to work with you."
}
11. Backend Architecture

The Go backend will follow a layered architecture.

HTTP Request
     │
     ▼
Router
     │
     ▼
Handler
     │
     ▼
Service
     │
     ▼
Repository
     │
     ▼
SQLite
Router

Responsible for mapping URLs to handlers.

Example:

GET  /api/projects
GET  /api/projects/:id
GET  /api/skills
POST /api/contact
Handlers

Handlers are responsible for:

Receiving HTTP requests
Reading request data
Validating basic request structure
Calling services
Returning HTTP responses

Handlers should not contain database logic.

Services

Services contain application/business logic.

Example:

ProjectService
ContactService
SkillService
Repository

Repositories handle database operations.

Example:

ProjectRepository
ContactRepository
SkillRepository

This keeps database logic separate from HTTP logic.

12. Backend Folder Structure
backend/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   │
│   ├── handlers/
│   │   ├── project_handler.go
│   │   ├── skill_handler.go
│   │   └── contact_handler.go
│   │
│   ├── services/
│   │   ├── project_service.go
│   │   ├── skill_service.go
│   │   └── contact_service.go
│   │
│   ├── repositories/
│   │   ├── project_repository.go
│   │   ├── skill_repository.go
│   │   └── contact_repository.go
│   │
│   ├── models/
│   │   ├── project.go
│   │   ├── skill.go
│   │   └── contact.go
│   │
│   └── database/
│       └── database.go
│
├── migrations/
│
├── go.mod
└── go.sum
13. Frontend Folder Structure
frontend/
│
├── public/
│   ├── images/
│   └── projects/
│
├── src/
│   │
│   ├── assets/
│   │
│   ├── components/
│   │   ├── Navbar.jsx
│   │   ├── Footer.jsx
│   │   ├── Button.jsx
│   │   ├── ProjectCard.jsx
│   │   ├── SkillCard.jsx
│   │   └── Loading.jsx
│   │
│   ├── pages/
│   │   ├── Home.jsx
│   │   ├── About.jsx
│   │   ├── Projects.jsx
│   │   ├── Skills.jsx
│   │   └── Contact.jsx
│   │
│   ├── services/
│   │   ├── projectService.js
│   │   ├── skillService.js
│   │   └── contactService.js
│   │
│   ├── hooks/
│   │
│   ├── utils/
│   │
│   ├── App.jsx
│   ├── main.jsx
│   └── index.css
│
├── package.json
└── vite.config.js
14. Database Architecture

SQLite will initially contain three main tables.

projects
skills
contacts
Projects
projects
──────────────
id
title
description
category
image
github_url
live_url
created_at
Skills
skills
──────────────
id
name
category
level
created_at
Contacts
contacts
──────────────
id
name
email
subject
message
created_at

The database should only be accessed by the Go backend.

The React application must never connect directly to SQLite.

15. API Architecture

The API will use REST principles.

Projects
GET /api/projects

Returns all projects.

GET /api/projects/:id

Returns one project.

Skills
GET /api/skills

Returns skills.

Contact
POST /api/contact

Creates a contact message.

16. API Response Structure

Successful responses should use JSON.

Example:

{
  "success": true,
  "data": []
}

Errors should follow a consistent structure.

{
  "success": false,
  "error": "Unable to process request"
}

The frontend should use the response status and JSON body to determine whether to display success, loading, or error states.

17. Frontend ↔ Backend Communication

The communication flow should be:

React Component
      │
      ▼
Service Function
      │
      ▼
Fetch API
      │
      ▼
Go REST API
      │
      ▼
Service Layer
      │
      ▼
Repository
      │
      ▼
SQLite

For example:

Projects.jsx
     ↓
projectService.js
     ↓
GET /api/projects
     ↓
Go ProjectHandler
     ↓
ProjectService
     ↓
ProjectRepository
     ↓
SQLite
18. UI State Architecture

Dynamic pages must account for three primary states.

Loading
Projects
   ↓
Loading projects...
Success
Projects
   ↓
Project cards
Error
Projects
   ↓
Unable to load projects.
[Try Again]

The Contact page should additionally have:

Idle
  ↓
Submitting
  ↓
Success / Error
19. Navigation Architecture

The navigation bar will remain consistent across the application.

Home
About
Projects
Skills
Contact

Desktop:

[NAME]       Home  About  Projects  Skills  Contact

Mobile:

[NAME]                                      [MENU]

The mobile menu should expand into a vertical navigation menu.

Navigation should be handled through React Router rather than manually reloading pages.

20. Responsive Architecture

The website must work across:

Mobile
   ↓
Tablet
   ↓
Desktop

Primary target sizes:

375px
768px
1024px
1440px+

The layout should adapt rather than simply shrinking the desktop interface.

For example:

Desktop

[ Project ] [ Project ] [ Project ]


Tablet

[ Project ] [ Project ]


Mobile

[ Project ]
[ Project ]
[ Project ]
21. Component Architecture

Components should be reusable.

For example:

ProjectCard

should receive project data rather than containing a specific project's information.

<ProjectCard project={project} />

Similarly:

<SkillCard skill={skill} />

The same component should be reusable throughout the application.

22. Architectural Rules

The following rules must be maintained throughout development.

Frontend
UI logic belongs in React components.
API communication belongs in service files.
Reusable UI belongs in components/.
Pages belong in pages/.
Do not place database logic in React components.
Do not hardcode API responses into components when the data is supposed to come from the backend.
Backend
HTTP handling belongs in handlers.
Business logic belongs in services.
Database operations belong in repositories.
Database models should be separated from HTTP handling.
Handlers should not directly contain SQL queries.
Database
SQLite is only accessed by Go.
Database queries must be parameterized.
Database schema changes should be tracked through migrations.
23. Security Architecture

Even though this is initially a portfolio website, basic security should be included.

The backend should:

Validate contact form input.
Validate email format.
Limit excessive contact submissions.
Sanitize or safely handle stored user input.
Use parameterized SQL queries.
Never expose database credentials or secrets.
Never expose private backend configuration to React.
Configure CORS appropriately.
Return safe error messages to clients.

Environment-specific configuration should use environment variables.

Example:

.env

PORT=8080
DATABASE_URL=./portfolio.db
FRONTEND_URL=http://localhost:5173

Secrets must never be committed to Git.

24. Development Environment

During development:

React + Vite
http://localhost:5173

Go API:

http://localhost:8080

Communication:

React
localhost:5173
      │
      │ HTTP
      ▼
Go API
localhost:8080

The Go server should be configured to allow requests from the frontend development origin.

25. Production Architecture

The final deployment can follow this structure:

                 INTERNET
                     │
                     ▼
              ┌──────────────┐
              │   Frontend   │
              │ React + Vite │
              └──────┬───────┘
                     │
                     │ HTTPS
                     ▼
              ┌──────────────┐
              │  Go Backend  │
              │   REST API   │
              └──────┬───────┘
                     │
                     ▼
              ┌──────────────┐
              │    SQLite    │
              └──────────────┘

The exact hosting provider will be decided separately and should not be hardcoded into the application architecture.

26. Overall User Journey

A typical visitor experience should be:

                VISITOR
                   │
                   ▼
                 HOME
                   │
        ┌──────────┼──────────┐
        ▼          ▼          ▼
      ABOUT     PROJECTS    SKILLS
        │          │          │
        └──────────┼──────────┘
                   ▼
                CONTACT
                   │
                   ▼
             Send Message
                   │
                   ▼
               Go API
                   │
                   ▼
                SQLite

The visitor should be able to understand the developer and navigate the entire portfolio without unnecessary interactions.

27. Final Architecture Principle

The project should follow a clear separation of responsibilities:

React
─────
Presentation
UI
Navigation
User Interaction
Frontend State

        ↓

Go API
──────
HTTP
Validation
Business Logic
API Responses

        ↓

SQLite
──────
Persistent Data

The central architectural principle is:

The frontend should know how to present and request data. The backend should know how to process and manage data. The database should only be accessed by the backend.
