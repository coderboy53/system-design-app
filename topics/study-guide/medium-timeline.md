---
id: medium-timeline
title: Medium Timeline Track
module: study-guide
order: 20
est_minutes: 4
timelines: [medium]
tags: [meta, track, planning]
prerequisites: [study-guide]
related: [short-timeline, long-timeline, interview-approach]
source: https://github.com/donnemartin/system-design-primer#study-guide
---

# Medium Timeline Track

> **Goal: breadth and some depth.** Practice by solving **many** interview questions.

Everything in the [short timeline](short-timeline.md), plus depth in the areas that come up most:
data storage, caching, and asynchronism.

## 1. Complete the short track

Do the full breadth pass and the two reference tables first. Depth without breadth leaves holes an
interviewer will find.

* [Short timeline track](short-timeline.md)

## 2. Go deep on data

This is where most design interviews spend their time. Read every sub-module, including the
disadvantages sections, and be able to justify picking one technique over another.

* [Relational database management system (RDBMS)](../data/rdbms.md) — know ACID cold
* [Replication](../data/replication.md) — master-slave vs master-master, and the disadvantages shared by both
* [Federation](../data/federation.md)
* [Sharding](../data/sharding.md) — including rebalancing and consistent hashing
* [Denormalization](../data/denormalization.md)
* [SQL tuning](../data/sql-tuning.md) — schema, indices, joins, partitioning, query cache
* [NoSQL](../data/nosql.md) — BASE, and all four store types: [key-value](../data/key-value-store.md), [document](../data/document-store.md), [wide column](../data/wide-column-store.md), [graph](../data/graph-database.md)
* [SQL or NoSQL](../data/sql-or-nosql.md) — be able to argue either side

## 3. Go deep on caching and asynchronism

* [Caching overview](../caching/caching-overview.md) and its disadvantages
* [Where to cache](../caching/cache-locations.md)
* [What to cache](../caching/what-to-cache.md) — query level vs object level
* [When to update the cache](../caching/cache-update-strategies.md) — all four strategies with their failure modes
* [Asynchronism](../application/asynchronism.md) — message queues, task queues, back pressure

## 4. Go deep on the trade-off vocabulary

You will be asked to defend a choice. These modules are the vocabulary for doing that.

* [Availability vs consistency (CAP theorem)](../fundamentals/cap-theorem.md)
* [Consistency patterns](../fundamentals/consistency-patterns.md)
* [Availability patterns](../fundamentals/availability-patterns.md) — including [availability in numbers](../fundamentals/availability-patterns.md#availability-in-numbers) and the parallel vs sequence math
* [Load balancer](../networking/load-balancer.md) — layer 4 vs layer 7, and [horizontal scaling](../networking/load-balancer.md#horizontal-scaling)
* [RPC](../networking/rpc.md) vs [REST](../networking/rest.md) — including the side-by-side call comparison

## 5. Work through **many** exercises

Aim for most of the system design set. Work each one end to end using the
[four-step framework](interview-approach.md) *before* reading the solution.

* [Design Pastebin.com (or Bit.ly)](../exercises/system-design/pastebin/README.md)
* [Design the Twitter timeline and search](../exercises/system-design/twitter/README.md)
* [Design a web crawler](../exercises/system-design/web-crawler/README.md)
* [Design Mint.com](../exercises/system-design/mint/README.md)
* [Design a key-value store for a search engine](../exercises/system-design/query-cache/README.md)
* [Design a system that scales to millions of users on AWS](../exercises/system-design/scaling-aws/README.md)

And **many** object-oriented design questions:

* [Hash map](../exercises/object-oriented-design/hash-map.md), [LRU cache](../exercises/object-oriented-design/lru-cache.md), [call center](../exercises/object-oriented-design/call-center.md), [deck of cards](../exercises/object-oriented-design/deck-of-cards.md)

## 6. Read the context material properly

* Several [real world architectures](../appendix/real-world-architectures.md) — identify shared principles, common technologies, and patterns; study what problems each component solves, where it works and where it doesn't; review the lessons learned
* [Company architectures](../appendix/company-architectures.md) for your target companies
* Articles from the [company engineering blogs](../appendix/engineering-blogs.md)
* Work through **many** of the [additional system design interview questions](../appendix/additional-questions.md)

## Self-check

1. Given a read-heavy workload with a 100:1 read-to-write ratio, which scaling techniques would you reach for and in what order?
2. When does write-through beat cache-aside, and when does the reverse hold?
3. Two components each at 99.9% availability: what is the total in sequence, and in parallel?
4. Name three reasons to pick NoSQL and three to pick SQL.
