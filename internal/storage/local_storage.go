package storage

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/ChinawatDc/014-go-api-uploadfile/internal/utils"
)

type LocalStorage struct {
	SingleDir       string
	MultipleDir     string
	MaxUploadMB     int64
	AllowedExt      map[string]bool
	StaticURLPrefix string // เช่น "/static"
}

func NewLocalStorage(singleDir, multipleDir string, maxUploadMB int64, allowed map[string]bool, staticPrefix string) *LocalStorage {
	return &LocalStorage{
		SingleDir:       singleDir,
		MultipleDir:     multipleDir,
		MaxUploadMB:     maxUploadMB,
		AllowedExt:      allowed,
		StaticURLPrefix: staticPrefix,
	}
}

func (s *LocalStorage) maxBytes() int64 {
	return s.MaxUploadMB * 1024 * 1024
}

func (s *LocalStorage) SaveSingle(c *gin.Context, fh *multipart.FileHeader) (utils.SavedFile, error) {
	return s.saveToDir(c, fh, s.SingleDir, "single")
}

func (s *LocalStorage) SaveMultiple(c *gin.Context, fh *multipart.FileHeader) (utils.SavedFile, error) {
	return s.saveToDir(c, fh, s.MultipleDir, "multiple")
}

func (s *LocalStorage) saveToDir(c *gin.Context, fh *multipart.FileHeader, dir string, category string) (utils.SavedFile, error) {
	if fh == nil {
		return utils.SavedFile{}, fmt.Errorf("nil file")
	}

	// per-file size check
	if fh.Size <= 0 {
		return utils.SavedFile{}, fmt.Errorf("empty file")
	}
	if fh.Size > s.maxBytes() {
		return utils.SavedFile{}, fmt.Errorf("file too large: %d bytes (max %d bytes)", fh.Size, s.maxBytes())
	}

	ext := utils.ExtLower(fh.Filename)
	if !utils.IsAllowedExt(ext, s.AllowedExt) {
		return utils.SavedFile{}, fmt.Errorf("file extension not allowed: %s", ext)
	}

	if err := utils.EnsureDir(dir); err != nil {
		return utils.SavedFile{}, err
	}

	safeName, err := utils.BuildSafeFilename(fh.Filename)
	if err != nil {
		return utils.SavedFile{}, err
	}

	dstPath := filepath.Join(dir, safeName)

	if err := c.SaveUploadedFile(fh, dstPath); err != nil {
		return utils.SavedFile{}, err
	}

	// URL ให้เข้ากับ static serve: /static/single/<filename> หรือ /static/multiple/<filename>
	url := fmt.Sprintf("%s/%s/%s", s.StaticURLPrefix, category, safeName)

	return utils.SavedFile{
		OriginalName: fh.Filename,
		SavedName:    safeName,
		Size:         fh.Size,
		Ext:          ext,
		Path:         dstPath,
		URL:          url,
		ContentType:  utils.DetectContentType(fh),
	}, nil
}
