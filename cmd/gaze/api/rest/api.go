package rest

import (
	"context"
	"github.com/ben-dow/Gaze/cmd/gaze/api/socket"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/logging"
	"net/http"
	"os"
	"path/filepath"
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

	spa := &sveleteHandler{
		staticPath: "/opt/web",
		indexPath:  "index.html",
	}
	mux.Handle("/", LoggingMiddleware(spa))

	healthHandler := http.HandlerFunc(Health)
	mux.Handle("/api/health", LoggingMiddleware(healthHandler))

	wsHandler := http.HandlerFunc(socket.WebSocketHttp)
	mux.Handle("/api/ws", LoggingMiddleware(wsHandler))

	logging.Trace("Registered Routes")

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &GazeRestApi{
		server: server,
	}
}

type sveleteHandler struct {
	staticPath string
	indexPath  string
}

func (h sveleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {

		// try appending html to file
		path = path + ".html"
		fi, err := os.Stat(path)
		if os.IsNotExist(err) || fi.IsDir() {
			// file does not exist or path is a directory, serve index.html
			http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
			return
		}
	}

	http.ServeFile(w, r, path)
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
