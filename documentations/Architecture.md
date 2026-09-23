# ARCHITECTURE.md

# Personal Portfolio Architecture

## 1. Project Overview

This is a personal developer portfolio built with:

- React
- Vite
- JavaScript
- CSS
- Go
- REST API
- SQLite

The website should visually follow the structure, composition, spacing,
section flow, typography hierarchy, image placement, and storytelling
approach of the provided reference portfolio:

https://syedameen-elonmuskporfolio.netlify.app/

The implementation must be original and use the developer's own content,
projects, images, branding, and information.

The portfolio will NOT include:

- Net Worth section
- Articles section
- Blog page
- Elon Musk-related content
- Unrelated corporate/company sections

---

# 2. Main Navigation

The website will contain five primary sections:

1. Home
2. About
3. Projects
4. Skills
5. Contact

The navigation should remain visible at the top of the website.

Desktop:

    [ VALENTINE ]     HOME   ABOUT   PROJECTS   SKILLS   CONTACT

Mobile:

    [ VALENTINE ]                         [ MENU ]

The navigation should smoothly scroll to the corresponding section.

---

# 3. Overall UI Concept

The website should feel like a personal story rather than a
traditional dashboard-style portfolio.

The design should use:

- Large typography
- Large visual sections
- Generous whitespace
- Strong image placement
- Minimal cards
- Clean navigation
- Subtle animations
- Smooth scrolling
- Strong section headings
- Full-width sections
- Timeline-style project presentation
- Professional developer aesthetic

Avoid:

- Excessive rounded cards
- Excessive gradients
- Generic AI-generated dashboard styling
- Excessive icons
- Emoji
- Unnecessary animations
- Overloaded UI

The website should feel like a premium personal portfolio.

---

# 4. Overall Page Structure

The complete website should follow this structure:

    ┌─────────────────────────────┐
    │          NAVBAR             │
    ├─────────────────────────────┤
    │                             │
    │          HOME               │
    │        HERO SECTION         │
    │                             │
    ├─────────────────────────────┤
    │                             │
    │          ABOUT              │
    │      PERSONAL STORY         │
    │                             │
    ├─────────────────────────────┤
    │                             │
    │        PROJECTS             │
    │      PROJECT TIMELINE       │
    │                             │
    ├─────────────────────────────┤
    │                             │
    │         SKILLS              │
    │      TECHNICAL AREAS        │
    │                             │
    ├─────────────────────────────┤
    │                             │
    │        CONTACT              │
    │      CONTACT SECTION        │
    │                             │
    ├─────────────────────────────┤
    │          FOOTER             │
    └─────────────────────────────┘

The application should behave primarily as a single-page portfolio.

React Router may still be used if individual routes are required later.

---

# 5. HOME SECTION

## Purpose

The Home section should create the first impression.

The layout should closely follow the reference site's hero composition.

## Layout

    ┌─────────────────────────────────────────────┐
    │                                             │
    │                  I'M                        │
    │                                             │
    │              VALENTINE                      │
    │                                             │
    │        FULL-STACK DEVELOPER                 │
    │                                             │
    │   I build modern web applications using     │
    │   React, JavaScript and Go.                 │
    │                                             │
    │        [ VIEW MY WORK ]                     │
    │                                             │
    │              [ IMAGE ]                      │
    │                                             │
    └─────────────────────────────────────────────┘

The hero should use a large profile image or developer photograph.

The image should have strong visual presence but should not overwhelm
the name and introduction.

## Hero Content

Example:

    I'M

    VALENTINE OMONDI AWILI

    FULL-STACK DEVELOPER

    I build modern, responsive and practical web applications
    using React, JavaScript and Go.

Primary CTA:

    VIEW MY PROJECTS

Secondary CTA:

    CONTACT ME

---

# 6. ABOUT SECTION

## Purpose

The About section should tell the developer's story.

It should use a large editorial layout rather than a traditional
profile card.

Layout:

    ┌─────────────────────────────────────────────┐
    │                                             │
    │              ABOUT ME                       │
    │                                             │
    │    [ LARGE IMAGE ]     WHO I AM             │
    │                                             │
    │                         Introduction        │
    │                         Education           │
    │                         Development journey │
    │                         Career interests    │
    │                                             │
    └─────────────────────────────────────────────┘

Below the introduction:

    WHAT I DO

    Frontend Development
    Backend Development
    Full-Stack Development
    Web Applications
    API Development

