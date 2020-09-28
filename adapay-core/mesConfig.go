package adapayCore

import (
	"errors"
)

type MsgConfig struct {
	Base_url       string
	Instance_id    string
	Access_key     string
	Group_id       string
	Client_id      string
	Broker_url     string
	Api_key        string
	Private_key    string
	Public_key     string
	MerchConfigId  string
	MerchSysConfig *MerchSysConfig
}


func InitConfig(merchConfigPath string) (*MsgConfig, error) {


	msc, configErr := ReadMerchConfig(merchConfigPath)
	if configErr != nil {
		return nil, errors.New("====== 读取配置文件异常 ======")
	}

	msg := &MsgConfig{
		Base_url:       "https://api.adapay.tech",
		Instance_id:    "post-cn-0pp18zowf0m",
		Access_key:     "LTAIOP5RkeiuXieW",
		Group_id:       "GID_CRHS_ASYN",
		Broker_url:     "post-cn-0pp18zowf0m.mqtt.aliyuncs.com",
		Api_key:        msc.ApiKeyLive,
		Private_key:    msc.RspPriKey,
		Public_key:     msc.RspPubKey,
		MerchSysConfig: msc,
	}
	return msg, nil
}
