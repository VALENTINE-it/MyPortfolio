# TASKS.md

# Personal Developer Portfolio

This document breaks the portfolio development into small, manageable
tasks.

The project uses:

- React
- Vite
- JavaScript
- CSS
- Go
- REST API
- SQLite

The visual design must follow `DESIGN.md`, and the technical structure
must follow `ARCHITECTURE.md`.

---

# Development Workflow

Every task should follow this process:

1. Read the relevant documentation.
2. Inspect the existing implementation.
3. Create a small implementation plan.
4. Implement only the requested task.
5. Test the implementation.
6. Fix any issues.
7. Review the result against `DESIGN.md`.
8. Mark the task as complete.
9. Commit the changes.

Do not implement multiple unrelated tasks at once.

---

# PHASE 1 — PROJECT FOUNDATION

## TASK-001 — Create Project Repository

- [x] Create the project root directory.
- [x] Initialize Git.
- [x] Create the basic documentation structure.
- [x] Add `.gitignore`.
- [x] Add `README.md`.

Expected structure:

    portfolio/
    ├── docs/
    │   ├── PRD.md
    │   ├── ARCHITECTURE.md
    │   ├── DESIGN.md
    │   └── TASKS.md
    │
    ├── frontend/
    ├── backend/
    ├── README.md
    └── .gitignore

---

## TASK-002 — Initialize React + Vite

- [x] Create React + Vite frontend.
- [x] Configure JavaScript.
- [x] Verify development server.
- [x] Remove default Vite content.
- [x] Create initial application structure.

Expected result:

    npm run dev

starts the frontend successfully.

---

## TASK-003 — Initialize Go Backend

- [x] Create Go module.
- [x] Create backend entry point.
- [x] Create HTTP server.
- [x] Configure server port.
- [x] Add initial health endpoint.

Expected endpoint:

    GET /api/health

Expected response:

    {
      "success": true,
      "message": "API is running"
    }

---

# PHASE 2 — FRONTEND ARCHITECTURE

## TASK-004 — Create Frontend Folder Structure

- [x] Create `components/`.
- [x] Create `pages/`.
- [x] Create `services/`.
- [x] Create `hooks/`.
- [x] Create `utils/`.
- [x] Create `assets/`.

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

## TASK-005 — Configure Global Styling

- [x] Create global CSS.
- [x] Configure typography.
- [x] Configure CSS variables.
- [x] Configure base spacing.
- [x] Configure responsive breakpoints.
- [x] Configure global transitions.
- [x] Remove default browser margins.

Use the design system defined in `DESIGN.md`.

---

# PHASE 3 — NAVIGATION

## TASK-006 — Build Navbar

- [x] Create `Navbar.jsx`.
- [x] Add developer name/logo.
- [x] Add Home navigation.
- [x] Add About navigation.
- [x] Add Projects navigation.
- [x] Add Skills navigation.
- [x] Add Contact navigation.
- [x] Add mobile menu.
- [x] Add active section state.
- [x] Add smooth scrolling.

Desktop:

    VALENTINE       HOME ABOUT PROJECTS SKILLS CONTACT

Mobile:

    VALENTINE                         MENU

---

## TASK-007 — Build Footer

- [x] Create `Footer.jsx`.
- [x] Add developer name.
- [x] Add professional title.
- [x] Add GitHub link.
- [x] Add LinkedIn link.
- [x] Add email.
- [x] Add copyright.
- [x] Make footer responsive.

---

# PHASE 4 — HOME / HERO

## TASK-008 — Build Hero Section

- [x] Create `Hero.jsx`.
- [x] Add "I'm".
- [x] Add developer name.
- [x] Add professional title.
- [x] Add short introduction.
- [x] Add profile image.
- [x] Add Projects CTA.
- [x] Add Contact CTA.
- [x] Add social links.

The developer name should be one of the largest elements
on the page.

---

## TASK-009 — Hero Animations

- [x] Add initial text reveal.
- [x] Add image reveal.
- [x] Add subtle movement.
- [x] Respect reduced-motion preferences.
- [x] Verify animation performance.

Animations must remain subtle.

---

# PHASE 5 — ABOUT

