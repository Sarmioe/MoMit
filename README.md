# MoMit encrypted protocol

[中文简体](https://github.com/Sarmioe/MoMit/blob/main/Mo%20Mit%E5%B7%A5%E4%BD%9C%E5%8E%9F%E7%90%86.md)

## Introducing

> This is an encrypted protocol.
>
> MoMit is different from any encryption protocol such as shandowsocks, obfs4, snowflake, etc., which use a single obfuscation mode.
>
> It uses multiple obfuscation modes, such as obfuscating into a video call in one minute and browsing a video in the next minute, etc.
>
> Made by Golang , Why? Because it have high performance and easy learn.

## How to run build?

> It support using makefile(GNU Makefile) and go build to run build.
>
> First , you need cd into the project root folder and golang V1.23.4
>
> Makefile build:
>
> Client: make client
>
> Server: make server
>
> Clean: make debuild
>
> Go Build:
>
> Client: cd ./MoMitClient && go build ./main.go ./utils.go
>
> Server: cd ./MoMitServer && go build ./main.go ./utils.go
>
> Run the "build-all.bat" or "build-all.sh" file to build the all versions.

## How to it work?

### Transport protocol

> The code indicates which protocol is used for transmission based on IV1

| Code | Protocol             |
| ---- | -------------------- |
| 1    | TCP RAW              |
| 2    | UDP defaults to QUIC |
| 3    | TLS                  |
| 4    | HTTPS                |
| 5    | DoT                  |
| 6    | DoH                  |
| 7    | mKCP                 |
| 8    | gRCP                 |

### Camouflage protocol

> What is the code IV2? It is the camouflage feature
>
> And it will add a certain random value. It will be agreed on where to place it before the data packet starts to be transmitted (multiple random values may be inserted in multiple places in a data packet)
>
> What makes the censors even more angry is that every time there is a random value (an integer with a minimum of 1 and a maximum of 20), this code will randomly generate a new one, so that the proxy server is also disguised as a reverse proxy, and the client pretends to use this reverse proxy to surf the Internet normally
>
> There is also TTL value camouflage, which makes the proxy server look more like a reverse proxy server
>
> There is also automatic rotation of the proxy server IP. The proxy server IP is changed every few minutes, which makes DPI angry. It also supports IPv6, which makes DPI even more angry
>
> What makes DPI even more angry is that this thing will also send several real and fake data packets. It will really create a request to access the website, but it will be deleted after reaching the proxy server
>
> If a certain IP replays the data packet to the proxy server, it will really return a legitimate page (or a legitimate website DNS query result) to this IP, and in order to avoid being exposed, this IP will also be recorded Each replay will return the same data

| Code | Disguise features                | Disguise behavior                                            |
| ---- | -------------------------------- | ------------------------------------------------------------ |
| 9    | Watch video websites             | Simulate CDN behavior Divide traffic into multiple small segments Randomly distribute on different IPs |
| 10   | Listen to music websites         | Simulate buffering and pausing Intermittently transmit traffic instead of continuous transmission |
| 11   | Download large files             | Simulate breakpoint resumption Re-request different parts of the file after a period of time |
| 12   | Log in to the cloudflare website | Small data packets Slightly longer interval Simulate logging in to a website |
| 13   | Play online games                | High-frequency small data packets Simulate UDP traffic Combined with Ping value randomization (but there is also a limit Maximum 500ms Minimum 100ms and slightly smaller fluctuations) |
| 14   | Video call                       | Simulate video conferencing protocols such as WebRTC or Zoom |
| 15   | Random data                      | Pre-disguise with OBFS4+Shandowsocks Then insert some useless data |

### Packet acceleration

> Mo Mit uses a variety of advanced packet size reduction modes to ensure that the packet size is smaller and more data is transmitted
>
> 1. Preamble deletion
>
> 2. Frame gap deletion
>
> 3. Try frame fragmentation as much as possible
>
> 4. If the protocol is based on TCP, BBR will be enabled
>
> 5. If the protocol is based on UDP, UOT will be enabled
>
> 6. Enable header compression
>
> 7. Enable server-side (to proxy server) Gzip compression and Brotli compression
>
> 8. If https is randomly reached, HTTP/3 will be enabled
>
> 9. Merge similar data
>
> 10. IPv4/v6 automatically selects the best
>
> 11. Enable video or image compression
>
> And more

### Simplified actual Internet operation process

> First, the client needs to get the server key and a server trust list
>
> The trust list can have a maximum of 10 entries at a time and a minimum of 2 entries. It can be a single IPv4\IPv6 address. IPv4 server is recommended for deployment. After all, compared to IPV6, this thing is more stable in transmission.

> 1. Communicate with the proxy server and establish a WS connection with it
>
> 2. Then exchange encryption keys. The key will be signed locally. Then confirm that the server is the one you want to connect to. If there is tampering by the middleman, immediately cut off the network (kill switch)
>
> 3. Then start the second data packet exchange: first generate a random value locally, an IV1, and then notify the server
>
> 4. After the server gets this value, it sends a received data packet to the client
>
> 5. After the client receives the data packet, cut off the WS connection
>
> 6. Use the new protocol to start transmitting data
>
> 7. After the new protocol is negotiated, start negotiating the random data position (IV3). Add a few random packets for the number Integer with a maximum of 10 and a minimum of 1
>
> 8. The client starts random data and uses the encryption key negotiated during the TCP protocol as the transmission encryption of the random data position notification data packet
>
> 9. After the server confirms that everything is fine, the client starts random IV2
>
> 10. After IV2 is randomized, encrypt the data to the server and then notify the server of the IV2 number
>
> 11. Get browser data and start using Gzip and Brotli and video compression to compress data
>
> 12. Start deleting the preamble and interframe gap
>
> 13. After deleting, start merging and header compression data
>
> 14. Complete other data packet optimization work
>
> 15. Transmit a small data packet to notify the server of the transmission mode for these few minutes
>
> 16. Locally connect to a new IP address
>
> 16.5. Announce the encryption method used to the next server abroad (encrypted notification)
>
> 17. End the connection to this server and open a new connection to start connecting to this new address
>
> 18. When the random time expires, the connection with this server will be cut off and the next server will be opened. This cycle will continue.

> Note that the first few data packets are encrypted. All data of the first server are transmitted in an encrypted environment.
>
> In addition, IV1 in the notification process is disguised as a DNS request (for example, if the random number reaches 0, the DNS result returns 0.0.0.0). IV3 is disguised as an encrypted QUIC request (for example, watching YouTube). IV2 is disguised as IMTP to read emails. The small data packet in step 14 is disguised as an encrypted NTP request and the value is wrapped in the time. It looks like a wrong time on the outside, but...
>
> Moreover, these requests can also be protected if they face replay attacks.
>
> Before switching to the next data server, the client will first send a request to this server. The content is roughly: Have you finished transmitting (WS Request)
>
> If the server says: I have finished transmitting, the client will treat it as the first server and continue to exchange keys.
>
> Then after this server has exchanged keys, the client will connect to the next server in the trust list.
