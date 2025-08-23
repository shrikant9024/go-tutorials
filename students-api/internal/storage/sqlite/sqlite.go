package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shrikant9024/go-tutorials/internal/config"
	"github.com/shrikant9024/go-tutorials/internal/types"
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

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM STUDENTS WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {

		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found witht the id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error %w", err)
	}

	return student, nil
}

func (s *Sqlite) GetStudentList() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM STUDENTS")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var item types.Student

		err := rows.Scan(&item.Id, &item.Name, &item.Email, &item.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, item)

	}

	return students, nil

}
