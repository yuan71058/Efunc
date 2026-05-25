// 权限管理模块
// 基于 casbin 库实现，支持 RBAC（基于角色的访问控制）和 ABAC（基于属性的访问控制）。
// 提供权限检查、策略管理、角色分配等完整功能。
package utils

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

// Permission_CreateFromFile 从模型文件和策略文件创建权限管理器。
// 模型文件定义权限模型（如 RBAC、ABAC），策略文件存储具体的权限规则。
//
// 常用 RBAC 模型文件内容:
//
//	[request_definition]
//	r = sub, obj, act
//	[policy_definition]
//	p = sub, obj, act
//	[role_definition]
//	g = _, _
//	[policy_effect]
//	e = some(where (p.eft == allow))
//	[matchers]
//	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
//
// 参数:
//   - modelPath: 权限模型文件路径，如 "model.conf"
//   - policyPath: 权限策略文件路径，如 "policy.csv"
//
// 返回:
//   - *casbin.Enforcer: 权限管理器实例
//   - error: 创建失败时返回错误
func Permission_CreateFromFile(modelPath string, policyPath string) (*casbin.Enforcer, error) {
	return casbin.NewEnforcer(modelPath, policyPath)
}

// Permission_CreateFromText 从模型文本和适配器创建权限管理器。
// 无需模型文件，直接传入模型文本字符串。
//
// 参数:
//   - modelText: 权限模型文本内容
//   - adapter: 策略存储适配器，nil 则使用内存适配器
//
// 返回:
//   - *casbin.Enforcer: 权限管理器实例
//   - error: 创建失败时返回错误
func Permission_CreateFromText(modelText string, adapter persist.Adapter) (*casbin.Enforcer, error) {
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		return nil, err
	}
	return casbin.NewEnforcer(m, adapter)
}

// Permission_CreateMemory 从模型文本创建纯内存权限管理器（不持久化）。
// 适用于临时权限检查或测试场景。
//
// 参数:
//   - modelText: 权限模型文本内容
//
// 返回:
//   - *casbin.Enforcer: 权限管理器实例
//   - error: 创建失败时返回错误
func Permission_CreateMemory(modelText string) (*casbin.Enforcer, error) {
	return Permission_CreateFromText(modelText, nil)
}

// Permission_CreateFileAdapter 创建文件适配器，用于策略持久化。
//
// 参数:
//   - policyPath: CSV 格式的策略文件路径
//
// 返回:
//   - *fileadapter.Adapter: 文件适配器实例
func Permission_CreateFileAdapter(policyPath string) *fileadapter.Adapter {
	return fileadapter.NewAdapter(policyPath)
}

// Permission_Enforce 检查主体是否有权限对资源执行操作。
// 基于 RBAC/ABAC 模型进行权限判断。
//
// 参数:
//   - enforcer: 权限管理器实例
//   - sub: 请求主体（用户/角色），如 "alice"
//   - obj: 目标资源，如 "data1"
//   - act: 请求操作，如 "read"
//
// 返回:
//   - bool: true 表示有权限，false 表示无权限
//   - error: 检查失败时返回错误
func Permission_Enforce(enforcer *casbin.Enforcer, sub string, obj string, act string) (bool, error) {
	return enforcer.Enforce(sub, obj, act)
}

// Permission_AddPolicy 添加一条权限策略。
// 添加后需调用 Permission_SavePolicy 持久化到文件。
//
// 参数:
//   - enforcer: 权限管理器实例
//   - sub: 主体标识
//   - obj: 资源标识
//   - act: 操作标识
//
// 返回:
//   - bool: true 表示策略是新添加的，false 表示策略已存在
//   - error: 添加失败时返回错误
func Permission_AddPolicy(enforcer *casbin.Enforcer, sub string, obj string, act string) (bool, error) {
	return enforcer.AddPolicy(sub, obj, act)
}

// Permission_RemovePolicy 删除一条权限策略。
//
// 参数:
//   - enforcer: 权限管理器实例
//   - sub: 主体标识
//   - obj: 资源标识
//   - act: 操作标识
//
// 返回:
//   - bool: true 表示策略被删除，false 表示策略不存在
//   - error: 删除失败时返回错误
func Permission_RemovePolicy(enforcer *casbin.Enforcer, sub string, obj string, act string) (bool, error) {
	return enforcer.RemovePolicy(sub, obj, act)
}

// Permission_AddRoleForUser 为用户分配角色。
// 在 RBAC 模型中使用，将用户绑定到角色。
//
// 参数:
//   - enforcer: 权限管理器实例
//   - user: 用户标识
//   - role: 角色标识
//
// 返回:
//   - bool: true 表示分配成功
//   - error: 分配失败时返回错误
func Permission_AddRoleForUser(enforcer *casbin.Enforcer, user string, role string) (bool, error) {
	return enforcer.AddRoleForUser(user, role)
}

// Permission_DeleteRoleForUser 移除用户的指定角色。
//
// 参数:
//   - enforcer: 权限管理器实例
//   - user: 用户标识
//   - role: 角色标识
//
// 返回:
//   - bool: true 表示移除成功
//   - error: 移除失败时返回错误
func Permission_DeleteRoleForUser(enforcer *casbin.Enforcer, user string, role string) (bool, error) {
	return enforcer.DeleteRoleForUser(user, role)
}

// Permission_GetRolesForUser 获取用户的所有角色列表。
//
// 参数:
//   - enforcer: 权限管理器实例
//   - user: 用户标识
//
// 返回:
//   - []string: 角色名称列表
//   - error: 查询失败时返回错误
func Permission_GetRolesForUser(enforcer *casbin.Enforcer, user string) ([]string, error) {
	return enforcer.GetRolesForUser(user)
}

// Permission_SavePolicy 将当前内存中的策略保存到持久化存储。
// 使用文件适配器时，策略会写入 CSV 文件。
//
// 参数:
//   - enforcer: 权限管理器实例
//
// 返回:
//   - error: 保存失败时返回错误
func Permission_SavePolicy(enforcer *casbin.Enforcer) error {
	return enforcer.SavePolicy()
}

// Permission_LoadPolicy 从持久化存储重新加载策略。
// 用于外部修改了策略文件后刷新内存中的策略。
//
// 参数:
//   - enforcer: 权限管理器实例
//
// 返回:
//   - error: 加载失败时返回错误
func Permission_LoadPolicy(enforcer *casbin.Enforcer) error {
	return enforcer.LoadPolicy()
}