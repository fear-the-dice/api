package models

type (
	DbOptions struct {
		Host     string
		Database string
	}
)

var (
	dbOptions *DbOptions
)

func SetConfig(dbo *DbOptions) {
	dbOptions = dbo
}
