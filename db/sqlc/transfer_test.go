package db

import (
	"context"
	"math/rand"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	account2, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	transfer, err := testQueries.CreateTransfer(context.Background(), createTransferParams(account1.ID, account2.ID))
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
}

func TestGetTransfer(t *testing.T) {
	numTransfers := rand.Intn(10)

	for i := 0; i < numTransfers; i++ {
		account1, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
		require.NoError(t, err)
		require.NotEmpty(t, account1)
		account2, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
		require.NoError(t, err)
		require.NotEmpty(t, account2)

		createdTransfer, err := testQueries.CreateTransfer(context.Background(), createTransferParams(account1.ID, account2.ID))
		require.NoError(t, err)
		require.NotEmpty(t, createdTransfer)

		retrievedTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
		require.NoError(t, err)
		require.NotEmpty(t, retrievedTransfer)

		require.Equal(t, createdTransfer.Amount, retrievedTransfer.Amount)
		require.Equal(t, createdTransfer.FromAccountID, retrievedTransfer.FromAccountID)
		require.Equal(t, createdTransfer.ToAccountID, retrievedTransfer.ToAccountID)
	}
}

func TestListTransfers(t *testing.T) {
	numTransfers := int(util.RandomInt(5, 10))
	account1, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	account2, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	// generate multiple transfers between same pair of accounts
	for i := 0; i < numTransfers; i++ {
		transfer, err := testQueries.CreateTransfer(context.Background(), createTransferParams(account1.ID, account2.ID))
		require.NoError(t, err)
		require.NotEmpty(t, transfer)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        int32(numTransfers) - 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func createTransferParams(sourceAccountID, targetAccountID int64) CreateTransferParams {
	return CreateTransferParams{
		FromAccountID: sourceAccountID,
		ToAccountID:   targetAccountID,
		Amount:        rand.Int63(),
	}
}
