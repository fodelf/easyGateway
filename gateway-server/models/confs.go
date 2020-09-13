package model

import InterfaceEntity "gateway/models/InterfaceEntity"

var Models = []interface{}{
	&InterfaceEntity.SumInfo{},
	&InterfaceEntity.ChartInfo{},
	&InterfaceEntity.ServiceInfo{},
	&InterfaceEntity.ConsulInfo{},
	&InterfaceEntity.RabbitMQInfo{},
}
