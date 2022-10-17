package server

import (
	"kettle/pkg/accounts"
	"kettle/pkg/accounts/model"
	"kettle/utils/response"
	"net/http"
)

type (
	accountAPI struct {
		account accounts.IAccountService
	}
)

func NewAccountAPI(account accounts.IAccountService) *accountAPI {
	return &accountAPI{account: account}
}

func (a accountAPI) CreateAccount(w http.ResponseWriter, r *http.Request) {
	c := response.C{W: w, R: r}

	var request model.AccountModel
	c.BindJSON(&request)

	c.Response200(request, "account created")
}

func (a accountAPI) GetAccount(w http.ResponseWriter, r *http.Request) {
	c := response.C{W: w, R: r}

	c.Response200(nil, "account created")
}
