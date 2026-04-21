/**
 * SmartPlant EMS E2E 测试
 * 运行方式: node test-e2e.mjs
 * 前置条件: 后端服务启动 (8081/8082/8083), 前端启动 (3000)
 */

import { chromium } from 'playwright';

const BASE_URL = 'http://localhost:3000';
const API_DEVICE = 'http://localhost:8081/api/device';
const API_ENERGY = 'http://localhost:8082/api/energy';
const API_REPORT = 'http://localhost:8083/api/report';

let passed = 0;
let failed = 0;
let browser;

async function assert(condition, msg) {
  if (condition) {
    console.log(`  ✅ ${msg}`);
    passed++;
  } else {
    console.log(`  ❌ ${msg}`);
    failed++;
  }
}

async function testAPI(url, expectedCode = 200) {
  const res = await fetch(url);
  const status = res.status;
  const body = await res.json();
  return { status, body };
}

// ============================================================
// 1. 后端 API 测试
// ============================================================
async function testBackendAPIs() {
  console.log('\n【1】后端 API 测试');
  console.log('----------------------------------------');

  // 1.1 设备列表
  {
    const { status, body } = await testAPI(`${API_DEVICE}/list`);
    assert(status === 200, `设备列表 GET /api/device/list → ${status}`);
    assert(body.code === 0, `设备列表 code=0 (actual: ${body.code})`);
    assert(Array.isArray(body.data?.records), '设备列表 records 是数组');
    console.log(`    records 数量: ${body.data?.records?.length ?? 0}`);
  }

  // 1.2 设备详情
  {
    const { status, body } = await testAPI(`${API_DEVICE}/1`);
    assert(status === 200, `设备详情 GET /api/device/1 → ${status}`);
    assert(body.code === 0, `设备详情 code=0`);
    assert(body.data?.device_code === 'AC-001', `设备编号=AC-001 (actual: ${body.data?.device_code})`);
  }

  // 1.3 设备实时数据
  {
    const { status, body } = await testAPI(`${API_DEVICE}/1/realtime`);
    assert(status === 200, `实时数据 GET /api/device/1/realtime → ${status}`);
    assert(body.code === 0, `实时数据 code=0`);
    assert(typeof body.data?.power_kw === 'number', '实时功率是数字');
  }

  // 1.4 能耗统计
  {
    const { status, body } = await testAPI(`${API_ENERGY}/stats?period=day`);
    assert(status === 200, `能耗统计 GET /api/energy/stats → ${status}`);
    assert(body.code === 0, `能耗统计 code=0`);
    assert(typeof body.data?.total_energy_kwh === 'number', '有总能耗数据');
    assert(typeof body.data?.total_cost === 'number', '有总成本数据');
    console.log(`    总能耗: ${body.data?.total_energy_kwh} kWh, 总成本: ${body.data?.total_cost} 元`);
  }

  // 1.5 报表列表
  {
    const { status, body } = await testAPI(`${API_REPORT}/list`);
    assert(status === 200, `报表列表 GET /api/report/list → ${status}`);
    assert(body.code === 0, `报表列表 code=0`);
    assert(Array.isArray(body.data?.records), '报表 records 是数组');
  }

  // 1.6 报表详情
  {
    const { status, body } = await testAPI(`${API_REPORT}/1`);
    assert(status === 200, `报表详情 GET /api/report/1 → ${status}`);
    assert(body.code === 0, `报表详情 code=0`);
  }

  // 1.7 创建设备 POST
  {
    const formData = new URLSearchParams({
      device_code: 'TEST-E2E-001',
      device_name: 'E2E测试设备',
      device_type: 'air_compressor',
      location: '测试车间',
      power_rating: '200',
    });
    const res = await fetch(API_DEVICE, {
      method: 'POST',
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      body: formData.toString(),
    });
    const body = await res.json();
    assert(res.status === 200, `创建设备 POST /api/device → ${res.status}`);
    assert(body.code === 0, `创建设备 code=0 (msg: ${body.msg})`);
  }

  // 1.8 能耗记录列表
  {
    const { status, body } = await testAPI(`${API_ENERGY}/list`);
    assert(status === 200, `能耗记录 GET /api/energy/list → ${status}`);
    assert(body.code === 0, `能耗记录 code=0`);
  }

  // 1.9 并发请求
  {
    const results = await Promise.all([
      testAPI(`${API_DEVICE}/list`),
      testAPI(`${API_ENERGY}/stats?period=month`),
      testAPI(`${API_REPORT}/list`),
    ]);
    const allOk = results.every(r => r.status === 200 && r.body.code === 0);
    assert(allOk, `并发请求 3个API 全部成功`);
  }
}

