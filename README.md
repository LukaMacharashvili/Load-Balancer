# Warning
### This is a toy project for educational purposes only. It is not intended for production use.

# How to run locally

## Install dependencies for the example server

```bash
cd example-server
npm install
cd ..
```

## Install dependencies for the LB

```bash
cd LB
go tidy
cd ..
```

## Run the example server twice

```bash
cd example-server

export PORT=3000
node index.js

export PORT=3001
node index.js
```

## Run the LB

```bash
cd LB
export TARGET_URLS=http://localhost:3000,http://localhost:3001
cd cmd/lb
go run .
```

## Test the LB

Send a request to the LB (http://localhost:3002)
You can check the logs of the example servers to see the requests being routed to the different servers

# Docker

## Pull the image

```bash
docker pull lukamacharashvili/my-load-balancer:latest
```

## Run the image

```bash
docker run --network="host" -e TARGET_URLS=http://localhost:3000,http://localhost:3001 lukamacharashvili/my-load-balancer:latest
```

## Test the LB

Send a request to the LB (http://localhost:3002)
You can check the logs of the example servers to see the requests being routed to the different servers