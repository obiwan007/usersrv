const { Tracer, ExplicitContext } = require("zipkin");
const { recorder } = require("./recorder");
// const {
//   jsonEncoder: { JSON_V2 }
// } = require("zipkin");

const ctxImpl = new ExplicitContext();
const localServiceName = "browser";
const tracer = new Tracer({
  ctxImpl,
  recorder: recorder(localServiceName),
  localServiceName
});

// instrument fetch
const wrapFetch = require("zipkin-instrumentation-fetch");
const remoteServiceName = "gql service";
export const zipkinFetch = wrapFetch(fetch, { tracer, remoteServiceName });
