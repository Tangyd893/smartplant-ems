-- SmartPlant EMS Database Init
CREATE DATABASE IF NOT EXISTS smartplant_ems CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE smartplant_ems;

-- 设备表
CREATE TABLE IF NOT EXISTS devices (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    device_code VARCHAR(64) NOT NULL UNIQUE COMMENT '设备编号',
    device_name VARCHAR(128) NOT NULL COMMENT '设备名称',
    device_type VARCHAR(32) NOT NULL COMMENT '设备类型：air_compressor空压机/injection_machine注塑机/laser激光切割机',
    location VARCHAR(128) COMMENT '安装位置',
    status TINYINT DEFAULT 1 COMMENT '状态：0停用1运行2维护3故障',
    power_rating DECIMAL(10,2) COMMENT '额定功率(kW)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_device_code (device_code),
    INDEX idx_device_type (device_type),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='设备表';

-- 能耗记录表（高频写入，按月分表）
CREATE TABLE IF NOT EXISTS energy_records (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    device_id BIGINT NOT NULL COMMENT '设备ID',
    record_time TIMESTAMP NOT NULL COMMENT '记录时间',
    power_kw DECIMAL(10,3) COMMENT '瞬时功率(kW)',
    energy_kwh DECIMAL(12,4) COMMENT '能耗(kWh)',
    voltage_v DECIMAL(8,2) COMMENT '电压(V)',
    current_a DECIMAL(8,3) COMMENT '电流(A)',
    power_factor DECIMAL(5,3) COMMENT '功率因数',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_device_time (device_id, record_time),
    INDEX idx_record_time (record_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='能耗记录表';

-- 能源成本表（按班次/工单统计）
CREATE TABLE IF NOT EXISTS cost_records (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    cost_date DATE NOT NULL COMMENT '成本日期',
    shift VARCHAR(16) COMMENT '班次：day/night',
    work_order VARCHAR(64) COMMENT '工单号',
    production_line VARCHAR(64) COMMENT '产线',
    total_energy_kwh DECIMAL(12,4) COMMENT '总能耗',
    unit_price DECIMAL(8,4) COMMENT '电价(元/kWh)',
    total_cost DECIMAL(12,2) COMMENT '总成本(元)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_cost_date (cost_date),
    INDEX idx_work_order (work_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='能源成本表';

-- 告警表
CREATE TABLE IF NOT EXISTS alerts (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    device_id BIGINT NOT NULL COMMENT '设备ID',
    alert_type VARCHAR(32) NOT NULL COMMENT '告警类型',
    alert_level TINYINT NOT NULL COMMENT '级别：1普通/2重要/3紧急',
    alert_msg TEXT COMMENT '告警信息',
    threshold_value DECIMAL(10,3) COMMENT '触发阈值',
    current_value DECIMAL(10,3) COMMENT '当前值',
    is_resolved TINYINT DEFAULT 0 COMMENT '是否解决',
    resolved_at TIMESTAMP NULL COMMENT '解决时间',
    resolved_by VARCHAR(64) NULL COMMENT '处理人',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_device_alerts (device_id, is_resolved),
    INDEX idx_alert_level (alert_level),
    INDEX idx_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='告警表';

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL UNIQUE,
    password_hash VARCHAR(256) NOT NULL,
    real_name VARCHAR(64) COMMENT '真实姓名',
    role VARCHAR(16) NOT NULL DEFAULT 'operator' COMMENT '角色：admin/admin/operator/technician',
    phone VARCHAR(32),
    email VARCHAR(128),
    status TINYINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_username (username),
    INDEX idx_role (role)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 报表表
CREATE TABLE IF NOT EXISTS reports (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    report_name VARCHAR(128) NOT NULL COMMENT '报表名称',
    report_type VARCHAR(32) NOT NULL COMMENT '报表类型',
    period_type VARCHAR(16) NOT NULL COMMENT '周期：daily/monthly/yearly',
    period_start DATE NOT NULL COMMENT '统计开始日期',
    period_end DATE NOT NULL COMMENT '统计结束日期',
    created_by BIGINT COMMENT '创建人',
    file_path VARCHAR(256) COMMENT '文件路径',
    status TINYINT DEFAULT 0 COMMENT '状态：0生成中/1已完成/2失败',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP NULL,
    INDEX idx_report_type (report_type),
    INDEX idx_period (period_start, period_end)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='报表表';

-- 插入默认管理员
INSERT INTO users (username, password_hash, real_name, role) VALUES
('admin', '$2a$10$XQyKBfFGgUFfOAKfN5K3OOfFPVJZBHxhZQJJ2xasFz8r0LhXwLq5u', '系统管理员', 'admin');