---
id: security-basics
title: Security Basics
module: security
order: 10
est_minutes: 3
timelines: [short, medium, long]
tags: [security, xss, sql-injection, least-privilege, encryption]
prerequisites: []
related: [rest, rdbms, reverse-proxy]
source: https://github.com/donnemartin/system-design-primer#security
---

# Security Basics

> **Note:** the primer marks this section as one that could use updates. Consider
> [contributing](https://github.com/donnemartin/system-design-primer#contributing) upstream.

Security is a broad topic. Unless you have considerable experience, a security background, or are
applying for a position that requires knowledge of security, you probably **won't need to know more
than the basics**:

* **Encrypt in transit and at rest.**
* **Sanitize all user inputs** or any input parameters exposed to the user, to prevent
  [XSS](https://en.wikipedia.org/wiki/Cross-site_scripting) and
  [SQL injection](https://en.wikipedia.org/wiki/SQL_injection).
* **Use parameterized queries** to prevent SQL injection.
* Use the principle of [**least privilege**](https://en.wikipedia.org/wiki/Principle_of_least_privilege).

## Where security shows up elsewhere in this path

* [Reverse proxies](../networking/reverse-proxy.md) hide backend servers, blacklist IPs, and limit connections per client.
* [Load balancers](../networking/load-balancer.md) and reverse proxies handle SSL termination, so backends don't each need [X.509 certificates](https://en.wikipedia.org/wiki/X.509).
* [DNS is a DDoS target](../networking/dns.md#disadvantages-dns).
* [Back pressure](../application/asynchronism.md#back-pressure) is an availability defense as much as a performance one.

## Source(s) and further reading

* [API security checklist](https://github.com/shieldfy/API-Security-Checklist)
* [Security guide for developers](https://github.com/FallibleInc/security-guide-for-developers)
* [OWASP top ten](https://www.owasp.org/index.php/OWASP_Top_Ten_Cheat_Sheet)

## Self-check

1. Name the four security basics every engineer should know.
2. What specifically prevents SQL injection, and why is sanitization alone not the answer?
3. State the principle of least privilege in one sentence.
