package tyranids

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
	The Tyranids are a terrifying alien race that exists outside the 
	galaxy, consisting of countless genetically engineered creatures 
	that function as a single, ravenous hive mind. They invade galaxies 
	in massive swarms called "Hive Fleets," consuming all organic matter 
	on planets to fuel their endless cycle of evolution and reproduction.
	Tyranid organisms range from small, swarming Genestealers to 
	gigantic, monstrous Biotitans, each adapted for a specific role in 
	the swarm. As they consume biomass, they incorporate the genetic 
	material of their prey into their own genome, constantly evolving 
	and adapting to new threats. The ultimate goal of the Tyranids is 
	to consume all life in the universe, making them an existential 
	threat to every other species in the Warhammer 40,000 galaxy. Their 
	hive mind is controlled by the powerful psychic will of the Hive 
	Queens, and their relentless advance has already consumed countless 
	worlds.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Short to mid range army with both shooting and great melee options.
	Hard to battle shock and okayishly tough.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Tyranids are at tier 4, with a current win-rate of 43%.

	Price: The Tyranids are a okayishly expensive faction to collect.
	`
	return searchResult
}