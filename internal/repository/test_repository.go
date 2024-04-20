package repository

import (
	"github.com/antony0016/sw-api/utils/logger"
	"gorm.io/gorm"
)

type TestRepository struct {
	db *gorm.DB
}

func NewTestRepository(db *gorm.DB) *TestRepository {
	return &TestRepository{db: db}
}

func (tr *TestRepository) TestInsert(id int, test string) (count int, err error) {
	// tr.db.Model(&Test{}).Create(&Test{Test: "test"})
	logger.Printf("Test insert id: %d, test: %s", id, test)
	return 1, nil
}
