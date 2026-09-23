Personal Developer Portfolio — Technical & Product Decision Log

Project: Valentine Omondi Awili Personal Portfolio
Status: Active
Document purpose: Record important product, design, architecture, security, and implementation decisions so future development remains consistent.

1. How to Use This Document

This file records decisions that materially affect the project.

Before introducing a major feature, dependency, architectural change, or design direction:

Check whether an existing decision already covers it.

Prefer the existing decision unless there is a clear reason to change it.

If the decision changes, update this document instead of silently overriding it.

Keep decisions practical and proportional to a personal portfolio.

This document explains why choices were made. It does not replace PRD.md, ARCHITECTURE.md, DESIGN.md, SECURITY.md, TASKS.md, or MEMORY.md.

2. Project Context

The project is a professional personal portfolio for Valentine Omondi Awili, an ICT/web developer based in Kenya.

The portfolio should present Valentine as a serious developer with practical experience across frontend development, backend development, APIs, databases, Linux, Git, networking, and related technologies.

The portfolio is intentionally designed as a premium, editorial, story-driven website rather than a generic developer dashboard.

3. Decision Summary

ID

Decision

Status

D001

React + Vite + JavaScript frontend

Accepted

D002

Go REST API backend

Accepted

D003

SQLite database

Accepted

D004

REST/JSON communication

Accepted

D005

Layered backend architecture

Accepted

D006

Component-based React architecture

Accepted

D007

Premium editorial visual direction

Accepted

D008

Reference website used for design direction only

Accepted

D009

Use original content and assets

Accepted

D010

Initial navigation: Home, About, Projects, Skills, Contact

Accepted

D011

No generic dashboard-style portfolio

Accepted

D012

No blog, net-worth, or celebrity sections initially

Accepted

D013

No authentication in the initial public portfolio

Accepted

D014

No cookies unless future functionality requires them

Accepted

D015

Future authentication should use secure server-side sessions

Accepted

D016

Contact form requires server-side validation

Accepted

D017

Contact endpoint requires rate limiting

Accepted

D018

CORS should be restricted

Accepted

D019

SQL queries must be parameterized

Accepted

D020

SQLite remains server-side only

Accepted

D021

Secrets and environment configuration stay outside Git

Accepted

D022

Production deployment must use HTTPS

Accepted

D023

Security headers should be applied

Accepted

D024

Responsive design is required

Accepted

D025

Accessibility is a first-class requirement

Accepted

D026

Performance and asset optimization matter

Accepted

D027

Avoid unnecessary dependencies

Accepted

D028

Avoid overengineering

Accepted

D029

Documentation must stay synchronized with major changes

Accepted

D030

Portfolio content must be factual and verifiable

Accepted

D031

Animation should support storytelling, not distract

Accepted

D032

Image-led sections should be used intentionally

Accepted

D033

Initial API remains small and focused

Accepted

D034

Admin dashboard is outside the initial scope

Accepted

D035

No AI chatbot in the initial portfolio

Accepted

D036

No CMS, payment system, or user-account system initially

Accepted

D037

Frontend/backend separation is maintained

Accepted

D038

Security, maintainability, and clarity take priority over unnecessary complexity

Accepted

4. Architecture Decisions

D001 — React + Vite + JavaScript

Decision: Use React with Vite and JavaScript for the frontend.

Reason:

React supports reusable UI components.

Vite provides a fast development environment and build process.

JavaScript matches the current learning and development direction.

The project does not require TypeScript at this stage.

Constraint: Do not introduce TypeScript merely for fashion or complexity. Reconsider only if the project grows enough that the benefits clearly justify migration.

D002 — Go REST API Backend

Decision: Use Go for the backend API.

Reason:

Go is lightweight and well suited to REST APIs.

It provides strong performance with relatively low operational complexity.

It aligns with the developer's current backend learning direction.

It creates a clear separation between frontend and backend responsibilities.

Constraint: The backend should remain simple and proportional to portfolio requirements.

D003 — SQLite Database

Decision: Use SQLite as the initial database.

Reason:

The portfolio has relatively small data requirements.

SQLite is simple to deploy and maintain.

It removes unnecessary database infrastructure.

It is sufficient for projects, skills, and contact-message storage at the initial scale.

Constraint: Do not introduce PostgreSQL, MongoDB, or multiple databases unless a real requirement appears.

D004 — REST/JSON API

Decision: The frontend communicates with the Go backend through REST endpoints using JSON.

Reason:

Simple and widely understood.

Easy to debug.

Suitable for the portfolio's limited API requirements.

Keeps frontend and backend loosely coupled.

D005 — Layered Backend Architecture

Decision: Structure the backend around layers such as:

HTTP Handler
    ↓
Service / Business Logic
    ↓
