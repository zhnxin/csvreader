# csvreader
csv reader is a simple decoder for decoding csv file to struct

# todo

- [ ] Custom parster

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