package user

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int8   `json:"age"`
	Social Social
}

type Social struct {
	FaceBook string
	Twitter  string
}