Repository / Data Access
    ↓
SQLite

Reason: This provides separation of concerns without introducing unnecessary enterprise architecture.

Constraint: Layers should remain lightweight. Do not create abstractions that add complexity without a practical benefit.

D006 — Component-Based React Architecture

Decision: Build the frontend from reusable React components.

Reason:

Improves maintainability.

Allows sections to evolve independently.

Makes responsive behavior easier to manage.

Prevents one large page component from becoming difficult to maintain.

Constraint: Components should represent meaningful UI or behavior rather than splitting every small HTML element into a component.

D037 — Frontend/Backend Separation

Decision: Keep frontend and backend responsibilities separate.

Frontend responsibilities:

Rendering UI.

Navigation.

User interaction.

Client-side presentation state.

Calling the API.

Backend responsibilities:

Validation.

Business rules.

Data persistence.

API responses.

Security controls.

Reason: This keeps the system understandable and allows either side to evolve independently.

5. Product and Scope Decisions

D010 — Initial Navigation

Decision: The initial portfolio contains:

Home

About

Projects

Skills

Contact

Reason: These sections cover the essential professional story without unnecessary navigation complexity.

D011 — No Generic Dashboard Portfolio

Decision: The portfolio should not look like a SaaS dashboard, admin panel, or generic AI-generated developer template.

Reason: The primary goal is personal storytelling and professional positioning.

The design should emphasize:

typography

whitespace

large imagery

editorial composition

storytelling

project narratives

accomplishments and progression

D012 — Exclude Unrelated Reference Sections

Decision: The portfolio will not initially include:

Net worth

Celebrity-related sections

Elon Musk content

Blog

Articles

Corporate/company sections unrelated to Valentine

Reason: These elements belong to the reference site's identity and do not represent the actual product requirements.

D034 — No Admin Dashboard Initially

Decision: Do not build an admin dashboard in the first version.

Reason: The initial objective is the public portfolio. An admin system would introduce authentication, authorization, session management, and additional security requirements without being necessary for launch.

Future condition: An admin dashboard may be considered if managing projects or contact messages manually becomes necessary.

D035 — No AI Chatbot Initially

Decision: Do not add an AI chatbot to the portfolio's initial version.

Reason: It is not necessary to communicate the developer's work and would introduce additional infrastructure, API costs, privacy considerations, and maintenance requirements.

D036 — No CMS, Payments, or User Accounts Initially

Decision: The initial portfolio will not include:

CMS functionality

Payment processing

User registration

User accounts

Subscription systems

Reason: These features are outside the portfolio's core purpose.

6. Design Decisions

D007 — Premium Editorial Design

Decision: Use a premium editorial design direction.

Characteristics:

Strong typography.

Large visual compositions.

Generous spacing.

Long-form storytelling.

Image-led sections.

Clear visual hierarchy.

Intentional transitions.

Professional and restrained interaction design.

Reason: The portfolio should communicate identity and capability rather than simply list technologies.

D008 — Reference Site as Design Direction Only

Decision: The provided portfolio reference site may influence layout philosophy, visual hierarchy, spacing, storytelling flow, and section composition.

Reason: The user wants a similar overall experience.

Important restriction:

Do not copy source code.

Do not copy text.

Do not copy proprietary images or assets.

Do not reproduce the reference site's personal identity.

Do not present the reference site's content as Valentine's content.

The implementation must be original.

D009 — Original Content and Assets

Decision: Use original portfolio content and appropriate original/licensed assets.

Reason: The website represents a real developer and must accurately represent his experience.

Do not invent:

jobs

clients

awards

testimonials

salaries

project metrics

certifications

employment history

company relationships

If information is unavailable, leave it out or ask the user.

D031 — Animation Supports Storytelling

Decision: Animation should be subtle and purposeful.

Appropriate uses include:

page transitions

section reveals

image movement

hover states

timeline progression

navigation transitions

Avoid:

constant motion

excessive parallax

distracting effects

animation that harms readability

animation that slows page performance

Respect prefers-reduced-motion where appropriate.

D032 — Intentional Image-Led Sections

Decision: Use large images where they strengthen the story.

Images should communicate:

projects

development work

environments

personal professional identity

accomplishments

Images should not be added merely to fill empty space.

7. Authentication and Session Decisions

D013 — No Authentication Initially

Decision: The initial public portfolio does not require authentication.

Reason: Visitors only need to view portfolio information and optionally submit a contact message.

Do not introduce login functionality simply because the backend exists.

D014 — No Cookies Initially

Decision: Do not introduce cookies unless a future feature genuinely requires them.

Reason: A public portfolio with no authenticated sessions does not need authentication cookies.

Avoid unnecessary session state and tracking.

D015 — Future Authentication Uses Secure Server-Side Sessions

