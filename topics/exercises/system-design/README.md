---
id: system-design-exercises
title: System Design Exercises
module: exercises/system-design
module_order: 100
order: 0
est_minutes: 3
timelines: [short, medium, long]
tags: [index, exercise, system-design, practice]
prerequisites: [interview-approach]
related: [exercises, additional-questions, interview-approach]
source: https://github.com/donnemartin/system-design-primer#system-design-interview-questions-with-solutions
---

# System Design Exercises

> Common system design interview questions with sample discussions, code, and diagrams.

Every solution follows the same four steps from the
[interview framework](../../study-guide/interview-approach.md), and links back into the
[topic modules](../../README.md) instead of repeating their content.

| Exercise | Time | Track | Core themes |
|---|---|---|---|
| [Design Pastebin.com (or Bit.ly)](pastebin/README.md) | 25 min | short · medium · long | Hashing, [SQL vs NoSQL](../../data/sql-or-nosql.md), object store, MapReduce analytics |
| [Design the Twitter timeline and search](twitter/README.md) | 25 min | short · medium · long | Fan-out on write, [denormalization](../../data/denormalization.md), [memory cache](../../caching/cache-locations.md#application-caching) |
| [Design a web crawler](web-crawler/README.md) | 25 min | medium · long | [Queues](../../application/asynchronism.md#message-queues), dedup, reverse index |
| [Design Mint.com](mint/README.md) | 25 min | medium · long | ETL, [task queues](../../application/asynchronism.md#task-queues), categorization |
| [Design the data structures for a social network](social-graph/README.md) | 25 min | medium · long | [Graphs](../../data/graph-database.md), shortest path, [sharding](../../data/sharding.md) by person |
| [Design a key-value store for a search engine](query-cache/README.md) | 25 min | medium · long | [LRU cache](../object-oriented-design/lru-cache.md), [cache update strategies](../../caching/cache-update-strategies.md) |
| [Design Amazon's sales ranking by category](sales-rank/README.md) | 25 min | long | MapReduce, offline aggregation, [SQL tuning](../../data/sql-tuning.md) |
| [Design a system that scales to millions of users on AWS](scaling-aws/README.md) | 30 min | short · medium · long | **The whole scaling ladder, iteratively** |

## Suggested order

1. **[Pastebin](pastebin/README.md)** — smallest surface, exercises all four steps cleanly. Start here.
2. **[Scaling to millions of users on AWS](scaling-aws/README.md)** — walks the entire
   [scaling ladder](../../data/README.md#the-scaling-ladder) step by step, and is explicitly the
   reference for how to scale *iteratively* rather than jumping to the final design.
3. Everything else, in any order.

## The one instruction that matters

> **Do not simply jump right into the final design from the initial design.**
>
> State that you would do this iteratively: 1) benchmark / load test, 2) profile for bottlenecks,
> 3) address bottlenecks while evaluating alternatives and trade-offs, and 4) repeat.
