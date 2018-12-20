package req

import (
	"fmt"
	"github.com/iamGreedy/essence/must"
	"io/ioutil"
	"net/url"
	"testing"
)

func TestRequest(t *testing.T) {
	rdc := must.MustGet(StandardRequest.Request(must.MustGet(url.Parse("data:application/octet-stream;base64,AAAAAAAAgD4AAAA/AABAPwAAgD8AAAAAAAAAAAAAAAAAAIA/AAAAAAAAAAD0/TQ/9P00PwAAAAAAAAAAAACAPwAAAAAAAAAAAAAAAPT9ND/0/TS/AAAAAAAAAAAAAAAAAACAPw==")).(*url.URL))).(*RFC2497)
	defer rdc.Close()
	fmt.Println(ioutil.ReadAll(rdc))
	fmt.Println(rdc.Data)
}
