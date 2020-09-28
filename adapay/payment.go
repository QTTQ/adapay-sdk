package adapay

import (
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type paymentInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Close(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	CreateConfirm(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryConfirm(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryConfirmList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	CreateReverse(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryReverse(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryReverseList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Payment struct {
	*Adapay
}

func (p *Payment) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_CREATE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(PAYMENT_QUERY, "{payment_id}", adapayCore.ToString(reqParam["payment_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) Close(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(PAYMENT_CLOSE, "{payment_id}", adapayCore.ToString(reqParam["payment_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) CreateConfirm(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_CONFIRM
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) QueryConfirm(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(PAYMENT_QUERY_CONFIRM, "{payment_confirm_id}", adapayCore.ToString(reqParam["payment_confirm_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) QueryConfirmList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_QUERY_CONFIRM_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) CreateReverse(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_REVERSE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) QueryReverse(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(PAYMENT_QUERY_REVERSE, "{reverse_id}", adapayCore.ToString(reqParam["reverse_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) QueryReverseList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_QUERY_REVERSE_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}
