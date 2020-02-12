# April Chaos Server (Docker)  
An chaos server implementation that terminates docker containers instances by API.

## What is a Chaos server ?
Chaos server hosts an API that terminantes instances. Aprils CLI runs its algorithm and asks the Chaos server to finish 
the selected instances. The API implementation lives in april/destroyer package, so chaos servers must include that package and
implement the Destroyer interface, which contain the business logic for terminate instances. 

