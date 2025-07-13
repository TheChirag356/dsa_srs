package storage

import (
	"time"

	"github.com/google/uuid"
)

// ConceptCard extends BaseCard
type ConceptCard struct {
	BaseCard
	Title string `json:"title"`
	Notes string `json:"notes"`
}

func NewConceptCard(title string, notes string) (ConceptCard, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return ConceptCard{}, err
	}

	return ConceptCard{
		BaseCard: BaseCard{
			ID:           id.String(),
			Type:         ConceptCardType,
			LastReviewed: time.Now(),
			NextReview:   time.Now().Add(time.Hour * 24),
			Repetition:   0,
			EaseFactor:   2.5,
			Interval: 1,
		},
		Title: title,
		Notes: notes,
	}, nil
}
