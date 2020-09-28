package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type unionInterface interface {
	UserIdentity(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Union struct {
	*Adapay
}

func (u *Union) UserIdentity(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + USER_IDENTITY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, u.HandleConfig(multiMerchConfigId...))
}
