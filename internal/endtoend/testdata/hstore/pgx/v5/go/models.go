// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package hstore

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Foo struct {
	Bar pgtype.Hstore
	Baz pgtype.Hstore
}
