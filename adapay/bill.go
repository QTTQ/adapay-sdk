package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type billInterface interface {
	Download(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Bill struct {
	*Adapay
}

func (b *Bill) Download(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + BILL_DOWNLOAD
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, b.HandleConfig(multiMerchConfigId...))
}
