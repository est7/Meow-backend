package initialize

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/pkg/log"
	"database/sql"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	viper2 "github.com/spf13/viper"
	"gorm.io/gorm"
)

var ConfigPath = "config/"

type AppEnvConfig struct {
	Name              string          `mapstructure:"Name"`
	Version           string          `mapstructure:"Version"`
	Port              string          `mapstructure:"Port"`
	PprofPort         string          `mapstructure:"PprofPort"`
	Mode              interfaces.Mode `mapstructure:"Mode"`
	CookieName        string          `mapstructure:"CookieName"`
	SSL               bool            `mapstructure:"SSL"`
	CtxDefaultTimeout int             `mapstructure:"CtxDefaultTimeout"`
	CSRF              bool            `mapstructure:"CSRF"`
	Debug             bool            `mapstructure:"Debug"`
	EnableTrace       bool            `mapstructure:"EnableTrace"`
	EnablePprof       bool            `mapstructure:"EnablePprof"`
	PGConfig          PGConfig        `mapstructure:"DB"`
	RedisConfig       RedisConfig     `mapstructure:"Redis"`
	JWTConfig         JWTConfig       `mapstructure:"Jwt"`
	OTelConfig        OTelConfig      `mapstructure:"OTel"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"JwtSecret"`
	ExpiresIn int    `mapstructure:"JwtTimeout"`
}

type OTelConfig struct {
	Endpoint string `mapstructure:"Endpoint"`
	Insecure bool   `mapstructure:"Insecure"`
}

var (
	EnvConfig AppEnvConfig
)

type AppInstance struct {
	Db          *sql.DB
	GormDb      *gorm.DB
	RedisClient *redis.Client
}

type AppOption func(*AppInstance)

func WithDB(db *sql.DB) AppOption {
	return func(a *AppInstance) {
		a.Db = db
	}
}

func WithGormDB(gormDb *gorm.DB) AppOption {
	return func(a *AppInstance) {
		a.GormDb = gormDb
	}
}

func WithRedisClient(redisClient *redis.Client) AppOption {
	return func(a *AppInstance) {
		a.RedisClient = redisClient
	}
}

func NewAppInstance(opts ...AppOption) *AppInstance {
	app := &AppInstance{}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

func (a *AppInstance) GetDB() *sql.DB {
	return a.Db
}

func (a *AppInstance) GetGormDB() *gorm.DB {
	return a.GormDb
}

func (a *AppInstance) GetRedisClient() *redis.Client {
	return a.RedisClient
}

var AppCtxInstance interfaces.AppContext = &AppInstance{
	Db:          nil,
	GormDb:      nil,
	RedisClient: nil,
}

func LoadConfig(path string) AppEnvConfig {
	// 从环境变量中获取配置
	viper := viper2.GetViper()
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper2.ConfigFileNotFoundError

		if errors.As(err, &configFileNotFoundError) {
			log.Fatalf("Config file not found: %v", err)
		} else {
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	// 将配置文件映射到结构体
	if err := viper.Unmarshal(&EnvConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	fmt.Printf("Loaded config: %+v\n", EnvConfig)

	return EnvConfig
}

func LoadLoggerConfig(path string, mode interfaces.Mode) interfaces.Logger {
	// 从环境变量中获取配置
	viper := viper2.GetViper()
	viper.AddConfigPath(path)
	viper.SetConfigName("logger")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	var err error
	if err = viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper2.ConfigFileNotFoundError

		if errors.As(err, &configFileNotFoundError) {
			log.Fatalf("Config file not found: %v", err)
		} else {
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	// 将配置文件映射到结构体
	var loggerConfig log.LoggerConfig
	if err = viper.Unmarshal(&loggerConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	fmt.Printf("Loaded config: %+v\n", loggerConfig)

	zapLogger := initZapLogger(&loggerConfig, mode)
	return zapLogger
}

func initZapLogger(cfg *log.LoggerConfig, mode interfaces.Mode) interfaces.Logger {
	return log.InitZapLogger(cfg, mode)
}

/**
* 在Go语言中，init 函数有特殊的行为。每个包可以包含多个 init 函数，这些函数在包被导入时自动执行，
* 且执行顺序是在所有全局变量声明之后，main 函数执行之前。你不需要显式地调用 init 函数；它们由Go运行时自动调用。
* 这使得 init 函数成为进行初始化设置，如设置全局变量、初始化数据库连接等操作的理想选择。
* 在你的项目中，init 函数定义在 main.go 文件中。当程序启动并导入包含 init 函数的包时，
* Go运行时会自动执行这些 init 函数。这意味着，尽管你在代码中没有看到直接的调用，init 函数仍然会在程序启动时执行。
 */
func init() {
	//config, err := LoadConfig("config/config.yaml")
	//if err != nil {
	//	log.Fatalf("Failed to initialize config: %v", err)
	//}
}
