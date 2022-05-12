package main

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"tutorial.sqlc.dev/app/tutorial"
)

func createAuthor(t *testing.T) tutorial.Author {

	args := tutorial.CreateAuthorParams{
		Name: "Linda",
		Bio:  sql.NullString{String: "She was gorgeous!", Valid: true},
	}

	author, err := testQueries.CreateAuthor(context.Background(), args)

	require.NoError(t, err)

	require.NotEmpty(t, author)

	require.Equal(t, args.Name, author.Name)

	require.Equal(t, args.Bio, author.Bio)

	require.NotZero(t, author.ID)

	return author

}

func TestGetAuthor(t *testing.T) {

	err := testQueries.DeleteAllAuthors(context.Background())

	require.NoError(t, err)

	firstAuthor := createAuthor(t)
	author2, err := testQueries.GetAuthor(context.Background(), firstAuthor.ID)

	require.NoError(t, err)

	require.NotEmpty(t, author2)

	require.Equal(t, firstAuthor.ID, author2.ID)
	require.Equal(t, firstAuthor.Name, author2.Name)
	require.Equal(t, firstAuthor.Bio, author2.Bio)

}

func TestUpdateAuthor(t *testing.T) {

	err := testQueries.DeleteAllAuthors(context.Background())

	require.NoError(t, err)

	initialAuthor := createAuthor(t)

	args := tutorial.UpdateAuthorParams{
		ID:   initialAuthor.ID,
		Name: "lilian",
		Bio:  sql.NullString{String: "I was rejected", Valid: true},
	}

	updatedAuthor, error := testQueries.UpdateAuthor(context.Background(), args)

	require.NoError(t, error)

	require.NotEmpty(t, updatedAuthor)

	require.Equal(t, initialAuthor.ID, updatedAuthor.ID)
	require.Equal(t, args.Name, updatedAuthor.Name)
	require.Equal(t, args.Bio, updatedAuthor.Bio)

}

func TestDeleteAuthor(t *testing.T) {

	err := testQueries.DeleteAllAuthors(context.Background())

	require.NoError(t, err)

	initialAuthor := createAuthor(t)

	err = testQueries.DeleteAuthor(context.Background(), initialAuthor.ID)

	require.NoError(t, err)

	author, err := testQueries.GetAuthor(context.Background(), initialAuthor.ID)

	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, author)
}

func TestListAuthors(t *testing.T) {

	err := testQueries.DeleteAllAuthors(context.Background())

	require.NoError(t, err)

	createAuthor(t)

	allAuthors, err := testQueries.ListAuthors(context.Background())

	require.NoError(t, err)

	require.Len(t, allAuthors, 1)

	for _, authors := range allAuthors {
		require.NotEmpty(t, authors)
	}
}
