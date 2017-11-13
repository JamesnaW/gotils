package gotils

import (
  "strings"
  "github.com/djimenez/iconv-go"
  "io"
  "bytes"
  // "log"
  // "gopkg.in/iconv.v1"
)

// func ReadLines(buff io.Reader, decodeFrom string, decodeTo string) ([]string, error) {
//   if buff == nil {
//     return nil, &errorString{"No input!"}
//   }

// 	var lines []string
//   // decode, err := iconv.NewReader(buff, decodeFrom, decodeTo)

//   cd, err := iconv.Open(decodeTo, decodeFrom) // convert utf-8 to gbk
// 	if err != nil {
// 		return nil, nil
// 	}
// 	defer cd.Close()

//   decode := iconv.NewReader(cd, buff, 0)

//   if err != nil {
//     return nil, err
//   }
//   buf := new(bytes.Buffer)
//   buf.ReadFrom(decode)

//   for _, line := range strings.Split(buf.String(), "\n") {
//     log.Println(line)
//     lines = append(lines, line)
//   }

//   return lines, nil


// }

func ReadLines(buff io.Reader, decodeFrom string, decodeTo string) ([]string, error) {
  if buff == nil {
    return nil, &errorString{"No input!"}
  }

  var lines []string
  if decodeFrom != "" && decodeTo != "" {
    decode, err := iconv.NewReader(buff, decodeFrom, decodeTo)
    if err != nil {
      return nil, err
    }
    buf := new(bytes.Buffer)
    buf.ReadFrom(decode)
    for _, line := range strings.Split(buf.String(), "\n") {
      // log.Println(line)
      lines = append(lines, line)
    }
  } else {
    buf := new(bytes.Buffer)
    buf.ReadFrom(buff)
    for _, line := range strings.Split(buf.String(), "\n") {
      // log.Println(line)
      lines = append(lines, line)
    }
  }


  return lines, nil
}

func ReadLinesMap(buff io.Reader, decodeFrom string, decodeTo string) (map[int]string, error) {
  if buff == nil {
    return nil, &errorString{"No input!"}
  }

	lines := make(map[int]string)
  decode, err := iconv.NewReader(buff, decodeFrom, decodeTo)
  if err != nil {
    return nil, err
  }
  buf := new(bytes.Buffer)
  buf.ReadFrom(decode)

  for n, line := range strings.Split(buf.String(), "\n") {
      lines[n] = line
  }

  return lines, nil
}
