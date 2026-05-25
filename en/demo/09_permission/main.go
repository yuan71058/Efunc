package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/en/utils"
)

func main() {
	fmt.Println("=== Permission Management Demo (RBAC) ===\n")

	modelText := `
[request_definition]
r = sub, obj, act
[policy_definition]
p = sub, obj, act
[role_definition]
g = _, _
[policy_effect]
e = some(where (p.eft == allow))
[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	enforcer, err := utils.Permission_CreateMemory(modelText)
	if err != nil {
		fmt.Println("Create error:", err)
		return
	}

	// Define policies
	utils.Permission_AddPolicy(enforcer, "admin", "data1", "read")
	utils.Permission_AddPolicy(enforcer, "admin", "data1", "write")
	utils.Permission_AddPolicy(enforcer, "admin", "data1", "delete")

	utils.Permission_AddPolicy(enforcer, "editor", "data1", "read")
	utils.Permission_AddPolicy(enforcer, "editor", "data1", "write")

	utils.Permission_AddPolicy(enforcer, "viewer", "data1", "read")

	// Assign roles
	utils.Permission_AddRoleForUser(enforcer, "alice", "admin")
	utils.Permission_AddRoleForUser(enforcer, "bob", "editor")
	utils.Permission_AddRoleForUser(enforcer, "charlie", "viewer")

	// Check permissions
	users := []string{"alice", "bob", "charlie", "unknown"}
	actions := []string{"read", "write", "delete"}

	for _, user := range users {
		roles, _ := utils.Permission_GetRolesForUser(enforcer, user)
		fmt.Printf("User: %s (roles: %v)\n", user, roles)
		for _, action := range actions {
			ok, _ := utils.Permission_Enforce(enforcer, user, "data1", action)
			status := "DENIED"
			if ok {
				status = "ALLOWED"
			}
			fmt.Printf("  %s on data1 -> %s\n", action, status)
		}
		fmt.Println()
	}

	fmt.Println("Done!")
}