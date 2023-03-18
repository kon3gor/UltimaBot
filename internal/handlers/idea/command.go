package idea

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const Cmd = "idea"

func ProcessCommand(context *appcontext.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
	guarded(context)
}

func guarded(context *appcontext.Context) {
	idea := selectIdea()
	context.TextAnswer(idea)
}

// todo: need to do a little drill down here to get all ideas
func selectIdea() string {
	req := ghclient.NewMyContentRequest("PersonalObsidian", "ideas")
	result, _ := ghclient.GetContent(req)

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(result))

	res, err := http.Get(result[index].DownloadUrl)
	if err != nil {
		return fmt.Sprintf("Error while fetching donwload: %s", err)
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintf("error parsing bytes: %s", err)
	}

	return strings.ReplaceAll(string(bodyBytes), "\t", "    ")
}
