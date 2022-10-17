package repo

import (
	"kettle/config/db"
	"kettle/pkg/accounts/model"
)

type (
	accountRepo struct {
		db *db.MongoDB
	}
)

func NewAccountRepo(db *db.MongoDB) *accountRepo {
	return &accountRepo{db: db}
}

func (a accountRepo) CreateAccount(request model.AccountModel) (model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountRepo) CreateVaultAccount(request model.AccountModel) (model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountRepo) GetAccount(accountModel model.AccountModel) (model.AccountModel, error) {
	//TODO implement me
	panic("implement me")
}
