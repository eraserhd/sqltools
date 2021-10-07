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
	inDoubleQuotedIdent
	inDollarQuotedString
)

func Remove(in io.Reader, out io.Writer) error {
	runes := bufio.NewReader(in)
	state := start
	for ch, _, err := runes.ReadRune(); err == nil; ch, _, err = runes.ReadRune() {
		next, _, err := runes.ReadRune()
		if err == nil {
			runes.UnreadRune()
		}

		// Single-quoted strings
		if state == start && ch == '\'' {
			out.Write([]byte{'\''})
			state = inSingleQuotedString
			continue
		}
		if state == inSingleQuotedString && ch == '\'' {
			out.Write([]byte{'\''})
			state = start
			continue
		}
		if state == inSingleQuotedString && ch == '\\' {
			runes.ReadRune()
			out.Write([]byte{'\\'})
			out.Write([]byte(string([]rune{next})))
			continue
		}

		// Dollar-quoted strings
		if state == start && ch == '$' && next == '$' {
			runes.ReadRune()
			out.Write([]byte{'$', '$'})
			state = inDollarQuotedString
			continue
		}
		if state == inDollarQuotedString && ch == '$' && next == '$' {
			runes.ReadRune()
			out.Write([]byte{'$', '$'})
			state = start
			continue
		}

		// Double-quoted identifiers
		if state == start && ch == '"' {
			out.Write([]byte{'"'})
			state = inDoubleQuotedIdent
			continue
		}
		if state == inDoubleQuotedIdent && ch == '"' {
			out.Write([]byte(string([]rune{ch})))
			state = start
			continue
		}

		// Single-line comments
		if state == start && ch == '-' && next == '-' {
			runes.ReadRune()
			state = inSingleLineComment
			continue
		}
		if state == inSingleLineComment && ch == '\n' {
			out.Write([]byte{'\n'})
			state = start
			continue
		}

		// Multi-line comments
		if state == start && ch == '/' && next == '*' {
			runes.ReadRune()
			state = inMultiLineComment
			continue
		}
		if state == inMultiLineComment && ch == '*' && next == '/' {
			runes.ReadRune()
			state = start
			continue
		}

		// Everything else
		if state != inSingleLineComment && state != inMultiLineComment {
			out.Write([]byte(string([]rune{ch})))
			continue
		}
	}
	return nil
}
