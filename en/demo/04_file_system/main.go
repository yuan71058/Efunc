package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/en/utils"
)

func main() {
	fmt.Println("=== File & System Demo ===")

	// File operations
	fileName := "demo_test_file.txt"

	fmt.Println("\n--- File Write/Read ---")
	err := utils.File_Write(fileName, []byte("Hello, Efunc File System!"))
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
	fmt.Println("File written successfully")

	content, err := utils.File_Read(fileName)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Printf("File content: %s\n", content)

	exists := utils.File_Exists(fileName)
	fmt.Printf("File exists: %v\n", exists)

	size := utils.File_GetSize(fileName)
	fmt.Printf("File size: %d bytes\n", size)

	// Clean up
	utils.File_Delete(fileName)

	// Directory operations
	fmt.Println("\n--- Directory ---")
	utils.Dir_Create("demo_dir")
	fmt.Printf("Dir exists: %v\n", utils.Dir_Exists("demo_dir"))

	utils.File_Write("demo_dir/test1.txt", []byte("file1"))
	utils.File_Write("demo_dir/test2.txt", []byte("file2"))

	files := utils.Dir_List("demo_dir")
	fmt.Printf("Files in dir: %v\n", files)

	utils.Dir_RemoveAll("demo_dir")

	// Program info
	fmt.Println("\n--- Program Info ---")
	programPath := utils.Program_GetPath()
	programDir := utils.Program_GetDir()
	fmt.Printf("Program path: %s\n", programPath)
	fmt.Printf("Program dir: %s\n", programDir)

	// Environment
	fmt.Println("\n--- Environment ---")
	fmt.Printf("GOPATH: %s\n", utils.Env_Get("GOPATH", "N/A"))

	// Logging
	fmt.Println("\n--- Logging ---")
	utils.Log_Info("info", "This is an info message")
	utils.Log_Debug("debug", "This is a debug message")
	utils.Log_Warn("warn", "This is a warning message")

	fmt.Println("\nDone!")
}