package dns

import (
	"net/url"
	"sync"
)

func NewRecordCollection() *RecordCollection {
	return &RecordCollection{
		mux:     sync.RWMutex{},
		records: make(map[string]*url.URL, 0),
	}
}

type RecordCollection struct {
	mux sync.RWMutex
	records map[string]*url.URL
}

func (recs *RecordCollection) All() []*url.URL {
	recs.mux.RLock()
	col := make([]*url.URL, 0, len(recs.records))

	for _, tx := range recs.records {
		col = append(col, tx)
	}
	recs.mux.RUnlock()
	return col
}

func (recs *RecordCollection) Store(url *url.URL) {
	recs.mux.Lock()
	recs.records[url.Host] = url
	recs.mux.Unlock()
}

func (recs *RecordCollection) Retrieve(host string) *url.URL {
	recs.mux.RLock()
	u := recs.records[host]
	recs.mux.RUnlock()
	return u
}

func (recs *RecordCollection) Remove(host string) {
	recs.mux.Lock()
	delete(recs.records, host)
	recs.mux.Unlock()
}