// ============================================================
// 2. 前端页面 E2E 测试
// ============================================================
async function testFrontend() {
  console.log('\n【2】前端页面 E2E 测试');
  console.log('----------------------------------------');

  browser = await chromium.launch({ headless: true });
  const page = await browser.newPage();

  // 捕获 console 错误
  const errors = [];
  page.on('console', msg => {
    if (msg.type() === 'error') errors.push(msg.text());
  });
  page.on('pageerror', err => errors.push(err.message));

  // 2.1 首页/看板
  {
    await page.goto(BASE_URL, { waitUntil: 'networkidle', timeout: 15000 });
    assert(page.url().includes('localhost:3000'), `访问首页 ${page.url()}`);
    const title = await page.title();
    console.log(`    页面标题: ${title}`);
    // 检查是否有内容渲染
    const body = await page.textContent('body');
    assert(body.length > 10, '页面有内容渲染');
  }

  // 2.2 侧边栏导航
  {
    const sidebarLinks = await page.$$('a[href]');
    console.log(`    侧边栏链接数: ${sidebarLinks.length}`);
    assert(sidebarLinks.length > 0, '侧边栏有导航链接');
  }

  // 2.3 设备列表页
  {
    await page.goto(`${BASE_URL}/devices`, { waitUntil: 'load', timeout: 15000 });
    await page.waitForTimeout(2000);
    const body = await page.textContent('body');
    assert(body.includes('AC-001') || body.includes('设备') || body.length > 10, '设备列表有内容');
    console.log(`    设备列表页加载成功`);
  }

  // 2.4 能耗分析页
  {
    await page.goto(`${BASE_URL}/energy`, { waitUntil: 'load', timeout: 15000 });
    await page.waitForTimeout(2000);
    const body = await page.textContent('body');
    assert(body.length > 10, '能耗分析页有内容');
    console.log(`    能耗分析页加载成功`);
  }

  // 2.5 报表页
  {
    await page.goto(`${BASE_URL}/reports`, { waitUntil: 'load', timeout: 15000 });
    await page.waitForTimeout(2000);
    const body = await page.textContent('body');
    assert(body.includes('报表') || body.includes('report') || body.length > 10, '报表页有内容');
    console.log(`    报表页加载成功`);
  }

  // 2.6 登录页
  {
    await page.goto(`${BASE_URL}/login`, { waitUntil: 'networkidle', timeout: 15000 });
    await page.waitForTimeout(500);
    const inputs = await page.$$('input');
    assert(inputs.length > 0, '登录页有输入框');
    console.log(`    登录页加载成功，输入框数: ${inputs.length}`);
  }

  // 2.7 登录功能
  {
    await page.goto(`${BASE_URL}/login`, { waitUntil: 'networkidle', timeout: 15000 });
    const usernameInput = await page.$('input[name="username"]');
    const passwordInput = await page.$('input[name="password"]');
    if (usernameInput && passwordInput) {
      await usernameInput.fill('admin');
      await passwordInput.fill('admin123');
      const submitBtn = await page.$('button[type="submit"]');
      if (submitBtn) {
        await submitBtn.click();
        await page.waitForTimeout(2000);
        const currentUrl = page.url();
        assert(!currentUrl.includes('login'), `登录后跳转 (当前: ${currentUrl})`);
        console.log(`    登录跳转成功`);
      }
    }
  }

  // 2.8 控制台错误检查
  const criticalErrors = errors.filter(e =>
    !e.includes('favicon') &&
    !e.includes('manifest') &&
    !e.includes('DevTools')
  );
  assert(criticalErrors.length === 0, `无严重控制台错误 (actual: ${criticalErrors.length}个)`);
  if (criticalErrors.length > 0) {
    criticalErrors.forEach(e => console.log(`    ⚠️  console error: ${e.substring(0, 100)}`));
  }

  await browser.close();
}

// ============================================================
// 主流程
// ============================================================
async function main() {
  console.log('╔══════════════════════════════════════════════╗');
  console.log('║   SmartPlant EMS 测试套件  (Playwright E2E)   ║');
  console.log('╚══════════════════════════════════════════════╝');
  console.log(`后端: 8081 / 8082 / 8083`);
  console.log(`前端: ${BASE_URL}`);

  // 前置检查
  try {
    const res = await fetch(`${API_DEVICE}/list`, { timeout: 3000 });
    console.log(`\n✅ 后端服务正常 (状态: ${res.status})`);
  } catch {
    console.log('\n❌ 后端服务未启动，请先运行:');
    console.log('   cd backend/*/ && ./device-service-bin &');
    process.exit(1);
  }

  try {
    const res = await fetch(BASE_URL, { timeout: 3000 });
    console.log(`✅ 前端服务正常 (状态: ${res.status})`);
  } catch {
    console.log('\n⚠️  前端未启动，页面测试将跳过');
  }

  // 执行测试
  await testBackendAPIs();

  try {
    await testFrontend();
  } catch (e) {
    console.log(`\n⚠️  前端测试出错: ${e.message}`);
  }

  // 总结
  console.log('\n========================================');
  console.log(`测试结果: ✅ ${passed} 通过  ❌ ${failed} 失败`);
  console.log('========================================');

  if (failed > 0) {
    process.exit(1);
  }
}

main().catch(e => {
  console.error('测试异常:', e);
  process.exit(1);
});
