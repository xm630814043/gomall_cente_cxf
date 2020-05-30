package api

import (
	"fmt"
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

// FindChannelList ...获取频道专题
func FindChannelList(ctx *web.Context) {
	srv := service.NewSubjectService(ctx.RequestContext)
	data := srv.ChannelList()
	ctx.ResponseData(e.SUCCESS, data)
}

// FindBySubjectList ...获取频道专题内容列表
func FindBySubjectList(ctx *web.Context) {
	subjectID, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	contentType := ctx.Query("contentType")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	srv := service.NewSubjectService(ctx.RequestContext)
	data := srv.SubjectListByID(subjectID, contentType, limit)
	ctx.ResponseData(e.SUCCESS, data)
}

// FindByUnSubjectList ...获取可以添加到频道专题的内容列表
func FindByUnSubjectList(ctx *web.Context) {
	subjectID, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	contentType := ctx.Query("contentType")
	keyWord := ctx.Query("keyWord")
	shopInput := ctx.Query("shopinput")
	brandInput := ctx.Query("brandinput")
	radio, _ := strconv.Atoi(ctx.Query("radio"))
	count, _ := strconv.Atoi(ctx.Query("count"))
	srv := service.NewSubjectService(ctx.RequestContext)
	data := srv.UnSubjectListByID(subjectID, contentType, keyWord, shopInput, brandInput, radio, count)
	ctx.ResponseData(e.SUCCESS, data)
}

// RemoveSubject ...删除专题对应产品关系
func RemoveSubject(ctx *web.Context) {
	objectID, _ := strconv.Atoi(ctx.Param("id"))
	subjectContent := &models.SubjectContentRelation{}
	_ = ctx.ShouldBindQuery(subjectContent)
	srv := service.NewSubjectService(ctx.RequestContext)
	code := srv.DeleteByID(objectID, subjectContent)
	ctx.Response(code)
}

// AddSubject ...添加专题对应产品关系
func AddSubject(ctx *web.Context) {
	form := &models.SubjectContentRelation{}
	if err := ctx.BindJSON(form); err != nil {
		fmt.Println("form", form)
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewSubjectService(ctx.RequestContext)
	code := srv.InsertSubject(form)
	ctx.Response(code)
}

//AddArticle 添加轮播图
func AddArticle(ctx *web.Context) {
	form := &models.Article{}
	ctx.BindJSON(&form)
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	form.SubjectID = 4
	srv := service.NewSubjectService(ctx.RequestContext)
	code := srv.InsertArticle(form)
	ctx.Response(code)
}

// FindBySubject  ...获取频道专题内容详情
func FindBySubject(ctx *web.Context) {
	objectID, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	contentType := ctx.Query("contentType")
	srv := service.NewSubjectService(ctx.RequestContext)
	data := srv.SubjectByID(objectID, contentType)
	ctx.ResponseData(e.SUCCESS, data)
}

//FindByName  获取模糊查询的店铺
func FindByName(ctx *web.Context) {
	like := ctx.Query("value")
	srv := service.NewSubjectService(ctx.RequestContext)
	srv.ShopListByLike(like)
}
