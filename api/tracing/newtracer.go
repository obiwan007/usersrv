package tracing

import (
	"github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin-contrib/zipkin-go-opentracing"
	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
)

// Zipkin stuff
const (
	host_url              = "0.0.0.0:0"
	service_name_call_get = "callGet"
)

func NewTracer(servicename string, url string) (opentracing.Tracer, zipkintracer.Collector, error) {
	collector, err := openzipkin.NewHTTPCollector(url)
	if err != nil {
		return nil, nil, err
	}
	recorder := openzipkin.NewRecorder(collector, true, host_url, servicename)
	tracer, err := openzipkin.NewTracer(
		recorder,
		openzipkin.ClientServerSameSpan(true))

	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer, collector, nil
}
