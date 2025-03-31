package drukhari

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
	The Drukhari, or Dark Eldar, are a sadistic and cruel Aeldari 
	faction that split off from their kin after the fall of the Eldar 
	empire. They reside in the dark city of Commorragh, hidden within 
	the Webway, where they indulge in piracy, slave raids, and torture 
	to sustain their souls and stave off the predations of Slaanesh. 
	The Drukhari are organized into Kabals, Wych Cults, and Haemonculus 
	Covens, each with its own specialties and hierarchy. They are known 
	for their lightning-fast raids, terrifying weaponry, and utter lack 
	of mercy. Despite their cruelty, the Drukhari are a key part of the 
	Aeldari race, and their skills in the arts of war and subterfuge 
	make them a force to be reckoned with in the 41st millennium.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Very fast movers, very fast melee hitters, but not many unique units
	and not too tough or tanky.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Drukhari are at tier 3, with a current win-rate of 50%.

	Price: The Drukhari are a okayishly expensive faction to collect.
	`
	return searchResult
}