package domain

import (
	"database/sql"
	"time"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	Username   string         `db:"username"`
	CustomerId string         `db:"customer_id"`
	Accounts   sql.NullString `db:"accounts"`
	Role       string         `db:"role"`
}
