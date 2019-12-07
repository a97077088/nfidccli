package models

type Jianyanxiangmu struct {
	V任务编号 string `gorm:"column:任务编号"`
	V项目名称 string `gorm:"column:项目名称"`
}

func (Jianyanxiangmu) TableName() string {
	return "检验项目"
}
