package daos

import (
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/Automation_Testing/postman/pkg/rest/server/daos/clients/sqls"
	"github.com/unnagirirao/Automation_Testing/postman/pkg/rest/server/models"
)

type Test123Dao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateTest123s(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS test123s(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Password TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewTest123Dao() (*Test123Dao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateTest123s(sqlClient)
	if err != nil {
		return nil, err
	}
	return &Test123Dao{
		sqlClient,
	}, nil
}

func (test123Dao *Test123Dao) CreateTest123(m *models.Test123) (*models.Test123, error) {
	insertQuery := "INSERT INTO test123s(Password)values(?)"
	res, err := test123Dao.sqlClient.DB.Exec(insertQuery, m.Password)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("test123 created")
	return m, nil
}

func (test123Dao *Test123Dao) UpdateTest123(id int64, m *models.Test123) (*models.Test123, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	test123, err := test123Dao.GetTest123(id)
	if err != nil {
		return nil, err
	}
	if test123 == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE test123s SET Password = ? WHERE Id = ?"
	res, err := test123Dao.sqlClient.DB.Exec(updateQuery, m.Password, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("test123 updated")
	return m, nil
}

func (test123Dao *Test123Dao) DeleteTest123(id int64) error {
	deleteQuery := "DELETE FROM test123s WHERE Id = ?"
	res, err := test123Dao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("test123 deleted")
	return nil
}

func (test123Dao *Test123Dao) ListTest123s() ([]*models.Test123, error) {
	selectQuery := "SELECT * FROM test123s"
	rows, err := test123Dao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var test123s []*models.Test123
	for rows.Next() {
		m := models.Test123{}
		if err = rows.Scan(&m.Id, &m.Password); err != nil {
			return nil, err
		}
		test123s = append(test123s, &m)
	}
	if test123s == nil {
		test123s = []*models.Test123{}
	}

	log.Debugf("test123 listed")
	return test123s, nil
}

func (test123Dao *Test123Dao) GetTest123(id int64) (*models.Test123, error) {
	selectQuery := "SELECT * FROM test123s WHERE Id = ?"
	row := test123Dao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Test123{}
	if err := row.Scan(&m.Id, &m.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("test123 retrieved")
	return &m, nil
}
