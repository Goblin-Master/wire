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

	/* 参数错误以及系统错误：10000 ~ 19999 */
	PARAM_NOT_VALID    = MsgCode{10001, "参数无效"}
	PARAM_IS_BLANK     = MsgCode{10002, "参数为空"}
	PARAM_TYPE_ERROR   = MsgCode{10003, "参数类型错误"}
	PARAM_NOT_COMPLETE = MsgCode{10004, "参数缺失"}

	/* 用户错误：20000 ~ 29999 */
	USER_NOT_LOGIN         = MsgCode{20001, "用户未登录"}
	USER_LOGIN_ERROR       = MsgCode{20002, "账号密码不正确"}
	USER_ACCOUNT_FORBIDDEN = MsgCode{20003, "账号已被禁用"}
	USER_NOT_EXIST         = MsgCode{20004, "用户不存在"}
	USER_CREATE_ERROR      = MsgCode{20004, "创建用户失败"}
	USER_UPDATE_ERROR      = MsgCode{20005, "更新用户信息失败"}
	USER_DELETE_ERROR      = MsgCode{20006, "删除用户失败"}
)
