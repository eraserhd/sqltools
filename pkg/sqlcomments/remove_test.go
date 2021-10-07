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

func Test_Remove_removes_single_line_comments(t *testing.T) {
	check(t, " x--hello", " x")
	check(t, "--hello", "")
	check(t, "--hello\nthere", "\nthere")
	check(t, "  x--hello\nx --there", "  x\nx ")
}

func Test_Remove_does_not_remove_single_slashes(t *testing.T) {
	check(t, "/ hello/ *world /", "/ hello/ *world /")
}

func Test_Remove_removes_multi_line_comments(t *testing.T) {
	check(t, " hello/* ???\n */ there", " hello there")
	check(t, " hello/* ?\n?? **/ there", " hello there")
}

func Test_Remove_skips_single_quoted_single_line_comments(t *testing.T) {
	check(t, " x'--hello'there", " x'--hello'there")
}

func Test_Remove_skips_single_quoted_multi_line_comments(t *testing.T) {
	check(t, " x'/*hello*/'there", " x'/*hello*/'there")
}

func Test_Remove_is_not_confused_by_escaped_single_quotes(t *testing.T) {
	check(t, "  'hello\\'/*there*/' ", "  'hello\\'/*there*/' ")
}

func Test_Remove_skips_double_quoted_comments(t *testing.T) {
	check(t, ` x"--hello"there`, ` x"--hello"there`)
}
