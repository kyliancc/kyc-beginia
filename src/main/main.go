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
		vDocs.GET("/create", docsHandler.CreateDoc)
	}

	log.Fatal(r.Run(":8080"))
	//todoDocRepo := repository.NewTodoDocsRepo(db)
	//
	//labels := []string{"电子", "工科"}
	//
	//doc := model.DocItem{Name: "数字电子技术基础", Priority: 2, Labels: labels}
	//id, err := todoDocRepo.CreateTodoDoc(&doc)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(id)
	//
	//docs, err := todoDocRepo.QueryAllTodoDocs()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, doc := range docs {
	//	fmt.Println(doc)
	//}
}
