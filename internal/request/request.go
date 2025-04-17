package request

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"errors"
)

type Request struct {
	RequestLine RequestLine
	state       int
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

const crlf = "\r\n"

const (
	parserInitialized = iota
	parserDone
)

func RequestFromReader(reader io.Reader) (*Request, error) {
    const bufferSize = 8
    buffer := make([]byte, bufferSize)
    readToIndex := 0

    req := &Request{
        state: parserInitialized,
    }

    for req.state != parserDone {
        if readToIndex == len(buffer) {
            newBuffer := make([]byte, len(buffer)*2)
            copy(newBuffer, buffer)
            buffer = newBuffer
        }

        readBytes, err := reader.Read(buffer[readToIndex:])
        if err != nil && err != io.EOF {
            return nil, err
        }

        readToIndex += readBytes

        parsedBytes, parsedErr := req.parse(buffer[:readToIndex])
        if parsedErr != nil {
            return nil, parsedErr
        }

        copy(buffer, buffer[parsedBytes:readToIndex])
        readToIndex -= parsedBytes

        if err == io.EOF {
            if req.state != parserDone {
                return nil, errors.New("incomplete request")
            }
            break 
        }
    }

    return req, nil
}

func parseRequestLine(data []byte, requestLine *RequestLine) (int, error) {
    idx := bytes.Index(data, []byte(crlf))
    if idx == -1 {
        // Can't find \r\n, need more data
        return 0, nil
    }
    
    requestLineText := string(data[:idx])
    parsedLine, err := requestLineFromString(requestLineText)
    if err != nil {
        return 0, err
    }
    *requestLine = *parsedLine
    
    return idx + 2, nil
}

func requestLineFromString(str string) (*RequestLine, error) {
	parts := strings.Split(str, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("poorly formatted request-line: %s", str)
	}

	method := parts[0]
	for _, c := range method {
		if c < 'A' || c > 'Z' {
			return nil, fmt.Errorf("invalid method: %s", method)
		}
	}

	requestTarget := parts[1]

	versionParts := strings.Split(parts[2], "/")
	if len(versionParts) != 2 {
		return nil, fmt.Errorf("malformed start-line: %s", str)
	}

	httpPart := versionParts[0]
	if httpPart != "HTTP" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", httpPart)
	}
	version := versionParts[1]
	if version != "1.1" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", version)
	}

	return &RequestLine{
		Method:        method,
		RequestTarget: requestTarget,
		HttpVersion:   versionParts[1],
	}, nil
}


func (r *Request) parse(data []byte) (int, error) {
    // Check the current state
    switch r.state {
    case parserInitialized:
        // If we're in the initialized state, try to parse the request line
        bytesRead, err := parseRequestLine(data, &r.RequestLine)
        if err != nil {
            return 0, err
        }
        if bytesRead == 0 {
            // Need more data, no error
            return 0, nil
        }
        // Successfully parsed request line, update state
        r.state = parserDone
        return bytesRead, nil
    case parserDone:
        // Parser is already done
        return 0, errors.New("error: trying to read data in a done state")
    default:
        // Unknown state
        return 0, errors.New("error: unknown state")
    }
}