package model

type Response struct {
	Data     interface{} `json:"data"`
	Metadata Metadata    `json:"metadata"`
}

type Metadata struct {
	Count int `json:"count"`
}
