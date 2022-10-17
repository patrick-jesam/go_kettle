package server

import (
	"context"
	"kettle/pkg/accounts"
	"kettle/pkg/accounts/model"
)

type (
	accountGRPC struct {
		account accounts.IAccountService
	}
)

func NewAccountGRPC(account accounts.IAccountService) *accountGRPC {
	return &accountGRPC{account: account}
}

func (a accountGRPC) CreateAccount(ctx context.Context, model *model.AccountModel) (*model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountGRPC) CreateVaultAccount(ctx context.Context, model *model.AccountModel) (*model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountGRPC) GetAccount(ctx context.Context, model *model.AccountModel) (*model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}
