package storage

import "time"

type CardStore struct {
	ConceptCards []ConceptCard `json:"concept_cards"`
	ProblemCards []ProblemCard `json:"problem_cards"`
}

func (store *CardStore) AppendConceptCard(card ConceptCard) {
	store.ConceptCards = append(store.ConceptCards, card)
}

func (store *CardStore) AppendProblemCard(card ProblemCard) {
	store.ProblemCards = append(store.ProblemCards, card)
}

type DueCard struct {
	Type CardType    `json:"card_type"`
	ID   string      `json:"id"`
	Card interface{} `json:"card"` // Can be ConceptCard or ProblemCard
}

// Load all cards and return only whose NextReview <= now.
func (s *CardStore) GetDueCards() []DueCard {
	now := time.Now()
	var due []DueCard

	for _, c := range s.ConceptCards {
		if c.NextReview.Before(now) || c.NextReview.Equal(now) {
			due = append(due, DueCard{
				Type: ConceptCardType,
				ID:   c.ID,
				Card: c,
			})
		}
	}

	for _, p := range s.ProblemCards {
		if p.NextReview.Before(now) || p.NextReview.Equal(now) {
			due = append(due, DueCard{
				Type: ProblemCardType,
				ID:   p.ID,
				Card: p,
			})
		}
	}

	return due
}
