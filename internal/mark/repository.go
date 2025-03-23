package mark

import "go/adv-demo/pkg/db"

type MarkRepository struct {
	Database *db.Db
}

func NewMarkRepository(database *db.Db) *MarkRepository {
	return &MarkRepository{
		Database: database,
	}
}

func (c *MarkRepository) Create(mark *Mark) (*Mark, error) {
	result := c.Database.DB.Create(mark)
	if result.Error != nil {
		return nil, result.Error
	}
	return mark, nil
}

func (c *MarkRepository) DeleteForId(id int) error {
	result := c.Database.DB.Delete(&Mark{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
