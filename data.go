// data.go
package main

import (
	"database/sql"
	"log"
)

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type allTask []task

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "ueeeod2kmittfzxk:giipyyFFxHqXs0yXuS0X@tcp(bkgfemq08we7wtdritzw-mysql.services.clever-cloud.com:3306)/bkgfemq08we7wtdritzw")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func getAllTasks(db *sql.DB) ([]task, error) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	tasks := []task{}
	for rows.Next() {
		var task task
		err := rows.Scan(&task.ID, &task.Name, &task.Content)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func getTaskByID(db *sql.DB, taskID int) (task, error) {
	row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", taskID)
	var t task
	err := row.Scan(&t.ID, &t.Name, &t.Content)
	if err != nil {
		log.Fatal(err)
		return task{}, err
	}
	return t, nil
}

func createNewTask(db *sql.DB, newTask task) (int, error) {
	stmt, err := db.Prepare("INSERT INTO tasks (name, content) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(newTask.Name, newTask.Content)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return int(lastInsertID), nil
}

func updateTaskByID(db *sql.DB, taskID int, updatedTask task) error {
	stmt, err := db.Prepare("UPDATE tasks SET name = ?, content = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedTask.Name, updatedTask.Content, taskID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func deleteTaskByID(db *sql.DB, taskID int) error {
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(taskID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
