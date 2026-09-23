Personal Developer Portfolio — Agent Memory & Project Context

1. Purpose

This file provides persistent project context for AI coding agents working on the portfolio.

The agent should read this file before making substantial changes. It describes the project's identity, goals, technology choices, design direction, constraints, and working preferences.

This file is context, not a replacement for the detailed project documentation.

2. Documentation Hierarchy

Before implementing features, use the documentation in this order:

MEMORY.md
   ↓
Project context and agent behavior

PRD.md
   ↓
What the product is and why it exists

ARCHITECTURE.md
   ↓
How the technical system is structured

DESIGN.md
   ↓
How the interface should look and behave

SECURITY.md
   ↓
How the application should be protected

DECISION.md
   ↓
Why important technical/product decisions were made

TASKS.md
   ↓
What should be implemented and in what order

If two documents appear to conflict, do not silently choose one. Identify the conflict and follow the most recent explicit project decision after reviewing the relevant files.

3. Project Identity

This is a personal developer portfolio for Valentine Omondi Awili.

The website should present the developer as a serious software developer and communicate:

Professional identity

Technical capabilities

Projects

Experience

Learning and growth

Problem-solving ability

Contact information

The portfolio should feel personal, intentional, modern, and professional.

It should not look like a generic AI-generated portfolio template.

4. Primary Technology Stack

Frontend

React
Vite
JavaScript
CSS

Backend

Go
REST API

Database

SQLite

Communication

HTTP
JSON
REST-style endpoints

Do not replace the stack with another framework or database without an explicit project decision.

5. Current Product Scope

The primary portfolio sections are:

Home
About
Projects
Skills
Contact

The website should be a long-form experience rather than a collection of unrelated dashboard cards.

6. Design Direction

The supplied reference portfolio is:

https://syedameen-elonmuskporfolio.netlify.app/

Use the reference as a source of design direction only.

The portfolio may take inspiration from its:

Visual hierarchy

Long-form storytelling

Large typography

Large imagery

Editorial spacing

Timeline/accomplishment presentation

Section transitions

Footer structure

Overall storytelling flow

The implementation must remain original.

Never copy:

Source code

Proprietary text

Images

Logos

Celebrity identity

Celebrity biography

Personal information from the reference

7. Visual Personality

The preferred visual direction is:

Premium

Editorial

Minimal but expressive

Professional

Image-led

Typography-led

Spacious

Story-driven

Modern

Intentional

Avoid making the site look like:

A SaaS dashboard

An admin panel

A generic developer template

An AI-generated landing page

A cryptocurrency dashboard

A collection of repetitive cards

8. Avoid Generic AI-Looking UI

Do not automatically add:

Random gradients

Decorative blobs

Excessive glassmorphism

Floating cards everywhere

Fake AI assistants

Chatbots without a product reason

Excessive badges

Emoji-heavy UI

Random glowing effects

Unnecessary animated statistics

Generic "AI-powered" labels

Repetitive three-card layouts

Every visual element should have a reason to exist.

9. Content Rules

Content should be based on the developer's real professional information.

Prioritize:

Real projects

Real technologies

Real experience

Real accomplishments

Real learning progress

Real professional goals

Do not invent:

Jobs

Clients

Companies

Awards

Certifications

Project results

User numbers

Revenue

Testimonials

Professional achievements

If required information is missing, use a clear placeholder or ask for the information rather than inventing it.

10. Current Developer Profile Context

Known professional context that may be useful when creating portfolio content:

Name: Valentine Omondi Awili

Location: Kisumu, Kenya

Education: ICT studies at Maseno University

Primary development direction: Web development / full-stack development

Strong frontend foundation: HTML and CSS

JavaScript: actively developing deeper proficiency

Backend interests: Go

Database interests: SQLite and other backend data technologies

Version control: Git

Operating environment: Linux / Ubuntu

Relevant technical areas include:

HTML

CSS

JavaScript

React

REST APIs

Go

SQLite

Git

Linux

Responsive web design

Debugging

Networking fundamentals

Only display skills on the final portfolio that accurately represent the current project/content state.

11. Known Project Experience

Potential portfolio projects include work such as:

Safeguarding Anonymous Reporting Application

A secure reporting platform designed around anonymous safeguarding reports.

Relevant concepts include:

React

Go/backend development

REST APIs

Database-backed reporting

Authentication/authorization concepts

Privacy

Security

Anonymous reporting workflows

Do not expose sensitive implementation details or real private data in the public portfolio.

