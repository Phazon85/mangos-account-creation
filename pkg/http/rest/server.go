package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fiseo/httpsrv"
	"github.com/husobee/vestigo"
	"github.com/phazon85/mangos-account-registration/pkg/acct"
	"go.uber.org/zap"
)

// Account ...
type Account interface {
	Register(req *acct.CreateRequest) error
	ResetPassword(req *acct.CreateRequest) error
}

// New implements the actions interface
func New(acct Account, logger *zap.Logger) *http.Server {
	r := newRouter(acct, logger)

	srv := httpsrv.NewWithDefault(r)
	return srv
}

// create account
// forgot password
//
// reset password
// remove/ban account

func newRouter(acct Account, logger *zap.Logger) http.Handler {
	router := vestigo.NewRouter()
	router.Get("/healthz", healthHandler())

	// API Endpoints
	router.Post("/account", postAccountHandler(logger, acct))
	router.Post("/account/resetpassword", postResetPasswordHandler(logger, acct))

	return router
}

func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("running"))
	}
}

func postAccountHandler(logger *zap.Logger, acctSvc Account) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var acctReq *acct.CreateRequest

		data, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Error(
				"post account handler error",
				zap.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(data, &acctReq)
		if err != nil {
			logger.Error(
				"post account handler error",
				zap.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = acctSvc.Register(acctReq)
		if err != nil {
			logger.Error(
				"post account handler error",
				zap.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "created"}`))
	}
}

func postResetPasswordHandler(logger *zap.Logger, acctSvc Account) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var acctReq *acct.CreateRequest

		data, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Error(
				"post account handler error",
				zap.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(data, &acctReq)
		if err != nil {
			logger.Error(
				"post account handler error",
				zap.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = acctSvc.ResetPassword(acctReq)
		if err != nil {
			logger.Error(
				"post reset password handler error",
				zap.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "updated"}`))
	}
}
