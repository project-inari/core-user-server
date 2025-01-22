package repository

import (
	"database/sql"

	"github.com/project-inari/core-user-server/dto"
)

type databaseRepository struct {
	database string
	client   *sql.DB
}

// DatabaseRepositoryConfig represents the configuration for wiremock API repository
type DatabaseRepositoryConfig struct {
	Database string
}

// DatabaseRepositoryDependencies represents the dependencies for wiremock API repository
type DatabaseRepositoryDependencies struct {
	Client *sql.DB
}

// NewDatabaseRepository creates a new wiremock API repository
func NewDatabaseRepository(c DatabaseRepositoryConfig, d DatabaseRepositoryDependencies) DatabaseRepository {
	return &databaseRepository{
		database: c.Database,
		client:   d.Client,
	}
}

// CreateNewUser creates a new user in the database
func (r *databaseRepository) CreateNewUser(newUser dto.UserEntity) error {
	query := `
		INSERT INTO tbl_users (username, uid, first_name, last_name, phone_no, email, selected_locale)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.client.Exec(query, newUser.Username, newUser.UID, newUser.FirstName, newUser.LastName, newUser.PhoneNo, newUser.Email, newUser.SelectedLocale)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByUsername gets a user by username from the database
func (r *databaseRepository) GetUserByUsername(username string) (*dto.UserEntity, error) {
	query := `
		SELECT username, uid, first_name, last_name, phone_no, email, selected_locale
		FROM tbl_users
		WHERE username = ?
	`

	user := new(dto.UserEntity)
	err := r.client.QueryRow(query, username).Scan(&user.Username, &user.UID, &user.FirstName, &user.LastName, &user.PhoneNo, &user.Email, &user.SelectedLocale)
	if err != nil {
		return &dto.UserEntity{}, err
	}

	return user, nil
}
