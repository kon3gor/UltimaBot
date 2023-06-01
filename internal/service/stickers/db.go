package stickers

import (
	"dev/kon3gor/ultima/internal/service/db"
	"log"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

const getStickersQuery = `
select file from Stickers;
`

const saveStickerQuery = `
declare $sticker as String;
declare $id as String;
insert into Stickers (id, file)
values ($id, $sticker);
`

func SaveSticker(uniqueId, sticker string) error {
	connection, err := db.Connect()
	if err != nil {
		connection.Release()
		return err
	}
	
	id := table.ValueParam("$id", types.BytesValue([]byte(uniqueId)))
	param := table.ValueParam("$sticker", types.BytesValue([]byte(sticker)))
	connection.Execute(saveStickerQuery, table.NewQueryParameters(id, param), nil)
	connection.Release()

	return nil
}

func GetStickers() ([]string, error) {
	stickers = make([]string, 0)
	connection, err := db.Connect()
	if err != nil {
		return nil, err
	}

	connection.Execute(getStickersQuery, nil, readResults)

	return stickers, nil
}

var stickers []string

func readResults(connection *db.YdbConnection, res result.Result) {
	if err := res.NextResultSetErr(connection.Context); err != nil {
		panic(err)
	}
	for res.NextRow() {
		var fileID string
		err := res.ScanNamed(named.OptionalWithDefault("file", &fileID))
		if err != nil {
			log.Println(err)
		}
		stickers = append(stickers, fileID)
	}
}
