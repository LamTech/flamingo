package response

const(
	//	不论怎样，反正错了
	ErrorAnyway = 40001
	//	缺少参数
	RequireParam = 10001
	//	手机号已经存在
	MobileExist = 10002
	//	创建token失败
	CreateTokenError = 10003
	//	json解析出错
	ParseJsonError = 10004
	//	缺少token
	TokenRequired = 10005
	//	token过期了
	TokenExpired = 10006
	//	数据库查询错误
	DbQueryError = 10007
)