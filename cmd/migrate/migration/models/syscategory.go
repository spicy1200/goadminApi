package models

type SysCategory struct {
	Model
	Name   string `json:"name" gorm:"type:varchar(255);comment:名称"`   //
	Img    string `json:"img" gorm:"type:varchar(255);comment:图标"`    //
	Sort   int    `json:"sort" gorm:"type:int(4);comment:排序"`         //
	Status int    `json:"status" gorm:"type:int(1);comment:状态"`       //
	Remark string `json:"remark" gorm:"type:varchar(255);comment:备注"` //
	ControlBy
	ModelTime
}

func (SysCategory) TableName() string {
	return "sys_category"
}
