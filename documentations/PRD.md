# PRD.md

# Personal Developer Portfolio

## 1. Product Overview

### Product Name

Valentine Omondi Awili — Personal Developer Portfolio

### Product Type

Personal professional portfolio website.

### Purpose

The portfolio is designed to present Valentine Omondi Awili as a professional Full-Stack Developer through a visually strong, editorial-style website.

The website will communicate:

- Who Valentine is
- What he does
- His development journey
- His technical skills
- His projects
- His education and experience
- How potential clients, employers, collaborators, or recruiters can contact him

The website should function as both a professional introduction and a visual representation of his development capabilities.

---

# 2. Product Vision

Create a premium personal portfolio that feels more like a digital editorial profile than a traditional developer portfolio.

The experience should combine:

- Personal storytelling
- Strong typography
- Large imagery
- Project storytelling
- Minimal navigation
- Editorial layouts
- Smooth interactions
- Responsive design
- Professional technical presentation

The final website should feel intentional, modern, personal, and original.

---

# 3. Reference Design

The visual direction is strongly inspired by:

https://syedameen-elonmuskporfolio.netlify.app/

The reference should influence:

- Page composition
- Visual hierarchy
- Typography scale
- Large imagery
- Section spacing
- Editorial storytelling
- Timeline presentation
- Long-scroll experience
- Navigation structure
- Footer treatment

The implementation must remain original.

The project must NOT copy:

- Source code
- Text
- Images
- Proprietary assets
- Celebrity content
- Branding
- Personal information from the reference site

The reference is used only as a design and layout inspiration.

---

# 4. Target Audience

The primary audiences are:

### Employers

Recruiters and companies looking for developers.

They should quickly understand:

- Valentine's technical background
- Technologies he works with
- Projects he has built
- His development interests
- How to contact him

### Potential Clients

Individuals, businesses, startups, NGOs, and organizations looking for web development services.

They should be able to understand:

- What Valentine can build
- His technical capabilities
- Previous project experience
- How to contact him

### Collaborators

Developers, designers, organizations, and technical communities interested in collaboration.

### Professional Network

People discovering Valentine through:

- GitHub
- LinkedIn
- Professional communities
- Search engines
- Direct referrals

---

# 5. Product Goals

## Primary Goals

### Goal 1 — Professional Identity

Clearly establish Valentine Omondi Awili as a Full-Stack Developer.

### Goal 2 — Showcase Projects

Present real projects through detailed, visually engaging storytelling.

### Goal 3 — Demonstrate Technical Skills

Show relevant frontend, backend, database, development, and tooling skills.

### Goal 4 — Tell a Personal Story

Explain Valentine's development journey, education, interests, and professional direction.

### Goal 5 — Generate Opportunities

Provide clear ways for employers, clients, and collaborators to contact Valentine.

### Goal 6 — Demonstrate Technical Ability

The portfolio itself should demonstrate practical knowledge of:

- React
- JavaScript
- Go
- REST APIs
- SQLite
- Git
- Responsive web development
- API integration
- Backend architecture

---

# 6. Non-Goals

The portfolio will NOT attempt to become:

- A SaaS dashboard
- A social network
- A blogging platform
- A news website
- A celebrity biography
- An AI chatbot platform
- A project management application
- A financial dashboard

The following reference-site features are explicitly excluded:

- Net Worth
- Articles
- Blog
- Celebrity content
- Celebrity quotes
- Celebrity/company showcase

---

# 7. Technology Requirements

## Frontend

The frontend will use:

- React
- Vite
- JavaScript
- CSS

React will be responsible for:

- User interface
- Navigation
- Section rendering
- Form interaction
- API communication
- UI state
- Animations

---

# 8. Backend

The backend will use:

- Go
- REST API

Go will handle:

- HTTP requests
- API routing
- Request validation
- Business logic
- Database communication
- Contact form submissions
- JSON responses

---

# 9. Database

The database will use:

- SQLite

SQLite will store:

- Projects
- Skills
- Contact messages

The frontend must never communicate directly with SQLite.

The architecture must be:

    React
       ↓
    REST API
       ↓
    Go
       ↓
    SQLite

