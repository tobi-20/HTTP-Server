package request

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	// Split the request into lines using \r\n as separator
	lines := strings.Split(string(data), "\r\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty request")
	}

	// Parse only the first line
	reqLine, err := ParseRequestLine(lines[0])
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *reqLine,
	}, nil
}

// parseRequestLine parses a single HTTP request line
func ParseRequestLine(line string) (*RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: must have 3 parts")
	}

	method := parts[0]
	for _, r := range method {
		if !unicode.IsUpper(r) {
			return nil, fmt.Errorf("invalid method: must be uppercase letters only")
		}
	}

	requestTarget := parts[1]

	httpVersion := parts[2]
	if httpVersion != "HTTP/1.1" {
		return nil, fmt.Errorf("unsupported HTTP version: %s", httpVersion)
	}

	return &RequestLine{
		Method:        method,
		RequestTarget: requestTarget,
		HttpVersion:   "1.1",
	}, nil
}
