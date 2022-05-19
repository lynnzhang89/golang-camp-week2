package dao

import "fmt"

type User struct {
	UserName string
	Password string
}

func GetUser(userId int, offset int, page int) ([]*User, error) {
	sql := `select xxx from user where userid = ? limit ?, ?`

	var result []*User
	err := sql.Query(conn, &result, sql, userId, offset, page)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("userId %d, offset %d, page %d, %w", userId, offset, page, err)
	}

	// ...

}
