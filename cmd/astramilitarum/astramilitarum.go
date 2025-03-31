package astramilitarum

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
	The Astra Militarum, also known as the Imperial Guard in colloquial 
	Low Gothic, is the largest coherent fighting force in the galaxy and 
	serves as the Imperium of Man's primary military force and first 
	line of defence from the myriad threats which endanger the existence 
	of the Human race in the 41st Millennium.

	It is comprised of countless billions of men and women -- hundreds 
	of thousands of different regiments, supported by a vast array of 
	light and heavy armoured vehicles that provide the Imperial Guard's 
	primary offensive punch. The Astra Militarum is usually the first 
	Imperial force to respond to a threat if a world's Planetary Defence 
	Force (PDF) fails to suppress it.

	They also garrison major locations of strategic or cultural interest 
	to the Imperium and are often found in defensive roles. Supported 
	by legions of heavy armour and thundering artillery, the Imperial 
	Guard fight a never-ending war for the survival of Mankind in an 
	unrelentingly hostile universe. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Large but very fragile and tank/vehicle heavy army. Heavily reliant
	on shooting and commissars issuing orders to add flexibility.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Astra Militarum are at tier 4, with a current win-rate of 44%.

	Price: The Astra Militarum are a very expensive army.
	`
	return searchResult
}