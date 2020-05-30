package models

import (
	"gomall-center/pkg/enum"

	"github.com/jinzhu/gorm"
)

// Company ...Company表
type Company struct {
	gorm.Model
	CompanyName     string             `json:"company_name"`
	BusinessLicense string             `json:"business_license"`  // 营业执照号
	LegalPersonName string             `json:"legal_person_name"` // 法人姓名
	LegalPersonID   string             `json:"legal_person_id"`   // 法人身份证
	Taxpayer        string             `json:"taxpayer"`          // 纳税识别号
	Contacts        string             `json:"contacts"`          // 联系人
	Tel             string             `json:"tel"`               // 联系电话
	Lat             string             `json:"lat"`               // 经度
	Lng             string             `json:"lng"`               // 维度
	Address         string             `json:"address"`           // 地址
	Province        string             `json:"province"`          // 省份
	City            string             `json:"city"`              // 城市
	Area            string             `json:"area"`              // 地区
	Logo            string             `json:"logo"`              // 企业logo
	CompanyType     enum.CompanyType   `json:"company_type"`      // 企业类型
	CompanyStatus   enum.CompanyStatus `json:"company_status"`
	LegalTypes      int                `json:"legal_types"` //1 法人 2 不是法人
	CompanyProfile  string             `json:"company_profile"`
	ProvinceID      int                `json:"province_id"`
	CityID          int                `json:"city_id"`
	AreaID          int                `json:"area_id"`
	CompanyNature   string             `json:"company_nature"`  //企业性质：可选性参见数据字典
	AuthVip         int                `json:"auth_vip"`        //大客户认证：0 否 1：是
	AuthAgent       int                `json:"auth_agent"`      //0: 没有代理权限 1：开通找代理权限
	AuthAllocation  int                `json:"auth_allocation"` //调拨 0 没有开启调拨 1 已开启调拨
	Shop            []*Shop            `json:"shop"`
	Accounts        []*Account
}

// CompanyAbs ...Company列表
type CompanyAbs struct {
	ID              int64              `json:"id"`
	CompanyName     string             `json:"company_name"`
	BusinessLicense string             `json:"business_license"`  // 营业执照号
	LegalPersonName string             `json:"legal_person_name"` // 法人姓名
	LegalPersonID   string             `json:"legal_person_id"`   // 法人身份证
	Taxpayer        string             `json:"taxpayer"`          // 纳税识别号
	Contacts        string             `json:"contacts"`          // 联系人
	Tel             string             `json:"tel"`               // 联系电话
	Lat             string             `json:"lat"`               // 经度
	Lng             string             `json:"lng"`               // 维度
	Address         string             `json:"address"`           // 地址
	Province        string             `json:"province"`          // 省份
	City            string             `json:"city"`              // 城市
	Area            string             `json:"area"`              // 地区
	Logo            string             `json:"logo"`              // 企业logo
	CompanyType     enum.CompanyType   `json:"company_type"`      // 企业类型
	CompanyStatus   enum.CompanyStatus `json:"company_status"`
}

// CompanyFile ...Company_File资质审核附件
type CompanyFile struct {
	gorm.Model
	FileName  string `gorm:"column:file_name;type:varchar(45)" json:"file_name"`
	FileURL   string `gorm:"column:file_url;type:varchar(45)" json:"file_url"`
	CompanyID int64  `gorm:"column:company_id;type:bigint(20)" json:"company_id"`
}

// Companys ...详情表
type Companys struct {
	ID              int64              `json:"id"`
	CompanyName     string             `json:"company_name"`
	BusinessLicense string             `json:"business_license"`  // 营业执照号
	LegalPersonName string             `json:"legal_person_name"` // 法人姓名
	LegalPersonID   string             `json:"legal_person_id"`   // 法人身份证
	Taxpayer        string             `json:"taxpayer"`          // 纳税识别号
	Contacts        string             `json:"contacts"`          // 联系人
	Tel             string             `json:"tel"`               // 联系电话
	Lat             string             `json:"lat"`               // 经度
	Lng             string             `json:"lng"`               // 维度
	Address         string             `json:"address"`           // 地址
	Province        string             `json:"province"`          // 省份
	City            string             `json:"city"`              // 城市
	Area            string             `json:"area"`              // 地区
	Logo            string             `json:"logo"`              // 企业logo
	CompanyType     enum.CompanyType   `json:"company_type"`      // 企业类型
	CompanyStatus   enum.CompanyStatus `json:"company_status"`
	CompanyFile     []*CompanyFile     `json:"company_file"`
	Shop            []*Shop            `json:"shop"`
	Accounts        []*Account         `json:"account"`
}

// CompanyPut ...Company修改企业状态，审核企业资质
type CompanyPut struct {
	CompanyStatus enum.CompanyStatus `json:"company_status" validate:"required"`
	//CompanyNature  string             `json:"company_nature" validate:"required"`
	AuthVip        int `json:"auth_vip" `
	AuthAgent      int `json:"auth_agent" `
	AuthAllocation int `json:"auth_allocation" `
	//CompanyProfile string             `json:"company_profile" validate:"required"`
}
