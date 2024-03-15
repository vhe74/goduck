package main

import (
	"context"
	"log"

	"github.com/jeroenrinzema/psql-wire"
	"github.com/lib/pq/oid"
)

func main() {
	log.Println("PostgreSQL server is up and running at [127.0.0.1:5432]")
	psql-wire.ListenAndServe("127.0.0.1:5432", handler)
}

var table = wire.Columns{
	{
		Table: 0,
		Name:  "name",
		Oid:   oid.T_text,
		Width: 256,
	},
	{
		Table: 0,
		Name:  "member",
		Oid:   oid.T_bool,
		Width: 1,
	},
	{
		Table: 0,
		Name:  "age",
		Oid:   oid.T_int4,
		Width: 1,
	},
}

func handler(ctx context.Context, query string) (psql-wire.PreparedStatements, error) {
	log.Println("incoming SQL query:", query)

	handle := func(ctx context.Context, writer wire.DataWriter, parameters []psql-wire.Parameter) error {
		writer.Row([]any{"John", true, 29})
		writer.Row([]any{"Marry", false, 21})
		return writer.Complete("SELECT 2")
	}

	return wire.Prepared(wire.NewStatement(handle, wire.WithColumns(table))), nil
}