The content should be presented as part of the page narrative rather
than as excessive UI cards.

---

# 7. PROJECTS SECTION

## Purpose

Projects are one of the most important parts of the portfolio.

The reference website presents accomplishments using large visual
sections and timeline-like storytelling.

The portfolio should use the same general concept for projects.

## Layout

    PROJECTS

    SELECTED WORK

    ─────────────────────────────────────

    01

    [ PROJECT IMAGE ]

    PROJECT NAME

    Short description explaining the problem,
    solution and purpose.

    React
    JavaScript
    Go
    SQLite

    [ GITHUB ]     [ LIVE PROJECT ]


    ─────────────────────────────────────

    02

    [ PROJECT IMAGE ]

    PROJECT NAME

    Description

    Technologies

    [ GITHUB ]     [ LIVE PROJECT ]


    ─────────────────────────────────────

    03

    [ PROJECT IMAGE ]

    PROJECT NAME

    Description

    Technologies

    [ GITHUB ]     [ LIVE PROJECT ]

Projects should appear as large sections rather than small generic
three-column cards.

## Project Data

Projects will be retrieved from the Go backend.

Example:

    GET /api/projects

Example project:

    {
      "id": 1,
      "title": "Safeguarding Reporting Platform",
      "description": "Anonymous safeguarding reporting platform.",
      "year": 2026,
      "technologies": [
        "React",
        "Go",
        "MongoDB"
      ],
      "image": "/projects/safeguarding.png",
      "githubUrl": "...",
      "liveUrl": "..."
    }

---

# 8. PROJECT TIMELINE

Projects should optionally contain a year.

Example:

    2026
       │
       ├── Safeguarding Reporting Platform
       │
       ├── Career Guidance Platform
       │
       └── Portfolio Website

This creates the same storytelling feeling as the reference
accomplishment sections without copying its content.

---

# 9. SKILLS SECTION

## Purpose

The Skills section should communicate technical capability visually.

Instead of using progress bars such as:

    React      90%
    JavaScript 80%

the portfolio should use grouped technical areas.

Layout:

    SKILLS

    FRONTEND

    HTML
    CSS
    JavaScript
    React
    Vite
    Responsive Design


    BACKEND

    Go
    REST APIs
    SQLite
    MongoDB


    TOOLS

    Git
    GitHub
    Linux
    VS Code


    OTHER

    Networking
    Debugging
    Problem Solving

The section should remain minimal and typography-focused.

---

# 10. TECHNOLOGY PRESENTATION

Technologies can be displayed as simple text labels.

Example:

    React  •  JavaScript  •  Go  •  REST API
    SQLite •  Git         •  Linux

Avoid excessive technology logos.

Icons may be used sparingly where they improve recognition.

---

# 11. CONTACT SECTION

## Purpose

The Contact section should provide a direct way to communicate.

Layout:

    ┌─────────────────────────────────────────────┐
    │                                             │
    │              LET'S TALK                     │
    │                                             │
    │      Have a project or opportunity?         │
    │                                             │
    │              [ CONTACT ME ]                 │
    │                                             │
    ├─────────────────────────────────────────────┤
    │                                             │
    │ Name                                        │
    │ [____________________________________]      │
    │                                             │
    │ Email                                       │
    │ [____________________________________]      │
    │                                             │
    │ Subject                                     │
    │ [____________________________________]      │
    │                                             │
    │ Message                                     │
    │ [____________________________________]      │
    │                                             │
    │              [ SEND MESSAGE ]               │
    │                                             │
    └─────────────────────────────────────────────┘

The form will communicate with:

    POST /api/contact

The Go backend will validate the request and store the message in
SQLite.

---

# 12. FOOTER

The footer should remain minimal.

Example:

    VALENTINE OMONDI AWILI

    Full-Stack Developer

    GitHub
    LinkedIn
    Email

    © 2026 Valentine Omondi Awili

The footer should not contain a blog, articles or unnecessary
navigation.

---

# 13. VISUAL HIERARCHY

The visual hierarchy should be:

    1. Developer name
    2. Professional title
    3. Large imagery
    4. Section titles
    5. Project names
    6. Descriptions
    7. Supporting information

Large headings should be used to create strong visual transitions
between sections.

Example:

    ABOUT

    PROJECTS

    SKILLS

    CONTACT

---

# 14. IMAGE ARCHITECTURE

