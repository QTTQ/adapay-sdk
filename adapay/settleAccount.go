package adapay

import (
	"errors"
	"fmt"
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type settleAccountInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type SettleAccount struct {
	*Adapay
}

func (s *SettleAccount) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + SETTLE_ACCOUNT_CREATE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, s.HandleConfig(multiMerchConfigId...))
}

func (s *SettleAccount) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {

	if reqParam["settle_account_id"] == nil || reqParam["settle_account_id"] == "" {
		return make(map[string]interface{}), nil, errors.New(fmt.Sprintf("请求参数错误 ==> %s", "settle_account_id"))
	}
	reqUrl := BASE_URL + strings.Replace(SETTLE_ACCOUNT_QUERY, "{settle_account_id}", adapayCore.ToString(reqParam["settle_account_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, s.HandleConfig(multiMerchConfigId...))
}

func (s *SettleAccount) Delete(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + SETTLE_ACCOUNT_DELETE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, s.HandleConfig(multiMerchConfigId...))
}

func (s *SettleAccount) QueryDetails(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + SETTLE_ACCOUNT_DETAIL_QUERY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, s.HandleConfig(multiMerchConfigId...))
}

func (s *SettleAccount) Modify(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + SETTLE_ACCOUNT_MODIFY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, s.HandleConfig(multiMerchConfigId...))
}

func (s *SettleAccount) QueryBalance(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + SETTLE_ACCOUNT_BALANCE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, s.HandleConfig(multiMerchConfigId...))
}
