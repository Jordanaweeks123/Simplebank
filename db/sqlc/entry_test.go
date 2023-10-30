package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomSign() * util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, entry.Amount, arg.Amount)

	return entry
}

// Test CreateEntry
func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

// Test GetEntry
func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
	require.Equal(t, entry1.ID, entry2.ID)
}

// Test UpdateEntry
func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	// Used to test changing AccountID value
	entry2 := createRandomEntry(t)

	arg := UpdateEntryParams{
		ID:      entry1.ID,
		AccountID: entry2.AccountID,
		Amount:    util.RandomSign() * util.RandomMoney(),
	}

	entry3, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry3)

	require.Equal(t, arg.AccountID, entry3.AccountID)
	require.Equal(t, arg.Amount, entry3.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry3.CreatedAt, time.Second)
	require.Equal(t, entry1.ID, entry3.ID)
}

// Test DeleteEntry
func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

// Test ListEntry
func TestListEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
