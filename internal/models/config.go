package models

// Config is a struct that holds the configuration values
type Config struct {
	Server struct {
		Port string
		Host string
	}
	Database struct {
		User string
		Pass string
	}
}
