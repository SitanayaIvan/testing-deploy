package config

type ConfigBody struct {
	Db  DBAccount
	App AppAccount
}

type DBAccount struct {
	DBUrl string
}

type AppAccount struct {
	Port string
}
