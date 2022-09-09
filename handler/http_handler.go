package handler

import (
	"fmt"
	"github.com/cloud-ace/geco"
	"log"
	"net/http"
)

type HttpHandler struct {
	ProjectId  string
	httpLogger geco.Geco
}

func NewHttpHandler(pid string) *HttpHandler {
	httpLogger, err := geco.Build(pid)
	if err != nil {
		log.Fatalf("geco init error (http): %s", err.Error())
	}
	return &HttpHandler{
		httpLogger: httpLogger,
		ProjectId:  pid,
	}
}

func (h HttpHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func (h HttpHandler) Debug(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.httpLogger.Debug(ctx, "Debug")
	fmt.Fprint(w, "Debug")
}

func (h HttpHandler) Info(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.httpLogger.Info(ctx, "Info")
	fmt.Fprint(w, "Info")
}

func (h HttpHandler) Error(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.httpLogger.Error(ctx, "Error")
	fmt.Fprint(w, "Error")
}

func (h HttpHandler) Critical(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.httpLogger.Critical(ctx, "Critical")
	fmt.Fprint(w, "Critical")
}

func (h HttpHandler) Warning(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.httpLogger.Warning(ctx, "Warning")
	fmt.Fprint(w, "Warning")
}
