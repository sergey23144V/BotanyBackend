package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/transect"
	trial_site "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/trial-site"
	type_plant "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/type-plant"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/user"
	"log"
	"os"
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

	db.AutoMigrate(ecomorph.EcomorphORM{}, ecomorph_entity.EcomorphsEntityORM{}, user.UserORM{}, type_plant.TypePlantORM{}, trial_site.TrialSiteORM{}, transect.TransectORM{})
	log.Println("migrant")

	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	return db, err
}
