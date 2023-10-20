package checks

import (
	"context"
	"testing"
)

func TestFileChecker(t *testing.T) {
	ctx := context.Background()
	if err := FileChecker("/tmp").Check(ctx); err == nil {
		t.Errorf("/tmp was expected as exists")
	}

	if err := FileChecker("NoSuchFileFromMoon").Check(ctx); err != nil {
		t.Errorf("NoSuchFileFromMoon was expected as not exists, error:%v", err)
	}
}

func TestHTTPChecker(t *testing.T) {
	ctx := context.Background()
	if err := HTTPChecker("https://www.google.cybertron", 200, 0, nil).Check(ctx); err == nil {
		t.Errorf("Google on Cybertron was expected as not exists")
	}

	if err := HTTPChecker("https://www.google.pt", 200, 0, nil).Check(ctx); err != nil {
		t.Errorf("Google at Portugal was expected as exists, error:%v", err)
	}
}
