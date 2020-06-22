# Web TCP Ping & ClI TCP Ping.

## Download your version

### Linux
tcp_ping-linux-386        
tcp_ping-linux-amd64  

### Mac
tcp_ping-darwin-386
tcp_ping-darwin-amd64

### Windows
tcp_ping-windows-386
tcp_ping-windows-amd64

## Run the tcp_ping script 

### WEB:
1. Web Ping with defauts (defaults: port=443, interval=1s, counter=4)
    > ./tcp_ping-linux-amd64 -web -file=example.txt
2. Web Ping with over port=80, interval=3s, counter=10
    > ./tcp_ping-linux-amd64 -web -file=example.txt -port 80 -interval 3s -counter 10

### CLI:
1. ping over tcp  with defaults (defaults: port=443, interval=1s, counter=4)
    > ./tcp_ping-linux-amd64 example.com
2. ping over tcp over custom port
    > ./tcp_ping-linux-amd64 -port 22 example.com 
3. ping over tcp using counter and interval
    > ./tcp_ping-linux-amd64 -port 80 -counter 3 -interval 3s example.com

## Example

### Web

```bash
./tcp_ping -web -file=example.txt
2020/06/22 07:53:34 API listening on port 8080. Open browser: http://localhost:8080
```

Open http://localhost:8080 in your browser. 

### CLI

```bash
./tcp_ping --counter 10 example.com
Ping example.com:443(93.184.216.34:443) - Connected - time=194.842495ms
Ping example.com:443(93.184.216.34:443) - Connected - time=213.846808ms
Ping example.com:443(93.184.216.34:443) - Connected - time=243.057291ms
Ping example.com:443(93.184.216.34:443) - Connected - time=159.409986ms
Ping example.com:443(93.184.216.34:443) - Connected - time=183.742792ms
Ping example.com:443(93.184.216.34:443) - Connected - time=207.211287ms
Ping example.com:443(93.184.216.34:443) - Connected - time=129.699002ms
Ping example.com:443(93.184.216.34:443) - Connected - time=152.944282ms
Ping example.com:443(93.184.216.34:443) - Connected - time=177.170497ms
Ping example.com:443(93.184.216.34:443) - Connected - time=201.085108ms
2020/06/22 07:50:58 ping done for site example.com

Ping statistics example.com:443
	10 probes sent.
	10 successful, 0 failed.
Approximate trip times:
	Minimum = 129.699002ms, Maximum = 243.057291ms, Average = 186.300954ms
```
