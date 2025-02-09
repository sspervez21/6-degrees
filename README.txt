### This server is meant to help you figure out how many degrees of separation an actor is from Kevin Bacon.
### The API call lets you specify an actor name and returns how many degrees of sepatation exist between that actor and Kevin Bacon.

### generate swagger files
swagger generate server -f swagger.yaml

### download dependences
dep init
dep ensure

### build the server
go build cmd/nr6degrees-server/main.go

### start the server
./main.exe --port=8086

### sample usage
curl -X GET http://localhost:8086/calculate_degrees/Sean+Penn
curl -X GET http://localhost:8086/calculate_degrees/Dennis+Quaid

### Testing limitations

### Due to time limitation the server has only been tested for actors up to 2 degrees of sepatation from Kevin Bacon.
### If you specify an actor with more than 2 degrees of separation from Kevin Bacon (I was unable to find one) the server may take an unspecified amount of time returning a response.

### Planned improvements

### The performance of this server could be improved greatly by adding local caching of results (movie casts) as well as bulk processing of requests. These have not been implemented due to time limitation.
