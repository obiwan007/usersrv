package tracing

import (
	"github.com/opentracing/opentracing-go"
	// zipkintracer "github.com/openzipkin-contrib/zipkin-go-opentracing"
	openzipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
)

// Zipkin stuff
const (
	service_name_call_get = "callGet"
)

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
	//opentracing.SetGlobalTracer(tracer)

	return tracer, collector, nil
}
