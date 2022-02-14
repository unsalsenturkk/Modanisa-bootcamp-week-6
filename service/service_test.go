package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"week-6-assignment/mock"

	"week-6-assignment/models"
)

func TestWallet_GetWallet(t *testing.T) {
	databaseReturn := &models.DataResponse{
		"username1": models.Wallet{
			Username: "username1",
			Balance:  100,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDatabase := mock.NewMockIDatabase(ctrl)
	mockDatabase.
		EXPECT().
		GetWallet().
		Return(databaseReturn, nil).
		Times(1)

	walletService := NewWalletService(mockDatabase)
	wallets, err := walletService.GetWallet()
	assert.Equal(t, (*models.ServiceResponse)(databaseReturn), wallets)
	assert.Nil(t, err)
}

func TestWallet_GetWalletByUsername(t *testing.T) {
	t.Run("when GetWalletbyUsername returns data properly", func(t *testing.T) {
		databaseReturn := &models.Wallet{
			Username: "username1",
			Balance:  100,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDatabase := mock.NewMockIDatabase(ctrl)
		mockDatabase.
			EXPECT().
			GetWalletByUsername("username1").
			Return(databaseReturn, nil).
			Times(1)

		walletService := NewWalletService(mockDatabase)
		wallets, err := walletService.GetWalletByUsername("username1")
		assert.Equal(t, databaseReturn, wallets)
		assert.Nil(t, err)
	})
	t.Run("when GetWalletbyUsername returns error", func(t *testing.T) {

		error := fmt.Errorf("database Error : %s user's wallet could not be found", "username1")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDatabase := mock.NewMockIDatabase(ctrl)
		mockDatabase.
			EXPECT().
			GetWalletByUsername("username1").
			Return(nil, error).
			Times(1)

		walletService := NewWalletService(mockDatabase)
		_, err := walletService.GetWalletByUsername("username1")
		assert.Equal(t, error, err)
	})

}

func TestWallet_PutWalletByUsername(t *testing.T) {
	t.Run("when PutWalletByUsername returns data properly", func(t *testing.T) {
		databaseReturn := &models.Wallet{
			Username: "username1",
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDatabase := mock.NewMockIDatabase(ctrl)
		mockDatabase.
			EXPECT().
			PutWalletByUsername("username1", (float64)(0)).
			Return(databaseReturn, nil).
			Times(1)

		walletService := NewWalletService(mockDatabase)
		wallets, err := walletService.PutWalletByUsername("username1")
		assert.Equal(t, databaseReturn, wallets)
		assert.Nil(t, err)
	})
	t.Run("when PutWalletByUsername returns error", func(t *testing.T) {
		error := fmt.Errorf("mock error")
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDatabase := mock.NewMockIDatabase(ctrl)
		mockDatabase.
			EXPECT().
			PutWalletByUsername("username1", (float64)(0)).
			Return(nil, error).
			Times(1)

		walletService := NewWalletService(mockDatabase)
		_, err := walletService.PutWalletByUsername("username1")
		assert.Equal(t, error, err)

	})

}

func TestWallet_PostWalletByUsername(t *testing.T) {

	t.Run("when there is enough balance in the wallet", func(t *testing.T) {
		userWallet := &models.Wallet{
			Username: "username1",
			Balance:  100,
		}

		databaseReturn := &models.Wallet{
			Username: "username1",
			Balance:  50,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDatabase := mock.NewMockIDatabase(ctrl)

		mockDatabase.
			EXPECT().
			GetWalletByUsername("username1").
			Return(userWallet, nil).
			Times(1)

		mockDatabase.
			EXPECT().
			PostWalletByUsername("username1", float64(50)).
			Return(databaseReturn, nil).
			Times(1)

		walletService := NewWalletService(mockDatabase)
		wallets, err := walletService.PostWalletByUsername("username1", float64(-50))
		assert.Equal(t, databaseReturn, wallets)
		assert.Nil(t, err)

	})

	t.Run("when there is not enough balance in the wallet", func(t *testing.T) {
		userWallet := &models.Wallet{
			Username: "username1",
			Balance:  100,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDatabase := mock.NewMockIDatabase(ctrl)

		mockDatabase.
			EXPECT().
			GetWalletByUsername("username1").
			Return(userWallet, nil).
			Times(1)

		walletService := NewWalletService(mockDatabase)
		_, err := walletService.PostWalletByUsername("username1", (float64)(-1000))
		assert.EqualError(t, err, "service Error : insufficient balance")

	})

	t.Run("when PostWalletByUsername returns error", func(t *testing.T) {
		userWallet := &models.Wallet{
			Username: "username1",
			Balance:  100,
		}

		error := fmt.Errorf("service Error : insufficient balance")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockDatabase := mock.NewMockIDatabase(ctrl)

		mockDatabase.
			EXPECT().
			GetWalletByUsername("username1").
			Return(userWallet, nil).
			Times(1)

		walletService := NewWalletService(mockDatabase)

		w, err := walletService.PostWalletByUsername("username1", float64(-1000))
		assert.Equal(t, error, err)
		assert.Nil(t, w)

	})

}
