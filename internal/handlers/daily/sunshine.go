package daily

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/db"
	"dev/kon3gor/ultima/internal/util"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

const sunshineDailyQuery = `
declare $date as Date;
select * from SunshineDaily
where date = $date
order by content;
`

var daily []string

func sunshineDaily(ctx *appcontext.Context) {
	daily = make([]string, 0)
	date, err := getCurrentDateAsParam()
	if err != nil {
		ctx.SmthWentWrong(err)
		return
	}
	connection, err := db.Connect()
	if err != nil {
		ctx.SmthWentWrong(err)
		return
	}

	connection.Execute(sunshineDailyQuery, table.NewQueryParameters(date), readResults)
	content := strings.Join(daily, "\n")
	if content == "" {
		content = "Nothing here"
	}
	content = util.EscapeFakeMarkdown(content)
	msg := tgbotapi.NewMessage(ctx.ChatID, content)
	msg.ParseMode = "MarkdownV2"
	ctx.CustomAnswer(msg)
}

func readResults(connection *db.YdbConnection, res result.Result) {
	if err := res.NextResultSetErr(connection.Context); err != nil {
		panic(err)
	}
	i := 1
	for res.NextRow() {
		var content string
		err := res.ScanNamed(named.OptionalWithDefault("content", &content))
		if err != nil {
			log.Println(err)
		}
		daily = append(daily, fmt.Sprintf("`%d.` %s", i, content))
		i++
	}
}

// todo: remove it from here, must be an utility
func getCurrentDateAsParam() (table.ParameterOption, error) {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, err
	}

	date := time.Now().In(tz).UnixMilli() / 86400000
	return table.ValueParam("$date", types.DateValue(uint32(date))), nil
}
