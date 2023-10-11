package exasol_test_setup_abstraction_go

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func downloadFile(url, localPath string) error {
	log.Printf("Downloading %q to local path %q...", url, localPath)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("download of %q failed with status %d: %q", url, resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	out, err := os.Create(localPath)
	defer func() {
		err = out.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close server file. Cause: %v", err))
		}
	}()
	if err != nil {
		return fmt.Errorf("failed to create file for: %w", err)
	}
	fileSize, err := io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to download from %q to %q: %w", url, localPath, err)
	}
	log.Printf("File downloaded with size %d bytes", fileSize)
	return nil
}
