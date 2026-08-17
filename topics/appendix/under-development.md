---
id: under-development
title: Under Development
module: appendix
order: 80
est_minutes: 2
timelines: [long]
tags: [reference, gaps, advanced]
prerequisites: [sharding]
related: [sharding, real-world-architectures, long-timeline]
source: https://github.com/donnemartin/system-design-primer#under-development
---

# Under Development

The primer flags these sections as in-progress. They are **not covered in depth anywhere in this
path**, and that is a gap you should close from primary sources if you are on the
[long timeline](../study-guide/long-timeline.md) — all three come up in real interviews.

## Distributed computing with MapReduce

Referenced in passing by the [pastebin](../exercises/system-design/pastebin/README.md) and
[sales rank](../exercises/system-design/sales-rank/README.md) exercises, which use MapReduce jobs to
compute analytics offline.

* [MapReduce paper (Google)](http://static.googleusercontent.com/media/research.google.com/zh-CN/us/archive/mapreduce-osdi04.pdf)
* [Spark architecture](http://www.slideshare.net/AGrishchenko/apache-spark-architecture)

## Consistent hashing

Referenced by [sharding](../data/sharding.md) as the technique that reduces how much data moves when
you rebalance shards.

* [The magic of consistent hashing](http://www.paperplanes.de/2011/12/9/the-magic-of-consistent-hashing.html)
* [Dynamo paper (Amazon)](http://www.read.seas.harvard.edu/~kohler/class/cs239-w08/decandia07dynamo.pdf) — consistent hashing in production

## Scatter gather

The pattern behind fan-out reads across shards or services: send the request to N nodes in parallel,
combine the responses. It is what makes [sharded](../data/sharding.md) queries and
[search](additional-questions.md) work, and its cost is that the slowest node sets the latency.

## Contributing

The primer welcomes [contributions](https://github.com/donnemartin/system-design-primer#contributing)
to complete these sections.

## Self-check

1. Why does consistent hashing reduce data transfer during shard rebalancing?
2. In a scatter-gather read, what determines the overall latency?
3. Where does MapReduce show up in the exercises in this path, and what job does it do there?