## TASK-010 — Build About Section

- [x] Create `About.jsx`.
- [x] Add About heading.
- [x] Add large personal image.
- [x] Add personal introduction.
- [x] Add education information.
- [x] Add development journey.
- [x] Add current professional focus.
- [x] Add "What I Do" content.

The layout should follow the editorial style defined in
`DESIGN.md`.

---

## TASK-011 — About Responsive Layout

- [x] Create desktop layout.
- [x] Create tablet layout.
- [x] Create mobile layout.
- [x] Test image positioning.
- [x] Test typography.
- [x] Test spacing.

---

# PHASE 6 — DATABASE

## TASK-012 — Configure SQLite

- [x] Add SQLite dependency.
- [x] Create database connection.
- [x] Configure database path.
- [x] Create database initialization.
- [x] Add error handling.

The database must only be accessed by the Go backend.

---

## TASK-013 — Create Database Schema

Create:

    projects
    skills
    contacts

Projects:

    id
    title
    description
    category
    year
    technologies
    image
    github_url
    live_url
    created_at

Skills:

    id
    name
    category
    created_at

Contacts:

    id
    name
    email
    subject
    message
    created_at

- [x] Create tables.
- [x] Add appropriate indexes.
- [x] Add timestamps.
- [x] Test database creation.

---

# PHASE 7 — GO BACKEND

## TASK-014 — Create Backend Models

- [x] Create `Project` model.
- [x] Create `Skill` model.
- [x] Create `Contact` model.
- [x] Add JSON serialization.
- [x] Verify model structure.

---

## TASK-015 — Create Project Repository

- [x] Create project repository.
- [x] Add `GetAllProjects`.
- [x] Add `GetProjectByID`.
- [x] Add database queries.
- [x] Use parameterized queries.
- [x] Handle database errors.

---

## TASK-016 — Create Skill Repository

- [x] Create skill repository.
- [x] Add `GetAllSkills`.
- [x] Add category filtering if required.
- [x] Handle database errors.

---

## TASK-017 — Create Contact Repository

- [x] Create contact repository.
- [x] Add message creation.
- [x] Validate database operations.
- [x] Handle database errors.

---

# PHASE 8 — BACKEND SERVICES

## TASK-018 — Create Project Service

- [x] Create `ProjectService`.
- [x] Retrieve projects.
- [x] Validate project data.
- [x] Connect service to repository.

---

## TASK-019 — Create Skill Service

- [x] Create `SkillService`.
- [x] Retrieve skills.
- [x] Group skills by category if necessary.

---

## TASK-020 — Create Contact Service

- [x] Create `ContactService`.
- [x] Validate contact data.
- [x] Validate email.
- [x] Reject invalid requests.
- [x] Save valid messages.

---

# PHASE 9 — REST API

## TASK-021 — Projects API

Implement:

    GET /api/projects

and:

    GET /api/projects/:id

- [x] Create handlers.
- [x] Create routes.
- [x] Return JSON.
- [x] Handle 404.
- [x] Handle database errors.
- [x] Test endpoints.

---

## TASK-022 — Skills API

Implement:

    GET /api/skills

- [x] Create handler.
- [x] Create route.
- [x] Return JSON.
- [x] Test endpoint.

---

## TASK-023 — Contact API

Implement:

    POST /api/contact

- [x] Create handler.
- [x] Parse JSON.
- [x] Validate request.
- [x] Call service.
- [x] Store message.
- [x] Return success response.
- [x] Return validation errors.

---

# PHASE 10 — PROJECTS UI

## TASK-024 — Build Project Timeline

- [x] Create `ProjectTimeline.jsx`.
- [x] Create `ProjectItem.jsx`.
- [x] Add year.
- [x] Add project title.
- [x] Add category.
- [x] Add description.
- [x] Add technologies.
- [x] Add project image.
- [x] Add GitHub link.
- [x] Add live project link.

Projects must NOT use a generic dashboard-style card grid.

The layout should follow the editorial/timeline design.

---

## TASK-025 — Connect Projects to API

Create:

    services/projectService.js

