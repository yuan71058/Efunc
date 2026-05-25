package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/en/utils"
)

func main() {
	fmt.Println("=== Core Conversion Demo ===")

	// Int to string
	str := utils.IntToStr(12345)
	fmt.Printf("IntToStr(12345) = %s\n", str)

	// String to int
	num := utils.StrToInt("67890")
	fmt.Printf("StrToInt(\"67890\") = %d\n", num)

	// Float to string
	f := utils.Float64ToStr(3.14159, 2)
	fmt.Printf("Float64ToStr(3.14159, 2) = %s\n", f)

	// String to float
	fv := utils.StrToFloat64("123.456")
	fmt.Printf("StrToFloat64(\"123.456\") = %f\n", fv)

	// Base conversion
	hex := utils.BaseToInt("FF", 16)
	fmt.Printf("BaseToInt(\"FF\", 16) = %d\n", hex)

	// Text extraction
	between := utils.Text_Between("hello [world] today", "[", "]")
	fmt.Printf("Text_Between: %s\n", between)

	left := utils.Text_Left("hello world", 5)
	fmt.Printf("Text_Left: %s\n", left)

	// Encoding
	encoded := utils.Encoding_Base64Encode([]byte("Hello, World!"))
	fmt.Printf("Base64: %s\n", encoded)

	// Regex
	match := utils.Regex_IsMatch(`\d{3}-\d{4}`, "Phone: 123-4567")
	fmt.Printf("Regex match: %v\n", match)

	// Atomic operations
	var counter int64 = 0
	utils.AtomicInc(&counter)
	utils.AtomicAdd(&counter, 10)
	fmt.Printf("Atomic counter: %d\n", utils.AtomicLoad(&counter))

	// Array utilities
	arr := []string{"a", "b", "c", "a", "d"}
	fmt.Printf("Contains 'b': %v\n", utils.ArrayContains(arr, "b"))
	fmt.Printf("Distinct: %v\n", utils.ArrayDistinct(arr))

	// Map utilities
	m := map[string]interface{}{"name": "Alice", "age": 30}
	fmt.Printf("Map keys: %v\n", utils.MapKeys(m))

	// IP utilities
	fmt.Printf("IsValid IP '192.168.1.1': %v\n", utils.IP_IsValid("192.168.1.1"))

	// Checksums
	fmt.Printf("MD5: %s\n", utils.Checksum_MD5("hello"))
	fmt.Printf("SHA256: %s\n", utils.Checksum_SHA256("hello"))

	fmt.Println("\nDone!")
}