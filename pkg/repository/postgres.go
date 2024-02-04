package repository

import (
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config содержит необходимую конфигурацию для подключения к базе данных
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectDB(cfg Config) (*gorm.DB, error) {

	dsn := "host=" + cfg.Host + " user=" + cfg.Username + " password=" + cfg.Password + " dbname=" + cfg.DBName + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode + " TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(g_rpc.EcomorphsEntity{}, g_rpc.Ecomorph{})
	println("log")

	return db, err
}
