package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huhx/common-go/logger"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

var (
	Instance *gorm.DB
	Redis    *redis.Client
)

func InitDb() {
	if gin.Mode() != gin.ReleaseMode {
		if err := godotenv.Load(); err != nil {
			log := logger.NewLogger()
			log.Warn("Can not load .env file.", zap.String("env", os.Getenv("GIN_MODE")))
		}
	}

	dbHost := os.Getenv("db_host")
	dbPort := "5432"
	dbUser := os.Getenv("db_username")
	dbPassword := os.Getenv("db_password")
	dbName := os.Getenv("db_name")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger:                                   gormLogger.Default.LogMode(gormLogger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("无法连接到数据库：" + err.Error())
	}

	// hook
	_ = db.Callback().Query().Before("gorm:query").Register("BeforeQuery", BeforeQueryCallback)
	_ = db.Callback().Create().Before("gorm:create").Register("BeforeCreate", BeforeCreateCallback)
	_ = db.Callback().Update().Before("gorm:update").Register("BeforeUpdate", BeforeUpdateCallback)

	Instance = db
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_host"),
		Password: os.Getenv("redis_password"),
		DB:       0, // 使用默认的数据库
	})
}
