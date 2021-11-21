package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func (rh *RequestHandler) serveHTTP(ctx context.Context) {
	r := mux.NewRouter()

	r.Path("/price").
		Methods(http.MethodGet).
		Queries("fsyms", "{fsyms}").
		Queries("tsyms", "{tsyms}").
		HandlerFunc(rh.priceHandler)

	server := &http.Server{Addr: ":8080", Handler: r}

	var err error
	go func() {
		if err = server.ListenAndServe(); err != nil {
			rh.l.Error("ListenAndServe error", zap.Error(err))
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		rh.l.Error("server.Shutdown error", zap.Error(err))
	}
}
