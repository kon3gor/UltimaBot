package save

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/db"
	"dev/kon3gor/ultima/internal/guard"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

const Cmd = "save"

func ProcessCommand(ctx *appcontext.Context) {
	if err := ctx.Guard(guard.DefaultUserNameGuard); err != nil {
		panic(err)
	}
	guarded(ctx)
}

func guarded(ctx *appcontext.Context) {
	connection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	text := ctx.RawUpdate.Message.Text
	_, solid, _ := strings.Cut(text, " ")
	dailies := strings.Split(solid, "\n")
	for _, daily := range dailies {
		if err = saveDaily(connection, daily); err != nil {
			connection.Release()
			panic(err)
		}
	}
	connection.Release()
}

const saveDailyQuery = `
DECLARE $daily AS String;
DECLARE $id as UInt64;
DECLARE $date as Date;
INSERT INTO SunshineDaily (id, content, date)
VALUES ($id, $daily, $date)
`

func saveDaily(connection *db.YdbConnection, daily string) error {
	id, err := generateId()
	if err != nil {
		return err
	}
	content := table.ValueParam("$daily", types.BytesValue([]byte(daily)))
	date, err := getCurrentDate()
	if err != nil {
		return err
	}
	params := table.NewQueryParameters(id, content, date)
	return connection.Execute(saveDailyQuery, params, nil)
}

func generateId() (table.ParameterOption, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return table.ValueParam("$id", types.Uint64Value(uint64(uid.ID()))), nil
}

func getCurrentDate() (table.ParameterOption, error) {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, err
	}

	date := time.Now().In(tz).UnixMilli() / 86400000
	return table.ValueParam("$date", types.DateValue(uint32(date))), nil
}
