package config

type Config struct {
	DBUser      string
	DBPassword  string
	DBURL       string
	DBPort      string
	DBName      string
	ClientID    string
	APISecret   string
	BearerToken string
	Namespace   string
	Repository  string
}

var AppConfig Config