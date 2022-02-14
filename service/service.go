package service

import (
	"fmt"
	"week-6-assignment/config"
	"week-6-assignment/database"
	"week-6-assignment/models"
)

type IWallet interface {
	GetWallet() (*models.ServiceResponse, error)
	GetWalletByUsername(username string) (*models.Wallet, error)
	PutWalletByUsername(username string) (*models.Wallet, error)
	PostWalletByUsername(username string, balance float64) (*models.Wallet, error)
}

type Wallet struct {
	Database database.IDatabase
}

func (s *Wallet) GetWallet() (*models.ServiceResponse, error) {
	databaseRes, _ := s.Database.GetWallet()
	return (*models.ServiceResponse)(databaseRes), nil
}

func (s *Wallet) GetWalletByUsername(username string) (*models.Wallet, error) {
	databaseRes, err := s.Database.GetWalletByUsername(username)
	if err != nil {
		return nil, err
	}

	return databaseRes, nil
}

func (s *Wallet) PutWalletByUsername(username string) (*models.Wallet, error) {
	cfg := config.NewConfig()
	initialBalance := cfg.Get().InitialBalanceAmount

	databaseRes, err := s.Database.PutWalletByUsername(username, (float64)(initialBalance))
	if err != nil {
		return nil, err
	}
	return databaseRes, nil
}

func (s *Wallet) PostWalletByUsername(username string, balance float64) (*models.Wallet, error) {
	cfg := config.NewConfig()
	minimumBalanceAmount := cfg.Get().MinimumBalanceAmount

	userWallet, err := s.GetWalletByUsername(username)
	if err != nil {
		return nil, err
	}

	if (userWallet.Balance + balance) < (float64)(minimumBalanceAmount) {
		return nil, fmt.Errorf("service Error : insufficient balance")
	} else {
		userWallet.Balance += balance
	}

	databaseRes, er := s.Database.PostWalletByUsername(username, userWallet.Balance)
	if er != nil {
		return nil, er
	}

	return databaseRes, nil

}

func NewWalletService(database database.IDatabase) IWallet {
	return &Wallet{Database: database}
}
