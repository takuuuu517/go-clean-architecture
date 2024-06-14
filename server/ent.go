package server

import (
	"cleanArchitecture/ent"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func newEntClient(dsn string) (*ent.Client, error) {
	var opt []ent.Option
	opt = append(opt, ent.Debug())
	drv, err := sql.Open(dialect.MySQL, dsn)
	if err != nil {
		return nil, err
	}

	client := ent.NewClient(append(opt, ent.Driver(drv))...)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, nil
}
