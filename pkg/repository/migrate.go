package repository

import (
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"gorm.io/gorm"
	"log"
)

func MigrateUser(db *gorm.DB) error {
	ormObj := api.UserORM{
		Id:       "9e958acb-5d06-8cbb-f204-71dc57529edd",
		Name:     "Admin",
		Email:    "Admin",
		Password: "686a7172686a7177313234363137616a6668616a734e7afebcfbae000b22c7c85e5560f89a2a0280b4",
		Role:     0,
	}
	if err := db.Create(&ormObj).Error; err != nil {
		return err
	}
	return nil
}

func MigrateEcomorph(db *gorm.DB) error {

	Ecomorphs := []api.EcomorphORM{
		{Id: "9d5a13c6-cc08-415e-bcb7-c96f3e0c040d", Title: "Гелиоморфы"},
		{Id: "caf44071-29ad-4dce-87cc-107e77b7e3aa", Title: "Трофоморфы"},
		{Id: "5d024049-0256-4992-9e24-b92701ea57ed", Title: "Гигроморфы"},
		{Id: "f1548852-c8d6-4ef1-a3e4-6130e9528559", Title: "Термоморфы"}}

	for _, ormObj := range Ecomorphs {

		if err := db.Create(&ormObj).Error; err != nil {
			return err
		}
	}

	return nil
}

