package commands

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"encoding/json"
	"fmt"
	"net/http"

	"strings"
)

const scheduleCmd = "schedule"

func schedule(context *context.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
	scheduleGuarded(context)
}

const schedule_api_url = "https://bba7l5a13h11phof7aqh.containers.yandexcloud.net/schedule"

type ScheduleResponse struct {
	Schedule []string `json:"schedule"`
}

func scheduleGuarded(context *context.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", schedule_api_url, nil)
	res, err := client.Do(req)
	if err != nil {
		context.TextAnswer(fmt.Sprint(err))
		return
	}

	var result ScheduleResponse
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		context.TextAnswer(fmt.Sprint(err))
		return
	}

	schedule := strings.Join(result.Schedule, "\n")
	context.TextAnswer(schedule)
}
