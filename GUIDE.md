# Go Project Structure Guidebook

A personal reference for structuring Go backend projects — how to find the domain
nouns, how to lay out `internal/`, and what goes in each file. Written against
this project (a system-design study app) but meant to be reused for any future
Go service.

---

## 1. Top-level layout

```
project-root/
├── cmd/
│   └── api/
│       └── main.go        # wiring only — no logic
├── internal/
│   ├── <noun>/             # one package per domain noun (see §2)
│   ├── platform/           # infra: db, config, http server, logging
│   └── apperrors/          # shared error types
├── pkg/                    # ONLY if something is meant to be imported by other repos
├── go.mod
└── go.sum
```

**Rules for the three top-level dirs:**

- `cmd/<binary>/main.go` — reads config, constructs concrete dependencies
  (db connection, loggers), injects them into services, starts the server.
  Should be short enough to read top to bottom in one sitting. If `main.go`
  has an `if` statement implementing a business rule, that rule leaked out of
  a service — move it back.
- `internal/` — default location for everything. Go's compiler enforces that
  nothing outside this module can import it, so this is where 95% of the code
  lives.
- `pkg/` — do not create this preemptively. Only add it the day you actually
  have a second consumer (another repo, a public SDK) for some piece of code.
  "Might reuse it later" is not that day.

---

## 2. Finding the domain nouns

Do this before writing any code, and redo it whenever a new feature is added.

### 2.1 Extract nouns from plain-English feature sentences

Write one sentence per feature, in language a non-programmer would use, then
underline the nouns.

> "A user logs in and sees the modules they last opened, can browse all
> modules, and see their level in each."

Nouns: user, module, level, "last opened" (a fact, not yet a noun).

Nouns that recur across multiple feature sentences, and that a domain expert
would say out loud unprompted, are your package candidates. If a candidate
name is something no stakeholder would ever say ("Orchestrator", "Manager",
"Handler", "Processor") it's a technical role in a noun costume — not a real
domain concept, and not a good package name.

### 2.2 Does it deserve to be its own package?

Ask: **does this thing have its own lifecycle, identity, and rules,
independent of its parents?**

- `Module` — created, updated, has an ID, has levels → yes, real noun.
- "last opened" — on its own it's just a timestamp field on a relationship
  between User and Module. It only becomes its own package (`progress`) once
  it accumulates independent rules: "progress can be reset," "progress has a
  completion percentage," "progress emits an event on level-up." If all it
  does is store one timestamp, it's a field on an existing struct, not a
  package.

### 2.3 Domain nouns vs. capability nouns

Two different kinds of package, both legitimate, kept separate because they
change for different reasons:

- **Domain nouns** — come from the business itself (`user`, `module`,
  `level`, `progress`). Change when the product changes.
- **Capability nouns** — things the system needs regardless of business
  specifics (`auth`, `session`, `notification`, `search`, `ratelimit`).
  Change when infrastructure/technical approach changes.

Never let a capability concern (e.g. how passwords are hashed) live inside a
domain package (e.g. `user`) — put it in its own package (`auth`) that
`user` depends on. This is why `auth` and `user` are separate packages in
this project even though "a user logs in" reads like one feature.

### 2.4 Splitting an overloaded noun

If a package's methods stop touching each other's data — e.g. `user`
accumulates password-reset-token logic, billing logic, and avatar-upload
logic that never call into one another — that's really three packages
(`auth`, `billing`, `profile`) that got merged by accident. Split along the
line where the methods stop sharing data.

### 2.5 The tiebreaker: co-change, not taste

When unsure whether something is its own package or belongs inside an
existing one:

> When this thing changes, does it drag changes into another package too, or
> does it stand alone?

Always-together → same package. Independently editable and testable →
separate package. This is checkable against real history (`git log -p` on
the relevant files) rather than a matter of opinion, and it's the rule to
fall back on whenever intuition and "conceptual purity" disagree.

