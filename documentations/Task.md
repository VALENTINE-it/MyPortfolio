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

- [ ] Create `components/`.
- [ ] Create `pages/`.
- [ ] Create `services/`.
- [ ] Create `hooks/`.
- [ ] Create `utils/`.
- [ ] Create `assets/`.

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

- [ ] Create global CSS.
- [ ] Configure typography.
- [ ] Configure CSS variables.
- [ ] Configure base spacing.
- [ ] Configure responsive breakpoints.
- [ ] Configure global transitions.
- [ ] Remove default browser margins.

Use the design system defined in `DESIGN.md`.

---

# PHASE 3 — NAVIGATION

## TASK-006 — Build Navbar

- [ ] Create `Navbar.jsx`.
- [ ] Add developer name/logo.
- [ ] Add Home navigation.
- [ ] Add About navigation.
- [ ] Add Projects navigation.
- [ ] Add Skills navigation.
- [ ] Add Contact navigation.
- [ ] Add mobile menu.
- [ ] Add active section state.
- [ ] Add smooth scrolling.

Desktop:

    VALENTINE       HOME ABOUT PROJECTS SKILLS CONTACT

Mobile:

    VALENTINE                         MENU

---

## TASK-007 — Build Footer

- [ ] Create `Footer.jsx`.
- [ ] Add developer name.
- [ ] Add professional title.
- [ ] Add GitHub link.
- [ ] Add LinkedIn link.
- [ ] Add email.
- [ ] Add copyright.
- [ ] Make footer responsive.

---

# PHASE 4 — HOME / HERO

## TASK-008 — Build Hero Section

- [ ] Create `Hero.jsx`.
- [ ] Add "I'm".
- [ ] Add developer name.
- [ ] Add professional title.
- [ ] Add short introduction.
- [ ] Add profile image.
- [ ] Add Projects CTA.
- [ ] Add Contact CTA.
- [ ] Add social links.

The developer name should be one of the largest elements
on the page.

---

## TASK-009 — Hero Animations

- [ ] Add initial text reveal.
- [ ] Add image reveal.
- [ ] Add subtle movement.
- [ ] Respect reduced-motion preferences.
- [ ] Verify animation performance.

Animations must remain subtle.

---

# PHASE 5 — ABOUT

## TASK-010 — Build About Section

- [ ] Create `About.jsx`.
- [ ] Add About heading.
- [ ] Add large personal image.
- [ ] Add personal introduction.
- [ ] Add education information.
- [ ] Add development journey.
- [ ] Add current professional focus.
- [ ] Add "What I Do" content.

The layout should follow the editorial style defined in
`DESIGN.md`.

---

## TASK-011 — About Responsive Layout

- [ ] Create desktop layout.
- [ ] Create tablet layout.
- [ ] Create mobile layout.
- [ ] Test image positioning.
- [ ] Test typography.
- [ ] Test spacing.

---

# PHASE 6 — DATABASE

## TASK-012 — Configure SQLite

- [ ] Add SQLite dependency.
- [ ] Create database connection.
- [ ] Configure database path.
- [ ] Create database initialization.
- [ ] Add error handling.

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

- [ ] Create tables.
- [ ] Add appropriate indexes.
- [ ] Add timestamps.
- [ ] Test database creation.

---

# PHASE 7 — GO BACKEND

## TASK-014 — Create Backend Models

- [ ] Create `Project` model.
- [ ] Create `Skill` model.
- [ ] Create `Contact` model.
- [ ] Add JSON serialization.
- [ ] Verify model structure.

---

## TASK-015 — Create Project Repository

- [ ] Create project repository.
- [ ] Add `GetAllProjects`.
- [ ] Add `GetProjectByID`.
- [ ] Add database queries.
- [ ] Use parameterized queries.
- [ ] Handle database errors.

---

## TASK-016 — Create Skill Repository

- [ ] Create skill repository.
- [ ] Add `GetAllSkills`.
- [ ] Add category filtering if required.
- [ ] Handle database errors.

---

## TASK-017 — Create Contact Repository

- [ ] Create contact repository.
- [ ] Add message creation.
- [ ] Validate database operations.
- [ ] Handle database errors.

---

# PHASE 8 — BACKEND SERVICES

