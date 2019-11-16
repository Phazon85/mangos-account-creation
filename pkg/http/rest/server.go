package rest

import (
	"net/http"

	"github.com/fiseo/httpsrv"
	"github.com/husobee/vestigo"
)

// Authorizor ...
type Authorizor interface {
	Register()
}

//NewHandler implements the actions interface
func New(auth Authorizor) *http.Server {
	r := newRouter(auth)

	srv := httpsrv.NewWithDefault(r)
	return srv
}

func newRouter(auth Authorizor) http.Handler {
	router := vestigo.NewRouter()
	router.Get("/healthz", healthHandler())

	return router
}

func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("running"))
	}
}
