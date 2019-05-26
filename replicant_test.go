package tyrellcorp

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testReplicant(t *testing.T) *Spec {
	in, err := ioutil.ReadFile(filepath.Join("test_data", "sample_spec.json"))
	if err != nil {
		t.Fatal(err)
	}

	s := &Spec{}
	if err := json.Unmarshal(in, &s); err != nil {
		t.Fatal(err)
	}

	return s
}

func TestGenReplicant(t *testing.T) {
	const testDir = "test_data"

	var (
		nexus = testReplicant(t)
		dir   = getDir(nexus.Repository, nexus.Package)
	)

	if err := GenerateReplicant(nexus); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	sfi, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	// this should be the total number of top level files + dir
	assert.Len(t, sfi, 16)

	// run through the files and make sure they're the same
	for _, fi := range sfi {
		assert.NoError(t, filepath.Walk(filepath.Join(dir, fi.Name()), func(p string, i os.FileInfo, err error) error {
			if !i.IsDir() {
				t.Run(i.Name(), func(t *testing.T) {
					got, err := ioutil.ReadFile(p)
					if err != nil {
						t.Fatal(err)
					}

					td := testDir
					if filepath.Dir(p) != dir {
						td = filepath.Join(testDir, filepath.Base(filepath.Dir(p)))
					}

					want, err := ioutil.ReadFile(filepath.Join(td, i.Name()))
					if err != nil {
						t.Fatal(err)
					}

					assert.Equal(t, want, got)
				})
				return nil
			}
			return nil
		}))
	}
}
