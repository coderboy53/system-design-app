---
id: key-value-store
title: Key-Value Store
module: data
order: 80
est_minutes: 3
timelines: [short, medium, long]
tags: [database, nosql, key-value, cache]
prerequisites: [nosql]
related: [document-store, caching-overview, cache-locations, service-discovery]
source: https://github.com/donnemartin/system-design-primer#key-value-store
---

# Key-Value Store

> **Abstraction: hash table**

A key-value store generally allows for **O(1) reads and writes** and is often backed by **memory or
SSD**. Data stores can maintain keys in
[lexicographic order](https://en.wikipedia.org/wiki/Lexicographical_order), allowing **efficient
retrieval of key ranges**. Key-value stores can allow for storing of **metadata with a value**.

Key-value stores provide **high performance** and are often used for **simple data models** or for
**rapidly-changing data**, such as an [in-memory cache layer](../caching/cache-locations.md#application-caching).
Since they offer only a limited set of operations, **complexity is shifted to the application
layer** if additional operations are needed.

A key-value store is the **basis for more complex systems** such as a
[document store](document-store.md), and in some cases, a [graph database](graph-database.md).

## Source(s) and further reading: key-value store

* [Key-value database](https://en.wikipedia.org/wiki/Key-value_database)
* [Disadvantages of key-value stores](http://stackoverflow.com/questions/4056093/what-are-the-disadvantages-of-using-a-key-value-table-over-nullable-columns-or)
* [Redis architecture](http://qnimate.com/overview-of-redis-architecture/)
* [Memcached architecture](https://adayinthelifeof.nl/2011/02/06/memcache-internals/)

## Self-check

1. What complexity do reads and writes have, and what backs the store?
2. What does maintaining keys in lexicographic order enable?
3. Where does complexity go when the store's operation set isn't rich enough?
4. Which other store types are built on top of a key-value store?
