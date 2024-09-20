package blockchain

import (
	"fmt"

	"github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models/blockchain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionSerialiser struct{}

func (*TransactionSerialiser) ReadTransaction(db *gorm.DB, transaction_id uuid.UUID) (*blockchain.Transaction, error) {
	var transaction []blockchain.Transaction
	db.Model(&blockchain.Transaction{}).Find(&transaction, &blockchain.Transaction{TransactionID: transaction_id})
	if len(transaction) != 1 {
		return &blockchain.Transaction{}, fmt.Errorf("invalid transaction id")
	}
	return &transaction[0], nil
}

func (t *TransactionSerialiser) CreateTransaction(db *gorm.DB, sender_id uuid.UUID, receiver_id uuid.UUID, title string, description string, amount float32) (string, error) {
	var status string
	status, _ = constants.CONTRACT_DRAFT.String()

	return database.CreateModel(db, &blockchain.Transaction{
		ID:                uuid.New(),
		TransactionID:     uuid.New(),
		SenderID:          sender_id,
		ReceiverID:        receiver_id,
		Title:             title,
		Description:       description,
		SenderSigniture:   "String",
		ReceiverSigniture: "String",
		TransactionStatus: status,
		Amount:            amount,
		Salt_value:        uuid.New(),
	})
}

func (*TransactionSerialiser) UpdateTransaction(db *gorm.DB, id uuid.UUID, transaction_data map[string]interface{}) (string, error) {
	var transaction blockchain.Transaction
	err := db.Model(&blockchain.Transaction{}).First(&transaction, id).Error
	if err != nil {
		return "", err
	}
	db.Model(transaction).Updates(transaction_data)
	return transaction.String(), nil
}

func (*TransactionSerialiser) DeleteTransaction(db *gorm.DB, id uuid.UUID) (string, error) {
	err := db.Delete(&blockchain.Transaction{}, id).Error
	if err != nil {
		return "", err
	}
	return id.String() + " Deleted", nil
}
