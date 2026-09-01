---
id: ood-exercises
title: Object-Oriented Design Exercises
module: exercises/object-oriented-design
module_order: 110
order: 0
est_minutes: 2
timelines: [short, medium, long]
tags: [index, exercise, ood, practice]
prerequisites: [interview-approach]
related: [exercises, system-design-exercises]
source: https://github.com/donnemartin/system-design-primer#object-oriented-design-interview-questions-with-solutions
---

# Object-Oriented Design Exercises

> Common object-oriented design interview questions with sample discussions and code.

> **Note:** the primer marks this section as under development. Solutions are Python, converted here
> from the original Jupyter notebooks.

Same opening move as a [system design question](../../study-guide/interview-approach.md#step-1-outline-use-cases-constraints-and-assumptions):
**establish constraints and assumptions before writing a single class.** Each exercise below opens
with the clarifying Q&A that scoped it.

| Exercise | Time | Track | Core themes |
|---|---|---|---|
| [Design a hash map](hash-map.md) | 20 min | short · medium · long | Chaining, buckets, [key-value abstraction](../../data/key-value-store.md) |
| [Design a least recently used cache](lru-cache.md) | 20 min | short · medium · long | Hash map + doubly linked list, [LRU eviction](../../caching/cache-locations.md#application-caching) |
| [Design a call center](call-center.md) | 25 min | medium · long | Escalation chain, inheritance, dispatch |
| [Design a deck of cards](deck-of-cards.md) | 20 min | medium · long | Generic base + game-specific extension |
| [Design a parking lot](parking-lot.md) | 25 min | medium · long | Vehicle/spot compatibility, multi-level allocation |
| [Design a chat server](online-chat.md) | 25 min | long | Users, friend requests, group vs private chats |

## Suggested order

Start with **[hash map](hash-map.md)** and **[LRU cache](lru-cache.md)** — they are the two that
also show up as components inside system design answers ([caching](../../caching/README.md),
[key-value stores](../../data/key-value-store.md), and the
[query cache exercise](../system-design/query-cache/README.md)).

## Contribute

The primer lists these as open: design a circular array, and
[add an object-oriented design question](https://github.com/donnemartin/system-design-primer#contributing).
