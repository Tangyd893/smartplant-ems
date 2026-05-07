# SmartPlant EMS - 智慧工厂能源管理系统

基于 Beego 2.x 微服务架构的前后端分离项目。

## 技术栈

| 组件 | 技术 |
|------|------|
| 后端框架 | Beego 2.x (Go) |
| 数据库 | MySQL 8 |
| 前端 | React 18 + Vite + TypeScript |
| 反向代理 | Nginx |
| 部署 | Docker Compose |

## 项目结构

```
smartplant-ems/
├── backend/
│   ├── device-service/   # 设备管理服务 (:8081)
│   ├── energy-service/    # 能耗分析服务 (:8082)
│   └── report-service/    # 报表服务 (:8083)
├── frontend/              # React 前端 (:3000)
├── docker/                # Docker 配置
│   ├── docker-compose.yml
│   ├── mysql/init.sql
│   └── nginx/
└── docs/                  # 项目文档
```

## 快速启动

### 1. 启动基础设施 (MySQL + Nginx)
```bash
cd docker
docker compose up -d
```

### 2. 启动后端服务（每个服务一个终端）
```bash
# 设备服务
cd backend/device-service
./device-service-bin

# 能耗服务
cd backend/energy-service
./device-service-bin

# 报表服务
cd backend/report-service
./device-service-bin
```

### 3. 启动前端
```bash
cd frontend
npm install
npm run dev
```

## 默认账号

- 用户名: admin
- 密码: admin123

## API 端口

| 服务 | 端口 |
|------|------|
| 设备服务 | 8081 |
| 能耗服务 | 8082 |
| 报表服务 | 8083 |
| Nginx API Gateway | 8080 |
| React 前端 | 3000 |
| MySQL | 3306 |

## 微服务 API

### 设备服务 `/api/device`
- `GET /api/device/list` - 设备列表
- `GET /api/device/:id` - 设备详情
- `POST /api/device` - 创建设备
- `PUT /api/device/:id` - 更新设备
- `DELETE /api/device/:id` - 删除设备
- `GET /api/device/:id/realtime` - 实时数据

### 能耗服务 `/api/energy`
- `GET /api/energy/list` - 能耗记录列表
- `GET /api/energy/:id` - 记录详情
- `POST /api/energy` - 创建记录
- `GET /api/energy/stats` - 能耗统计

### 报表服务 `/api/report`
- `GET /api/report/list` - 报表列表
- `GET /api/report/:id` - 报表详情
- `POST /api/report` - 生成报表
- `GET /api/report/:id/download` - 下载报表
- `DELETE /api/report/:id` - 删除报表