package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (svr *Server) Run(port string, handler http.Handler) error {
	svr.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return svr.httpServer.ListenAndServe()
}

func (svr *Server) Shutdown(ctx context.Context) error {
	return svr.httpServer.Shutdown(ctx)
}
