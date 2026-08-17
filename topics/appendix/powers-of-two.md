---
id: powers-of-two
title: Powers of Two Table
module: appendix
order: 20
est_minutes: 2
timelines: [short, medium, long]
tags: [reference, estimation, math]
prerequisites: []
related: [latency-numbers, back-of-the-envelope, interview-approach]
source: https://github.com/donnemartin/system-design-primer#powers-of-two-table
---

# Powers of Two Table

```
Power           Exact Value         Approx Value        Bytes
---------------------------------------------------------------
7                             128
8                             256
10                           1024   1 thousand           1 KB
16                         65,536                       64 KB
20                      1,048,576   1 million            1 MB
30                  1,073,741,824   1 billion            1 GB
32                  4,294,967,296                        4 GB
40              1,099,511,627,776   1 trillion           1 TB
```

## How to use it

The whole table collapses to one fact worth memorizing: **2^10 ≈ 1 thousand**, and every additional
10 in the exponent multiplies by another thousand. So 2^20 ≈ 1 million (1 MB), 2^30 ≈ 1 billion
(1 GB), 2^40 ≈ 1 trillion (1 TB).

Two more worth knowing on sight:

* 2^8 = 256 — one byte, and why `VARCHAR(255)` is [the largest length countable in 8 bits](../data/sql-tuning.md#tighten-up-the-schema).
* 2^32 ≈ 4 billion — the range of a 32-bit `INT`, and why [`INT` handles numbers up to about 4 billion](../data/sql-tuning.md#tighten-up-the-schema).

## Source(s) and further reading

* [Powers of two](https://en.wikipedia.org/wiki/Power_of_two)

## Self-check

1. Approximately what value is 2^20, and how many bytes is that?
2. How many rows can a 32-bit auto-increment primary key address?
3. A record is 1.27 KB and you write 10 million per month. Roughly how many GB per month?
