package utils

const (
	SUCCESS         = 200
	SUCCESS_MESSAGE = "操作成功"

	PRODUCT_OPERATING         = 201
	PRODUCT_OPERATING_MESSAGE = "产品正在进行操作"

	LOGIN         = 307
	LOGIN_MESSAGE = "重新登录"

	REQUEST_PARAMETERS         = 400
	REQUEST_PARAMETERS_MESSAGE = "请求参数不能为空"

	PARAMETERS_TYPE_ERROR         = 401
	PARAMETERS_TYPE_ERROR_MESSAGE = "参数类型不合法"

	FINANCIAL_NO_POWER         = 402
	FINANCIAL_NO_POWER_MESSAGE = "无财务权限"

	NO_CUSTOMER         = 403
	NO_CUSTOMER_MESSAGE = "此客户不存在"

	ERROR         = 500
	ERROR_MESSAGE = "服务器异常"

	RETURN_DATA_NULL         = 501
	RETURN_DATA_NULL_MESSAGE = "返回值为空"

	FAIL         = 502
	FAIL_MESSAGE = "操作失败"

	TOKEN_IS_NULL         = 600
	TOKEN_IS_NULL_MESSAGE = "获取token为空"

	//########################注册登录####################################
	REGISTER_SUCCESS         = 1000
	REGISTER_SUCCESS_MESSAGE = "注册成功"

	REGISTER_FAIL         = 1001
	REGISTER_FAIL_MESSAGE = "注册失败"

	NAME_NOT_NIL         = 1002
	NAME_NOT_NIL_MESSAGE = "用户名不合法"

	NAME_REGISTERD         = 1002
	NAME_REGISTERD_MESSAGE = "用户名已被注册"

	EMAIL_NOT_NIL         = 1003
	EMAIL_NOT_NIL_MESSAGE = "邮箱不合法"

	EMAIL_REGISTERD         = 1002
	EMAIL_REGISTERD_MESSAGE = "邮箱已被注册"

	PHONE_NOT_NIL         = 1004
	PHONE_NOT_NIL_MESSAGE = "手机号不合法"

	PHONE_REGISTERD         = 1002
	PHONE_REGISTERD_MESSAGE = "手机号已被注册"

	PASSWORD_NOT_NIL         = 1004
	PASSWORD_NOT_NIL_MESSAGE = "密码不合法"

	CONFIRMPASSWORD_NOT_NIL         = 1005
	CONFIRMPASSWORD_NOT_NIL_MESSAGE = "确认密码不合法"

	PASSWORD_NOT_SAME         = 1006
	PASSWORD_NOT_SAME_MESSAGE = "密码不相同"

	PHONECODE_NOT_NIL         = 1008
	PHONECODE_NOT_NIL_MESSAGE = "手机验证码不能为空"

	VERIFICATION_CODE_FAIL=1009
	VERIFICATION_CODE_FAIL_MESSAGE="验证码失败"

	VERIFICATION_CODE_NOT_SAME=1010
	VERIFICATION_CODE_NOT_SAME_MESSAGE="验证码不匹配"

	PASSWORD_MD5_FAIL=1011
	PASSWORD_MD5_FAIL_MESSAGE="处理客户信息异常"

	RESEND_CODE=1012
	RESEND_CODE_MESSAGE="请重新发送验证码"

	LOGIN_FAIL=1013
	LOGIN_FAIL_MESSAGE="登录失败"
	//############################################################


	//#################获取前端json数据错误#########################
	REQUESTBODY_FAIL         = 1009
	REQUESTBODY_FAIL_MESSAGE = "请求体获取失败"

	//网络提示信息（1300-1400）
	PRIVATE_NNETWORK_EXIST         = 1300
	PRIVATE_NNETWORK_EXIST_MESSAGE = "此名字已存在"

	PRIVATE_NNETWORK_CREATE_FAIL         = 1301
	PRIVATE_NNETWORK_CREATE_FAIL_MESSAGE = "创建私有网络失败"

	VPC_CIDR_ILLEGALITY         = 1302
	VPC_CIDR_ILLEGALITY_MESSAGE = "VPC网段不合法"

	PRIVATE_NNETWORK_DELETE_FAIL         = 1303
	PRIVATE_NNETWORK_DELETE_FAIL_MESSAGE = "删除失败,已绑定云主机"

	PRIVATE_NNETWORK_HAVE_PORT         = 1304
	PRIVATE_NNETWORK_HAVE_PORT_MESSAGE = "该私网有内网ip，不可删除"

	PRIVATE_NNETWORK_NO_EXIST         = 1305
	PRIVATE_NNETWORK_NO_EXIST_MESSAGE = "该私网不存在"

	GET_PORTID_IP_FAIL         = 1306
	GET_PORTID_IP_FAIL_MESSAGE = "获取port和ip失败"

	RETURN_NULL         = 1307
	RETURN_NULL_MESSAGE = "调用创建vpc返回null"

	//项目成员（1400-1500）
	CREATE_PROJECT_ERROR         = 1400
	CREATE_PROJECT_ERROR_MESSAGE = "创建项目失败"

	CREATE_PROJECT_LIMIT         = 1401
	CREATE_PROJECT_LIMIT_MESSAGE = "项目已经达到上限，无法继续创建"

	EDITE_PROJECT_FAIL         = 1402
	EDITE_PROJECT_FAIL_MESSAGE = "项目修改失败"

	DELETE_PROJECT_FAIL         = 1403
	DELETE_PROJECT_FAIL_MESSAGE = "项目包含成员，请删除成员"

	DELETE_PROJECT_FAIL1         = 1404
	DELETE_PROJECT_FAIL1_MESSAGE = "项目包含产品，请删除产品"

	NEED_ADD_MEMBER         = 1405
	NEED_ADD_MEMBER_MESSAGE = "当前项目下没有可选成员，是否新建成员"

	USER_EXIT         = 1406
	USER_EXIT_MESSAGE = "用户已经存在，请重新输入"

	USER_ADD_SUCCESS         = 1407
	USER_ADD_SUCCESS_MESSAGE = "添加成员成功"

	USER_ADD_FAIL         = 1408
	USER_ADD_FAIL_MESSAGE = "添加成员失败"

	REMOVE_MEMBER_FAIL         = 1409
	REMOVE_MEMBER_FAIL_MESSAGE = "移除成员失败"

	NAME_EXIT         = 1410
	NAME_EXIT_MESSAGE = "用户名已经存在，不允许重复"

	MAIL_EXIT         = 1411
	MAIL_EXIT_MESSAGE = "邮箱已经存在，不允许重复"

	PHONE_EXIT         = 1412
	PHONE_EXIT_MESSAGE = "手机号已经存在，不允许重复"

	CREATE_MEM_SUCCESS         = 1413
	CREATE_MEM_SUCCESS_MESSAGE = "新建成员成功"

	CREATE_MEM_FAIL         = 1414
	CREATE_MEM_FAIL_MESSAGE = "新建成员失败"

	PROJECT_NOT_EXIT         = 1415
	PROJECT_NOT_EXIT_MESSAGE = "项目不存在"

	MEMBER_AND_PROJECT         = 1416
	MEMBER_AND_PROJECT_MESSAGE = "该成员不属于此项目"

	CUSTOMER_NOW_USING         = 1417
	CUSTOMER_NOW_USING_MESSAGE = "该用户登录状态正常,无需恢复"

	FAIL_TO_DISTRIBUTION         = 1418
	FAIL_TO_DISTRIBUTION_MESSAGE = "默认分配项目配额失败"

	DELETE_PROJECTS_FAIL         = 1419
	DELETE_PROJECTS_FAIL_MESSAGE = "删除项目失败"

	DELETE_DEFAULT_FAIL         = 1420
	DELETE_DEFAULT_FAIL_MESSAGE = "默认项目不允许删除"

	MEMBER_IN_PROJECT         = 1421
	MEMBER_IN_PROJECT_MESSAGE = "成员已添加,禁止重复添加"

	PASSWORD_NOT_EQUAL         = 1422
	PASSWORD_NOT_EQUAL_MESSAGE = "两次密码不统一"

	REMOVE_FAIL         = 1423
	REMOVE_FAIL_MESSAGE = "当前成员已分配项目,请在项目中移除"

	MEMBER_NOT_EXIT         = 1424
	MEMBER_NOT_EXIT_MESSAGE = "成员不存在"

	CREATE_LOANDBALANCE_FAIL         = 1425
	CREATE_LOANDBALANCE_FAIL_MESSAGE = "添加负载均衡的配额失败"

	DELETE_LOANDBALANCE_FAIL         = 1426
	DELETE_LOANDBALANCE_FAIL_MESSAGE = "删除负载均衡的配额失败"

	DELETE_DMEMBER_FAIL         = 1427
	DELETE_DMEMBER_FAIL_MESSAGE = "调用docker删除成员失败"

	DELETE_DOCKERMEMBER_FAIL         = 1428
	DELETE_DOCKERMEMBER_FAIL_MESSAGE = "该成员下含有容器产品无法删除"

	EDITE_PROJECT_SUCCESS         = 1429
	EDITE_PROJECT_SUCCESS_MESSAGE = "项目修改成功"
)
