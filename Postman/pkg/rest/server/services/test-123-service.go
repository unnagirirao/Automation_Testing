package services

import (
	"github.com/unnagirirao/Automation_Testing/postman/pkg/rest/server/daos"
	"github.com/unnagirirao/Automation_Testing/postman/pkg/rest/server/models"
)

type Test123Service struct {
	test123Dao *daos.Test123Dao
}

func NewTest123Service() (*Test123Service, error) {
	test123Dao, err := daos.NewTest123Dao()
	if err != nil {
		return nil, err
	}
	return &Test123Service{
		test123Dao: test123Dao,
	}, nil
}

func (test123Service *Test123Service) CreateTest123(test123 *models.Test123) (*models.Test123, error) {
	return test123Service.test123Dao.CreateTest123(test123)
}

func (test123Service *Test123Service) UpdateTest123(id int64, test123 *models.Test123) (*models.Test123, error) {
	return test123Service.test123Dao.UpdateTest123(id, test123)
}

func (test123Service *Test123Service) DeleteTest123(id int64) error {
	return test123Service.test123Dao.DeleteTest123(id)
}

func (test123Service *Test123Service) ListTest123s() ([]*models.Test123, error) {
	return test123Service.test123Dao.ListTest123s()
}

func (test123Service *Test123Service) GetTest123(id int64) (*models.Test123, error) {
	return test123Service.test123Dao.GetTest123(id)
}
