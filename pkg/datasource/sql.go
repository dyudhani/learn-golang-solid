package datasource

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"learn-golang-solid/pkg/utils"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type (
	Conn interface {
		BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
		PingContext(ctx context.Context) (err error)
		io.Closer
		ConnTx
	}

	ConnTx interface {
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
		QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	}

	Exec interface {
		Scan(rowsAffected, lastInsertId *int64) (err error)
	}

	Query interface {
		Scan(row func(i int) utils.Array) (err error)
	}

	exec struct {
		sqlResult sql.Result
		err       error
	}

	query struct {
		sqlRows *sqlx.Rows
		err     error
	}

	DataSource struct{}
)

var (
	_   Conn   = (*sqlx.Conn)(nil)
	_   Conn   = (*sqlx.DB)(nil)
	_   ConnTx = (*sqlx.Tx)(nil)
	log        = zerolog.New(os.Stdout)
)

var (
	ErrNoClumnReturned    = errors.New("no columns returned")
	ErrDataNotFound       = errors.New("data not found")
	ErrInvalidArguments   = errors.New("invalid arguments for scan")
	ErrInvalidTransaction = errors.New("invalid transaction")
)

func (x exec) Scan(rowsAffected, lastInsertId *int64) error {
	if x.err != nil {
		log.Err(x.err)
	}
}
