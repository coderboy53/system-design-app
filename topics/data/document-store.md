---
id: document-store
title: Document Store
module: data
order: 90
est_minutes: 3
timelines: [short, medium, long]
tags: [database, nosql, document-store]
prerequisites: [key-value-store]
related: [key-value-store, wide-column-store, sql-or-nosql]
source: https://github.com/donnemartin/system-design-primer#document-store
---

# Document Store

> **Abstraction: key-value store with documents stored as values**

A document store is centered around **documents** (XML, JSON, binary, etc), where a document stores
**all information for a given object**. Document stores provide APIs or a query language to query
based on the **internal structure of the document itself**.

*Note: many [key-value stores](key-value-store.md) include features for working with a value's
metadata, blurring the lines between these two storage types.*

Based on the underlying implementation, documents are organized by **collections, tags, metadata, or
directories**. Although documents can be organized or grouped together, **documents may have fields
that are completely different from each other**.

Some document stores like [MongoDB](https://www.mongodb.com/mongodb-architecture) and
[CouchDB](https://blog.couchdb.org/2016/08/01/couchdb-2-0-architecture/) also provide a SQL-like
language to perform complex queries.
[DynamoDB](http://www.read.seas.harvard.edu/~kohler/class/cs239-w08/decandia07dynamo.pdf) supports
both key-values and documents.

Document stores provide **high flexibility** and are often used for working with **occasionally
changing data**.

## Source(s) and further reading: document store

* [Document-oriented database](https://en.wikipedia.org/wiki/Document-oriented_database)
* [MongoDB architecture](https://www.mongodb.com/mongodb-architecture)
* [CouchDB architecture](https://blog.couchdb.org/2016/08/01/couchdb-2-0-architecture/)
* [Elasticsearch architecture](https://www.elastic.co/blog/found-elasticsearch-from-the-bottom-up)

## Self-check

1. What does a document store let you query that a plain key-value store does not?
2. What does it mean that documents in a collection may have completely different fields?
3. What workload shape are document stores good for?
4. Name two document stores that also offer a SQL-like query language.
