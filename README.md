# A simple REST service to check if a message is palindrome or not

A microserver exposes an API api/v1/message 


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

GET `/api/v1/messages` Returns all the messages

GET `/api/v1/messages/{id}` Returns a message with ID and also tells if the message is a palindrome or not 
```js
{"messageText": "Amore, roma", "isPalindrome":"true"}
```
POST `/api/v1/messages` Adds a new message to the list of messages to be requested later