Expect to get this wrong at least once early on and need to merge or split a
package a few months in — that's normal and cheap (`gopls`/`gorename` handle
the mechanical part). Don't over-invest in a perfect split on day one.

### 2.6 Composition packages

Some packages don't represent a domain noun at all — they compose several
domains for one screen/use case. Example: `dashboard` isn't a noun with its
own lifecycle; it's a read-only view that pulls from `user`, `module`, and
`progress`. Give it a thin package (usually just `service.go` + `handler.go`,
no `model.go`/`repository.go` of its own) that calls into the other domains'
services. Never let this orchestration logic bleed into one of the
underlying domain packages instead.

---

## 3. Inside a domain package: what each file does

Default shape for a domain package (e.g. `internal/module/`):

```
module/
├── model.go        # types + invariants
├── repository.go   # persistence
├── service.go       # business logic / orchestration
└── handler.go       # HTTP transport
```

| File | Owns | Must NOT contain | Depends on |
|---|---|---|---|
| `model.go` | Structs for this noun, and validation/invariant methods on the struct itself (e.g. `(l Level) IsValidDifficulty() bool`) | Any import of `database/sql`, `net/http`, or any other package in this project | nothing internal |
| `repository.go` | All persistence: SQL/ORM calls, mapping DB rows ↔ structs from `model.go` | Business rules, decisions ("is this user allowed to..."), HTTP concerns | `model.go` |
| `service.go` | Business logic: enforces rules, orchestrates one or more repository calls, calls into *other domains'* services (e.g. `module.service` calling `progress.service`) | Raw SQL, `http.Request`/`http.ResponseWriter`, JSON encoding | `model.go`, `repository.go`, other domains' service interfaces |
| `handler.go` | HTTP-layer only: decode request → call service → encode response, map domain errors to HTTP status codes | `if` statements implementing business rules, direct DB access | `service.go` |

**Layer-swap test.** Each file should be replaceable without the others
noticing:

- Swap Postgres for an in-memory fake in tests → only `repository.go`
  changes, because `service.go` talks to a small interface (defined in
  `service.go`, satisfied by the concrete `repository.go` struct), not the
  concrete type.
- Swap REST for gRPC → only `handler.go` (or a new `grpc_handler.go`)
  changes; `service.go` never imports a transport package.

**Where to catch a misplaced concern:**

- `handler.go` has business logic (e.g. "user can only see levels they've
  unlocked") → move it into `service.go`.
- `service.go` has a raw SQL string → move it into `repository.go`.
- `model.go` imports `database/sql` or `net/http` → something transport- or
  persistence-specific leaked into the pure type definition; move it out.

**When to split further.** Only when a file does two distinct jobs, not
just because it crossed a line-count threshold. E.g. once `module/service.go`
grows both catalog-browsing logic and access-control logic, split into
`service.go` + `access_control.go`. A trivial noun (a read-only reference
table with no writes) can collapse to a single file — don't force the
four-file shape where it adds nothing.

---

## 4. `internal/platform/`: infrastructure, not business logic

```
platform/
├── database/
│   └── postgres.go     # connection pool setup, migrations trigger
├── config/
│   └── config.go        # env var / file parsing into a typed Config struct
├── httpserver/
│   ├── router.go        # route table: wires each domain's handler to a path
│   └── server.go         # http.Server lifecycle (start/shutdown)
└── logger/
    └── logger.go          # logger construction
```

Rule: **domain packages may import `platform/`, `platform/` must never
import a domain package.** If `platform/httpserver/router.go` needs to know
about `module.Handler`, that's fine (it's wiring, lives close to `main.go`
in spirit) — but `platform/database/postgres.go` should never reference
`module.Module`. Keep persistence-of-a-specific-struct code in that struct's
own `repository.go`, not in `platform/database/`.

---

## 5. `internal/apperrors/`: shared error vocabulary

