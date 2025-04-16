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
	}, nil
}

func parseRequestLine(b []byte) (RequestLine, error) {
	fullRequest := strings.Split(string(b), "\r\n")
	if len(fullRequest) == 0 {
		return RequestLine{}, fmt.Errorf("empty request")
	}

	requestLineParts := strings.Split(fullRequest[0], " ")
	if len(requestLineParts) != 3 {
		return RequestLine{}, fmt.Errorf("malformed request")
	}

	method := requestLineParts[0]
	for _, char := range method {
		if char < 'A' || char > 'Z' {
			return RequestLine{}, fmt.Errorf("method must contain only uppercase letters")
		}
	}
	target := requestLineParts[1]

	httpVersionFull := requestLineParts[2] // This would be something like "HTTP/1.1"
	httpVersionParts := strings.Split(httpVersionFull, "/")
	if len(httpVersionParts) != 2 || httpVersionParts[0] != "HTTP" {
		return RequestLine{}, fmt.Errorf("incorrect http version format")
	}
	httpVersion := httpVersionParts[1] // This gives you just "1.1"
	if httpVersion != "1.1" {
		return RequestLine{}, fmt.Errorf("incorrect http version")
	}

	requestlineStruct := RequestLine{
		HttpVersion:   httpVersion,
		RequestTarget: target,
		Method:        method,
	}

	return requestlineStruct, nil
}
