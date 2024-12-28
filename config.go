package main

var Config struct {
	Port   string
	JWTKey string
	DB     DBConfig
	SMTP   SMTPConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type SMTPConfig struct {
	Server   string
	Port     string
	Mail     string
	Password string
}
