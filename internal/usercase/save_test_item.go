package usercase

import (
	"github.com/antony0016/sw-api/internal/repository"
	"github.com/antony0016/sw-api/utils/logger"
)

func SaveTestItem(testRepository *repository.TestRepository, id int, test string) (success bool) {
	count, err := testRepository.TestInsert(id, test)

	if err != nil {
		logger.Panic("Can't save test item.", err)
	} else {
		logger.Println("Save test item success.")
	}
	return count == 1
}
