package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

type Server struct {
	Mux          *mux.Router
	port string
}


func NewServer(port string) *Server {
	s := &Server{
		Mux:          mux.NewRouter(),
		port: 		port,
	}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

func NewMainHandler(m *mux.Router, port string, timeout time.Duration) {

	handler := &Handler{port: port, timeout:timeout}
	m.HandleFunc("/api/data", handler.HandleMain)
	m.HandleFunc("/", handler.HandleMain)
	m.Handle("/metrics", promhttp.Handler())
}


func (h *Handler) HandleMain(w http.ResponseWriter, r *http.Request) {
	time.Sleep(h.timeout * time.Millisecond)
	RequestCount.Inc()
	Respond(w, r, http.StatusOK, "hello from server" + h.port + " timeout: " + h.timeout.String())
}

type Handler struct{
	port string
	timeout time.Duration
}