---
id: short-timeline
title: Short Timeline Track
module: study-guide
order: 10
est_minutes: 4
timelines: [short]
tags: [meta, track, planning]
prerequisites: [study-guide]
related: [medium-timeline, long-timeline, interview-approach]
source: https://github.com/donnemartin/system-design-primer#study-guide
---

# Short Timeline Track

> **Goal: breadth.** Practice by solving **some** interview questions.

You have days, not weeks. Optimize for being able to say something sensible about any component
that comes up, and for having a repeatable process when you get a question cold.

## 1. Learn the process first

Do this before any topic reading — it is what you are actually graded on.

* [How to approach a system design interview question](interview-approach.md)

## 2. Breadth pass over the topics

Read every module once. Do not stop to go deep. The goal is that no term in the interview is new.

* [Fundamentals](../fundamentals/README.md) — [start here](../fundamentals/start-here.md), [performance vs scalability](../fundamentals/performance-vs-scalability.md), [latency vs throughput](../fundamentals/latency-vs-throughput.md), [CAP theorem](../fundamentals/cap-theorem.md), [consistency patterns](../fundamentals/consistency-patterns.md), [availability patterns](../fundamentals/availability-patterns.md)
* [Networking & edge](../networking/README.md) — [DNS](../networking/dns.md), [CDN](../networking/cdn.md), [load balancer](../networking/load-balancer.md), [reverse proxy](../networking/reverse-proxy.md), [HTTP](../networking/http.md), [TCP and UDP](../networking/tcp-and-udp.md), [RPC](../networking/rpc.md), [REST](../networking/rest.md)
* [Application layer](../application/README.md) — [application layer](../application/application-layer.md), [microservices](../application/microservices.md), [service discovery](../application/service-discovery.md), [asynchronism](../application/asynchronism.md)
* [Data](../data/README.md) — [RDBMS](../data/rdbms.md), [replication](../data/replication.md), [federation](../data/federation.md), [sharding](../data/sharding.md), [denormalization](../data/denormalization.md), [SQL tuning](../data/sql-tuning.md), [NoSQL](../data/nosql.md), [SQL or NoSQL](../data/sql-or-nosql.md)
* [Caching](../caching/README.md) — [overview](../caching/caching-overview.md), [where](../caching/cache-locations.md), [what](../caching/what-to-cache.md), [when to update](../caching/cache-update-strategies.md)
* [Security](../security/security-basics.md)

## 3. Memorize the two reference tables

You may be asked to estimate by hand. These are the only two things worth rote-memorizing.

* [Powers of two table](../appendix/powers-of-two.md)
* [Latency numbers every programmer should know](../appendix/latency-numbers.md)
* [Back-of-the-envelope calculations](../appendix/back-of-the-envelope.md)

## 4. Work through **some** exercises

Pick two or three. [Pastebin](../exercises/system-design/pastebin/README.md) and
[scaling to millions of users on AWS](../exercises/system-design/scaling-aws/README.md) give the
best coverage per hour — the first exercises the whole four-step process on a small surface, the
second walks the entire scaling ladder.

* [Design Pastebin.com (or Bit.ly)](../exercises/system-design/pastebin/README.md)
* [Design a system that scales to millions of users on AWS](../exercises/system-design/scaling-aws/README.md)
* [Design the Twitter timeline and search](../exercises/system-design/twitter/README.md)

And **some** object-oriented design questions:

* [Design an LRU cache](../exercises/object-oriented-design/lru-cache.md)
* [Design a hash map](../exercises/object-oriented-design/hash-map.md)

## 5. Skim, don't study, the context material

* A few [real world architectures](../appendix/real-world-architectures.md)
* A few articles from the [company engineering blogs](../appendix/engineering-blogs.md) for the companies you are interviewing with
* Skim the [additional system design interview questions](../appendix/additional-questions.md) list so the problem statements aren't a surprise

## Self-check

1. Can you name the four steps of the interview framework without looking?
2. For each component in the breadth list, can you give one sentence on what problem it solves and one disadvantage?
3. Can you convert "10 million writes per month" into writes per second from memory?
