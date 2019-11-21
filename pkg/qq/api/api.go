package api

import (
	"github.com/denyu95/life/pkg/http"
	"github.com/denyu95/life/pkg/setting"
)

var qqApi http.Api

func init() {
	qqApi = http.NewApi(setting.Api.QQBaseUrl)
}

// @title	SendPrivateMsg
// @description	发送私聊消息
// @param	user_id		number	"对方 QQ 号"
// @param	message		message	"要发送的内容"
// @param	auto_escape	boolean "消息内容是否作为纯文本发送（即不解析 CQ 码）"
// @return	message_id	number	"消息 ID"
func SendPrivateMsg(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("send_private_msg", param)
}

// @title	SendGroupMsg
// @description	发送群消息
// @param	group_id	number	"群号"
// @param	message		message	"要发送的内容"
// @param	auto_escape	boolean "消息内容是否作为纯文本发送（即不解析 CQ 码）"
// @return	message_id	number	"消息 ID"
func SendGroupMsg(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("send_group_msg", param)
}

// @title	SendDiscussMsg
// @description	发送讨论组消息
// @param	discuss_id	number	"讨论组 ID（正常情况下看不到，需要从讨论组消息上报的数据中获得）"
// @param	message		message	"要发送的内容"
// @param	auto_escape	boolean "消息内容是否作为纯文本发送（即不解析 CQ 码）"
// @return	message_id	number	"消息 ID"
func SendDiscussMsg(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("send_discuss_msg", param)
}

// @title	SendMsg
// @description	发送消息
// @param	message_type	string	"消息类型，支持 private、group、discuss，分别对应私聊、群组、讨论组，如不传入，则根据传入的 *_id 参数判断"
// @param	user_id			number	"对方 QQ 号"
// @param	group_id		number	"群号"
// @param	discuss_id		number	"讨论组 ID"
// @param	message			message	"要发送的内容"
// @param	auto_escape		boolean "消息内容是否作为纯文本发送（即不解析 CQ 码）"
// @return	message_id		number	"消息 ID"
func SendMsg(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("send_msg", param)
}

// @title	DeleteMsg
// @description	撤回消息
// @param	message_id	number	"消息 ID"
// @return	无
func DeleteMsg(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("delete_msg", param)
}

// @title	SendLike
// @description	发送好友赞
// @param	user_id	number	"对方 QQ 号"
// @param	times	number	"赞的次数，每个好友每天最多 10 次"
// @return	无
func SendLike(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("send_like", param)
}

// @title	SetGroupBan
// @description	群组单人禁言
// @param	group_id	number	"群号"
// @param	user_id		number	"要禁言的 QQ 号"
// @param	duration	number	"禁言时长，单位秒，0 表示取消禁言"
// @return	无
func SetGroupBan(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_group_ban", param)
}

// @title	SetGroupWholeBan
// @description	群组全员禁言
// @param	group_id	number	"群号"
// @param	enable		boolean	"是否禁言"
// @return	无
func SetGroupWholeBan(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_group_whole_ban", param)
}

// @title	SetGroupAdmin
// @description	群组设置管理员
// @param	group_id	number	"群号"
// @param	user_id		number	"要设置管理员的 QQ 号"
// @param	enable		boolean	"true 为设置，false 为取消"
// @return	无
func SetGroupAdmin(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_group_admin", param)
}

// @title	SetGroupAnonymous
// @description	群组匿名
// @param	group_id	number	"群号"
// @param	enable		boolean	"是否允许匿名聊天"
// @return	无
func SetGroupAnonymous(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_group_anonymous", param)
}

// @title	SetGroupCard
// @description	设置群名片（群备注）
// @param	group_id	number	"群号"
// @param	user_id		number	"要设置的 QQ 号"
// @param	card		string	"群名片内容，不填或空字符串表示删除群名片"
// @return	无
func SetGroupCard(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_group_card", param)
}

// @title	SetGroupLeave
// @description	退出群组
// @param	group_id	number	"群号"
// @param	is_dismiss	boolean	"是否解散，如果登录号是群主，则仅在此项为 true 时能够解散"
// @return	无
func SetGroupLeave(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_group_leave", param)
}

// @title	SetGroupSpecialTitle
// @description	设置群组专属头衔
// @param	group_id		number	"群号"
// @param	user_id			number	"要设置的 QQ 号"
// @param	special_title	string	"专属头衔，不填或空字符串表示删除专属头衔"
// @param	duration		number	"专属头衔有效期，单位秒，-1 表示永久，不过此项似乎没有效果，可能是只有某些特殊的时间长度有效，有待测试"
// @return	无
func SetGroupSpecialTitle(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_group_leave", param)
}

