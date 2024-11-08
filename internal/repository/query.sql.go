// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"
)

const checkEmailIfMember = `-- name: CheckEmailIfMember :one
SELECT email FROM members WHERE email = ?
`

func (q *Queries) CheckEmailIfMember(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRowContext(ctx, checkEmailIfMember, email)
	err := row.Scan(&email)
	return email, err
}

const getMemberInfo = `-- name: GetMemberInfo :one
SELECT m.email, m.full_name, c.committee_name, d.division_name, p.position_name 
FROM members m
JOIN committees c ON m.committee_id = c.committee_id
JOIN divisions d ON c.division_id = d.division_id
JOIN positions p ON m.position_id = p.position_id
WHERE m.email = ?
`

type GetMemberInfoRow struct {
	Email         string
	FullName      string
	CommitteeName string
	DivisionName  string
	PositionName  string
}

func (q *Queries) GetMemberInfo(ctx context.Context, email string) (GetMemberInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getMemberInfo, email)
	var i GetMemberInfoRow
	err := row.Scan(
		&i.Email,
		&i.FullName,
		&i.CommitteeName,
		&i.DivisionName,
		&i.PositionName,
	)
	return i, err
}

const listMembers = `-- name: ListMembers :many
SELECT id, full_name, nickname, email, telegram, position_id, committee_id, college, program, discord FROM members ORDER BY email
`

func (q *Queries) ListMembers(ctx context.Context) ([]Member, error) {
	rows, err := q.db.QueryContext(ctx, listMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Member
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Nickname,
			&i.Email,
			&i.Telegram,
			&i.PositionID,
			&i.CommitteeID,
			&i.College,
			&i.Program,
			&i.Discord,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
