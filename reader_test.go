package csvreader

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type testStruct struct {
	Name     string
	UserName string
	ID       int
	Enable   bool
	Type     CustomeType
}

type CustomeType int

func (c *CustomeType) FromString(str string) error {
	switch str {
	case "tcp":
		*c = 0
	case "udp":
		*c = 1
	default:
		return fmt.Errorf("unknown type:%s", str)
	}
	return nil
}

func TestCus(t *testing.T) {
	var c CustomeType
	v := reflect.ValueOf(c)
	if v.CanAddr() {
		t.Log("yes")
		if m, ok := v.Addr().Interface().(CsvMarshal); ok {
			t.Log("ok")
			m.FromString("udp")
		}
	}
	if m, ok := v.Interface().(CsvMarshal); ok {
		m.FromString("udp")
		t.Log(c)
	} else {
		t.Log("no")
	}
}

func TestSnakeName(t *testing.T) {
	reader := csv.NewReader(bytes.NewReader([]byte("zhengxin,zhnxin,0,false\nxinzheng,zhnxin,1,true")))
	bean := []testStruct{}
	if err := New().WithHeader([]string{"name", "user_name", "id", "enable"}).UnMarshal(reader, &bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	json.Unmarshal([]byte(""), bean)
	t.Log(string(b))
}

func TestLowerName(t *testing.T) {
	reader := csv.NewReader(bytes.NewReader([]byte("zhengxin,zhnxin,0,false\nxinzheng,zhnxin,1,true")))
	bean := []*testStruct{}
	if err := New().WithHeader([]string{"NAME", "USERNAME", "ID", "ENABLE"}).UnMarshal(reader, &bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	t.Log(string(b))
}

func TestCustom(t *testing.T) {
	reader := csv.NewReader(bytes.NewReader([]byte("zhengxin,zhnxin,udp,false\nxinzheng,zhnxin,udp,true")))
	bean := []*testStruct{}
	if err := New().WithHeader([]string{"NAME", "USERNAME", "type", "ENABLE"}).UnMarshal(reader, &bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	t.Log(string(b))
}
