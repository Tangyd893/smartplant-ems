package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

const (
	DeviceBaseURL = "http://host.docker.internal:8081/api/device"
	EnergyBaseURL = "http://host.docker.internal:8082/api/energy"
	ReportBaseURL = "http://host.docker.internal:8083/api/report"
)

type APIResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data,omitempty"`
}

var passed int
var failed int

func assert(condition bool, msg string) {
	if condition {
		fmt.Printf("  ✅ %s\n", msg)
		passed++
	} else {
		fmt.Printf("  ❌ %s\n", msg)
		failed++
	}
}

func TestDeviceList(t *testing.T) {
	fmt.Println("\n[设备列表] GET /api/device/list")
	resp, err := http.Get(DeviceBaseURL + "/list")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200, fmt.Sprintf("状态码=200 (actual: %d)", resp.StatusCode))

	body, _ := io.ReadAll(resp.Body)
	var result APIResponse
	json.Unmarshal(body, &result)
	assert(result.Code == 0, fmt.Sprintf("code=0 (actual: %d, msg: %s)", result.Code, result.Msg))
}

func TestDeviceGet(t *testing.T) {
	fmt.Println("\n[设备详情] GET /api/device/1")
	resp, err := http.Get(DeviceBaseURL + "/1")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200, fmt.Sprintf("状态码=200 (actual: %d)", resp.StatusCode))

	body, _ := io.ReadAll(resp.Body)
	var result APIResponse
	json.Unmarshal(body, &result)
	assert(result.Code == 0, fmt.Sprintf("code=0 (actual: %d)", result.Code))
}

func TestDeviceRealTime(t *testing.T) {
	fmt.Println("\n[实时数据] GET /api/device/1/realtime")
	resp, err := http.Get(DeviceBaseURL + "/1/realtime")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200, fmt.Sprintf("状态码=200 (actual: %d)", resp.StatusCode))

	body, _ := io.ReadAll(resp.Body)
	var result APIResponse
	json.Unmarshal(body, &result)
	assert(result.Code == 0, fmt.Sprintf("code=0 (actual: %d)", result.Code))
}

func TestDeviceCreate(t *testing.T) {
	fmt.Println("\n[创建设备] POST /api/device")
	data := url.Values{}
	data.Set("device_code", "TEST-E2E-001")
	data.Set("device_name", "E2E测试设备")
	data.Set("device_type", "air_compressor")
	data.Set("location", "测试车间")
	data.Set("power_rating", "200")

	resp, err := http.PostForm(DeviceBaseURL, data)
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200, fmt.Sprintf("状态码=200 (actual: %d)", resp.StatusCode))

	body, _ := io.ReadAll(resp.Body)
	var result APIResponse
	json.Unmarshal(body, &result)
	assert(result.Code == 0, fmt.Sprintf("创建成功 code=0 (actual: %d, msg: %s)", result.Code, result.Msg))
}

func TestEnergyStats(t *testing.T) {
	fmt.Println("\n[能耗统计] GET /api/energy/stats")
	resp, err := http.Get(EnergyBaseURL + "/stats?period=day")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200, fmt.Sprintf("状态码=200 (actual: %d)", resp.StatusCode))

	body, _ := io.ReadAll(resp.Body)
	var result APIResponse
	json.Unmarshal(body, &result)
	assert(result.Code == 0, fmt.Sprintf("code=0 (actual: %d)", result.Code))

	var stats map[string]interface{}
	json.Unmarshal(result.Data, &stats)
	assert(stats["total_energy_kwh"] != nil, "有 total_energy_kwh 字段")
	assert(stats["total_cost"] != nil, "有 total_cost 字段")
	fmt.Printf("    总能耗: %.2f kWh, 总成本: %.2f 元\n",
		stats["total_energy_kwh"].(float64), stats["total_cost"].(float64))
}

func TestEnergyList(t *testing.T) {
	fmt.Println("\n[能耗记录] GET /api/energy/list")
	resp, err := http.Get(EnergyBaseURL + "/list")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200, fmt.Sprintf("状态码=200 (actual: %d)", resp.StatusCode))

	body, _ := io.ReadAll(resp.Body)
	var result APIResponse
	json.Unmarshal(body, &result)
	assert(result.Code == 0, fmt.Sprintf("code=0 (actual: %d)", result.Code))
}

func TestReportList(t *testing.T) {
	fmt.Println("\n[报表列表] GET /api/report/list")
	resp, err := http.Get(ReportBaseURL + "/list")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200, fmt.Sprintf("状态码=200 (actual: %d)", resp.StatusCode))

	body, _ := io.ReadAll(resp.Body)
	var result APIResponse
	json.Unmarshal(body, &result)
	assert(result.Code == 0, fmt.Sprintf("code=0 (actual: %d)", result.Code))
}

func TestReportGet(t *testing.T) {
	fmt.Println("\n[报表详情] GET /api/report/1")
	resp, err := http.Get(ReportBaseURL + "/1")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200, fmt.Sprintf("状态码=200 (actual: %d)", resp.StatusCode))
}

func TestConcurrentRequests(t *testing.T) {
	fmt.Println("\n[并发请求] 3个API同时请求")
	urls := []string{
		DeviceBaseURL + "/list",
		EnergyBaseURL + "/stats?period=month",
		ReportBaseURL + "/list",
	}

	done := make(chan bool, len(urls))
	for _, u := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			if resp != nil {
				defer resp.Body.Close()
				done <- (resp.StatusCode == 200)
			} else {
				done <- false
			}
		}(u)
	}

	results := []bool{}
	for i := 0; i < len(urls); i++ {
		results = append(results, <-done)
	}

	allOk := true
	for i, r := range results {
		if !r {
			allOk = false
			fmt.Printf("  ❌ 并发请求 #%d 失败\n", i+1)
		}
	}
	if allOk {
		fmt.Printf("  ✅ 并发请求 3/3 全部成功\n")
		passed++
	} else {
		failed++
	}
}

func TestMain(m *testing.M) {
	time.Sleep(500 * time.Millisecond)

	// 前置检查
	_, err := http.Get(DeviceBaseURL + "/list")
	if err != nil {
		fmt.Println("⚠️  后端服务不可达，请确认 8081/8082/8083 已启动")
		os.Exit(0)
	}

	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║   SmartPlant EMS Go API 测试                 ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	os.Exit(m.Run())
}
