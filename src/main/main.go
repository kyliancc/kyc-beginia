package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kyliancc/kyc-beginia/src/handler"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("sqlite3", "./beginia.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	docsHandler := handler.NewDocsHandler(db)

	docs := r.Group("/api/v1/docs")
	{
		docs.POST("/create", docsHandler.CreateDoc)
		docs.POST("/update_todo", docsHandler.UpdateTodoDoc)
		docs.POST("/update_cplt", docsHandler.UpdateCpltDoc)
		docs.DELETE("/delete_todo", docsHandler.DeleteTodoDoc)
		docs.DELETE("/delete_cplt", docsHandler.DeleteCpltDoc)
		docs.GET("/get_todo", docsHandler.GetTodoDoc)
		docs.GET("/get_cplt", docsHandler.GetCpltDoc)
		docs.GET("/get_all_todo", docsHandler.GetAllTodoDocs)
		docs.GET("/get_all_cplt", docsHandler.GetAllCpltDocs)
		docs.GET("/get_all", docsHandler.GetAllDocs)
		docs.POST("/complete", docsHandler.CompleteDoc)
		docs.POST("/switch", docsHandler.SwitchTodoPriority)
	}

	log.Fatal(r.Run(":8080"))
}
