---
id: data
title: Data
module: data
module_order: 50
order: 0
est_minutes: 3
timelines: [short, medium, long]
tags: [index, database, storage]
prerequisites: [fundamentals]
related: [rdbms, nosql, sql-or-nosql]
source: https://github.com/donnemartin/system-design-primer#database
---

# Data

<p align="center">
  <img src="../assets/images/Xkm5CXz.png" alt="Database">
  <br/>
  <i><a href="https://www.youtube.com/watch?v=kKjm4ehYiMs">Source: Scaling up to your first 10 million users</a></i>
</p>

Where the state lives. Most system design interviews spend most of their time here — this is the
component that is hardest to scale and hardest to change later.

## Relational

| # | Module | Time | What it answers |
|---|---|---|---|
| 1 | [Relational database management system (RDBMS)](rdbms.md) | 4 min | What does ACID guarantee, and what are the ways to scale SQL? |
| 2 | [Replication](replication.md) | 7 min | Master-slave vs master-master, and the costs both share |
| 3 | [Federation](federation.md) | 4 min | Splitting databases **by function** |
| 4 | [Sharding](sharding.md) | 5 min | Splitting a dataset **by key** across databases |
| 5 | [Denormalization](denormalization.md) | 4 min | Trading write performance for read performance |
| 6 | [SQL tuning](sql-tuning.md) | 7 min | Schema, indices, joins, partitioning, query cache |

## Non-relational

| # | Module | Time | What it answers |
|---|---|---|---|
| 7 | [NoSQL](nosql.md) | 4 min | What is BASE, and how does it differ from ACID? |
| 8 | [Key-value store](key-value-store.md) | 3 min | Abstraction: hash table |
| 9 | [Document store](document-store.md) | 3 min | Abstraction: key-value store with documents as values |
| 10 | [Wide column store](wide-column-store.md) | 4 min | Abstraction: nested map of column families |
| 11 | [Graph database](graph-database.md) | 3 min | Abstraction: graph |

## Choosing

| # | Module | Time | What it answers |
|---|---|---|---|
| 12 | [SQL or NoSQL](sql-or-nosql.md) | 4 min | Which one, and how do I defend the choice? |

## The scaling ladder

The relational modules are roughly in the order you would apply them:

1. **Tune what you have** — [SQL tuning](sql-tuning.md): indices, schema, query cache.
2. **Add read capacity** — [master-slave replication](replication.md#master-slave-replication) plus a [cache](../caching/README.md).
3. **Add write capacity / availability** — [master-master replication](replication.md#master-master-replication).
4. **Split by function** — [federation](federation.md).
5. **Split by key** — [sharding](sharding.md).
6. **Stop joining** — [denormalization](denormalization.md).
7. **Move the workloads that don't fit** — [NoSQL](nosql.md).

Every step adds hardware and complexity. Do not skip to step 5 in an interview.
