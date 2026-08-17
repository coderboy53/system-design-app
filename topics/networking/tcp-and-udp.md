---
id: tcp-and-udp
title: TCP and UDP
module: networking
order: 60
est_minutes: 6
timelines: [short, medium, long]
tags: [networking, protocols, tcp, udp, transport]
prerequisites: [http]
related: [http, rpc, load-balancer, consistency-patterns]
source: https://github.com/donnemartin/system-design-primer#transmission-control-protocol-tcp
---

# TCP and UDP

The transport layer under [HTTP](http.md), and the layer that
[layer 4 load balancing](load-balancer.md#layer-4-load-balancing) operates on.

## Transmission control protocol (TCP)

<p align="center">
  <img src="../assets/images/JdAsdvG.jpg" alt="TCP">
  <br/>
  <i><a href="http://www.wildbunny.co.uk/blog/2012/10/09/how-to-make-a-multi-player-game-part-1/">Source: How to make a multiplayer game</a></i>
</p>

TCP is a **connection-oriented** protocol over an [IP network](https://en.wikipedia.org/wiki/Internet_Protocol).
Connection is established and terminated using a [handshake](https://en.wikipedia.org/wiki/Handshaking).
All packets sent are **guaranteed to reach the destination in the original order and without
corruption** through:

* Sequence numbers and [checksum fields](https://en.wikipedia.org/wiki/Transmission_Control_Protocol#Checksum_computation) for each packet
* [Acknowledgement](https://en.wikipedia.org/wiki/Acknowledgement_(data_networks)) packets and automatic retransmission

If the sender does not receive a correct response, it will **resend** the packets. If there are
multiple timeouts, the connection is dropped. TCP also implements
[flow control](https://en.wikipedia.org/wiki/Flow_control_(data)) and
[congestion control](https://en.wikipedia.org/wiki/Network_congestion#Congestion_control). These
guarantees cause **delays** and generally result in **less efficient transmission** than UDP.

To ensure high throughput, web servers can keep a large number of TCP connections open, resulting in
**high memory usage**. It can be expensive to have a large number of open connections between web
server threads and, say, a [memcached](https://memcached.org/) server.
[Connection pooling](https://en.wikipedia.org/wiki/Connection_pool) can help, in addition to
switching to UDP where applicable.

TCP is useful for applications that require **high reliability but are less time critical**. Some
examples include web servers, database info, SMTP, FTP, and SSH.

**Use TCP over UDP when:**

* You need all of the data to arrive intact
* You want to automatically make a best estimate use of the network throughput

## User datagram protocol (UDP)

<p align="center">
  <img src="../assets/images/yzDrJtA.jpg" alt="UDP">
  <br/>
  <i><a href="http://www.wildbunny.co.uk/blog/2012/10/09/how-to-make-a-multi-player-game-part-1/">Source: How to make a multiplayer game</a></i>
</p>

UDP is **connectionless**. Datagrams (analogous to packets) are guaranteed only at the datagram
level. Datagrams might reach their destination **out of order or not at all**. UDP does **not**
support congestion control. Without the guarantees that TCP supports, UDP is generally **more
efficient**.

UDP can **broadcast**, sending datagrams to all devices on the subnet. This is useful with
[DHCP](https://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol) because the client has not
yet received an IP address, thus preventing a way for TCP to stream without the IP address.

UDP is **less reliable** but works well in real time use cases such as **VoIP, video chat,
streaming, and realtime multiplayer games** — the same use cases as
[weak consistency](../fundamentals/consistency-patterns.md#weak-consistency).

**Use UDP over TCP when:**

* You need the lowest latency
* Late data is worse than loss of data
* You want to implement your own error correction

## Source(s) and further reading: TCP and UDP

* [Networking for game programming](https://gafferongames.com/post/udp_vs_tcp/)
* [Key differences between TCP and UDP protocols](http://www.cyberciti.biz/faq/key-differences-between-tcp-and-udp-protocols/)
* [Difference between TCP and UDP](http://stackoverflow.com/questions/5970383/difference-between-tcp-and-udp)
* [Transmission control protocol](https://en.wikipedia.org/wiki/Transmission_Control_Protocol)
* [User datagram protocol](https://en.wikipedia.org/wiki/User_Datagram_Protocol)
* [Scaling memcache at Facebook](http://www.cs.bu.edu/~jappavoo/jappavoo.github.com/451/papers/memcache-fb.pdf)

## Self-check

1. What two mechanisms give TCP its ordering and integrity guarantees?
2. Why does keeping many TCP connections open cost memory, and what mitigates it?
3. Give three use cases where UDP is the right choice and say why.
4. Why can UDP broadcast when TCP cannot, and where does that matter?
