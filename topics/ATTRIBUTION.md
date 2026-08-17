---
id: attribution
title: Attribution and License
module: root
order: 900
est_minutes: 3
timelines: [short, medium, long]
tags: [meta, license, attribution]
prerequisites: []
related: [root]
source: https://github.com/donnemartin/system-design-primer#license
---

# Attribution

All content under `topics/` is adapted from **[The System Design Primer](https://github.com/donnemartin/system-design-primer)**
by [Donne Martin](https://github.com/donnemartin).

```
Copyright 2017 Donne Martin

Creative Commons Attribution 4.0 International License (CC BY 4.0)

http://creativecommons.org/licenses/by/4.0/
```

The full license text as shipped with the source repository is preserved at
[`assets/PRIMER-LICENSE.txt`](assets/PRIMER-LICENSE.txt).

## What changed in the adaptation

The primer is a single ~1,800 line `README.md` plus a `solutions/` folder. This path restructures
that material without altering its substance:

* **Split** the monolithic README into one file per concept, so the app can serve a module at a time.
* **Added** YAML frontmatter to every file (`id`, `title`, `module`, `order`, `est_minutes`,
  `timelines`, `tags`, `prerequisites`, `related`, `source`) for indexing and routing.
* **Rewrote** intra-document anchor links (`#sharding`) as relative file links (`../data/sharding.md`)
  so every cross-reference resolves inside this content tree.
* **Localized** all images. The primer's `images/` folder plus the diagrams the solution files
  hot-linked from `i.imgur.com` are vendored into [`assets/images/`](assets/images). No module
  fetches from a remote host.
* **Converted** the object-oriented design Jupyter notebooks to markdown with their Python solutions inline.
* **Added** a `Self-check` section to each concept module. These questions are derived from the
  primer's own content and exist to feed a spaced-repetition queue.
* **Added** [`manifest.json`](manifest.json), a machine-readable index of every module.

Every module's frontmatter carries a `source:` field pointing at the exact section of the upstream
primer it derives from.

## Onward attribution from the primer

The primer itself credits:

* [Hired in tech](http://www.hiredintech.com/system-design/the-system-design-process/)
* [Cracking the coding interview](https://www.amazon.com/dp/0984782850/)
* [High scalability](http://highscalability.com/)
* [checkcheckzz/system-design-interview](https://github.com/checkcheckzz/system-design-interview)
* [shashank88/system_design](https://github.com/shashank88/system_design)
* [mmcgrana/services-engineering](https://github.com/mmcgrana/services-engineering)
* [System design cheat sheet](https://gist.github.com/vasanthk/485d1c25737e8e72759f)
* [A distributed systems reading list](http://dancres.github.io/Pages/)
* [Cracking the system design interview](http://www.puncsky.com/blog/2016-02-13-crack-the-system-design-interview)

Individual diagrams retain their original source attribution inline in the module that uses them.
