package handler

import (
	"github.com/gin-gonic/gin"

	"GoHttp/util"
)

/**
 * for example
 *
 * @api {post} /api/example
 * @apiName examplef
 * @apiVersion 1.0.0
 * @apiGroup example
 *
 * @apiDescription 这是一个接口示例
 *
 * @apiParam {String} access_token                   access_token
 * @apiParam {String} uid                            UID
 *
 * @apiSuccess (返回值) {String} error_code           返回码
 * @apiSuccess (返回值) {String} error                返回信息
 *
 * @apiSuccessExample {json} 请求示例
 * {
 *     "uid": "UID",
 * }
 *
 * @apiSuccessExample {json} 返回示例
 *  {
 *      "error_code": "0",
 *      "error": "成功",
 *  }
 */
func (*Servlet) Example(c *gin.Context) {
	util.ResponseJSON(c, "ok")
}
