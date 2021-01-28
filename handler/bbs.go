package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"GoHttp/model"
	"GoHttp/util"
)

type (
	createMeowTagReq struct {
		Name        string `json:"name" binding:"required"`
		HostID      string `json:"hostId"`
		Description string `json:"description"`
		Background  string `json:"background"`
	}

	meowTag struct {
		ID          int64  `json:"id,string"`
		Name        string `json:"name"`
		HostID      string `json:"hostId"`
		Description string `json:"description"`
		Background  string `json:"background"`
		TotalCnt    int32  `json:"totalCnt"`
	}

	createMeowTieziReq struct {
		ID      int64  `json:"id,string"`
		Title   string `json:"title" binding:"required"`
		Content string `json:"content"`
		Weight  int32  `json:"weight"`
		HostID  string `json:"hostId"`
		TagID   int64  `json:"tagId,string"`
	}

	getBbsListReq struct {
		PageSize int32 `form:"pageSize" binding:"required,min=1,max=20"`
		Page     int32 `form:"page" binding:"required,min=1"`
		TagID    int64 `form:"tagId" binding:"required,min=1"`
	}

	comment struct {
		ID               int64  `json:"id,string"`
		Content          string `json:"content"`
		CreatedTimestamp int64  `json:"createdTimestamp,string"`
		Nickname         string `json:"nickname"`
		Avatar           string `json:"avatar"`
	}

	tiezi struct {
		ID               int64     `json:"id,string"`
		Title            string    `json:"title"`
		Content          string    `json:"content"`
		Weight           int32     `json:"weight"`
		ReadCnt          int32     `json:"readCnt"`
		CommentCnt       int32     `json:"commentCnt"`
		CreatedTimestamp int64     `json:"createdTimestamp,string"`
		Nickname         string    `json:"nickname"`
		Avatar           string    `json:"avatar"`
		Comments         []comment `json:"comments"`
	}

	getBbsListResp struct {
		Cnt    int64   `json:"cnt,string"`
		Tiezis []tiezi `json:"tiezis"`
	}

	getTieziDetailReq struct {
		TieziID int64 `form:"tieziId" binding:"required,min=1"`
	}

	createCommentReq struct {
		HostID  string `json:"hostId"`
		Content string `json:"content"`
		TieziID int64  `json:"tieziId,string"`
	}
)

/**
 *
 * @api {get} /api/bbs/tags
 * @apiName 论坛标签列表
 * @apiVersion 1.0.0
 * @apiGroup 论坛模块
 *
 * @apiDescription 展示出论坛全部标签，不分页
 *
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  [{
 *      "id": 1,
 *      "name": "test1",
 *      "hostId": "test1",
 *      "description": "test1",
 *      "background": "test1",
 *      "totalCnt": 1
 *  },{
 *      "id": 2,
 *      "name": "test2",
 *      "hostId": "test2",
 *      "description": "test2",
 *      "background": "test2",
 *      "totalCnt": 2
 *  }]
 */
