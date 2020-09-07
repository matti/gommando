# gommando

Run any command and run hooks when the output (stdout, stderr or stdboth) matches

See [examples/ping/main.go](examples/ping/main.go)

which outputs:
```
$ go run examples/ping/main.go
PING google.com (172.217.20.46): 56 data bytes
64 bytes from 172.217.20.46: icmp_seq=0 ttl=116 time=20.072 ms

found 'google', I'll never say this again because I'm after .Once, but I'll trigger the next...
... .Then and I'll trigger the next .Once which will start to match now:
found 'icmp_seq=' seen and I continue to say this because I'm after .Every
64 bytes from 172.217.20.46: icmp_seq=1 ttl=116 time=26.538 ms
found 'icmp_seq=' seen and I continue to say this because I'm after .Every

found 'ttl' and I'll never say this again because I'm after .Once
64 bytes from 172.217.20.46: icmp_seq=2 ttl=116 time=21.320 ms
found 'icmp_seq=' seen and I continue to say this because I'm after .Every
64 bytes from 172.217.20.46: icmp_seq=3 ttl=116 time=25.697 ms
found 'icmp_seq=' seen and I continue to say this because I'm after .Every
64 bytes from 172.217.20.46: icmp_seq=4 ttl=116 time=18.994 ms

--- google.com ping statistics ---
5 packets transmitted, 5 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 18.994/22.524/26.538/3.037 ms
```

## More examples

See [examples/](examples/)
