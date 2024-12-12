package sqlc

// パッケージ変数としてQueriesを保持
var (
	queries *Queries
)

func SetQueries(db DBTX) {
	queries = New(db)
}

func GetQueries() *Queries {
	return queries
}
