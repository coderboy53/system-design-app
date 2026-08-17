---
id: asynchronism
title: Asynchronism
module: application
order: 40
est_minutes: 7
timelines: [short, medium, long]
tags: [application, asynchronism, queues, throughput, back-pressure]
prerequisites: [application-layer, latency-vs-throughput]
related: [latency-vs-throughput, microservices, caching-overview]
source: https://github.com/donnemartin/system-design-primer#asynchronism
---

# Asynchronism

<p align="center">
  <img src="../assets/images/54GYsSx.png" alt="Asynchronism">
  <br/>
  <i><a href="http://lethain.com/introduction-to-architecting-systems-for-scale/#platform_layer">Source: Intro to architecting systems for scale</a></i>
</p>

Asynchronous workflows help **reduce request times for expensive operations** that would otherwise
be performed in-line. They can also help by doing **time-consuming work in advance**, such as
periodic aggregation of data.

This is the primary lever for trading
[latency against throughput](../fundamentals/latency-vs-throughput.md).

## Message queues

Message queues **receive, hold, and deliver messages**. If an operation is too slow to perform
inline, you can use a message queue with the following workflow:

* An application **publishes a job** to the queue, then notifies the user of job status
* A **worker picks up the job** from the queue, processes it, then signals the job is complete

The user is **not blocked** and the job is processed in the background. During this time, the client
might optionally do a small amount of processing to make it seem like the task has completed. For
example, if posting a tweet, the tweet could be instantly posted to your timeline, but it could take
some time before your tweet is actually delivered to all of your followers.

**[Redis](https://redis.io/)** is useful as a simple message broker but **messages can be lost**.

**[RabbitMQ](https://www.rabbitmq.com/)** is popular but requires you to adapt to the 'AMQP' protocol
and **manage your own nodes**.

**[Amazon SQS](https://aws.amazon.com/sqs/)** is hosted but can have **high latency** and has the
possibility of **messages being delivered twice**.

## Task queues

Task queues **receive tasks and their related data, run them, then deliver their results**. They can
support **scheduling** and can be used to run **computationally-intensive jobs** in the background.

**[Celery](https://docs.celeryproject.org/en/stable/)** has support for scheduling and primarily has
Python support.

## Back pressure

If queues start to grow significantly, the **queue size can become larger than memory**, resulting
in cache misses, disk reads, and even slower performance.
[Back pressure](http://mechanical-sympathy.blogspot.com/2012/05/apply-back-pressure-when-overloaded.html)
can help by **limiting the queue size**, thereby maintaining a high throughput rate and good
response times for jobs already in the queue.

Once the queue fills up, clients get a **server busy or HTTP 503** status code to try again later.
Clients can retry the request at a later time, perhaps with
[exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff).

> Back pressure is a deliberate choice to fail some requests fast rather than degrade every request.
> Without it, an overloaded queue takes the whole system down with it.

## Disadvantage(s): asynchronism

* Use cases such as **inexpensive calculations and realtime workflows** might be better suited for
  synchronous operations, as introducing queues can add **delays and complexity**.

## Source(s) and further reading

* [It's all a numbers game](https://www.youtube.com/watch?v=1KRYH75wgy4)
* [Applying back pressure when overloaded](http://mechanical-sympathy.blogspot.com/2012/05/apply-back-pressure-when-overloaded.html)
* [Little's law](https://en.wikipedia.org/wiki/Little%27s_law)
* [What is the difference between a message queue and a task queue?](https://www.quora.com/What-is-the-difference-between-a-message-queue-and-a-task-queue-Why-would-a-task-queue-require-a-message-broker-like-RabbitMQ-Redis-Celery-or-IronMQ-to-function)

## Self-check

1. Describe the publish/worker workflow of a message queue, and what the user sees while it runs.
2. Name the trade-off of Redis, RabbitMQ, and SQS as brokers — one weakness each.
3. What is the difference between a message queue and a task queue?
4. What goes wrong when a queue grows past memory, and what does back pressure do about it?
5. What status code do clients get when back pressure kicks in, and how should they respond?
6. When is asynchronism the wrong choice?
