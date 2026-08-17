---
id: ood-lru-cache
title: Design an LRU Cache
module: exercises/object-oriented-design
order: 20
est_minutes: 20
timelines: [short, medium, long]
tags: [exercise, ood, cache, lru, data-structures]
prerequisites: [interview-approach]
related: [cache-locations, cache-update-strategies, caching-overview]
source: https://github.com/donnemartin/system-design-primer/blob/master/solutions/object_oriented_design/lru_cache/lru_cache.ipynb
---

# Design an LRU Cache

> **How to use this exercise:** clarify the constraints and assumptions first, sketch the
> classes yourself, then compare against the solution below. Object-oriented design questions
> follow the same opening move as [system design questions](../../study-guide/interview-approach.md#step-1-outline-use-cases-constraints-and-assumptions)
> — establish scope before writing anything.

## Constraints and assumptions

* What are we caching?
    * We are caching the results of web queries
* Can we assume inputs are valid or do we have to validate them?
    * Assume they're valid
* Can we assume this fits memory?
    * Yes

## Solution

```python
class Node(object):

    def __init__(self, results):
        self.results = results
        self.prev = None
        self.next = None


class LinkedList(object):

    def __init__(self):
        self.head = None
        self.tail = None

    def move_to_front(self, node):  # ...
    def append_to_front(self, node):  # ...
    def remove_from_tail(self):  # ...


class Cache(object):

    def __init__(self, MAX_SIZE):
        self.MAX_SIZE = MAX_SIZE
        self.size = 0
        self.lookup = {}  # key: query, value: node
        self.linked_list = LinkedList()

    def get(self, query)
        """Get the stored query result from the cache.
        
        Accessing a node updates its position to the front of the LRU list.
        """
        node = self.lookup.get(query)
        if node is None:
            return None
        self.linked_list.move_to_front(node)
        return node.results

    def set(self, results, query):
        """Set the result for the given query key in the cache.
        
        When updating an entry, updates its position to the front of the LRU list.
        If the entry is new and the cache is at capacity, removes the oldest entry
        before the new entry is added.
        """
        node = self.lookup.get(query)
        if node is not None:
            # Key exists in cache, update the value
            node.results = results
            self.linked_list.move_to_front(node)
        else:
            # Key does not exist in cache
            if self.size == self.MAX_SIZE:
                # Remove the oldest entry from the linked list and lookup
                self.lookup.pop(self.linked_list.tail.query, None)
                self.linked_list.remove_from_tail()
            else:
                self.size += 1
            # Add the new key and value
            new_node = Node(results)
            self.linked_list.append_to_front(new_node)
            self.lookup[query] = new_node
```

## Self-check

1. Which two data structures are combined, and what does each one give you?
2. Why must the linked list be doubly linked rather than singly linked?
3. Where does this exact structure show up in a real system?
