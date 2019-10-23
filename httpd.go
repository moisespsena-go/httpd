package httpd

import (
	"net/http"

	http_render "github.com/moisespsena-go/http-render"

	"github.com/moisespsena-go/http-render/ropt"
)

type Server struct {
	Handler                  http.Handler
	RootDir                  string
	RenderOrNotFoundDisabled bool
	Render                   http_render.Render
	RenderFunc               func(w http.ResponseWriter, r *http.Request) http_render.Render
}

func New(handler http.Handler) *Server {
	return &Server{Handler: handler}
}

func (this Server) GetOrCreateRender() (Render http_render.Render) {
	return this.Render.Option(ropt.Dir(this.RootDir))
}

func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writed := &writed{ResponseWriter: w}
	this.Handler.ServeHTTP(writed, r)
	if !writed.writed {
		var Render http_render.Render
		if this.RenderFunc != nil {
			Render = this.RenderFunc(w, r)
		} else {
			Render = this.GetOrCreateRender().Option(ropt.Request(r))
		}
		r := Render.Option(ropt.FileNames(r.URL.Path))
		r.Status = http.StatusOK
		if this.RenderOrNotFoundDisabled {
			r.Render(w)
		} else {
			r.MustRenderOrNotFound(w)
		}
	}
}

type writed struct {
	http.ResponseWriter
	writed bool
}

func (this *writed) WriteHeader(s int) {
	this.writed = true
	this.ResponseWriter.WriteHeader(s)
}

func (this *writed) Write(p []byte) (n int, err error) {
	this.writed = true
	return this.ResponseWriter.Write(p)
}
