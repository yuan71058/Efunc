package utils

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

// Q权限_从文件创建 从模型文件和策略文件创建权限管理器。
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
//   - 模型文件路径: 权限模型文件路径，如 "model.conf"
//   - 策略文件路径: 权限策略文件路径，如 "policy.csv"
//
// 返回:
//   - *casbin.Enforcer: 权限管理器实例
//   - error: 创建失败时返回错误
func Q权限_从文件创建(模型文件路径 string, 策略文件路径 string) (*casbin.Enforcer, error) {
	return casbin.NewEnforcer(模型文件路径, 策略文件路径)
}

// Q权限_从文本创建 从模型文本和适配器创建权限管理器。
// 无需模型文件，直接传入模型文本字符串。
//
// 参数:
//   - 模型文本: 权限模型文本内容
//   - 适配器: 策略存储适配器，nil 则使用内存适配器
//
// 返回:
//   - *casbin.Enforcer: 权限管理器实例
//   - error: 创建失败时返回错误
func Q权限_从文本创建(模型文本 string, 适配器 persist.Adapter) (*casbin.Enforcer, error) {
	m, err := model.NewModelFromString(模型文本)
	if err != nil {
		return nil, err
	}
	if 适配器 == nil {
		return casbin.NewEnforcer(m)
	}
	return casbin.NewEnforcer(m, 适配器)
}

// Q权限_创建内存 从模型文本创建纯内存权限管理器（不持久化）。
// 适用于临时权限检查或测试场景。
//
// 参数:
//   - 模型文本: 权限模型文本内容
//
// 返回:
//   - *casbin.Enforcer: 权限管理器实例
//   - error: 创建失败时返回错误
func Q权限_创建内存(模型文本 string) (*casbin.Enforcer, error) {
	return Q权限_从文本创建(模型文本, nil)
}

// Q权限_创建文件适配器 创建文件适配器，用于策略持久化。
//
// 参数:
//   - 策略文件路径: CSV 格式的策略文件路径
//
// 返回:
//   - *fileadapter.Adapter: 文件适配器实例
func Q权限_创建文件适配器(策略文件路径 string) *fileadapter.Adapter {
	return fileadapter.NewAdapter(策略文件路径)
}

// Q权限_检查权限 检查主体是否有权限对资源执行操作。
// 基于 RBAC/ABAC 模型进行权限判断。
//
// 参数:
//   - 管理器: 权限管理器实例
//   - 主体: 请求主体（用户/角色），如 "alice"
//   - 资源: 目标资源，如 "data1"
//   - 操作: 请求操作，如 "read"
//
// 返回:
//   - bool: true 表示有权限，false 表示无权限
//   - error: 检查失败时返回错误
func Q权限_检查权限(管理器 *casbin.Enforcer, 主体 string, 资源 string, 操作 string) (bool, error) {
	return 管理器.Enforce(主体, 资源, 操作)
}

// Q权限_添加策略 添加一条权限策略。
// 添加后需调用 Q权限_保存策略 持久化到文件。
//
// 参数:
//   - 管理器: 权限管理器实例
//   - 主体: 主体标识
//   - 资源: 资源标识
//   - 操作: 操作标识
//
// 返回:
//   - bool: true 表示策略是新添加的，false 表示策略已存在
//   - error: 添加失败时返回错误
func Q权限_添加策略(管理器 *casbin.Enforcer, 主体 string, 资源 string, 操作 string) (bool, error) {
	return 管理器.AddPolicy(主体, 资源, 操作)
}

// Q权限_删除策略 删除一条权限策略。
//
// 参数:
//   - 管理器: 权限管理器实例
//   - 主体: 主体标识
//   - 资源: 资源标识
//   - 操作: 操作标识
//
// 返回:
//   - bool: true 表示策略被删除，false 表示策略不存在
//   - error: 删除失败时返回错误
func Q权限_删除策略(管理器 *casbin.Enforcer, 主体 string, 资源 string, 操作 string) (bool, error) {
	return 管理器.RemovePolicy(主体, 资源, 操作)
}

// Q权限_添加角色 为用户分配角色。
// 在 RBAC 模型中使用，将用户绑定到角色。
//
// 参数:
//   - 管理器: 权限管理器实例
//   - 用户: 用户标识
//   - 角色: 角色标识
//
// 返回:
//   - bool: true 表示分配成功
//   - error: 分配失败时返回错误
func Q权限_添加角色(管理器 *casbin.Enforcer, 用户 string, 角色 string) (bool, error) {
	return 管理器.AddRoleForUser(用户, 角色)
}

// Q权限_删除角色 移除用户的指定角色。
//
// 参数:
//   - 管理器: 权限管理器实例
//   - 用户: 用户标识
//   - 角色: 角色标识
//
// 返回:
//   - bool: true 表示移除成功
//   - error: 移除失败时返回错误
func Q权限_删除角色(管理器 *casbin.Enforcer, 用户 string, 角色 string) (bool, error) {
	return 管理器.DeleteRoleForUser(用户, 角色)
}

// Q权限_取用户角色 获取用户的所有角色列表。
//
// 参数:
//   - 管理器: 权限管理器实例
//   - 用户: 用户标识
//
// 返回:
//   - []string: 角色名称列表
//   - error: 查询失败时返回错误
func Q权限_取用户角色(管理器 *casbin.Enforcer, 用户 string) ([]string, error) {
	return 管理器.GetRolesForUser(用户)
}

// Q权限_保存策略 将当前内存中的策略保存到持久化存储。
// 使用文件适配器时，策略会写入 CSV 文件。
//
// 参数:
//   - 管理器: 权限管理器实例
//
// 返回:
//   - error: 保存失败时返回错误
func Q权限_保存策略(管理器 *casbin.Enforcer) error {
	return 管理器.SavePolicy()
}

// Q权限_加载策略 从持久化存储重新加载策略。
// 用于外部修改了策略文件后刷新内存中的策略。
//
// 参数:
//   - 管理器: 权限管理器实例
//
// 返回:
//   - error: 加载失败时返回错误
func Q权限_加载策略(管理器 *casbin.Enforcer) error {
	return 管理器.LoadPolicy()
}
