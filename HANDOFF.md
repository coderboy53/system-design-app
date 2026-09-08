# System Design App — session handoff

**Temporary scratch doc.** Captures the design analysis from a Claude session on 2026-09-01 so other
sessions can catch up. Delete once these decisions land in `topics/README.md`, `GUIDE.md`, or real
code.

> **Frontend session:** read [Scope split](#scope-split), [Part 1 — Content layer](#part-1--content-layer-shared-truth),
> and [Part 3 — Frontend](#part-3--frontend-delegated-session). Part 2 is backend-only; skim it for
> context on where your data comes from, but don't act on it.

## Scope split

| Who | Scope |
|---|---|
| Soham (repo owner) | `backend/` — Go/Gin API, Postgres schema, content ingestion |
| Delegated Claude session | `frontend/` — React + Vite app |

Nobody but Soham edits `backend/`. Nobody but the frontend session edits `frontend/`. The content
under `topics/` is shared truth — changes there need agreement.

---

## Decisions made this session

| # | Question | Decision |
|---|---|---|
| 1 | How is module ordering stored? | New `module_order` frontmatter field on module index files, ingested into a DB column. **Implemented** — see below. |
| 2 | Derive module order from `manifest.json` array position instead? | No. Array position is implicit and one `json.Unmarshal` into a map away from silently scrambling the curriculum. Declare explicitly, sort in SQL. |
| 3 | Separate `topics` package from `modules`? | No. One package, kept as `internal/modules/` (not renamed to `content/`) — high level this domain is about modules, and a module has no behavior without its topics, and ingestion parses both in one pass. |
| 4 | One model or two? | Two structs, `Module` and `Topic`, in the same package. |
| 5 | `prerequisites` at module or topic level? | Topic level. Already present and already complete — no content change needed. Module-level prereqs are a derived query. |
| 6 | `pq.StringArray` or `[]string`? | `[]string` in the struct, `pq.Array(&f)` at the scan site — keeps the driver out of the models. |
| 7 | Add `body` to the struct when `manifest.json` has none? | Yes. The body's source of truth is the `.md` file; the manifest is a derived index. |
| 8 | Where does ingestion read from? | Walk `topics/**/*.md`. Everything in `manifest.json` is derivable from the files; the manifest is a cross-check and a frontend fixture. |
| 9 | Module page layout? | README prose at top, topic list below on scroll. |

## Current state of the repo

**Works:** Google OAuth login end to end (popup + Redis-backed session cookie), Vite dev proxy.

**Does not exist yet:** any content API, DB schema, migration tooling, or ingestion.
`backend/internal/modules/` is three near-empty stub files (`models.go` has `Module` with `Id` and
`Title`). No `.sql` files anywhere. `manifest.json` is hand-maintained — there is no generator script.

Everything in Part 2 and the API contract in Part 3 is **agreed design, not shipped code.**

---

## Part 1 — Content layer (shared truth)

All learning content is committed markdown under `topics/` — **12 modules, 75 entries** — adapted from
The System Design Primer, CC BY 4.0. See `topics/ATTRIBUTION.md`; **attribution must stay visible in
the UI.** `topics/manifest.json` (69 KB) is a machine-readable index of every entry.

Two levels only: a **module** is a directory; **topics** are the `.md` files in it.

### Frontmatter contract

Documented in `topics/README.md` → "Content conventions".

```yaml
id: cache-locations          # stable slug, THE join key (filenames may move)
title: Where to Cache
module: caching              # parent module id
module_order: 60             # index files (*/README.md) only — sorts modules
order: 20                    # sorts topics within the module
est_minutes: 5
timelines: [short, medium, long]
tags: [caching, architecture, redis, memcached]
prerequisites: [caching-overview]
related: [cdn, reverse-proxy, key-value-store, what-to-cache]
source: https://github.com/donnemartin/...
```

`manifest.json` entries add four derived fields: `path`, `word_count`, `sections` (the `##` headings)
and `has_self_check`.

### Ordering

- `module_order` sorts **modules**; `order` sorts **topics within** a module.
- Both **spaced by 10** so an item can be inserted without renumbering. Gaps are intentional — the
  values are sort keys, not indices.
- `order: 0` is reserved for a module's own `README.md`. `order: 900` parks an entry last (only
  `ATTRIBUTION.md`).
- **Always sort by `(module_order, order)`.** Never rely on `manifest.json` array position, never
  sort alphabetically.

`module_order` was added this session to all 12 module index files (line 5 of each frontmatter,
right after `module:`), mirrored into `manifest.json` under both `modules[]` and the matching index
entry, and documented in `topics/README.md`. The diff was purely additive — 44 insertions, 0
deletions — because array position and `module_order` agree today.

### The modules

| `module_order` | id | title | entries |
|---|---|---|---|
| 0 | `root` | Root | 2 |
| 10 | `study-guide` | Study Guide | 6 |
| 20 | `fundamentals` | Fundamentals | 7 |
| 30 | `networking` | Networking & Edge | 9 |
| 40 | `application` | Application Layer | 5 |
| 50 | `data` | Data | 13 |
| 60 | `caching` | Caching | 5 |
| 70 | `security` | Security | 2 |
| 80 | `appendix` | Appendix | 9 |
| 90 | `exercises` | Exercises | 1 |
| 100 | `exercises/system-design` | System Design Exercises | 9 |
| 110 | `exercises/object-oriented-design` | Object-Oriented Design Exercises | 7 |

- **`root` is not a module card.** It's `topics/README.md` — landing-page copy.
- **8 top-level cards:** Study Guide → Fundamentals → Networking & Edge → Application Layer → Data →
  Caching → Security → Appendix.
- **The `exercises` tree (rows 90/100/110) is out of scope.** This is a theory-only app; those 3
  modules / 17 topics are excluded from `topics/db-export/` and are not modelled as modules.
  Interactive exercises, if ever built, are a separate unattached feature with their own shape — so
  **`modules` is a flat list**, no parent/child nesting, and no card drills into sub-modules.
- Common confusion: **`dns` is not a top-level module.** It's a topic at `order: 10` inside
  `networking`, with `cdn`, `load-balancer`, `reverse-proxy`, `http`, `tcp-and-udp`, `rpc`, `rest`.

### The prerequisite graph

`prerequisites` and `related` reference other topic `id`s, cross-module. Verified this session:

- All **26 distinct prerequisite values resolve to a valid topic `id`** — zero dangling refs.
- The graph is a **DAG** (no cycles), so it is safe to topologically sort for a "recommended path".
- **24 edges cross module boundaries**, e.g. `data/nosql → fundamentals/cap-theorem`,
  `networking/load-balancer → fundamentals/availability-patterns`, and 15 exercises →
  `study-guide/interview-approach`.
- **12 of 63 leaf topics have no prerequisites** (`http`, `security-basics`, `powers-of-two`, the
  reference/blog pages…), so any gating UI needs a clean "no prerequisites" state.

**Module-level prerequisites already exist implicitly**, because a module's index README is itself a
topic with its own `prerequisites`:

```
networking/README.md   prerequisites: [fundamentals]
application/README.md  prerequisites: [networking]
data/README.md         prerequisites: [fundamentals]
caching/README.md      prerequisites: [data]
```

⚠️ **The field is polymorphic by accident.** For all 9 in-scope modules the index topic's `id` is
byte-identical to the module `id` (module `caching`, index topic `caching`), which is why
`prerequisites: [fundamentals]` reads as a module but resolves as a topic. (The only two manifest
entries that diverged — `exercises/system-design` → `system-design-exercises` and
`exercises/object-oriented-design` → `ood-exercises` — are in the excluded exercises tree.)

So: **resolve prerequisites against topic ids only**, and don't add a constraint that assumes
`module.id == index_topic.id` — the collision is a coincidence of naming, not a rule. Four leaf
topics point at module-id lookalikes (`short-timeline`/`medium-timeline`/`long-timeline` →
`study-guide`, `caching-overview` → `data`); they resolve fine as topic ids and need no cleanup.

### Bodies and assets

- **347 KB of markdown across 75 files.** Median 3.1 KB, largest 19.5 KB (`exercise-mint`, 1992
  words). Small enough to store rendered in Postgres without thinking about it.
- **9.6 MB under `topics/assets/`** — serve as static files, never from the DB.
- `manifest.json` has **no `body` field**, by design. `path` points at the file that holds it.

---

## Part 2 — Backend design (Soham)

### Package layout

One package, not a `modules` + `topics` split. Rationale: strip the topics out and `Module` is
`{id, title, module_order, count}` — a grouping row, not a domain. A module's index page *is* a topic.
Ingestion produces both from one frontmatter parse. Every read path joins both.

```
backend/internal/modules/
    models.go     Module, Topic
    ingest.go     walk topics/, parse frontmatter, upsert   ← the shared parser
    modules.go    ListModules, GetModule, GetTopic, Path(timeline), Next(topicID)
    handlers.go   GET /api/modules, /api/modules/:id, /api/topics/:id
```

Matches the existing `auth`/`user` convention (`models.go` / `<domain>.go` / `handlers.go`). Kept the
existing `modules` package name rather than renaming to `content` — at a high level this domain is
about modules, even though the package also holds `Topic`.

**If you split anyway:** keep the dependency one-directional, `modules → topics`, with
`Topic.ModuleID` as a plain `string` (which is what `module: fundamentals` already is). Put a
`*Module` on `Topic` and you have an import cycle with no clean fix short of a third types package.

**Later features that *should* be their own packages:** `internal/progress/`, `internal/flashcards/`,
exercise submissions. Those are separate domains with user-scoped writes; they import `modules`.

### Models

```go
type Module struct {
	Id          string  `json:"id"                 db:"id"`
	Title       string  `json:"title"              db:"title"`
	ModuleOrder int     `json:"module_order"       db:"module_order"`
	TopicCount  int     `json:"topic_count"        db:"topic_count"`
	Overview    string  `json:"overview,omitempty" db:"overview"`   // rendered README, table stripped
	Topics      []Topic `json:"topics,omitempty"`
}

type Topic struct {
	Id            string   `json:"id"             db:"id"`
	Title         string   `json:"title"          db:"title"`
	ModuleId      string   `json:"module_id"      db:"module_id"`
	Order         int      `json:"order"          db:"sort_order"`
	Path          string   `json:"path"           db:"path"`
	EstMinutes    int      `json:"est_minutes"    db:"est_minutes"`
	Summary       string   `json:"summary"        db:"summary"`      // from the README table
	Timelines     []string `json:"timelines"      db:"timelines"`
	Tags          []string `json:"tags"           db:"tags"`
	Prerequisites []string `json:"prerequisites"`
	Related       []string `json:"related"`
	Sections      []string `json:"sections"       db:"sections"`
	HasSelfCheck  bool     `json:"has_self_check" db:"has_self_check"`
	WordCount     int      `json:"word_count"     db:"word_count"`
	Source        string   `json:"source"         db:"source"`
	Body          string   `json:"body,omitempty" db:"body_html"`
}
```

Notes:

- **`order` is a SQL reserved word** → column `sort_order`. Quoting `"order"` works in Postgres but
  you'll forget the quotes once.
- **Arrays:** `[]string` doesn't implement `sql.Scanner`, so a bare `rows.Scan(&t.Tags)` fails at
  *runtime* (`unsupported Scan, storing driver.Value type []uint8 into type *[]string`) — Postgres
  sends `text[]` as one literal, `{a,b,"c d"}`. Use `pq.Array(&t.Tags)` at the scan site rather than
  `pq.StringArray` in the struct: these structs are also the JSON response types, and `pgx` v5 (the
  modern default; `lib/pq` is in maintenance mode) handles `[]string` natively, so the wrapper is the
  cheaper thing to remove later.
- **`prerequisites` / `related` are not `text[]`** — use join tables
  `topic_prerequisites(topic_id, prereq_id)` and `topic_related(...)`, both FKs to `topics(id)`, so
  the DAG is traversable in SQL. That leaves `tags`, `timelines`, `sections` as the only real arrays.
- **`Body` / `body_html`:** store rendered HTML with links already rewritten. Postgres TOASTs values
  over ~2 KB out of line, so at a 3.1 KB median the list query pays nothing for the column existing —
  just don't name it in list queries. One struct with `omitempty` beats duplicating 14 fields into a
  `TopicSummary`; revisit if read paths multiply.
- Module-level prerequisites are a **derived query**, not a column:
  ```sql
  SELECT p.prereq_id FROM topic_prerequisites p
  JOIN topics t ON t.id = p.topic_id
  WHERE t.module_id = $1 AND t.sort_order = 0;
  ```
- `module_order INT NOT NULL` (a `UNIQUE` makes gaps explicitly intentional, at the cost of making
  insertions multi-row migrations).

### Ingestion pipeline

**Walk `topics/**/*.md`** — don't read `manifest.json` as the source of truth. Everything in the
manifest is derivable from the files (`path` from the walk, `word_count`/`sections`/`has_self_check`
from the body), the manifest is hand-maintained so it *can* drift, and the body and `summary` exist
only in the files. Use the manifest as a cross-check: fail the build if the walk and the manifest
disagree.

Four transforms the parser owns:

1. **Frontmatter → columns**, plus the four derived fields.
2. **Contents-table extraction** — 7 module READMEs embed a table listing their own topics. Match
   rows to topics by link target (`cache-locations.md` → `path`), lift the last column into
   `Topic.summary`, then **strip the table** from what you store as `Module.Overview`. That last
   column ("Why cache at all, and what does it cost?") exists nowhere else and is the intended
   subtitle for a topic row. 4 modules have no table → `summary` is `""`.
3. **Link rewriting** — bodies contain `caching-overview.md`, `../networking/cdn.md`, and images as
   *raw HTML* (`<p align="center"><img src="../assets/images/Q6z24La.png">`), not markdown `![]()`. A
   markdown-link-only rewriter will miss the images. Map `*.md` → app route via `path` → `id`, and
   image `src` → the static asset path.
4. **`order: 0` handling** — that entry is the module's intro prose, not a lesson. Ingest it into
   `Module.Overview` and list only `sort_order > 0` as topics. Otherwise the user opens **Caching**
   and the first item inside is a topic also called **Caching**.

---

## Part 3 — Frontend (delegated session)

### What already works — use it as the reference

`frontend/src/App.jsx` has a working Google OAuth flow:

- `GET {backend}/api/login` opened in a **popup**; the app polls `/auth/status` every 1.5s until 200.
- Session is a Redis-backed cookie (`gin-contrib/sessions`): 6h `MaxAge`, `HttpOnly`, `Path=/`,
  `SameSite=Lax`, `Secure: false` (local HTTP only).
- `POST /api/logout` clears it.
- `frontend/vite.config.js` proxies `/api` → `localhost:9000` and rewrites `/auth/status` → `GET /`
  on the backend, so the cookie stays same-origin and no CORS is needed.
- **Every new content call goes through the proxy** — `/api/...` with `credentials: 'include'`.

Stack: React 19 + Vite 8, `oxlint`. Backend on `localhost:9000`. **No router, no state library, no CSS
framework installed.**

### The three screens

1. **Module grid** (post-login) — the 9 cards from Part 1, in `module_order`. Title, topic count,
   maybe summed `est_minutes`. No topic data needed.
2. **Module page** — `Module.overview` (the rendered README prose) at the top, then the topic list on
   scroll.
3. **Topic page** — the rendered body.

### API contract (proposed — not built)

```
GET /api/modules          → [Module]              (no topics, no bodies)
GET /api/modules/:id      → Module + topics[]      (topic metadata, no bodies)
GET /api/topics/:id       → Topic (with body)
```

```jsonc
// Module
{
  "id": "caching",
  "title": "Caching",
  "module_order": 60,
  "topic_count": 4,          // excludes the order:0 index entry
  "overview": "<rendered README prose, contents table stripped>",
  "topics": [ /* Topic — omitted on the list endpoint */ ]
}

// Topic
{
  "id": "cache-locations",
  "title": "Where to Cache",
  "module_id": "caching",
  "order": 20,
  "est_minutes": 5,
  "summary": "Client, CDN, web server, database, or application?",
  "timelines": ["short", "medium", "long"],
  "tags": ["caching", "architecture"],
  "prerequisites": ["caching-overview"],
  "related": ["cdn", "reverse-proxy"],
  "sections": ["Client caching", "CDN caching", "..."],
  "has_self_check": true,
  "word_count": 358,
  "source": "https://github.com/donnemartin/...",
  "body": "<rendered HTML — only on GET /api/topics/:id>"
}
```

**Build against fixtures derived from `topics/manifest.json`** — it has every field except `body`,
`summary` and `overview`, needs no markdown parsing, and is only 69 KB. Stub those three by hand for
now.

### Frontend gotchas

1. **`overview` arrives table-free.** If a contents table shows up in it, that's a backend bug to
   report — don't hide it with CSS.
2. **`summary` is `""` for 14 of 58 topic rows** — every module's order:0 index entry (which the
   module-detail query filters out anyway), plus all of `root` and 5 of 6 `study-guide` topics,
   since those READMEs have no contents table to lift a summary from. The topic row must look right
   without a subtitle.