func MigrateEcomorphsEntity(db *gorm.DB) error {

	Ecomorphs := []api.EcomorphsEntityORM{
		{Id: "de0ed4eb-e0c3-47ca-8261-56cbb535e79e", Title: "Ультрасциофиты", Score: 0, DisplayTable: "USc", Ecomorphs: &api.EcomorphORM{Id: "9d5a13c6-cc08-415e-bcb7-c96f3e0c040d"}},
		{Id: "b2183bc3-a897-4cf7-af47-979f7104f764", Title: "Сциофиты", Score: 1, DisplayTable: "Sc", Ecomorphs: &api.EcomorphORM{Id: "9d5a13c6-cc08-415e-bcb7-c96f3e0c040d"}},
		{Id: "ae530d0b-fb0d-4f52-ad1d-a03e9d8c8616", Title: "Гелиосциофиты", Score: 2, DisplayTable: "HeSc", Ecomorphs: &api.EcomorphORM{Id: "9d5a13c6-cc08-415e-bcb7-c96f3e0c040d"}},
		{Id: "68661247-9fe6-4ef6-b1fb-670bd736a8d6", Title: "Сциогелиофиты", Score: 3, DisplayTable: "ScHe", Ecomorphs: &api.EcomorphORM{Id: "9d5a13c6-cc08-415e-bcb7-c96f3e0c040d"}},
		{Id: "9b5f3a4f-0490-4abe-8f99-4007e3624061", Title: "Гелиофиты", Score: 4, DisplayTable: "He", Ecomorphs: &api.EcomorphORM{Id: "9d5a13c6-cc08-415e-bcb7-c96f3e0c040d"}},
		{Id: "6e2e7980-0c84-43de-be67-b820defb2213", Title: "Ультраолиготрофы", Score: 5, DisplayTable: "UHe", Ecomorphs: &api.EcomorphORM{Id: "9d5a13c6-cc08-415e-bcb7-c96f3e0c040d"}},

		{Id: "cc1771e8-197d-4c8c-828e-13aea16dd276", Title: "Ультраолиготрофы", Score: 0, DisplayTable: "UOgTr", Ecomorphs: &api.EcomorphORM{Id: "caf44071-29ad-4dce-87cc-107e77b7e3aa"}},
		{Id: "467312b0-3db0-4dd7-be7d-eba7879056a4", Title: "Олиготрофы", Score: 1, DisplayTable: "OgTr", Ecomorphs: &api.EcomorphORM{Id: "caf44071-29ad-4dce-87cc-107e77b7e3aa"}},
		{Id: "6787c6f2-0ad4-4c44-8c85-34f444e19fa0", Title: "Мезотрофы", Score: 2, DisplayTable: "MsTr", Ecomorphs: &api.EcomorphORM{Id: "caf44071-29ad-4dce-87cc-107e77b7e3aa"}},
		{Id: "b2f0ce81-3212-445e-919c-17c7b3e791ce", Title: "Мегатрофы", Score: 3, DisplayTable: "MgTr", Ecomorphs: &api.EcomorphORM{Id: "caf44071-29ad-4dce-87cc-107e77b7e3aa"}},
		{Id: "ad273087-1772-4e76-acae-d93fa33688da", Title: "Галомегатрофы", Score: 4, DisplayTable: "HMgTr", Ecomorphs: &api.EcomorphORM{Id: "caf44071-29ad-4dce-87cc-107e77b7e3aa"}},
		{Id: "77c711ac-1d63-4061-b69b-c3bcaf32fb9c", Title: "Галофиты", Score: 5, DisplayTable: "Hal", Ecomorphs: &api.EcomorphORM{Id: "caf44071-29ad-4dce-87cc-107e77b7e3aa"}},

		{Id: "5cb29094-f94e-43cf-a5b4-c9d7bc92d715", Title: "Ксерофиты", Score: 0, DisplayTable: "Ks", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},
		{Id: "5dba0f42-e01b-4525-b8c0-023a86bcc961", Title: "Мезоксерофиты", Score: 1, DisplayTable: "MsKs", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},
		{Id: "ad36a4be-62e3-4fbe-8c7f-5cf2ce9151ee", Title: "Ксеромезофиты", Score: 1, DisplayTable: "KsMs", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},
		{Id: "ae81710b-c6cb-4fa3-b103-307bd39f6735", Title: "Мезофиты", Score: 2, DisplayTable: "Ms", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},
		{Id: "4f553174-1bd8-4ae1-a0ef-079ac75e64b3", Title: "Гигромезофиты", Score: 2, DisplayTable: "HgrMs", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},
		{Id: "74c99e53-f58d-494f-847c-70387e8f98f7", Title: "Мезогигрофиты", Score: 3, DisplayTable: "MsHgr", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},
		{Id: "61cb47a4-24e9-48fb-adf4-5ee762524fb2", Title: "Гигрофиты", Score: 4, DisplayTable: "Hgr", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},
		{Id: "54887fdf-9835-4e12-a57e-85cc4495f3ac", Title: "Ультрагигрофиты", Score: 5, DisplayTable: "UHgr", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},
		{Id: "a1b0c2ee-747d-461b-9a98-a4469c178dd0", Title: "Гидрофиты", Score: 6, DisplayTable: "Hd", Ecomorphs: &api.EcomorphORM{Id: "5d024049-0256-4992-9e24-b92701ea57ed"}},

		{Id: "d0967720-ae2d-4328-9176-fc6478c45a2b", Title: "Ультраолиготермы ", Score: 1, DisplayTable: "UOgT", Ecomorphs: &api.EcomorphORM{Id: "f1548852-c8d6-4ef1-a3e4-6130e9528559"}},
		{Id: "0e969b2f-14cc-423e-aa31-0c3a9e8d037d", Title: "Олиготермы ", Score: 2, DisplayTable: "OgT", Ecomorphs: &api.EcomorphORM{Id: "f1548852-c8d6-4ef1-a3e4-6130e9528559"}},
		{Id: "31f9837d-1042-45a2-95b2-2aebc13069f0", Title: "Мезотермы ", Score: 3, DisplayTable: "MsT", Ecomorphs: &api.EcomorphORM{Id: "f1548852-c8d6-4ef1-a3e4-6130e9528559"}},
		{Id: "f7ba4fd3-0540-4409-8f2b-53cc6c9805c0", Title: "Мегатермы ", Score: 4, DisplayTable: "MgT", Ecomorphs: &api.EcomorphORM{Id: "f1548852-c8d6-4ef1-a3e4-6130e9528559"}},
		{Id: "b80678f7-de0e-4fe7-a10a-196a501b349f", Title: "Ультрамегатермы", Score: 5, DisplayTable: "UMgT", Ecomorphs: &api.EcomorphORM{Id: "f1548852-c8d6-4ef1-a3e4-6130e9528559"}}}

	for _, ormObj := range Ecomorphs {

		if err := db.Create(&ormObj).Error; err != nil {
			return err
		}
	}

	return nil
}

func Migrate(db *gorm.DB) {

	db.AutoMigrate(api.ImgORM{}, api.EcomorphORM{}, api.EcomorphsEntityORM{}, api.UserORM{}, api.TypePlantORM{}, api.TransectORM{}, api.TrialSiteORM{}, api.PlantORM{}, api.AnalysisORM{})
	MigrateUser(db)
	MigrateEcomorph(db)
	MigrateEcomorphsEntity(db)
	log.Println("migrant")

}
