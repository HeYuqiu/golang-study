package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type RequestInfo struct {
	verb      string
	url       string
	userAgent string
	client    string
	totalTime int
	etcdKey   string
	etcdLimit string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <logfile>")
		return
	}

	logfile := os.Args[1]
	file, err := os.Open(logfile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	noListRequests := make(map[string]int)
	listRequestsByUrlUa := make(map[string]int)
	listRequestsByClient := make(map[string]int)
	long5Requests := make(map[string]int)
	long50Requests := make(map[string]int)
	etcdRequests := make(map[string]int)
	etcdLong5Requests := make(map[string]int)

	var startTime, endTime time.Time

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// APIServer 请求
		if strings.Contains(line, "url:") && strings.Contains(line, "user-agent:") && strings.Contains(line, "total time:") {
			logTime := extractLogTime(line)
			if startTime.IsZero() || startTime.After(logTime) {
				startTime = logTime
			}
			if endTime.IsZero() || endTime.Before(logTime) {
				endTime = logTime
			}

			info := parseLogLine(line)
			keyNoList := fmt.Sprintf("【Verb】:%s,【URL】:%s,【UserAgent】:%s", info.verb, info.url, info.userAgent)
			keyList := fmt.Sprintf("【URL】:%s,【UserAgent】:%s", info.url, info.userAgent)
			keyClientIP := fmt.Sprintf("【ClientIP】:%s,【UserAgent】:%s", info.client, info.userAgent)
			if strings.ToLower(info.verb) != "list" {
				noListRequests[keyNoList]++
			} else {
				listRequestsByUrlUa[keyList]++
				listRequestsByClient[keyClientIP]++
			}
			if info.totalTime > 5000 {
				long5Requests[keyNoList]++
			}
			if info.totalTime > 50000 {
				long50Requests[keyNoList]++
			}
		}

		// etcd 请求
		if strings.Contains(line, "etcd3") && strings.Contains(line, "key:") && strings.Contains(line, "limit:") {
			info := parseLogLineETCD(line)
			etcdKey := fmt.Sprintf("【Key】:%s,【Limit】:%s", info.etcdKey, info.etcdLimit)
			etcdRequests[etcdKey]++
			if info.totalTime > 5000 {
				etcdLong5Requests[etcdKey]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("\n\n—————————————————————————————————[%s] ~ [%s] Kube-APIServer 日志分析（时延超过500ms的请求）—————————————————————————————————\n", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))

	printTopN("非List请求 Top10:", noListRequests, 10)
	printTopN("List请求 Top10:", listRequestsByUrlUa, 10)
	printTopN("List请求IP来源 Top8:", listRequestsByClient, 8)
	printTopN("时延超过5秒的请求 Top8:", long5Requests, 8)
	printTopN("时延超过50秒的请求 Top6:", long50Requests, 6)
	printTopN("ETCD请求 Top8:", etcdRequests, 8)
	printTopN("ETCD请求时延超过5秒的请求 Top6:", etcdLong5Requests, 6)
}

func parseLogLine(line string) *RequestInfo {
	verb := extractValue(line, `""([^" ]+)""`)
	url := extractValue(line, `url:([^,]+)`)
	userAgent := extractValue(line, `user-agent:([^,]+)`)
	client := extractValue(line, `client:([^, ]+)`)
	totalTimeStr := extractValue(line, `total time: ([0-9]+)`)
	totalTime, _ := strconv.Atoi(totalTimeStr)

	return &RequestInfo{
		verb:      verb,
		url:       url,
		userAgent: userAgent,
		client:    client,
		totalTime: totalTime,
	}
}
func parseLogLineETCD(line string) *RequestInfo {
	etcdKey := extractValue(line, `key:([^,]+)`)
	etcdLimit := extractValue(line, `limit:([^,]+)`)
	var totalTimeStr string
	if strings.Contains(line, "total time:") {
		totalTimeStr = extractValue(line, `total time: ([0-9]+)`)
	} else {
		totalTimeStr = extractValue(line, `\s(\d+)ms\s`)

	}
	totalTime, _ := strconv.Atoi(totalTimeStr)
	return &RequestInfo{
		totalTime: totalTime,
		etcdKey:   etcdKey,
		etcdLimit: etcdLimit,
	}
}

func extractValue(line, pattern string) string {
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractLogTime(line string) time.Time {
	re := regexp.MustCompile(`I(\d{4}) (\d{2}:\d{2}:\d{2})`)
	match := re.FindStringSubmatch(line)
	if len(match) > 2 {
		monthDay := match[1]
		timeOfDay := match[2]
		logTime, err := time.Parse("0102 15:04:05", monthDay+" "+timeOfDay)
		if err == nil {
			// Assuming the logs are from the current year
			logTime = logTime.AddDate(time.Now().Year(), 0, 0).Add(8 * time.Hour)
			return logTime
		}
	}
	return time.Time{}
}

func printTopN(title string, data map[string]int, n int) {
	fmt.Println("\n" + title)
	type kv struct {
		Key   string
		Value int
	}
	var sortedData []kv
	for k, v := range data {
		sortedData = append(sortedData, kv{k, v})
	}
	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i].Value > sortedData[j].Value
	})
	for i, kv := range sortedData {
		if i >= n {
			break
		}
		fmt.Printf("%d.【请求次数】:%d, %s\n", i+1, kv.Value, kv.Key)
	}
}
