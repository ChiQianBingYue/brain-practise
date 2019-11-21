package playground

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type statusType uint8

const (
	doing statusType = iota // default
	done
	drop
	plan
)

type userWorkProfile struct {
	gorm.Model
	UserID           uint      `gorm:"not null;index"`
	Hours            int       `gorm:"default:0"`
	Goals            []goal    `gorm:"foreignkey:UserWorkProfileID"`
	IndepentMissions []mission `gorm:"foreignkey:UserWorkProfileID"`
	Records          []record  `gorm:"foreignkey:UserWorkProfileID"`
}

type goal struct {
	gorm.Model
	UserWorkProfileID uint        `gorm:"not null;index:idx_status"`
	Status            *statusType `gorm:"index:idx_status;default:0"`
	Title             *string     `gorm:"not null;size:80"`
	Result            string      `gorm:"size:255"`
	Hours             int         `gorm:"not null;default:0"`
	Missions          []mission   `gorm:"foreignkey:GoalID"`
	Records           []record    `gorm:"foreignkey:GoalID"`
}

type mission struct {
	gorm.Model
	Title             sql.NullString `gorm:"not null"` // 0 进行中, 1 已完成, 2 已放弃
	Description       string         `gorm:"size:255"`
	Status            uint           `sql:"index"`
	UserWorkProfileID uint           `gorm:"not null,index"`
	Hours             int
	Goal              goal
	GoalID            uint
	Records           []record `gorm:"foreignkey:MissionID"`
}

// Record 记录的model
type record struct {
	gorm.Model
	Content           sql.NullString `gorm:"not null;size:255"`
	Thoughts          sql.NullString `gorm:"size:255"`
	Hours             int
	UserWorkProfileID uint `gorm:"not null;index:scope_idx"`
	Goal              goal
	GoalID            *uint `gorm:"index:scope_idx"`
	Mission           mission
	MissionID         *uint `gorm:"index:scope_idx"`
}

func PlayGorm() {
	db, err := gorm.Open("mysql", "qy:871127@tcp(localhost:3306)/everest_test?parseTime=true")
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// db.AutoMigrate(&userWorkProfile{})
	// db.AutoMigrate(&goal{})
	// db.AutoMigrate(&mission{})
	// db.AutoMigrate(&record{})
	// // db.Delete(g)
	// // str := ""
	// st := done
	// g := goal{Title: nil, Result: "fdsaf", Status: &st}
	// g.ID = 1
	// db.Model(&g).Updates(g)
	// for _, r := range g.Records {

	// }
}