- [x] Fetch projects from Go API.
- [x] Display loading state.
- [x] Display projects.
- [x] Display error state.
- [x] Add retry functionality.

Data flow:

    Projects.jsx
         ↓
    projectService.js
         ↓
    GET /api/projects
         ↓
    Go API
         ↓
    SQLite

---

## TASK-026 — Project Animations

- [x] Add scroll reveal.
- [x] Add image reveal.
- [x] Add hover effects.
- [x] Add subtle link animations.
- [x] Test reduced-motion behavior.

---

# PHASE 11 — SKILLS UI

## TASK-027 — Build Skills Section

- [x] Create `Skills.jsx`.
- [x] Create skill categories.
- [x] Display frontend skills.
- [x] Display backend skills.
- [x] Display tools.
- [x] Display other technical skills.

Do not use percentage-based skill bars.

---

## TASK-028 — Connect Skills to API

Create:

    services/skillService.js

- [x] Fetch skills.
- [x] Display loading state.
- [x] Display skills.
- [x] Handle errors.
- [x] Group skills by category.

---

# PHASE 12 — CONTACT

## TASK-029 — Build Contact UI

- [x] Create `ContactForm.jsx`.
- [x] Add name input.
- [x] Add email input.
- [x] Add subject input.
- [x] Add message textarea.
- [x] Add submit button.
- [x] Add GitHub link.
- [x] Add LinkedIn link.
- [x] Add email.

The form should follow the editorial design.

Avoid dashboard-style form cards.

---

## TASK-030 — Contact Form Validation

Validate:

- [x] Name is required.
- [x] Email is required.
- [x] Email format is valid.
- [x] Subject is required.
- [x] Message is required.
- [x] Message has reasonable length.

Display useful validation messages.

---

## TASK-031 — Connect Contact Form to API

Create:

    services/contactService.js

Flow:

    ContactForm
        ↓
    Validation
        ↓
    POST /api/contact
        ↓
    Go Backend
        ↓
    SQLite
        ↓
    Response
        ↓
    Success/Error UI

---

# PHASE 13 — VISUAL REFINEMENT

## TASK-032 — Match Reference Layout

Compare the implementation with `DESIGN.md`.

Check:

- [x] Hero proportions.
- [x] Typography scale.
- [x] Section spacing.
- [x] Image sizes.
- [x] Project composition.
- [x] Timeline layout.
- [x] Navigation.
- [x] Footer.
- [x] Overall visual rhythm.

The website should maintain the editorial visual direction
throughout the entire page.

---

## TASK-033 — Typography Refinement

- [x] Configure primary font.
- [x] Configure hero typography.
- [x] Configure section headings.
- [x] Configure project headings.
- [x] Configure body text.
- [x] Configure metadata.
- [x] Verify mobile typography.

Use fluid typography with `clamp()` where appropriate.

---

## TASK-034 — Image Optimization

- [x] Compress large images.
- [x] Use appropriate formats.
- [x] Add image dimensions.
- [x] Add alt text.
- [x] Lazy-load non-critical images.
- [x] Verify mobile image performance.

---

# PHASE 14 — RESPONSIVE DESIGN

## TASK-035 — Desktop Testing

Test at:

    1440px
    1366px
    1280px
    1024px

Verify:

- [x] Navigation.
- [x] Hero.
- [x] About.
- [x] Projects.
- [x] Skills.
- [x] Contact.
- [x] Footer.

---

## TASK-036 — Tablet Testing

Test at:

    768px
    820px
    912px

Verify:

- [x] Navigation.
- [x] Typography.
- [x] Images.
- [x] Project layout.
- [x] Forms.
- [x] Spacing.

---

## TASK-037 — Mobile Testing

Test at:

    375px
    390px
    414px

Verify:

- [x] Mobile navigation.
- [x] Hero.
- [x] Profile image.
- [x] About.
- [x] Project timeline.
- [x] Skills.
- [x] Contact form.
- [x] Footer.

No horizontal scrolling should occur.

---

# PHASE 15 — ACCESSIBILITY

## TASK-038 — Accessibility Review

