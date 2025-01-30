package database

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database() {
	db, err := gorm.Open(sqlite.Open("healthtrack.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.Paciente{}, &entity.Medico{}, &entity.Recepcionista{}, &entity.Agendamento{}, &entity.Receita{}, &entity.Usuario{})
	if err != nil {
		panic(err)
	}
	super, err := entity.NewUsuario("admin@admin", "admin", "super")
	if err != nil {
		panic(err)
	}
	err = db.Create(super).Error
	if err != nil {
		panic(err)
	}
}
