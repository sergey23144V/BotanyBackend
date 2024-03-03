package repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
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
	var (
		db  *gorm.DB
		err error
	)
	dst := "host=" + cfg.Host + " user=" + cfg.Username + " password=" + cfg.Password + " dbname=" + cfg.DBName + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode + " TimeZone=Asia/Shanghai"

	for {
		db, err = gorm.Open("postgres", dst)

		dbSQL := db.DB()

		err = dbSQL.Ping()
		if err == nil {
			log.Println("Соединение с базой данных установлено!")
			return db, nil
		}

		log.Println("Соединение с базой данных не установлено. Повторная проверка через 5 секунд...")
		time.Sleep(5 * time.Second)
	}

}
