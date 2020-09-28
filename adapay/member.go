package adapay

import (
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type memberInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Update(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Member struct {
	*Adapay
}

func (m *Member) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + MEMBER_CREATE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, m.HandleConfig(multiMerchConfigId...))
}

func (m *Member) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(MEMBER_QUERY, "{member_id}", adapayCore.ToString(reqParam["member_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, m.HandleConfig(multiMerchConfigId...))
}

func (m *Member) Update(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + MEMBER_UPDATE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, m.HandleConfig(multiMerchConfigId...))
}

func (m *Member) QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + MEMBER_QUERY_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, m.HandleConfig(multiMerchConfigId...))
}
