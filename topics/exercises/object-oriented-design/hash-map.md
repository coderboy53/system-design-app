---
id: ood-hash-map
title: Design a Hash Map
module: exercises/object-oriented-design
order: 10
est_minutes: 20
timelines: [short, medium, long]
tags: [exercise, ood, hash-table, data-structures]
prerequisites: [interview-approach]
related: [key-value-store, caching-overview]
source: https://github.com/donnemartin/system-design-primer/blob/master/solutions/object_oriented_design/hash_table/hash_map.ipynb
---

# Design a Hash Map

> **How to use this exercise:** clarify the constraints and assumptions first, sketch the
> classes yourself, then compare against the solution below. Object-oriented design questions
> follow the same opening move as [system design questions](../../study-guide/interview-approach.md#step-1-outline-use-cases-constraints-and-assumptions)
> — establish scope before writing anything.

## Constraints and assumptions

* For simplicity, are the keys integers only?
    * Yes
* For collision resolution, can we use chaining?
    * Yes
* Do we have to worry about load factors?
    * No
* Can we assume inputs are valid or do we have to validate them?
    * Assume they're valid
* Can we assume this fits memory?
    * Yes

## Solution

```python
class Item(object):

    def __init__(self, key, value):
        self.key = key
        self.value = value


class HashTable(object):

    def __init__(self, size):
        self.size = size
        self.table = [[] for _ in range(self.size)]

    def _hash_function(self, key):
        return key % self.size

    def set(self, key, value):
        hash_index = self._hash_function(key)
        for item in self.table[hash_index]:
            if item.key == key:
                item.value = value
                return
        self.table[hash_index].append(Item(key, value))

    def get(self, key):
        hash_index = self._hash_function(key)
        for item in self.table[hash_index]:
            if item.key == key:
                return item.value
        raise KeyError('Key not found')

    def remove(self, key):
        hash_index = self._hash_function(key)
        for index, item in enumerate(self.table[hash_index]):
            if item.key == key:
                del self.table[hash_index][index]
                return
        raise KeyError('Key not found')
```

## Self-check

1. What collision resolution strategy does this implementation use, and what is the alternative?
2. What is the average-case and worst-case lookup complexity here?
3. What would change if you had to handle load factors and resizing?
