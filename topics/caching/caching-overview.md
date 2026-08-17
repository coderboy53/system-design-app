---
id: caching-overview
title: Caching Overview
module: caching
order: 10
est_minutes: 4
timelines: [short, medium, long]
tags: [caching, performance, invalidation]
prerequisites: [data]
related: [cache-locations, what-to-cache, cache-update-strategies, cdn]
source: https://github.com/donnemartin/system-design-primer#cache
---

# Caching Overview

<p align="center">
  <img src="../assets/images/Q6z24La.png" alt="Cache">
  <br/>
  <i><a href="http://horicky.blogspot.com/2010/10/scalable-system-design-patterns.html">Source: Scalable system design patterns</a></i>
</p>

Caching **improves page load times** and can **reduce the load on your servers and databases**. In
this model, the dispatcher will first look up if the request has been made before and try to find
the previous result to return, in order to save the actual execution.

Databases often benefit from a **uniform distribution of reads and writes** across their partitions.
**Popular items can skew the distribution, causing bottlenecks.** Putting a cache in front of a
database can help **absorb uneven loads and spikes in traffic**.

That second paragraph is the argument interviewers want to hear: a cache is not only a latency
optimization, it is load-shape insurance for the [database](../data/README.md) behind it — and it is
what makes [sharding](../data/sharding.md) hot spots survivable.

## The three decisions

* [**Where** to cache](cache-locations.md) — client, [CDN](../networking/cdn.md), web server, database, application
* [**What** to cache](what-to-cache.md) — database queries or objects
* [**When** to update](cache-update-strategies.md) — cache-aside, write-through, write-behind, refresh-ahead

## Disadvantage(s): cache

* Need to **maintain consistency** between caches and the source of truth such as the database,
  through [cache invalidation](https://en.wikipedia.org/wiki/Cache_algorithms).
* **Cache invalidation is a difficult problem**; there is additional complexity associated with when to update the cache.
* Need to make **application changes** such as adding Redis or Memcached.

## Source(s) and further reading

* [From cache to in-memory data grid](http://www.slideshare.net/tmatyashovsky/from-cache-to-in-memory-data-grid-introduction-to-hazelcast)
* [Scalable system design patterns](http://horicky.blogspot.com/2010/10/scalable-system-design-patterns.html)
* [Introduction to architecting systems for scale](http://lethain.com/introduction-to-architecting-systems-for-scale/)
* [Scalability, availability, stability, patterns](http://www.slideshare.net/jboner/scalability-availability-stability-patterns/)
* [Scalability](https://web.archive.org/web/20230126233752/https://www.lecloud.net/post/9246290032/scalability-for-dummies-part-3-cache)
* [AWS ElastiCache strategies](http://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/Strategies.html)
* [Wikipedia](https://en.wikipedia.org/wiki/Cache_(computing))

## Self-check

1. Name the two benefits of caching stated up front.
2. Why do popular items cause a problem for a partitioned database, and how does a cache help?
3. What are the three disadvantages of introducing a cache?
4. What are the three decisions every caching design has to make?
