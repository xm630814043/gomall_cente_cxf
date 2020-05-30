package service

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
)

//定义栏目主题
const (
	SubjectProduct = "product"
	SubjectShop    = "shop"
	SubjectArticle = "article"
)

//SubjectService ...
type SubjectService struct {
	Service
}

//NewSubjectService ...
func NewSubjectService(context *web.RequestContext) *SubjectService {
	c := &SubjectService{InitService(context)}
	return c
}

// ChannelList ...获取频道专题列表
func (c *SubjectService) ChannelList() []*models.Channel {
	var channels []*models.Channel
	var subjects []*models.Subject
	c.Find(&channels)
	c.Find(&subjects)
	for _, c := range channels {
		for _, s := range subjects {
			if int64(c.ID) == s.ChannelID {
				c.Subjects = append(c.Subjects, s)
			}
		}
	}
	return channels
}

//UnSubjectListByID 根据专题Id 获取不在专题的商品、店铺或文章
func (c *SubjectService) UnSubjectListByID(subjectID int64, contentType string, keyWord string, shopInput string, brandInput string, radio int, limit int) interface{} {
	var count int
	db := c.Where("t_"+contentType+".id not in (select object_id from t_subject_content_relation where subject_id = ? )", subjectID)
	if contentType == SubjectProduct {
		var products []*models.Product
		//根据店铺名称，品牌名称，商品名称筛选查询商品
		db = db.Select("t_product.*,f.favourite")
		db = db.Joins("LEFT JOIN (SELECT COUNT(*) favourite, object_id FROM t_favorites WHERE object_type=1 GROUP BY object_id )f ON id =f.object_id")

		db = db.Where("product_name like ? AND publish_status=2 AND id IN ( SELECT p.id FROM t_product p WHERE p.shop_id IN(SELECT id FROM t_shop WHERE shop_name LIKE ? AND shop_status=1)) AND brand_name LIKE ?", "%"+keyWord+"%", "%"+shopInput+"%", "%"+brandInput+"%")
		if radio == 3 {
			db = db.Order("id ")
		} else if radio == 6 {
			db = db.Order("click_num desc")
		} else if radio == 9 {
			db = db.Order("sell_num desc")
		} else {
			db = db.Order("f.favourite  desc")
		}
		db.Model(&models.Product{}).Count(&count)
		db.Limit(limit).Find(&products)
		loaddata := &models.PagerData{
			Count: int(count),
			Data:  products,
		}
		return loaddata
	} else if contentType == SubjectShop {
		var shops []*models.Shop
		db = db.Select("t_shop.*,f.favourite,c.logo")
		db = db.Joins("LEFT JOIN (SELECT COUNT(*) favourite, object_id FROM t_favorites WHERE object_type=2 GROUP BY object_id )f ON t_shop.id =f.object_id LEFT JOIN (SELECT t_company.id,t_company.logo FROM t_company )c ON t_shop.company_id=c.id")
		db = db.Where("shop_status=1 AND shop_name like ?", "%"+keyWord+"%")
		if radio == 3 {
			db = db.Order("id ")
		} else if radio == 6 {
			db = db.Order("hot desc")
		} else {
			db = db.Order("f.favourite desc")
		}
		db.Model(&models.Shop{}).Count(&count)
		db.Limit(limit).Find(&shops)
		loaddata := &models.PagerData{
			Count: int(count),
			Data:  shops,
		}
		return loaddata
	} else if contentType == SubjectArticle {
		var articles []*models.Article
		db.Where("title like ?", "%"+keyWord+"%").Find(&articles)
		return articles
	}
	return nil
}

//SubjectListByID ...根据频道id,专题id,获取频道专题的列表
func (c *SubjectService) SubjectListByID(subjectID int64, contentType string, limit int) interface{} {
	db := c.Where("t_"+contentType+".id in (select object_id from t_subject_content_relation where subject_id = ? )", subjectID)
	if contentType == SubjectProduct {
		var products []*models.Product
		db = db.Select("t_product.*,f.favourite")
		db = db.Joins("LEFT JOIN (SELECT COUNT(*) favourite, object_id FROM t_favorites WHERE object_type=1 GROUP BY object_id )f ON id =f.object_id")
		db.Find(&products)
		return products
	} else if contentType == SubjectShop {
		//var companyID []int64
		var shop []*models.Shop
		db = db.Select("t_shop.*,f.favourite,c.logo")
		db = db.Joins("LEFT JOIN (SELECT COUNT(*) favourite, object_id FROM t_favorites WHERE object_type=2 GROUP BY object_id )f ON t_shop.id =f.object_id LEFT JOIN (SELECT t_company.id,t_company.logo FROM t_company )c ON t_shop.company_id=c.id")
		db.Find(&shop)
		return shop
	} else if contentType == SubjectArticle {
		var articles []*models.Article
		db.Find(&articles)
		return articles
	}
	return nil
}

//DeleteByID ...	根据专题id,专题内容id,删除专题对应产品关系
func (c *SubjectService) DeleteByID(objectID int, subjects *models.SubjectContentRelation) int {
	tx := c.Begin()
	if err := c.Where("object_id = ? and subject_id = ?", objectID, subjects.SubjectID).Delete(&models.SubjectContentRelation{}).Error; err != nil {
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//InsertSubject ...	根据专题id,专题内容id,添加专题对应产品关系
func (c *SubjectService) InsertSubject(subject *models.SubjectContentRelation) int {
	var subjectContent = models.SubjectContentRelation{SubjectID: subject.SubjectID, ObjectID: subject.ObjectID}
	if err := c.Table("t_subject_content_relation").Create(&subjectContent).Error; err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

//InsertArticle ...创建轮播图
func (c *SubjectService) InsertArticle(article *models.Article) int {
	tx := c.Begin()
	if err := c.Create(&article).Error; err != nil {
		tx.Rollback()
		return e.ERROR
	}
	c.Last(&article)
	subjectContent := &models.SubjectContentRelation{
		ID:        0,
		SubjectID: 4,
		ObjectID:  int64(article.ID),
	}
	if err := c.Create(&subjectContent).Error; err != nil {
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//SubjectByID ...根据专题内容id,获取专题内容详情
func (c *SubjectService) SubjectByID(objectID int64, contentType string) interface{} {
	if contentType == SubjectProduct {
		products := &models.Product{}
		c.Where("id = ? ", objectID).Find(&products)
		return products
	} else if contentType == SubjectShop {
		shops := &models.Shop{}
		c.Where("id = ? ", objectID).Find(&shops)
		return shops
	} else if contentType == SubjectArticle {
		articles := &models.Article{}
		c.Where("id = ? ", objectID).Find(&articles)
		return articles
	}
	return nil
}

//ShopListByLike 模糊查询提示
func (c *SubjectService) ShopListByLike(value string) []*models.Shop {
	var shops []*models.Shop
	c.Where("shop_name like ? AND shop_status=1", "%"+value+"%").Limit(5).Find(&shops)
	return shops
}
