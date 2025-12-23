package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ChinawatDc/014-go-api-uploadfile/internal/config"
	"github.com/ChinawatDc/014-go-api-uploadfile/internal/http/handlers"
	"github.com/ChinawatDc/014-go-api-uploadfile/internal/storage"
	"github.com/ChinawatDc/014-go-api-uploadfile/internal/utils"
)

func main() {
	cfg := config.Load()

	// ensure dirs
	if err := utils.EnsureDir(cfg.UploadSingleDir); err != nil {
		log.Fatal(err)
	}
	if err := utils.EnsureDir(cfg.UploadMultipleDir); err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// จำกัดขนาดรวม multipart ที่ gin จะเก็บใน memory ต่อ request
	r.MaxMultipartMemory = cfg.MaxUploadMB * 1024 * 1024

	// เสิร์ฟไฟล์ static:
	// /static/single/<file>
	// /static/multiple/<file>
	r.Static(cfg.StaticURLPath+"/single", cfg.UploadSingleDir)
	r.Static(cfg.StaticURLPath+"/multiple", cfg.UploadMultipleDir)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	store := storage.NewLocalStorage(
		cfg.UploadSingleDir,
		cfg.UploadMultipleDir,
		cfg.MaxUploadMB,
		cfg.AllowedExt,
		cfg.StaticURLPath,
	)

	uh := handlers.NewUploadHandler(store)

	r.POST("/upload/single", uh.UploadSingle)
	r.POST("/upload/multiple", uh.UploadMultiple)

	// optional: list allowed extensions
	r.GET("/upload/allowed", func(c *gin.Context) {
		var exts []string
		for k := range cfg.AllowedExt {
			exts = append(exts, k)
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "allowed_ext": exts, "max_upload_mb": cfg.MaxUploadMB})
	})

	log.Println("listening on :" + cfg.AppPort)
	_ = r.Run(":" + cfg.AppPort)
}
