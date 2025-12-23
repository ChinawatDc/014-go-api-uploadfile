package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}

func ExtLower(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	return strings.TrimPrefix(ext, ".")
}

func IsAllowedExt(ext string, allowed map[string]bool) bool {
	if ext == "" {
		return false
	}
	return allowed[strings.ToLower(ext)]
}

func RandomHex(nBytes int) (string, error) {
	b := make([]byte, nBytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func BuildSafeFilename(originalName string) (string, error) {
	base := filepath.Base(originalName) // กัน path traversal
	base = sanitizeFilename(base)

	r, err := RandomHex(6)
	if err != nil {
		return "", err
	}

	ts := time.Now().Format("20060102_150405")
	return fmt.Sprintf("%s_%s_%s", ts, r, base), nil
}

func sanitizeFilename(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "..", ".")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "\\", "_")
	return name
}

type SavedFile struct {
	OriginalName string `json:"original_name"`
	SavedName    string `json:"saved_name"`
	Size         int64  `json:"size"`
	Ext          string `json:"ext"`
	Path         string `json:"path"`
	URL          string `json:"url"`
	ContentType  string `json:"content_type"`
}

func DetectContentType(fh *multipart.FileHeader) string {
	if fh.Header != nil {
		if v := fh.Header.Get("Content-Type"); v != "" {
			return v
		}
	}
	return "application/octet-stream"
}
