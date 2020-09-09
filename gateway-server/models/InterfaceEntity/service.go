package InterfaceEntity

import (
	conf "gateway/conf"
)

//汇总实体类
type ServiceInfo struct {
	conf.Model
	ServerId            string `json:"serverId"`
	ServiceName         string `json:"serviceName"`
	ServiceType         string `json:"serviceType"`
	ServiceAddress      string `json:"serviceAddress"`
	ServicePort         int    `json:"servicePort"`
	ServiceLimit        int    `json:"serviceLimit"`
	ServiceBreak        int    `json:"serviceBreak"`
	ServiceRules        string `json:"serviceRules"`
	UseConsulId         string `json:"useConsulId"`
	UseConsulTag        string `json:"useConsulTag"`
	UseConsulCheckPath  string `json:"useConsulCheckPath"`
	UseConsulPort       int    `json:"useConsulPort"`
	UseConsulInterval   int    `json:"useConsulInterval"`
	UseConsulTimeout    int    `json:"useConsulTimeout"`
	DingdingAccessToken string `json:"dingdingAccessToken"`
	DingdingSecret      string `json:"dingdingSecret"`
	DingdingList        string `json:"dingdingList"`
}
