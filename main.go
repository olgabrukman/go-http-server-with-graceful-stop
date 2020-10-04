package main

import (
	//"net/http"
	"net/http"
	"sync"

	"go-http-server-with-graceful-stop/src/httpserverwithgracefulstop"
)

func main() {
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)

	port := 8081
	srv := httpserverwithgracefulstop.StartMockHTTPServer(port, httpServerExitDone, "{}", http.StatusOK)

	httpserverwithgracefulstop.StopMockHTTPServerGracefully(srv, httpServerExitDone)
}
