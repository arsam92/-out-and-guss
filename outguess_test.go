package outguess

import (
 "bytes"
 "io"
 "testing"
)
func TestGetRejectsNonJPEGInput(t *testing.T) {
 if err := Get(bytes.NewReader([]byte("not an image")), io.Discard, nil); err == nil { t.Fatal("expected invalid input to be rejected") }
}