# “Word of Wisdom”  TCP client and server

## TCP server
### Configure
Following environment variables are used by server application:
- TCP_SERVER_PORT=8008 - tcp server port
- TCP_SERVER_HOST=localhost - tcp server port
- TCP_CONNECTION_DEADLINE=5s - (s - second | m - minute ) tcp connection deadline is an absolute time after which I/O operations fail instead of blocking.
- QUOTATION_FILE="quotations.txt" - text files with quotations(one per line)
- DIFFICULTY=5 - difficulty of pow algorithm
 
.env-file in root directory is used by docker-compose
### Run
```shell
docker-compose up tcp-server
```

## TCP client
### Configure
Following environment variables are used by server application:
- TCP_SERVER_PORT=8008 - tcp server port
- TCP_SERVER_HOST=localhost - tcp server port

### Run
```shell
docker-compose up tcp-client
```