package repositories

import (
	"fmt"
	"log"

	"github.com/gitish/polyglot_training/models"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user-related data operations.
type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

// userRepositoryImpl implements the UserRepository interface using gorm.DB
type userRepositoryImpl struct {
	dbConn *gorm.DB
}

// NewUserRepository creates a new instance of userRepositoryImpl.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{dbConn: db}
}

// CreateUser adds a new user to the database.
func (repo *userRepositoryImpl) CreateUser(user models.User) (models.User, error) {
	result := repo.dbConn.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		log.Printf("Failed to insert user due to %v", result.Error)
		return models.User{}, fmt.Errorf("failed to create user")
	}
	log.Printf("Successfully inserted %d user with id %d", result.RowsAffected, user.ID)

	return user, nil
}

// GetAllUsers retrieves all users from the database.
func (repo *userRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := repo.dbConn.Find(&users) // Get all records (SELECT * FROM users;)
	if result.Error != nil {
		log.Printf("Failed to get all users due to %v", result.Error)
		return []models.User{}, fmt.Errorf("failed to get user(s)")
	}
	log.Printf("Successfully retrieved %d user(s)", result.RowsAffected)

	return users, nil
}

// GetUserByID retrieves a user by their ID.
func (repo *userRepositoryImpl) GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := repo.dbConn.First(&user, id) // find product with integer primary key (SELECT * FROM users WHERE id = 10;)
	if result.Error != nil {
		log.Printf("Failed to get user by id (%d) due to %v", id, result.Error)
		return models.User{}, fmt.Errorf("failed to get user")
	}
	log.Printf("Successfully retrieved %d user", result.RowsAffected)

	return user, nil
}

// UpdateUser updates an existing user in the database.
func (repo *userRepositoryImpl) UpdateUser(user models.User) (models.User, error) {
	result := repo.dbConn.Save(&user) // Save is a combination function. If save value does not contain primary key, it will execute Create, otherwise it will execute Update (with all fields) (INSERT INTO `users` (`name`,`email`,`update_at`) VALUES ("binzhu",binzhu@mail.com,"0000-00-00 00:00:00")   OR   UPDATE `users` SET `name`="jinzhu", `email`="jinzhu@email.com",`update_at`="0000-00-00 00:00:00" WHERE `id` = 1)
	if result.Error != nil {
		log.Printf("Failed to update user with id %d due to %v", user.ID, result.Error)
		return models.User{}, fmt.Errorf("failed to update user")
	}
	log.Printf("Successfully updated %d user with id %d", result.RowsAffected, user.ID)

	return user, nil
}

// DeleteUser deletes a user from the database by their ID.
func (repo *userRepositoryImpl) DeleteUser(id uint) error {
	result := repo.dbConn.Delete(&models.User{}, id) // DELETE FROM users WHERE id = 10;
	if result.Error != nil {
		log.Printf("Failed to delete user with id %d due to %v", id, result.Error)
		return fmt.Errorf("failed to delete user")
	}
	log.Printf("Successfully deleted %d user with id %d", result.RowsAffected, id)

	return nil
}
