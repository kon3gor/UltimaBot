package db

import (
	"context"
	"os"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result"
	yc "github.com/ydb-platform/ydb-go-yc"
)

type YdbConnection struct {
	connection ydb.Connection
	cancelFunc context.CancelFunc
	Context    context.Context
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
		Context:    ctx,
	}, nil
}

type ResultConsumer func(*YdbConnection, result.Result)

func (self *YdbConnection) Execute(query string, params *table.QueryParameters, consumer ResultConsumer) error {
	queryErr := self.connection.Table().Do(self.Context, func(ctx context.Context, s table.Session) (err error) {
		_, res, err := s.Execute(ctx, table.DefaultTxControl(), query, params)
		if consumer != nil {
			consumer(self, res)
		}
		if err != nil {
			return err
		}
		defer res.Close()
		return res.Err() // for driver retry if not nil
	})
	return queryErr
}

func (self *YdbConnection) Release() {
	self.cancelFunc()
	self.connection.Close(self.Context)
}
