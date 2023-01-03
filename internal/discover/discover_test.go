package discover

import (
	"docker-init/internal/types"
	"testing"
	"testing/fstest"
)

func TestScanFolderForTargetsGo(t *testing.T) {
	m := fstest.MapFS{
		"go_project/go.mod": {},
	}

	got, err := ScanFolderForTargets(m)
	if err != nil {
		t.Fatal("expected no error when detecting target of type Go")
	}

	want := types.Target{
		TargetType: types.Go,
		Path:       "go_project",
	}

	if len(got) < 1 || got[0] != want {
		t.Fatal("expected one target of type Go, but none found")
	}
}

func TestScanFolderForTargetsFolderDoesntExist(t *testing.T) {
	m := fstest.MapFS{}

	got, err := ScanFolderForTargets(m)
	if len(got) != 0 {
		t.Fatal("expected no targets when directory doesn't exist")
	}

	if err == nil {
		t.Fatal("expected error if directory doesn't exist")
	}
}

func TestScanFolderNoTargetFound(t *testing.T) {
	m := fstest.MapFS{
		"rando_directory/a.txt": {Data: []byte("My amazing text file")},
		"rando_directory/b.txt": {Data: []byte("Another life-changing file")},
	}

	got, err := ScanFolderForTargets(m)
	if err != nil {
		t.Fatal("expected no error when scanning folder with no targets")
	}

	if len(got) > 0 {
		t.Fatal("expected no targets when scanning folder with no targets")
	}
}
