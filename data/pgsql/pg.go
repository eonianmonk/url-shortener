package pgsql

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"

	"github.com/eonianmonk/url-shortener/data"
	"github.com/eonianmonk/url-shortener/types"
)

type pgStorage struct {
	ctx context.Context
	db  *sql.DB
	e   types.Encoder
}

type urlsEntry struct {
	urlId   int32
	fullUrl string
}

var (
	urlTableName         = "urls"
	urlTableSq           = sq.Select("*").From(urlTableName)
	urlIdColumnName      = "url_id"
	urlFullUrlColumnName = "full_url"
)

func NewPgStorage(user string, password string, dbname string, address string, encoder types.Encoder, ctx context.Context) data.Storage {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, address, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to establish database connection: %s", err.Error()))
	}
	return &pgStorage{db: db, e: encoder, ctx: ctx}
}

func (pg *pgStorage) Put(u types.Url) (types.ShortUrl, error) {
	failed := func(msg string, err error) (types.ShortUrl, error) {
		return "", fmt.Errorf("%s failed: %s", msg, err.Error())
	}

	tx, err := pg.db.BeginTx(pg.ctx, nil)
	if err != nil {
		return failed("tx initiation", err)
	}
	defer tx.Rollback()

	sqlCheckExistance, _, err := urlTableSq.Where(sq.Eq{urlFullUrlColumnName: string(u)}).ToSql()
	if err != nil {
		return failed("sql for existance check compilation", err)
	}
	r := tx.QueryRow(sqlCheckExistance)
	var e urlsEntry
	err = r.Scan(&e)
	if err == nil {
		return types.ShortUrl(e.urlId), nil
	}
	if err != sql.ErrNoRows {
		return failed("existance check sql execution", err)
	}

	sqlCreateEntry, _, err := sq.Insert(urlTableName).Columns(urlFullUrlColumnName).Values(string(u)).ToSql()
	if err != nil {
		return failed("sql create entry", err)
	}
	r = tx.QueryRow(sqlCreateEntry)
	err = r.Scan(&e)
	if err == nil {
		return types.ShortUrl(e.urlId), nil
	}
	if err != nil {
		return failed("create entry", err)
	}
	return "", nil
}

func (pg *pgStorage) Get(su types.ShortUrl) (types.Url, error) {
	id, err := pg.e.Decode(su)
	if err != nil {
		return "", fmt.Errorf("failed to decode short url: %s", err.Error())
	}
	// SELECT * from urls where id == id
	sql, _, err := urlTableSq.Where(sq.Eq{urlIdColumnName: int32(id)}).ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to compile sql: %s", err.Error())
	}
	rows, err := pg.db.Query(sql)
	if err != nil {
		return "", fmt.Errorf("failed to receive sql responce: %s", err.Error())
	}
	if !rows.Next() {
		return "", fmt.Errorf("no short url registered")
	}
	var e urlsEntry
	if err := rows.Scan(&e); err != nil {
		return "", fmt.Errorf("failed to read row: %s", err.Error())
	}

	return types.Url(e.fullUrl), nil
}
