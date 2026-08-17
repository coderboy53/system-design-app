---
id: interview-approach
title: How to Approach a System Design Interview Question
module: study-guide
order: 5
est_minutes: 8
timelines: [short, medium, long]
tags: [interview-prep, framework, process]
prerequisites: []
related: [study-guide, back-of-the-envelope, powers-of-two, latency-numbers]
source: https://github.com/donnemartin/system-design-primer#how-to-approach-a-system-design-interview-question
---

# How to Approach a System Design Interview Question

> How to tackle a system design interview question.

The system design interview is an **open-ended conversation**. **You are expected to lead it.**

Use the following steps to guide the discussion. To solidify the process, work through the
[system design exercises](../exercises/system-design/README.md) using these steps.

## Step 1: Outline use cases, constraints, and assumptions

Gather requirements and scope the problem. Ask questions to clarify use cases and constraints.
Discuss assumptions.

* Who is going to use it?
* How are they going to use it?
* How many users are there?
* What does the system do?
* What are the inputs and outputs of the system?
* How much data do we expect to handle?
* How many requests per second do we expect?
* What is the expected read to write ratio?

## Step 2: Create a high level design

Outline a high level design with all important components.

* Sketch the main components and connections
* Justify your ideas

## Step 3: Design core components

Dive into details for each core component. For example, if you were asked to
[design a url shortening service](../exercises/system-design/pastebin/README.md), discuss:

* Generating and storing a hash of the full url
    * [MD5](../exercises/system-design/pastebin/README.md) and [Base62](../exercises/system-design/pastebin/README.md)
    * Hash collisions
    * [SQL or NoSQL](../data/sql-or-nosql.md)
    * Database schema
* Translating a hashed url to the full url
    * Database lookup
* API and object-oriented design

## Step 4: Scale the design

Identify and address bottlenecks, given the constraints. For example, do you need the following to
address scalability issues?

* [Load balancer](../networking/load-balancer.md)
* [Horizontal scaling](../networking/load-balancer.md#horizontal-scaling)
* [Caching](../caching/README.md)
* [Database sharding](../data/sharding.md)

Discuss potential solutions and trade-offs. **Everything is a trade-off.** Address bottlenecks using
[principles of scalable system design](../README.md#index-of-system-design-topics).

> Do not jump straight to the final design. State that you would do this iteratively:
> 1) **benchmark / load test**, 2) **profile** for bottlenecks, 3) address bottlenecks while
> evaluating alternatives and trade-offs, and 4) repeat. See
> [scaling to millions of users on AWS](../exercises/system-design/scaling-aws/README.md) for a
> worked example of iterative scaling.

## Back-of-the-envelope calculations

You might be asked to do some estimates by hand. Refer to:

* [Back-of-the-envelope calculations](../appendix/back-of-the-envelope.md)
* [Powers of two table](../appendix/powers-of-two.md)
* [Latency numbers every programmer should know](../appendix/latency-numbers.md)
* [Use back of the envelope calculations](http://highscalability.com/blog/2011/1/26/google-pro-tip-use-back-of-the-envelope-calculations-to-choo.html)

## Source(s) and further reading

Check out the following links to get a better idea of what to expect:

* [How to ace a systems design interview](https://web.archive.org/web/20210505130322/https://www.palantir.com/2011/10/how-to-rock-a-systems-design-interview/)
* [The system design interview](http://www.hiredintech.com/system-design)
* [Intro to Architecture and Systems Design Interviews](https://www.youtube.com/watch?v=ZgdS0EUmn70)
* [System design template](https://leetcode.com/discuss/career/229177/My-System-Design-Template)

## Self-check

1. List the four steps in order, and the one-line goal of each.
2. Name five clarifying questions you would ask in step 1 for any product-shaped question.
3. Why is jumping straight to the scaled design a mistake even if the final design is correct?
4. What are the four stages of the iterative scaling loop?
