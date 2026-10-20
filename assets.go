package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func getAssetPath(videoID uuid.UUID, contentType string) string {
	ext := contentTypeToExt(contentType)
	return fmt.Sprintf("%s%s", videoID, ext)
}

func contentTypeToExt(contentType string) any {
	parts := strings.Split(contentType, "")
	if len(parts) != 2 {
		return ".bin"
	}
	return "." + parts[1]
}

func (cfg *apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg *apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

func validThumbnailMimeType(contentType string) bool {
	mimeType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}
	return (mimeType == "image/jpeg" || mimeType == "image/png")
}