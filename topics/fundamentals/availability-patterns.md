---
id: availability-patterns
title: Availability Patterns
module: fundamentals
order: 60
est_minutes: 9
timelines: [short, medium, long]
tags: [availability, failover, replication, reliability, math]
prerequisites: [cap-theorem]
related: [consistency-patterns, replication, load-balancer]
source: https://github.com/donnemartin/system-design-primer#availability-patterns
---

# Availability Patterns

There are two complementary patterns to support high availability: **fail-over** and **replication**.

## Fail-over

### Active-passive

With active-passive fail-over, **heartbeats** are sent between the active and the passive server on
standby. If the heartbeat is interrupted, the passive server takes over the active's IP address and
resumes service.

The length of downtime is determined by whether the passive server is already running in **'hot'
standby** or whether it needs to start up from **'cold' standby**. Only the active server handles
traffic.

Active-passive failover can also be referred to as **master-slave failover**.

### Active-active

In active-active, **both servers are managing traffic**, spreading the load between them.

If the servers are public-facing, the [DNS](../networking/dns.md) would need to know about the
public IPs of both servers. If the servers are internal-facing, application logic would need to know
about both servers.

Active-active failover can also be referred to as **master-master failover**.

Both modes are commonly used to protect the [load balancer](../networking/load-balancer.md) itself,
which would otherwise be a single point of failure.

### Disadvantage(s): failover

* Fail-over adds more hardware and additional complexity.
* There is a potential for **loss of data** if the active system fails before any newly written data can be replicated to the passive.

## Replication

### Master-slave and master-master

This topic is discussed in depth in the [data module](../data/replication.md):

* [Master-slave replication](../data/replication.md#master-slave-replication)
* [Master-master replication](../data/replication.md#master-master-replication)

## Availability in numbers

Availability is often quantified by **uptime** (or downtime) as a percentage of time the service is
available. Availability is generally measured in **number of 9s** — a service with 99.99%
availability is described as having **four 9s**.

### 99.9% availability — three 9s

| Duration            | Acceptable downtime |
|---------------------|---------------------|
| Downtime per year   | 8h 45min 57s        |
| Downtime per month  | 43m 49.7s           |
| Downtime per week   | 10m 4.8s            |
| Downtime per day    | 1m 26.4s            |

### 99.99% availability — four 9s

| Duration            | Acceptable downtime |
|---------------------|---------------------|
| Downtime per year   | 52min 35.7s         |
| Downtime per month  | 4m 23s              |
| Downtime per week   | 1m 5s               |
| Downtime per day    | 8.6s                |

### Availability in parallel vs in sequence

If a service consists of multiple components prone to failure, the service's overall availability
depends on whether the components are **in sequence** or **in parallel**.

#### In sequence

Overall availability **decreases** when two components with availability < 100% are in sequence:

```
Availability (Total) = Availability (Foo) * Availability (Bar)
```

If both `Foo` and `Bar` each had 99.9% availability, their total availability in sequence would be
**99.8%**.

#### In parallel

Overall availability **increases** when two components with availability < 100% are in parallel:

```
Availability (Total) = 1 - (1 - Availability (Foo)) * (1 - Availability (Bar))
```

If both `Foo` and `Bar` each had 99.9% availability, their total availability in parallel would be
**99.9999%**.

> This is the whole argument for redundancy in one formula. Every component you add to the request
> path multiplies availability down; every component you duplicate alongside an existing one
> multiplies *unavailability* down.

## Self-check

1. What is the difference between hot and cold standby, and how does it affect downtime?
2. In active-active, what has to know about both servers when they are public-facing? When internal-facing?
3. What is the data-loss risk inherent in any fail-over setup?
4. Three components at 99.9% each, in sequence — roughly what is the total availability?
5. Two components at 99% each, in parallel — what is the total availability?
6. How much downtime per month does four 9s allow?
