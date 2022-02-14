package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
	"week-6-assignment/models"
	"week-6-assignment/service"
)

var (
	GetWallet            = regexp.MustCompile(`^[\/]$`)
	GetWalletByUsername  = regexp.MustCompile(`^\/(\D+)$`)
	PutWalletByUsername  = regexp.MustCompile(`^\/(\D+)$`)
	PostWalletByUsername = regexp.MustCompile(`^\/(\D+)$`)
)

type IController interface {
	WalletHandler(w http.ResponseWriter, r *http.Request)
	GetWallet(w http.ResponseWriter, r *http.Request)
	GetWalletByUsername(w http.ResponseWriter, r *http.Request)
	PutWalletByUsername(w http.ResponseWriter, r *http.Request)
	PostWalletByUsername(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	Service service.IWallet
}

func (c *Controller) WalletHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && GetWallet.MatchString(r.URL.Path):
		c.GetWallet(w, r)
		return
	case r.Method == http.MethodGet && GetWalletByUsername.MatchString(r.URL.Path):
		c.GetWalletByUsername(w, r)
		return
	case r.Method == http.MethodPut && PutWalletByUsername.MatchString(r.URL.Path):
		c.PutWalletByUsername(w, r)
		return
	case r.Method == http.MethodPost && PostWalletByUsername.MatchString(r.URL.Path):
		c.PostWalletByUsername(w, r)
		return
	default:
		return
	}
}

func (c *Controller) GetWallet(w http.ResponseWriter, r *http.Request) {
	wallets, _ := c.Service.GetWallet()

	walletsBytes, _ := json.Marshal(wallets)

	w.WriteHeader(http.StatusOK)
	w.Write(walletsBytes)
}

func (c *Controller) GetWalletByUsername(w http.ResponseWriter, r *http.Request) {
	wallets, err := c.Service.GetWalletByUsername(strings.ReplaceAll(r.URL.Path, "/", ""))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	walletsBytes, _ := json.Marshal(wallets)

	w.WriteHeader(http.StatusOK)
	w.Write(walletsBytes)
}

func (c *Controller) PutWalletByUsername(w http.ResponseWriter, r *http.Request) {
	wallets, err := c.Service.PutWalletByUsername(strings.ReplaceAll(r.URL.Path, "/", ""))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	walletsBytes, _ := json.Marshal(wallets)

	w.WriteHeader(http.StatusOK)
	w.Write(walletsBytes)
}

func (c *Controller) PostWalletByUsername(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	wlt := models.Wlt{}
	jsonErr := json.Unmarshal(body, &wlt)
	if jsonErr != nil {
		http.Error(w, "can't read body ex:{'balance':100}", http.StatusBadRequest)
		return
	}

	wallets, err := c.Service.PostWalletByUsername(strings.ReplaceAll(r.URL.Path, "/", ""), wlt.Balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	walletsBytes, _ := json.Marshal(wallets)

	w.WriteHeader(http.StatusOK)
	w.Write(walletsBytes)

}

func NewWalletController(service service.IWallet) IController {
	return &Controller{Service: service}
}
