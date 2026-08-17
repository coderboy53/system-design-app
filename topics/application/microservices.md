---
id: microservices
title: Microservices
module: application
order: 20
est_minutes: 3
timelines: [short, medium, long]
tags: [application, architecture, microservices]
prerequisites: [application-layer]
related: [service-discovery, rpc, rest, application-layer]
source: https://github.com/donnemartin/system-design-primer#microservices
---

# Microservices

Related to the [application layer](application-layer.md) discussion are
[microservices](https://en.wikipedia.org/wiki/Microservices), which can be described as a suite of
**independently deployable, small, modular services**. Each service runs a unique process and
communicates through a well-defined, lightweight mechanism to serve a business
goal.<sup>[1](https://smartbear.com/learn/api-design/what-are-microservices)</sup>

Pinterest, for example, could have the following microservices: user profile, follower, feed,
search, photo upload, etc.

## What this implies

* Services need a way to locate each other — see [service discovery](service-discovery.md).
* Internal communication is typically [RPC](../networking/rpc.md) for performance; external
  interfaces are typically [REST](../networking/rest.md).
* Each service can be [scaled horizontally](../networking/load-balancer.md#horizontal-scaling) on
  its own, which is the point.
* Long-running work between services belongs on a [queue](asynchronism.md#message-queues), not on
  a synchronous call chain.

## Disadvantage(s)

* Microservices add **complexity in terms of deployments and operations**.
* Loosely coupled services require a **different architectural, operations, and process approach**
  than a monolith.

## Source(s) and further reading

* [What are microservices](https://smartbear.com/learn/api-design/what-are-microservices)
* [Here's what you need to know about building microservices](https://cloudncode.wordpress.com/2016/07/22/msa-getting-started/)
* [Service oriented architecture](https://en.wikipedia.org/wiki/Service-oriented_architecture)

## Self-check

1. Define a microservice in terms of deployability, size, and communication.
2. Give a plausible microservice decomposition for a photo-sharing product.
3. What operational cost do you accept in exchange for independent deployability?
