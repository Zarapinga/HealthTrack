package database

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"gorm.io/gorm"
)

type Recepcionista struct {
	DB *gorm.DB
}

func NewRecepcionista(db *gorm.DB) *Recepcionista {
	return &Recepcionista{DB: db}
}

func (r *Recepcionista) CreateRecepcionista(recepcionista *entity.Recepcionista) error {
	return r.DB.Create(recepcionista).Error
}

func (r *Recepcionista) FindAllRecepcionista() ([]entity.Recepcionista, error) {
	var recepcionista []entity.Recepcionista
	err := r.DB.Find(&recepcionista).Error
	return recepcionista, err
}

func (r *Recepcionista) FindByRecepcionistaEmail(email string) (*entity.Recepcionista, error) {
	var recepcionista entity.Recepcionista
	err := r.DB.First(&recepcionista, "email = ?", email).Error
	return &recepcionista, err
}

func (r *Recepcionista) UpdateRecepcionista(recepcionista *entity.Recepcionista) error {
	_, err := r.FindByRecepcionistaEmail(recepcionista.Email)
	if err != nil {
		return err
	}
	return r.DB.Save(recepcionista).Error
}

func (r *Recepcionista) DeleteRecepcionista(email string) error {
	recepcionista, err := r.FindByRecepcionistaEmail(email)
	if err != nil {
		return err
	}
	return r.DB.Delete(recepcionista).Error
}
