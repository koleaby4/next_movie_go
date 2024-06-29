package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestUpdateHighWatermark(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString("TMDB_BACKLOAD_HIGH_WATERMARK_DATE=2022-01-01")
	if err != nil {
		t.Fatal(err)
	}
	tempFile.Close()

	// initial date
	testDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	err = updateHighWatermark(testDate, tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := "TMDB_BACKLOAD_HIGH_WATERMARK_DATE=" + testDate.Format("2006-01-02")
	if !strings.Contains(string(data), expected) {
		t.Errorf("Expected to find %s, but it was not found", expected)
	}
}
