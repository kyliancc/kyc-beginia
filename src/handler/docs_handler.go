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

func (h *DocsHandler) UpdateDoc(c *gin.Context) {}

func (h *DocsHandler) DeleteDoc(c *gin.Context) {}

func (h *DocsHandler) GetAllDocs(c *gin.Context) {}

func (h *DocsHandler) GetDoc(c *gin.Context) {}

func (h *DocsHandler) CompleteDoc(c *gin.Context) {}

func (h *DocsHandler) SwitchPriority(c *gin.Context) {}
