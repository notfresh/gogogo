package main


import (
	"fmt"
	"math/rand"
	"time"
)

func PrintRandInt(min, max, n int){
	// 设置种子
	rand.Seed(time.Now().UnixNano())

	// // 定义范围
	// min := -100
	// max := 500

	// 用 map 记录已生成的随机数
	used := make(map[int]bool)

	// 生成随机数
	for i := 0; i < n; i++ {
		num := generateRandomNumber(min, max, used)
		fmt.Println(num)
	}
}

func generateRandomNumber(min, max int, used map[int]bool) int {
	for {
		num := rand.Intn(max-min) + min
		if !used[num] {
			used[num] = true
			return num
		}
	}
}