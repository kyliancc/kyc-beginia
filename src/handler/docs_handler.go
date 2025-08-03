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
	var docItem model.TodoDocItem
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

func (h *DocsHandler) UpdateTodoDoc(c *gin.Context) {
	var docItem model.TodoDocItem
	if err := c.ShouldBindJSON(&docItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.docsService.UpdateTodoDoc(&docItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *DocsHandler) UpdateCpltDoc(c *gin.Context) {
	var docItem model.CpltDocItem
	if err := c.ShouldBindJSON(&docItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.docsService.UpdateCpltDoc(&docItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *DocsHandler) DeleteTodoDoc(c *gin.Context) {
	var idParam model.IDParam
	if err := c.ShouldBindUri(&idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.docsService.DeleteTodoDoc(idParam.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *DocsHandler) DeleteCpltDoc(c *gin.Context) {
	var idParam model.IDParam
	if err := c.ShouldBindUri(&idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.docsService.DeleteCpltDoc(idParam.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *DocsHandler) GetAllTodoDocs(c *gin.Context) {
	todo, err := h.docsService.GetAllTodoDocs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, todo)
}

func (h *DocsHandler) GetAllCpltDocs(c *gin.Context) {
	cplt, err := h.docsService.GetAllCpltDocs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, cplt)
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

func (h *DocsHandler) GetTodoDoc(c *gin.Context) {
	var idParam model.IDParam
	if err := c.ShouldBindUri(&idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	doc, err := h.docsService.GetTodoDoc(idParam.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, doc)
}

func (h *DocsHandler) GetCpltDoc(c *gin.Context) {
	var idParam model.IDParam
	if err := c.ShouldBindUri(&idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	doc, err := h.docsService.GetCpltDoc(idParam.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, doc)
}

func (h *DocsHandler) CompleteDoc(c *gin.Context) {
	var idParam model.IDParam
	if err := c.ShouldBindUri(&idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := h.docsService.CompleteDoc(idParam.ID)
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
