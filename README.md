# Docker Chaos Server  
A Chaos Server implementation that terminates Docker containers, used by [April](https://github.com/barbosaigor/april) tool.

## What is a Chaos Server ?
Chaos server hosts an API which terminates instances. It is used by [April](https://github.com/barbosaigor/april) CLI, 
which runs its algorithm and asks the Chaos Server to finish any selected instances. 
All Chaos Servers implementations must implement the interface defined in april/destroyer package, so CSs must include that package and
implement the Destroyer interface, where the business logic to terminate instances should be defined.  

## Installation  
```bash 
go get -u github.com/barbosaigor/dockercs/...
```  

DockerCS hosts an HTTP API    
-u username for chaos server auth  
-s password for chaos server auth  
-p port number (Default is 7071)  
```bash 
dockercs -u myusername -s mysecret  
``` 

_DockerCS requests dockerd locally, make sure that DockerCS has access to dockerd locally._  
