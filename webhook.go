package main

import (
	"fmt"
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := gitlab.New(gitlab.Options.Secret("FxHn6PJ6FG1STYiJpS2G"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, gitlab.MergeRequestEvents)
		if err != nil {
			if err == gitlab.ErrEventNotFound {
				fmt.Printf("event not found %s", err)
			}
		}
		switch payload.(type) {
		case gitlab.MergeRequestEventPayload:
			mergeRequest := payload.(gitlab.MergeRequestEventPayload)
			lastCommitID := mergeRequest.ObjectAttributes.LastCommit.ID

			fmt.Printf("%+v", lastCommitID)
		}
	})
	http.ListenAndServe(":3000", nil)
}