One file, `errors.go`, defining sentinel/typed errors used across domains
(e.g. `ErrNotFound`, `ErrUnauthorized`, `ErrConflict`). Domain
`repository.go` files return these instead of raw driver errors; `handler.go`
files map them to HTTP status codes in one place (usually a shared
`apperrors.ToHTTPStatus(err) int` helper). This is what lets every domain's
`handler.go` do the same one-line error mapping instead of reinventing
status-code logic per package.

---

## 6. Authentication vs. authorization vs. middleware

These three are easy to conflate but split cleanly:

- **Authentication** ("who is this?") — verifying credentials, issuing and
  validating session tokens. This is real domain-ish logic with its own
  rules and lives in a proper package: `internal/auth/` (`service.go` with
  `Login(creds) (Token, error)` and `ValidateToken(token) (Claims, error)`,
  `repository.go` for credential lookups, `handler.go` for `POST /login` /
  `POST /logout`). It is a **capability noun** (§2.3), not a domain noun.
- **Middleware** — the *mechanism* that applies authentication uniformly to
  every protected route. It should stay dumb: call `auth.ValidateToken` on
  the incoming request, and either attach the resulting `Claims`/`UserID` to
  the request context and pass through, or short-circuit with 401. It does
  not decide what the user is allowed to do — only whether they are who
  they claim to be. It can live in `internal/auth/middleware.go` (keeps the
  package self-contained) or in `internal/platform/httpserver/middleware.go`
  (keeps route wiring visible in one place) — either is fine, this is a
  judgment call, not a rule.
- **Authorization** ("can this user do X?") — fine-grained, per-action
  permission checks (e.g. "can this user edit this module," "is this user's
  level high enough to unlock this content"). This is domain-specific and
  belongs inside the relevant domain's `service.go` (e.g. `module.service`,
  `progress.service`), never centralized in the middleware or in `auth`.

**Why not centralize authorization in the middleware or in `auth`:** the
middleware/`auth` package would then need to know the rules of every domain
in the system, turning it into a god-object and defeating the entire point
of splitting packages by noun (§2). Middleware answers "is there a valid
session at all" (a coarse, uniform gate); each domain service answers
"is this particular action allowed for this particular user" (a rule
specific to that noun).

---

## 7. Sample complete project structure

A fully fleshed-out version of this project once `user`, `module`,
`progress`, and `dashboard` exist alongside `auth`:

```
system-design-app/
├── cmd/
│   └── api/
│       └── main.go            # loads config, builds db/logger, constructs
│                                # each domain's repo→service→handler, builds
│                                # the router via platform/httpserver, starts
│                                # the server. No business logic.
│
├── internal/
│   ├── auth/                   # CAPABILITY package — login/session/token logic
│   │   ├── model.go             # Token, Claims, Credentials types
│   │   ├── repository.go        # credential lookups (hashed password, etc.)
│   │   ├── service.go            # Login(), ValidateToken(), Logout()
│   │   ├── handler.go            # POST /login, POST /logout
│   │   └── middleware.go          # RequireAuth(): validates token, attaches
│   │                               # Claims to request context, 401 on failure
│   │
│   ├── user/                    # DOMAIN noun — account/profile data
│   │   ├── model.go              # User struct + invariants
│   │   ├── repository.go          # CRUD against the users table
│   │   ├── service.go              # profile update rules, calls auth for
│   │   │                            # credential changes if needed
│   │   └── handler.go              # GET/PATCH /users/:id
│   │
│   ├── module/                  # DOMAIN noun — the catalog of study modules
│   │   ├── model.go              # Module, Level structs + invariants
│   │   ├── repository.go          # module/level queries
│   │   ├── service.go              # list-all, get-by-id, per-action
│   │   │                            # authorization checks (§6) live here
│   │   └── handler.go              # GET /modules, GET /modules/:id
│   │
│   ├── progress/                # DOMAIN noun — user↔module relationship:
│   │   │                          # last-opened, completion, unlocked levels
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── service.go              # RecordOpened(), MarkComplete(), rules
│   │   │                            # around resets/unlocks
│   │   └── handler.go              # POST /progress, GET /progress/:moduleID
│   │
│   ├── dashboard/                # COMPOSITION package — no model/repository
│   │   │                          # of its own; reads from user, module,
│   │   │                          # progress services for one screen
│   │   ├── service.go              # BuildDashboard(userID) — calls into
│   │   │                            # user.Service, module.Service,
│   │   │                            # progress.Service and assembles the view
│   │   └── handler.go              # GET /dashboard
│   │
│   ├── platform/                 # infra — technology-specific, business-agnostic
│   │   ├── database/
│   │   │   └── postgres.go         # connection pool, migration trigger
│   │   ├── config/
│   │   │   └── config.go            # env/file parsing into typed Config
│   │   ├── httpserver/
│   │   │   ├── router.go             # gin.Engine route table; wires each
│   │   │   │                          # domain's handler + auth middleware
│   │   │   └── server.go              # server start/graceful shutdown
│   │   └── logger/
│   │       └── logger.go               # logger construction
│   │
│   └── apperrors/
│       └── errors.go              # ErrNotFound, ErrUnauthorized, ErrConflict,
│                                    # ToHTTPStatus(err) helper
│
├── pkg/                          # absent until a second consumer actually
│                                  # needs to import something from this repo
├── go.mod
└── go.sum
```

