package models

import "time"

type Jianyanxiangmu struct {
	Id      string
	V任务编号   string    `gorm:"column:任务编号"`
	V抽样委托单号 string    `gorm:"column:抽样委托单号"`
	V报告编号   string    `gorm:"column:报告编号"`
	V检验类型   string    `gorm:"column:检验类型"`
	V委托日期   time.Time `gorm:"column:委托日期"`
	V样品名称br string    `gorm:"column:样品名称br"`
	V规格型号   string    `gorm:"column:规格型号"`
	V样品数    string    `gorm:"column:样品数"`
	//V商定完成日期 string `gorm:"column:商定完成日期"`
	//V要求完成日期 string `gorm:"column:要求完成日期"`
	V样品等级 string `gorm:"column:样品等级"`
	V样品类型 string `gorm:"column:样品类型"`
	//V生产日期 string `gorm:"column:生产日期"`
	V保存条件    string `gorm:"column:保存条件"`
	V商标      string `gorm:"column:商标"`
	V生产批号    string `gorm:"column:生产批号"`
	V样品状态    string `gorm:"column:样品状态"`
	V保质期     string `gorm:"column:保质期"`
	V样品状态3   string `gorm:"column:样品状态3"`
	V业务受理人   string `gorm:"column:业务受理人"`
	V销售员     string `gorm:"column:销售员"`
	V委托单位    string `gorm:"column:委托单位"`
	V委托单位地址  string `gorm:"column:委托单位地址"`
	V委托人     string `gorm:"column:委托人"`
	V委托单位电话  string `gorm:"column:委托单位电话"`
	V受检单位    string `gorm:"column:受检单位"`
	V地址      string `gorm:"column:地址"`
	V联系人     string `gorm:"column:联系人"`
	V电话      string `gorm:"column:电话"`
	V生产单位    string `gorm:"column:生产单位"`
	V生产单位地址  string `gorm:"column:生产单位地址"`
	V生产单位联系人 string `gorm:"column:生产单位联系人"`
	V生产单位电话  string `gorm:"column:生产单位电话"`
	V检验依据    string `gorm:"column:检验依据"`
	V抽样单号    string `gorm:"column:抽样单号"`
	V抽送样人    string `gorm:"column:抽送样人"`
	V抽样方式    string `gorm:"column:抽样方式"`
	V抽样地点    string `gorm:"column:抽样地点"`
	V抽样基数    string `gorm:"column:抽样基数"`
	//V抽到样日期 string `gorm:"column:抽到样日期"`
	V备注 string `gorm:"column:备注"`
}

func (Jianyanxiangmu) TableName() string {
	return "检验任务"
}
