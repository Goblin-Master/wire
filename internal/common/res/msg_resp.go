package res

type MsgCode struct {
	Code int
	Msg  string
}

var (
	/* 成功 */
	SUCCESS = MsgCode{Code: 20000, Msg: "成功"}

	/* 默认失败 */
	COMMON_FAIL = MsgCode{-43960, "失败"}

	/* 参数错误：10000 ~ 19999 */
	PARAM_NOT_VALID    = MsgCode{10001, "参数无效"}
	PARAM_IS_BLANK     = MsgCode{10002, "参数为空"}
	PARAM_TYPE_ERROR   = MsgCode{10003, "参数类型错误"}
	PARAM_NOT_COMPLETE = MsgCode{10004, "参数缺失"}
	MEMBER_NOT_EXIST   = MsgCode{10005, "用户不存在"}
)
