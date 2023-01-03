package discover

import (
	"docker-init/internal/types"
	"os"
	"testing"
	"testing/fstest"
)

func TestScanFolderForTargetsGo(t *testing.T) {
	m := fstest.MapFS{
		"go.mod": {},
	}

	got, err := ScanFolderForTargets(m)
	if err != nil {
		t.Fatal("expected no error when detecting target of type Go")
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
	m := fstest.MapFS{
		"a.txt": {Data: []byte("My amazing text file")},
		"b.txt": {Data: []byte("Another life-changing file")},
	}

	got, err := ScanFolderForTargets(m)
	if err != nil {
		t.Fatal("expected no error when scanning folder with no targets")
	}

	if len(got) > 0 {
		t.Fatal("expected no targets when scanning folder with no targets")
	}
}
