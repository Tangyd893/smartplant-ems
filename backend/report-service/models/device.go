package models

import (
	"time"
)

<<<<<<< HEAD
type Report struct {
	ID          int64     `orm:"auto" json:"id"`
	ReportName  string    `orm:"size(128)" json:"report_name"`
	ReportType  string    `orm:"size(32)" json:"report_type"`
	PeriodType  string    `orm:"size(16)" json:"period_type"`
	PeriodStart time.Time `orm:"type(date)" json:"period_start"`
	PeriodEnd   time.Time `orm:"type(date)" json:"period_end"`
	CreatedBy   int64     `json:"created_by"`
	FilePath    string    `orm:"size(256)" json:"file_path"`
	Status      int       `orm:"default(0)" json:"status"`
	CreatedAt   time.Time `orm:"auto_now_add" json:"created_at"`
	CompletedAt time.Time `orm:"null" json:"completed_at"`
}

func (r *Report) TableName() string {
	return "reports"
}

type Alert struct {
	ID             int64      `orm:"auto" json:"id"`
	DeviceID       int64      `orm:"index" json:"device_id"`
	AlertType      string     `orm:"size(32)" json:"alert_type"`
	AlertLevel     int        `json:"alert_level"`
	AlertMsg       string     `orm:"type(text)" json:"alert_msg"`
	ThresholdValue float64    `orm:"digits(10);decimals(3)" json:"threshold_value"`
	CurrentValue   float64    `orm:"digits(10);decimals(3)" json:"current_value"`
	IsResolved     int        `orm:"default(0)" json:"is_resolved"`
	ResolvedAt     *time.Time `orm:"null" json:"resolved_at"`
	ResolvedBy     string     `orm:"size(64)" json:"resolved_by"`
	CreatedAt      time.Time  `orm:"auto_now_add" json:"created_at"`
}

func (a *Alert) TableName() string {
	return "alerts"
}

type ReportCreate struct {
	ReportName  string `json:"report_name" valid:"Required"`
	ReportType  string `json:"report_type" valid:"Required"`
	PeriodType  string `json:"period_type" valid:"Required"`
	PeriodStart string `json:"period_start" valid:"Required"`
	PeriodEnd   string `json:"period_end" valid:"Required"`
=======
type Device struct {
	ID           int64     `json:"id"`
	DeviceCode   string    `json:"device_code"`
	DeviceName   string    `json:"device_name"`
	DeviceType   string    `json:"device_type"`
	Location     string    `json:"location"`
	Status       int       `json:"status"`
	PowerRating  float64   `json:"power_rating"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DeviceCreate struct {
	DeviceCode  string  `json:"device_code" valid:"Required"`
	DeviceName  string  `json:"device_name" valid:"Required"`
	DeviceType  string  `json:"device_type" valid:"Required"`
	Location    string  `json:"location"`
	PowerRating float64 `json:"power_rating"`
}

type DeviceUpdate struct {
	ID          int64   `json:"id" valid:"Required"`
	DeviceName  string  `json:"device_name"`
	Location    string  `json:"location"`
	Status      int     `json:"status"`
	PowerRating float64 `json:"power_rating"`
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
}

type PageResult struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
<<<<<<< HEAD
=======
}

type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
}