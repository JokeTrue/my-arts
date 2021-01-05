package utils

import (
	"database/sql"
	"encoding/json"
)

type NullTimeWithJSON struct {
	sql.NullTime
}

func (n *NullTimeWithJSON) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return n.Time.MarshalJSON()
	}
	return json.Marshal(nil)
}

func (n *NullTimeWithJSON) UnmarshalJSON(value []byte) error {
	return n.Scan(value)
}
