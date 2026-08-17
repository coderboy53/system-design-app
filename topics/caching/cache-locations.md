---
id: cache-locations
title: Where to Cache
module: caching
order: 20
est_minutes: 5
timelines: [short, medium, long]
tags: [caching, architecture, redis, memcached]
prerequisites: [caching-overview]
related: [cdn, reverse-proxy, key-value-store, what-to-cache]
source: https://github.com/donnemartin/system-design-primer#client-caching
---

# Where to Cache

Caches can be located on the **client side** (OS or browser), **server side**, or in a **distinct
cache layer**.

## Client caching

Caches on the client — the OS or the browser. Free to you, but you control invalidation only
indirectly, through headers and TTLs.

## CDN caching

[CDNs](../networking/cdn.md) are considered a type of cache. Push and pull CDNs are two different
cache-population strategies for the same layer.

## Web server caching

[Reverse proxies](../networking/reverse-proxy.md) and caches such as
[Varnish](https://www.varnish-cache.org/) can serve **static and dynamic content directly**. Web
servers can also cache requests, returning responses **without having to contact application
servers**.

## Database caching

Your database usually includes some level of caching in a default configuration, optimized for a
generic use case. **Tweaking these settings for specific usage patterns** can further boost
performance. See also [tuning the query cache](../data/sql-tuning.md#tune-the-query-cache).

## Application caching

In-memory caches such as **Memcached** and **Redis** are [key-value stores](../data/key-value-store.md)
between your application and your data storage. Since the data is held in RAM, it is **much faster
than typical databases** where data is stored on disk.

RAM is more limited than disk, so [cache invalidation](https://en.wikipedia.org/wiki/Cache_algorithms)
algorithms such as [least recently used (LRU)](https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU))
can help invalidate 'cold' entries and keep 'hot' data in RAM. (Implement one yourself in the
[LRU cache exercise](../exercises/object-oriented-design/lru-cache.md).)

**Redis** has the following additional features over Memcached:

* Persistence option
* Built-in data structures such as sorted sets and lists

## Levels of caching

There are multiple levels you can cache, falling into two general categories —
**[database queries](what-to-cache.md#caching-at-the-database-query-level)** and
**[objects](what-to-cache.md#caching-at-the-object-level)**:

* Row level
* Query-level
* Fully-formed serializable objects
* Fully-rendered HTML

Generally, you should **try to avoid file-based caching**, as it makes cloning and auto-scaling more
difficult — the same reason [horizontally scaled servers must be stateless](../networking/load-balancer.md#horizontal-scaling).

## Self-check

1. List the five places a cache can live, from closest to the user to closest to the data.
2. What does a web server cache save you that an application cache does not?
3. What two features does Redis add over Memcached?
4. Why is file-based caching discouraged?
5. What role does LRU play in an in-memory cache?
