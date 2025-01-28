package config

import ()

var DefaultConfig = []byte(`
bool:
        enabled: true
mongo:
      uri: "mongodb://localhost:27017"
postgres:
      uri: "postgresql://postgres:1234@localhost:5432/postgres?sslmode=disable"
port:
      port: 8080
`)

type Config struct {
	Mongo    Mongo    `koanf:"mongo"`
	Postgres Postgres `koanf:"postgres"`
	Port     Port     `koanf:"port"`
	Bool     Bool     `koanf:"bool"`
}

type Port struct {
	Port int `koanf:"port"`
}

type Mongo struct {
	URI string `koanf:"uri"`
}

type Postgres struct {
	URI string `koanf:"uri"`
}

type Bool struct {
	Enabled bool `koanf:"enabled"`
}
