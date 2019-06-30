package slack

import (
	"fmt"

	"github.com/nlopes/slack"
)

type BlockData struct {
	Emoji  string
	Host   string
	Id     string
	Status string
}

func MsgBlock(b *BlockData) *slack.SectionBlock {
	txt := slack.NewTextBlockObject(
		"mrkdwn",
		fmt.Sprintf(
			">%s    Workflow <%s/%s|%s> %s",
			b.Emoji, b.Host, b.Id, b.Id, b.Status,
		),
		false, false,
	)
	return slack.NewSectionBlock(
		txt, nil, nil, slack.SectionBlockOptionBlockID("argo-block"),
	)
}
