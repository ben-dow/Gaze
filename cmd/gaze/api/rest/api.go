package rest

import (
	"github.com/ben-dow/Gaze/cmd/gaze/api/socket"
	"net/http"
)

func NewRestApi() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", Health)
	mux.Handle("/", http.FileServer(http.Dir("/opt/web")))
	mux.HandleFunc("/ws", socket.WebSocketHttp)

	return mux
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
