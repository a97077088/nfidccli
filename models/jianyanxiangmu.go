package models

type Jianyanxiangmu struct {
	Id string
	V抽样委托单号 string `gorm:"column:抽样委托单号"`
}

func (Jianyanxiangmu) TableName() string {
	return "检验任务"
}

