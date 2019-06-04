# csvreader
csv reader is a simple decoder for decoding csv file to struct

# todo

- [x] Custom parster
- [ ] pointer attribute

# install
```
go get github.com/zhnxin/csvreader
```

# usage

## simple
As the default,the first line of csv file would be regared as the header.Take a look as follows.

`file.csv`
```csv
hosname,ip
redis,172.17.0.2
mariadb,172.17.0.3
```

```go
type Info struct{
    Hostname string
    IP string
}

//struc slice
infos := []Info{}
_ = csvreader.New().UnMarshalFile("file.csv",&infos)
body,_ := json.Marshal(infos)
fmt.Println(string(body))

//point slice
infos := []*Info{}
_ = csvreader.New().UnMarshalFile("file.csv",&infos)
body,_ := json.Marshal(infos)
fmt.Println(string(body))
```

If the csv file do not contain headers,you can also set it.

```
_ = csvreader.New().WithHeader([]string{"hostname","ip"}).UnMarshalFile("file.csv",&infos)
```

## custom parseter

Just like enum,we need implement our own parster. The example is shown as follows.

```go
type NetProtocol uint32
const(
    NetProtocol_TCP NetProtocol = iota
    NetProtocol_UDP
    NetProtocol_DCCP
    NetProtocol_SCTP
)

type ServiceInfo struct{
    Host string
    Port string
    Protocol NetProtocol
}
```

It's inconvenient to edit a csv file with the custome enum type. It's greate to convert _top_ or _TCP_ to NetProtocol_TCP automatically.Just to implement the _CsvMarshal_ interface


    type CsvMarshal interface {
	    FromString(string) error
    }

As the case above, the simple solution is this.
```go
func (p *NetProtocol)FromString(protocol string) error{
    switch strings.ToLower(protocol){
        case "tcp":
            *p = NetProtocol_TCP
        case "udp":
            *p = NetProtocol_UDP
        case "dccp":
            *p = NetProtocol_DCCP
        case "sctp":
            *p = NetProtocol_SCTP
        default:
            return fmt.Errorf("unknown protocoal:%s",protocol)
    }
    return nil
}

```
There is another exampler you can refer to [TestCustom](./reader_test.go#TestCustom)