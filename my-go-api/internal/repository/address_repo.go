package repository

import (
	"database/sql"
	"my-go-api/internal/models"
)

type AddressRepository interface {
	GetByUserID(userID uint) ([]models.Address, error)
	Create(addr *models.Address) error
	Update(addr *models.Address, addressID uint, userID uint) (int64, error)
	Delete(addressID uint, userID uint) (int64, error)
}

type addressRepo struct {
	DB *sql.DB
}

func NewAddressRepository(db *sql.DB) AddressRepository {
	return &addressRepo{DB: db}
}

func (r *addressRepo) GetByUserID(userID uint) ([]models.Address, error) {
	query := `SELECT id, user_id, address_line1, address_line2, locality, city, district, state, country, is_active 
	          FROM addresses WHERE user_id = $1 AND is_active = true`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var addr models.Address
		err := rows.Scan(&addr.ID, &addr.UserID, &addr.AddressLine1, &addr.AddressLine2, &addr.Locality, &addr.City, &addr.District, &addr.State, &addr.Country, &addr.IsActive)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}
	return addresses, nil
}

func (r *addressRepo) Create(addr *models.Address) error {
	query := `INSERT INTO addresses (user_id, address_line1, address_line2, locality, city, district, state, country) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	return r.DB.QueryRow(query, addr.UserID, addr.AddressLine1, addr.AddressLine2, addr.Locality, addr.City, addr.District, addr.State, addr.Country).Scan(&addr.ID)
}

func (r *addressRepo) Update(addr *models.Address, addressID uint, userID uint) (int64, error) {
	query := `UPDATE addresses SET address_line1=$1, address_line2=$2, locality=$3, city=$4, district=$5, state=$6, country=$7, updated_at=NOW() 
	          WHERE id=$8 AND user_id=$9`

	res, err := r.DB.Exec(query, addr.AddressLine1, addr.AddressLine2, addr.Locality, addr.City, addr.District, addr.State, addr.Country, addressID, userID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *addressRepo) Delete(addressID uint, userID uint) (int64, error) {
	query := `DELETE FROM addresses WHERE id = $1 AND user_id = $2`
	res, err := r.DB.Exec(query, addressID, userID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
