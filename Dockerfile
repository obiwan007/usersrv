# FROM golang:alpine as builder

# WORKDIR /go/src/app
# COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...

# EXPOSE 10000

# CMD ["usersrv"]


FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
# RUN go get -d -v ./...
RUN cd cli/usersrv && go build -v
FROM alpine
# RUN adduser -S -D -H -h /app appuser
# USER appuser
COPY --from=builder /build/cli/usersrv/usersrv /app/
WORKDIR /app
EXPOSE 10000
CMD ["./usersrv"]