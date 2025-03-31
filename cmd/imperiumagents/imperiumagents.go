package imperiumagents

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
	Throughout the Imperium there exist numerous martial organisations
	and shadowy institutions. Bodies of armed warriors or solitary agents
	from these groups possess specialist skills, unusual equipment and
	vested interests that lead them to be attached to larger Imperial
	armies. Some are requistioned by the army's commander for their
	particular abilities, others are assigned for their hidden masters
	to achieve singular agendas. The most powerful have the authority
	and reputation to enforce their presence on the field of battle.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Neither the Ordo Xenos, Ordo Hereticus nor Ordo Malleus have
	a defined playstyle. 

	Tier: N/A

	Price: N/A
	`
	return searchResult
}