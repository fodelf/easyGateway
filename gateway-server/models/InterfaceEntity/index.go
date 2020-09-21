package InterfaceEntity

//汇总实体类
type SumInfo struct {
	ServerSum  int `json:"serverSum"`
	WarningSum int `json:"warningSum"`
	RequestSum int `json:"requestSum"`
	FailSum    int `json:"failSum"`
}

// 图表实体类
type ChartInfo struct {
	ChartID  int    `json:"chart_id" gorm:"index"`
	Time     string `json:"time"`
	Total    int    `json:"total"`
	Success  int    `json:"success"`
	Fail     int    `json:"fail"`
	ServerId string `json:"serverId"`
}

// 图表列表实体类
// type ChartsInfo struct {
// 	TimeList    []string `json:"timeList"`
// 	TotalList   []int    `json:"totalList"`
// 	SuccessList []int    `json:"successList"`
// 	FailList    []int    `json:"failList"`
// }
// 图表实体类
type WarningInfo struct {
	WarningID int    `json:"warning_id" gorm:"index"`
	Time      string `json:"time"`
	System    string `json:"system"`
}