## TASK-018 — Create Project Service

- [ ] Create `ProjectService`.
- [ ] Retrieve projects.
- [ ] Validate project data.
- [ ] Connect service to repository.

---

## TASK-019 — Create Skill Service

- [ ] Create `SkillService`.
- [ ] Retrieve skills.
- [ ] Group skills by category if necessary.

---

## TASK-020 — Create Contact Service

- [ ] Create `ContactService`.
- [ ] Validate contact data.
- [ ] Validate email.
- [ ] Reject invalid requests.
- [ ] Save valid messages.

---

# PHASE 9 — REST API

## TASK-021 — Projects API

Implement:

    GET /api/projects

and:

    GET /api/projects/:id

- [ ] Create handlers.
- [ ] Create routes.
- [ ] Return JSON.
- [ ] Handle 404.
- [ ] Handle database errors.
- [ ] Test endpoints.

---

## TASK-022 — Skills API

Implement:

    GET /api/skills

- [ ] Create handler.
- [ ] Create route.
- [ ] Return JSON.
- [ ] Test endpoint.

---

## TASK-023 — Contact API

Implement:

    POST /api/contact

- [ ] Create handler.
- [ ] Parse JSON.
- [ ] Validate request.
- [ ] Call service.
- [ ] Store message.
- [ ] Return success response.
- [ ] Return validation errors.

---

# PHASE 10 — PROJECTS UI

## TASK-024 — Build Project Timeline

- [ ] Create `ProjectTimeline.jsx`.
- [ ] Create `ProjectItem.jsx`.
- [ ] Add year.
- [ ] Add project title.
- [ ] Add category.
- [ ] Add description.
- [ ] Add technologies.
- [ ] Add project image.
- [ ] Add GitHub link.
- [ ] Add live project link.

Projects must NOT use a generic dashboard-style card grid.

The layout should follow the editorial/timeline design.

---

## TASK-025 — Connect Projects to API

Create:

    services/projectService.js

- [ ] Fetch projects from Go API.
- [ ] Display loading state.
- [ ] Display projects.
- [ ] Display error state.
- [ ] Add retry functionality.

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

- [ ] Add scroll reveal.
- [ ] Add image reveal.
- [ ] Add hover effects.
- [ ] Add subtle link animations.
- [ ] Test reduced-motion behavior.

---

# PHASE 11 — SKILLS UI

## TASK-027 — Build Skills Section

- [ ] Create `Skills.jsx`.
- [ ] Create skill categories.
- [ ] Display frontend skills.
- [ ] Display backend skills.
- [ ] Display tools.
- [ ] Display other technical skills.

Do not use percentage-based skill bars.

---

## TASK-028 — Connect Skills to API

Create:

    services/skillService.js

- [ ] Fetch skills.
- [ ] Display loading state.
- [ ] Display skills.
- [ ] Handle errors.
- [ ] Group skills by category.

---

# PHASE 12 — CONTACT

## TASK-029 — Build Contact UI

- [ ] Create `ContactForm.jsx`.
- [ ] Add name input.
- [ ] Add email input.
- [ ] Add subject input.
- [ ] Add message textarea.
- [ ] Add submit button.
- [ ] Add GitHub link.
- [ ] Add LinkedIn link.
- [ ] Add email.

The form should follow the editorial design.

Avoid dashboard-style form cards.

---

## TASK-030 — Contact Form Validation

Validate:

- [ ] Name is required.
- [ ] Email is required.
- [ ] Email format is valid.
- [ ] Subject is required.
- [ ] Message is required.
- [ ] Message has reasonable length.

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

- [ ] Hero proportions.
- [ ] Typography scale.
- [ ] Section spacing.
- [ ] Image sizes.
- [ ] Project composition.
- [ ] Timeline layout.
- [ ] Navigation.
- [ ] Footer.
- [ ] Overall visual rhythm.

The website should maintain the editorial visual direction
throughout the entire page.

---

## TASK-033 — Typography Refinement

- [ ] Configure primary font.
- [ ] Configure hero typography.
- [ ] Configure section headings.
- [ ] Configure project headings.
- [ ] Configure body text.
- [ ] Configure metadata.
- [ ] Verify mobile typography.

