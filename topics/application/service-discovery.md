---
id: service-discovery
title: Service Discovery
module: application
order: 30
est_minutes: 3
timelines: [medium, long]
tags: [application, architecture, service-discovery, operations]
prerequisites: [microservices]
related: [microservices, key-value-store, http, application-layer]
source: https://github.com/donnemartin/system-design-primer#service-discovery
---

# Service Discovery

Systems such as [Consul](https://www.consul.io/docs/index.html),
[Etcd](https://coreos.com/etcd/docs/latest), and
[Zookeeper](http://www.slideshare.net/sauravhaloi/introduction-to-apache-zookeeper) can help
services find each other by **keeping track of registered names, addresses, and ports**.

[Health checks](https://www.consul.io/intro/getting-started/checks.html) help verify service
integrity and are often done using an [HTTP](../networking/http.md) endpoint.

Both Consul and Etcd have a built-in [key-value store](../data/key-value-store.md) that can be
useful for storing config values and other shared data.

## Why it exists

In a [microservices](microservices.md) deployment, instances come and go — autoscaling, deploys,
failures. Hardcoding addresses stops working the moment servers are
[cloned](../networking/load-balancer.md#horizontal-scaling). The discovery layer is the moving
address book, and the health check is what keeps dead instances out of it — the same job a
[load balancer](../networking/load-balancer.md) does for a fixed pool.

## Source(s) and further reading

* [Consul documentation](https://www.consul.io/docs/index.html)
* [Etcd documentation](https://coreos.com/etcd/docs/latest)
* [Introduction to Zookeeper](http://www.slideshare.net/sauravhaloi/introduction-to-apache-zookeeper)

## Self-check

1. What three pieces of information does a service registry track?
2. How are health checks typically implemented, and what do they prevent?
3. Why do Consul and Etcd ship a key-value store alongside discovery?
