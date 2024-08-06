package initialize

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type PGConfig struct {
	Driver          string        `mapstructure:"Driver"`
	Name            string        `mapstructure:"Name"`
	Host            string        `mapstructure:"Host"`
	Port            int           `mapstructure:"Port"`
	UserName        string        `mapstructure:"UserName"`
	Password        string        `mapstructure:"Password"`
	ShowLog         bool          `mapstructure:"ShowLog"`
	MaxIdleConn     int           `mapstructure:"MaxIdleConn"`
	MaxOpenConn     int           `mapstructure:"MaxOpenConn"`
	Timeout         int           `mapstructure:"Timeout"`
	ReadTimeout     int           `mapstructure:"ReadTimeout"`
	ConnMaxLifeTime time.Duration `mapstructure:"ConnMaxLifeTime"`
	SlowThreshold   time.Duration `mapstructure:"SlowThreshold"`
}

func InitDB(dbConfig PGConfig, ctx context.Context) (*gorm.DB, *sql.DB, error) {
	// 初始化数据库连接
	dbDSN := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.UserName, dbConfig.Password, dbConfig.Name)

	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicf("open db failed. Host: %s, database name: %s, err: %+v", dbConfig.Host, dbConfig.Name, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panicf("database connection failed. database name: %s, err: %+v", dbConfig.Name, err)
	}
	// 设置数据库连接池参数
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)

	sqlDB.SetConnMaxLifetime(dbConfig.ConnMaxLifeTime)

	return db, sqlDB, nil
}

func PingTest(dbConfig PGConfig, ctx context.Context) (*sql.DB, error) {
	// 初始化数据库连接
	dbDSN := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.UserName, dbConfig.Password, dbConfig.Name)

	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// 设置数据库连接池参数
	db.SetMaxIdleConns(dbConfig.MaxIdleConn)
	db.SetMaxOpenConns(dbConfig.MaxOpenConn)

	// 测试数据库连接
	if err := db.PingContext(ctx); err != nil {
		log.Printf("PingContext failed: %v", err)
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func CloseDB(db *sql.DB) error {
	return db.Close()
}