If authentication is added later, use server-managed sessions with secure cookie settings.

Expected defaults include:

HttpOnly
Secure
SameSite=Lax

The exact settings may be tightened depending on the deployment and authentication flow.

Avoid storing sensitive authentication credentials or long-lived authentication tokens in localStorage.

8. API Decisions

D033 — Keep the Initial API Small

The initial API should focus on real portfolio requirements.

Expected endpoints include:

GET  /api/projects
GET  /api/projects/:id
GET  /api/skills
POST /api/contact

Additional endpoints should only be introduced when a real feature requires them.

D016 — Server-Side Validation for Contact Form

Decision: Validate contact submissions on the backend even if the frontend already validates them.

Validate at minimum:

required fields

valid email format

sensible length limits

message length

unexpected input

Reason: Client-side validation can be bypassed.

D017 — Rate Limit Contact Endpoint

Decision: The public contact endpoint must be rate limited.

Reason: Public forms are common targets for spam and automated abuse.

The exact limits may change after deployment based on observed traffic.

9. Security Decisions

D018 — Restricted CORS

Decision: CORS must allow only trusted frontend origins in production.

Avoid permissive production configurations such as allowing every origin without a clear reason.

D019 — Parameterized SQL

Decision: All user-controlled values used in SQLite queries must use parameterized queries.

Do not build SQL statements by concatenating untrusted input.

D020 — SQLite Server-Side Only

Decision: The SQLite database file must never be directly exposed to the browser.

The browser communicates with the Go API; the API communicates with SQLite.

D021 — Secrets Outside Git

Decision: Secrets and environment-specific configuration must not be committed to the repository.

Examples:

API keys

database credentials where applicable

session secrets

private service credentials

production-specific secrets

Use environment variables or deployment secret management.

.env files containing secrets should be ignored by Git.

D022 — HTTPS in Production

Decision: Production traffic should use HTTPS.

This protects:

contact submissions

authentication if added later

general communication between browser and server

D023 — Security Headers

Decision: Apply appropriate HTTP security headers in production.

Depending on deployment, this may include controls such as:

Content-Security-Policy

X-Content-Type-Options

Referrer-Policy

frame-ancestors / clickjacking protection

appropriate permissions policy

Exact configuration should be tested rather than copied blindly.

D038 — Security, Maintainability, and Clarity Over Complexity

Security controls should be practical and proportional to the application.

The project should avoid both extremes:

insecure shortcuts

unnecessary enterprise-level complexity

The preferred approach is:

Simple
  ↓
Correct
  ↓
Secure
  ↓
Maintainable

10. Database Decisions

The initial SQLite schema should remain small.

Potential core entities include:

projects
skills
messages

Future tables should only be added when required.

Contact messages should contain only the information genuinely needed to respond to the visitor.

Avoid collecting unnecessary personal information.

11. Performance Decisions

D026 — Optimize Performance

The portfolio is image-heavy and visually ambitious, so performance must be considered during implementation rather than after completion.

Priorities:

appropriately sized images

modern image formats where practical

lazy loading for below-the-fold media

avoiding unnecessary JavaScript

avoiding excessive animation

efficient API requests

sensible caching

responsive image sizing

Do not sacrifice usability merely for visual effects.

12. Responsive Design Decisions

D024 — Responsive by Default

The portfolio must work across:

mobile phones

tablets

laptops

desktop monitors

Desktop design should not simply be scaled down for mobile.

Important sections should be deliberately redesigned where necessary for smaller screens.

13. Accessibility Decisions

D025 — Accessibility Is Required

The portfolio should support accessible interaction from the beginning.

Requirements include:

semantic HTML

keyboard navigation

visible focus states

useful labels

meaningful alt text

sufficient text contrast

logical heading hierarchy

accessible form errors

reduced-motion consideration

Accessibility should not be treated as a final cosmetic pass.

14. Dependency Decisions

D027 — Avoid Unnecessary Dependencies

Only add a package when it provides meaningful value.

Before installing a dependency, consider:

Can the requirement be implemented cleanly with the existing stack?

Is the dependency actively maintained?

Does it significantly increase bundle size or complexity?

Does it introduce security or licensing concerns?

Will the project actually use most of its functionality?

Avoid adding libraries simply because they are popular.

15. Avoid Overengineering

D028 — Keep the Architecture Proportional

This is a personal portfolio, not a large enterprise platform.

Avoid introducing:

microservices

message queues

Kubernetes

multiple databases

complex event systems

unnecessary caching infrastructure

elaborate state-management frameworks

complex authentication systems

unnecessary third-party platforms

unless a real project requirement appears.

The simplest architecture that meets the requirements should be preferred.

16. Content Accuracy

D030 — Portfolio Content Must Be Factual

