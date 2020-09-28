package adapay

import (
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type refundInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Refund struct {
	*Adapay
}

func (r *Refund) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(REFUND_CREATE, "{payment_id}", adapayCore.ToString(reqParam["payment_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, r.HandleConfig(multiMerchConfigId...))
}

func (r *Refund) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + REFUND_QUERY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, r.HandleConfig(multiMerchConfigId...))
}
