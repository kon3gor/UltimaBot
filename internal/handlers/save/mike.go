package save

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/util"
	"fmt"
	"net/http"
)

func saveMyDaily(ctx *appcontext.Context, dailies []string) {
}

func getExistingFutureDaily(shift int) (string, error) {
	date, err := util.GetDateAsString(shift)
	if err != nil {
		return "", err
	}
	p := fmt.Sprintf("plans/daily/%s.md", date)
	req := ghclient.NewMyContentRequest("PersonalObsidian", p)
	if _, err := ghclient.GetContent(http.DefaultClient, req); err != nil {
		return "", err
	}

	return "", nil
}

func formatMyDailies(dailies []string) {

}
