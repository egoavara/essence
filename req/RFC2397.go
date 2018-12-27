package req

import (
	"encoding/base64"
	"errors"
	"github.com/iamGreedy/essence/must"
	"io"
	"net/url"
	"strings"
)

// https://tools.ietf.org/html/rfc2397
type RFC2497 struct {
	Mime   string
	B64    bool
	Option map[string]string
	Data   []byte
	//
	index int
}

func (s *RFC2497) Reset() {
	s.Seek(0, io.SeekStart)
}
func (s *RFC2497) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		s.index = int(offset)
	case io.SeekCurrent:
		s.index += int(offset)
	case io.SeekEnd:
		s.index = len(s.Data) - int(offset)
	default:
		return 0, errors.New("invalid whence")
	}
	if s.index < 0 {
		s.index = 0
	}
	if s.index > len(s.Data) {
		s.index = len(s.Data)
	}
	return int64(s.index), nil
}
func (s *RFC2497) Read(p []byte) (n int, err error) {
	if s.index >= len(s.Data) {
		return 0, io.EOF
	}
	n = copy(p, s.Data[s.index:])
	s.index += n
	return n, nil
}
func (s *RFC2497) Close() error {
	return nil
}

func NewRFC2497(src *url.URL) *RFC2497 {
	if src.Scheme != "data" {
		return nil
	}
	datastart := strings.Index(src.Opaque, ",")
	options := strings.FieldsFunc(src.Opaque[:datastart], func(r rune) bool {
		return r == ';'
	})
	data := src.Opaque[datastart+1:]
	res := new(RFC2497)
	res.Option = make(map[string]string)
	for _, v := range options {
		if len(v) <= 0 {
			continue
		}
		if v == "base64" {
			res.B64 = true
			continue
		}
		if i := strings.Index(v, "="); i > 0 {
			res.Option[v[:i]] = v[i+1:]
		} else {
			res.Mime = v
		}
	}
	if res.B64 {
		res.Data = must.MustGet(base64.StdEncoding.DecodeString(data)).([]byte)
	} else {
		res.Data = []byte(data)
	}
	return res
}
