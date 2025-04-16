package request

import (
	"fmt"
	"io"
	"strings"
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
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	line, err := parseRequestLine(b)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: line,
		},
	}, nil
}

func parseRequestLine(b []byte) (string, error) {
	fullRequest := strings.Split(string(b), "\r\n")

	requestLine := strings.Split(fullRequest[0], " ")

	requestlineStruct := RequestLine{
		HttpVersion:   str[0],
		RequestTarget: str[1],
		Method:        str[2],
	}

	return requestline, nil
}