// @title	SetDiscussLeave
// @description	退出讨论组
// @param	discuss_id	number	"讨论组 ID（正常情况下看不到，需要从讨论组消息上报的数据中获得）"
// @return	无
func SetDiscussLeave(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_discuss_leave", param)
}

// @title	SetFriendAddRequest
// @description 处理加好友请求
// @param	flag	string	"加好友请求的 flag（需从上报的数据中获得）"
// @param	approve	boolean	"是否同意请求"
// @param	remark	string	"添加后的好友备注（仅在同意时有效）"
// @return	无
func SetFriendAddRequest(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_friend_add_request", param)
}

// @title	SetGroupAddRequest
// @description 处理加群请求／邀请
// @param	flag				string	"加群请求的 flag（需从上报的数据中获得）"
// @param	sub_type 或 type		string	"add 或 invite，请求类型（需要和上报消息中的 sub_type 字段相符）"
// @param	approve				boolean	"是否同意请求／邀请"
// @param	reason				string	"拒绝理由（仅在拒绝时有效）"
// @return	无
func SetGroupAddRequest(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_group_add_request", param)
}

// @title	GetLoginInfo
// @description	获取登录号信息
// @return	user_id		number	"QQ 号"
// @return	nickname	string	"QQ 昵称"
func GetLoginInfo(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_login_info", param)
}

// @title	GetStrangerInfo
// @description	获取陌生人信息
// @param	user_id		number	"QQ 号"
// @param	no_cache	boolean	"是否不使用缓存（使用缓存可能更新不及时，但响应更快）"
// @return	user_id		number	"QQ 号"
// @return	nickname	string	"昵称"
// @return	sex			string	"性别，male 或 female 或 unknown"
// @return	age			number	"年龄"
func GetStrangerInfo(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_stranger_info", param)
}

// @title	GetFriendList
// @description	获取好友列表
// @return	user_id		number	"QQ 号"
// @return	nickname	string	"昵称"
// @return	remark		string	"备注名"
func GetFriendList(param map[string]interface{}) []map[string]interface{} {
	return qqApi.PostReturnList("get_friend_list", param)
}

// @title	GetGroupList
// @description	获取群列表
// @return	group_id	number	"群号"
// @return	group_name	string	"群名称"
func GetGroupList(param map[string]interface{}) []map[string]interface{} {
	return qqApi.PostReturnList("get_group_list", param)
}

// @title	GetGroupInfo
// @description	获取群信息
// @param	group_id			number	"群号"
// @param	no_cache			boolean	"是否不使用缓存（使用缓存可能更新不及时，但响应更快）"
// @return	group_id			number	"群号"
// @return	group_name			string	"群名称"
// @return	member_count		number	"成员数"
// @return	max_member_count	number	"最大成员数（群容量）"
func GetGroupInfo(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_group_info", param)
}

// @title	GetGroupMemberInfo
// @description	获取群成员信息
// @param	group_id			number	"群号"
// @param	user_id				number	"QQ 号"
// @param	no_cache			boolean	"是否不使用缓存（使用缓存可能更新不及时，但响应更快）"
// @return	group_id			number	"群号"
// @return	user_id				number	"QQ 号"
// @return	nickname			string	"昵称"
// @return	card				string	"群名片／备注"
// @return	sex					string	"性别，male 或 female 或 unknown"
// @return	age					number	"年龄"
// @return	area				string	"地区"
// @return	join_time			number	"加群时间戳"
// @return	last_sent_time		number	"最后发言时间戳"
// @return	level				string	"成员等级"
// @return	role				string	"角色，owner 或 admin 或 member"
// @return	unfriendly			boolean	"是否不良记录成员"
// @return	title				string	"专属头衔"
// @return	title_expire_time	number	"专属头衔过期时间戳"
// @return	card_changeable		boolean	"是否允许修改群名片"
func GetGroupMemberInfo(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_group_member_info", param)
}

