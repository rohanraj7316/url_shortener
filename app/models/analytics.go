package models

import "time"

// AdAnalyticsSchema mongo schema
type AdAnalyticsSchema struct {
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	ProductName string    `bson:"productName" json:"productName"`
	URLHash     string    `bson:"requestURL" json:"requestURL"`
	UserAgent   string    `bson:"userAgent" json:"userAgent"`
	IP          string    `bson:"ipAddress" json:"ipAddress"`
}
