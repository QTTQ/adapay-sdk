package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type walletInterface interface {
	WalletLogin(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Pay(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Wallet struct {
	*Adapay
}

func (w *Wallet) WalletLogin(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + WALLET_LOGIN
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, w.HandleConfig(multiMerchConfigId...))
}

func (w *Wallet) Pay(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + WALLET_PAY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, w.HandleConfig(multiMerchConfigId...))
}

func (w *Wallet) CreateCheckOut(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + WALLET_CHECKOUT
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, w.HandleConfig(multiMerchConfigId...))
}
