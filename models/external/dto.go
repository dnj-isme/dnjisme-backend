package external

type FetchPlayerDto struct {
	Action  string             `json:"action,omitempty"` // checksubscription or getprofile
	Payload FetchPlayerPayload `json:"payload,omitempty"`
}

type FetchPlayerPayload struct {
	AllyCode string `json:"allyCode,omitempty"` // enter ally code here
}
