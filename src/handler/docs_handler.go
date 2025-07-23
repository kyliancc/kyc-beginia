package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/kyliancc/kyc-beginia/src/model"
	"github.com/kyliancc/kyc-beginia/src/service"
	"net/http"
)

type DocsHandler struct {
	// Services objs
	docsService *service.DocsService
}

func NewDocsHandler(db *sql.DB) *DocsHandler {
	return &DocsHandler{
		docsService: service.NewDocsService(db),
	}
}

func (h *DocsHandler) CreateDoc(c *gin.Context) {
	var docItem model.DocItem
	if err := c.ShouldBindJSON(&docItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.docsService.CreateDoc(&docItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *DocsHandler) UpdateDoc(c *gin.Context) {
	var docItem model.DocItem
	if err := c.ShouldBindJSON(&docItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.docsService.UpdateDoc(&docItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *DocsHandler) DeleteDoc(c *gin.Context) {
	var docItem model.DocItem
	if err := c.ShouldBindJSON(&docItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.docsService.DeleteDoc(&docItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *DocsHandler) GetAllDocs(c *gin.Context) {
	todo, cplt, err := h.docsService.GetAllDocs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
		"cplt": cplt,
	})
}

func (h *DocsHandler) GetDoc(c *gin.Context) {
	var docItem model.DocItem
	if err := c.ShouldBindJSON(&docItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	doc, err := h.docsService.GetDoc(&docItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, doc)
}

func (h *DocsHandler) CompleteDoc(c *gin.Context) {
	var data map[string]int
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := data["id"]
	err := h.docsService.CompleteDoc(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *DocsHandler) SwitchTodoPriority(c *gin.Context) {
	var data map[string][][]int
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pairs := data["pairs"]
	err := h.docsService.SwitchTodoPriority(pairs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}
