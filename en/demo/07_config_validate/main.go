package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/en/utils"
)

type LoginForm struct {
	Username string `validate:"required,min=3,max=20"`
	Password string `validate:"required,min=6"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"gte=0,lte=130"`
}

func main() {
	fmt.Println("=== Validation & Config Demo ===")

	// Validation
	fmt.Println("\n--- Data Validation ---")
	valid := LoginForm{
		Username: "alice",
		Password: "secret123",
		Email:    "alice@example.com",
		Age:      25,
	}
	err := utils.Validation_Struct(valid)
	if err != nil {
		fmt.Println("Validation errors:", err)
	} else {
		fmt.Println("Valid data: OK")
	}

	invalid := LoginForm{
		Username: "ab",
		Password: "123",
		Email:    "not-an-email",
		Age:      999,
	}
	err = utils.Validation_Struct(invalid)
	if err != nil {
		fmt.Println("Invalid data errors:", err)
	}

	// Config (Viper) - in-memory
	fmt.Println("\n--- Config ---")
	cfg := utils.Config_New()
	utils.Config_SetDefault(cfg, "app.name", "EfuncDemo")
	utils.Config_SetDefault(cfg, "app.version", "1.0.0")
	utils.Config_SetDefault(cfg, "server.port", 8080)
	utils.Config_SetDefault(cfg, "database.host", "localhost")

	fmt.Printf("App name: %s\n", utils.Config_GetString(cfg, "app.name"))
	fmt.Printf("App version: %s\n", utils.Config_GetString(cfg, "app.version"))
	fmt.Printf("Server port: %d\n", utils.Config_GetInt(cfg, "server.port"))

	// Expression evaluation
	fmt.Println("\n--- Expression Eval ---")
	result, err := utils.Expression_Eval("2 + 3 * 4")
	if err != nil {
		fmt.Println("Eval error:", err)
	} else {
		fmt.Printf("2 + 3 * 4 = %v\n", result)
	}

	result2, err := utils.Expression_EvalWithVars("x * y + z", map[string]interface{}{
		"x": 10, "y": 20, "z": 5,
	})
	if err != nil {
		fmt.Println("Eval error:", err)
	} else {
		fmt.Printf("x*y+z (10*20+5) = %v\n", result2)
	}

	// Time utilities
	fmt.Println("\n--- Time Utilities ---")
	now := utils.Time_Now()
	fmt.Printf("Now: %s\n", utils.Time_Format(now, "2006-01-02 15:04:05"))
	fmt.Printf("Unix: %d\n", utils.Time_Unix(now))
	fmt.Printf("StartOfDay: %s\n", utils.Time_Format(utils.Time_StartOfDay(now), "15:04:05"))

	// Date parsing
	parsed, err := utils.DateParse_Any("2024-06-15")
	if err != nil {
		fmt.Println("Date parse error:", err)
	} else {
		fmt.Printf("Parsed date: %s\n", utils.Time_Format(parsed, "2006-01-02"))
	}

	// Template rendering
	fmt.Println("\n--- Template ---")
	tmpl := utils.Template_New("demo")
	tmpl = utils.Template_Parse(tmpl, "Hello, {{.Name}}! Your score is {{.Score}}.")
	rendered := utils.Template_Render(tmpl, map[string]interface{}{
		"Name":  "Alice",
		"Score": 95,
	})
	fmt.Printf("Rendered: %s\n", rendered)

	// Struct merge
	fmt.Println("\n--- Struct Merge ---")
	type ConfigA struct {
		Host string
		Port int
	}
	type ConfigB struct {
		Port    int
		Timeout int
	}

	a := ConfigA{Host: "localhost", Port: 8080}
	b := ConfigB{Port: 9090, Timeout: 30}

	var merged ConfigA
	utils.StructMerge_Merge(&merged, a)
	utils.StructMerge_MergeWithOverride(&merged, b)
	fmt.Printf("Merged config: Host=%s Port=%d\n", merged.Host, merged.Port)

	fmt.Println("\nDone!")
}