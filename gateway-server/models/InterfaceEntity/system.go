package InterfaceEntity

//注册中心
type ConsulInfo struct {
	ConsulId           string `json:"consulId"`
	ConsulAddress      string `json:"address"`
	ConsulPort         int    `json:"port"`
	Type                 string    `json:"type"`
}
//RabbitMQ中心
type RabbitMQInfo struct {
	RabbitMQId           string `json:"rabbitMQId"`
	RabbitMQAddress      string `json:"address"`
	RabbitMQPort         int    `json:"port"`
	RabbitMQUserName     string    `json:"userName"`
	RabbitMQPassword     string    `json:"password"`
	Type                 string    `json:"type"`
}
