package repository

import (
	"database/sql"
	"my-go-api/internal/models"
	"time"
)

type UserRepository interface {
	GetProfileByID(userID uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User, plainPassword string) error

	// UPDATED: Added session contract signatures for multi-platform support
	SaveSession(userID uint, token string, platform string, expiresAt time.Time) error
	IsSessionValid(token string) (bool, error)
	DeleteSession(userID uint, platform string, allDevices bool, currentToken string) error
}

type userRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{DB: db}
}

// 1. IMPLEMENTATION FOR GetProfileByID
func (r *userRepo) GetProfileByID(userID uint) (*models.User, error) {
	user := &models.User{}

	// Fetch core user data
	userQuery := `SELECT id, first_name, last_name, email, phone_number, profile_image, is_active FROM users WHERE id = $1`
	err := r.DB.QueryRow(userQuery, userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.ProfileImage, &user.IsActive)
	if err != nil {
		return nil, err
	}

	// Fetch dynamic roles assigned to user via mapping join table
	roleQuery := `SELECT r.name FROM roles r 
	              JOIN user_roles ur ON r.id = ur.role_id 
	              WHERE ur.user_id = $1`

	rows, err := r.DB.Query(roleQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var roleName string
		if err := rows.Scan(&roleName); err == nil {
			user.Roles = append(user.Roles, roleName)
		}
	}

	return user, nil
}

// 2. IMPLEMENTATION FOR GetByEmail
func (r *userRepo) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, first_name, last_name, email, password_hash, phone_number, is_active FROM users WHERE email = $1`

	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.PhoneNumber, &user.IsActive)
	if err != nil {
		return nil, err
	}

	// Fetch user roles
	roleQuery := `SELECT r.name FROM roles r JOIN user_roles ur ON r.id = ur.role_id WHERE ur.user_id = $1`
	rows, err := r.DB.Query(roleQuery, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var roleName string
		if err := rows.Scan(&roleName); err == nil {
			user.Roles = append(user.Roles, roleName)
		}
	}
	return user, nil
}

// 3. IMPLEMENTATION FOR Create (With secure SQL Transactions)
func (r *userRepo) Create(user *models.User, plainPassword string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Step A: Insert user record into the users table
	userQuery := `INSERT INTO users (first_name, last_name, email, password_hash, phone_number, profile_image, is_active) 
	              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = tx.QueryRow(userQuery, user.FirstName, user.LastName, user.Email, user.PasswordHash, user.PhoneNumber, user.ProfileImage, true).Scan(&user.ID)
	if err != nil {
		return err
	}

	// Step B: Ensure the "USER" role exists in your master roles directory table
	var roleID uint
	roleQuery := `SELECT id FROM roles WHERE name = $1`
	err = tx.QueryRow(roleQuery, "USER").Scan(&roleID)
	if err != nil {
		err = tx.QueryRow(`INSERT INTO roles (name) VALUES ($1) RETURNING id`, "USER").Scan(&roleID)
		if err != nil {
			return err
		}
	}

	// Step C: Map the newly created user ID with the target Role ID inside the junction table
	mappingQuery := `INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)`
	_, err = tx.Exec(mappingQuery, user.ID, roleID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// 4. UPDATED: IMPLEMENTATION FOR SaveSession
func (r *userRepo) SaveSession(userID uint, token string, platform string, expiresAt time.Time) error {
	query := `INSERT INTO user_sessions (user_id, token, platform, expires_at) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, userID, token, platform, expiresAt)
	return err
}

// 5. UPDATED: IMPLEMENTATION FOR IsSessionValid
func (r *userRepo) IsSessionValid(token string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM user_sessions WHERE token = $1 AND expires_at > NOW())`
	err := r.DB.QueryRow(query, token).Scan(&exists)
	return exists, err
}

// 6. UPDATED: IMPLEMENTATION FOR DeleteSession
func (r *userRepo) DeleteSession(userID uint, platform string, allDevices bool, currentToken string) error {
	if allDevices {
		// Terminate all sessions across every single device/platform profile completely
		query := `DELETE FROM user_sessions WHERE user_id = $1`
		_, err := r.DB.Exec(query, userID)
		return err
	}

	// Terminate only the specific device type platform or current explicit token signature
	query := `DELETE FROM user_sessions WHERE user_id = $1 AND (platform = $2 OR token = $3)`
	_, err := r.DB.Exec(query, userID, platform, currentToken)
	return err
}
