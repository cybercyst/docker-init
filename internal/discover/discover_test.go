package discover

import (
	"docker-init/internal/types"
	"os"
	"strings"
	"testing"

	"github.com/spf13/afero"
)

func TestScanFolderForTargetsGo(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "go.mod", []byte(strings.TrimSpace(`
module go-test

go 1.19
`)), 0644)

	got, err := ScanFolderForTargets(fs)
	if err != nil {
		t.Fatalf("expected no error when detecting target of type Go, got %s", err)
	}

	currDir, err := os.Getwd()
	if err != nil {
		t.Fatal("expected no error when getting current directory")
	}

	want := types.Target{
		TargetType: types.Go,
		Path:       currDir,
	}

	if len(got) < 1 || got[0] != want {
		t.Fatal("expected one target of type Go, but none found")
	}
}

func TestScanFolderNoTargetFound(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "a.txt", []byte("My amazing text file"), 0644)
	afero.WriteFile(fs, "b.txt", []byte("Another life-changing file"), 0644)

	got, err := ScanFolderForTargets(fs)
	if err != nil {
		t.Fatal("expected no error when scanning folder with no targets")
	}

	if len(got) > 0 {
		t.Fatal("expected no targets when scanning folder with no targets")
	}
}