HopeReach NGO Website

A responsive NGO website demonstrating frontend development using:

HTML

CSS

JavaScript

Responsive design

Career Guidance Application

A career guidance platform concept involving student performance, career recommendations, and educational guidance.

Only include features that are actually implemented or clearly label them as planned concepts.

12. Feature Scope — Initial Release

The initial release should focus on the portfolio itself.

Included:

Home

About

Projects

Skills

Contact

Responsive design

Backend API where required

SQLite persistence where required

Not included initially:

Blog

Articles

Net worth section

Celebrity content

Admin dashboard

User accounts

Authentication

Complex CMS

Unnecessary AI features

These may be reconsidered later, but should not be added automatically.

13. Authentication and Cookies

Authentication is not required for the initial release.

Cookies are not required for the initial public portfolio.

Do not introduce cookies simply because a backend exists.

If an admin dashboard or other protected functionality is introduced later, the preferred approach is:

Server-side session
        ↓
HttpOnly cookie
        ↓
Secure cookie in production
        ↓
SameSite=Lax by default

Avoid storing sensitive authentication tokens in localStorage when a secure HttpOnly session cookie can be used.

See SECURITY.md for the complete security requirements.

14. Backend Architecture

Use a clear layered structure:

HTTP Handler
      ↓
   Service
      ↓
 Repository
      ↓
   SQLite

Handlers should handle HTTP concerns.

Services should contain application/business logic.

Repositories should handle database operations.

Do not put all backend logic into a single main.go file.

15. Frontend Architecture

Use reusable React components.

A reasonable structure is:

src/
├── components/
├── sections/
├── pages/
├── assets/
├── services/
├── hooks/
└── styles/

Do not create an unnecessarily complex architecture for a portfolio of this size.

Prefer simple, understandable components.

16. API Expectations

Likely initial API resources include:

GET  /api/projects
GET  /api/projects/:id
GET  /api/skills
POST /api/contact

The exact endpoints should follow ARCHITECTURE.md and TASKS.md once implementation begins.

The API should use JSON and predictable HTTP status codes.

17. Contact Form

The contact form should initially collect:

Name

Email

Subject

Message

The backend must validate these fields.

The public contact endpoint must have rate limiting.

Contact submissions must never become publicly readable.

Do not create a public endpoint such as:

GET /api/contacts

unless a future authenticated admin feature explicitly requires it.

18. Security Rules

Security requirements are defined in SECURITY.md.

Important baseline rules:

Validate input on the backend.

Use parameterized SQL.

Restrict CORS.

Rate-limit public abuse-prone endpoints.

Limit request sizes.

Protect the SQLite file.

Never expose .env or secrets.

Never commit database files.

Do not expose stack traces in production responses.

Use HTTPS in production.

Configure appropriate security headers.

Do not log passwords, tokens, API keys, or session credentials.

Frontend validation is for user experience, not security.

19. Database Rules

SQLite is server-side only.

The database must never be placed in the frontend's public/static directory.

Recommended Git exclusions:

.env
.env.*
*.db
*.sqlite
*.sqlite3

Use parameterized queries for user-controlled values.

20. Responsive Design

The portfolio must work on:

Desktop

Laptop

Tablet

Mobile

Do not simply shrink the desktop layout.

Design mobile layouts intentionally, especially for:

Large typography

Navigation

Images

Timeline sections

Project layouts

Footer

21. Performance

The portfolio may use large visual assets, so performance matters.

Prefer:

Optimized images

Modern image formats where appropriate

Correct image dimensions

Lazy loading below the fold

Minimal dependencies

Efficient React rendering

Small API responses

Do not optimize prematurely at the expense of maintainability.

22. Accessibility

Accessibility is part of the implementation, not a later feature.

Use:

Semantic HTML

Proper heading hierarchy

Accessible navigation

Keyboard-friendly interactions

Visible focus states

Meaningful alt text

Sufficient text contrast

Form labels

Useful error messages

Do not use animation that prevents users from accessing content.

Respect reduced-motion preferences where animations are significant.

23. Animation Guidelines

Animations should support storytelling and navigation.

Use motion for:

Section transitions

Revealing content

Navigation transitions

Image presentation

Subtle interaction feedback

Avoid:

Constant movement

Excessive parallax

Distracting loops

Animations that slow navigation

Animation used only because it looks impressive

The site should still be understandable if animations are disabled.

24. Image Guidelines

Images should support the story of the portfolio.

Prefer:

High-quality personal/project imagery

Relevant screenshots

