package domain

type (
	// Cursor はページネーションのカーソルです。
	Cursor string
)

// Empty はカーソルが空か判定します。
func (c Cursor) Empty() bool {
	return c == ""
}
