// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package override

import (
	"database/sql"

	orm "database/sql"
	fuid "github.com/gofrs/uuid"
	uuid "github.com/gofrs/uuid"
	null "github.com/volatiletech/null/v8"
	null_v4 "gopkg.in/guregu/null.v4"
)

type Bar struct {
	ID      uuid.UUID
	OtherID fuid.UUID
	MoreID  fuid.UUID
	Age     sql.NullInt32
	Balance sql.NullFloat64
	Bio     sql.NullString
	About   sql.NullString
}

type Foo struct {
	ID      uuid.UUID
	OtherID fuid.UUID
	Age     orm.NullInt32
	Balance null.Float32
	Bio     null_v4.String
	About   *string
}
