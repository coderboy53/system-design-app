---
id: back-of-the-envelope
title: Back-of-the-Envelope Calculations
module: appendix
order: 10
est_minutes: 4
timelines: [short, medium, long]
tags: [reference, estimation, interview-prep, math]
prerequisites: []
related: [powers-of-two, latency-numbers, interview-approach]
source: https://github.com/donnemartin/system-design-primer#back-of-the-envelope-calculations
---

# Back-of-the-Envelope Calculations

You might be asked to do some estimates by hand — for example, how long it will take to generate
100 image thumbnails from disk, or how much memory a data structure will take.

**Clarify with your interviewer whether you should run back-of-the-envelope usage calculations**
before spending time on them.

## The method

This is [step 1 of the interview framework](../study-guide/interview-approach.md#step-1-outline-use-cases-constraints-and-assumptions)
made numeric. Work in this order:

1. **State the assumptions.** Number of users, reads per month, writes per month, read-to-write ratio.
2. **Size a single unit.** Bytes per row, per paste, per tweet — sum the individual fields.
3. **Multiply out to storage.** Per month, then extrapolate to 3 or 5 years.
4. **Convert to per-second rates.** This is where the conversion guide below does the work.
5. **Sanity-check against latency.** Use the [latency numbers](latency-numbers.md) to see whether
   the resulting rate is comfortable, needs a [cache](../caching/README.md), or needs
   [sharding](../data/sharding.md).

## Conversion guide

* **2.5 million seconds per month**
* 1 request per second = **2.5 million requests per month**
* 40 requests per second = **100 million requests per month**
* 400 requests per second = **1 billion requests per month**

Memorize the first line and the rest is division. (2.5 million ≈ 30 days × 86,400 s.)

## Worked example

From the [pastebin exercise](../exercises/system-design/pastebin/README.md):

* **Assumptions:** 10 million users, 10 million paste writes/month, 100 million paste reads/month, 10:1 read-to-write.
* **Size per paste:** 1 KB content + `shortlink` 7 bytes + `expiration_length_in_minutes` 4 bytes +
  `created_at` 5 bytes + `paste_path` 255 bytes ≈ **1.27 KB**.
* **Storage:** 1.27 KB × 10 million = **12.7 GB of new paste content per month** → **~450 GB in
  3 years**, ~360 million shortlinks in 3 years.
* **Rates:** 10 million writes/month ÷ 2.5 million s = **4 writes per second**; 100 million
  reads/month = **40 reads per second**.
* **Sanity check:** 40 reads/s of 1.27 KB is trivially cacheable; 4 writes/s fits a single
  [master-slave](../data/replication.md#master-slave-replication) pair. No sharding needed yet — and
  saying so out loud is the point of the exercise.

## Reference tables

* [Powers of two table](powers-of-two.md) — byte conversions
* [Latency numbers every programmer should know](latency-numbers.md) — what is fast

## Source(s) and further reading

* [Use back of the envelope calculations](http://highscalability.com/blog/2011/1/26/google-pro-tip-use-back-of-the-envelope-calculations-to-choo.html)

## Self-check

1. How many seconds are in a month, to the precision you'd use in an interview?
2. Convert 1 billion requests per month into requests per second.
3. A record is 500 bytes and you write 50 million per month. How much new storage per month, and per year?
4. After computing a rate, what is the next thing you check it against?
