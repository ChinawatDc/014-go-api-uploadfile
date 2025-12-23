package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	MaxUploadMB int64
	AllowedExt  map[string]bool

	UploadDir         string
	UploadSingleDir   string
	UploadMultipleDir string
	StaticURLPath     string
}

func Load() Config {
	_ = godotenv.Load(".env")

	maxMB := getenvInt64("MAX_UPLOAD_MB", 20)
	allowed := parseAllowedExt(getenv("ALLOWED_EXT", "jpg,jpeg,png,pdf,txt"))

	cfg := Config{
		AppPort: getenv("APP_PORT", "8080"),

		MaxUploadMB: maxMB,
		AllowedExt:  allowed,

		UploadDir:         getenv("UPLOAD_DIR", "uploads"),
		UploadSingleDir:   getenv("UPLOAD_SINGLE_DIR", "uploads/single"),
		UploadMultipleDir: getenv("UPLOAD_MULTIPLE_DIR", "uploads/multiple"),
		StaticURLPath:     getenv("STATIC_URL_PATH", "/static"),
	}

	log.Println("APP_PORT:", cfg.AppPort)
	log.Println("MAX_UPLOAD_MB:", cfg.MaxUploadMB)
	log.Println("UPLOAD_SINGLE_DIR:", cfg.UploadSingleDir)
	log.Println("UPLOAD_MULTIPLE_DIR:", cfg.UploadMultipleDir)
	log.Println("STATIC_URL_PATH:", cfg.StaticURLPath)

	return cfg
}

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}

func getenvInt64(k string, def int64) int64 {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Printf("invalid %s, using default=%d", k, def)
		return def
	}
	return n
}

func parseAllowedExt(csv string) map[string]bool {
	m := map[string]bool{}
	for _, p := range strings.Split(csv, ",") {
		p = strings.TrimSpace(strings.ToLower(p))
		if p == "" {
			continue
		}
		p = strings.TrimPrefix(p, ".")
		m[p] = true
	}
	return m
}
