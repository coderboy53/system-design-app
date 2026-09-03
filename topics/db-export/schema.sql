-- Schema for the temporary bulk-load API that reads topics/db-export/*.json.
-- Current load is theory-only: the `exercises` module tree (3 modules, 17
-- topics) is excluded from db-export/ for now, but the schema still supports
-- it (parent_module_id) since interactive exercises are a later feature.
-- Matches HANDOFF.md Part 2, plus:
--   * `parent_module_id` on modules, to hold the exercises -> exercises/system-design
--     and exercises -> exercises/object-oriented-design tree (see modules/*.json).
--   * `is_index` on topics: the order:0 entry of every module is stored as a real
--     topics row (so topic_prerequisites/topic_related edges pointing at it, e.g.
--     `networking -> fundamentals`, resolve as a real FK) but the module-detail API
--     must filter it out with `WHERE sort_order > 0 AND NOT is_index`.
--
-- Load order (respects FKs): modules (by module_order, so a parent module id
-- exists before its children) -> topics -> topic_prerequisites / topic_related.

CREATE TABLE modules (
    id               TEXT PRIMARY KEY,
    title            TEXT NOT NULL,
    module_order     INT NOT NULL UNIQUE,
    parent_module_id TEXT REFERENCES modules(id),
    topic_count      INT NOT NULL DEFAULT 0,
    overview         TEXT NOT NULL DEFAULT ''
);

CREATE TABLE topics (
    id             TEXT PRIMARY KEY,
    title          TEXT NOT NULL,
    module_id      TEXT NOT NULL REFERENCES modules(id),
    sort_order      INT NOT NULL,
    is_index       BOOLEAN NOT NULL DEFAULT false,
    path           TEXT NOT NULL,
    est_minutes    INT NOT NULL,
    summary        TEXT NOT NULL DEFAULT '',
    timelines      TEXT[] NOT NULL DEFAULT '{}',
    tags           TEXT[] NOT NULL DEFAULT '{}',
    sections       TEXT[] NOT NULL DEFAULT '{}',
    has_self_check BOOLEAN NOT NULL DEFAULT false,
    word_count     INT NOT NULL DEFAULT 0,
    source         TEXT NOT NULL DEFAULT '',
    body           TEXT NOT NULL DEFAULT '',
    UNIQUE (module_id, sort_order)
);

CREATE TABLE topic_prerequisites (
    topic_id  TEXT NOT NULL REFERENCES topics(id) ON DELETE CASCADE,
    prereq_id TEXT NOT NULL REFERENCES topics(id) ON DELETE CASCADE,
    PRIMARY KEY (topic_id, prereq_id)
);

CREATE TABLE topic_related (
    topic_id   TEXT NOT NULL REFERENCES topics(id) ON DELETE CASCADE,
    related_id TEXT NOT NULL REFERENCES topics(id) ON DELETE CASCADE,
    PRIMARY KEY (topic_id, related_id)
);
