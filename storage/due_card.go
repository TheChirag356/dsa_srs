package storage

type DueCard struct {
	Type CardType    `json:"card_type"`
	ID   string      `json:"id"`
	Card interface{} `json:"card"` // Can be ConceptCard or ProblemCard
}