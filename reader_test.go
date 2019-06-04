package csvreader_test

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/zhnxin/csvreader"
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

func TestSnakeName(t *testing.T) {
	reader := csv.NewReader(bytes.NewReader([]byte("zhengxin,zhnxin,0,false\nxinzheng,zhnxin,1,true")))
	bean := []testStruct{}
	if err := csvreader.New().WithHeader([]string{"name", "user_name", "id", "enable"}).UnMarshal(reader, &bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	t.Log(string(b))
}

func TestLowerName(t *testing.T) {
	reader := csv.NewReader(bytes.NewReader([]byte("zhengxin,zhnxin,0,false\nxinzheng,zhnxin,1,true")))
	bean := []*testStruct{}
	if err := csvreader.New().WithHeader([]string{"NAME", "USERNAME", "ID", "ENABLE"}).UnMarshal(reader, &bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	t.Log(string(b))
}

func TestCustom(t *testing.T) {
	reader := csv.NewReader(bytes.NewReader([]byte("zhengxin,zhnxin,udp,false\nxinzheng,zhnxin,tcp,true")))
	bean := []*testStruct{}
	if err := csvreader.New().WithHeader([]string{"NAME", "USERNAME", "type", "ENABLE"}).UnMarshal(reader, &bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	t.Log(string(b))
}
