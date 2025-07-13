package storage

import (
	"time"
)

type CardType string

const (
	ProblemCardType CardType = "problem"
	ConceptCardType CardType = "concept"
)

type BaseCard struct {
	ID           string    `json:"id"`
	Type         CardType  `json:"type"`
	LastReviewed time.Time `json:"last_reviewed"`
	NextReview   time.Time `json:"next_review"`
	Repetition   int       `json:"repetition"`
	EaseFactor   float64   `json:"ease_factor"`
}