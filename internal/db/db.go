package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Open abre a conexão com o MySQL e valida com um Ping.
//
// Intencionalmente NÃO ajusta o pool (SetMaxOpenConns, SetMaxIdleConns,
// SetConnMaxLifetime...): a v0 usa os defaults do database/sql. Ajustar
// isso é uma decisão a ser tomada quando a carga mostrar que é um gargalo.
func Open(ctx context.Context, dsn string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := conn.PingContext(pingCtx); err != nil {
		conn.Close()
		return nil, fmt.Errorf("ping: %w", err)
	}
	return conn, nil
}