Images should be used as major visual elements.

Suggested assets:

    public/
    └── images/
        ├── profile.jpg
        ├── about.jpg
        └── projects/
            ├── project-1.jpg
            ├── project-2.jpg
            └── project-3.jpg

Images should use consistent aspect ratios.

Project images should be large enough to visually communicate what
each project is about.

---

# 15. ANIMATION

Animations should be subtle.

Use:

- Fade-in
- Slide-in
- Image reveal
- Hover transitions
- Smooth scrolling
- Navigation transitions

Example:

    Section enters viewport
            ↓
       Fade + Slide
            ↓
       Content visible

Do not animate every element.

Animation should support the storytelling of the portfolio.

---

# 16. RESPONSIVE DESIGN

Desktop:

    Large hero
    Large typography
    Side-by-side layouts
    Large project imagery


Tablet:

    Reduced typography
    Flexible columns
    Reduced spacing


Mobile:

    Single-column layout
    Mobile navigation
    Full-width images
    Smaller typography
    Stacked project content

Example:

Desktop:

    [ IMAGE ]       [ PROJECT INFORMATION ]


Mobile:

    [ IMAGE ]

    [ PROJECT INFORMATION ]

---

# 17. FRONTEND COMPONENT ARCHITECTURE

    src/
    │
    ├── components/
    │   ├── Navbar.jsx
    │   ├── Hero.jsx
    │   ├── SectionTitle.jsx
    │   ├── About.jsx
    │   ├── ProjectTimeline.jsx
    │   ├── ProjectItem.jsx
    │   ├── Skills.jsx
    │   ├── ContactForm.jsx
    │   └── Footer.jsx
    │
    ├── pages/
    │   └── Home.jsx
    │
    ├── services/
    │   ├── projectService.js
    │   ├── skillService.js
    │   └── contactService.js
    │
    ├── assets/
    │
    ├── App.jsx
    └── main.jsx

---

# 18. BACKEND ARCHITECTURE

    backend/
    │
    ├── cmd/
    │   └── server/
    │       └── main.go
    │
    ├── internal/
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
    └── go.mod

---

# 19. DATA FLOW

Projects:

    React
      ↓
    GET /api/projects
      ↓
    Go Handler
      ↓
    Project Service
      ↓
    Project Repository
      ↓
    SQLite
      ↓
    JSON
      ↓
    React
      ↓
    Project Timeline


Contact:

    Contact Form
       ↓
    React Validation
       ↓
    POST /api/contact
       ↓
    Go Handler
       ↓
    Contact Service
       ↓
    Contact Repository
       ↓
    SQLite
       ↓
    Success Response
       ↓
    Success Message

---

# 20. RESPONSIBILITY SEPARATION

React is responsible for:

- UI
- Navigation
- Animations
- Form interaction
- Displaying API data

Go is responsible for:

- API
- Validation
- Business logic
- Database operations
- Contact submissions

SQLite is responsible for:

- Project data
- Skills data
- Contact messages

React must never communicate directly with SQLite.

---

# 21. IMPORTANT DESIGN RULE

The portfolio should be inspired by the reference website's:

- Layout philosophy
- Visual hierarchy
- Large typography
- Full-width sections
- Storytelling approach
- Image placement
- Timeline presentation
- Minimal navigation
- Section spacing
- Animation style

However, the actual implementation must use original:

- Text
- Images
- Branding
- Developer information
- Project content
- Skills
- Colors where appropriate
- Component implementation

Do not copy the reference site's source code or proprietary assets.

---

# 22. EXCLUDED SECTIONS

The following sections from the reference concept must NOT be implemented:

- Net Worth
- Articles
- Blog
- News
- Elon Musk biography
- Elon Musk accomplishments
- Corporate company showcase
- Celebrity quotations

The portfolio should remain focused on the developer and their work.

---

# 23. FINAL UI FLOW

The final experience should be:

    NAVBAR
       ↓
    HERO
       ↓
    ABOUT
       ↓
    PROJECTS
       ↓
    SKILLS
       ↓
    CONTACT
       ↓
    FOOTER

The website should feel like one continuous professional story rather
than five unrelated pages.

The visitor should move naturally from:

    "Who is this developer?"

            ↓

    "What does this developer do?"

            ↓

    "What has this developer built?"

            ↓

    "What technologies does this developer use?"

            ↓

    "How can I contact this developer?"