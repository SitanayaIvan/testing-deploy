package config

type ConfigBody struct {
	Db  DBAccount
	App AppAccount
}

type DBAccount struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
	DBUrl    string
}

type AppAccount struct {
	Port string
}
