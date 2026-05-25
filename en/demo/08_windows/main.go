//go:build windows

package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/en/utils"
)

func main() {
	fmt.Println("=== Windows Features Demo ===\n")

	// Window management
	fmt.Println("--- Window Management ---")
	desktop := utils.Window_GetDesktop()
	fmt.Printf("Desktop window: 0x%X\n", desktop)

	// Find calculator window
	calc := utils.Window_Find("CalcFrame", "")
	if calc != 0 {
		fmt.Println("Found Calculator window")
		title := utils.Window_GetTitle(calc)
		fmt.Printf("  Title: %s\n", title)

		className := utils.Window_GetClassName(calc)
		fmt.Printf("  Class: %s\n", className)

		rect, ok := utils.Window_GetRect(calc)
		if ok {
			fmt.Printf("  Rect: (%d,%d) %dx%d\n", rect.Left, rect.Top,
				rect.Right-rect.Left, rect.Bottom-rect.Top)
		}
	} else {
		fmt.Println("Calculator not running")
	}

	// Process enumeration
	fmt.Println("\n--- Process Enumeration ---")
	processes, err := utils.Memory_EnumProcesses()
	if err != nil {
		fmt.Println("Enum error:", err)
	} else {
		fmt.Printf("Total processes: %d\n", len(processes))
		for i, pe := range processes {
			if i >= 5 {
				fmt.Printf("  ... and %d more\n", len(processes)-5)
				break
			}
			name := utils.UTF16ArrayToString(pe.SzExeFile[:])
			fmt.Printf("  PID=%d %s\n", pe.Th32ProcessID, name)
		}
	}

	// Input simulation example (commented out for safety)
	fmt.Println("\n--- Input Simulation (commented out) ---")
	fmt.Println("// utils.Input_KeyDown(utils.VK_CONTROL)")
	fmt.Println("// utils.Input_KeyDown('C')")
	fmt.Println("// utils.Input_KeyUp('C')")
	fmt.Println("// utils.Input_KeyUp(utils.VK_CONTROL)")

	// System info
	fmt.Println("\n--- System Info ---")
	cpuPercent := utils.SysInfo_CPUPercent()
	fmt.Printf("CPU usage: %.1f%%\n", cpuPercent)

	memInfo := utils.SysInfo_MemoryInfo()
	fmt.Printf("Memory total: %.2f GB\n", float64(memInfo.Total)/(1024*1024*1024))
	fmt.Printf("Memory used: %.2f GB\n", float64(memInfo.Used)/(1024*1024*1024))

	// Foreground window
	foreground := utils.Window_GetForeground()
	if foreground != 0 {
		title := utils.Window_GetTitle(foreground)
		fmt.Printf("\nForeground window: %s\n", title)
	}

	// Clipboard
	utils.SysCmd_SetClipboard("Hello from Efunc!")
	text := utils.SysCmd_GetClipboard()
	fmt.Printf("Clipboard: %s\n", text)

	fmt.Println("\nDone!")
}