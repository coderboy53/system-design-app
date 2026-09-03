#!/usr/bin/env python3
"""Generate per-module / per-topic JSON fixtures from topics/manifest.json + the
markdown files themselves, for bulk-loading into Postgres via a temporary API.

Theory-only phase: modules under EXCLUDED_MODULE_PREFIXES (currently `exercises`,
covering exercises/system-design and exercises/object-oriented-design) are skipped
entirely, and any prerequisite/related edge pointing at a skipped topic is dropped.
Interactive exercises come back once that feature is built.

Reads: topics/manifest.json, topics/**/*.md
Writes: topics/db-export/modules/*.json, topics/db-export/topics/*.json
"""
import json
import re
import os

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
TOPICS = os.path.dirname(SCRIPT_DIR)
OUT = os.path.join(TOPICS, "db-export")

# Theory-only phase: interactive exercises come later. Any module id equal to or
# nested under one of these is skipped entirely (modules + their topics).
EXCLUDED_MODULE_PREFIXES = ("exercises",)


def is_excluded_module(module_id):
    return any(
        module_id == prefix or module_id.startswith(prefix + "/")
        for prefix in EXCLUDED_MODULE_PREFIXES
    )

FRONTMATTER_RE = re.compile(r"^---\n(.*?)\n---\n(.*)$", re.DOTALL)
LINK_RE = re.compile(r"\[([^\]]+)\]\(([^)]+\.md)\)")


def strip_frontmatter(raw):
    m = FRONTMATTER_RE.match(raw)
    if not m:
        raise ValueError("no frontmatter")
    return m.group(2).lstrip("\n")


def parent_module_id(module_id):
    if "/" in module_id:
        return module_id.rsplit("/", 1)[0]
    return None


def resolve_link_target(current_dir, target):
    """Resolve a relative .md link target (from a README) to a topics/-relative path."""
    joined = os.path.normpath(os.path.join(current_dir, target))
    return joined.replace(os.sep, "/")


def strip_contents_table_and_extract_summaries(body, current_dir, sibling_paths_by_path):
    """Find a markdown table whose row links resolve to sibling topic paths in this
    module. Lift the last cell of each matching row as that topic's summary, then
    remove the whole table block from the returned body.
    """
    lines = body.split("\n")
    summaries = {}
    out_lines = []
    i = 0
    removed_any = False
    while i < len(lines):
        line = lines[i]
        if line.strip().startswith("|"):
            # collect the whole table block
            start = i
            block = []
            while i < len(lines) and lines[i].strip().startswith("|"):
                block.append(lines[i])
                i += 1
            # does this table reference >=1 sibling topic path?
            matched = {}
            for row in block[2:]:  # skip header + separator row
                link_m = LINK_RE.search(row)
                if not link_m:
                    continue
                target = resolve_link_target(current_dir, link_m.group(2))
                if target in sibling_paths_by_path:
                    cells = [c.strip() for c in row.strip().strip("|").split("|")]
                    last_cell = cells[-1] if cells else ""
                    last_cell = LINK_RE.sub(r"\1", last_cell).strip("`* ")
                    matched[target] = last_cell
            if matched:
                summaries.update(matched)
                removed_any = True
                # drop a single blank line before/after too, if present
                continue
            else:
                out_lines.extend(block)
                continue
        else:
            out_lines.append(line)
            i += 1
    if removed_any:
        # drop a heading that now has nothing under it before the next heading/EOF
        cleaned = []
        n = len(out_lines)
        i = 0
        while i < n:
            line = out_lines[i]
            if line.strip().startswith("#"):
                j = i + 1
                while j < n and out_lines[j].strip() == "":
                    j += 1
                if j >= n or out_lines[j].strip().startswith("#"):
                    i += 1
                    continue
            cleaned.append(line)
            i += 1
        out_lines = cleaned

    new_body = "\n".join(out_lines)
    # collapse 3+ blank lines left behind by table removal
    new_body = re.sub(r"\n{3,}", "\n\n", new_body).strip("\n") + "\n"
    return new_body, summaries, removed_any


