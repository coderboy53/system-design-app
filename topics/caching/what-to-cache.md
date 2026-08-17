---
id: what-to-cache
title: What to Cache
module: caching
order: 30
est_minutes: 4
timelines: [short, medium, long]
tags: [caching, invalidation, granularity]
prerequisites: [cache-locations]
related: [cache-update-strategies, denormalization, caching-overview]
source: https://github.com/donnemartin/system-design-primer#caching-at-the-database-query-level
---

# What to Cache

Two general granularities: **database queries** and **objects**.

## Caching at the database query level

Whenever you query the database, **hash the query as a key** and store the result to the cache.

This approach suffers from **expiration issues**:

* Hard to **delete a cached result with complex queries**
* If one piece of data changes such as a table cell, you need to **delete all cached queries that might include the changed cell**

That second point is the killer. Query-level caching looks simple and becomes unmaintainable exactly
when the schema gets interesting.

## Caching at the object level

See your data as an **object**, similar to what you do with your application code. Have your
application **assemble the dataset** from the database into a class instance or data structure(s):

* **Remove the object from cache if its underlying data has changed**
* Allows for **asynchronous processing**: workers assemble objects by consuming the latest cached object

Object-level invalidation is tractable because the ownership is clear — one object, one owner, one
place to invalidate.

## Suggestions of what to cache

* User sessions
* Fully rendered web pages
* Activity streams
* User graph data

## Self-check

1. How is the cache key formed when caching at the query level?
2. Give the two expiration problems query-level caching runs into.
3. What makes object-level invalidation more tractable?
4. What does object-level caching enable for background workers?
5. Name four things worth caching.
