package accounts

import (
	"context"
	"kettle/pkg/accounts/model"
)

type (
	IAccountRepo interface {
		CreateAccount(request model.AccountModel) (model.AccountModel, error)
		CreateVaultAccount(request model.AccountModel) (model.AccountModel, error)
		GetAccount(accountModel model.AccountModel) (model.AccountModel, error)
	}

	IAccountService interface {
		IAccountRepo
	}
)

type MockGRPCInterface interface {
	CreateAccount(context.Context, *model.AccountModel) (*model.AccountModel, error)
	CreateVaultAccount(context.Context, *model.AccountModel) (*model.AccountModel, error)
	GetAccount(context.Context, *model.AccountModel) (*model.AccountModel, error)
}
