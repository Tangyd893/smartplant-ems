package models

import (
	"time"
)

type Device struct {
	ID          int64     `orm:"auto" json:"id"`
	DeviceCode  string    `orm:"size(64);unique" json:"device_code"`
	DeviceName  string    `orm:"size(128)" json:"device_name"`
	DeviceType  string    `orm:"size(32)" json:"device_type"`
	Location    string    `orm:"size(128)" json:"location"`
	Status      int       `orm:"default(1)" json:"status"`
	PowerRating float64   `orm:"digits(10);decimals(2)" json:"power_rating"`
	CreatedAt   time.Time `orm:"auto_now_add" json:"created_at"`
	UpdatedAt   time.Time `orm:"auto_now" json:"updated_at"`
}

func (d *Device) TableName() string {
	return "devices"
}

type User struct {
	ID           int64     `orm:"auto" json:"id"`
	Username     string    `orm:"size(64);unique" json:"username"`
	PasswordHash string    `orm:"size(256)" json:"-"`
	RealName     string    `orm:"size(64)" json:"real_name"`
	Role         string    `orm:"size(16)" json:"role"`
	Phone        string    `orm:"size(32)" json:"phone"`
	Email        string    `orm:"size(128)" json:"email"`
	Status       int       `orm:"default(1)" json:"status"`
	CreatedAt    time.Time `orm:"auto_now_add" json:"created_at"`
	UpdatedAt    time.Time `orm:"auto_now" json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
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