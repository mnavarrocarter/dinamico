package dns

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net"
	"net/http"
	"net/url"
)

const googleScheme = "google-domains"
const googleUrl = "https://domains.google.com/nic/update?hostname=%s&myip=%s"

type GoogleRecordUpdater struct {
	client *http.Client
}

func (p *GoogleRecordUpdater) Update(record *url.URL, ip net.IP) error {
	if record.Scheme != googleScheme {
		return ErrUnsupportedScheme
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(googleUrl, record.Host, ip.String()), nil)
	if err != nil {
		return errors.Wrap(ErrUnexpected, err.Error())
	}

	if record.User == nil {
		return errors.Wrap(ErrInvalidInput, "username is required")
	}

	username := record.User.Username()
	password, ok := record.User.Password()
	if !ok {
		return errors.Wrap(ErrInvalidInput, "password is required")
	}

	req.SetBasicAuth(username, password)

	res, err := p.client.Do(req)
	if err != nil {
		return errors.Wrap(ErrUnexpected, err.Error())
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return errors.Wrap(ErrUnexpected, "http response was not 200")
	}

	return nil
}
