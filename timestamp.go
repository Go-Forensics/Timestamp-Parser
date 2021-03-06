// Copyright (c) 2020 Alec Randazzo

package timestamp

import (
	"fmt"
	bin "github.com/Go-Forensics/BinaryTransforms"
	"time"
)

// RawTimestamp is a []byte alias
type RawTimestamp []byte

// Parse a byte slice containing a unix timestamp and convert it to a timestamp string.
func (rawTimestamp RawTimestamp) Parse() (timestamp time.Time, err error) {

	// verify that we are getting the bytes we need
	if len(rawTimestamp) != 8 {
		err = fmt.Errorf("timestamp.parse() received %d bytes, not 8 bytes", len(rawTimestamp))
		timestamp = time.Time{} // time.Time nil equivalent
		return
	}

	var delta = time.Date(1970-369, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
	// Convert the byte slice to little endian int64 and then convert it to a string
	timestampInt64, _ := bin.LittleEndianBinaryToInt64(rawTimestamp)

	timestamp = time.Unix(0, timestampInt64*100+delta).UTC()

	return
}
