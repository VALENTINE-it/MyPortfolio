# SECURITY.md

# Personal Developer Portfolio — Security Requirements

## 1. Purpose

This document defines the security requirements for the personal developer portfolio.

The application consists of:

    React + Vite
          ↓
       REST API
          ↓
          Go
          ↓
        SQLite

The portfolio is primarily public, but the backend must still protect:

- Contact submissions
- Database access
- Application configuration
- Future authenticated functionality
- Any sensitive server-side information

Security should follow a defense-in-depth approach.

The frontend must never be treated as the final security boundary.

---

# 2. Security Goals

The application should:

- Protect contact submissions from abuse.
- Prevent SQL injection.
- Validate incoming data.
- Prevent accidental exposure of secrets.
- Protect authentication state if authentication is introduced.
- Secure cookies when cookies are used.
- Restrict cross-origin requests.
- Reduce common HTTP security risks.
- Avoid exposing internal server details.
- Protect the SQLite database from direct public access.
- Provide safe error responses.
- Keep dependencies reasonably up to date.

---

# 3. Security Boundaries

The main security boundary is the Go backend.

Architecture:

    Browser
       |
       | HTTPS
       ↓
    React Frontend
       |
       | JSON / HTTP
       ↓
    Go REST API
       |
       ↓
    Service Layer
       |
       ↓
    Repository Layer
       |
       ↓
    SQLite

The browser and frontend code must be treated as untrusted.

The backend must independently validate:

- Request method
- Content type
- Input length
- Input format
- Required fields
- IDs
- Authentication state where applicable
- Authorization where applicable

---

# 4. HTTPS

Production traffic must use HTTPS.

Requirements:

- Authentication cookies must never be sent over plain HTTP in production.
- Production deployments must use TLS.
- Local development may use HTTP.
- HTTP-to-HTTPS redirection should be enabled where supported.

Production authentication cookies should use:

    Secure

---

# 5. Cookies

Cookies are NOT required for the initial public portfolio functionality.

The current portfolio can operate without cookies because:

- Visitors do not need accounts.
- The contact form does not require authentication.
- Projects and skills are publicly accessible.
- There is currently no admin dashboard.

However, cookies may become necessary if the project later introduces:

- Admin authentication
- Protected project management
- Admin dashboard
- Session-based authentication
- Protected API endpoints

If authentication is introduced, secure server-managed sessions should be preferred over storing authentication credentials in `localStorage`.

---

# 6. Authentication Cookies

If authentication is introduced, the preferred session cookie should use:

    HttpOnly
    Secure
    SameSite=Lax

Example:

    Set-Cookie:
      session=<opaque-session-id>;
      HttpOnly;
      Secure;
      SameSite=Lax;
      Path=/;
      Max-Age=86400

The exact configuration depends on the production deployment architecture.

### HttpOnly

Authentication cookies should be `HttpOnly`.

This prevents JavaScript from directly reading the cookie.

Example:

    HttpOnly

This reduces the impact of certain XSS attacks attempting to steal session credentials.

### Secure

Authentication cookies should use:

    Secure

in production.

This ensures browsers send the cookie only over HTTPS.

### SameSite

Use:

    SameSite=Lax

as the default for normal session authentication.

Use:

    SameSite=Strict

when the application works correctly with stricter cross-site behavior.

Use:

    SameSite=None

only when cross-site cookies are genuinely required.

`SameSite=None` must be combined with:

    Secure

### Cookie Scope

Cookies should be restricted as narrowly as practical.

Use:

- Appropriate `Path`
- Appropriate `Domain`
- Reasonable expiration time

Avoid unnecessarily broad cookie scopes.

---

# 7. CSRF Protection

If authentication uses cookies, CSRF protection must be considered because browsers automatically send cookies with qualifying requests.

State-changing authenticated requests include:

    POST
    PUT
    PATCH
    DELETE

The backend should use appropriate CSRF defenses when required.

Possible approach:

    SameSite cookies
          +
    CSRF token
          +
    Origin / Referer validation

For the current public contact endpoint, CSRF protection is not required merely because cookies may exist.

If the endpoint later becomes authenticated or depends on an authentication cookie, the requirement must be reassessed.

---

# 8. Authentication

The current portfolio does not require authentication.

Do not introduce authentication simply because the application has a backend.

Authentication should only be introduced when protected functionality requires it.

If authentication is added:

