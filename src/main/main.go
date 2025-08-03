package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/kyliancc/kyc-beginia/src/handler"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("sqlite3", "./beginia.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	docsHandler := handler.NewDocsHandler(db)

	vDocs := r.Group("/api/v1/docs")
	{
		vDocs.POST("/create", docsHandler.CreateDoc)
		vDocs.POST("/update_todo", docsHandler.UpdateTodoDoc)
		vDocs.POST("/update_cplt", docsHandler.UpdateCpltDoc)
		vDocs.DELETE("/delete_todo", docsHandler.DeleteTodoDoc)
		vDocs.DELETE("/delete_cplt", docsHandler.DeleteCpltDoc)
		vDocs.GET("/get_todo", docsHandler.GetTodoDoc)
		vDocs.GET("/get_cplt", docsHandler.GetCpltDoc)
		vDocs.GET("/get_all_todo", docsHandler.GetAllTodoDocs)
		vDocs.GET("/get_all_cplt", docsHandler.GetAllCpltDocs)
		vDocs.GET("/get_all", docsHandler.GetAllDocs)
		vDocs.POST("complete", docsHandler.CompleteDoc)
		vDocs.POST("switch", docsHandler.SwitchTodoPriority)
	}

	log.Fatal(r.Run(":8080"))
}