---

# 10. Application Architecture

High-level architecture:

    User
      ↓
    React + Vite
      ↓
    HTTP / JSON
      ↓
    Go REST API
      ↓
    Services
      ↓
    Repositories
      ↓
    SQLite

Responsibilities:

### React

Responsible for presentation and user interaction.

### API Handlers

Responsible for HTTP requests and responses.

### Services

Responsible for business logic and validation.

### Repositories

Responsible for database operations.

### SQLite

Responsible for persistence.

---

# 11. Website Structure

The website will be a long-form single-page portfolio.

Primary sections:

1. Home
2. About
3. Projects
4. Skills
5. Contact
6. Footer

Navigation should allow users to move between these sections.

---

# 12. Navigation Requirements

## Desktop

Navigation should contain:

    VALENTINE

    HOME
    ABOUT
    PROJECTS
    SKILLS
    CONTACT

The navigation should remain visually minimal.

## Mobile

The navigation should transform into:

    VALENTINE                         MENU

The mobile menu should provide access to all primary sections.

Requirements:

- Keyboard accessible
- Responsive
- Smooth transitions
- Clear active state
- No unnecessary UI elements

---

# 13. Home / Hero Section

The hero is the first major visual experience.

It should immediately communicate:

- Identity
- Professional role
- Personality
- Development focus

### Content

The hero should contain:

- "I'm"
- VALENTINE OMONDI AWILI
- Full-Stack Developer
- Short introduction
- Profile image
- Projects CTA
- Contact CTA
- Social links

### Design

The name should be one of the largest elements on the page.

The layout should use:

- Large typography
- Generous whitespace
- Strong image composition
- Minimal controls
- Subtle animation

The hero must not look like a generic SaaS landing page.

---

# 14. About Section

The About section should provide a deeper personal introduction.

### Required Content

- Personal biography
- Education
- Development journey
- Current professional focus
- What I Do
- Large personal image

### Purpose

The section should answer:

- Who is Valentine?
- How did he get into technology?
- What is he currently learning/building?
- What type of developer is he becoming?
- What does he enjoy building?

The section should feel editorial rather than like a resume table.

---

# 15. Projects Section

Projects are one of the most important sections of the portfolio.

Projects should demonstrate real development experience.

## Presentation

Projects must NOT be displayed as a generic grid of dashboard cards.

Instead, use an editorial/timeline storytelling layout.

Each project should have:

- Year
- Project name
- Category
- Description
- Technologies
- Project image
- GitHub link
- Live project link when available

Projects may alternate between:

- Image left / text right
- Text left / image right

This creates visual rhythm throughout the page.

---

# 16. Initial Project Portfolio

Potential projects include:

### Safeguarding Reporting Platform

A secure reporting platform designed to allow users to submit safeguarding reports.

Possible technologies:

- React
- JavaScript
- Go
- REST APIs
- SQLite
- MongoDB
- Authentication
- Security mechanisms

### Career Guidance Platform

A platform that helps students explore career paths based on academic performance and interests.

Possible technologies:

- HTML
- CSS
- JavaScript
- React
- Backend APIs
- Database

### HopeReach NGO Website

A responsive website for an NGO.

Possible technologies:

- HTML
- CSS
- JavaScript
- Responsive Web Design

### Personal Portfolio

The current portfolio itself.

Technologies:

- React
- Vite
- JavaScript
- CSS
- Go
- REST API
- SQLite

Only projects that accurately represent actual experience should be published.

---

# 17. Skills Section

The Skills section should communicate technical capabilities without using artificial rankings.

## Frontend

- HTML
- CSS
- JavaScript
- React
- Vite

## Backend

- Go
- REST APIs
- SQLite
- MongoDB

## Tools

- Git
- GitHub
- Linux
- VS Code

## Other Technical Skills

- Networking
- Debugging
- Responsive Web Design
- Problem Solving
- Analytical Thinking
- IFMIS / Government Systems

---

# 18. Skills Presentation Rules

Do NOT use:

- Fake percentages
- Progress bars
- Skill ratings
- Stars
- Arbitrary proficiency scores

Skills should be presented through typography, grouping, and visual hierarchy.

