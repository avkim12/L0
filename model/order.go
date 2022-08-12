package model

import "encoding/json"

type Order struct {
	Id    string          `json:"id"`
	Model json.RawMessage `json:"model"`
}