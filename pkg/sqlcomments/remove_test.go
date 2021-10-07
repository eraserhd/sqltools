package sqlcomments

import (
	"bytes"
	"testing"
)

func check(t *testing.T, in string, out string) {
	reader := bytes.NewBuffer([]byte(in))
	var writer bytes.Buffer
	err := Remove(reader, &writer)
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}
	result := string(writer.Bytes())
	if out != result {
		t.Errorf("wanted Remove(%q) = %q; got %q", in, out, result)
	}
}

func Test_Remove_does_not_remove_single_dashes(t *testing.T) {
	check(t, "-  hel-lo -world -", "-  hel-lo -world -")
}