The purpose is to communicate experience rather than create artificial rankings.

---

# 19. Contact Section

The Contact section should provide a simple way for visitors to reach Valentine.

### Main Heading

Suggested:

    LET'S BUILD SOMETHING TOGETHER.

### Form Fields

- Name
- Email
- Subject
- Message

### Additional Contact Methods

- Email
- GitHub
- LinkedIn

---

# 20. Contact Form Behavior

User submits:

    Name
    Email
    Subject
    Message

Frontend validates the data.

Then:

    POST /api/contact

The Go backend:

1. Receives the request.
2. Validates the request.
3. Sanitizes/validates input.
4. Stores the message in SQLite.
5. Returns JSON.

The frontend then displays:

- Success state
- Validation errors
- Server errors

---

# 21. API Requirements

## Health

    GET /api/health

Response:

    {
      "success": true,
      "message": "API is running"
    }

## Projects

    GET /api/projects

Returns all projects.

## Project

    GET /api/projects/:id

Returns one project.

## Skills

    GET /api/skills

Returns available skills.

## Contact

    POST /api/contact

Creates a contact message.

---

# 22. Database Requirements

## Projects Table

Fields:

- id
- title
- description
- category
- year
- technologies
- image
- github_url
- live_url
- created_at

## Skills Table

Fields:

- id
- name
- category
- created_at

## Contacts Table

Fields:

- id
- name
- email
- subject
- message
- created_at

---

# 23. Visual Design Requirements

The overall design should follow a premium editorial aesthetic.

### Typography

Preferred:

- Inter
- Manrope
- System fallback

Large headings should use fluid typography.

Example:

    clamp()

Typography should create a strong hierarchy between:

- Hero title
- Section titles
- Project titles
- Body text
- Metadata

---

# 24. Color System

Primary palette:

    White
    #FFFFFF

    Primary text
    #111111

    Secondary text
    #555555

    Muted text
    #888888

    Border
    #E5E5E5

    Accent
    #111111

Optional dark sections may be introduced when useful.

Avoid unnecessary gradients.

---

# 25. Buttons

Buttons should be:

- Minimal
- Professional
- Clear
- Editorial

Avoid excessive pill-shaped buttons.

Buttons should have:

- Clear hover states
- Keyboard focus states
- Good contrast
- Accessible labels

---

# 26. Animation Requirements

Animations should be subtle.

Possible animations:

- Text reveal
- Fade-in
- Slide-in
- Image reveal
- Hover movement
- Link underline transitions
- Section reveal

Animations must not distract from content.

The website must respect:

    prefers-reduced-motion

When reduced motion is enabled, unnecessary animations should be disabled or reduced.

---

# 27. Responsive Requirements

The website must support:

## Mobile

- 375px
- 390px
- 414px

## Tablet

- 768px
- 820px
- 912px

## Desktop

- 1024px
- 1280px
- 1366px
- 1440px

There must be no horizontal scrolling.

All major sections must remain readable and visually balanced at different screen sizes.

---

# 28. Accessibility Requirements

The website must include:

- Semantic HTML
- Proper heading hierarchy
- Alt text
- Form labels
- Keyboard navigation
- Visible focus states
- Accessible navigation
- Accessible mobile menu
- Sufficient color contrast
- Reduced-motion support

Interactive elements must be usable without a mouse.

---

# 29. Performance Requirements

The website should:

- Optimize images
- Lazy-load non-critical images
- Minimize unnecessary JavaScript
- Avoid unnecessary dependencies
- Avoid unnecessary React re-renders
- Produce a successful production build

The backend should:

- Use efficient SQL queries
- Use appropriate indexes
- Avoid unnecessary database calls

---

# 30. Security Requirements

The backend must:

- Validate all incoming data
- Use parameterized SQL queries
- Configure CORS
- Validate content types
- Limit contact submissions
- Avoid exposing internal server errors
- Keep secrets outside source code

Environment variables should be used for configuration.

Example:

    PORT=8080
    DATABASE_URL=./portfolio.db
    FRONTEND_URL=http://localhost:5173

A `.env.example` file must be provided.

---

