package daos

type PostgresDao interface {
	Open() error
	Close() error
}
