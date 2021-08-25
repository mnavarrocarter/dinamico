package dns

import (
	"bufio"
	"io"
	"net/url"
)


func ParseRecords(r io.Reader) (*RecordCollection, error) {
	records := NewRecordCollection()
	s := bufio.NewScanner(r)
	for s.Scan() {
		str := string(s.Bytes())
		if str == "" {
			continue
		}
		u, err := url.Parse(str)
		if err != nil {
			return nil, err
		}
		records.Store(u)
	}
	return records, nil
}
