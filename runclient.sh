echo "Running client against kubernetes server service"
cd cli/gqlapi
go run main.go --server_addr 192.168.64.3:30000 --zipkin http://192.168.64.3:30011/api/v2/spans
