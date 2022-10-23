package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65-project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrete the schema
	database.AutoMigrate(
		&Officer{},
		&Collegeyear{},
		&Faculty{},
		&Teacher{},
		&Student{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("Prachya1234"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("Sanploy12345"), 14)

	db.Model(&Officer{}).Create(&Officer{
		Name:     "Pratchaya",
		Email:    "pratchaya@gmail.com",
		Password: string(password),
	})
	db.Model(&Officer{}).Create(&Officer{
		Name:     "Sanploy",
		Email:    "sanploy @gmail.com",
		Password: string(password2),
	})

	var Pratchaya Officer
	var Sanploy Officer
	db.Raw("SELECT * FROM users WHERE email = ?", "pratchaya@gmail.com").Scan(&Pratchaya)
	db.Raw("SELECT * FROM users WHERE email = ?", "sanploy @gmail.com").Scan(&Sanploy)

	// Resolution Data    // ของเรา Faculty Data
	science := Faculty{
		Name: "สำนักวิชาวิทยาศาสตร์",
	}
	db.Model(&Faculty{}).Create(&science)

	agricultural_technology := Faculty{
		Name: "สำนักวิชาเทคโนโลยีการเกษตร",
	}
	db.Model(&Faculty{}).Create(&agricultural_technology)

	socialtechnology := Faculty{
		Name: "สำนักวิชาเทคโนโลยีสังคม",
	}
	db.Model(&Faculty{}).Create(&socialtechnology)

	engineering := Faculty{
		Name: "สำนักวิชาวิศวกรรมศาสตร์",
	}
	db.Model(&Faculty{}).Create(&engineering)

	medicine := Faculty{
		Name: "สำนักวิชาแพทย์ศาสตร์",
	}
	db.Model(&Faculty{}).Create(&medicine)

	Nursing := Faculty{
		Name: "สำนักวิชาพยาบาลศาสตร์",
	}
	db.Model(&Faculty{}).Create(&Nursing)

	Dentist := Faculty{
		Name: "สำนักวิชาทันตแพทย์ศาสตร์",
	}
	db.Model(&Faculty{}).Create(&Dentist)

	//     // ของเรา collegeyear  Collegeyear Data
	P1 := Collegeyear{
		Name: "ปี 1",
	}
	db.Model(&Collegeyear{}).Create(&P1)

	P2 := Collegeyear{
		Name: "ปี 2",
	}
	db.Model(&Collegeyear{}).Create(&P2)

	P3 := Collegeyear{
		Name: "ปี 3",
	}
	db.Model(&Collegeyear{}).Create(&P3)

	P4 := Collegeyear{
		Name: "ปี 4",
	}
	db.Model(&Collegeyear{}).Create(&P4)

	// --- Video Data  // ของเรา Teacher Data
	T5001 := Teacher{
		Name:    "สมชาย ทันสมัย",
		Faculty: engineering,
	}
	db.Model(&Teacher{}).Create(&T5001)

	T5002 := Teacher{
		Name:    "สมหญิง จุลทล",
		Faculty: science,
	}
	db.Model(&Teacher{}).Create(&T5002)

	T5003 := Teacher{
		Name:    "มาสาย ประจำ",
		Faculty: engineering,
	}
	db.Model(&Teacher{}).Create(&T5003)

}
