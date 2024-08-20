package main

import (
	"fmt"
	"math"
)

func main() {
	inqueueLength := 89
	maxQueueLength := 100
	ss := float64(inqueueLength) / float64(maxQueueLength)
	fmt.Println(ss)
	if float64(inqueueLength/maxQueueLength) > 0.9 {

	}

	memoryLimit := int64(100)

	memoryUsage := int64(78)

	memoryUsagePercent := (float64(memoryUsage) / float64(memoryLimit)) * 100

	// 3. 根据算法模型计算限流值
	// TODO 单测覆盖每个分支
	resultAcs := 15
	switch {
	case memoryUsagePercent <= 50:
		// memory使用量小于50%时，限流值acs无变化
		resultAcs = 15
	case memoryUsagePercent >= 80:
		// memory使用量大于80%时，最大力度限流
		resultAcs = 1
	default:
		// memory使用量在50%到80%时，根据限流算法限流：250-3*m  (四舍五入取整数)
		resultAcs = int(math.Round((250 - 3*memoryUsagePercent) / 100 * 15))
	}

	fmt.Println(resultAcs)

	fmt.Println("===================== arrayFunc")
	arrays()
	fmt.Println("===================== slice")
	slice()

}

// 数组
func arrays() {
	var test1 [6]int
	fmt.Println("内容：", test1[1:4])
}

// 切片
func slice() {
	// 比数组更强大的序列接口,比数组更常用
	primes := [6]int{2, 3, 5, 1, 7}
	var s []int = primes[1:4]
	fmt.Println(s)
}
