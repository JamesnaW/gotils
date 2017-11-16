package gotils

import (
	"io"
  "encoding/csv"
  "reflect"
  "strconv"
  "time"
  "strings"
)

var layout string = "2006-01-02"

func CsvReader(buff io.Reader) ([][]string, error) {
  if buff == nil {
    return nil, &errorString{"No buffer!"}
  }
  r := csv.NewReader(buff)
  return r.ReadAll()
}

func CsvToMap(records [][]string) ([]map[string]interface{}, error) {
  var a []map[string]interface{}
  if len(records) == 0 {
    return nil, &errorString{"No csv records!"}
  }
  for i, e := range records {
    if i != 0 {
      // r, err := ParseStruct(setVal(records[0], e), result)
      // if err != nil {
      //   continue
      // }
      r := setVal(records[0], e)
      a = append(a, r)
    }
  }

  return a, nil
}

func setKey(keys []string) map[string]interface{} {
  m := make(map[string]interface{})
  for _, e := range keys {
    if e == "" {
      continue
    }
    m[e] = nil
  }
  return m
}
func setVal(keys []string, val []string) map[string]interface{} {
  m := setKey(keys)
  for i, e := range val {
    if strings.Trim(e, "") == "" {
      continue
    }
    m[keys[i]] = val[i]
  }
  return m
}

// func mapValue(value []string, obj interface{}) interface{} {
//   // v := reflect.ValueOf(&obj)
//   // for i, e := range value {
//     // setType(v.Field(i), e)
//     // v.Field(i).Set(reflect.ValueOf(e))

//   // }
//   val := reflect.ValueOf(&obj).Elem()

//   for i := 0; i < val.NumField(); i++ {
//     valueField := val.Field(i)
//     typeField := val.Type().Field(i)
//     // if key == strings.ToLower(typeField.Name) {
//     //   return valueField.Interface().(string)
//     // }
//   }
//   return &obj
// }

func setType(f reflect.Value, v string) {
  switch f.Interface().(type){
  case string:
    f.SetString(v)
  case int:
    i, _ := strconv.ParseInt(v, 64, 10)
    f.SetInt(i)
  case float64:
    i, _ := strconv.ParseFloat(v, 10)
    f.SetFloat(i)
  case time.Time:
    t, _ :=time.Parse(layout, v)
    f.Set(reflect.ValueOf(t))
  }
}
