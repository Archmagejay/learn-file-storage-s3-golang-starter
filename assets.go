package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

// Get the asset path for a given mediaType
func getAssetPath(mediaType string) string {
	ext := contentTypeToExt(mediaType)
	id := getRandomAssetName()
	return fmt.Sprintf("%s%s", id, ext)
}

// Convert from mime type to file extension
func contentTypeToExt(contentType string) any {
	parts := strings.Split(contentType, "/")
	if len(parts) != 2 {
		return ".bin"
	}
	return "." + parts[1]
}

// Get the provided asset's disk path
func (cfg *apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

// Get the provided asset's local URL
func (cfg *apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

// Get the provided object's S3 URL
func (cfg *apiConfig) getObjectURL(key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", cfg.s3Bucket, cfg.s3Region, key)
}

// Generate a random 32 byte base 64 encoded string
func getRandomAssetName() string {
	base := make([]byte, 32)
	_, err := rand.Read(base)
	if err != nil {
		panic("failed to generate random bytes")
	}
	return base64.RawURLEncoding.EncodeToString(base)
}
