package save

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/util"
	"fmt"
	"net/http"
	"strings"
)

func saveMyDaily(ctx *appcontext.Context, dailies []string, shift int) {
	daily, err := getExistingFutureDaily(shift)
	if err != nil {
		// panic(err)
	}
	res := strings.Join(append(dailies, daily), "\n")
	filePathTemplate := "plans/daily/%s.md"
	currentDate, err := util.GetDateAsString(shift)
	if err != nil {
		panic(err)
	}
	filePath := fmt.Sprintf(filePathTemplate, currentDate)
	req := ghclient.NewMyPushRequest("PersonalObsidian", "main", filePath, res)
	ghclient.PushContent(http.DefaultClient, req)
}

// todo: I can put it in the ghclient i guess
func getExistingFutureDaily(shift int) (string, error) {
	filePathTemplate := "plans/daily/%s.md"
	currentDate, err := util.GetDateAsString(shift)
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf(filePathTemplate, currentDate)
	client := &http.Client{}
	req := ghclient.NewMyContentRequest("PersonalObsidian", filePath)
	return ghclient.GetFile(client, req)
}

func formatMyDailies(dailies []string) {

}
