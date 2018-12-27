package req

import (
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Requester interface {
	Request(u *url.URL) (io.ReadCloser, error)
}
type FnRequester func(u *url.URL) (io.ReadCloser, error)

func (s FnRequester) Request(u *url.URL) (io.ReadCloser, error) {
	return s(u)
}

var Standard FnRequester = func(u *url.URL) (io.ReadCloser, error) {
	switch u.Scheme {
	case "http":
		// http server
		fallthrough
	case "https":
		// http TLS server
		res, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		return res.Body, nil
	case "":
		fallthrough
	case "file":
		f, err := os.Open(u.Path)
		if err != nil {
			return nil, err
		}
		return f, nil
	case "data":
		sd := NewRFC2497(u)
		if sd != nil {
			return sd, nil
		}
	}
	return nil, errors.Errorf("Unsupported scheme '%s'", u.Scheme)
}
