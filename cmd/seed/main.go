// Seed pequeno da Fase 0: só valida que schema + API funcionam com dados.
//
// Ingênuo de propósito: INSERTs individuais e distribuição uniforme.
// Seed em massa (lote, power-law/Zipfian) é assunto da Fase 2.
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"time"

	"github.com/meninn/db-scale-journal/internal/db"
)

func main() {
	users := flag.Int("users", 1000, "quantidade de usuários")
	postsPerUser := flag.Int("posts-per-user", 5, "posts por usuário")
	followsPerUser := flag.Int("follows-per-user", 20, "usuários seguidos por usuário")
	truncate := flag.Bool("truncate", true, "limpa as tabelas antes de popular")
	flag.Parse()

	if *followsPerUser >= *users {
		log.Fatalf("follows-per-user (%d) precisa ser menor que users (%d)", *followsPerUser, *users)
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "app:app@tcp(localhost:3306)/journal?parseTime=true"
	}

	ctx := context.Background()
	conn, err := db.Open(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if *truncate {
		for _, t := range []string{"follows", "posts", "users"} {
			must(exec(ctx, conn, "TRUNCATE TABLE "+t))
		}
	}

	start := time.Now()

	ids := make([]int64, 0, *users)
	for i := range *users {
		res, err := conn.ExecContext(ctx, "INSERT INTO users (username) VALUES (?)", fmt.Sprintf("user_%d", i))
		must(err)
		id, err := res.LastInsertId()
		must(err)
		ids = append(ids, id)
	}
	log.Printf("users: %d", len(ids))

	nPosts := 0
	for _, uid := range ids {
		for j := range *postsPerUser {
			must(exec(ctx, conn, "INSERT INTO posts (user_id, body) VALUES (?, ?)",
				uid, fmt.Sprintf("post %d do usuário %d", j, uid)))
			nPosts++
		}
	}
	log.Printf("posts: %d", nPosts)

	nFollows := 0
	for _, follower := range ids {
		seen := map[int64]bool{follower: true}
		for len(seen) <= *followsPerUser {
			followee := ids[rand.IntN(len(ids))]
			if seen[followee] {
				continue
			}
			seen[followee] = true
			must(exec(ctx, conn, "INSERT INTO follows (follower_id, followee_id) VALUES (?, ?)", follower, followee))
			nFollows++
		}
	}
	log.Printf("follows: %d", nFollows)

	log.Printf("concluído em %s", time.Since(start).Round(time.Millisecond))
}

func exec(ctx context.Context, conn *sql.DB, q string, args ...any) error {
	_, err := conn.ExecContext(ctx, q, args...)
	return err
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
