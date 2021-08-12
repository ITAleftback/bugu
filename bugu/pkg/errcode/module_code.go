/**
 * @Author: Anpw
 * @Description:
 * @File:  module_code
 * @Version: 1.0.0
 * @Date: 2021/5/27 22:55
 */

package errcode



var (
	ErrorCreateUserFail = NewError(20010001, "创建用户失败")
	ErrorLoginUserFail = NewError(20010002,"用户登录失败")
	ErrorUpdateCodeFail = NewError(20010003,"修改密码失败")

	ErrorCreateChipFail = NewError(20020001, "创建芯片失败")
	ErrorUpdateChipFail = NewError(20020002, "更新芯片失败")
	ErrorDeleteChipFail = NewError(20020003, "删除芯片失败")
	ErrorGetChipFail = NewError(20020004, "获取芯片失败")
	ErrorGetChipsFail = NewError(20020005,"获取芯片列表失败")

	ErrorCreatePluginFail = NewError(20030001, "创建插件失败")
	ErrorUpdatePluginFail = NewError(20030002, "更新插件失败")
	ErrorDeletePluginFail = NewError(20030003, "删除插件失败")
	ErrorGetPluginFail = NewError(20030004, "获取插件失败")
	ErrorGetPluginsFail = NewError(20320005,"获取插件列表失败")

	ErrorCreateModuleFail = NewError(20040001, "创建模块失败")
	ErrorUpdateModuleFail = NewError(20040002, "更新模块失败")
	ErrorDeleteModuleFail = NewError(20040003, "删除模块失败")
	ErrorGetModuleFail = NewError(20040004, "获取模块失败")
	ErrorGetModulesFail = NewError(20040005,"获取模块列表失败")

	ErrorCreateOverloadFail = NewError(20050001, "创建重载失败")
	ErrorUpdateOverloadFail = NewError(20050002, "更新重载失败")
	ErrorDeleteOverloadFail = NewError(20050003, "删除重载失败")
	ErrorGetOverloadFail = NewError(20050004, "获取重载失败")
	ErrorGetOverloadsFail = NewError(20050005,"获取重载列表失败")
)