// @title	GetGroupMemberInfo
// @description	获取群成员信息
// @param	group_id			number	"群号"
// @return	group_id			number	"群号"
// @return	user_id				number	"QQ 号"
// @return	nickname			string	"昵称"
// @return	card				string	"群名片／备注"
// @return	sex					string	"性别，male 或 female 或 unknown"
// @return	age					number	"年龄"
// @return	area				string	"地区"
// @return	join_time			number	"加群时间戳"
// @return	last_sent_time		number	"最后发言时间戳"
// @return	level				string	"成员等级"
// @return	role				string	"角色，owner 或 admin 或 member"
// @return	unfriendly			boolean	"是否不良记录成员"
// @return	title				string	"专属头衔"
// @return	title_expire_time	number	"专属头衔过期时间戳"
// @return	card_changeable		boolean	"是否允许修改群名片"
// @remark	获取列表时和获取单独的成员信息时，某些字段可能有所不同，例如 area、title 等字段在获取列表时无法获得，具体应以单独的成员信息为准。
func GetGroupMemberList(param map[string]interface{}) []map[string]interface{} {
	return qqApi.PostReturnList("get_group_member_list", param)
}

// @title	GetCookies
// @description	获取 Cookies
// @param	domain	string	"需要获取 cookies 的域名"
// @return	cookies	string	"Cookies"
func GetCookies(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_cookies", param)
}

// @title	GetCsrfToken
// @description	获取 CSRF Token
// @return	token	number	"CSRF Token"
func GetCsrfToken(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_csrf_token", param)
}

// @title	GetCredentials
// @description	获取 QQ 相关接口凭证
// @return	cookies		string	"Cookies"
// @return	csrf_token	number	"CSRF Token"
func GetCredentials(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_credentials", param)
}

// @title	GetRecord
// @description	获取语音
// @param	file		string	"收到的语音文件名（CQ 码的 file 参数），如 0B38145AA44505000B38145AA4450500.silk"
// @param	out_format	string	"要转换到的格式，目前支持 mp3、amr、wma、m4a、spx、ogg、wav、flac"
// @param	full_path	boolean	"是否返回文件的绝对路径（Windows 环境下建议使用，Docker 中不建议）"
// @return	file		string	"转换后的语音文件名或路径，如 0B38145AA44505000B38145AA4450500.mp3，如果开启了 full_path，则如 C:\Apps\CoolQ\data\record\0B38145AA44505000B38145AA4450500.mp3"
func GetRecord(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_record", param)
}

// @title	GetImage
// @description	获取图片
// @param	file		string	"收到的图片文件名（CQ 码的 file 参数），如 6B4DE3DFD1BD271E3297859D41C530F5.jpg"
// @return	file		string	"下载后的图片文件路径，如 C:\Apps\CoolQ\data\image\6B4DE3DFD1BD271E3297859D41C530F5.jpg"
func GetImage(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_image", param)
}

// @title	CanSendImage
// @description	检查是否可以发送图片
// @return	yes	boolean	"是或否"
func CanSendImage(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("can_send_image", param)
}

// @title	CanSendRecord
// @description	检查是否可以发送语音
// @return	yes	boolean	"是或否"
func CanSendRecord(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("can_send_record", param)
}

// @title	GetStatus
// @description	获取插件运行状态
// @return	app_initialized	boolean	"HTTP API 插件已初始化"
// @return	app_enabled		boolean	"HTTP API 插件已启用"
// @return	plugins_good	object	"HTTP API 的各内部插件是否正常运行"
// @return	app_good		boolean	"HTTP API 插件正常运行（已初始化、已启用、各内部插件正常运行）"
// @return	online			boolean	"当前 QQ 在线，null 表示无法查询到在线状态"
// @return	good			boolean	"HTTP API 插件状态符合预期，意味着插件已初始化，内部插件都在正常运行，且 QQ 在线"
func GetStatus(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_status", param)
}

// @title	GetVersionInfo
// @description	获取 酷Q 及 HTTP API 插件的版本信息
// @return	coolq_directory				string	"酷Q 根目录路径"
// @return	coolq_edition				string	"酷Q 版本，air 或 pro"
// @return	plugin_version				string	"HTTP API 插件版本，例如 2.1.3"
// @return	plugin_build_number			number	"HTTP API 插件 build 号"
// @return	plugin_build_configuration	string	"HTTP API 插件编译配置，debug 或 release"
func GetVersionInfo(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("get_version_info", param)
}

// @title	SetRestartPlugin
// @description	重启 HTTP API 插件
// @return	delay	number	"要延迟的毫秒数，如果默认情况下无法重启，可以尝试设置延迟为 2000 左右"
func SetRestartPlugin(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("set_restart_plugin", param)
}

// @title	CleanDataDir
// @description	清理数据目录
// @param	data_dir	string	"收到清理的目录名，支持 image、record、show、bface"
func CleanDataDir(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("clean_data_dir", param)
}

// @title	CleanPluginLog
// @description	清理插件日志
func CleanPluginLog(param map[string]interface{}) map[string]interface{} {
	return qqApi.PostReturnMap("clean_plugin_log", param)
}