- Passwords must never be stored in plaintext.
- Use Argon2id or bcrypt for password hashing.
- Use secure server-side sessions.
- Use generic authentication error messages.
- Apply login rate limiting.
- Expire sessions.
- Implement secure logout.
- Invalidate sessions when necessary.
- Protect authenticated routes with middleware.

---

# 9. Authorization

Authentication and authorization are separate concepts.

Authentication answers:

    Who are you?

Authorization answers:

    Are you allowed to perform this action?

Every future protected endpoint must enforce authorization on the backend.

Do not rely on:

- React route guards alone
- Hidden frontend buttons
- Client-side role checks
- User-controlled role values
- Query parameters containing authorization information

---

# 10. Input Validation

All incoming data must be validated by the Go backend.

Validation should include:

- Required fields
- Maximum length
- Minimum length where appropriate
- Email format
- Allowed values
- IDs
- Content type
- Request body size

The frontend may perform validation for user experience.

However:

    Frontend validation ≠ security validation

The backend remains responsible for security validation.

---

# 11. Contact Form Security

The contact endpoint is:

    POST /api/contact

This is likely to be the primary public abuse target.

The backend should protect the endpoint using:

- Strict validation
- Request body limits
- Rate limiting
- Content-Type validation
- Email validation
- Message length limits
- Spam protection if necessary

Suggested initial limits:

    Name:       1–100 characters
    Email:      3–254 characters
    Subject:    1–200 characters
    Message:    1–5000 characters

These limits can be adjusted during implementation.

---

# 12. Rate Limiting

Public endpoints should have appropriate rate limits.

At minimum:

    POST /api/contact

must be rate limited.

The exact rate limit should be determined during implementation.

For example:

    Multiple requests
           ↓
    Same IP/client
           ↓
    Rate limit exceeded
           ↓
    429 Too Many Requests

The application should not expose internal rate-limiter details.

---

# 13. SQL Injection Prevention

All SQLite queries must use parameterized statements or prepared statements.

Never construct SQL queries using direct string concatenation with user input.

Unsafe:

    "SELECT * FROM contacts WHERE email = '" + email + "'"

Safe:

    Query(
        "SELECT * FROM contacts WHERE email = ?",
        email,
    )

The repository layer is responsible for safe database access.

---

# 14. Database Security

SQLite must never be directly accessible from the browser.

The architecture must remain:

    Browser
       ↓
    Go API
       ↓
    SQLite

The SQLite database file must:

- Remain on the server.
- Never be placed inside the React `public/` directory.
- Never be served as a static file.
- Not be committed to Git.
- Have appropriate filesystem permissions.
- Be backed up securely when production data matters.

Recommended `.gitignore` entries:

    *.db
    *.sqlite
    *.sqlite3

---

# 15. Sensitive Data

The portfolio should collect the minimum information required.

The contact form should collect:

- Name
- Email
- Subject
- Message

Do not collect unnecessary sensitive information.

Avoid storing:

- Passwords
- Authentication tokens
- API keys
- Private credentials
- Payment information
- Unnecessary personal information

---

# 16. Environment Variables

Secrets and environment-specific configuration must not be hard-coded.

Use environment variables where appropriate.

Example:

    PORT=8080
    DATABASE_URL=./portfolio.db
    FRONTEND_URL=http://localhost:5173

Production secrets must never be committed to Git.

Provide:

    .env.example

Do not commit:

    .env

---

# 17. CORS

The Go API must configure CORS explicitly.

Development may allow:

    http://localhost:5173

Production should allow only the real frontend origin.

Do not use unrestricted CORS such as:

    Access-Control-Allow-Origin: *

for authenticated APIs.

CORS is not an authentication mechanism.

It is an additional browser security control.

---

# 18. HTTP Security Headers

The production server should provide appropriate security headers.

Recommended headers include:

    Content-Security-Policy
    X-Content-Type-Options: nosniff
    Referrer-Policy
    Permissions-Policy
    Strict-Transport-Security

A Content Security Policy should be introduced carefully so it does not break:

- Scripts
- Fonts
- Images
- API requests
- Other legitimate resources

Avoid:

    unsafe-eval

unless technically required.

Minimize:

    unsafe-inline

where practical.

---

# 19. XSS Protection

The application must prevent untrusted content from becoming executable HTML or JavaScript.

React's normal rendering behavior provides escaping for ordinary text content.

Do not use:

    dangerouslySetInnerHTML

with untrusted content.

If raw HTML is ever required, it must be sanitized using a maintained HTML sanitization library before rendering.

---

# 20. Error Handling

Production errors must not expose:

