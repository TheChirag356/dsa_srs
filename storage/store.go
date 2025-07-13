package storage

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