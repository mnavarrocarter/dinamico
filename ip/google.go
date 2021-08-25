package ip

import (
	"github.com/pkg/errors"
	"io"
	"net"
	"net/http"
)

const ipCheckUrl = "https://domains.google.com/checkip"

type GoogleProvider struct {
	client *http.Client
}

func NewGoogleProvider() *GoogleProvider {
	return &GoogleProvider{http.DefaultClient}
}

func (p *GoogleProvider) ObtainIp() (net.IP, error) {
	resp, err := p.client.Get(ipCheckUrl)
	if err != nil {
		return nil, errors.Wrap(ErrFecthingIp, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(ErrFecthingIp, "http response code was %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(ErrFecthingIp, err.Error())
	}

	ip := net.ParseIP(string(b))
	if ip == nil {
		return nil, errors.Wrap(ErrFecthingIp, "could not parse ip address")
	}
	return ip, nil
}
