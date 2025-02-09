### generate swagger files
swagger generate server -f swagger.yml

### download dependences
dep init
dep ensure

### build the server
go build cmd/nr6degrees-server/main.go

### start the server
./main.exe --port=8086

### sample usage
curl -X GET http://localhost:8086/calculate_degrees/Sean+Penn/10
curl -X GET http://localhost:8086/calculate_degrees/Dennis+Quaid/10
