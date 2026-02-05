package repository

import (
	"database/sql"
	"eratani_assesment_test/TestCase_3/database"
	model "eratani_assesment_test/TestCase_3/internal/domain/credit_card/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type CreditCardRepository interface {
	GetAllUser() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) (int, error)
	GetTransactions() ([]model.TransactionWithUser, error)
	GetCreditCardTypes() ([]map[string]interface{}, error)
}

type CreditCardRepositoryPostgres struct {
	DBRead  *sqlx.DB
	DBWrite *sqlx.DB
}

func ProvideUserRepository(db *database.PostgresConn) *CreditCardRepositoryPostgres {
	return &CreditCardRepositoryPostgres{
		DBRead:  db.Read,
		DBWrite: db.Write,
	}
}

func (r *CreditCardRepositoryPostgres) GetAllUser() ([]model.User, error) {
	query := `SELECT id, country, credit_card_type, credit_card_number, first_name, last_name FROM users`
	log.Info().Msg(query)
	rows, err := r.DBRead.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *CreditCardRepositoryPostgres) GetUserByID(id int) (*model.User, error) {
	query := `SELECT id, country, credit_card_type, credit_card_number, first_name, last_name FROM users WHERE id = $1`
	row := r.DBRead.QueryRowx(query, id)

	var user model.User
	if err := row.StructScan(&user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *CreditCardRepositoryPostgres) CreateUser(user *model.User) (int, error) {
	query := `INSERT INTO users (country, credit_card_type, credit_card_number, first_name, last_name)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := r.DBWrite.QueryRow(query, user.Country, user.CreditCardType, user.CreditCardNumber, user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CreditCardRepositoryPostgres) GetTransactions() ([]model.TransactionWithUser, error) {
	query := `SELECT t.id, t.id_user, t.total_buy, u.first_name, u.last_name, u.country
	          FROM user_transactions t
	          JOIN users u ON t.id_user = u.id`
	rows, err := r.DBRead.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.TransactionWithUser
	for rows.Next() {
		var t model.TransactionWithUser
		if err := rows.StructScan(&t); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (r *CreditCardRepositoryPostgres) GetCreditCardTypes() ([]map[string]interface{}, error) {
	query := `SELECT credit_card_type, COUNT(*) as total
	          FROM users
	          GROUP BY credit_card_type
	          ORDER BY total DESC`
	rows, err := r.DBRead.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var cardType string
		var total int
		if err := rows.Scan(&cardType, &total); err != nil {
			return nil, err
		}
		results = append(results, map[string]interface{}{
			"credit_card_type": cardType,
			"total":            total,
		})
	}
	return results, nil
}
