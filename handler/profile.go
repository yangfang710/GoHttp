package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"GoHttp/model"
	"GoHttp/util"
)

type Response struct {
	Success bool `json:"success"`
}

/**
 * 喵星人入驻
 *
 * @api {post} /api/cat/profile
 * @apiName 创建喵星人
 * @apiVersion 1.0.0
 * @apiGroup cat profile
 *
 * @apiDescription 通过本接口创建喵星人
 *
 * @apiParam {String}  host_id                     主人标识(16位随机字符串)
 * @apiParam {Boolean} [is_worked]                 是否有工作经历
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 *     "hid": "abc",
 *     "is_worked": true
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  {
 *      "error_code": "0",
 *      "error": "成功",
 *  }
 */

type addCatProfileReq struct {
	HostID   string `json:"host_id" binding:"required"`
	IsWorked bool   `json:"is_worked"`
}

func (*Servlet) AddCatProfile(c *gin.Context) {

	var req addCatProfileReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if err := model.CatProfileStatic.Insert(c.Request.Context(), &model.CatProfile{
		HostID:   req.HostID,
		Nickname: fmt.Sprintf("🐱-%s", util.RandString(8, util.LetterBytes)),
		Avatar:   fmt.Sprintf("avatar_%d", r.Intn(10)),
		MeowID:   fmt.Sprintf("meow-%s", util.RandString(11, util.NumberLetterBytes)),
		Level:    1,
		Money:    666.66,
		IsWorked: req.IsWorked,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseJSON(c, Response{Success: true})
}

/**
 * 喵星人入驻
 *
 * @api {get} /api/cat/profile
 * @apiName 获取喵星人信息
 * @apiVersion 1.0.0
 * @apiGroup cat profile
 *
 * @apiDescription 通过本接口获取喵星人信息
 *
 * @apiParam {String}  host_id                       主人标识(16位随机字符串)
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 返回示例
 *	{
 *	    "id": 1,
 *	    "host_id": "abcd",
 *	    "nickname": "🐱",
 *	    "avatar": "qq",
 *	    "meow_id": "meowid",
 *	    "level": 1,
 *	    "money": 999.99,
 *	    "is_worked": true,
 *	    "created_at": "2020-12-10T11:43:54+08:00",
 *	    "updated_at": "2020-12-10T11:43:54+08:00"
 *	}
 */

type getCatProfileReq struct {
	HostID string `form:"host_id" json:"host_id" binding:"required"`
}

func (*Servlet) GetCatProfile(c *gin.Context) {

	var req getCatProfileReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	catProfile, err := model.CatProfileStatic.GetByHostID(c.Request.Context(), req.HostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if catProfile == nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseJSON(c, catProfile)
}
