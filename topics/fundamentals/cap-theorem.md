---
id: cap-theorem
title: Availability vs Consistency (CAP Theorem)
module: fundamentals
order: 40
est_minutes: 6
timelines: [short, medium, long]
tags: [tradeoffs, distributed-systems, consistency, availability, cap]
prerequisites: [latency-vs-throughput]
related: [consistency-patterns, availability-patterns, nosql, sql-or-nosql]
source: https://github.com/donnemartin/system-design-primer#availability-vs-consistency
---

# Availability vs Consistency (CAP Theorem)

<p align="center">
  <img src="../assets/images/bgLMI2u.png" alt="CAP theorem">
  <br/>
  <i><a href="https://robertgreiner.com/cap-theorem-revisited">Source: CAP theorem revisited</a></i>
</p>

In a distributed computer system, you can only support **two of the following three guarantees**:

* **Consistency** — every read receives the most recent write or an error
* **Availability** — every request receives a response, without guarantee that it contains the most recent version of the information
* **Partition Tolerance** — the system continues to operate despite arbitrary partitioning due to network failures

*Networks aren't reliable, so you'll need to support partition tolerance. You'll need to make a
software trade-off between consistency and availability.*

That last point is the practical takeaway: **P is not optional**, so the real choice is CP or AP.

## CP — consistency and partition tolerance

Waiting for a response from the partitioned node might result in a **timeout error**. CP is a good
choice if your business needs require **atomic reads and writes**.

Think: financial ledgers, inventory decrements, anything where serving stale data is worse than
serving an error. See [strong consistency](consistency-patterns.md#strong-consistency) and
[RDBMS transactions](../data/rdbms.md).

## AP — availability and partition tolerance

Responses return the **most readily available version** of the data available on any node, which
might not be the latest. Writes might take some time to propagate when the partition is resolved.

AP is a good choice if the business needs to allow for
[eventual consistency](consistency-patterns.md#eventual-consistency), or when the system needs to
continue working despite external errors.

Think: social feeds, view counters, [DNS](../networking/dns.md). See [NoSQL and BASE](../data/nosql.md), which explicitly
chooses availability over consistency.

## Source(s) and further reading

* [CAP theorem revisited](https://robertgreiner.com/cap-theorem-revisited/)
* [A plain english introduction to CAP theorem](http://ksat.me/a-plain-english-introduction-to-cap-theorem)
* [CAP FAQ](https://github.com/henryr/cap-faq)
* [The CAP theorem](https://www.youtube.com/watch?v=k-Yaq8AHlFA)

## Self-check

1. State the three CAP guarantees and define each precisely.
2. Why is partition tolerance effectively mandatory in a real distributed system?
3. What does a CP system do when a node is partitioned away, and when is that the right behavior?
4. What does an AP system return during a partition, and what happens once the partition heals?
