// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package generated

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int64              `json:"id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
