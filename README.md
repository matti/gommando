# gommando

Run any command and run hooks when the output (stdout, stderr or stdboth) matches

See [examples/ping/main.go](examples/ping/main.go)

which outputs:
```
$ go run examples/ping/main.go
PING google.com (172.217.22.174): 56 data bytes
64 bytes from 172.217.22.174: icmp_seq=0 ttl=116 time=18.189 ms
found 'icmp_seq=' seen and I continue to say this because I'm after .Every

found 'google', I'll never say this again because I'm after .Once, but I'll trigger the next...
... .Then and I'll trigger the next .Once which will start to match now:
64 bytes from 172.217.22.174: icmp_seq=1 ttl=116 time=20.714 ms

found 'ttl' and I'll never say this again because I'm after .Once
found 'icmp_seq=' seen and I continue to say this because I'm after .Every
64 bytes from 172.217.22.174: icmp_seq=2 ttl=116 time=24.447 ms

--- google.com ping statistics ---
3 packets transmitted, 3 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 18.189/21.117/24.447/2.571 ms
found 'icmp_seq=' seen and I continue to say this because I'm after .Every
```

## More examples

See [examples/](examples/)
