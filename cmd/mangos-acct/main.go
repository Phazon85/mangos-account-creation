package main

import (
	"fmt"
	"log"
	"os"

	"github.com/phazon85/mangos-account-registration/pkg/acct"
	"github.com/phazon85/mangos-account-registration/pkg/http/rest"
	"github.com/phazon85/mangos-account-registration/pkg/repository/pgsql"
	"github.com/phazon85/mangos-account-registration/pkg/sqldb"
	"go.uber.org/zap"
)

const (
	// DBCONN ...
	DBCONN = "MANGOS_ACCT_DBCONN"
	// PORT ...
	PORT = "MANGOS_ACCT_PORT"
	// DEFAULTPORT ...
	DEFAULTPORT = "9000"
	// ENV ...
	ENV = "MANGOS_ACCT_ENV"
	// DEFAULTENV ...
	DEFAULTENV = "dev"
)

func main() {
	// load environment
	dbconn := os.Getenv(DBCONN)
	if dbconn == "" {
		log.Fatal("db connection required")
	}

	port := os.Getenv(PORT)
	if port == "" {
		port = DEFAULTPORT
	}

	env := os.Getenv(ENV)
	if env == "" {
		env = DEFAULTENV
	}

	var logger *zap.Logger
	var err error
	switch env {
	case "dev":
		logger, err = zap.NewDevelopment()
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatal(err)
	}

	db := sqldb.New(dbconn)
	pgsql := pgsql.New(db)
	acc := acct.New(pgsql)

	api := rest.New(acc, logger)
	api.Addr = fmt.Sprintf(":%s", port)
	logger.Info("server start", zap.String("port", port))
	log.Fatal(api.ListenAndServe())
}
