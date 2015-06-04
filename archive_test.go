package archive

import "testing"

func TestExtension(t *testing.T) {
	curLen := 26
	if len(Extensions) != curLen {
		t.Fatalf("Length doesn't match. Expected %d, Got %d", curLen, len(Extensions))
	}
}

func TestIs(t *testing.T) {
	if Is("a/b/c/bar.css") {
		t.Fatal("Wrong detection. Must not be archive")
	}

	if !Is("a/b/c/bar.egg") {
		t.Fatal("Wrong detection. Must be archive")
	}

	if !Is("a/b/c/bar.ZIP") {
		t.Fatal("Wrong detection. Must be archive")
	}

	if Is("a/b/c/7z") {
		t.Fatal("Wrong detection. Must not be archive")
	}
}