- Stack traces
- SQL queries
- Database paths
- Environment variables
- File paths
- Internal service names
- Authentication details

The API should return safe JSON.

Example:

    {
      "success": false,
      "message": "Unable to process your request."
    }

Detailed errors should be logged server-side rather than returned to users.

---

# 21. Logging

Logs should help diagnose problems without exposing secrets.

Never log:

- Passwords
- Session tokens
- Authentication cookies
- API keys
- Database credentials
- Sensitive request bodies

Useful non-sensitive logs may include:

- HTTP method
- Endpoint
- Status code
- Request duration
- Error identifier

Production logs should avoid unnecessary personal information.

---

# 22. Request Size Limits

The backend should limit request body sizes.

This protects the server from excessively large requests.

The contact endpoint should have a relatively small request limit because it only accepts short text fields.

---

# 23. HTTP Method Restrictions

Endpoints should accept only the HTTP methods they require.

Examples:

    GET /api/projects
    GET /api/projects/:id
    GET /api/skills
    POST /api/contact

Unexpected methods should return:

    405 Method Not Allowed

where appropriate.

---

# 24. Content-Type Validation

JSON endpoints should require:

    Content-Type: application/json

Requests using unexpected content types should be rejected where appropriate.

---

# 25. ID Validation

Endpoints such as:

    GET /api/projects/:id

must validate the ID before querying SQLite.

For example:

    /api/projects/abc

should produce a controlled client error if the API expects an integer ID.

The server must not expose database errors.

---

# 26. Authentication Token Storage

If token-based authentication is introduced later, avoid storing sensitive authentication tokens in:

    localStorage

when a secure HttpOnly cookie-based session can be used.

Preferred future architecture:

    Browser
       ↓
    HttpOnly Secure Session Cookie
       ↓
    Go Authentication Middleware
       ↓
    Protected API

The frontend should determine authentication state through API responses rather than reading the session cookie directly.

---

# 27. Session Management

If sessions are introduced:

- Use unpredictable session IDs.
- Store sessions securely.
- Set session expiration.
- Rotate sessions after authentication where appropriate.
- Invalidate sessions on logout.
- Consider session invalidation after sensitive account changes.
- Do not rely solely on client-controlled information.

---

# 28. Password Security

If an admin account is introduced:

Passwords must:

- Never be stored in plaintext.
- Never be logged.
- Never be returned through an API.
- Be hashed using Argon2id or bcrypt.

Password reset functionality must use:

- Short-lived tokens
- Unpredictable tokens
- One-time-use tokens where appropriate

---

# 29. Dependency Security

Keep project dependencies reasonably up to date.

Before production deployment:

- Review npm dependencies.
- Review Go dependencies.
- Remove unused packages.
- Run security audits where practical.
- Avoid unnecessary third-party dependencies.

Frontend:

    npm audit

Go dependencies should also be reviewed regularly.

---

# 30. Git Security

Never commit:

    .env
    *.db
    *.sqlite
    *.sqlite3
    Private keys
    API keys
    Session secrets
    Passwords
    Credentials

Recommended `.gitignore`:

    .env
    .env.*
    !.env.example

    *.db
    *.sqlite
    *.sqlite3

If a secret is accidentally committed, deleting the file is not enough.

The secret should be considered compromised and rotated.

---

# 31. File Uploads

The current portfolio does not require file uploads.

If file uploads are introduced later, they must be treated as a separate security feature.

Requirements should include:

- File size limits
- Allowed MIME types
- Extension validation
- Generated storage names
- Storage outside executable web directories
- Malware scanning where appropriate
- No direct execution of uploaded files
- Authorization checks
- Safe download handling

Do not introduce uploads unless required.

---

# 32. API Security

API responses should:

- Use correct HTTP status codes.
- Return predictable JSON.
- Avoid exposing internal implementation details.
- Validate all input.
- Enforce authorization on protected routes.
- Rate-limit abuse-prone endpoints.

Suggested status codes:

    200 OK
    201 Created
    400 Bad Request
    401 Unauthorized
    403 Forbidden
    404 Not Found
    405 Method Not Allowed
    409 Conflict
    429 Too Many Requests
    500 Internal Server Error

---

# 33. Contact Data Privacy

Contact messages may contain personal information.

The backend should:

- Store only required fields.
- Restrict database access.
- Never expose contact messages through public API endpoints.
- Avoid returning stored messages unnecessarily.
- Establish a retention policy if the portfolio becomes production-facing.

