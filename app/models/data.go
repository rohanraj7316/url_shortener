package models

import "time"

// URLDataCollectionName collection name
const URLDataCollectionName = "URLData"

// URLData mongo schema
type URLData struct {
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt" validate:"required"`
	UpdateAt    time.Time `bson:"updateAt" json:"updateAt" validate:"required"`
	ExpireAt    time.Time `bson:"expireAt" json:"expireAt" validate:"required" default:""`
	URLHash     string    `bson:"urlHash" json:"urlHash" validate:"required"`
	URLOriginal string    `bson:"urlOriginal" json:"urlOriginal" validate:"required"`
	IsWebToWeb  bool      `bson:"isWebToWeb" json:"isWebToWeb" validate:"required"`
	IsWebToMob  bool      `bson:"isWebToMob" json:"isWebToMob" validate:"required"`
	IsMobToApp  bool      `bson:"isMobToApp" json:"isMobToApp" validate:"required"`
	IsAppToApp  bool      `bson:"isAppToApp" json:"isAppToApp" validate:"required"`
}
