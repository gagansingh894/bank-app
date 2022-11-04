package db

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"simple-bank/util"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := generateCreateAccountParams()

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	numAccounts := rand.Intn(10)
	for i := 0; i < numAccounts; i++ {
		createdAccount, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
		require.NoError(t, err)
		retrievedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
		require.NoError(t, err)

		require.Equal(t, createdAccount.ID, retrievedAccount.ID)
		require.Equal(t, createdAccount.Owner, retrievedAccount.Owner)
		require.Equal(t, createdAccount.Balance, retrievedAccount.Balance)
		require.Equal(t, createdAccount.Currency, retrievedAccount.Currency)
		require.WithinDuration(t, createdAccount.CreatedAt, retrievedAccount.CreatedAt, time.Second)

	}
}

func TestUpdateAccount(t *testing.T) {
	createdAccount, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
	require.NoError(t, err)

	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: util.RandomInt(0, 100),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, createdAccount.ID, updatedAccount.ID)
	require.Equal(t, createdAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, createdAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	createdAccount, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
	require.NoError(t, err)

	err = testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)

	account, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}

func TestListAccounts(t *testing.T) {
	numAccounts := int(util.RandomInt(5, 10))
	for i := 0; i < numAccounts; i++ {
		createdAccount, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
		require.NoError(t, err)
		require.NotEmpty(t, createdAccount)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: int32(numAccounts) - 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func generateCreateAccountParams() CreateAccountParams {
	return CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomInt(0, 10000),
		Currency: util.RandomCurrency(),
	}
}
