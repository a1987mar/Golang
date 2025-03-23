package auth

import "go/adv-demo/pkg/db"

type RegRepository struct {
	Database *db.Db
}

func NewRegRepository(database *db.Db) *RegRepository {
	return &RegRepository{
		Database: database,
	}
}

func (repo *RegRepository) CreateUser(reg *Reg) (*Reg, error) {
	result := repo.Database.DB.Create(reg)
	if result.Error != nil {
		return nil, result.Error
	}
	return reg, nil
}

func (repo *RegRepository) UpdateById(reg *Reg) (*Reg, error) {
	result := repo.Database.DB.Updates(reg)
	if result.Error != nil {
		return nil, result.Error
	}
	return reg, nil
}

func (repo *RegRepository) DeleteById(id int) error {
	result := repo.Database.DB.Delete(&Reg{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
