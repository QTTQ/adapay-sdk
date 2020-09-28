package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type cashsInterface interface {
	CreateCashs(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Cashs struct {
	*Adapay
}

func (c *Cashs) CreateCashs(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + BILL_DOWNLOAD
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, c.HandleConfig(multiMerchConfigId...))
}

func (c *Cashs) QueryCashsStat(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + QUERY_CASHS_STAT
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, c.HandleConfig(multiMerchConfigId...))
}
