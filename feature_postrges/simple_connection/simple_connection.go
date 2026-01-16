package simple_connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func TestDb() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:5636@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		panic("ping failed: " + err.Error())
	}

	fmt.Println("Я успешно подключился к БД")
}
