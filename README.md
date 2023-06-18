# MineMonitor
It is the fastest Minecraft Server Checker, which 2000 server checks at the same time using Goroutine.
![image](https://github.com/Leaf48/MineMonitor/assets/58620209/92168c96-6135-4029-973c-9bef56565d86)

# Configuration
maxGoroutines is the maximum goroutine instance value.

# checklist.yaml
IPs can be set that is going to check. </br>
Those ips are used for checking where is changed after the second dot. </br>
```json
ips:
  - 160.251.0.0
  - 157.7.0.0
  - 118.27.0.0
  - 118.240.0.0
...
```

# Result (as txt file)
```txt
0: 160.251.xxx.xxx
	Version: 1.20
	Players: 0/20
	Player name: []
	MOTD: A Minecraft Server

1: 160.251.xxx.xxx
	Version: 1.17.1
	Players: 0/20
	Player name: []
	MOTD: A Minecraft Server

2: 160.251.xxx.xxx
	Version: 1.19.4
	Players: 0/20
	Player name: []
	MOTD: A Minecraft Server
...
```
