package db

import (
	"context"
	"os"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	yc "github.com/ydb-platform/ydb-go-yc"
)

type YdbConnection struct {
	connection ydb.Connection
	cancelFunc context.CancelFunc
	context    context.Context
}

type YdbQuery func(context.Context, table.Session) error

func Connect() (*YdbConnection, error) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := ydb.Open(ctx,
		"grpcs://ydb.serverless.yandexcloud.net:2135/ru-central1/b1g1mak5k5l52mp3rc74/etn3jvn0er8t77urapjm",
		yc.WithInternalCA(),
		yc.WithServiceAccountKeyFileCredentials(os.Getenv("SA_PATH")),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	return &YdbConnection{
		connection: db,
		cancelFunc: cancel,
		context:    ctx,
	}, nil
}

func (self *YdbConnection) Execute(query table.Operation) error {
	queryErr := self.connection.Table().Do(self.context, query)
	if queryErr != nil {
		return queryErr
	}
	return nil
}

func (self *YdbConnection) Release() {
	self.cancelFunc()
	self.connection.Close(self.context)
}
