// Code generated by goctl. DO NOT EDIT.
package types

type Message struct {
	From        uint64 `json:"from,optional"`
	To          uint64 `json:"to"`
	Type        uint8  `json:"type"`
	ContentType uint8  `json:"contentType"`
	Time        string `json:"time"`
	Content     string `json:"content"`
}