Use fluid typography with `clamp()` where appropriate.

---

## TASK-034 — Image Optimization

- [ ] Compress large images.
- [ ] Use appropriate formats.
- [ ] Add image dimensions.
- [ ] Add alt text.
- [ ] Lazy-load non-critical images.
- [ ] Verify mobile image performance.

---

# PHASE 14 — RESPONSIVE DESIGN

## TASK-035 — Desktop Testing

Test at:

    1440px
    1366px
    1280px
    1024px

Verify:

- [ ] Navigation.
- [ ] Hero.
- [ ] About.
- [ ] Projects.
- [ ] Skills.
- [ ] Contact.
- [ ] Footer.

---

## TASK-036 — Tablet Testing

Test at:

    768px
    820px
    912px

Verify:

- [ ] Navigation.
- [ ] Typography.
- [ ] Images.
- [ ] Project layout.
- [ ] Forms.
- [ ] Spacing.

---

## TASK-037 — Mobile Testing

Test at:

    375px
    390px
    414px

Verify:

- [ ] Mobile navigation.
- [ ] Hero.
- [ ] Profile image.
- [ ] About.
- [ ] Project timeline.
- [ ] Skills.
- [ ] Contact form.
- [ ] Footer.

No horizontal scrolling should occur.

---

# PHASE 15 — ACCESSIBILITY

## TASK-038 — Accessibility Review

- [ ] Use semantic HTML.
- [ ] Add image alt text.
- [ ] Add form labels.
- [ ] Test keyboard navigation.
- [ ] Add visible focus states.
- [ ] Check color contrast.
- [ ] Check heading hierarchy.
- [ ] Test mobile menu with keyboard.
- [ ] Add reduced-motion support.

---

# PHASE 16 — SECURITY

## TASK-039 — Backend Security

- [ ] Configure CORS.
- [ ] Validate all API inputs.
- [ ] Use parameterized SQL queries.
- [ ] Limit contact submissions.
- [ ] Avoid exposing internal errors.
- [ ] Validate request content type.
- [ ] Keep secrets out of source code.

---

## TASK-040 — Environment Configuration

Create:

    .env.example

Example:

    PORT=8080
    DATABASE_URL=./portfolio.db
    FRONTEND_URL=http://localhost:5173

- [ ] Create `.env.example`.
- [ ] Add `.env` to `.gitignore`.
- [ ] Load environment variables safely.
- [ ] Verify production configuration.

---

# PHASE 17 — PERFORMANCE

## TASK-041 — Frontend Performance

- [ ] Optimize images.
- [ ] Lazy-load non-critical images.
- [ ] Minimize unnecessary React re-renders.
- [ ] Avoid unnecessary dependencies.
- [ ] Check Vite production build.

---

## TASK-042 — Backend Performance

- [ ] Check database queries.
- [ ] Add appropriate indexes.
- [ ] Avoid unnecessary database calls.
- [ ] Verify API response times.

---

# PHASE 18 — TESTING

## TASK-043 — Frontend Functional Testing

Test:

- [ ] Navigation works.
- [ ] Mobile menu works.
- [ ] Project data loads.
- [ ] Skills data loads.
- [ ] Contact form works.
- [ ] Loading states work.
- [ ] Error states work.
- [ ] Links work.

---

## TASK-044 — Backend API Testing

Test:

    GET /api/health
    GET /api/projects
    GET /api/projects/:id
    GET /api/skills
    POST /api/contact

Verify:

- [ ] Successful requests.
- [ ] Invalid requests.
- [ ] Missing data.
- [ ] Invalid IDs.
- [ ] Database errors.
- [ ] Correct HTTP status codes.

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

- [ ] Add developer name.
- [ ] Add professional title.
- [ ] Add biography.
- [ ] Add education.
- [ ] Add development journey.
- [ ] Add professional links.
- [ ] Add email.
- [ ] Add profile image.

---

## TASK-047 — Add Projects

Add initial projects.

For every project:

- [ ] Project name.
- [ ] Description.
- [ ] Year.
- [ ] Category.
- [ ] Technologies.
- [ ] Project image.
- [ ] GitHub URL.
- [ ] Live URL if available.

---

