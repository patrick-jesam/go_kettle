package service

import (
	"kettle/pkg/accounts"
	"kettle/pkg/accounts/model"
)

type (
	accountService struct {
		repo accounts.IAccountService
	}
)

func NewAccountService(repo accounts.IAccountService) *accountService {
	return &accountService{repo: repo}
}

func (a accountService) CreateAccount(request model.AccountModel) (model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) CreateVaultAccount(request model.AccountModel) (model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) GetAccount(accountModel model.AccountModel) (model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}
