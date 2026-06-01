package lib

import "testing"

func TestEntryAddPrefixIgnoresCommentLines(t *testing.T) {
	entry := NewEntry("test")

	for _, line := range []string{"", "   ", "# comment", "// comment", "/* comment */"} {
		if err := entry.AddPrefix(line); err != nil {
			t.Fatalf("AddPrefix(%q) returned error: %v", line, err)
		}
	}

	if err := entry.AddPrefix("192.0.2.1"); err != nil {
		t.Fatalf("AddPrefix returned error: %v", err)
	}

	prefixes, err := entry.MarshalText()
	if err != nil {
		t.Fatalf("MarshalText returned error: %v", err)
	}
	if len(prefixes) != 1 || prefixes[0] != "192.0.2.1/32" {
		t.Fatalf("MarshalText = %v, want [192.0.2.1/32]", prefixes)
	}
}

func TestEntryRemovePrefixIgnoresCommentLines(t *testing.T) {
	entry := NewEntry("test")
	if err := entry.AddPrefix("192.0.2.0/24"); err != nil {
		t.Fatalf("AddPrefix returned error: %v", err)
	}

	for _, line := range []string{"", "   ", "# comment", "// comment", "/* comment */"} {
		if err := entry.RemovePrefix(line); err != nil {
			t.Fatalf("RemovePrefix(%q) returned error: %v", line, err)
		}
	}

	prefixes, err := entry.MarshalText()
	if err != nil {
		t.Fatalf("MarshalText returned error: %v", err)
	}
	if len(prefixes) != 1 || prefixes[0] != "192.0.2.0/24" {
		t.Fatalf("MarshalText = %v, want [192.0.2.0/24]", prefixes)
	}
}
