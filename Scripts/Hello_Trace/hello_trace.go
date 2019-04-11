package hello_trace

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

var exporter *stackdriver.Exporter

func init() {
	var err error
	exporter, err = stackdriver.NewExporter(stackdriver.Options{})
	if err != nil {
		log.Println(err)
	}
	trace.RegisterExporter(exporter)
	//trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

//Hello respond with "Hello World."
func Hello(w http.ResponseWriter, r *http.Request) {
	ctx, sp := trace.StartSpan(r.Context(), "example.com/hello_trace.Hello")
	defer sp.End()

	switch rand.Intn(3) {
	case 0:
		sp.AddAttributes(trace.BoolAttribute("slow", true))
		slow(ctx)
	case 1:
		sp.AddAttributes(trace.BoolAttribute("fib", true))
		fibonacci(31)
	}
	fmt.Fprintln(w, "Hello World.")
}

func slow(ctx context.Context) {
	_, sp := trace.StartSpan(ctx, "example.com/hello_trace.slow")
	defer sp.End()
	time.Sleep(300 * time.Millisecond)
}

func fibonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
