package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ChinawatDc/014-go-api-uploadfile/internal/storage"
)

type UploadHandler struct {
	Store *storage.LocalStorage
}

func NewUploadHandler(store *storage.LocalStorage) *UploadHandler {
	return &UploadHandler{Store: store}
}

// POST /upload/single  field: file
func (h *UploadHandler) UploadSingle(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "missing file field 'file'"})
		return
	}

	saved, err := h.Store.SaveSingle(c, fh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "upload failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "file": saved})
}

// POST /upload/multiple  field: files
func (h *UploadHandler) UploadMultiple(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil || form == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid multipart form"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "missing files field 'files'"})
		return
	}

	var savedFiles []any
	var errors []string

	for _, fh := range files {
		saved, e := h.Store.SaveMultiple(c, fh)
		if e != nil {
			errors = append(errors, fh.Filename+": "+e.Error())
			continue
		}
		savedFiles = append(savedFiles, saved)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"files":   savedFiles,
		"errors":  errors,
	})
}
