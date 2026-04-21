#!/bin/bash
# SmartPlant EMS API 测试脚本 (Go-less 版本)
# 用 curl 测试所有 API 端点

BASE_DEV="http://127.0.0.1:8081/api/device"
BASE_EN="http://127.0.0.1:8082/api/energy"
BASE_RPT="http://127.0.0.1:8083/api/report"

PASSED=0
FAILED=0

assert() {
  if [ $1 -eq 0 ]; then
    echo "  ✅ $2"
    ((PASSED++))
  else
    echo "  ❌ $2 (exit: $1)"
    ((FAILED++))
  fi
}

echo "╔══════════════════════════════════════════════╗"
echo "║   SmartPlant EMS API 测试 (curl版)           ║"
echo "╚══════════════════════════════════════════════╝"

# 检查服务可用性
echo -e "\n[前置检查]"
curl -s -o /dev/null -w "%{http_code}" "$BASE_DEV/list" | grep -q "200"
assert $? "device-service :8081 可达"

# ===== 设备服务 =====
echo -e "\n【设备服务】"

CODE=$(curl -s -o /tmp/dev_list.json -w "%{http_code}" "$BASE_DEV/list")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "GET /api/device/list → $CODE"
assert $(grep -q '"code": 0' /tmp/dev_list.json && echo 0 || echo 1) "设备列表 code=0"
assert $(grep -q 'AC-001' /tmp/dev_list.json && echo 0 || echo 1) "设备编号 AC-001 存在"

CODE=$(curl -s -o /tmp/dev_get.json -w "%{http_code}" "$BASE_DEV/1")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "GET /api/device/1 → $CODE"
assert $(grep -q '"code": 0' /tmp/dev_get.json && echo 0 || echo 1) "设备详情 code=0"

CODE=$(curl -s -o /tmp/dev_rt.json -w "%{http_code}" "$BASE_DEV/1/realtime")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "GET /api/device/1/realtime → $CODE"
assert $(grep -q '"code": 0' /tmp/dev_rt.json && echo 0 || echo 1) "实时数据 code=0"
assert $(grep -q 'power_kw' /tmp/dev_rt.json && echo 0 || echo 1) "实时数据包含 power_kw"

# POST 创建设备
CODE=$(curl -s -o /tmp/dev_post.json -w "%{http_code}" -X POST \
  -d "device_code=TEST-001&device_name=测试设备&device_type=air_compressor&location=测试车间&power_rating=200" \
  "$BASE_DEV")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "POST /api/device → $CODE"
assert $(grep -q '"code": 0' /tmp/dev_post.json && echo 0 || echo 1) "创建设备 code=0"

# PUT 更新设备
CODE=$(curl -s -o /tmp/dev_put.json -w "%{http_code}" -X PUT \
  -d "device_name=已更新" \
  "$BASE_DEV/1")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "PUT /api/device/1 → $CODE"
assert $(grep -q '"code": 0' /tmp/dev_put.json && echo 0 || echo 1) "更新设备 code=0"

# DELETE 删除设备
CODE=$(curl -s -o /tmp/dev_del.json -w "%{http_code}" -X DELETE "$BASE_DEV/99")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "DELETE /api/device/99 → $CODE"

# ===== 能耗服务 =====
echo -e "\n【能耗服务】"

CODE=$(curl -s -o /tmp/en_list.json -w "%{http_code}" "$BASE_EN/list")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "GET /api/energy/list → $CODE"
assert $(grep -q '"code": 0' /tmp/en_list.json && echo 0 || echo 1) "能耗列表 code=0"

CODE=$(curl -s -o /tmp/en_stats.json -w "%{http_code}" "$BASE_EN/stats?period=day")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "GET /api/energy/stats?period=day → $CODE"
assert $(grep -q '"code": 0' /tmp/en_stats.json && echo 0 || echo 1) "能耗统计 code=0"
assert $(grep -q 'total_energy_kwh' /tmp/en_stats.json && echo 0 || echo 1) "统计包含 total_energy_kwh"
assert $(grep -q 'total_cost' /tmp/en_stats.json && echo 0 || echo 1) "统计包含 total_cost"

CODE=$(curl -s -o /tmp/en_stats2.json -w "%{http_code}" "$BASE_EN/stats?period=month")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "GET /api/energy/stats?period=month → $CODE"

CODE=$(curl -s -o /tmp/en_post.json -w "%{http_code}" -X POST \
  -d "device_id=1&power_kw=120.5&energy_kwh=38500" \
  "$BASE_EN")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "POST /api/energy → $CODE"

# ===== 报表服务 =====
echo -e "\n【报表服务】"

CODE=$(curl -s -o /tmp/rpt_list.json -w "%{http_code}" "$BASE_RPT/list")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "GET /api/report/list → $CODE"
assert $(grep -q '"code": 0' /tmp/rpt_list.json && echo 0 || echo 1) "报表列表 code=0"

CODE=$(curl -s -o /tmp/rpt_get.json -w "%{http_code}" "$BASE_RPT/1")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "GET /api/report/1 → $CODE"
assert $(grep -q '"code": 0' /tmp/rpt_get.json && echo 0 || echo 1) "报表详情 code=0"

CODE=$(curl -s -o /tmp/rpt_post.json -w "%{http_code}" -X POST \
  -d "report_name=测试报表&report_type=daily&period_type=daily&period_start=2026-04-20&period_end=2026-04-20" \
  "$BASE_RPT")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "POST /api/report → $CODE"

CODE=$(curl -s -o /tmp/rpt_del.json -w "%{http_code}" -X DELETE "$BASE_RPT/99")
assert $([ "$CODE" = "200" ] && echo 0 || echo 1) "DELETE /api/report/99 → $CODE"

# ===== 并发测试 =====
echo -e "\n【并发测试】"
curl -s "$BASE_DEV/list" > /tmp/c1.json &
curl -s "$BASE_EN/stats?period=month" > /tmp/c2.json &
curl -s "$BASE_RPT/list" > /tmp/c3.json &
wait
R1=$(grep -c '"code": 0' /tmp/c1.json 2>/dev/null || echo 0)
R2=$(grep -c '"code": 0' /tmp/c2.json 2>/dev/null || echo 0)
R3=$(grep -c '"code": 0' /tmp/c3.json 2>/dev/null || echo 0)
assert $([ "$R1" -ge 1 ] && [ "$R2" -ge 1 ] && [ "$R3" -ge 1 ] && echo 0 || echo 1) \
  "并发 3 API 全部成功 (dev:$R1 en:$R2 rpt:$R3)"

# ===== 清理 =====
rm -f /tmp/dev_*.json /tmp/en_*.json /tmp/rpt_*.json /tmp/c[123].json

echo -e "\n========================================"
echo "测试结果: ✅ $PASSED 通过  ❌ $FAILED 失败"
echo "========================================"

[ $FAILED -gt 0 ] && exit 1 || exit 0
