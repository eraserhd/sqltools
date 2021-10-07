package sqlcomments

import (
	"bufio"
	"io"
)

const (
	start = iota
	afterDash
	inSingleLineComment
	afterSlash
	inMultiLineComment
	inMultiLineCommentAfterStar
	inSingleQuotedString
)

func Remove(in io.Reader, out io.Writer) error {
	runes := bufio.NewReader(in)
	state := start
	for ch, _, err := runes.ReadRune(); err == nil; ch, _, err = runes.ReadRune() {
	restart:
		switch state {
		case start:
			switch ch {
			case '-':
				state = afterDash
			case '/':
				state = afterSlash
			case '\'':
				out.Write([]byte{'\''})
				state = inSingleQuotedString
			default:
				out.Write([]byte(string([]rune{ch})))
			}
		case afterDash:
			if ch == '-' {
				state = inSingleLineComment
			} else {
				out.Write([]byte{'-'})
				state = start
				goto restart
			}
		case inSingleLineComment:
			if ch == '\n' {
				out.Write([]byte{'\n'})
				state = start
			}
		case afterSlash:
			if ch == '*' {
				state = inMultiLineComment
			} else {
				out.Write([]byte{'/'})
				state = start
				goto restart
			}
		case inMultiLineComment:
			if ch == '*' {
				state = inMultiLineCommentAfterStar
			}
		case inMultiLineCommentAfterStar:
			if ch == '/' {
				state = start
			} else {
				state = inMultiLineComment
				goto restart
			}
		case inSingleQuotedString:
			out.Write([]byte(string([]rune{ch})))
			if ch == '\'' {
				state = start
			}
		}
	}
	switch state {
	case afterDash:
		out.Write([]byte{'-'})
	case afterSlash:
		out.Write([]byte{'/'})
	}
	return nil
}