The portfolio is professional documentation, so content must reflect actual experience.

Known profile information can include:

Valentine Omondi Awili

ICT student/developer

Maseno University

Kisumu, Kenya

web development

frontend development

backend/API development

JavaScript

React

Go

SQLite

Git

Linux

networking fundamentals

Specific claims about jobs, organizations, achievements, project outcomes, certifications, or measurable results must be supported by information provided by the user.

Do not fabricate professional credibility.

17. Known Project Portfolio Direction

Projects that may be represented when appropriate include:

Safeguarding Anonymous Reporting Application

A secure reporting-oriented application focused on anonymous safeguarding reports and controlled follow-up communication.

Known technologies and concepts have included:

React

backend APIs

MongoDB/SQLite experimentation

authentication concepts

rate limiting

CORS

security controls

safeguarding workflow

Descriptions must remain accurate to the actual implementation.

HopeReach NGO Website

A responsive NGO-oriented website demonstrating frontend development using technologies such as HTML, CSS, and JavaScript.

Career Guidance Application

A career guidance platform concept involving academic results, career recommendations, educational pathways, and related guidance features.

Only describe features that were actually implemented or clearly label concepts as planned.

18. Git and Change Management

Major changes should be committed in logical units.

Prefer commit structure such as:

feat: add project API
feat: build projects section
fix: validate contact form
style: refine hero typography
docs: update architecture decisions

Avoid large commits that mix unrelated changes.

Do not commit:

secrets

database files when they should remain local

build artifacts unless intentionally required

unnecessary generated files

19. Documentation Synchronization

D029 — Keep Documentation Current

When a major decision changes, update the relevant documentation.

For example:

Change

Documentation to review

New backend architecture

ARCHITECTURE.md + DECISION.md

Major UI direction change

DESIGN.md + DECISION.md

New security mechanism

SECURITY.md + DECISION.md

New feature

PRD.md + TASKS.md + possibly DECISION.md

Agent behavior/context

MEMORY.md

The documentation should describe the actual system, not an outdated planned system.

20. Agent Development Rules

Any AI agent working on this project should follow these rules:

Rule 1 — Read before changing

Review the relevant project documentation before making major architectural or design changes.

Rule 2 — Preserve existing decisions

Do not replace an established decision simply because another technology is newer or more fashionable.

Rule 3 — Ask before expanding scope

If a proposed feature materially expands the product, it should be discussed before implementation.

Rule 4 — Do not invent information

Never create fake professional experience, clients, statistics, testimonials, awards, or achievements.

Rule 5 — Keep the UI original

The reference website is inspiration for structure and visual direction, not a source to copy.

Rule 6 — Security belongs in implementation

Do not postpone basic validation, parameterized queries, CORS restrictions, rate limiting, or secret handling until the end.

Rule 7 — Do not overengineer

Prefer the smallest solution that correctly satisfies the requirement.

Rule 8 — Protect working features

Before changing existing functionality, understand how it currently works and avoid breaking unrelated sections.

Rule 9 — Test important changes

At minimum, verify:

frontend builds

backend starts

API endpoints work

database operations work

contact validation works

responsive layout remains usable

no obvious console/server errors remain

Rule 10 — Update documentation when decisions change

Do not allow the codebase and documentation to drift apart.

21. Rejected or Deferred Ideas

The following ideas are intentionally deferred from the initial release:

Admin dashboard

User authentication

User accounts

AI chatbot

Blog/CMS

Payment system

Subscription system

Multiple databases

Microservices

Complex analytics platform

Unnecessary third-party integrations

Net-worth section

Celebrity/reference-site content

These may be reconsidered only when there is a clear product requirement.

22. Decision Change Process

When a decision needs to change:

Identify the existing decision ID.

Explain the new requirement.

Record why the old decision is no longer sufficient.

Define the replacement decision.

Update related documentation.

Update implementation.

Test the affected functionality.

Example:

Old decision:
SQLite is sufficient for initial portfolio scale.

New requirement:
The application now needs a hosted multi-user database with concurrent write-heavy workloads.

Action:
Evaluate PostgreSQL and document the migration decision.

Do not silently change major architectural choices.

23. Definition of a Good Decision

A good project decision should be:

understandable

justified

proportional to the project

maintainable

secure

reversible where practical

consistent with the portfolio's purpose

The project should optimize for clarity, correctness, security, maintainability, performance, and professional presentation rather than technical complexity for its own sake.

24. Final Principle

The portfolio should feel like a carefully designed professional product, not a collection of technologies.

Every major implementation decision should answer three questions:

Does this help communicate Valentine's professional identity?

Does this improve the actual product or user experience?

Is the added complexity justified?

If the answer is no, the feature or architectural change should normally be avoided or deferred.