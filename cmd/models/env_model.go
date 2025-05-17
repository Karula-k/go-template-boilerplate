package models

type EnvModel struct {
	DatabaseUrl string `json:"database_url"`
	JwtSecret   string `json:"jwt_secret"`
	Port        string `json:"port"`
}
