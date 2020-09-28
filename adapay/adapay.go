package adapay

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type Adapay struct {

	MultiMerchSysConfigs map[string]*adapayCore.MerchSysConfig

	DefaultMerchSysConfig *adapayCore.MerchSysConfig
}


func InitDefaultMerchSysConfig(filePath string) (*Adapay, error) {

	config, err := adapayCore.ReadMerchConfig(filePath)
	if err != nil {
		return nil, err
	}

	ada := &Adapay{
		DefaultMerchSysConfig: config,
	}

	return ada, nil
}


func InitMultiMerchSysConfigs(fileDir string) (*Adapay, error) {

	dirs, _ := ioutil.ReadDir(fileDir)

	configs := map[string]*adapayCore.MerchSysConfig{}

	for _, f := range dirs {

		ext := filepath.Ext(f.Name())
		if ext == ".json" {
			config, err := adapayCore.ReadMerchConfig(fileDir + f.Name())
			if err != nil {
				continue
			}

			key := strings.Replace(f.Name(), ".json", "", -1)
			configs[key] = config
		}
	}

	ada := &Adapay{
		MultiMerchSysConfigs: configs,
	}

	return ada, nil
}


func (a *Adapay) HandleConfig(multiMerchConfigId ...string) *adapayCore.MerchSysConfig {
	if multiMerchConfigId == nil {
		return a.DefaultMerchSysConfig
	} else {
		return a.MultiMerchSysConfigs[multiMerchConfigId[0]]
	}
}


func (a *Adapay) Payment() *Payment {
	return &Payment{Adapay: a}
}


func (a *Adapay) SettleAccount() *SettleAccount {
	return &SettleAccount{Adapay: a}
}


func (a *Adapay) Bill() *Bill {
	return &Bill{Adapay: a}
}


func (a *Adapay) Cashs() *Cashs {
	return &Cashs{Adapay: a}
}


func (a *Adapay) CorpMember() *CorpMember {
	return &CorpMember{Adapay: a}
}


func (a *Adapay) Member() *Member {
	return &Member{Adapay: a}
}


func (a *Adapay) Refund() *Refund {
	return &Refund{Adapay: a}
}


func (a *Adapay) Union() *Union {
	return &Union{Adapay: a}
}


func (a *Adapay) Wallet() *Wallet {
	return &Wallet{Adapay: a}
}


func (a *Adapay) Version() string {
	return "1.1.1"
}