- [x] Use semantic HTML.
- [x] Add image alt text.
- [x] Add form labels.
- [x] Test keyboard navigation.
- [x] Add visible focus states.
- [x] Check color contrast.
- [x] Check heading hierarchy.
- [x] Test mobile menu with keyboard.
- [x] Add reduced-motion support.

---

# PHASE 16 — SECURITY

## TASK-039 — Backend Security

- [x] Configure CORS.
- [x] Validate all API inputs.
- [x] Use parameterized SQL queries.
- [x] Limit contact submissions.
- [x] Avoid exposing internal errors.
- [x] Validate request content type.
- [x] Keep secrets out of source code.

---

## TASK-040 — Environment Configuration

Create:

    .env.example

Example:

    PORT=8080
    DATABASE_URL=./portfolio.db
    FRONTEND_URL=http://localhost:5173

- [x] Create `.env.example`.
- [x] Add `.env` to `.gitignore`.
- [x] Load environment variables safely.
- [x] Verify production configuration.

---

# PHASE 17 — PERFORMANCE

## TASK-041 — Frontend Performance

- [x] Optimize images.
- [x] Lazy-load non-critical images.
- [x] Minimize unnecessary React re-renders.
- [x] Avoid unnecessary dependencies.
- [x] Check Vite production build.

---

## TASK-042 — Backend Performance

- [x] Check database queries.
- [x] Add appropriate indexes.
- [x] Avoid unnecessary database calls.
- [x] Verify API response times.

---

# PHASE 18 — TESTING

## TASK-043 — Frontend Functional Testing

Test:

- [x] Navigation works.
- [x] Mobile menu works.
- [x] Project data loads.
- [x] Skills data loads.
- [x] Contact form works.
- [x] Loading states work.
- [x] Error states work.
- [x] Links work.

---

## TASK-044 — Backend API Testing

Test:

    GET /api/health
    GET /api/projects
    GET /api/projects/:id
    GET /api/skills
    POST /api/contact

Verify:

- [x] Successful requests.
- [x] Invalid requests.
- [x] Missing data.
- [x] Invalid IDs.
- [x] Database errors.
- [x] Correct HTTP status codes.

---

## TASK-045 — End-to-End Testing

Verify the complete flows.

### Projects

    User
      ↓
    Projects section
      ↓
    API request
      ↓
    Go
      ↓
    SQLite
      ↓
    Project displayed

### Contact

    User
      ↓
    Contact form
      ↓
    Submit
      ↓
    Go API
      ↓
    SQLite
      ↓
    Success message

---

# PHASE 19 — CONTENT

## TASK-046 — Add Personal Information

- [x] Add developer name.
- [x] Add professional title.
- [x] Add biography.
- [x] Add education.
- [x] Add development journey.
- [x] Add professional links.
- [x] Add email.
- [x] Add profile image.

---

## TASK-047 — Add Projects

Add initial projects.

For every project:

- [x] Project name.
- [x] Description.
- [x] Year.
- [x] Category.
- [x] Technologies.
- [x] Project image.
- [x] GitHub URL.
- [x] Live URL if available.

---

## TASK-048 — Add Skills

Add:

### Frontend

- [x] HTML
- [x] CSS
- [x] JavaScript
- [x] React
- [x] Vite

### Backend

- [x] Go
- [x] REST APIs
- [x] SQLite
- [x] MongoDB

### Tools

- [x] Git
- [x] GitHub
- [x] Linux
- [x] VS Code

### Other

- [x] Networking
- [x] Debugging
- [x] Problem Solving

Only include technologies that accurately represent the developer's
current experience.

---

# PHASE 20 — FINAL REVIEW

## TASK-049 — Design Review

Compare the completed website against:

    DESIGN.md

Check:

- [x] Visual hierarchy.
- [x] Typography.
- [x] Spacing.
- [x] Images.
- [x] Project storytelling.
- [x] Navigation.
- [x] Responsive behavior.
- [x] Animations.
- [x] Contact experience.

---

## TASK-050 — Architecture Review

Compare implementation against:

    ARCHITECTURE.md

Verify:

- [x] React handles UI.
- [x] Services handle API communication.
- [x] Go handles API requests.
- [x] Services contain business logic.
- [x] Repositories handle database operations.
- [x] SQLite is only accessed by Go.
- [x] Components remain reusable.

