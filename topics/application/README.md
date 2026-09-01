---
id: application
title: Application Layer
module: application
module_order: 40
order: 0
est_minutes: 2
timelines: [short, medium, long]
tags: [index, application, architecture]
prerequisites: [networking]
related: [application-layer, microservices, service-discovery, asynchronism]
source: https://github.com/donnemartin/system-design-primer#application-layer
---

# Application Layer

<p align="center">
  <img src="../assets/images/yB5SYwm.png" alt="Application layer">
  <br/>
  <i><a href="http://lethain.com/introduction-to-architecting-systems-for-scale/#platform_layer">Source: Intro to architecting systems for scale</a></i>
</p>

What sits behind the [load balancer](../networking/load-balancer.md): the platform tier, how it is
decomposed into services, how those services find each other, and how work gets pushed off the
request path.

## Modules

| # | Module | Time | What it answers |
|---|---|---|---|
| 1 | [Application layer](application-layer.md) | 4 min | Why separate the web tier from the platform tier? |
| 2 | [Microservices](microservices.md) | 3 min | How small should the services be? |
| 3 | [Service discovery](service-discovery.md) | 3 min | How do services find each other? |
| 4 | [Asynchronism](asynchronism.md) | 7 min | How do I get slow work off the request path? |

## Where this leads

* Services are stateless so they can be [horizontally scaled](../networking/load-balancer.md#horizontal-scaling); state moves to the [database](../data/README.md) or a [cache](../caching/README.md).
* Internal service calls typically use [RPC](../networking/rpc.md); public ones use [REST](../networking/rest.md).
* Asynchronism is the main lever for trading [latency for throughput](../fundamentals/latency-vs-throughput.md).
