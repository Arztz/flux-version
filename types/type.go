package types

type Project struct {
	Project  string     `json:"project"`
	Category []Category `json:"category"`
}

type Category struct {
	Name    string    `json:"name"`
	Service []Service `json:"service"`
}

type Service struct {
	Name    string `json:"name"`
	NonProd string `json:"nonprod"`
	UAT     string `json:"uat"`
	Prod    string `json:"prod"`
}
