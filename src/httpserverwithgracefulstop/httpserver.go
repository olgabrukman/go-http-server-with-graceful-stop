package httpserverwithgracefulstop

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func StartMockHTTPServer(port int, wg *sync.WaitGroup, response string, responseStatus int) *http.Server {
	address := fmt.Sprintf(":%d", port)
	srv := &http.Server{Addr: address}

	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(responseStatus)
		w.Header().Set("Content-Type", "application/json")
		n, err := w.Write([]byte(response))
		if err != nil || n == 0 {
			log.Printf("Error writing to the http server: %v", err)
		}
	})

	go func() {
		defer wg.Done() // let main know we are done cleaning up

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			panic(err)
		}
	}()

	log.Println("started mock http server")

	return srv
}

func StopMockHTTPServerGracefully(srv *http.Server, wg *sync.WaitGroup) {
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	// wait for goroutine started in startHttpServer() to stop
	wg.Wait()

	http.DefaultServeMux = new(http.ServeMux)

	log.Println("stopped mock http server")
}
