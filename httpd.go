package httpd

import (
	"net/http"

	"github.com/moisespsena-go/http-render/rrhandler"

	"github.com/moisespsena-go/httpu"
)

type handler http.Handler

type Server struct {
	Handler       http.Handler
	RenderHandler *rrhandler.RequestRenderHandler
}

func New(srv ...*Server) (s *Server) {
	for _, s = range srv {
	}
	if s == nil {
		s = &Server{}
	}
	s.init()
	return
}

func (this *Server) init() {
	if this.RenderHandler == nil {
		this.RenderHandler = &rrhandler.RequestRenderHandler{}
	}
	if this.Handler == nil {
		this.Handler = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	}
}

func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpu.Fallback(
		this.Handler,
		this.RenderHandler,
	).ServeHTTP(w, r)
}