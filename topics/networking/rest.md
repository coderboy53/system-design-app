---
id: rest
title: Representational State Transfer (REST)
module: networking
order: 80
est_minutes: 6
timelines: [short, medium, long]
tags: [networking, communication, rest, apis]
prerequisites: [http, rpc]
related: [rpc, http, load-balancer]
source: https://github.com/donnemartin/system-design-primer#representational-state-transfer-rest
---

# Representational State Transfer (REST)

REST is an **architectural style** enforcing a client/server model where the client acts on a set of
**resources** managed by the server. The server provides a representation of resources and actions
that can either manipulate or get a new representation of resources. All communication must be
**stateless and cacheable**.

## Four qualities of a RESTful interface

* **Identify resources (URI in HTTP)** — use the same URI regardless of any operation.
* **Change with representations (verbs in HTTP)** — use verbs, headers, and body.
* **Self-descriptive error message (status response in HTTP)** — use status codes, don't reinvent the wheel.
* **[HATEOAS](http://restcookbook.com/Basics/hateoas/) (HTML interface for HTTP)** — your web service should be fully accessible in a browser.

## Sample REST calls

```
GET /someresources/anId

PUT /someresources/anId
{"anotherdata": "another value"}
```

**REST is focused on exposing data.** It minimizes the coupling between client/server and is often
used for **public HTTP APIs**. REST uses a more generic and uniform method of exposing resources
through URIs, [representation through headers](https://github.com/for-GET/know-your-http-well/blob/master/headers.md),
and actions through [verbs](http.md#common-http-verbs) such as GET, POST, PUT, DELETE, and PATCH.

Being **stateless**, REST is great for [horizontal scaling](load-balancer.md#horizontal-scaling) and
partitioning.

## Disadvantage(s): REST

* With REST being focused on exposing data, it might **not be a good fit if resources are not
  naturally organized or accessed in a simple hierarchy**. For example, returning all updated
  records from the past hour matching a particular set of events is not easily expressed as a path.
  With REST, it is likely to be implemented with a combination of URI path, query parameters, and
  possibly the request body.
* REST typically relies on a **few verbs** which sometimes don't fit your use case. For example,
  moving expired documents to the archive folder might not cleanly fit within these verbs.
* Fetching complicated resources with nested hierarchies requires **multiple round trips** between
  the client and server to render single views — e.g. fetching a blog entry and its comments. For
  mobile applications operating in variable network conditions, these multiple round trips are
  highly undesirable.
* Over time, more fields might be added to an API response and **older clients will receive all new
  data fields**, even those they do not need, bloating payload size and leading to larger latencies.

## RPC and REST calls comparison

The same seven operations expressed both ways. [RPC](rpc.md) names the *operation*; REST names the
*resource* and lets the verb carry the operation.

| Operation | RPC | REST |
|---|---|---|
| Signup | **POST** /signup | **POST** /persons |
| Resign | **POST** /resign<br/>{<br/>"personid": "1234"<br/>} | **DELETE** /persons/1234 |
| Read a person | **GET** /readPerson?personid=1234 | **GET** /persons/1234 |
| Read a person's items list | **GET** /readUsersItemsList?personid=1234 | **GET** /persons/1234/items |
| Add an item to a person's items | **POST** /addItemToUsersItemsList<br/>{<br/>"personid": "1234";<br/>"itemid": "456"<br/>} | **POST** /persons/1234/items<br/>{<br/>"itemid": "456"<br/>} |
| Update an item | **POST** /modifyItem<br/>{<br/>"itemid": "456";<br/>"key": "value"<br/>} | **PUT** /items/456<br/>{<br/>"key": "value"<br/>} |
| Delete an item | **POST** /removeItem<br/>{<br/>"itemid": "456"<br/>} | **DELETE** /items/456 |

<p align="center">
  <i><a href="https://apihandyman.io/do-you-really-know-why-you-prefer-rest-over-rpc/">Source: Do you really know why you prefer REST over RPC</a></i>
</p>

## Source(s) and further reading: REST and RPC

* [Do you really know why you prefer REST over RPC](https://apihandyman.io/do-you-really-know-why-you-prefer-rest-over-rpc/)
* [When are RPC-ish approaches more appropriate than REST?](http://programmers.stackexchange.com/a/181186)
* [REST vs JSON-RPC](http://stackoverflow.com/questions/15056878/rest-vs-json-rpc)
* [Debunking the myths of RPC and REST](https://web.archive.org/web/20170608193645/http://etherealbits.com/2012/12/debunking-the-myths-of-rpc-rest/)
* [What are the drawbacks of using REST](https://www.quora.com/What-are-the-drawbacks-of-using-RESTful-APIs)
* [Crack the system design interview](http://www.puncsky.com/blog/2016-02-13-crack-the-system-design-interview)
* [Thrift](https://code.facebook.com/posts/1468950976659943/)
* [Why REST for internal use and not RPC](http://arstechnica.com/civis/viewtopic.php?t=1190508)

## Self-check

1. List the four qualities of a RESTful interface.
2. Why does statelessness make REST a good fit for horizontal scaling?
3. Give two cases where REST's resource model fits badly, and say what you would do instead.
4. Express "add an item to a person's list" as both an RPC call and a REST call.
