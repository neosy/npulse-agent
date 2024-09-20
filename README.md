# npulse-agent
Agent for monitoring the operation of servers and workstations

## Installation
### Clone and install
```
cd /tmp && git clone https://github.com/neosy/npulse-agent.git && cd npulse-agent && make install
```
## Launching
```
npulse-agent -a <url1,url2,url...> -p <port>
```
### Example
```
npulse-agent -a "http://127.0.0.1" -p 8080
```