---

## TASK-051 — Remove Unwanted Features

Confirm the website does NOT contain:

- [x] Net Worth
- [x] Articles
- [x] Blog
- [x] Celebrity content
- [x] Unrelated sections
- [x] Generic AI dashboard components
- [x] Fake skill percentages

---

# PHASE 21 — PRODUCTION PREPARATION

## TASK-052 — Production Build

Frontend:

    npm run build

- [x] Build succeeds.
- [x] No console errors.
- [x] No broken assets.
- [x] Verify production output.

Backend:

    go build

- [x] Build succeeds.
- [x] No compilation errors.
- [x] Environment configuration works.

---

## TASK-053 — Final Browser Testing

Test the production build in:

- [x] Chrome
- [x] Firefox
- [x] Edge

Check:

- [x] Desktop.
- [x] Tablet.
- [x] Mobile.
- [x] Navigation.
- [x] Images.
- [x] Forms.
- [x] API communication.

---

# PHASE 22 — DEPLOYMENT

## TASK-054 — Frontend Deployment

- [ ] Select hosting provider.
- [ ] Configure production environment.
- [ ] Build frontend.
- [ ] Deploy frontend.
- [ ] Configure API URL.
- [ ] Test deployed site.

---

## TASK-055 — Backend Deployment

- [ ] Select backend hosting provider.
- [ ] Configure environment variables.
- [ ] Deploy Go API.
- [ ] Configure SQLite/database storage.
- [ ] Configure CORS.
- [ ] Test production API.

---

## TASK-056 — Domain Configuration

Optional:

- [ ] Purchase/configure domain.
- [ ] Connect domain to frontend.
- [ ] Configure HTTPS.
- [ ] Verify API communication.
- [ ] Test all pages.

---

# PHASE 23 — FINAL RELEASE

## TASK-057 — Final Portfolio Review

- [x] All sections work.
- [x] All projects display correctly.
- [x] Contact form works.
- [x] Mobile layout works.
- [x] Desktop layout works.
- [x] No console errors.
- [x] No broken links.
- [x] No missing images.
- [x] No exposed secrets.
- [x] Accessibility reviewed.
- [x] Performance reviewed.

---

# COMPLETION CHECKLIST

## Frontend

- [x] React + Vite configured
- [x] Navbar completed
- [x] Hero completed
- [x] About completed
- [x] Projects completed
- [x] Skills completed
- [x] Contact completed
- [x] Footer completed
- [x] Responsive design completed
- [x] Animations completed

## Backend

- [x] Go server completed
- [x] SQLite configured
- [x] Models completed
- [x] Repositories completed
- [x] Services completed
- [x] Handlers completed
- [x] API routes completed
- [x] Validation completed
- [x] Security reviewed

## Documentation

- [x] PRD.md
- [x] ARCHITECTURE.md
- [x] DESIGN.md
- [x] TASKS.md
- [x] README.md

## Final

- [x] Production build works
- [ ] Website deployed
- [ ] API deployed
- [ ] Domain configured
- [x] Final QA completed

---

# CURRENT TASK

Current task:

    PHASE 22 — DEPLOYMENT (Awaiting production host/infrastructure selection)

---

# PROJECT STATUS

Status:

    READY FOR DEPLOYMENT

Completed Phases:

    Phase 1 — Project Foundation
    Phase 2 — Frontend Architecture
    Phase 3 — Navigation
    Phase 4 — Home / Hero
    Phase 5 — About
    Phase 6 — Database
    Phase 7 — Go Backend
    Phase 8 — Backend Services
    Phase 9 — REST API
    Phase 10 — Projects UI
    Phase 11 — Skills UI
    Phase 12 — Contact
    Phase 13 — Visual Refinement
    Phase 14 — Responsive Design
    Phase 15 — Accessibility
    Phase 16 — Security
    Phase 17 — Performance
    Phase 18 — Testing
    Phase 19 — Content
    Phase 20 — Final Review
    Phase 21 — Production Preparation

Next Phase:

    Phase 22 — Deployment (TASK-054, TASK-055, TASK-056)