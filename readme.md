# Generating proto models
`sh protogen.sh`


# Testing the gRPC interface

## gRPC Testing App 
`brew cask install bloomrpc`

Download via https://github.com/uw-labs/bloomrpc


# Service discovery with etcd
`etcdctl put my-service/1.2.3.4 '{"Addr":"localhost:10000","Metadata":"..."}'`
## fixing some strange dependency
`rm -r $GOPATH/src/go.etcd.io/etcd/vendor/golang.org/x/net/trace`

# Tracing with zipkin:
`docker run -d -p 9411:9411 openzipkin/zipkin`

# Cert generation
`openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"`

# Running the App
## run server

`go run server.go -tls -cert_file ../certs/cert.pem -key_file ../certs/key.pem`

## run client
`go run main.go -tls -ca_file ../certs/cert.pem`