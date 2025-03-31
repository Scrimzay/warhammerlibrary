package tauempire

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
	The T'au Empire is a young, technologically advanced faction in the 
	Warhammer 40,000 universe. Founded by the T'au, a blue-skinned, 
	humanoid race, the Empire is driven by the philosophy of the 
	Greater Good, which emphasizes unity, cooperation, and the 
	collective well-being of all species within the Empire.

	The T'au are known for their advanced pulse weaponry, sleek
	battlesuits, and highly disciplined Fire Warriors. They employ a 
	mix of advanced technology and sophisticated battlefield tactics, 
	relying on mobility, ranged firepower, and precise strikes to 
	overcome their foes.

	The T'au Empire is a rapidly expanding power, seeking to absorb 
	other races into their fold through diplomacy, coercion, or 
	military conquest. The Empire is led by the Ethereals, enigmatic 
	leaders who guide the T'au's society and ensure adherence to the 
	Greater Good.

	On the tabletop, T'au armies are known for their devastating ranged 
	firepower, with units like Fire Warriors, Pathfinders, and Crisis 
	Battlesuits unleashing precise volleys of pulse fire. They are 
	supported by a range of advanced drones, vehicles, and towering 
	Battlesuits like the Riptide and Stormsurge.

	While the T'au Empire claims to seek a unified and peaceful galaxy, 
	their relentless expansion and unyielding belief in the Greater Good 
	often brings them into conflict with the other major factions of 
	the Warhammer 40,000 universe.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Possibly the strongest shooting army in the game but the weakest melee.
	Use guide drones to imporve their shooting and to help protect your units.
	A fast moving army as well.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, T'au Empire are at tier 3, with a current win-rate of 49%.

	Price: The T'au Empire are a mid to high cost army.
	`
	return searchResult
}