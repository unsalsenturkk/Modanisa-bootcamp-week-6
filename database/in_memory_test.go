package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"week-6-assignment/models"
)

func TestDatabase_GetWallet(t *testing.T) {

	t.Run("when the wallet is in memory then get all", func(t *testing.T) {
		wallet := models.DataResponse{
			"username1": models.Wallet{
				Username: "username1",
				Balance:  0,
			},
			"username2": models.Wallet{
				Username: "username2",
				Balance:  100,
			},
			"username3": models.Wallet{
				Username: "username3",
				Balance:  1000,
			},
		}

		data := NewDatabase(wallet)
		response, err := data.GetWallet()

		assert.Nil(t, err)
		assert.Equal(t, &wallet, response)
	})
	t.Run("when the wallet is not in memory then get all", func(t *testing.T) {
		wallet := models.DataResponse{}

		data := NewDatabase(wallet)
		response, err := data.GetWallet()

		assert.Nil(t, err)
		assert.Equal(t, &wallet, response)
	})
}

func TestDatabase_GetWalletByUsername(t *testing.T) {

	type fields struct {
		wallet models.DataResponse
	}
	type args struct {
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.Wallet
	}{
		{
			name: "when there is a wallet owned by the username then get wallet by username ",
			fields: fields{wallet: models.DataResponse{
				"username1": models.Wallet{
					Username: "username1",
					Balance:  100,
				},
			}},
			args: args{username: "username1"},
			want: &models.Wallet{
				Username: "username1",
				Balance:  100,
			},
		}, {
			name: "when there is no wallet owned by the username then get wallet by username",
			fields: fields{wallet: models.DataResponse{
				"username1": models.Wallet{
					Username: "username1",
					Balance:  100,
				},
			}},
			args: args{username: "username2"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Database{
				wallet: tt.fields.wallet,
			}
			got, err := d.GetWalletByUsername(tt.args.username)
			if err != nil {
				assert.EqualError(t, err, "database Error : "+tt.args.username+" user's wallet could not be found")
			}

			assert.Equalf(t, tt.want, got, "GetWalletByUsername(%v)", tt.args.username)
		})
	}
}

func TestDatabase_PutWalletByUsername(t *testing.T) {
	type fields struct {
		wallet models.DataResponse
	}
	type args struct {
		username string
		balance  float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.Wallet
	}{
		{
			name: "when there is a wallet owned by the username then put wallet by username",
			fields: fields{wallet: models.DataResponse{
				"username1": models.Wallet{
					Username: "username1",
					Balance:  100,
				},
			}},
			args: args{username: "username1", balance: 200},
			want: &models.Wallet{
				Username: "username1",
				Balance:  100,
			},
		}, {
			name: "when there is no wallet owned by the username then put wallet by username",
			fields: fields{wallet: models.DataResponse{
				"username1": models.Wallet{
					Username: "username1",
					Balance:  100,
				},
			}},
			args: args{username: "username2", balance: 200},
			want: &models.Wallet{
				Username: "username2",
				Balance:  200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Database{
				wallet: tt.fields.wallet,
			}
			got, _ := d.PutWalletByUsername(tt.args.username, tt.args.balance)
			assert.Equalf(t, tt.want, got, "PutWalletByUsername(%v, %v)", tt.args.username, tt.args.balance)
		})
	}
}

func TestDatabase_PostWalletByUsername(t *testing.T) {
	type fields struct {
		wallet models.DataResponse
	}
	type args struct {
		username string
		balance  float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.Wallet
	}{
		{
			name: "when there is a wallet owned by the username then post wallet by username",
			fields: fields{wallet: models.DataResponse{
				"username1": models.Wallet{
					Username: "username1",
					Balance:  100,
				},
			}},
			args: args{username: "username1", balance: 200},
			want: &models.Wallet{
				Username: "username1",
				Balance:  200,
			},
		}, {
			name: "when there is no wallet owned by the username then post wallet by username",
			fields: fields{wallet: models.DataResponse{
				"username1": models.Wallet{
					Username: "username1",
					Balance:  100,
				},
			}},
			args: args{username: "username2", balance: 200},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Database{
				wallet: tt.fields.wallet,
			}
			got, err := d.PostWalletByUsername(tt.args.username, tt.args.balance)
			if err != nil {
				assert.EqualError(t, err, "database Error : "+tt.args.username+" user's wallet could not be found")
			}
			assert.Equalf(t, tt.want, got, "PostWalletByUsername(%v, %v)", tt.args.username, tt.args.balance)
		})
	}
}
