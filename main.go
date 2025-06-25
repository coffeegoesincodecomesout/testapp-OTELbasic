package main

import (
   "fmt"
   "net/http"
   "context"
   "time"
   "log"

   "go.opentelemetry.io/otel"
   "go.opentelemetry.io/otel/codes"
   "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
   "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	ctx := context.Background()
	exp, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	tracerProvider := trace.NewTracerProvider(trace.WithBatcher(exp))
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	otel.SetTracerProvider(tracerProvider)
        
        http.HandleFunc("/ping" , func(w http.ResponseWriter, r *http.Request) {
		trace := otel.Tracer("http-server")
		_, span := trace.Start(r.Context(), "handleRequest")
		defer span.End()

		time.Sleep(1 * time.Second)
		span.SetStatus(codes.Ok, "Status 200")

                fmt.Fprintf(w,"pong")
	})

	log.Printf("Starting http-server...")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("failed to start http-server: %v", err)
	}
}