There should NOT be a public endpoint such as:

    GET /api/contacts

unless a future authenticated admin feature specifically requires it.

---

# 34. Admin Features

An admin dashboard is currently excluded from the product.

If an admin dashboard is added later, it must introduce:

- Authentication
- Authorization
- Secure sessions
- CSRF protection where required
- Rate limiting
- Protected admin routes
- Appropriate audit logging

Admin endpoints must never rely on frontend route protection alone.

---

# 35. Production Security Baseline

The production deployment should aim for:

    HTTPS
    Secure cookies
    HttpOnly authentication cookies
    SameSite=Lax
    Strict CORS
    Content-Security-Policy
    X-Content-Type-Options: nosniff
    Referrer-Policy
    Permissions-Policy
    HSTS
    Rate limiting
    Request size limits
    Input validation
    Parameterized SQL

The exact configuration depends on the hosting environment and reverse proxy.

---

# 36. Security Testing

Before production deployment, test:

## Frontend

- Navigation
- Forms
- Invalid input
- API errors
- Keyboard interaction
- Mobile behavior

## Backend

Test:

- Invalid JSON
- Missing fields
- Oversized requests
- Invalid IDs
- Unsupported methods
- Wrong content types
- SQL injection attempts
- Repeated contact submissions
- CORS behavior
- Authentication if introduced
- Authorization if introduced
- Cookie attributes if authentication is introduced

---

# 37. Security Checklist

Before production:

- [ ] HTTPS enabled
- [ ] CORS restricted
- [ ] Input validation implemented
- [ ] SQL queries parameterized
- [ ] SQLite file protected
- [ ] Database excluded from Git
- [ ] `.env` excluded from Git
- [ ] `.env.example` provided
- [ ] Contact rate limiting implemented
- [ ] Request body limits implemented
- [ ] Safe error responses implemented
- [ ] Security headers configured
- [ ] Dependencies reviewed
- [ ] No secrets in source code
- [ ] No sensitive data in logs
- [ ] Contact data not publicly exposed
- [ ] Production build tested

If authentication is introduced:

- [ ] Password hashing implemented
- [ ] Secure session management implemented
- [ ] HttpOnly cookies used
- [ ] Secure cookies enabled in production
- [ ] SameSite configured
- [ ] CSRF protection implemented where required
- [ ] Authorization enforced server-side
- [ ] Login rate limiting implemented
- [ ] Logout invalidates sessions
- [ ] Session expiration configured

---

# 38. Security Development Process

Security should be considered during development rather than added after the application is completed.

For every new feature, ask:

1. What data does this feature accept?
2. Who can access it?
3. What happens if the input is malicious?
4. Does it require authentication?
5. Does it require authorization?
6. Could it expose personal information?
7. Does it introduce a new database query?
8. Does it require cookies?
9. Does it require CSRF protection?
10. Can the feature be abused automatically?

Every new backend endpoint should have a documented security expectation before implementation.

---

# 39. Current Security Decision

For the initial portfolio release:

    Authentication: Not required
    Cookies: Not required
    Admin dashboard: Excluded
    Public contact form: Yes
    Rate limiting: Required
    HTTPS: Required in production
    CORS: Restricted
    SQLite: Server-side only
    SQL parameterization: Required
    CSRF: Re-evaluate if authenticated cookie-based features are introduced

Cookies should therefore NOT be added simply for the sake of having cookies.

They should be introduced when the application gains a feature that actually benefits from secure session state.

The most likely future use is:

    Admin authentication
          ↓
    Secure session
          ↓
    HttpOnly + Secure + SameSite cookie

---

# 40. Future Authentication Architecture

If a protected admin area is introduced:

    Browser
       |
       | HTTPS
       ↓
    React
       |
       | authenticated request
       ↓
    HttpOnly + Secure + SameSite Cookie
       |
       ↓
    Go Authentication Middleware
       |
       ↓
    Authorization
       |
       ↓
    Admin Handler
       |
       ↓
    Service
       |
       ↓
    Repository
       |
       ↓
    SQLite

The frontend should never need to read the authentication cookie directly.

---

# 41. Final Security Standard

The portfolio should remain simple, but simplicity must not mean ignoring security.

The initial implementation should prioritize:

    Validation
    Rate Limiting
    SQL Safety
    CORS
    HTTPS
    Security Headers
    Safe Error Handling
    Database Protection
    Secret Management

Cookies and authentication should be introduced only when required by future protected functionality.

This keeps the first release secure without adding unnecessary complexity.
