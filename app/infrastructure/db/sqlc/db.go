// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.fetchAllUserStmt, err = db.PrepareContext(ctx, fetchAllUser); err != nil {
		return nil, fmt.Errorf("error preparing query FetchAllUser: %w", err)
	}
	if q.findUserByEmailStmt, err = db.PrepareContext(ctx, findUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByEmail: %w", err)
	}
	if q.insertUserStmt, err = db.PrepareContext(ctx, insertUser); err != nil {
		return nil, fmt.Errorf("error preparing query InsertUser: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.fetchAllUserStmt != nil {
		if cerr := q.fetchAllUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fetchAllUserStmt: %w", cerr)
		}
	}
	if q.findUserByEmailStmt != nil {
		if cerr := q.findUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByEmailStmt: %w", cerr)
		}
	}
	if q.insertUserStmt != nil {
		if cerr := q.insertUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertUserStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                  DBTX
	tx                  *sql.Tx
	deleteUserStmt      *sql.Stmt
	fetchAllUserStmt    *sql.Stmt
	findUserByEmailStmt *sql.Stmt
	insertUserStmt      *sql.Stmt
	updateUserStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                  tx,
		tx:                  tx,
		deleteUserStmt:      q.deleteUserStmt,
		fetchAllUserStmt:    q.fetchAllUserStmt,
		findUserByEmailStmt: q.findUserByEmailStmt,
		insertUserStmt:      q.insertUserStmt,
		updateUserStmt:      q.updateUserStmt,
	}
}
