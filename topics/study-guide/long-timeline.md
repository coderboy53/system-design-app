---
id: long-timeline
title: Long Timeline Track
module: study-guide
order: 30
est_minutes: 5
timelines: [long]
tags: [meta, track, planning]
prerequisites: [study-guide]
related: [short-timeline, medium-timeline, interview-approach]
source: https://github.com/donnemartin/system-design-primer#study-guide
---

# Long Timeline Track

> **Goal: breadth and more depth.** Practice by solving **most** interview questions.

Everything in the [medium timeline](medium-timeline.md), plus primary sources, plus the full
exercise set, plus the topics the primer itself lists as unfinished.

## 1. Complete the medium track

* [Medium timeline track](medium-timeline.md)

## 2. Read the primary sources, not just the summaries

Every module ends with a **Source(s) and further reading** section. On a long timeline those links
are the point — the module is the index, the sources are the material. Prioritize:

* The [scalability lecture and article](../fundamentals/start-here.md) in full, not skimmed
* [Scalability, availability, stability, patterns](http://www.slideshare.net/jboner/scalability-availability-stability-patterns/) — referenced by half the modules here
* [Intro to architecting systems for scale](http://lethain.com/introduction-to-architecting-systems-for-scale)
* [Scaling up to your first 10 million users](https://www.youtube.com/watch?v=kKjm4ehYiMs)
* The original papers linked from [real world architectures](../appendix/real-world-architectures.md): MapReduce, Bigtable, Dynamo, GFS, Spanner, Chubby, Dapper

## 3. Cover every module, including the ones that are easy to skip

* [HTTP](../networking/http.md), [TCP and UDP](../networking/tcp-and-udp.md) — the protocol layer that layer 4 vs layer 7 load balancing actually rests on
* [Service discovery](../application/service-discovery.md)
* [Back pressure](../application/asynchronism.md#back-pressure) and exponential backoff
* [Security](../security/security-basics.md)
* [Graph database](../data/graph-database.md) and [wide column store](../data/wide-column-store.md) — the two store types candidates most often can't speak to

## 4. Study the topics the primer flags as under development

These are genuinely common in interviews despite being unfinished upstream. Work them from primary
sources.

* [Under development](../appendix/under-development.md) — distributed computing with MapReduce, consistent hashing, scatter gather

## 5. Work through **most** exercises

All eight system design exercises, all six object-oriented design exercises. Then re-do the first
two from scratch without notes — the second pass is where the framework becomes automatic.

* [Pastebin](../exercises/system-design/pastebin/README.md)
* [Twitter timeline and search](../exercises/system-design/twitter/README.md)
* [Web crawler](../exercises/system-design/web-crawler/README.md)
* [Mint.com](../exercises/system-design/mint/README.md)
* [Social network data structures](../exercises/system-design/social-graph/README.md)
* [Key-value store for a search engine](../exercises/system-design/query-cache/README.md)
* [Amazon's sales ranking by category](../exercises/system-design/sales-rank/README.md)
* [Scaling to millions of users on AWS](../exercises/system-design/scaling-aws/README.md)
* All of [object-oriented design](../exercises/object-oriented-design/README.md)

## 6. Work through **most** of the additional questions

The [additional system design interview questions](../appendix/additional-questions.md) list has no
written solutions — only reference links. On a long timeline, that is a feature: design each one
yourself first, then read the references and diff your design against reality.

## 7. Build the habit of reading architectures

* Work through [real world architectures](../appendix/real-world-architectures.md) systematically
* Read [company architectures](../appendix/company-architectures.md) for the companies you are interviewing with
* Subscribe to a few [company engineering blogs](../appendix/engineering-blogs.md) and keep reading after the interview

## Self-check

1. Pick any exercise and design it end to end on a blank page in 45 minutes, then diff against the solution.
2. Explain consistent hashing and why it reduces data movement when rebalancing shards.
3. Walk through what happens, layer by layer, from a user typing a URL to bytes rendering — naming every module in this path that participates.
4. For any component you would propose, state its disadvantage before the interviewer asks.
