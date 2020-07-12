package models

import "time"

// AdAnalyticsSchema mongo schema
type AdAnalyticsSchema struct {
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdateAt    time.Time `bson:"updatedAt" json:"updatedAt"`
	ExpireAt    time.Time `bson:"expireAt" json:"expireAt"`
	ProductName string    `bson:"productName" json:"productName"`
	RequestURL  string    `bson:"requestURL" json:"requestURL"`
	RedirectURL string    `bson:"redirectURL" json:"redirectURL"`
	UserAgent   string    `bson:"userAgent" json:"userAgent"`
	IP          string    `bson:"ipAddress" json:"ipAddress"`
}
