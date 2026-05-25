package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/en/utils"
)

func main() {
	fmt.Println("=== Concurrency Demo ===")

	// Goroutine Pool
	fmt.Println("\n--- Goroutine Pool ---")
	pool := utils.Pool_New(5)
	defer utils.Pool_Release(pool)

	for i := 0; i < 20; i++ {
		idx := i
		utils.Pool_Submit(pool, func() {
			fmt.Printf("Task %d running on goroutine\n", idx)
		})
	}
	poolRunning := utils.Pool_Running(pool)
	fmt.Printf("Pool running goroutines: %d\n", poolRunning)

	// Message Bus
	fmt.Println("\n--- Message Bus ---")
	utils.Bus_Subscribe("order.created", func(msg interface{}) {
		fmt.Printf("[Subscriber 1] Order created: %v\n", msg)
	})
	utils.Bus_Subscribe("order.created", func(msg interface{}) {
		fmt.Printf("[Subscriber 2] Send confirmation: %v\n", msg)
	})

	utils.Bus_Publish("order.created", map[string]interface{}{
		"id":     1001,
		"amount": 99.99,
	})

	// Cron
	fmt.Println("\n--- Cron Demo (single execution) ---")
	cron := utils.Cron_New()
	utils.Cron_AddJob(cron, "@every 1s", func() {
		fmt.Println("Cron tick:", utils.Time_Now())
	})
	utils.Cron_Start(cron)

	// Let it run for 3 seconds
	utils.Helper_Wait(3000)
	utils.Cron_Stop(cron)

	// Byte Buffer Pool
	fmt.Println("\n--- ByteBuffer Pool ---")
	buf := utils.BufferPool_Get()
	buf.WriteString("Hello from buffer pool!")
	fmt.Println(buf.String())
	utils.BufferPool_Put(buf)

	fmt.Println("\nDone!")
}