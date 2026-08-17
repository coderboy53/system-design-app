---
id: denormalization
title: Denormalization
module: data
order: 50
est_minutes: 4
timelines: [short, medium, long]
tags: [database, sql, scaling, reads, joins]
prerequisites: [sharding]
related: [sql-tuning, federation, sharding, what-to-cache]
source: https://github.com/donnemartin/system-design-primer#denormalization
---

# Denormalization

Denormalization attempts to **improve read performance at the expense of some write performance**.
Redundant copies of the data are written in multiple tables to **avoid expensive joins**. Some RDBMS
such as [PostgreSQL](https://en.wikipedia.org/wiki/PostgreSQL) and Oracle support
[materialized views](https://en.wikipedia.org/wiki/Materialized_view) which handle the work of
storing redundant information and keeping redundant copies consistent.

Once data becomes distributed with techniques such as [federation](federation.md) and
[sharding](sharding.md), managing joins **across data centers** further increases complexity.
Denormalization might circumvent the need for such complex joins.

**In most systems, reads can heavily outnumber writes 100:1 or even 1000:1.** A read resulting in a
complex database join can be very expensive, spending a significant amount of time on disk
operations. That ratio is the entire argument for the trade.

## Disadvantage(s): denormalization

* **Data is duplicated.**
* **Constraints** can help redundant copies of information stay in sync, which increases complexity of the database design.
* A denormalized database under **heavy write load** might perform worse than its normalized counterpart.

## Source(s) and further reading: denormalization

* [Denormalization](https://en.wikipedia.org/wiki/Denormalization)

## Self-check

1. What is traded for what in denormalization, and what read-to-write ratio makes the trade worthwhile?
2. Why does denormalization become more attractive after you federate or shard?
3. What do materialized views automate?
4. Under what workload does a denormalized schema perform worse than a normalized one?
