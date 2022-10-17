package pkg

import (
	"github.com/gorilla/mux"
	"kettle/config/db"
	"kettle/pkg/accounts"
	"kettle/pkg/accounts/repo"
	"kettle/pkg/accounts/server"
	accountService "kettle/pkg/accounts/service"
)

func InitializeAccountService(db *db.MongoDB, router *mux.Router) accounts.IAccountService {
	service := accountService.NewAccountService(repo.NewAccountRepo(db))
	server.NewAccountAPI(service).Router(router)
	server.NewAccountGRPC(service)
	return service
}