def main():
    with open(os.path.join(TOPICS, "manifest.json")) as f:
        manifest = json.load(f)

    entries = [e for e in manifest["entries"] if not is_excluded_module(e["module"])]
    modules_meta = {
        m["id"]: m for m in manifest["modules"] if not is_excluded_module(m["id"])
    }
    kept_topic_ids = {e["id"] for e in entries}

    entries_by_path = {}
    for e in entries:
        current_dir = os.path.dirname(e["path"])
        entries_by_path[e["path"]] = e

    # read raw bodies (frontmatter stripped) for every entry, keyed by id
    raw_body_by_id = {}
    for e in entries:
        full_path = os.path.join(TOPICS, e["path"])
        with open(full_path, encoding="utf-8") as f:
            raw = f.read()
        raw_body_by_id[e["id"]] = strip_frontmatter(raw)

    # group entries by module id
    entries_by_module = {}
    for e in entries:
        entries_by_module.setdefault(e["module"], []).append(e)

    os.makedirs(os.path.join(OUT, "modules"), exist_ok=True)
    os.makedirs(os.path.join(OUT, "topics"), exist_ok=True)

    topic_summary_overrides = {}
    index_overview_by_id = {}

    modules_out = []
    topics_out = []

    for module_id, mod_entries in entries_by_module.items():
        mod_entries_sorted = sorted(mod_entries, key=lambda e: e.get("order", 0))
        index_entry = next((e for e in mod_entries_sorted if e.get("order", 0) == 0), None)
        meta = modules_meta[module_id]

        overview = ""
        summaries_for_module = {}
        if index_entry is not None:
            current_dir = os.path.dirname(index_entry["path"])
            sibling_paths = {e["path"] for e in mod_entries_sorted if e is not index_entry}
            body = raw_body_by_id[index_entry["id"]]
            overview, summaries_for_module, _ = strip_contents_table_and_extract_summaries(
                body, current_dir, sibling_paths
            )
            # map resolved paths back to topic ids
            path_to_id = {e["path"]: e["id"] for e in mod_entries_sorted}
            for path, summ in summaries_for_module.items():
                tid = path_to_id.get(path)
                if tid:
                    topic_summary_overrides[tid] = summ
            index_overview_by_id[index_entry["id"]] = overview

        module_record = {
            "id": module_id,
            "title": meta["title"],
            "module_order": meta["module_order"],
            "parent_module_id": parent_module_id(module_id),
            "topic_count": meta["count"] - 1 if index_entry is not None else meta["count"],
            "overview": overview,
        }
        modules_out.append(module_record)

    for module_record in modules_out:
        fname = module_record["id"].replace("/", "--") + ".json"
        with open(os.path.join(OUT, "modules", fname), "w", encoding="utf-8") as f:
            json.dump(module_record, f, indent=2, ensure_ascii=False)
            f.write("\n")

    for e in entries:
        is_index = e.get("order", 0) == 0
        # Index (order:0) entries still get a topics row — the API's module-detail
        # endpoint filters them out with `WHERE sort_order > 0`, but they stay in
        # storage so prerequisite/related edges pointing at a module's index page
        # (e.g. `networking -> fundamentals`) resolve as a real topics.id FK, per
        # HANDOFF.md's derived module-prerequisites query.
        body = index_overview_by_id[e["id"]] if is_index else raw_body_by_id[e["id"]]
        topic_record = {
            "id": e["id"],
            "title": e["title"],
            "module_id": e["module"],
            "order": e.get("order", 0),
            "is_index": is_index,
            "path": e["path"],
            "est_minutes": e["est_minutes"],
            "summary": topic_summary_overrides.get(e["id"], ""),
            "timelines": e["timelines"],
            "tags": e["tags"],
            "prerequisites": [p for p in e["prerequisites"] if p in kept_topic_ids],
            "related": [r for r in e["related"] if r in kept_topic_ids],
            "sections": e["sections"],
            "has_self_check": e["has_self_check"],
            "word_count": e["word_count"],
            "source": e["source"],
            "body": body,
        }
        topics_out.append(topic_record)
        with open(os.path.join(OUT, "topics", e["id"] + ".json"), "w", encoding="utf-8") as f:
            json.dump(topic_record, f, indent=2, ensure_ascii=False)
            f.write("\n")

    print(f"wrote {len(modules_out)} module files, {len(topics_out)} topic files")
    with_summary = sum(1 for t in topics_out if t["summary"])
    print(f"{with_summary} topics got a summary lifted from a contents table")


if __name__ == "__main__":
    main()
