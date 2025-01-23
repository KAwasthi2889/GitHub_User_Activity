package main

type Event struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload,omitempty"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action       string `json:"action,omitempty"`
	Created_type string `json:"ref_type,omitempty"`
	Reason       string `json:"reason,omitempty"`
	Size         int    `json:"distinct_size,omitempty"`
	Number       int    `json:"number,omitempty"`
}
