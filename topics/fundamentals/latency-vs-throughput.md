---
id: latency-vs-throughput
title: Latency vs Throughput
module: fundamentals
order: 30
est_minutes: 2
timelines: [short, medium, long]
tags: [tradeoffs, performance, metrics]
prerequisites: [performance-vs-scalability]
related: [cap-theorem, asynchronism, latency-numbers]
source: https://github.com/donnemartin/system-design-primer#latency-vs-throughput
---

# Latency vs Throughput

**Latency** is the time to perform some action or to produce some result.

**Throughput** is the number of such actions or results per unit of time.

Generally, you should aim for **maximal throughput** with **acceptable latency**.

## Where the trade-off shows up

* [Batching and asynchronism](../application/asynchronism.md) raise throughput by accepting higher
  latency per request. A [message queue](../application/asynchronism.md#message-queues) makes the
  user's request return immediately while the actual work happens later.
* [Back pressure](../application/asynchronism.md#back-pressure) protects throughput by rejecting
  work rather than letting queues grow past memory and degrade everything.
* [TCP](../networking/tcp-and-udp.md#transmission-control-protocol-tcp) guarantees delivery at the
  cost of latency; [UDP](../networking/tcp-and-udp.md#user-datagram-protocol-udp) drops the
  guarantees to minimize it.

For concrete numbers to reason with, see
[latency numbers every programmer should know](../appendix/latency-numbers.md).

## Source(s) and further reading

* [Understanding latency vs throughput](https://community.cadence.com/cadence_blogs_8/b/fv/posts/understanding-latency-vs-throughput)

## Self-check

1. Define latency and throughput in one sentence each.
2. What is the general goal stated in terms of both?
3. Give an example of a change that increases throughput while increasing latency.
