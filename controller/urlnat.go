package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"urlnat/model"
	"urlnat/service"
)

var (
	reNum  = regexp.MustCompile(`^\d+$`)
	reUrl  = regexp.MustCompile(`^(ftp|http|https)://.*$`)
	reDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	rePath = regexp.MustCompile(`^/([^/]*)(.*)$`)
)

func AddUrlNat(ctx *gin.Context) {
	var data model.UrlNat
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid data"})
		return
	}
	if !reUrl.MatchString(data.Url) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid url"})
		return
	}
	if !reDate.MatchString(data.Date) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid date"})
		return
	}
	if rows := service.AddUrlNat(data); rows == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "add fail"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "add success"})
}

func DeleteUrlNat(ctx *gin.Context) {
	var data model.UrlNat
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid data"})
		return
	}
	if rows := service.DeleteUrlNat(data); rows == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "delete fail"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "delete success"})
}

func UpdateUrlNat(ctx *gin.Context) {
	var data model.UrlNat
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid data"})
		return
	}
	if data.Url != "" && !reUrl.MatchString(data.Url) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid url"})
		return
	}
	if data.Date != "" && !reDate.MatchString(data.Date) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid date"})
		return
	}
	if rows := service.UpdateUrlNat(data); rows == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "update fail"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "update success"})
}

func GetUrlNat(ctx *gin.Context) {
	path := ctx.Param("path")
	match := rePath.FindStringSubmatch(path)
	code := match[1]
	param := match[2]
	query := ctx.Request.URL.RawQuery
	data := service.GetUrlNat(code)
	if data.Url == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid code"})
		return
	}
	url := fmt.Sprintf("%s%s", data.Url, param)
	if query != "" {
		url = fmt.Sprintf("%s?%s", url, query)
	}
	ctx.Redirect(http.StatusFound, url)
}

func GetUrlNatList(ctx *gin.Context) {
	pages := ctx.DefaultQuery("page", "1")
	sizes := ctx.DefaultQuery("size", "10")
	query := ctx.DefaultQuery("query", "")
	if !reNum.MatchString(pages) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid page"})
		return
	}
	if !reNum.MatchString(sizes) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "invalid size"})
		return
	}
	data := service.GetUrlNatList(pages, sizes, query)
	var list []model.Response
	for _, item := range data {
		list = append(list, model.Response{
			Code: item.Code,
			Note: item.Note,
			Url:  item.Url,
			Date: item.Date,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "get success", "data": data})
}
