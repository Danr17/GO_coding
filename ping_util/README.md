# Web TCP Ping & ClI TCP Ping.

## Download your version
In the archive you should find:
* an executable (like tcp_ping-linux-386)
* index.html **Should be in the same folder with the executable, otherwise the program will crash**
* an example.txt file that can be modified for your needs

### Linux
tcp_ping-linux-386.tar.gz       
tcp_ping-linux-amd64.tar.gz  

### Mac
tcp_ping-darwin-386.tar.gz
tcp_ping-darwin-amd64.tar.gz

### Windows
tcp_ping-windows-386.tar.gz
tcp_ping-windows-amd64.tar.gz

## Run the tcp_ping script 

### WEB:
1. Web Ping with defauts (defaults: port=443, interval=1s, counter=4)
    > ./ping_util-linux-amd64 -p web -file example.txt
2. Web Ping with count 10, interval 3s, timeout 3s
    > ./tcp_ping-linux-amd64 -p web -file example.txt -c 10 -i 3s -t 3s

### CLI:
1. ping over TCP  with defaults (defaults: port=443, interval=1s, counter=4)
    > ping_util-linux-amd64 example.com
2. ping over TCP over with custom port, counter and interval
    > ping_util-linux-amd64 -p tcp -port 80 -c 3 -i 3s example.com

3. ping over UDP with defaults
	> ping_util-linux-amd64 -p udp example.com

4. ping over ICMP in Privilege (run as super-user) mode !!!
	> sudo ./ping_util -p icmp -privilege example.com
5. ping over ICMP without Privilege mode, is actually over UDP
	> sudo ./ping_util -p icmp -privilege example.com

## Example

### Web

```bash
./ping_util-linux-amd64 -p web -file example.txt
2020/06/22 07:53:34 API listening on port 8080. Open browser: http://localhost:8080
```

Open http://localhost:8080 in your browser. 

In Privileged mode:

```bash
$ sudo ./ping_util-linux-amd64 -p web -privileged -file example.txt
2020/06/26 14:53:59 API listening on port 8080. Open browser: http://localhost:8080

```

### CLI

#### TCP

```bash
$ ./ping_util-linux-amd64 example.com
Ping example.com:443(93.184.216.34:443) - Connected - time=270.931215ms
Ping example.com:443(93.184.216.34:443) - Connected - time=534.147984ms
Ping example.com:443(93.184.216.34:443) - Connected - time=353.278657ms
Ping example.com:443(93.184.216.34:443) - Connected - time=125.262357ms
2020/06/26 14:47:35 ping done for site example.com

Ping statistics example.com:443
	4 probes sent.
	4 successful, 0 failed.
Approximate trip times:
	Minimum = 125.262357ms, Maximum = 534.147984ms, Average = 320.905053ms
```

Ping over TCP over with custom port, counter and interval

```bash
$ ./ping_util-linux-amd64 -p tcp -port 80 -c 3 -i 3s example.com
Ping example.com:80(93.184.216.34:80) - Connected - time=428.117245ms
Ping example.com:80(93.184.216.34:80) - Connected - time=295.362951ms
Ping example.com:80(93.184.216.34:80) - Connected - time=162.165586ms
2020/06/26 14:48:34 ping done for site example.com

Ping statistics example.com:80
	3 probes sent.
	3 successful, 0 failed.
Approximate trip times:
	Minimum = 162.165586ms, Maximum = 428.117245ms, Average = 295.21526ms
```

#### UDP

```bash
$ ./ping_util -p udp example.com
Ping example.com:443(93.184.216.34:443) - Connected - time=3.046954ms
Ping example.com:443(93.184.216.34:443) - Connected - time=1.396042ms
Ping example.com:443(93.184.216.34:443) - Connected - time=1.938326ms
Ping example.com:443(93.184.216.34:443) - Connected - time=1.638996ms
2020/06/26 14:50:19 ping done for site example.com

Ping statistics example.com:443
	4 probes sent.
	4 successful, 0 failed.
Approximate trip times:
	Minimum = 1.396042ms, Maximum = 3.046954ms, Average = 2.005079ms
```

#### ICMP

Not in Provileged mode, actually using UDP:
```bash
$ sudo ./ping_util -p icmp example.com
PING example.com (93.184.216.34):
24 bytes from 93.184.216.34: icmp_seq=0 time=464.020671ms ttl=54
24 bytes from 93.184.216.34: icmp_seq=1 time=282.503073ms ttl=54
24 bytes from 93.184.216.34: icmp_seq=2 time=318.137142ms ttl=54
24 bytes from 93.184.216.34: icmp_seq=3 time=330.844197ms ttl=54

--- example.com ping statistics ---
4 packets transmitted, 4 packets received, 0% packet loss
round-trip min/avg/max/stddev = 282.503073ms/348.87627ms/464.020671ms/68.799861ms
```

**Privileged mode** (to permit creation of [ICMP Ras Sockets](https://sturmflut.github.io/linux/ubuntu/2015/01/17/unprivileged-icmp-sockets-on-linux/))
```bash
$ sudo ./ping_util -p icmp -privileged example.com
PING example.com (93.184.216.34):
24 bytes from 93.184.216.34: icmp_seq=0 time=386.175003ms ttl=54
24 bytes from 93.184.216.34: icmp_seq=1 time=230.539546ms ttl=54
24 bytes from 93.184.216.34: icmp_seq=2 time=459.186792ms ttl=54
24 bytes from 93.184.216.34: icmp_seq=3 time=278.273387ms ttl=54

--- example.com ping statistics ---
4 packets transmitted, 4 packets received, 0% packet loss
round-trip min/avg/max/stddev = 230.539546ms/338.543682ms/459.186792ms/89.611529ms
```
