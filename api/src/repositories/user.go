package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users = repository user
type Users struct {
	db *sql.DB
}

// NewRepositoriesUser = create repository user
func NewRepositoriesUser(db *sql.DB) *Users {
	return &Users{db}
}

// Create = create new user
func (repository Users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"INSERT INTO users (name, nick, email, password) values(?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

//Search = return all user with filter name or nick
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repository.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// Find = get user by id
func (repository Users) Find(ID uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		ID,
	)

	if erro != nil {
		return models.User{}, erro
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

//Update = update data user
func (repository Users) Update(ID uint64, user models.User) error {
	statement, erro := repository.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Delete = remove user
func (repository Users) Delete(ID uint64) error {
	statement, erro := repository.db.Prepare(
		"DELETE FROM  users WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// GetEmail = search user by email
func (repository Users) GetEmail(email string) (models.User, error) {
	line, erro := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)

	if erro != nil {
		return models.User{}, erro
	}

	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// Follow = Follow users
func (repository Users) Follow(userID, followerID uint64) error {
	statement, erro := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) values(?, ?)",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

// UnFollow = Unfollow users
func (repository Users) UnFollow(userID, followerID uint64) error {
	statement, erro := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

// Followers = Followers in the users
func (repository Users) Followers(userID uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM users u INNER JOIN followers s ON u.id = s.follower_id  WHERE s.user_id  = ?`,
		userID,
	)

	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// Following = Following in the users
func (repository Users) Following(userID uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM users u INNER JOIN followers s ON u.id = s.user_id  WHERE s.follower_id  = ?`,
		userID,
	)

	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchPassword = Search password by id
func (repository Users) SearchPassword(userID uint64) (string, error) {
	line, erro := repository.db.Query("SELECT password FROM users WHERE id = ?", userID)

	if erro != nil {
		return "", erro
	}

	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Password); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil
}

// UpdatePassword = Update password user
func (repository Users) UpdatePassword(userID uint64, password string) error {
	statement, erro := repository.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(password, userID); erro != nil {
		return erro
	}

	return nil
}
