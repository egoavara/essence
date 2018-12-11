package version

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math"
	"regexp"
	"strconv"
)

const _Invalid_text = "invalid"
const Undefined = math.MaxUint32

//
var ErrorVersionString = errors.New("_Invalid_text version string")
var Invalid = Version{
	Undefined,
	Undefined,
	Undefined,
}

func IsInvalid(v Version) bool {
	return v == Invalid
}

var (
	re_Version           = regexp.MustCompile(`^(?P<major>[0-9]+)(\.(?P<minor>[0-9]+)(\.(?P<patch>[0-9]+))?)?$`)
	re_Version_idx_major = 1
	re_Version_idx_minor = 3
	re_Version_idx_patch = 5
)

//
type Version struct {
	Major uint
	Minor uint
	Patch uint
}

func New(v ... uint) Version {
	switch len(v) {
	case 0:
		return Invalid
	case 1:
		return Version{
			Major:v[0],
			Minor:Undefined,
			Patch:Undefined,
		}
	case 2:
		return Version{
			Major:v[0],
			Minor:v[1],
			Patch:Undefined,
		}
	case 3:
		fallthrough
	default:
		return Version{
			Major:v[0],
			Minor:v[1],
			Patch:v[2],
		}

	}
}
func Parse(version string) Version {
	temp := New()
	if err := temp.UnmarshalJSON([]byte(`"` + version + `"`)); err != nil{
		return Invalid
	}
	return temp
}

func (s Version) String() string {
	if s.Major == Undefined {
		return _Invalid_text
	}
	if s.Minor == Undefined {
		return fmt.Sprintf("%d", s.Major)
	}
	if s.Patch == Undefined {
		return fmt.Sprintf("%d.%d", s.Major, s.Minor)
	}
	return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}

func (s Version) MarshalJSON() ([]byte, error) {
	if v := s.String(); v == _Invalid_text {
		return nil, errors.New("_Invalid_text version")
	} else {
		return []byte(`"` + v + `"`), nil
	}
}
func (s *Version) UnmarshalJSON(data []byte) error {
	var datastr string
	if err := json.Unmarshal(data, &datastr); err != nil{
		return errors.WithMessage(ErrorVersionString, string(data))
	}
	//
	var (
		temp   = re_Version.FindStringSubmatch(datastr)
		parsed uint64
		err    error
	)
	if temp == nil {
		return errors.WithMessage(ErrorVersionString, string(data))
	}
	// Major parse
	parsed, err = strconv.ParseUint(temp[re_Version_idx_major], 10, 32)
	if err == nil {
		s.Major = uint(parsed)
	}

	// Minor parse
	parsed, err = strconv.ParseUint(temp[re_Version_idx_minor], 10, 32)
	if err == nil {
		s.Minor = uint(parsed)
	}
	// Patch parse
	parsed, err = strconv.ParseUint(temp[re_Version_idx_patch], 10, 32)
	if err == nil {
		s.Patch = uint(parsed)
	}
	return nil
}
