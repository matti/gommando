# gommando

Run any command and run hooks when the output (stdout, stderr or stdboth) matches
Allows you to start scanning for strings like "error" *only* after something else has been seen.

See [examples/input/main.go](examples/input/main.go) which outputs:

```
$ go run examples/input/main.go
Enter your name:
  ^-- now it asks my name
      writing it (robot)
hello robot
  ^-- now it greeted me hello robot

  exited with 0
```

See [examples/ping/main.go](examples/ping/main.go) outputs:
```
$ go run examples/ping/main.go
PING google.com (172.217.20.46): 56 data bytes
  found 'google', I'll never say this again because I'm after .Once, but I'll trigger the next...
  ... .Then and I'll trigger the next .Once which will start to match now:
64 bytes from 172.217.20.46: seq=0 ttl=37 time=19.362 ms
  found 'ttl' and I'll never say this again because I'm after .Once
  found 'seq=' seen and I continue to say this because I'm after .Every
64 bytes from 172.217.20.46: seq=1 ttl=37 time=26.139 ms
  found 'seq=' seen and I continue to say this because I'm after .Every
64 bytes from 172.217.20.46: seq=2 ttl=37 time=19.728 ms
  found 'seq=' seen and I continue to say this because I'm after .Every
64 bytes from 172.217.20.46: seq=3 ttl=37 time=18.014 ms
  3 pings sent, exiting
  ping exited
  found 'seq=' seen and I continue to say this because I'm after .Every
```

## More examples

See [examples/](examples/)
