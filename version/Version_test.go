package version

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestParseVersion(t *testing.T) {
	t.Run("case major.minor.patch", func(t *testing.T) {
		v0 := Parse("1.2.3")
		if v0 != New(1, 2, 3) {
			t.Error("version parse Temp fail")
		}
	})
	t.Run("case major.minor", func(t *testing.T) {
		v1 := Parse("1.2")
		if v1 != New(1, 2) {
			t.Error("version parse Temp fail")
		}
	})
	t.Run("case major", func(t *testing.T) {
		v1 := Parse("1")
		if v1 != New(1) {
			t.Error("version parse Temp fail")
		}
	})
	t.Run("case invalud", func(t *testing.T) {
		v1 := Parse("1.2.a")
		if !IsInvalid(v1) {
			t.Error("version parse Temp fail")
		}
	})
}
func TestVersion_UnmarshalJSON(t *testing.T) {
	type Temp struct {
		I int     `json:"i"`
		S string  `json:"s"`
		V Version `json:"v"`
	}
	//
	t.Run("Success test", func(t *testing.T) {
		dec := json.NewDecoder(strings.NewReader(`
{
	"i" : 1, 
	"s" : "hello, world!", 
	"v" : "1.2.3"
}
`))
		var temp Temp
		if err := dec.Decode(&temp); err != nil{
			t.Error(err)
		}
		if temp.I != 1 || temp.S != "hello, world!" || temp.V != New(1,2,3){
			t.Error(temp, "Invalid Parsed")
		}
	})
	t.Run("Fail test 0", func(t *testing.T) {
		dec := json.NewDecoder(strings.NewReader(`
{
	"i" : 1, 
	"s" : "hello, world!", 
	"v" : "1.2.3a"
}
`))
		var temp Temp
		if err := dec.Decode(&temp); err == nil{
			t.Error("Must be fail, but success", temp)
		}
	})
	t.Run("Fail test 1", func(t *testing.T) {
		dec := json.NewDecoder(strings.NewReader(`
{
	"i" : 1, 
	"s" : "hello, world!", 
	"v" : 1
}
`))
		var temp Temp
		if err := dec.Decode(&temp); err == nil{
			t.Error("Must be fail, but success", temp)
		}
	})
}
func TestVersion_MarshalJSON(t *testing.T) {
	type Temp struct {
		I int     `json:"i"`
		S string  `json:"s"`
		V Version `json:"v"`
	}
	t.Run("Success test", func(t *testing.T) {
		var testset = `{"i":1,"s":"hello, world!","v":"1.2.3"}`
		var temp  = Temp{
			I: 1,
			S: "hello, world!",
			V: New(1,2,3),
		}
		if data, err := json.Marshal(temp); err != nil && string(data) == testset{
			t.Error(err)
		}
	})
	t.Run("Fail test", func(t *testing.T) {
		var temp  = Temp{
			I: 1,
			S: "hello, world!",
			V: New(),
		}

		if data, err := json.Marshal(temp); err == nil{
			t.Error("Parsing success, but must be fail " + string(data))
		}
	})

}
func TestVersion_Stringer(t *testing.T) {
	if New().String() != _Invalid_text{
		t.Error("invalid fail")
	}
	if New(1).String() != "1"{
		t.Error("'1' fail")
	}
	if New(1, 2).String() != "1.2"{
		t.Error("'1.2' fail")
	}
	if New(1, 2, 3).String() != "1.2.3"{
		t.Error("'1.2.3' fail")
	}
}
