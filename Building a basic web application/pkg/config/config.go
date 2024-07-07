package config

import "github.com/alexedwards/scs/v2"

// AppConfig holds the whole application config
type AppConfig struct{
	Name string
	IsRich bool
	InProd bool
	Session *scs.SessionManager
}
