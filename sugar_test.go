package sugar

import (
	"github.com/gosexy/to"
	"testing"
)

func TestGetSet(t *testing.T) {
	tuple := Tuple{}

	tuple.Set("foo/bar/baz", 1)
	tuple.Set("foo/bar/baz", 2)
	tuple.Set("foo/bar/bad", 3)
	tuple.Set("foo/bax/bad", "4")
	tuple.Set("foo/bax/Bad", "5")

	if to.Int(tuple.Get("foo/bar/baz")) != 2 {
		t.Errorf("Test failed.\n")
	}

	if tuple["foo"].(map[string]interface{})["bax"].(map[string]interface{})["Bad"].(string) != "5" {
		t.Errorf("Test failed.\n")
	}

}
