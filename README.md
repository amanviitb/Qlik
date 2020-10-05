# A simple REST service to check if a message is palindrome or not

The service exposes a simple REST API to check if a given message(a REST resource) is a palindrome or not. 


# Running Locally

Clone repository 
``` 
git clone https://github.com/amanviitb/Qlik
```

Go to the `src`  folder
```
cd Qlik/src
```

Execute the following command for building the service
```
make build
```
It will create an executable binary in the `/bin` folder


# Testing
Go to the `src` folder 

Each packages has its own unit tests written inside them.
Please run the following command to run all unit tests
```
make test
```
# Docker Image
A `Dockerfile` is present in the `src` folder which can be used to create the docker image of the project

### Create Docker Image
To build a docker image run the following command
```
docker build -t simple-service .
```

### Running the project as a Docker container
To run the docker image use the following command
```
docker run -p 9091:9091 simple-service
```
It will run the service on port 9091 and map it to port 9091 of the container

### API endpoints
GET `/api/v1/health` Returns the health and running state of the service
```js
 {
 "timeStampUTC":"2020-10-05 20:40:26.207180202 +0000 UTC",
 "serviceName":"palindrome-service",
 "serviceProvider":"Some Server",
 "serviceVersion":"v1.1.0",
 "serviceStatus":"Running",
 "connectionStatus":"Active","
 "hostname":"9c9d3c110cbe",
 "OS":"linux"
 }

```
GET `/api/v1/messages` Returns all the messages

GET `/api/v1/messages/{id}` Returns a message with ID and also tells if the message is a palindrome or not 
```js
{
"messageText": "Amore, roma",
"isPalindrome": "true"
}
```
POST `/api/v1/messages` Adds a new message to the list of messages to be requested later

DELETE `api/v1/messages{id}` Removes a message with ID or returns `404` if the message doesn't exist

### Observability
This is done through **_middlewares_** that are responsible for *logging* every incoming request and attaches some metrics to the request. Also, every request is tagged with a unique ID (unless otherwise sent as part of request header `X-Request-ID` for **_tracing_** the request should anything go wrong with that request.
