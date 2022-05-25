package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

// 假设在 dao 层中遇到一个 sql.ErrNoRows ，是否应该 Wrap 这个 error，抛给上层
// 无需 wrap，如有需要，底层打个日志即可，实际处理时可返回 nil 数据
func init() {
	var err error
	db, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
}

type school struct {
	id   int
	name string
}

type student struct {
	id   int
	name string
}

func main() {
	var err = db.Ping()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer db.Close()

	_, err = getStudent(1)
	if err != nil {
		log.Print(err)
	} else {
		log.Print("get Student success")
	}
	_, err = getSchool(1)
	if err != nil {
		log.Print(err)
	} else {
		log.Print("get School success")
	}
	getFromErrTable()
}

func getStudent(studentId int) (*student, error) {
	var (
		id   int
		name string
	)
	err := db.QueryRow("select * from student where id = ?", studentId).Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("no student")
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &student{
		id:   id,
		name: name,
	}, nil
}

func getSchool(schoolId int) (*school, error) {
	var (
		id   int
		name string
	)
	err := db.QueryRow("select * from school where id = ?", schoolId).Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("no school")
			return nil, nil
		} else {
			return nil, err
		}
	}
	log.Printf("get school id: %d, name: %s", id, name)

	return &school{
		id:   id,
		name: name,
	}, nil
}

func getFromErrTable() {
	var (
		id   int
		name string
	)
	err := db.QueryRow("select * from studet").Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("no student")
		} else {
			log.Fatal(err)
		}
	}
}
