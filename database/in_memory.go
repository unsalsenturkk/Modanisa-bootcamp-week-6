package database

import (
	"fmt"
	"week-6-assignment/models"
)

type IDatabase interface {
	GetWallet() (*models.DataResponse, error)
	GetWalletByUsername(username string) (*models.Wallet, error)
	PutWalletByUsername(username string, balance float64) (*models.Wallet, error)
	PostWalletByUsername(username string, balance float64) (*models.Wallet, error)
}

type Database struct {
	wallet models.DataResponse
}

func (d *Database) GetWallet() (*models.DataResponse, error) {
	return &d.wallet, nil
}

func (d *Database) GetWalletByUsername(username string) (*models.Wallet, error) {

	v, ok := d.wallet[username]
	if !ok {
		return nil, fmt.Errorf("database Error : %s user's wallet could not be found", username)
	} else {
		return &v, nil
	}

}

func (d *Database) PutWalletByUsername(username string, balance float64) (*models.Wallet, error) {

	v, ok := d.wallet[username]
	if !ok {
		d.wallet[username] = models.Wallet{
			Username: username,
			Balance:  balance,
		}
		v = d.wallet[username]
		return &v, nil
	} else {
		return &v, nil
	}

}

func (d *Database) PostWalletByUsername(username string, balance float64) (*models.Wallet, error) {
	v, ok := d.wallet[username]
	if !ok {
		return nil, fmt.Errorf("database Error : %s user's wallet could not be found", username)
	} else {
		d.wallet[username] = models.Wallet{
			Username: username,
			Balance:  balance,
		}
		v = d.wallet[username]
		return &v, nil
	}
}

func NewDatabase(wallet models.DataResponse) IDatabase {
	return &Database{wallet: wallet}
}
