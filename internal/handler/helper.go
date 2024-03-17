package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func UnmarshalBody(r io.Reader, v interface{}) error {
	res, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	if err := json.Unmarshal(res, v); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	return nil
}
