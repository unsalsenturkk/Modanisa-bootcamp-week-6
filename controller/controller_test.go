package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"week-6-assignment/database"
	"week-6-assignment/mock"
	"week-6-assignment/models"
	"week-6-assignment/service"
)

func TestController_GetWallet(t *testing.T) {
	t.Run("when service data returns", func(t *testing.T) {
		serviceReturn := &models.ServiceResponse{
			"username1": models.Wallet{
				Username: "username1",
				Balance:  0,
			}}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)
		mockService.EXPECT().GetWallet().Return(serviceReturn, nil).Times(1)

		controller := NewWalletController(mockService)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		controller.GetWallet(w, r)

		actual := &models.ServiceResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})

}

func TestController_GetWalletByUsername(t *testing.T) {
	t.Run("when service data returns", func(t *testing.T) {
		serviceReturn := &models.Wallet{
			Username: "username1",
			Balance:  100,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)
		mockService.EXPECT().GetWalletByUsername("username1").Return(serviceReturn, nil).Times(1)

		controller := NewWalletController(mockService)
		r := httptest.NewRequest(http.MethodGet, "/username1", nil)
		w := httptest.NewRecorder()
		controller.GetWalletByUsername(w, r)

		actual := &models.Wallet{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})
	t.Run("when service error returns", func(t *testing.T) {

		error := fmt.Errorf("database Error : %s user's wallet could not be found", "username1")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)
		mockService.EXPECT().GetWalletByUsername("username1").Return(nil, error).Times(1)

		controller := NewWalletController(mockService)
		r := httptest.NewRequest(http.MethodGet, "/username1", nil)
		w := httptest.NewRecorder()
		controller.GetWalletByUsername(w, r)

		actual := w.Body.String()

		assert.Equal(t, error.Error(), actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	})
}

func TestController_PostWalletByUsername(t *testing.T) {
	t.Run("when service error returns", func(t *testing.T) {
		error := fmt.Errorf("can't read body ex:{'balance':100}\n")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)
		mockService.EXPECT().PostWalletByUsername("username1", float64(100)).Return(nil, error).Times(0)

		controller := NewWalletController(mockService)
		r := httptest.NewRequest(http.MethodPost, "/username1", nil)
		w := httptest.NewRecorder()
		controller.PostWalletByUsername(w, r)

		actual := w.Body.String()

		assert.Equal(t, error.Error(), actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	})
	t.Run("when service data returns", func(t *testing.T) {
		wallet := models.DataResponse{
			"username1": models.Wallet{
				Username: "username1",
				Balance:  0,
			},
		}

		data := database.NewDatabase(wallet)
		svc := service.NewWalletService(data)
		controller := NewWalletController(svc)

		r := httptest.NewRequest(http.MethodPost, "/username1", strings.NewReader("{\"balance\":100}"))
		w := httptest.NewRecorder()

		controller.PostWalletByUsername(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)

	})

}

func TestController_PutWalletByUsername(t *testing.T) {
	t.Run("when service data returns", func(t *testing.T) {
		serviceReturn := &models.Wallet{
			Username: "username1",
			Balance:  100,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)
		mockService.EXPECT().PutWalletByUsername("username1").Return(serviceReturn, nil).Times(1)

		controller := NewWalletController(mockService)
		r := httptest.NewRequest(http.MethodPut, "/username1", nil)
		w := httptest.NewRecorder()
		controller.PutWalletByUsername(w, r)

		actual := &models.Wallet{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})
}

func TestController_WalletHandler(t *testing.T) {

	t.Run("when all user wallets are listed ", func(t *testing.T) {
		serviceReturn := &models.ServiceResponse{
			"username1": models.Wallet{
				Username: "username1",
				Balance:  0,
			}}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)

		mockService.EXPECT().GetWallet().Return(serviceReturn, nil).Times(1)

		controller := NewWalletController(mockService)

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		controller.WalletHandler(w, r)

		actual := &models.ServiceResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, "application/json", w.Header().Get("content-type"))
		assert.True(t, GetWallet.MatchString(r.URL.Path))
		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})

	t.Run("when a single user wallet is listed", func(t *testing.T) {
		serviceReturn := &models.Wallet{
			Username: "username",
			Balance:  100,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)

		mockService.EXPECT().GetWalletByUsername("username").Return(serviceReturn, nil).Times(1)

		controller := NewWalletController(mockService)

		r := httptest.NewRequest(http.MethodGet, "/username", nil)
		w := httptest.NewRecorder()
		controller.WalletHandler(w, r)

		actual := &models.Wallet{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, "application/json", w.Header().Get("content-type"))
		assert.True(t, GetWalletByUsername.MatchString(r.URL.Path))
		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)

	})

	t.Run("when new wallet is added", func(t *testing.T) {
		serviceReturn := &models.Wallet{
			Username: "username1",
			Balance:  100,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)

		mockService.EXPECT().PutWalletByUsername("username").Return(serviceReturn, nil).Times(1)

		controller := NewWalletController(mockService)

		r := httptest.NewRequest(http.MethodPut, "/username", nil)
		w := httptest.NewRecorder()
		controller.WalletHandler(w, r)

		actual := &models.Wallet{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, "application/json", w.Header().Get("content-type"))
		assert.True(t, PutWalletByUsername.MatchString(r.URL.Path))
		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})

	t.Run("when adding money", func(t *testing.T) {
		serviceReturn := &models.Wallet{
			Username: "username1",
			Balance:  100,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := mock.NewMockIWallet(ctrl)

		mockService.EXPECT().PostWalletByUsername("username", float64(100)).Return(serviceReturn, nil).Times(1)

		controller := NewWalletController(mockService)

		r := httptest.NewRequest(http.MethodPost, "/username", nil)
		w := httptest.NewRecorder()
		r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("{\n    \"balance\" : 100\n}")))
		controller.WalletHandler(w, r)

		actual := &models.Wallet{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, "application/json", w.Header().Get("content-type"))
		assert.True(t, PostWalletByUsername.MatchString(r.URL.Path))
		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})

}

func TestNewWalletController(t *testing.T) {
	type args struct {
		service service.IWallet
	}
	tests := []struct {
		name string
		args args
		want IController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWalletController(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWalletController() = %v, want %v", got, tt.want)
			}
		})
	}
}
