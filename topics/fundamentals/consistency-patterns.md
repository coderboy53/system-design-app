---
id: consistency-patterns
title: Consistency Patterns
module: fundamentals
order: 50
est_minutes: 5
timelines: [short, medium, long]
tags: [consistency, distributed-systems, tradeoffs]
prerequisites: [cap-theorem]
related: [availability-patterns, replication, nosql, cache-update-strategies]
source: https://github.com/donnemartin/system-design-primer#consistency-patterns
---

# Consistency Patterns

With multiple copies of the same data, we are faced with options on how to synchronize them so
clients have a consistent view of the data. Recall the definition of consistency from the
[CAP theorem](cap-theorem.md) — every read receives the most recent write or an error.

## Weak consistency

After a write, reads **may or may not** see it. A best effort approach is taken.

This approach is seen in systems such as **memcached**. Weak consistency works well in real time use
cases such as **VoIP, video chat, and realtime multiplayer games**. For example, if you are on a
phone call and lose reception for a few seconds, when you regain connection you do not hear what was
spoken during connection loss.

Compare with [UDP](../networking/tcp-and-udp.md#user-datagram-protocol-udp), which makes the same
bargain at the protocol level: late data is worse than lost data.

## Eventual consistency

After a write, reads will **eventually** see it (typically within milliseconds). Data is replicated
**asynchronously**.

This approach is seen in systems such as [**DNS**](../networking/dns.md) and **email**. Eventual consistency works well in
**highly available** systems.

This is the consistency model of [AP systems](cap-theorem.md#ap--availability-and-partition-tolerance),
of [BASE / NoSQL stores](../data/nosql.md), and of
[master-slave replication](../data/replication.md#master-slave-replication) with async replicas.

## Strong consistency

After a write, reads **will** see it. Data is replicated **synchronously**.

This approach is seen in **file systems and RDBMSes**. Strong consistency works well in systems
that need **transactions**.

See [ACID](../data/rdbms.md) and [CP systems](cap-theorem.md#cp--consistency-and-partition-tolerance).

## At a glance

| Pattern | After a write, reads… | Replication | Seen in | Good for |
|---|---|---|---|---|
| **Weak** | may or may not see it | best effort | memcached | VoIP, video chat, realtime games |
| **Eventual** | eventually see it (ms) | asynchronous | DNS, email | highly available systems |
| **Strong** | see it | synchronous | file systems, RDBMS | systems needing transactions |

## Source(s) and further reading

* [Transactions across data centers](http://snarfed.org/transactions_across_datacenters_io.html)

## Self-check

1. Order the three patterns from weakest to strongest guarantee, and give the replication style of each.
2. Why is weak consistency acceptable for a phone call but not for a bank balance?
3. Which consistency pattern does DNS use, and which does an RDBMS use?
4. Which pattern does an [AP system](cap-theorem.md) tend toward, and which does a [CP system](cap-theorem.md)?
