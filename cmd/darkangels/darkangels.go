package darkangels

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#62e411")).Bold(true)
)

type Selection struct {
	Choice string
}

func (s *Selection) Update(value string) {
	s.Choice = value
}

type SummaryModel struct {
	summary string
}

func (m SummaryModel) Init() tea.Cmd {
	return nil
}

func (m SummaryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m SummaryModel) View() string {
	return fmt.Sprintf("Faction Summary:\n\n%s", m.summary)
}

func FactionSummary() string {
	summary := `
	The Dark Angels are considered amongst the most powerful and 
	secretive of the Loyalist Space Marine Chapters. They were the 1st 
	Legion of the original 20 Space Marine Legions to be created during 
	the First Founding of the 30th Millennium.

	Though they claim complete allegiance and service to the Emperor of 
	Mankind, their actions and secret goals at times seem at odds with 
	that professed loyalty, as the Dark Angels strive above all other 
	things to atone for an ancient crime of betrayal committed over 
	10,000 standard years ago against the trust of the Emperor during 
	the time of the Horus Heresy.
	
	The Dark Angels stand first amongst the Space Marine Chapters, as 
	they have done since their very inception as the Ist Legion. They 
	are a proud Chapter, with traditions and rituals that date back to 
	the earliest days of the Imperium of Man. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Good mix of resilient hard to shift deathwing. Fast movers.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Dark Angels are at tier 4, and worst, with a current win-rate of 37%.

	Price: The Dark Angels are a reasonable cheap army to collect.
	`
	return searchResult
}