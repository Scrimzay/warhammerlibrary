package emperorschildren

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fe0000")).Bold(true)
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fe0000")).Bold(true)
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
	The Emperor's Children, also sometimes known after their fall as the 
	Lords of Profligacy, are a Traitor Legion of Chaos Space Marines who 
	devote themselves solely to the service of the Chaos God Slaanesh, 
	the Prince of Pleasure, though they were originally the Imperium of 
	Man's proud IIIrd Legion of Astartes.
	
	The Emperor's Children was the only Space Marine Legion to bear the 
	Emperor's own name and His own icon -- the Palatine Aquila -- 
	granted to them by His hand as a symbol of the Legion's martial 
	perfection. Few were ever so honoured amongst the ancient Space 
	Marine Legions and given less cause to betray the Master of Mankind 
	than the Emperor's Children.
	
	Given the plaudits and accolades accorded them, few could doubt that 
	they were the embodiment of what the Emperor had intended the 
	Legiones Astartes to be: noble in action and aspect, excelling in all 
	matters, strong, civilised, firm of purpose and loyal to the core. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Not currently known.

	Tier: Not currently known.

	Price: Not currently known.
	`
	return searchResult
}