package db

import (
	"context"
	"database/sql"
	"math/rand"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	accountArg := generateCreateAccountParams()
	account, err := testQueries.CreateAccount(context.Background(), accountArg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	entryArg := generateCreateEntryParams(account.ID)
	entry, err := testQueries.CreateEntry(context.Background(), entryArg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entryArg.AccountID.Int64, entry.AccountID.Int64)
	require.Equal(t, entryArg.Amount.Int64, entry.Amount.Int64)

	require.NotZero(t, entry.AccountID.Int64)
	require.NotZero(t, entry.AccountID.Int64)
}

func TestGetEntry(t *testing.T) {
	numEntries := rand.Intn(10)
	for i := 0; i < numEntries; i++ {
		accountArg := generateCreateAccountParams()
		account, err := testQueries.CreateAccount(context.Background(), accountArg)
		require.NoError(t, err)
		require.NotEmpty(t, account)

		createdEntry, err := testQueries.CreateEntry(context.Background(), generateCreateEntryParams(account.ID))
		require.NoError(t, err)
		require.NotEmpty(t, createdEntry)

		retrievedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
		require.NoError(t, err)
		require.NotEmpty(t, retrievedEntry)

		require.Equal(t, createdEntry.AccountID.Int64, retrievedEntry.AccountID.Int64)
		require.Equal(t, createdEntry.Amount.Int64, retrievedEntry.Amount.Int64)
		require.WithinDuration(t, createdEntry.CreatedAt.Time, retrievedEntry.CreatedAt.Time, time.Second)
	}
}

func TestListEntries(t *testing.T) {
	numAccounts := int(util.RandomInt(5, 10))
	for i := 0; i < numAccounts; i++ {
		createdAccount, err := testQueries.CreateAccount(context.Background(), generateCreateAccountParams())
		require.NoError(t, err)
		require.NotEmpty(t, createdAccount)

		createdEntry, err := testQueries.CreateEntry(context.Background(), generateCreateEntryParams(createdAccount.ID))
		require.NoError(t, err)
		require.NotEmpty(t, createdEntry)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: int32(numAccounts) - 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func generateCreateEntryParams(accountID int64) CreateEntryParams {
	return CreateEntryParams{
		AccountID: sql.NullInt64{
			Int64: accountID,
			Valid: true,
		},
		Amount: util.RandomSQLNullInt64(),
	}
}
