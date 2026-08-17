---
id: replication
title: Database Replication
module: data
order: 20
est_minutes: 7
timelines: [short, medium, long]
tags: [database, sql, replication, availability, scaling]
prerequisites: [rdbms, availability-patterns]
related: [federation, sharding, availability-patterns, consistency-patterns]
source: https://github.com/donnemartin/system-design-primer#master-slave-replication
---

# Database Replication

Replication is one of the two complementary
[availability patterns](../fundamentals/availability-patterns.md), and one of the primary techniques
for scaling a [relational database](rdbms.md).

## Master-slave replication

The master serves **reads and writes**, replicating writes to one or more slaves, which serve
**only reads**. Slaves can also replicate to additional slaves in a **tree-like fashion**. If the
master goes offline, the system can continue to operate in **read-only mode** until a slave is
promoted to a master or a new master is provisioned.

<p align="center">
  <img src="../assets/images/C9ioGtn.png" alt="Master-slave replication">
  <br/>
  <i><a href="http://www.slideshare.net/jboner/scalability-availability-stability-patterns/">Source: Scalability, availability, stability, patterns</a></i>
</p>

This is the read-scaling workhorse: it fits the common 100:1 or 1000:1 read-to-write ratio, and
pairs naturally with [active-passive fail-over](../fundamentals/availability-patterns.md#active-passive).

### Disadvantage(s): master-slave replication

* Additional logic is needed to **promote a slave to a master**.
* See [Disadvantage(s): replication](#disadvantages-replication) for points related to **both** master-slave and master-master.

## Master-master replication

Both masters serve **reads and writes** and coordinate with each other on writes. If either master
goes down, the system can continue to operate with **both reads and writes**.

<p align="center">
  <img src="../assets/images/krAHLGg.png" alt="Master-master replication">
  <br/>
  <i><a href="http://www.slideshare.net/jboner/scalability-availability-stability-patterns/">Source: Scalability, availability, stability, patterns</a></i>
</p>

This is the database analogue of [active-active fail-over](../fundamentals/availability-patterns.md#active-active).

### Disadvantage(s): master-master replication

* You'll need a [load balancer](../networking/load-balancer.md) or you'll need to change your
  application logic to determine **where to write**.
* Most master-master systems are either **loosely consistent** (violating [ACID](rdbms.md)) or have
  **increased write latency** due to synchronization.
* **Conflict resolution** comes more into play as more write nodes are added and as latency increases.
* See [Disadvantage(s): replication](#disadvantages-replication) for points related to **both** master-slave and master-master.

## Disadvantage(s): replication

These apply to **both** topologies:

* There is a potential for **loss of data** if the master fails before any newly written data can be replicated to other nodes.
* Writes are replayed to the read replicas. If there are a lot of writes, the **read replicas can get bogged down with replaying writes** and can't do as many reads.
* The more read slaves, the more you have to replicate, which leads to **greater replication lag**.
* On some systems, writing to the master can spawn multiple threads to write in parallel, whereas **read replicas only support writing sequentially with a single thread**.
* Replication adds **more hardware and additional complexity**.

## At a glance

| | Master-slave | Master-master |
|---|---|---|
| Writes served by | master only | both masters |
| Reads served by | master + slaves | both masters |
| If a node dies | read-only until promotion | reads and writes continue |
| Extra machinery | promotion logic | write routing, conflict resolution |
| Consistency cost | replication lag on reads | loose consistency or write latency |

## Source(s) and further reading: replication

* [Scalability, availability, stability, patterns](http://www.slideshare.net/jboner/scalability-availability-stability-patterns/)
* [Multi-master replication](https://en.wikipedia.org/wiki/Multi-master_replication)

## Self-check

1. What happens to writes in a master-slave setup when the master goes offline?
2. Why does adding more read replicas eventually make replication lag worse rather than better?
3. What two problems does master-master introduce that master-slave does not?
4. Name the data-loss window that both topologies share.
5. Why can replicas fall behind on replaying writes even when they have spare capacity for reads?