func (*Servlet) GetTags(c *gin.Context) {

	ctx := c.Request.Context()

	meowTags, err := model.MeowTagStatic.FindMeowTags(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tags := make([]meowTag, len(meowTags))
	for i, tag := range meowTags {

		cnt, err := model.MeowTieziStatic.CountByTagID(ctx, tag.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		tags[i] = meowTag{
			ID:          tag.ID,
			Name:        tag.Name,
			HostID:      tag.HostID,
			Description: tag.Description,
			Background:  tag.Background,
			TotalCnt:    int32(cnt),
		}
	}

	util.ResponseJSON(c, tags)
}

/**
 *
 * @api {post} /api/bbs/tag
 * @apiName 论坛标签新增/修改
 * @apiVersion 1.0.0
 * @apiGroup 论坛模块
 *
 * @apiDescription 新增/修改标签
 *
 * @apiParam {String} name                   标签名称（不存在为新增， 已存在为修改）
 * @apiParam {String} hostId                 标签创建者
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 * 		"name": "test",
 * 		"hostId": "xxxx",
 *      "description": "test1",
 *      "background": "test1",
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  HTTP/1.1 200 OK
 *  {
 *		 "error_code": "0",
 *       "error": "成功"
 *  }
 */
func (*Servlet) CreateMeowTag(c *gin.Context) {

	var req createMeowTagReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx := c.Request.Context()

	err = model.MeowTagStatic.InsertOrUpdate(ctx, &model.MeowTag{
		Name:        req.Name,
		HostID:      req.HostID,
		Description: req.Description,
		Background:  req.Background,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseJSON(c, Response{Success: true})
}

/**
 *
 * @api {post} /api/bbs/tiezi
 * @apiName 创建帖子
 * @apiVersion 1.0.0
 * @apiGroup 论坛模块
 *
 * @apiDescription 新增/修改帖子
 *
 * @apiParam {String} id                   主键ID，不传该参数为新增记录
 * @apiParam {String} title                帖子标题
 * @apiParam {String} content              帖子内容
 * @apiParam {String} weight               帖子权重
 * @apiParam {String} tagId                帖子分类
 * @apiParam {String} hostId               帖子创建者
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 * 		"id": "1",   // 不传为新增
 * 		"title": "test",
 * 		"content": "xxxx",
 * 		"weight": 1,
 * 		"tagId": "1",
 * 		"hostId": "xxxx"
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  {
 *		 "error_code": "0",
 *       "error": "成功"
 *  }
 */
func (*Servlet) CreateMeowTiezi(c *gin.Context) {

	var req createMeowTieziReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	err = model.MeowTieziStatic.InsertOrUpdate(ctx, &model.MeowTiezi{
		ID:        req.ID,
		Title:     req.Title,
		Content:   req.Content,
		HostID:    req.HostID,
		Weight:    req.Weight,
		TagID:     req.TagID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseJSON(c, Response{Success: true})
}

/**
 *
 * @api {get} /api/bbs/list
 * @apiName 帖子列表
 * @apiVersion 1.0.0
 * @apiGroup 论坛模块
 *
 * @apiDescription 展示出论坛内容列表，分页
 *
 * @apiParam {String} page                   当前页码
 * @apiParam {String} pageSize               每页大小
 * @apiParam {String} tagId                  标签分类
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 * 		"page": 1,
 * 		"pageSize": 10,
 * 		"tagId": "1",
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  {
 * 		"cnt": 10,
 *      "tieZis": [{
 *      	"id": 1,
 *      	"title": "test1",
 *      	"content": "test1",
 *      	"createdTimestamp": "1234444",
 *      	"weight": 1,
 *      	"hostId": "test1",
 *      	"readCnt": 1,
 *      	"commentCnt": 1
 *      },{
 * 			"id": 2,
 *      	"title": "test2",
 *      	"content": "test2",
 *      	"createdTimestamp": "1234444",
 *      	"weight": 2,
 *      	"hostId": "test1",
 *      	"readCnt": 1,
 *      	"commentCnt": 1
 *      }]
 *  }
 */
func (*Servlet) GetBbsList(c *gin.Context) {

	var req getBbsListReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	meowTiezis, cnt, err := model.MeowTieziStatic.FindMeowTiezis(ctx, req.TagID, int(req.Page), int(req.PageSize))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tiezis := make([]tiezi, len(meowTiezis))
	for i, meowTiezi := range meowTiezis {

		catProfile, err := model.CatProfileStatic.GetByHostID(ctx, meowTiezi.HostID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		var nickname, avatar string
		if catProfile != nil {
			nickname = catProfile.Nickname
			avatar = catProfile.Avatar
		}

		_, totalCnt, err := model.MeowCommentStatic.FindMeowCommentByTieziID(ctx, meowTiezi.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		tiezis[i] = tiezi{
			ID:               meowTiezi.ID,
			Title:            meowTiezi.Title,
			Content:          meowTiezi.Content,
			Weight:           meowTiezi.Weight,
			ReadCnt:          meowTiezi.ReadCount,
			CreatedTimestamp: meowTiezi.CreatedAt.Unix(),
			CommentCnt:       int32(totalCnt),
			Nickname:         nickname,
			Avatar:           avatar,
		}
	}

	util.ResponseJSON(c, &getBbsListResp{
		Cnt:    cnt,
		Tiezis: tiezis,
	})
}

/**
 *
 * @api {get} /api/bbs/tiezi
 * @apiName 帖子详情
 * @apiVersion 1.0.0
 * @apiGroup 论坛模块
 *
 * @apiDescription 展示出论坛内容详情
 *
 * @apiParam {String} tieziId                  帖子ID
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 * 		"tieziId": "1",
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  {
 *     	"id": 1,
 *     	"title": "test1",
 *      "content": "test1",
 *      "createdTimestamp": "1234444",
 *      "weight": 1,
 *      "nickname": "test1",
 *      "avatar": "test1",
 *      "readCnt": 1,
 *      "commentCnt": 1,
 *   	"comments":[{
 *			"id": 1,
 *			"nickname": "test",
 *			"avatar": "test",
 *			"content": "test",
 *			"createTimestamp": "12345"
 *      }]
 *  }
 */
func (*Servlet) GetTieziDetail(c *gin.Context) {

	var req getTieziDetailReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	meowTiezi, err := model.MeowTieziStatic.GetMeowTieziByID(ctx, req.TieziID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	cat, err := model.CatProfileStatic.GetByHostID(ctx, meowTiezi.HostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var tieziNickname, tieziAvatar string
	if cat != nil {
		tieziNickname = cat.Nickname
		tieziAvatar = cat.Avatar
	}

	meowComments, totalCnt, err := model.MeowCommentStatic.FindMeowCommentByTieziID(ctx, req.TieziID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var nickname, avatar string
	comments := make([]comment, len(meowComments))
	for i, co := range meowComments {

		catProfile, err := model.CatProfileStatic.GetByHostID(ctx, co.HostID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if catProfile != nil {
			nickname = catProfile.Nickname
			avatar = catProfile.Avatar
		}

		comments[i] = comment{
			ID:               co.ID,
			Content:          co.Content,
			CreatedTimestamp: co.CreatedAt.Unix(),
			Nickname:         nickname,
			Avatar:           avatar,
		}

	}

	util.ResponseJSON(c, tiezi{
		ID:               meowTiezi.ID,
		Title:            meowTiezi.Title,
		Content:          meowTiezi.Content,
		Weight:           meowTiezi.Weight,
		ReadCnt:          meowTiezi.ReadCount,
		CommentCnt:       int32(totalCnt),
		CreatedTimestamp: meowTiezi.CreatedAt.Unix(),
		Nickname:         tieziNickname,
		Avatar:           tieziAvatar,
		Comments:         comments,
	})
}

/**
 *
 * @api {post} /api/bbs/tiezi
 * @apiName 创建评论
 * @apiVersion 1.0.0
 * @apiGroup 论坛模块
 *
 * @apiDescription 新增评论
 *
 * @apiParam {String} content                评论内容
 * @apiParam {String} hostId                 评论创建者
 * @apiParam {String} tieziId                评论的帖子ID
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 * 		"content": "xxxx",
 * 		"hostId": "xxxx",
 * 		"tieziId": "1"
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  {
 *		 "error_code": "0",
 *       "error": "成功"
 *  }
 */
func (*Servlet) CreateMeowComment(c *gin.Context) {

	var req createCommentReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	err = model.MeowCommentStatic.InsertOrUpdate(ctx, &model.MeowComment{
		Content:   req.Content,
		HostID:    req.HostID,
		TieziID:   req.TieziID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseJSON(c, Response{Success: true})
}

/**
 *
 * @api {get} /api/bbs/like
 * @apiName 帖子点赞
 * @apiVersion 1.0.0
 * @apiGroup 论坛模块
 *
 * @apiDescription 点赞
 *
 * @apiParam {String} tieziId                  帖子ID
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 * 		"tieziId": "1",
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  {
 *		 "error_code": "0",
 *       "error": "成功"
 *  }
 */
func (*Servlet) ClickGood(c *gin.Context) {

	var req getTieziDetailReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	_, err = model.MeowTieziStatic.Increase(ctx, req.TieziID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseJSON(c, Response{Success: true})
}