# 31. Frontend Folder Structure

Expected structure:

    src/
    ├── components/
    ├── pages/
    ├── services/
    ├── hooks/
    ├── utils/
    ├── assets/
    ├── App.jsx
    ├── main.jsx
    └── index.css

---

# 32. Backend Folder Structure

Expected structure:

    backend/
    ├── cmd/
    │   └── server/
    │       └── main.go
    │
    ├── internal/
    │   ├── handlers/
    │   ├── services/
    │   ├── repositories/
    │   ├── models/
    │   ├── database/
    │   └── migrations/
    │
    └── go.mod

---

# 33. User Experience Requirements

Visitors should be able to understand Valentine's professional identity within the first few seconds.

The experience should follow this progression:

    Identity
       ↓
    Introduction
       ↓
    Personal Story
       ↓
    Projects
       ↓
    Skills
       ↓
    Contact

The website should feel like a continuous story rather than a collection of disconnected pages.

---

# 34. Content Principles

All portfolio content should be:

- Accurate
- Professional
- Personal
- Concise
- Easy to understand
- Relevant to development

Do not exaggerate experience.

Only technologies and projects that accurately represent current experience should be presented.

---

# 35. SEO Requirements

The website should include:

- Descriptive page title
- Meta description
- Semantic HTML
- Descriptive image alt text
- Open Graph metadata where appropriate
- Descriptive links
- Clean content hierarchy

Suggested title:

    Valentine Omondi Awili | Full-Stack Developer

---

# 36. Browser Support

The website should be tested in:

- Google Chrome
- Mozilla Firefox
- Microsoft Edge

Testing should include:

- Desktop
- Tablet
- Mobile

---

# 37. Testing Requirements

## Frontend

Test:

- Navigation
- Mobile menu
- Hero
- About
- Projects
- Skills
- Contact
- Loading states
- Error states
- Links
- Responsive layouts

## Backend

Test:

    GET /api/health
    GET /api/projects
    GET /api/projects/:id
    GET /api/skills
    POST /api/contact

Test:

- Valid requests
- Invalid requests
- Missing fields
- Invalid IDs
- Database errors
- HTTP status codes

---

# 38. Success Criteria

The project will be considered successful when:

- The portfolio clearly communicates Valentine's identity.
- The design follows the editorial direction.
- The website is responsive.
- Projects are presented professionally.
- Skills are presented accurately.
- The contact form works.
- The frontend communicates correctly with the Go API.
- The Go API communicates correctly with SQLite.
- The application handles errors appropriately.
- The website is accessible.
- The website has no major console errors.
- The production build succeeds.
- No secrets are exposed.
- No unwanted reference-site features are included.

---

# 39. Explicitly Excluded Features

The following must NOT be added unless the project requirements are intentionally changed:

- Net Worth section
- Blog
- Articles
- Celebrity biography
- Celebrity quotes
- Celebrity/company showcase
- Fake testimonials
- Fake client logos
- Fake statistics
- Fake skill percentages
- Generic AI chatbot
- SaaS dashboard
- Unnecessary authentication
- Admin dashboard unless separately specified

---

# 40. Development Workflow

Development should follow the task sequence defined in:

    TASKS.md

Technical decisions should follow:

    ARCHITECTURE.md

Visual decisions should follow:

    DESIGN.md

Before implementing a task:

1. Identify the task.
2. Read the relevant documentation.
3. Inspect the current implementation.
4. Plan the smallest required change.
5. Implement the task.
6. Test the task.
7. Review the result.
8. Mark the task as complete.
9. Move to the next task.

Do not implement unrelated tasks together.

---

# 41. Current Project Status

Status:

    NOT STARTED

Current Phase:

    Phase 1 — Project Foundation

Current Task:

    TASK-001 — Create Project Repository

Next Task:

    TASK-001 — Create Project Repository

---

# 42. Product Definition

The final product is a professional personal developer portfolio that combines:

    Personal Brand
          +
    Editorial Storytelling
          +
    Project Showcase
          +
    Technical Skills
          +
    Contact Experience

The website should communicate not only what Valentine can code, but also how he approaches technology, solves problems, and builds useful digital products.
