package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
	"log"
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

	dst := "host=" + cfg.Host + " user=" + cfg.Username + " password=" + cfg.Password + " dbname=" + cfg.DBName + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode + " TimeZone=Asia/Shanghai"

	db, err := gorm.Open("postgres", dst)
	log.Println("Open db")

	db.AutoMigrate(ecomorph_entity.EcomorphsEntity{}, ecomorph.Ecomorph{})
	log.Println("migrant")

	return db, err
}
