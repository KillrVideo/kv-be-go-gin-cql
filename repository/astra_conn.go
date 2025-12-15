package repository

import (
	"fmt"
	"path/filepath"

	apachegocql "github.com/apache/cassandra-gocql-driver/v2"
)

type AstraConfig struct {
	Token    string
	Keyspace string
	ScbDir   string
	Hostname string
}

func NewAstraSession(cfg AstraConfig) (*apachegocql.Session, error) {

	port := 29042
	username := "token"
	password := cfg.Token

	caPath, _ := filepath.Abs(cfg.ScbDir + "ca.crt")
	certPath, _ := filepath.Abs(cfg.ScbDir + "cert")
	keyPath, _ := filepath.Abs(cfg.ScbDir + "key")

	cluster := apachegocql.NewCluster(cfg.Hostname)
	cluster.Port = port
	cluster.Keyspace = cfg.Keyspace
	cluster.ProtoVersion = 4

	// security
	cluster.Authenticator = &apachegocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}

	cluster.SslOpts = &apachegocql.SslOptions{
		CertPath:               certPath,
		KeyPath:                keyPath,
		CaPath:                 caPath,
		EnableHostVerification: false,
	}

	cluster.Keyspace = cfg.Keyspace
	cluster.ProtoVersion = 4

	session, err2 := cluster.CreateSession()

	if err2 != nil {
		return nil, fmt.Errorf("unable to create Astra session: %w", err2)
	}

	return session, nil
}
