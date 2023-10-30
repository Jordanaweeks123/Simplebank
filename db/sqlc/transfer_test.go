package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer
}

// Test CreateTransfer
func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

// Test GetTransfer
func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

// Test Delete Transfer
func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.Empty(t, transfer2)
}

// Test Update Transfer
func TestUpdateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID: transfer1.ID,
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, arg.Amount, transfer2.Amount)
	require.Equal(t, arg.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.ID, transfer2.ID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

// Test List Transfer
func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
