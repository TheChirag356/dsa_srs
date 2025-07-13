package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/TheChirag356/dsa_srs/constants"
	"github.com/TheChirag356/dsa_srs/storage"
	tea "github.com/charmbracelet/bubbletea"
)

type reviewModel struct {
	store      *storage.CardStore
	due        []storage.DueCard
	current    int
	showAnswer bool
	done       bool
	errMsg     string
}

func newReviewModel() reviewModel {
	store, _ := storage.LoadCards(constants.CardFilePath)
	due := store.GetDueCards()
	return reviewModel{
		store:   store,
		due:     due,
		current: 0,
		done:    len(due) == 0,
		errMsg:  "",
	}
}

func (m *reviewModel) applyFeedback(quality string) {
	card := &m.due[m.current]
	now := time.Now()

	var ef float64 = 2.5
	var rep int
	var nextReview time.Time
	var c storage.ConceptCard
	var p storage.ProblemCard
	var interval int

	switch card.Type {
	case storage.ConceptCardType:
		c = card.Card.(storage.ConceptCard)
		rep = c.Repetition
		ef = c.EaseFactor
	case storage.ProblemCardType:
		p = card.Card.(storage.ProblemCard)
		rep = p.Repetition
		ef = p.EaseFactor
	}

	q := map[string]int{"1": 1, "2": 3, "3": 5}[quality]

	// SM-2
	if q < 3 {
		rep = 0
		interval = 1
		nextReview = now.Add(24 * time.Hour)
	} else {
		switch rep {
		case 0:
			interval = 1
			nextReview = now.Add(24 * time.Hour)
		case 1:
			interval = 6
			nextReview = now.Add(6 * 24 * time.Hour)
		default:
			intervalF := float64(1)
			switch card.Type {
			case storage.ConceptCardType:
				intervalF = float64(c.Interval)
			case storage.ProblemCardType:
				intervalF = float64(p.Interval)
			}
			intervalF = intervalF * ef
			interval = int(intervalF)
			nextReview = now.Add(time.Duration(intervalF) * 24 * time.Hour)
		}
		rep++
		ef = max(1.3, ef-0.8+0.28*float64(q)-0.02*float64(q)*float64(q))
	}

	// Save updates to card
	switch card.Type {
	case storage.ConceptCardType:
		for i := range m.store.ConceptCards {
			if m.store.ConceptCards[i].ID == card.ID {
				m.store.ConceptCards[i].EaseFactor = ef
				m.store.ConceptCards[i].Repetition = rep
				m.store.ConceptCards[i].NextReview = nextReview
				m.store.ConceptCards[i].Interval = interval
				m.store.ConceptCards[i].LastReviewed = now
			}
		}
	case storage.ProblemCardType:
		for i := range m.store.ProblemCards {
			if m.store.ProblemCards[i].ID == card.ID {
				m.store.ProblemCards[i].EaseFactor = ef
				m.store.ProblemCards[i].Repetition = rep
				m.store.ProblemCards[i].NextReview = nextReview
				m.store.ProblemCards[i].Interval = interval
				m.store.ProblemCards[i].LastReviewed = now
			}
		}
	}
}

func (m reviewModel) Init() tea.Cmd {
	return nil
}

func (m reviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.done {
			return initialModel(), waitAndReturn()
		}

		key := msg.String()

		if !m.showAnswer {
			m.showAnswer = true
			return m, nil
		}

		switch key {
		case "1", "2", "3":
			m.applyFeedback(key)
			m.current++
			m.showAnswer = false

			err := storage.SaveCards(constants.CardFilePath, m.store)
			if err != nil {
				m.errMsg = fmt.Sprintf("Error: %v", err)
				return m, waitAndReturn()
			}

			if m.current >= len(m.due) {
				m.done = true
			}

			return m, nil
		case "q":
			return initialModel(), nil
		}
	}
	return m, nil
}

func (m reviewModel) View() string {
	if m.done {
		if len(m.due) == 0 {
			return successStyle.Render("🎉 No cards due today!\n\nPress any key to return.")
		}
		if m.errMsg != "" {
			return errorStyle.Render("Error: "+m.errMsg) + "\n\nReturning to main menu in 3 seconds."
		}
		return successStyle.Render("✅ All due cards reviewed!\n\nPress any key to return.")
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("Due Cards") + "\n\n")
	b.WriteString(fmt.Sprintf("Card %d of %d\n\n", m.current+1, len(m.due)))

	card := m.due[m.current]

	switch card.Type {
	case storage.ConceptCardType:
		c := card.Card.(storage.ConceptCard)
		b.WriteString(formPadding.Render(c.Title) + "\n")
		b.WriteString(formPadding.Render(c.Notes) + "\n\n")
	case storage.ProblemCardType:
		p := card.Card.(storage.ProblemCard)
		b.WriteString(formPadding.Render(p.Title) + "\n")
		b.WriteString(formPadding.Render(p.Topic) + "\n")
		b.WriteString(formPadding.Render(p.Link) + "\n\n")
	}

	b.WriteString("\n" + helpStyle.Render("[1] Again • [2] Good • [3] Easy • [Ctrl+C or q] Quit"))
	return b.String()
}
