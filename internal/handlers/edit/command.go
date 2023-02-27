package edit

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/db"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

const Cmd = "edit"

func ProcessCommand(ctx *appcontext.Context) {
	if err := ctx.Guard(guard.DefaultUserNameGuard); err != nil {
		return
	}
	guarded(ctx)
}

func guarded(ctx *appcontext.Context) {
	ind, newc, _ := strings.Cut(ctx.RawUpdate.Message.CommandArguments(), " ")
	index, err := strconv.Atoi(ind)
	if err != nil {
		ctx.SmthWentWrong(err)
		return
	}

	updateSunshineDaily(ctx, index, newc)
}

const query = `
declare $ind as Int;
declare $content as String;
declare $date as Date;
$ordered = select row_number() over (order by content) as ind, content, id
from SunshineDaily where date = $date;
$id = select id from $ordered where ind = $ind;
`

var deleteQuery = fmt.Sprintf(`
%s
delete from SunshineDaily where id = $id;
`, query)

var updateQuery = fmt.Sprintf(`
%s
update SunshineDaily set content = $content where id = $id;
`, query)

func updateSunshineDaily(ctx *appcontext.Context, index int, content string) {
	connection, err := db.Connect()
	if err != nil {
		ctx.SmthWentWrong(err)
		return
	}
	defer connection.Release()

	contentParam := table.ValueParam("$content", types.BytesValueFromString(content))
	indParam := table.ValueParam("$ind", types.Int32Value(int32(index)))
	dateParam, err := getCurrentDateAsParam()
	if err != nil {
		ctx.SmthWentWrong(err)
		return
	}
	var q string
	if content == "" {
		q = deleteQuery
	} else {
		q = updateQuery
	}
	err = connection.Execute(q, table.NewQueryParameters(contentParam, indParam, dateParam), nil)
	if err != nil {
		ctx.SmthWentWrong(err)
		return
	}
}

func getCurrentDateAsParam() (table.ParameterOption, error) {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, err
	}

	date := time.Now().In(tz).UnixMilli() / 86400000
	return table.ValueParam("$date", types.DateValue(uint32(date))), nil
}
