---
id: fundamentals
title: Fundamentals
module: fundamentals
order: 0
est_minutes: 2
timelines: [short, medium, long]
tags: [index, tradeoffs]
prerequisites: []
related: [start-here]
source: https://github.com/donnemartin/system-design-primer#index-of-system-design-topics
---

# Fundamentals

> **Everything is a trade-off.**

The high-level trade-offs that every later module refers back to. Read these in order — the
networking, application, data, and caching modules all assume this vocabulary.

## Modules

| # | Module | Time | What it answers |
|---|---|---|---|
| 1 | [System design topics: start here](start-here.md) | 5 min | Where do I begin if I know nothing? |
| 2 | [Performance vs scalability](performance-vs-scalability.md) | 3 min | Is my system slow, or slow *under load*? |
| 3 | [Latency vs throughput](latency-vs-throughput.md) | 2 min | Am I optimizing time-per-action or actions-per-time? |
| 4 | [Availability vs consistency (CAP theorem)](cap-theorem.md) | 6 min | Which two guarantees can I actually have? |
| 5 | [Consistency patterns](consistency-patterns.md) | 5 min | How synchronized do my copies need to be? |
| 6 | [Availability patterns](availability-patterns.md) | 9 min | How do I stay up when a machine dies? |

## Where this leads

* Consistency and availability choices drive the [database](../data/README.md) and [caching](../caching/README.md) modules.
* Fail-over and replication reappear as [master-slave](../data/replication.md#master-slave-replication) and [master-master](../data/replication.md#master-master-replication) replication.
* The latency/throughput distinction is what [asynchronism](../application/asynchronism.md) trades between.