Original project visuals

Purposeful editorial imagery

Do not use random stock images simply to fill empty space.

Do not use copyrighted assets without appropriate permission or licensing.

25. Agent Behavior Rules

When working on this project, the coding agent should:

Read the relevant documentation before making substantial changes.

Preserve existing architecture unless there is a documented reason to change it.

Avoid adding unnecessary dependencies.

Avoid introducing new frameworks without approval.

Avoid inventing content.

Avoid changing unrelated files.

Keep changes focused on the requested task.

Reuse existing components where appropriate.

Follow the security requirements.

Keep code readable and maintainable.

Explain significant architectural changes.

Update documentation when a major decision changes.

26. Agent Decision Rules

Before adding a new feature, ask:

Does the PRD require it?
        ↓
Does the user explicitly request it?
        ↓
Does it solve a real product problem?
        ↓
Does it fit the existing architecture?
        ↓
Does it introduce unnecessary complexity?

If the feature is not required, do not add it automatically.

27. Do Not Overengineer

This is a portfolio, not a large enterprise platform.

Avoid introducing unnecessary:

Microservices

Message queues

Complex caching systems

Multiple databases

Kubernetes

Large state-management frameworks

Complex authentication systems

Unnecessary AI services

Overly abstract design patterns

Use the simplest architecture that correctly solves the requirement.

28. Do Not Break Existing Work

Before modifying an existing feature:

Understand how it currently works.

Identify dependencies.

Make the smallest appropriate change.

Test the affected area.

Avoid unrelated refactoring.

Do not rewrite working sections simply to make them look different.

29. Code Quality

Prefer:

Small focused functions

Clear variable names

Reusable components

Explicit error handling

Predictable API contracts

Consistent formatting

Simple control flow

Avoid:

Giant components

Giant functions

Duplicate logic

Hard-coded secrets

Magic values without explanation

Unused dependencies

Dead code

30. Git Practices

Use meaningful commits when Git history is part of the workflow.

Examples:

feat: add projects section
feat: add contact API
fix: validate contact request
style: refine hero layout
security: add contact rate limiting
docs: update architecture decisions

Do not commit:

Secrets

Database files

Build artifacts that should be ignored

Personal credentials

31. Documentation Update Rules

Update documentation when a major decision changes.

Examples:

If the backend changes:

ARCHITECTURE.md
DECISION.md
MEMORY.md

If security requirements change:

SECURITY.md
DECISION.md
MEMORY.md

If product scope changes:

PRD.md
TASKS.md
DECISION.md
MEMORY.md

Do not modify every documentation file for trivial implementation changes.

32. Current Project Priorities

Prioritize work in this general order:

Correct project structure

Core visual design

Home/hero experience

About section

Projects storytelling

Skills presentation

Contact experience

Backend/API integration

SQLite persistence where required

Security hardening

Responsive refinement

Accessibility refinement

Performance optimization

Final polish

Follow TASKS.md when a more specific implementation sequence exists.

33. Definition of Done

A feature should generally be considered complete when:

It matches the PRD.

It follows the architecture.

It matches the design system.

It works on desktop and mobile where applicable.

It has appropriate accessibility behavior.

It handles errors reasonably.

It does not introduce obvious security problems.

It does not add unnecessary dependencies.

Existing functionality still works.

Relevant documentation is updated when necessary.

34. Important Exclusions

The following should not be introduced without explicit approval:

Blog

Net worth tracking

Celebrity biography/content

User accounts

Admin dashboard

Authentication

AI chatbot

Payment system

CMS

Multiple databases

Microservice architecture

Unnecessary third-party analytics

These exclusions keep the initial product focused.

35. Long-Term Direction

The portfolio should be capable of evolving into a stronger professional platform without forcing complexity into the first release.

Possible future capabilities may include:

Authenticated admin dashboard

Project management

More detailed case studies

Analytics

Blog/articles

Private content

Additional API resources

These are future possibilities, not requirements for the initial release.

36. Final Agent Principle

Build the portfolio as if it were a real professional product, but keep the architecture proportional to its actual needs.

The agent should optimize for:

Clarity
   ↓
Correctness
   ↓
Security
   ↓
Maintainability
   ↓
Performance
   ↓
Visual quality

Do not optimize for the number of technologies used.

Do not add features to make the project appear more complicated.

Do not sacrifice usability for visual effects.

Do not sacrifice security for convenience.

Do not sacrifice maintainability for clever code.

The goal is a polished, original, technically credible personal developer portfolio.
