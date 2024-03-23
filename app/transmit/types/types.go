package types

type Message struct {
	Id          string `json:"id"`
	From        uint64 `json:"from,optional"`
	To          uint64 `json:"to"`
	Type        uint32 `json:"type"`
	ContentType uint32 `json:"contentType"`
	Time        string `json:"time"`
	Content     string `json:"content"`
}

type Transmit struct {
	Ip string
	Message
}
