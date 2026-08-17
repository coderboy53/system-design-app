---
id: sql-or-nosql
title: SQL or NoSQL
module: data
order: 120
est_minutes: 4
timelines: [short, medium, long]
tags: [database, sql, nosql, tradeoffs, decision]
prerequisites: [rdbms, nosql]
related: [rdbms, nosql, cap-theorem, key-value-store]
source: https://github.com/donnemartin/system-design-primer#sql-or-nosql
---

# SQL or NoSQL

<p align="center">
  <img src="../assets/images/wXGqG5f.png" alt="SQL vs NoSQL">
  <br/>
  <i><a href="https://www.infoq.com/articles/Transition-RDBMS-NoSQL/">Source: Transitioning from RDBMS to NoSQL</a></i>
</p>

## Reasons for SQL

* Structured data
* Strict schema
* Relational data
* Need for complex joins
* Transactions
* Clear patterns for scaling
* More established: developers, community, code, tools, etc
* Lookups by index are very fast

## Reasons for NoSQL

* Semi-structured data
* Dynamic or flexible schema
* Non-relational data
* No need for complex joins
* Store many TB (or PB) of data
* Very data intensive workload
* Very high throughput for IOPS

## Sample data well-suited for NoSQL

* Rapid ingest of clickstream and log data
* Leaderboard or scoring data
* Temporary data, such as a shopping cart
* Frequently accessed ('hot') tables
* Metadata/lookup tables

> In an interview, the answer is rarely "all of one". A realistic design keeps transactional data in
> [SQL](rdbms.md) and moves the clickstream, session, and hot-lookup workloads to
> [NoSQL](nosql.md) — that's what the pastebin, Twitter, and Mint exercises all end up doing.

## Source(s) and further reading: SQL or NoSQL

* [Scaling up to your first 10 million users](https://www.youtube.com/watch?v=kKjm4ehYiMs)
* [SQL vs NoSQL differences](https://www.sitepoint.com/sql-vs-nosql-differences/)

## Self-check

1. Give four reasons to choose SQL and four to choose NoSQL.
2. Name three data shapes that are well suited to NoSQL, and say why for each.
3. Your system has transactional orders plus a high-volume clickstream. What do you propose?
