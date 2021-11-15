package parseutil

import (
    "io"
)

// If the reader's leading bytes are ascii character of the form 'x' (or
// '\t' for escaped characters), pop those bytes off the reader and return
// the value.  Otherwise, return a nil slice.
func MaybeTokenizeCharacter(reader *LocationReader) ([]byte, Location, error) {
    bytes, err := reader.Peek(4)
    if err != nil {
        return nil, Location{}, err
    }

    if len(bytes) < 3 {
        return nil, Location{}, nil
    }

    if bytes[0] != '\'' {
        return nil, Location{}, nil
    }

    numBytes := 3
    if bytes[1] == '\\' { // c escape
        if len(bytes) < 4 || bytes[3] != '\'' {
            return nil, Location{}, nil
        }

        switch bytes[2] {
        case 't', 'n', '\'', '\\':
        default:
            return nil, Location{}, nil
        }

        numBytes = 4
    } else {
        if bytes[2] != '\'' {
            return nil, Location{}, nil
        }

        valid := false

        char := bytes[1]
        if ('a' <= char && char <= 'z') ||
            ('A' <= char && char <= 'Z') ||
            ('0' <= char && char <= '9') {

            valid = true
        } else {
            switch char {
            case '`', '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
                '-', '_', '=', '+', '[', '{', ']', '}', '|', ';', ':', '"',
                ',', '<', '.', '>', '/', '?', ' ':
                valid = true

            }
        }

        if !valid {
            return nil, Location{}, nil
        }
    }

    loc := reader.Location

    bytes = bytes[:numBytes]
    n, err := reader.Read(bytes)
    if len(bytes) != n || err != nil {
        panic(err) // should never happen
    }

    return bytes, loc, nil
}

// If the reader's leading bytes are identifer of the form [a-zA-Z_]\w* ,
// pop those bytes off the reader and return the value.  Otherwise, return a
// nil slice.
func MaybeTokenizeIdentifier(reader *LocationReader) ([]byte, Location, error) {
    peekRange := 32
    prevLen := 0
    checkIdx := 0

    var bytes []byte
    var err error
    for {
        bytes, err = reader.Peek(peekRange)
        if err != nil && err != io.EOF {
            return nil, Location{}, err
        }

        if len(bytes) == 0 {
            return nil, Location{}, nil
        }

        if checkIdx == 0 {
            char := bytes[0]
            if !(('a' <= char && char <= 'z') ||
                ('A' <= char && char <= 'Z') ||
                char == '_') {

                return nil, Location{}, nil
            }

            checkIdx = 1
        }

        if prevLen == len(bytes) {  // ran out of bytes to read
            break
        }

        foundEnd := false
        for checkIdx < len(bytes) {
            char := bytes[checkIdx]
            if !(('a' <= char && char <= 'z') ||
                ('A' <= char && char <= 'Z') ||
                ('0' <= char && char <= '9') ||
                char == '_') {

                foundEnd = true
                break
            }
            checkIdx += 1
        }

        if foundEnd {
            break
        }

        prevLen = len(bytes)
        peekRange *= 2
    }

    loc := reader.Location

    bytes = bytes[:checkIdx]
    n, err := reader.Read(bytes)
    if len(bytes) != n || err != nil {
        panic(err) // should never happen
    }

    return bytes, loc, nil
}
