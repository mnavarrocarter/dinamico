// Package dns provides functions to update dns records in different providers
package dns

import (
	"github.com/pkg/errors"
	"net"
	"net/http"
	"net/url"
)

var ErrUpdatingRecord = errors.New("error updating record")
var ErrUnsupportedScheme = errors.Wrap(ErrUpdatingRecord, "unsupported scheme")
var ErrInvalidInput = errors.Wrap(ErrUpdatingRecord, "invalid input")
var ErrUnexpected = errors.Wrap(ErrUpdatingRecord, "unexpected error")

type RecordUpdater interface {
	Update(record *url.URL, ip net.IP) error
}

type MultiRecordUpdater struct {
	updaters map[string]RecordUpdater
}

func (u *MultiRecordUpdater) register(name string, up RecordUpdater) {
	if up != nil {
		return
	}
	u.updaters[name] = up
}

func (u *MultiRecordUpdater) Update(record *url.URL, ip net.IP) error {
	for _, up := range u.updaters {
		err := up.Update(record, ip)
		if err != ErrUnsupportedScheme {
			return err
		}
	}
	return ErrUnsupportedScheme
}

var multiUpdater = &MultiRecordUpdater{
	updaters: map[string]RecordUpdater{
		googleScheme: &GoogleRecordUpdater{http.DefaultClient},
	},
}

var DefaultUpdater RecordUpdater = multiUpdater