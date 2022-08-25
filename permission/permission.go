package permission

import "strconv"

type Permission uint64

const (
	ManageChannels  Permission = 1 << 0  // 管理频道与分组
	EditChannels    Permission = 1 << 1  // 编辑频道
	ManageMembers   Permission = 1 << 2  // 管理成员
	Administrator   Permission = 1 << 3  // 超级管理员
	ModifyNickname  Permission = 1 << 4  // 修改群昵称
	ManageNickname  Permission = 1 << 5  // 管理群昵称
	ViewChannel     Permission = 1 << 6  // 查看频道
	ManageRoles     Permission = 1 << 7  // 管理权限与身份组
	ManageEmojis    Permission = 1 << 8  // 管理群表情包
	Mention         Permission = 1 << 9  // @所有人和身份组
	SendMessages    Permission = 1 << 10 // 发送消息
	ManageMessages  Permission = 1 << 11 // 管理消息
	CreateReaction  Permission = 1 << 12 // 添加新反应
	PublishArticles Permission = 1 << 13 // 发布帖子
	ManageArticles  Permission = 1 << 14 // 管理帖子
	DeleteArticles  Permission = 1 << 15 // 删除帖子
	Connect         Permission = 1 << 16 // 连接
	Speak           Permission = 1 << 17 // 说话
	ManageVoices    Permission = 1 << 18 // 管理语音
	MoveVoiceMember Permission = 1 << 19 // 移动成员加入语音频道
	SearchArticle   Permission = 1 << 20 // 搜索帖子权限
	CommentArticle  Permission = 1 << 21 // 帖子评论权限
	ManageRes       Permission = 1 << 22 // 资料管理权限
)

// CalculatePermission calculates a bunch of permissions into a Permission object.
//
// 将传入的一组 Permission 计算成一个结果
func CalculatePermission(permissions ...Permission) Permission {
	var perm Permission
	if len(permissions) == 0 {
		return perm
	}
	for _, permission := range permissions {
		perm |= permission
	}
	return perm
}

// CheckPermissionExist checks the permission that contains target Permission or not.
//
// 检查 target 是否存在于 permission 中
func CheckPermissionExist(permission Permission, target Permission) bool {
	return (permission & target) == target
}

// ParseToHexadecimalString converts a Permission object into a hexadecimal string,
// the result which you can use to build the request like model.CreateRoleReq or model.EditRoleReq, with `Permission` parameter.
//
// 转换 Permission 类型的权限值到 16 进制的字符串
// 可以在构建 model.CreateRoleReq#Permission 或 model.EditRoleReq#Permission 时使用
func ParseToHexadecimalString(permission Permission) string {
	hex := strconv.FormatUint(uint64(permission), 16)
	return hex
}

// ParseHexadecimalStringToPermission converts the hexadecimal string into a Permission object,
// usually you can use this function to handle the `Permission` value response by server in model.RoleElement.
// When the permString received an invalid hexadecimal string, an error will be occurred.
//
// 转换字符串类型的 16 进制权限值到 Permission 类型的权限值
// 通常在拿到服务器返回的 model.RoleElement#Permission 数据后校验权限时会用上
// 当收到不正确的 16 进制字符串时，会返回 0 值，所以需要仔细处理 error
func ParseHexadecimalStringToPermission(permString string) (Permission, error) {
	permission, err := strconv.ParseUint(permString, 16, 64)
	if err != nil {
		return 0, err
	}
	return Permission(permission), nil
}