3. **`body` and `overview` are server-rendered HTML with links already rewritten.** Confirm before
   adding a markdown renderer — if they arrive as raw markdown instead, that changes your deps.
4. **Bodies contain raw inline HTML** — `<p align="center">`, `<img>`, deprecated `align` attributes,
   tables. Style defensively; don't assume clean semantic markup.
5. **Images** come from wherever the backend maps `topics/assets/` (9.6 MB, path TBD — ask).
6. **12 of 63 topics have no prerequisites** — design the "no prerequisites" state.
7. **`timelines` is a track filter.** Every topic is tagged `short`/`medium`/`long`, and
   `manifest.json` has a `timelines` block with a label, goal, practice note and entry point per
   track. A track selector is expected UI.
8. **`prerequisites` / `related` are a real cross-module DAG** — good material for an "up next" rail
   or a dependency map.
9. **Many topics end in `## Self-check`** (`has_self_check`), and `topics/study-guide/flashcards.md`
   references Anki decks. Spaced repetition is a likely later feature — don't design the topic page as
   if the body is the only thing on it.
10. **Say "topic", not "module", in UI copy.** The primer's own prose calls each *lesson* a "module"
    (hence the `| Module |` table column and three `## Modules` headings), while our schema uses
    "module" for the group. Expect stray "module" wording inside rendered content.
