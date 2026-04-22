package db

import (
	"context"

	"github.com/google/uuid"
)

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID uuid.UUID `json:"from_account_id" binding:"required,uuid"`
	ToAccountID   uuid.UUID `json:"to_account_id" binding:"required,uuid"`
	Amount        int64 	`json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			ID: uuid.New(),
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			ID: uuid.New(),
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			ID: uuid.New(),
			AccountID: arg.FromAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		//Lock from account & Update account's balance
		fromAccount, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID: arg.FromAccountID,
			Balance: fromAccount.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}


		//Lock to account & Update account's balance
		toAccount, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID: arg.ToAccountID,
			Balance: toAccount.Balance + arg.Amount,
		})
		if err != nil {
			return err
		}
		
		return err
	})
	return  result, err
	 
}