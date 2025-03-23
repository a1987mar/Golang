package link

import (
	"go/adv-demo/pkg/db"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	res := repo.Database.DB.First(&link, "hash=?", hash)
	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}

func (repo *LinkRepository) GetById(id int) (*Link, error) {
	var link Link
	res := repo.Database.DB.First(&link, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}

func (repo *LinkRepository) DeleteById(id int) error {
	res := repo.Database.DB.Delete(&Link{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (repo *LinkRepository) UpdateByHash(link *Link) (*Link, error) {
	res := repo.Database.DB.Updates(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (repo *LinkRepository) Count() int64 {
	var count int64
	repo.Database.
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	return count

}

func (repo *LinkRepository) GetALL(limit, offset uint) []Link {
	var links []Link
	repo.Database.
		Table("Links").
		Where("deleted_at is null").
		Order("is asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Scan(links)
	return links
}
