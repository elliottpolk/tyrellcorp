package tyrellcorp

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
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
	nexus := testReplicant(t)

	if err := GenerateReplicant(nexus); err != nil {
		t.Fatal(err)
	}

	// TODO:
	// - include checks on generated content
}
