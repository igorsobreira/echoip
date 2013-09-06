# echoip

Small TCP server that echoes the client IP

Why? I needed a way to test our NAT configuration. We have a setup where outgoing
connections can bind to a local ip:port to choose what is going to be the external IP.

## Usage

Start the server somewhere outside your LAN

    $ echoip -b 0.0.0.0:8001

within you LAN connect to it specifying a local ip and port

    $ nc -s 192.168.100.17 -p 9001 igorsobreira.com 8081

the response will be the external IP your connection used, so you can verify if
your NAT configuration is correct.

## Installing

Just download one the binaries from `bin/`
