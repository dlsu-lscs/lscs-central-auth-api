// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"database/sql"
)

type Committee struct {
	CommitteeID         string
	CommitteeName       string
	CommitteeHead       sql.NullInt32
	CommitteeDivisionID sql.NullString
}

type Division struct {
	DivisionID   string
	DivisionName string
	DivisionHead sql.NullInt32
}

type Member struct {
	ID          int32
	FullName    string
	Nickname    sql.NullString
	Email       string
	Telegram    sql.NullString
	PositionID  sql.NullString
	CommitteeID sql.NullString
}

type Position struct {
	PositionID   string
	PositionName string
}