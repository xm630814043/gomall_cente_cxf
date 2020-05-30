package service

import (
	"gomall-center/models"
	"gomall-center/pkg/enum"
	"gomall-center/pkg/web"
)

// CategoryService 产品所有分类业务
type CategoryService struct {
	Service
}

//NewCategoryService ...
func NewCategoryService(ctx *web.RequestContext) *CategoryService {
	return &CategoryService{
		Service: InitService(ctx),
	}
}

// CategoryList ...获取所有种类
func (c *CategoryService) CategoryList() *models.CategoryNode {
	categoryList := make([]*models.CategoryNode, 0)
	list := make([]*models.Category, 0)
	c.Where("category_status = ?", enum.CATEGORY_ON_LINE).Find(&list)
	categoryList = buildTree(0, list)
	result := &models.CategoryNode{
		ID:           0,
		CategoryName: "所有种类",
		Children:     categoryList,
		ParentID:     -1,
	}
	return result
}

//buildTree 树
func buildTree(parentID int, data []*models.Category) []*models.CategoryNode {
	result := make([]*models.CategoryNode, 0)
	for _, node := range data {
		if node.ParentID == parentID {
			item := &models.CategoryNode{
				ID:           node.ID,
				ParentID:     node.ParentID,
				CategoryName: node.CategoryName,
				Icon:         node.Icon,
				Level:        node.Level,
				Children:     buildTree(int(node.ID), data),
			}
			result = append(result, item)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

//PropList ...所有种类的属性
func (c *CategoryService) PropList(args *models.PagerArgs) *models.PagerData {
	var result []*models.ProductProp
	var count int
	db := c.DB
	db = db.Where("id IN(SELECT prop_id FROM t_category_prop) ")
	db.Table("t_product_prop").Count(&count)
	db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&result)
	pagerData := &models.PagerData{
		Count: int(count),
		Data:  result,
	}
	return pagerData
}

//GetPropList ...根据种类ID获取属性
func (c *CategoryService) GetPropList(categoryID int, parentId int, args *models.PagerArgs) *models.PagerData {
	var result []*models.ProductProp
	var count int
	db := c.DB
	db = db.Table("t_category_prop c").Select(" distinct p.id,p.prop_name,p.prop_type,p.input_type,t_category.parent_id").Joins("INNER JOIN t_category ON c.category_id=t_category.id INNER JOIN t_product_prop p ON p.id=c.prop_id")
	db = db.Where("FIND_IN_SET(?,t_category.path) OR t_category.id=? OR c.category_id=0", categoryID, parentId)
	db.Group("p.prop_name").Count(&count)
	db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&result)
	pagerDate := &models.PagerData{
		Count: int(count),
		Data:  result,
	}
	return pagerDate
}
