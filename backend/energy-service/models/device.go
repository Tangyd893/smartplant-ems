package models

import (
	"time"
)

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
}

type PageResult struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
}

type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}