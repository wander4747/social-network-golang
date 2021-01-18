package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications = repository publication
type Publications struct {
	db *sql.DB
}

// NewRepositoriesPublication = create repository publication
func NewRepositoriesPublication(db *sql.DB) *Publications {
	return &Publications{db}
}

// Create = Create new publication
func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO publications (title, content, author_id) VALUES (?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(publication.Title, publication.Content, publication.AuthorID)

	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

// Find = get publication by id
func (repository Publications) Find(ID uint64) (models.Publication, error) {
	line, erro := repository.db.Query(`
		SELECT publications.id, title, content, publications.created_at, users.id, users.nick
		  FROM publications
	INNER JOIN users ON users.id = publications.author_id 
     	 WHERE publications.id = ?`,
		ID,
	)

	if erro != nil {
		return models.Publication{}, erro
	}

	defer line.Close()

	var publication models.Publication

	if line.Next() {
		if erro = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.CreatedAt,
			&publication.AuthorID,
			&publication.AuthorNick,
		); erro != nil {
			return models.Publication{}, erro
		}
	}

	return publication, nil
}

//All return all user with filter name or nick
func (repository Publications) All(ID uint64) ([]models.Publication, error) {
	lines, erro := repository.db.Query(`
	SELECT DISTINCT p.*, u.nick FROM publications p 
	     INNER JOIN users u on u.id = p.author_id 
	     INNER JOIN followers s on p.author_id = s.user_id 
	     WHERE u.id = ? OR s.follower_id = ?
	  ORDER BY 1 desc
	`, ID, ID)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if erro = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Update data publication
func (repository Publications) Update(publicationID uint64, publication models.Publication) error {
	statement, erro := repository.db.Prepare("UPDATE publications SET title = ?, content = ? WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publication.Title, publication.Content, publicationID); erro != nil {
		return erro
	}

	return nil
}

// Delete data publication
func (repository Publications) Delete(publicationID uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM publications WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}

// FindPublicationByUser = Get all publications by user
func (repository Publications) FindPublicationByUser(userID uint64) ([]models.Publication, error) {
	lines, erro := repository.db.Query(`
		SELECT p.*, u.nick FROM publications p
		INNER  JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?
	`, userID)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if erro = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Like = add a like publication
func (repository Publications) Like(publicationID uint64) error {
	statement, erro := repository.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}

// Unlike = add a ulike publication
func (repository Publications) Unlike(publicationID uint64) error {
	statement, erro := repository.db.Prepare(`
		UPDATE publications SET likes = 
		CASE WHEN likes > 0 THEN likes - 1 
		ELSE 0 END
		WHERE id = ?`)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}