## TASK-048 — Add Skills

Add:

### Frontend

- [ ] HTML
- [ ] CSS
- [ ] JavaScript
- [ ] React
- [ ] Vite

### Backend

- [ ] Go
- [ ] REST APIs
- [ ] SQLite
- [ ] MongoDB

### Tools

- [ ] Git
- [ ] GitHub
- [ ] Linux
- [ ] VS Code

### Other

- [ ] Networking
- [ ] Debugging
- [ ] Problem Solving

Only include technologies that accurately represent the developer's
current experience.

---

# PHASE 20 — FINAL REVIEW

## TASK-049 — Design Review

Compare the completed website against:

    DESIGN.md

Check:

- [ ] Visual hierarchy.
- [ ] Typography.
- [ ] Spacing.
- [ ] Images.
- [ ] Project storytelling.
- [ ] Navigation.
- [ ] Responsive behavior.
- [ ] Animations.
- [ ] Contact experience.

---

## TASK-050 — Architecture Review

Compare implementation against:

    ARCHITECTURE.md

Verify:

- [ ] React handles UI.
- [ ] Services handle API communication.
- [ ] Go handles API requests.
- [ ] Services contain business logic.
- [ ] Repositories handle database operations.
- [ ] SQLite is only accessed by Go.
- [ ] Components remain reusable.

---

## TASK-051 — Remove Unwanted Features

Confirm the website does NOT contain:

- [ ] Net Worth
- [ ] Articles
- [ ] Blog
- [ ] Celebrity content
- [ ] Unrelated sections
- [ ] Generic AI dashboard components
- [ ] Fake skill percentages

---

# PHASE 21 — PRODUCTION PREPARATION

## TASK-052 — Production Build

Frontend:

    npm run build

- [ ] Build succeeds.
- [ ] No console errors.
- [ ] No broken assets.
- [ ] Verify production output.

Backend:

    go build

- [ ] Build succeeds.
- [ ] No compilation errors.
- [ ] Environment configuration works.

---

## TASK-053 — Final Browser Testing

Test the production build in:

- [ ] Chrome
- [ ] Firefox
- [ ] Edge

Check:

- [ ] Desktop.
- [ ] Tablet.
- [ ] Mobile.
- [ ] Navigation.
- [ ] Images.
- [ ] Forms.
- [ ] API communication.

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

- [ ] All sections work.
- [ ] All projects display correctly.
- [ ] Contact form works.
- [ ] Mobile layout works.
- [ ] Desktop layout works.
- [ ] No console errors.
- [ ] No broken links.
- [ ] No missing images.
- [ ] No exposed secrets.
- [ ] Accessibility reviewed.
- [ ] Performance reviewed.

---

# COMPLETION CHECKLIST

## Frontend

- [ ] React + Vite configured
- [ ] Navbar completed
- [ ] Hero completed
- [ ] About completed
- [ ] Projects completed
- [ ] Skills completed
- [ ] Contact completed
- [ ] Footer completed
- [ ] Responsive design completed
- [ ] Animations completed

## Backend

- [ ] Go server completed
- [ ] SQLite configured
- [ ] Models completed
- [ ] Repositories completed
- [ ] Services completed
- [ ] Handlers completed
- [ ] API routes completed
- [ ] Validation completed
- [ ] Security reviewed

## Documentation

- [ ] PRD.md
- [ ] ARCHITECTURE.md
- [ ] DESIGN.md
- [ ] TASKS.md
- [ ] README.md

## Final

- [ ] Production build works
- [ ] Website deployed
- [ ] API deployed
- [ ] Domain configured
- [ ] Final QA completed

---

# CURRENT TASK

The AI development agent must always identify the current task before
making changes.

Current task:

    TASK-004

After completing a task:

1. Test it.
2. Review it.
3. Mark it `[x]`.
4. Move to the next task.

Do not skip tasks without documenting the reason.

---

# PROJECT STATUS

Status:

    IN PROGRESS

Completed:

    TASK-001 — Create Project Repository
    TASK-002 — Initialize React + Vite
    TASK-003 — Initialize Go Backend

Current Phase:

    Phase 2 — Frontend Architecture

Next Task:

    TASK-004 — Create Frontend Folder Structure