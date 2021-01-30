package main_test
import (
	cfg "github.com/nil-labs/configify"
	"testing"
)
func TestMain(t *testing.T) {
    want := "hello world"
    if got := cfg.Greet(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}