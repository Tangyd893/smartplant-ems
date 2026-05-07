package models

import (
	"time"
)

type EnergyRecord struct {
	ID          int64     `orm:"auto" json:"id"`
	DeviceID    int64     `orm:"index" json:"device_id"`
	RecordTime  time.Time `orm:"index" json:"record_time"`
	PowerKW     float64   `orm:"digits(10);decimals(3)" json:"power_kw"`
	EnergyKWH   float64   `orm:"digits(12);decimals(4)" json:"energy_kwh"`
	VoltageV    float64   `orm:"digits(8);decimals(2)" json:"voltage_v"`
	CurrentA    float64   `orm:"digits(8);decimals(3)" json:"current_a"`
	PowerFactor float64   `orm:"digits(5);decimals(3)" json:"power_factor"`
	CreatedAt   time.Time `orm:"auto_now_add" json:"created_at"`
}

func (e *EnergyRecord) TableName() string {
	return "energy_records"
}

type CostRecord struct {
	ID             int64     `orm:"auto" json:"id"`
	CostDate       time.Time `orm:"type(date);index" json:"cost_date"`
	Shift          string    `orm:"size(16)" json:"shift"`
	WorkOrder      string    `orm:"size(64);index" json:"work_order"`
	ProductionLine string    `orm:"size(64)" json:"production_line"`
	TotalEnergyKWH float64   `orm:"digits(12);decimals(4)" json:"total_energy_kwh"`
	UnitPrice      float64   `orm:"digits(8);decimals(4)" json:"unit_price"`
	TotalCost      float64   `orm:"digits(12);decimals(2)" json:"total_cost"`
	CreatedAt      time.Time `orm:"auto_now_add" json:"created_at"`
}

func (c *CostRecord) TableName() string {
	return "cost_records"
}

type EnergyStats struct {
	DeviceID       int64   `json:"device_id"`
	DeviceName    string  `json:"device_name"`
	TotalEnergyKWH float64 `json:"total_energy_kwh"`
	AvgPowerKW     float64 `json:"avg_power_kw"`
	MaxPowerKW     float64 `json:"max_power_kw"`
	MinPowerKW     float64 `json:"min_power_kw"`
	RecordCount    int64   `json:"record_count"`
}