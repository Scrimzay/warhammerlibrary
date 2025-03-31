package chaosmarines

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
	The Chaos Space Marines are the twisted and corrupted members of the 
	Adeptus Astartes who have turned their backs on the Emperor of 
	Mankind and pledged their allegiance to the dark gods of Chaos. Once 
	loyal and noble warriors, these former Space Marines have been 
	seduced by the promises of power, immortality, and the freedom to 
	indulge their darkest desires. Led by the Traitor Primarchs and 
	other powerful Chaos Lords, they wage an unending war against the 
	Imperium of Man and seek to plunge the galaxy into eternal darkness.

	Each Chaos Space Marine Legion and warband is unique, having 
	embraced different aspects of the Ruinous Powers. Some, like the 
	World Eaters, revel in bloodshed and mindless violence, while others, 
	like the Thousand Sons, delve deep into the sorcerous arts and the 
	pursuit of forbidden knowledge. Despite their differences, all 
	Chaos Space Marines are united in their hatred of the Imperium and 
	their desire to see the galaxy burn in the fires of Chaos. With 
	their ancient power armor, daemonic weapons, and unholy blessings, 
	they are a terrifying force on the battlefield, capable of 
	unleashing untold destruction and mayhem upon their foes.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Aggressive force with good durability at the cost of sacrifice
	to shooting but to hit very hard in combat. Dark pacts increase power to shooting
	and melee at the possible cost of health.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Chaos Marines are at tier 3, with a current win-rate of 50%.

	Price: The Chaos Marines are a middle tier cost army.
	`
	return searchResult
}