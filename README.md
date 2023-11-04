# MineMonitor
It is the fastest Minecraft Server Checker checking 2000 servers at the same time.
![image](https://github.com/Leaf48/MineMonitor/assets/58620209/92168c96-6135-4029-973c-9bef56565d86)

# Configuration
maxGoroutines is the maximum goroutine instance value.

# checklist.yaml
You can configure what IPs are subject to check. </br>
In those Ips, the numbers following second dot are subject to change. </br>
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
