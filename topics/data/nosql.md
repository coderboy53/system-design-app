---
id: nosql
title: NoSQL
module: data
order: 70
est_minutes: 4
timelines: [short, medium, long]
tags: [database, nosql, base, consistency]
prerequisites: [rdbms, cap-theorem]
related: [key-value-store, document-store, wide-column-store, graph-database, sql-or-nosql]
source: https://github.com/donnemartin/system-design-primer#nosql
---

# NoSQL

NoSQL is a collection of data items represented in a **key-value store**, **document store**,
**wide column store**, or a **graph database**. Data is [**denormalized**](denormalization.md), and **joins are generally
done in the application code**. Most NoSQL stores lack true [ACID](rdbms.md) transactions and favor
[eventual consistency](../fundamentals/consistency-patterns.md#eventual-consistency).

## BASE

**BASE** is often used to describe the properties of NoSQL databases. In comparison with the
[CAP theorem](../fundamentals/cap-theorem.md), **BASE chooses availability over consistency**.

* **Basically available** — the system guarantees availability.
* **Soft state** — the state of the system may change over time, even without input.
* **Eventual consistency** — the system will become consistent over a period of time, given that the system doesn't receive input during that period.

## The four store types

In addition to choosing between [SQL or NoSQL](sql-or-nosql.md), it is helpful to understand which
type of NoSQL database best fits your use case(s):

| Store | Abstraction | Module |
|---|---|---|
| Key-value store | hash table | [key-value-store.md](key-value-store.md) |
| Document store | key-value store with documents stored as values | [document-store.md](document-store.md) |
| Wide column store | nested map `ColumnFamily<RowKey, Columns<ColKey, Value, Timestamp>>` | [wide-column-store.md](wide-column-store.md) |
| Graph database | graph | [graph-database.md](graph-database.md) |

Note the layering: a key-value store is the basis for more complex systems such as a document store,
and in some cases, a graph database.

## Source(s) and further reading: NoSQL

* [Explanation of base terminology](http://stackoverflow.com/questions/3342497/explanation-of-base-terminology)
* [NoSQL databases: a survey and decision guidance](https://medium.com/baqend-blog/nosql-databases-a-survey-and-decision-guidance-ea7823a822d#.wskogqenq)
* [Scalability](https://web.archive.org/web/20220602114024/https://www.lecloud.net/post/7994751381/scalability-for-dummies-part-2-database)
* [Introduction to NoSQL](https://www.youtube.com/watch?v=qI_g07C_Q5I)
* [NoSQL patterns](http://horicky.blogspot.com/2009/11/nosql-patterns.html)

## Self-check

1. Expand BASE and define each term.
2. Which side of the CAP trade-off does BASE take, and which does ACID take?
3. Where do joins happen in a NoSQL system?
4. Name the four store types and the abstraction each one presents.
