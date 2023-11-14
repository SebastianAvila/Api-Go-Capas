// handlers.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	withDB(func(db *sql.DB) error {
		tasks, err := getAllTasks(db)
		if err != nil {
			return err
		}
		json.NewEncoder(w).Encode(tasks)
		return nil
	})(w, r)
}

func getOneTaskHandler(w http.ResponseWriter, r *http.Request) {
	withDB(func(db *sql.DB) error {
		vars := mux.Vars(r)
		taskID, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}

		t, err := getTaskByID(db, taskID)
		if err != nil {
			return err
		}

		json.NewEncoder(w).Encode(t)
		return nil
	})(w, r)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	withDB(func(db *sql.DB) error {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		var newTask task
		err = json.Unmarshal(reqBody, &newTask)
		if err != nil {
			return err
		}

		lastInsertID, err := createNewTask(db, newTask)
		if err != nil {
			return err
		}

		newTask.ID = lastInsertID
		json.NewEncoder(w).Encode(newTask)
		return nil
	})(w, r)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	withDB(func(db *sql.DB) error {
		vars := mux.Vars(r)
		taskID, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		var updatedTask task
		err = json.Unmarshal(reqBody, &updatedTask)
		if err != nil {
			return err
		}

		err = updateTaskByID(db, taskID, updatedTask)
		if err != nil {
			return err
		}

		fmt.Fprintf(w, "Tarea actualizada")
		return nil
	})(w, r)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	withDB(func(db *sql.DB) error {
		vars := mux.Vars(r)
		taskID, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}

		err = deleteTaskByID(db, taskID)
		if err != nil {
			return err
		}

		fmt.Fprintf(w, "Tarea eliminada")
		return nil
	})(w, r)
}

func indexRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi API 3")
}
