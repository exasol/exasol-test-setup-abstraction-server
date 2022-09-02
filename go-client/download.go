package exasol_test_setup_abstraction_go

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url, localPath string) error {
	out, err := os.Create(localPath)
	defer func() {
		err = out.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close server file. Cause: %v", err.Error()))
		}
	}()
	if err != nil {
		return fmt.Errorf("failed to create file for: %v", err.Error())
	}
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %v", err.Error())
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("download of %q failed with status %d: %q", url, resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to download from %q to %q: %v", url, localPath, err.Error())
	}
	return nil
}
