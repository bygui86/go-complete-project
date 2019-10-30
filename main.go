package main

import "github.com/bygui86/go-complete-project/api"

func main() {
	// TODO implement Prometheus metrics

	// TODO implement Kubernetes probes

	// TODO this is a blocking call, make it non-blocking
	api.Run()
	// TODO implement shutdown method and defer it

	// TODO implement method to listen for os interrupt signals
}
