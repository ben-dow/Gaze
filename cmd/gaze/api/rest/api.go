package rest

import (
	"github.com/ben-dow/Gaze/cmd/gaze/api/socket"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/logging"
	"net/http"
)

func NewRestApi() *http.ServeMux {
	mux := http.NewServeMux()

	healthHandler := http.HandlerFunc(Health)
	mux.Handle("/health", LoggingMiddleware(healthHandler))

	wsHandler := http.HandlerFunc(socket.WebSocketHttp)
	mux.Handle("/ws", LoggingMiddleware(wsHandler))

	fileServer := http.FileServer(http.Dir("/opt/web"))
	mux.Handle("/", LoggingMiddleware(fileServer))

	return mux
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		logging.Info("%s %s", r.Method, r.RequestURI)
	})
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
