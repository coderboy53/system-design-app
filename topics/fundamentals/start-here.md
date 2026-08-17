---
id: start-here
title: "System Design Topics: Start Here"
module: fundamentals
order: 10
est_minutes: 5
timelines: [short, medium, long]
tags: [foundations, scalability, orientation]
prerequisites: []
related: [performance-vs-scalability, latency-vs-throughput, cap-theorem]
source: https://github.com/donnemartin/system-design-primer#system-design-topics-start-here
---

# System Design Topics: Start Here

New to system design?

First, you'll need a basic understanding of common principles: what they are, how they are used,
and their pros and cons.

## Step 1: Review the scalability video lecture

[Scalability Lecture at Harvard](https://www.youtube.com/watch?v=-W9F__D3oY4)

Topics covered:

* [Vertical scaling](../networking/load-balancer.md#horizontal-scaling) — scaling *up* a single server on more expensive hardware
* [Horizontal scaling](../networking/load-balancer.md#horizontal-scaling) — scaling *out* using commodity machines
* [Caching](../caching/README.md)
* [Load balancing](../networking/load-balancer.md)
* [Database replication](../data/replication.md)
* Database partitioning — see [sharding](../data/sharding.md) and [federation](../data/federation.md)

## Step 2: Review the scalability article

[Scalability](https://web.archive.org/web/20221030091841/http://www.lecloud.net/tagged/scalability/chrono)

Topics covered:

* [Clones](https://web.archive.org/web/20220530193911/https://www.lecloud.net/post/7295452622/scalability-for-dummies-part-1-clones) — see [horizontal scaling](../networking/load-balancer.md#horizontal-scaling)
* [Databases](https://web.archive.org/web/20220602114024/https://www.lecloud.net/post/7994751381/scalability-for-dummies-part-2-database) — see [data](../data/README.md)
* [Caches](https://web.archive.org/web/20230126233752/https://www.lecloud.net/post/9246290032/scalability-for-dummies-part-3-cache) — see [caching](../caching/README.md)
* [Asynchronism](https://web.archive.org/web/20220926171507/https://www.lecloud.net/post/9699762917/scalability-for-dummies-part-4-asynchronism) — see [asynchronism](../application/asynchronism.md)

## Next steps

Next, we'll look at high-level trade-offs:

* **Performance** vs **scalability** → [module](performance-vs-scalability.md)
* **Latency** vs **throughput** → [module](latency-vs-throughput.md)
* **Availability** vs **consistency** → [module](cap-theorem.md)

Keep in mind that **everything is a trade-off**.

Then we'll dive into more specific topics such as [DNS](../networking/dns.md),
[CDNs](../networking/cdn.md), and [load balancers](../networking/load-balancer.md).

## Self-check

1. What is the difference between vertical and horizontal scaling?
2. Name the four "scalability for dummies" pillars covered by the article.
3. What are the three high-level trade-offs you look at before diving into specific components?
