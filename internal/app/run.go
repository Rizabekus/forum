package app

import (
	"database/sql"
	"forum/internal/controllers"
	"forum/internal/repo"
	"forum/internal/services"
	"log"
)

func Run() {
	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repo.RepositoryInstance(db)
	service := services.ServiceInstance(repo)
	controller := controllers.ControllersInstance(service)
	Routes(controller)
	defer db.Close()
}
