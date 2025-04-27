# climate-check

## Architecture
![Architecture](docs/architecture.png)

## Overview
This repository consists of several simple projects that form client-server system.
These are:
- arduino board with DHT sensor (collects temperature & humidity data)
- serial "logger" (reads data from board & sends it to server)
- server (GET/POST, interacts with database)
- client (CLI, outputs histograms based on data from server)


## Configuration
Create .env files in project directories roots to configure them.
Client:
```
SERVER_IP=<server_ip>
SERVER_PORT=<server_port>
```

Serial-logger:
```
SERVER_IP=<server_ip>
SERVER_PORT=<server_port>
```

Server:
```
SERVER_IP=<server_ip>
SERVER_PORT=<server_port>
DB_PORT=<port>
DB_NAME=<name>
DB_URL=postgres://<username>:<password>@${SERVER_IP}:${DB_PORT}/${DB_NAME}?sslmode=disable
```
