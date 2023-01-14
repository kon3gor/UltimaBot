package commands

import (
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const ideaCmd = "idea"

func idea(context *context.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
	ideaGuarded(context)
}

func ideaGuarded(context *context.Context) {
	idea := selectIdea()
	context.TextAnswer(idea)
}

func selectIdea() string {
	client := &http.Client{}
	req := ghclient.NewMyContentRequest("PersonalObsidian", "ideas")
	result, _ := ghclient.GetContent(client, req)

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(result))

	res, err := client.Get(result[index].DownloadUrl)
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
