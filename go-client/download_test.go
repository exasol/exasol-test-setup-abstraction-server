package exasol_test_setup_abstraction_go

import (
	"os"
	"path"
	"testing"
)

func TestDownloadFails(t *testing.T) {
	dir := t.TempDir()
	localPath := path.Join(dir, "file")
	var tests = []struct {
		name          string
		url           string
		expectedError string
	}{
		{"invalid scheme", "invalid", "download failed: Get \"invalid\": unsupported protocol scheme \"\""},
		{"connection refused", "http://0.0.0.0/", "download failed: Get \"http://0.0.0.0/\": dial tcp 0.0.0.0:80: connect: connection refused"},
		{"file not found", "https://example.com/non-existing-file", "download of \"https://example.com/non-existing-file\" failed with status 404: \"404 Not Found\""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := downloadFile(test.url, localPath)
			if err == nil {
				t.Error("expected download to fail")
			} else if err.Error() != test.expectedError {
				t.Errorf("expected error %q but got %q", test.expectedError, err.Error())
			}
		})
	}
}

func TestDownloadSucceeds(t *testing.T) {
	dir := t.TempDir()
	localPath := path.Join(dir, "file")
	err := downloadFile("https://example.com", localPath)
	if err != nil {
		t.Errorf("download failed: %v", err)
		return
	}
	if fileInfo, err := os.Stat(localPath); err != nil {
		t.Errorf("failed to stat: %v", err)
	} else {
		t.Logf("download file size: %d", fileInfo.Size())
		if fileInfo.Size() <= 0 {
			t.Errorf("expected file size >0 but got %d", fileInfo.Size())
		}
	}
}
