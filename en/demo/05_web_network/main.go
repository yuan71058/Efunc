package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/en/utils"
)

func main() {
	fmt.Println("=== Web & Network Demo ===")

	// HTTP Client (simple)
	fmt.Println("\n--- HTTP Client ---")
	resp, err := utils.HTTPClient_Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("HTTP GET error:", err)
	} else {
		fmt.Printf("Status: %d\n", resp.StatusCode())
		fmt.Printf("Body (first 200 chars): %s...\n", resp.String()[:200])
	}

	// HTTP Client with headers
	resp2, err := utils.HTTPClient_GetWithHeaders("https://httpbin.org/headers", map[string]string{
		"User-Agent": "Efunc/1.0",
	})
	if err != nil {
		fmt.Println("HTTP error:", err)
	} else {
		fmt.Printf("Headers response: %s\n", resp2.String()[:200])
	}

	// POST request
	resp3, err := utils.HTTPClient_Post("https://httpbin.org/post", map[string]interface{}{
		"name":  "Efunc",
		"value": 123,
	})
	if err != nil {
		fmt.Println("HTTP POST error:", err)
	} else {
		fmt.Printf("POST status: %d\n", resp3.StatusCode())
	}

	// Web URL parsing
	fmt.Println("\n--- URL Utils ---")
	domain := utils.WebUtils_GetDomain("https://www.example.com/path/to/page")
	fmt.Printf("Domain: %s\n", domain)

	// Cookie handling
	oldCookie := "session=abc123; user=alice"
	newCookie := "session=xyz789; token=secret"
	merged := utils.WebUtils_MergeCookies(oldCookie, newCookie)
	fmt.Printf("Merged cookies: %s\n", merged)

	cookieVal := utils.WebUtils_GetCookie(merged, "session")
	fmt.Printf("Session cookie: %s\n", cookieVal)

	// IP utilities
	fmt.Println("\n--- IP Utilities ---")
	fmt.Printf("192.168.1.1 valid: %v\n", utils.IP_IsValid("192.168.1.1"))
	fmt.Printf("10.0.0.1 is private: %v\n", utils.IP_IsPrivate("10.0.0.1"))

	longIP := utils.IP_StringToLong("192.168.1.1")
	backToStr := utils.IP_LongToString(longIP)
	fmt.Printf("192.168.1.1 -> long: %d -> back: %s\n", longIP, backToStr)

	fmt.Println("\nDone!")
}