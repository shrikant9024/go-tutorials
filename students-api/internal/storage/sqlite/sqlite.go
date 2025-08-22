package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shrikant9024/go-tutorials/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.Storage_path)
	if err != nil {
		return nil, err
	}

	// query := `CREATE TABLE IF NOT EXISTS students(
	// id INTEGER PRIMARY KEY AUTOINCREMENT,
	// name TEXT,
	// EMAIL TEXT,
	// age INTEGER,
	// );`

	// fmt.Println(query)

	// _, err = db.Exec(query)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	EMAIL TEXT,
	age INTEGER
	);`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil

}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil

}
