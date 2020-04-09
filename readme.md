Generating proto models
* sh protogen.sh


Testing the app directly

* gRPC Testing App 
brew cask install bloomrpc
https://github.com/uw-labs/bloomrpc

* fixing rm -r $GOPATH/src/go.etcd.io/etcd/vendor/golang.org/x/net/trace

https://golang.github.io/dep/docs/new-project.html

etcdctl put my-service/1.2.3.4 '{"Addr":"localhost:10000","Metadata":"..."}'

tracing:
docker run -d -p 9411:9411 openzipkin/zipkin

openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"

  run server

  go run server.go -tls -cert_file ../certs/cert.pem -key_file ../certs/key.pem

go run main.go -tls -ca_file ../certs/cert.pem 