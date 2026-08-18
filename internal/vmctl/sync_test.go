package vmctl

import (
	"encoding/xml"
	"os"
	"strings"
	"testing"
)

func TestInjectVirtiofs(t *testing.T) {
	raw, err := os.ReadFile("testdata/sample.xml")
	if err != nil {
		t.Fatal(err)
	}
	original := string(raw)

	if strings.Contains(original, "memoryBacking") {
		t.Fatal("fixture unexpectedly already has memoryBacking, test assumption broken")
	}

	result, already := injectVirtiofs(original, "/home/chris/Projects/OSS/zedlum")
	if already {
		t.Fatal("expected alreadyPresent=false on first injection")
	}

	if err := xml.Unmarshal([]byte(result), new(any)); err != nil {
		t.Fatalf("result is not well-formed XML: %v", err)
	}
	if !strings.Contains(result, "<memoryBacking>") {
		t.Error("missing memoryBacking block")
	}
	if !strings.Contains(result, "dir='zedshare'") {
		t.Error("missing filesystem target dir")
	}
	if !strings.Contains(result, "dir='/home/chris/Projects/OSS/zedlum'") {
		t.Error("missing filesystem source dir")
	}
	if strings.Count(result, "<filesystem") != 1 {
		t.Error("expected exactly one filesystem device")
	}

	// idempotency: running again against the already-modified doc must be a no-op
	again, already2 := injectVirtiofs(result, "/home/chris/Projects/OSS/zedlum")
	if !already2 {
		t.Error("expected alreadyPresent=true on second injection")
	}
	if again != result {
		t.Error("second injection should not modify the document")
	}
}
