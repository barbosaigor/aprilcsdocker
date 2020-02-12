# Docker Chaos Server  
An chaos server implementation that terminates docker containers instances by API.

## What is a Chaos Server ?
Chaos server hosts an API that terminantes instances. It is used in [Aprils](https://github.com/barbosaigor/april) CLI, which runs its algorithm and asks the Chaos server to finish 
the selected instances. The API implementation lives in april/destroyer package, so chaos servers must include that package and
implement the Destroyer interface, which contain the business logic for terminate instances.  

Dockercs hosts API    
-u username for chaos server auth  
-s password for chaos server auth  
-p port number (Default is 7071)  
```bash 
go run cmd/dockercs -u bob -s mysecret
``` 
