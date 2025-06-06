package res

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
)

type RespError struct {
	JsonMsgResult
}

// NewRespError 包装响应错误类型，简化返回信息流程。
func ErrResp(err error, result MsgCode) error {
	respError := &RespError{}
	respError.Code = result.Code
	respError.Message = result.Msg
	respError.Data = err
	return respError
}

func (r RespError) Error() string {
	return r.Message
}

type JsonMsgResponse struct {
	Ctx *gin.Context
}

type JsonMsgResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type nilStruct struct{}

const SUCCESS_CODE = 20000
const SUCCESS_MSG = "成功"
const ERROR_MSG = "错误"

const code200 = 200

// Response 更加通用的返回方法 以后不用直接使用gin的返回方法
func Response(c *gin.Context, data interface{}, err error) {
	if err != nil {
		// 如果出现错误，判断是否是RespError类型
		respErr := &RespError{}
		// 判断响应是否含有RespError ，如果有则返回错误信息
		if ok := errors.As(err, &respErr); ok {
			c.JSON(code200, JsonMsgResult{
				Code:    respErr.Code,
				Message: respErr.Message,
				Data:    nil,
			})
			return
		} else {
			// 更加通用的类型错误返回
			c.JSON(code200, JsonMsgResult{
				Code:    COMMON_FAIL.Code,
				Message: COMMON_FAIL.Msg,
				Data:    err.Error(),
			})
			return
		}
	} else {
		// 正常返回
		c.JSON(code200, JsonMsgResult{
			Code:    SUCCESS_CODE,
			Message: SUCCESS_MSG,
			Data:    data,
		})
	}
}

func NewResponse(c *gin.Context) *JsonMsgResponse {
	return &JsonMsgResponse{Ctx: c}
}

func (r *JsonMsgResponse) Success(data interface{}) {
	res := JsonMsgResult{}
	res.Code = SUCCESS_CODE
	res.Message = SUCCESS_MSG
	res.Data = data
	r.Ctx.JSON(code200, res)
}

func (r *JsonMsgResponse) Error(mc MsgCode) {
	r.error(mc.Code, mc.Msg)
}

func (r *JsonMsgResponse) error(code int, message string) {
	if message == "" {
		message = ERROR_MSG
	}
	res := JsonMsgResult{}
	res.Code = code
	res.Message = message
	res.Data = nilStruct{}
	r.Ctx.JSON(code200, res)
}

func SSESuccess(data any, c *gin.Context) {
	byteData, _ := json.Marshal(gin.H{"code": SUCCESS_CODE, "data": data, "msg": "成功"})
	c.SSEvent("", string(byteData))
	c.Writer.Flush()
}
func SSEFail(msg string, c *gin.Context) {
	byteData, _ := json.Marshal(gin.H{"code": COMMON_FAIL.Code, "data": map[string]any{}, "msg": msg})
	c.SSEvent("", string(byteData))
	c.Writer.Flush()
}
