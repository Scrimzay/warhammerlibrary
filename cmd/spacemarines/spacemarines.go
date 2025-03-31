package spacemarines

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
	The Space Marines or Adeptus Astartes are foremost amongst the 
	defenders of Humanity, the greatest of the Emperor of Mankind's 
	warriors. They are barely Human at all, but superhuman; having been 
	made superior in all respects to a normal man by a harsh regime of 
	genetic modification, psycho-conditioning and rigorous training.

	Space Marines are untouched by plague or any natural disease and can 
	suffer wounds that would kill a lesser being several times over, and 
	live to fight again. Clad in ancient power armour and wielding the 
	most potent weapons known to man, the Space Marines are terrifying 
	foes and their devotion to the Emperor and the Imperium of Man is 
	unyielding. They are the God-Emperor's Angels of Death, and they 
	know no fear.

	The Astartes are physically stronger, far more resilient and often 
	mentally far removed from the lot of most normal Human beings. In the 
	presence of the Astartes, most people feel a combination of awe and 
	fear, and many cultures on the more primitive worlds simply worship 
	them outright as demigods or angels of the God-Emperor made flesh. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Versatile and well rounded force in all aspect of the game. Straight forward.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Space Marines are at tier 4, with a current win-rate of 40%.

	Price: The Space Marines are a reasonable cheap army to collect.
	`
	return searchResult
}