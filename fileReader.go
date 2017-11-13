package gotils

import (
  "os"
  "io"
)

func Open(filepath string) (io.Reader, error) {
  if filepath == "" {
    return nil, &errorString{"No file path!"}
  }
	buff, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
  return buff, nil
}
