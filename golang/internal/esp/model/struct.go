package model

type Wifi2 struct {
	IndexStart int    `json:"IndexStart"`
	IndexEnd   int    `json:"IndexEnd"`
	Valor      string `json:"Valor"`
	Output     string `json:"Output"`
}

type Wifi struct {
	Count    int       `json:"count"`
	Networks []Network `json:"networks"`
}
type Network struct {
	Ssi       string `json:"ssi"`
	Rssi      int    `json:"rssi"`
	Encripted bool   `json:"encripted"`
}

type RequestNetwork struct {
	Network  string `json:"network"`
	Password string `json:"password"`
}
