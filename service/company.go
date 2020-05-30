package service

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
)

//CompanyService ...
type CompanyService struct {
	Service
}

const (
	// 列表查询返回字段
	selectListFields = "id, company_name,company_status,business_license, legal_person_name,legal_person_id"
)

//NewCompanyService new CompanyService
func NewCompanyService(ctx *web.RequestContext) *CompanyService {
	return &CompanyService{
		Service: InitService(ctx),
	}
}

//CompanyList ...获取企业列表
func (c *CompanyService) CompanyList(args *models.PagerArgs) *models.PagerData {
	var commands []*models.CompanyAbs
	var count int
	db := c.Table("t_company")
	if status, ok := args.Filter["company_status"]; ok {
		db = c.Table("t_company").Select(selectListFields).Where("company_status = ? AND company_name like ?", status, "%"+args.KeyWord+"%")
	}
	db.Count(&count)
	db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&commands)
	results := &models.PagerData{
		Count: count,
		Data:  commands,
	}
	return results
}

//CompanyByID ...根据企业id获取企业详情
func (c *CompanyService) CompanyByID(id int64) *models.Companys {
	result := &models.Company{}
	shop := make([]*models.Shop, 0)
	account := make([]*models.Account, 0)
	companyFile := make([]*models.CompanyFile, 0)
	c.Table("t_company").Where("id = ?", id).First(&result)
	c.Table("t_company_file").Where("company_id = ?", id).Find(&companyFile)
	//c.Table("t_shop").Where("company_id = ?", id).Find(&shop)
	c.Table("t_account").Where("company_id = ?", id).Find(&account)
	company := &models.Companys{
		ID:              id,
		CompanyName:     result.CompanyName,
		BusinessLicense: result.BusinessLicense,
		LegalPersonName: result.LegalPersonName,
		LegalPersonID:   result.LegalPersonID,
		Taxpayer:        result.Taxpayer,
		Contacts:        result.Contacts,
		Tel:             result.Tel,
		Lat:             result.Lat,
		Lng:             result.Lng,
		Address:         result.Address,
		Province:        result.Province,
		City:            result.City,
		Area:            result.Area,
		Logo:            result.Logo,
		CompanyType:     result.CompanyType,
		CompanyStatus:   result.CompanyStatus,
		CompanyFile:     companyFile,
		Shop:            shop,
		Accounts:        account,
	}
	return company
}

//UpdateCompany ...根据企业id修改未审核企业状态
func (c *CompanyService) UpdateCompany(id int64, form *models.CompanyPut) int {
	tx := c.Begin()
	result := &models.Company{}
	db := c.DB.Table("t_company").Where("id = ?", id)
	/*upCompany := &models.Company{
		CompanyStatus:  form.CompanyStatus,
		CompanyNature:  form.CompanyNature,
		AuthVip:        form.AuthVip,
		AuthAgent:      form.AuthAgent,
		AuthAllocation: form.AuthAllocation,
		CompanyProfile: form.CompanyProfile,
	}*/
	if form.CompanyStatus == 1 {
		if err := db.Update(map[string]interface{}{"company_status": form.CompanyStatus, "auth_vip": form.AuthVip, "auth_agent": form.AuthAgent, "auth_allocation": form.AuthAllocation}).Error; err != nil {
			tx.Rollback()
			return e.BAD_REQUEST
		}
		db.First(&result)
		shops := &models.Shops{
			CompanyID: id,
			ShopName:  result.CompanyName,
		}
		if err := c.Table("t_shop").Create(&shops).Error; err != nil {
			tx.Rollback()
			return e.BAD_REQUEST
		}
	} else {
		if err := db.Update("company_status", form.CompanyStatus).Error; err != nil {
			tx.Rollback()
			return e.BAD_REQUEST
		}
	}
	tx.Commit()
	return e.SUCCESS
}
