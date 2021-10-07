package sqlcomments

import (
	"bufio"
	"io"
)

const (
	start = iota
	afterDash
	inSingleLineComment
)

func Remove(in io.Reader, out io.Writer) error {
	runes := bufio.NewReader(in)
	state := start
	for ch, _, err := runes.ReadRune(); err == nil; ch, _, err = runes.ReadRune() {
	restart:
		switch state {
		case start:
			if ch == '-' {
				state = afterDash
			} else {
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
				state = start
			}
		}
	}
	switch state {
	case afterDash:
		out.Write([]byte{'-'})
	}
	return nil
}
