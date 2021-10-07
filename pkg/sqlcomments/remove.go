package sqlcomments

import (
	"bufio"
	"io"
)

const (
	start = iota
	inSingleLineComment
	inMultiLineComment
	inSingleQuotedString
)

func Remove(in io.Reader, out io.Writer) error {
	runes := bufio.NewReader(in)
	state := start
	for ch, _, err := runes.ReadRune(); err == nil; ch, _, err = runes.ReadRune() {
		next, _, err := runes.ReadRune()
		if err == nil {
			runes.UnreadRune()
		}

		if state == start && ch == '\'' {
			out.Write([]byte{'\''})
			state = inSingleQuotedString
		} else if state == inSingleQuotedString && ch == '\'' {
			out.Write([]byte{'\''})
			state = start
		} else if state == start && ch == '-' && next == '-' {
			runes.ReadRune()
			state = inSingleLineComment
		} else if state == inSingleLineComment && ch == '\n' {
			out.Write([]byte{'\n'})
			state = start
		} else if state == start && ch == '/' && next == '*' {
			runes.ReadRune()
			state = inMultiLineComment
		} else if state == inMultiLineComment && ch == '*' && next == '/' {
			runes.ReadRune()
			state = start
		} else if state != inSingleLineComment && state != inMultiLineComment {
			out.Write([]byte(string([]rune{ch})))
		}
	}
	return nil
}
