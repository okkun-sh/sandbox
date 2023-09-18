package sample_bun_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	sample_bun "sample-bun"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func TestFindByName(t *testing.T) {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Addr:   fmt.Sprintf("%s:%v", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		Net:    "tcp",
	}

	m, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	db := bun.NewDB(m, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	db.RegisterModel((*sample_bun.User)(nil))

	tests := []struct {
		name string
		arg  string
		want func(*dbfixture.Fixture, *testing.T) *sample_bun.User
	}{
		{
			name: "simple",
			arg:  "Bar",
			want: func(f *dbfixture.Fixture, t *testing.T) *sample_bun.User {
				u := f.MustRow("User.bar").(*sample_bun.User)
				return u
			},
		},
		{
			name: "simple2",
			arg:  "Bar",
			want: func(f *dbfixture.Fixture, t *testing.T) *sample_bun.User {
				u, err := f.Row("User.bar")
				if err != nil {
					t.Fatal(err)
				}
				return u.(*sample_bun.User)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			uRepo := sample_bun.NewUserRepository(ctx, db)

			fixture := dbfixture.New(db, dbfixture.WithTruncateTables())
			if err := fixture.Load(ctx, os.DirFS("./testdata"), "fixture.yml"); err != nil {
				t.Fatal(err)
			}

			user, err := uRepo.FindByName(tt.arg)
			if err != nil {
				t.Fatal(err)
			}

			if user.Name != tt.want(fixture, t).Name {
				t.Errorf("got %v, want %v", user.Name, tt.want(fixture, t).Name)
			}
		})
	}
}
