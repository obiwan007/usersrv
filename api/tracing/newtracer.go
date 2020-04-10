package tracing

import (
	"github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin-contrib/zipkin-go-opentracing"
	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
)

const (
	endpoint_url              = "http://localhost:9411/api/v1/spans"
	host_url                  = "localhost:5051"
	service_name_cache_client = "cache service client"
	service_name_call_get     = "callGet"
)

func NewTracer() (opentracing.Tracer, zipkintracer.Collector, error) {
	collector, err := openzipkin.NewHTTPCollector(endpoint_url)
	if err != nil {
		return nil, nil, err
	}
	recorder := openzipkin.NewRecorder(collector, true, host_url, service_name_cache_client)
	tracer, err := openzipkin.NewTracer(
		recorder,
		openzipkin.ClientServerSameSpan(true))

	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer, collector, nil
}
