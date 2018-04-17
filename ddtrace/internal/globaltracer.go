package internal // import "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/internal"

import "gopkg.in/DataDog/dd-trace-go.v1/ddtrace"

// GlobalTracer holds the currently active tracer. It's "zero value" should
// always be the NoopTracer.
var GlobalTracer ddtrace.Tracer = &ddtrace.NoopTracer{}

// Testing is set to true when the mock tracer is active. It usually signifies that we are in a test
// environment. This value is used by tracer.Start to prevent overriding the GlobalTracer in tests.
var Testing = false