11. **CC BY 4.0 attribution must be visible** somewhere in the UI — see `topics/ATTRIBUTION.md`.

---

## Open decisions — ask, don't assume

**Frontend:** router, state management, styling approach (nothing installed yet).

**Shared:**
- ~~Are the `exercises` sub-modules a special-cased screen, or do modules become a 2-level tree with
  `Module.ParentId`?~~ **Decided: neither.** Theory-only app; exercises are excluded from the module
  tree entirely and `modules` stays flat. Any future exercises feature is unattached, with its own
  tables and its own approach — not the primer's `exercises/*` tree.
- Is `body` HTML or markdown over the wire?
- Asset URL scheme for `topics/assets/images/`.

**Backend:** progress tracking (per-topic completion) is anticipated but has no schema or endpoints;
whether `manifest.json` becomes a generated artifact instead of hand-maintained.

## Verification snapshot (2026-09-01)

Facts checked against the tree, not assumed: 12 modules / 75 entries · all 26 distinct prerequisite
values resolve to valid topic ids · prerequisite graph is a DAG · 24 cross-module prerequisite edges ·
12 leaf topics with no prerequisites · 7 of 11 module READMEs carry a contents table · index topic id
== module id for 10 of 12 modules · 347 KB markdown (median 3.1 KB, max 19.5 KB) · 9.6 MB assets ·
`module_order` diff was 44 insertions / 0 deletions.
