package leetcode

import (
	"strings"
)

func lengthLongestPath(input string) int {
	p := pathParser{remaining: input}
	maxLen := 0
	// longestPath := ""
	for p.Scan() {
		// fmt.Println(p.Path())
		if t := p.PathLen(); t > maxLen {
			maxLen = t
			// longestPath = p.Path()
		}
	}
	return maxLen
}

type pathParser struct {
	remaining      string
	path           []string
	pathLenNoSlash int
}

func (p *pathParser) Scan() bool {
	for {
		e, ok := p.readNextLine()
		if !ok {
			return false
		}
		for e.depth < len(p.path) {
			last := len(p.path) - 1
			p.pathLenNoSlash -= len(p.path[last])
			p.path = p.path[:last]
		}
		if e.depth != len(p.path) {
			panic("what")
		}
		p.path = append(p.path, e.name)
		p.pathLenNoSlash += len(e.name)
		if e.IsFile() {
			return true
		}
	}
}

func (p *pathParser) PathLen() int { return p.pathLenNoSlash + len(p.path) - 1 }
func (p *pathParser) Path() string { return strings.Join(p.path, "/") }

type lineEntry struct {
	depth int
	name  string
}

func (e lineEntry) IsFile() bool {
	return strings.Contains(e.name, ".")
}

func (p *pathParser) readNextLine() (lineEntry, bool) {
	lineEnd := strings.IndexByte(p.remaining, '\n')
	if lineEnd < 0 {
		if len(p.remaining) > 0 {
			lineEnd = len(p.remaining)
		} else {
			return lineEntry{}, false
		}
	}
	line := p.remaining[:lineEnd]
	if lineEnd == len(p.remaining) {
		p.remaining = ""
	} else {
		p.remaining = p.remaining[lineEnd+1:]
	}
	var e lineEntry
	for len(line) > 0 && line[0] == '\t' {
		e.depth++
		line = line[1:]
	}
	e.name = line // stripped leading \t and end-of-line
	return e, true
}
