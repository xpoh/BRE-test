package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	swagger "github.com/xpoh/BRE-test/cmd/go"
)

const (
	host     = "10.10.10.100"
	port     = 5432
	user     = "user"
	password = "user"
	dbname   = "db"
)

type dbStorage struct {
}

func (d *dbStorage) AddToStorage(content swagger.Content) error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	defer db.Close()

	s := `insert into service values($1, $2, $3)`
	_, err = db.Exec(s, content.Id, content.Name, content.Body)
	if err != nil {
		return err
	}
	return err
}

func (d *dbStorage) ReadFromStorage(id string) (string, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return "", err
	}

	defer db.Close()

	rows, err1 := db.Query(`SELECT content FROM service WHERE id=$1`, id)
	if err1 != nil {
		return "", err
	}
	for rows.Next() {
		var ans string
		err = rows.Scan(&ans)
		if err == nil {
			return ans, nil
		} else {
			break
		}
	}
	return "", err
}

func (d *dbStorage) RemoveFromStorage(id string) error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	defer db.Close()

	s := `delete from service where id=$1`
	_, err = db.Exec(s, id)
	if err != nil {
		return err
	}
	return err
}

func NewDbStorage() *dbStorage {
	return &dbStorage{}
}
