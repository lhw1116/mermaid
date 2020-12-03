package model

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisConn *redis.Client
var Logger *logrus.Logger
