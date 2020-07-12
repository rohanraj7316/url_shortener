package models

import "time"

// URLData mongo schema
type URLData struct {
	CreatedAt   time.Time `bson:"expireAt" json:"createdAt"`
	UpdateAt    time.Time `bson:"expireAt" json:"updateAt"`
	ExpireAt    time.Time `bson:"expireAt" json:"expireAt"`
	URLHash     string    `bson:"urlHash" json:"urlHash"`
	URLOriginal string    `bson:"urlOriginal" json:"urlOriginal"`
	IsWebToWeb  bool      `bson:"isWebToWeb" json:"isWebToWeb"`
	IsWebToMob  bool      `bson:"isWebToMob" json:"isWebToMob"`
	IsMobToApp  bool      `bson:"isMobToApp" json:"isMobToApp"`
	IsAppToApp  bool      `bson:"isAppToApp" json:"isAppToApp"`
}
