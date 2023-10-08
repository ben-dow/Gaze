package rest

import (
	"context"
	"github.com/ben-dow/Gaze/cmd/gaze/api/socket"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/logging"
	"net/http"
	"sync"
)

type GazeRestApi struct {
	server *http.Server
}

func (a *GazeRestApi) Start(wg *sync.WaitGroup) {
	logging.Trace("Starting Rest Api")

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
			logging.Error("Rest Api Shutdown:  %v", err)
		}
	}()
}

func (a *GazeRestApi) Stop(ctx context.Context) error {
	logging.Trace("Stopping Rest Api")
	return a.server.Shutdown(ctx)
}

func NewRestApi(addr string) *GazeRestApi {
	mux := http.NewServeMux()

	healthHandler := http.HandlerFunc(Health)
	mux.Handle("/health", LoggingMiddleware(healthHandler))

	wsHandler := http.HandlerFunc(socket.WebSocketHttp)
	mux.Handle("/ws", LoggingMiddleware(wsHandler))

	fileServer := http.FileServer(http.Dir("/opt/web"))
	mux.Handle("/", LoggingMiddleware(fileServer))

	logging.Trace("Registered Routes with ServeMux")

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &GazeRestApi{
		server: server,
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	logging.Trace("creating a logging middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		logging.Info("%s %s", r.Method, r.RequestURI)
	})
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
