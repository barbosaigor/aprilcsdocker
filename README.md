# Docker Chaos Server  
An chaos server implementation that terminates docker containers instances by API.  

## What is a Chaos Server ?
Chaos server hosts an API that terminantes instances. It is used in [April](https://github.com/barbosaigor/april) CLI, which run its algorithm and asks the Chaos server to finish 
the selected instances. The API implementation lives in april/destroyer package, so chaos servers must include that package and
implement the Destroyer interface, which contain the business logic for terminate instances.  

## Installation  
```bash 
go get -u github.com/barbosaigor/dockercs  
cd $GOPATH/src/github.com/barbosaigor/dockercs/cmd  
go install dockercs.go  
# Or  
go build -o $YOURPATH/dockercs dockercs.go  
```  

Dockercs hosts API    
-u username for chaos server auth  
-s password for chaos server auth  
-p port number (Default is 7071)  
```bash 
dockercs -u myusername -s mysecret  
``` 