Quick-reference for what belongs where, one line each:

- **`cmd/api/main.go`** — wiring, nothing else.
- **`<domain>/model.go`** — the noun's shape and self-contained invariants.
- **`<domain>/repository.go`** — talking to the database for that noun only.
- **`<domain>/service.go`** — that noun's business rules and authorization
  decisions; orchestrates its own repository and other domains' services.
- **`<domain>/handler.go`** — HTTP in/out for that noun; no decisions.
- **`auth/middleware.go`** — identity check only, no per-domain rules.
- **`dashboard/`** (or any composition package) — orchestration across
  domains for one screen/use case; never owns persisted data of its own.
- **`platform/*`** — infrastructure that doesn't know the business exists.
- **`apperrors/errors.go`** — the one shared error vocabulary every domain
  speaks, so error→HTTP-status mapping is written once.

---

## 8. Checklist for starting or extending a project

1. Write feature sentences in plain English; underline nouns (§2.1).
2. For each noun, ask the lifecycle/identity question (§2.2) — reject
   attribute-only nouns.
3. Separate domain nouns from capability nouns (§2.3).
4. For each accepted noun, create `model.go` (start here), then add
   `repository.go`, `service.go`, `handler.go` as the feature needs each
   layer — don't scaffold all four before there's code to put in them.
5. If a feature spans multiple domains, give it a thin composition package
   instead of shoving it into one of the underlying domains (§2.6).
6. Put anything technology-specific and business-agnostic in `platform/`.
7. Wire everything in `cmd/api/main.go` last.
8. When a package or file feels wrong, apply the co-change tiebreaker
   (§2.5) and the layer-swap test (§3) before restructuring.

---

## 9. Current project mapping (system-design-app)

As of this writing the project has:

- `internal/auth/` — `auth.go`, `handlers.go`. This is a **capability**
  package (login/session logic), not a domain noun. As it grows, expect it to
  split along the file rules in §3: move HTTP decode/encode into
  `handlers.go` (already separate — good), keep token/session logic in
  `auth.go`, and if password hashing or persistence of credentials grows
  non-trivial, pull a `repository.go` out of `auth.go`. Per §6, add a
  `middleware.go` here that validates the token and attaches `Claims` to the
  request context — it should not contain any per-domain authorization
  rules; those belong in `module.service` / `progress.service` etc.

Still to add as real features arrive: `user/`, `module/`, `progress/`,
`dashboard/` (composition), and `platform/` for the shared db/config/server
setup currently likely sitting in `cmd/api/main.go`.
