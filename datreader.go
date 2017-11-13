package gotils

import (
	"strings"
  // "fmt"
  "github.com/djimenez/iconv-go"
)

type DatReads []DatRead
type DatRead struct {
  Name string
  ReadAt ReadAt
}
type ReadAt struct {
  StartAt int
  EndAt int
}


func DatReader(data string, prefix DatReads) map[string]interface{} {
  str := make(map[string]interface{})
  for _, q := range prefix {
    get := strings.TrimSpace(data[q.ReadAt.StartAt: q.ReadAt.EndAt])
    if get == "" {
      str[q.Name] = ""
      continue
    }
  }
  return str
}

func DatReaderDecode(data string, prefix DatReads, decodeFrom string, decodeTo string) map[string]interface{} {
  m := make(map[string]interface{})
  for _, q := range prefix {
    get := strings.TrimSpace(data[q.ReadAt.StartAt: q.ReadAt.EndAt])
    if get == "" {
      m[q.Name] = ""
      continue
    }
    str, _ := iconv.ConvertString(get, decodeFrom, decodeTo)
    m[q.Name] = str
  }
  return m
}
