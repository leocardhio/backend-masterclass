package db

import (
	"context"
	"database/sql"
	"masterclass/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	tf, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tf)

	require.Equal(t, arg.FromAccountID, tf.FromAccountID)
	require.Equal(t, arg.ToAccountID, tf.ToAccountID)
	require.Equal(t, arg.Amount, tf.Amount)

	require.NotZero(t, tf.ID)
	require.NotZero(t, tf.CreatedAt)

	return tf
}

func createSpecifiedAccountTransfer(t *testing.T, sender, receiver Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: sender.ID,
		ToAccountID:   receiver.ID,
		Amount:        util.RandomMoney(),
	}

	tf, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tf)

	require.Equal(t, arg.FromAccountID, tf.FromAccountID)
	require.Equal(t, arg.ToAccountID, tf.ToAccountID)
	require.Equal(t, arg.Amount, tf.Amount)

	require.NotZero(t, tf.ID)
	require.NotZero(t, tf.CreatedAt)

	return tf
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	tf1 := createRandomTransfer(t)

	tf2, err := testQueries.GetTransfer(context.Background(), tf1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tf2)

	require.Equal(t, tf1.ID, tf2.ID)
	require.Equal(t, tf1.FromAccountID, tf2.FromAccountID)
	require.Equal(t, tf1.ToAccountID, tf2.ToAccountID)
	require.Equal(t, tf1.Amount, tf2.Amount)

	require.WithinDuration(t, tf1.CreatedAt, tf2.CreatedAt, time.Second)
}

func TestGetTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := GetTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	tfs, err := testQueries.GetTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tfs, 5)

	for _, tf := range tfs {
		require.NotEmpty(t, tf)
	}
}

func TestGetTransfersBySenderId(t *testing.T) {
	sender := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		receiver := createRandomAccount(t)
		createSpecifiedAccountTransfer(t, sender, receiver)
	}

	arg := GetTransfersBySenderIdParams{
		FromAccountID: sender.ID,
		Limit:         5,
		Offset:        5,
	}

	tfs, err := testQueries.GetTransfersBySenderId(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tfs, 5)

	for _, tf := range tfs {
		require.NotEmpty(t, tf)
		require.Equal(t, sender.ID, tf.FromAccountID)
	}
}

func TestGetTransfersByReceiverId(t *testing.T) {
	receiver := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		sender := createRandomAccount(t)
		createSpecifiedAccountTransfer(t, sender, receiver)
	}

	arg := GetTransfersByReceiverIdParams{
		ToAccountID: receiver.ID,
		Limit:       5,
		Offset:      5,
	}

	tfs, err := testQueries.GetTransfersByReceiverId(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tfs, 5)

	for _, tf := range tfs {
		require.NotEmpty(t, tf)
		require.Equal(t, receiver.ID, tf.ToAccountID)
	}
}

func TestUpdateTransfer(t *testing.T) {
	tf1 := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     tf1.ID,
		Amount: util.RandomMoney(),
	}

	tf2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tf2)

	require.Equal(t, tf1.ID, tf2.ID)
	require.Equal(t, arg.Amount, tf2.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	tf1 := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), tf1.ID)
	require.NoError(t, err)

	tf2, err := testQueries.GetTransfer(context.Background(), tf1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, tf2)
}
