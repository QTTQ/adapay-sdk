package adapay

import (
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type corpMemberInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type CorpMember struct {
	*Adapay
}

func (c *CorpMember) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + CORP_MEMBERS_CREATE

	return adapayCore.UploadAdaPay(reqUrl, reqParam, "attach_file", c.HandleConfig(multiMerchConfigId...))
}

func (c *CorpMember) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(CORP_MEMBERS_QUERY, "{member_id}", adapayCore.ToString(reqParam["member_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, c.HandleConfig(multiMerchConfigId...))
}
