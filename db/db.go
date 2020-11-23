package db

import (
	"github.com/gocql/gocql"
	"github.com/nightLord189/ad_practice4/config"
)

func Launch(cfg *config.Config) *gocql.Session {
	cluster := gocql.NewCluster(cfg.Host)
	cluster.Keyspace = cfg.Keyspace
	cluster.Port = cfg.Port
	cluster.Consistency = gocql.Any
	session, err := cluster.CreateSession()
	if err != nil {
		panic("db connect err: " + err.Error())
	}
	return session
}
