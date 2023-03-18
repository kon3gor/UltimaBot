package save

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/db"
	"dev/kon3gor/ultima/internal/guard"
	"dev/kon3gor/ultima/internal/util"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

const Cmd = "save"

func ProcessCommand(ctx *appcontext.Context) {
	if err := ctx.Guard(guard.DefaultUserNameGuard); err != nil {
		return
	}
	guarded(ctx)
}

func guarded(ctx *appcontext.Context) {
	text := ctx.RawUpdate.Message.Text
	_, solid, _ := strings.Cut(text, " ")
	dailies := strings.Split(solid, "\n")
	shift, err := strconv.Atoi(dailies[0])
	if err != nil {
		panic(err)
	}

	if ctx.UserName == "zosuku" {
		sunhineSave(ctx, dailies, shift)
	} else if ctx.UserName == "eshendo" {
		saveMyDaily(ctx, dailies[1:], shift)
	}
}

func sunhineSave(ctx *appcontext.Context, dailies []string, shift int) {
	connection, err := db.Connect()
	if err != nil {
		ctx.SmthWentWrong(err)
		return
	}
	defer connection.Release()

	for _, daily := range dailies {
		if err = saveDaily(connection, daily, shift); err != nil {
			connection.Release()
			ctx.SmthWentWrong(err)
		}
	}
}

const saveDailyQuery = `
DECLARE $daily AS String;
DECLARE $id as UInt64;
DECLARE $date as Date;
INSERT INTO SunshineDaily (id, content, date)
VALUES ($id, $daily, $date)
`

func saveDaily(connection *db.YdbConnection, daily string, shift int) error {
	id, err := generateId()
	if err != nil {
		return err
	}
	content := table.ValueParam("$daily", types.BytesValue([]byte(daily)))
	date, err := getCurrentDate(shift)
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

func getCurrentDate(shift int) (table.ParameterOption, error) {
	date, err := util.GetCurrentDateAsMillis(shift)
	if err != nil {
		return nil, err
	}
	return table.ValueParam("$date", types.DateValue(uint32(date))), nil
}
