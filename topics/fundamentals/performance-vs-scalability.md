---
id: performance-vs-scalability
title: Performance vs Scalability
module: fundamentals
order: 20
est_minutes: 3
timelines: [short, medium, long]
tags: [tradeoffs, scalability, performance]
prerequisites: [start-here]
related: [latency-vs-throughput, cap-theorem, load-balancer]
source: https://github.com/donnemartin/system-design-primer#performance-vs-scalability
---

# Performance vs Scalability

A service is **scalable** if it results in increased **performance** in a manner proportional to
resources added. Generally, increasing performance means serving more units of work, but it can
also be to handle larger units of work, such as when datasets grow.<sup>[1](http://www.allthingsdistributed.com/2006/03/a_word_on_scalability.html)</sup>

Another way to look at performance vs scalability:

* If you have a **performance** problem, your system is **slow for a single user**.
* If you have a **scalability** problem, your system is **fast for a single user but slow under heavy load**.

## Why the distinction matters

The two problems have different fixes, and applying the wrong one wastes effort:

* A performance problem is usually addressed in the code path itself — a bad query, a missing
  [index](../data/sql-tuning.md#use-good-indices), an N+1 round trip, a synchronous call that should
  be [asynchronous](../application/asynchronism.md).
* A scalability problem is addressed structurally — [horizontal scaling](../networking/load-balancer.md#horizontal-scaling),
  [caching](../caching/README.md), [sharding](../data/sharding.md), [federation](../data/federation.md).

Adding servers to a system that is slow for one user makes it slow for many users on more hardware.

## Source(s) and further reading

* [A word on scalability](http://www.allthingsdistributed.com/2006/03/a_word_on_scalability.html)
* [Scalability, availability, stability, patterns](http://www.slideshare.net/jboner/scalability-availability-stability-patterns/)

## Self-check

1. Define "scalable" in terms of resources added and performance gained.
2. Your API responds in 3 seconds with one user and 3 seconds with a thousand. Performance problem or scalability problem?
3. Your API responds in 50 ms with one user and 8 seconds at peak. Which is it, and what class of fix applies?
