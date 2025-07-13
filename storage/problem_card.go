package storage

import (
	"github.com/google/uuid"
	"time"
)

// ProblemCard extends BaseCard
type ProblemCard struct {
	BaseCard
	Title string `json:"title"`
	Link  string `json:"link"`  // leetcode url
	Topic string `json:"topic"` // eg. "binary search"
}

func NewProblemCard(title string, link string, topic string) (ProblemCard, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return ProblemCard{}, err
	}

	return ProblemCard{
		BaseCard: BaseCard{
			ID:           id.String(),
			Type:         ProblemCardType,
			LastReviewed: time.Now(),
			NextReview:   time.Now().Add(time.Hour * 24),
			Repetition:   0,
			EaseFactor:   2.5,
		},
		Title: title,
		Link:  link,
		Topic: topic,
	}, nil
}
