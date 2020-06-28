package main

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/leodotcloud/log"
	"github.com/leodotcloud/log/server"
	"github.com/opentracing/opentracing-go"

	"github.com/namsral/flag"

	"github.com/coreos/etcd/clientv3"
	"github.com/rs/cors"

	"github.com/common-nighthawk/go-figure"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/trace"
	gql "github.com/obiwan007/usersrv/gqlsrv/api"
	"github.com/obiwan007/usersrv/pkg"
	pb "github.com/obiwan007/usersrv/proto"

	// zipkintracer "github.com/openzipkin-contrib/zipkin-go-opentracing"
	openzipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
)

var (
	config             = flag.String("config", "", "Using configfile")
	balancer           = flag.Bool("ectd", false, "Using etcd")
	tlsForGrpc         = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	gClientID          = flag.String("clientid", "xxxxxxxxxxxxxx", "OAuth2 client Id")
	gClientSecr        = flag.String("clientsecret", "yyyyyyyyyyyyy", "OAuth2 client secret")
	gRedirect          = flag.String("redirect", "http://localhost:3000/auth/callback", "OAuth2 client redirect callback")
	aClientID          = flag.String("aclientid", "xxxxxxxxxxxxxx", "OAuth2 client Id")
	aClientSecr        = flag.String("aclientsecret", "yyyyyyyyyyyyy", "OAuth2 client secret")
	aRedirect          = flag.String("aredirect", "http://localhost:3000/auth/callbackA", "OAuth2 client redirect callback")
	mClientID          = flag.String("mclientid", "xxxxxxxxxxxxxx", "OAuth2 client Id")
	mClientSecr        = flag.String("mclientsecret", "yyyyyyyyyyyyy", "OAuth2 client secret")
	mTenant            = flag.String("mtenant", "common", "OAuth2 Azure Tenant Id")
	mRedirect          = flag.String("mredirect", "http://localhost:3000/auth/callbackA", "OAuth2 client redirect callback")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "usersrv:10000", "The server address in the format of host:port")
	userSrv            = flag.String("usersrv", "usersrv:10000", "The server address in the format of host:port")
	timerSrv           = flag.String("timersrv", "timersrv:10001", "The server address in the format of host:port")
	projectSrv         = flag.String("projectsrv", "projectsrv:10002", "The server address in the format of host:port")
	clientSrv          = flag.String("clientsrv", "clientsrv:10003", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	zipkin             = flag.String("zipkin", "http://zipkin:9411/api/v1/spans", "Zipkin URL")
	signingKey         = flag.String("signingkey", "captainjacksparrowsayshi", "JWT Key")
)

// const (
// 	endpoint_url              = "http://localhost:9411/api/v1/spans"
// 	host_url                  = "localhost:5051"
// 	service_name_cache_client = "cache service client"
// 	service_name_call_get     = "callGet"
// )

func NewTracer(servicename string, grpcurl, zipkinurl string) (opentracing.Tracer, openzipkin.Collector, error) {
	collector, err := openzipkin.NewHTTPCollector(zipkinurl)
	if err != nil {
		return nil, nil, err
	}
	recorder := openzipkin.NewRecorder(collector, true, grpcurl, servicename)
	tracer, err := openzipkin.NewTracer(
		recorder, openzipkin.ClientServerSameSpan(true))

	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer, collector, nil
}

func main() {
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	server.StartServerWithDefaults()

	myFigure := figure.NewFigure("GQLSRV", "", true)
	myFigure.Print()
	flag.Parse()

	gql.SigningKey = []byte(*signingKey)

	gqlTracer, _, err := NewTracer("gql service", "0.0.0.0:0", *zipkin)

	params := &pkg.CommandParams{TlsForGrpc: tlsForGrpc, Balancer: balancer, CertFile: caFile, Zipkin: zipkin, Tracename: "service"}

	s, err := getSchema("../schema/schema.graphql")
	if err != nil {
		panic(err)
	}

	// for GraphQL/http endpoint
	tlsCfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("fail to connect to etcd: %v", err)
	}
	defer cli.Close()
	p1 := *params
	p1.Tracename = "userservice"
	connUserSrv, tracer, err := pkg.CreateClient(&p1, *userSrv)
	defer connUserSrv.Close()
	gqlUserSrvClient := pb.NewUserServiceClient(connUserSrv)

	p2 := *params
	p2.Tracename = "timerservice"
	connTimerSrv, tracer, err := pkg.CreateClient(&p2, *timerSrv)
	defer connTimerSrv.Close()
	gqlTimerSrvClient := pb.NewTimerServiceClient(connTimerSrv)

	p3 := *params
	p3.Tracename = "projectservice"
	connProjectSrv, tracer, err := pkg.CreateClient(&p3, *projectSrv)
	defer connProjectSrv.Close()
	gqlProjectSrvClient := pb.NewProjectServiceClient(connProjectSrv)

	p4 := *params
	p4.Tracename = "clientservice"
	connClientSrv, tracer, err := pkg.CreateClient(&p4, *clientSrv)
	defer connClientSrv.Close()
	gqlClientSrvClient := pb.NewClientServiceClient(connClientSrv)

	log.Infof("Tracer: %v", tracer)

	gql.UserSrvClient = gqlUserSrvClient

	resolver := gql.NewResolver(
		gqlUserSrvClient,
		gqlTimerSrvClient,
		gqlProjectSrvClient,
		gqlClientSrvClient)

	// schema := graphql.MustParseSchema(s, resolver, graphql.UseStringDescriptions(), graphql.Tracer(trace.OpenTracingTracer{}))
	schema := graphql.MustParseSchema(s, resolver, graphql.UseStringDescriptions(), graphql.Tracer(trace.OpenTracingTracer{}))
	mux := gql.NewRouter(schema, gqlTracer,
		&gql.AuthSecret{ClientID: gClientID, ClientSecr: gClientSecr, RedirectUrl: gRedirect},
		&gql.AuthSecret{ClientID: aClientID, ClientSecr: aClientSecr, RedirectUrl: aRedirect},
		&gql.AuthSecret{ClientID: mClientID, ClientSecr: mClientSecr, RedirectUrl: mRedirect, Tenant: mTenant})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	// mux2 := cors.Default().Handler(mux)
	mux2 := cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://foo.com", "http://foo.com:8080"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:3000", "https://frontend-obiwan007.cloud.okteto.net"},
		// Enable Debugging for testing, consider disabling in production
		AllowedHeaders: []string{"Authorization", "Content-Type", "X-B3-Sampled", "X-B3-Spanid", "X-B3-Traceid"},
		Debug:          false,
	}).Handler(mux)

	srv := &http.Server{
		Addr:         ":8090",
		Handler:      mux2,
		TLSConfig:    tlsCfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("Listening on port 8090")
	log.Fatalf("Exit", srv.ListenAndServe())
	// log.Fatal(srv.ListenAndServeTLS("server.rsa.crt", "server.rsa.key"))
}

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("username")
		if cookie != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
