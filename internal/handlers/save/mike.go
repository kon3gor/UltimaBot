package save

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/util"
	"fmt"
	"strings"
)

func saveMyDaily(ctx *appcontext.Context, dailies []string, shift int) {
	daily, _ := getExistingFutureDaily(shift)
	res := fmt.Sprintf("%s\n%s", formatMyDailies(dailies), daily)
	filePathTemplate := "plans/daily/%s.md"
	currentDate, err := util.GetDateAsString(shift)
	if err != nil {
		panic(err)
	}
	filePath := fmt.Sprintf(filePathTemplate, currentDate)
	req := ghclient.NewPersonalObsidianRequest(filePath, res)
	ghclient.PushContent(req)
	ctx.TextAnswer("Saved")
}

// todo: I can put it in the ghclient i guess
func getExistingFutureDaily(shift int) (string, error) {
	filePathTemplate := "plans/daily/%s.md"
	currentDate, err := util.GetDateAsString(shift)
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf(filePathTemplate, currentDate)
	req := ghclient.NewMyContentRequest("PersonalObsidian", filePath)
	return ghclient.GetFile(req)
}

func formatMyDailies(dailies []string) string {
	r := strings.Builder{}
	for _, v := range dailies[:len(dailies)-1] {
		r.WriteString("- [ ] ")
		r.WriteString(v)
		r.WriteRune('\n')
	}
	r.WriteString("- [ ] ")
	r.WriteString(dailies[len(dailies)-1])
	return r.String()
}
