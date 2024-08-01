package initialize

import (
	"context"
	"database/sql"
	"fmt"
)

type PGConfig struct {
	Driver          string `mapstructure:"Driver"`
	Name            string `mapstructure:"Name"`
	Host            string `mapstructure:"Host"`
	Port            int    `mapstructure:"Port"`
	UserName        string `mapstructure:"UserName"`
	Password        string `mapstructure:"Password"`
	ShowLog         bool   `mapstructure:"ShowLog"`
	MaxIdleConn     int    `mapstructure:"MaxIdleConn"`
	MaxOpenConn     int    `mapstructure:"MaxOpenConn"`
	Timeout         int    `mapstructure:"Timeout"`
	ReadTimeout     int    `mapstructure:"ReadTimeout"`
	ConnMaxLifeTime string `mapstructure:"ConnMaxLifeTime"`
	SlowThreshold   string `mapstructure:"SlowThreshold"`
}

func InitDB(dbConfig PGConfig, ctx context.Context) (*sql.DB, error) {
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
		db.Close() // 关闭连接
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func CloseDB(db *sql.DB) error {
	return db.Close()
}
