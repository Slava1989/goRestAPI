package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Slava1989/goRestAPI/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID string
	Slug sql.NullString
	Body sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID: c.ID,
		Slug: c.Slug.String,
		Body: c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	_, err := d.Client.ExecContext(ctx, "SELECT pg_sleep(16)")
	if err != nil {
		return comment.Comment{}, err
	}
	row := d.Client.QueryRowContext(
		ctx,
		`Select id, slug, body, author
		From comments
		Where id = $1`,
		uuid,
	)
	err = row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching the comment by uuid: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID: 	cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body: 	sql.NullString{String: cmt.Body, Valid: true},

	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, slug, author, body)
		VALUES
		(:id, :slug, :author, :body)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to delete comment from database: %w", err)
	}

	return nil
}

func (d *Database) UpdateComment(ctx context.Context, 
	id string, 
	cmt comment.Comment) (comment.Comment, error) {
	
	commentRow := CommentRow {
		ID: id,
		Slug: sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body: sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		 slug = :slug,
		 author = :author,
		 body = :body
		 WHERE id = :id`,
		commentRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error updating the comment by uuid: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCommentRowToComment(commentRow), nil
}