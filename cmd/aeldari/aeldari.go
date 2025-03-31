package aeldari

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
	The Aeldari, also known as the Eldar, are an ancient and highly 
	advanced alien race in the Warhammer 40,000 universe. They once 
	ruled a vast empire but fell into decline due to their own 
	hedonistic excesses, which led to the birth of the Chaos God 
	Slaanesh. The Aeldari are now divided into several distinct factions:

	Craftworld Aeldari: They live on massive, planet-sized ships called 
	Craftworlds and follow a strict "Path" system to control their 
	emotions and avoid falling into depravity.

	Drukhari (Dark Eldar): These Aeldari revel in piracy, slavery, 
	and torture, living in the hidden city of Commorragh within the 
	Webway, a labyrinthine dimension that allows faster-than-light 
	travel.

	Harlequins: Enigmatic warrior-dancers who serve the Aeldari god of 
	death, Cegorach, and act as mediators between the various Aeldari 
	factions.

	Exodites: Aeldari who rejected the decadence of their society 
	before the fall and chose to live a simpler, agrarian lifestyle 
	on remote worlds.

	Ynnari: A relatively new faction that seeks to unite the various 
	Aeldari groups to fight against the forces of Chaos and ultimately 
	destroy Slaanesh.

	The Aeldari are known for their advanced technology, psychic 
	abilities, and swift, elegant fighting style. They often see 
	themselves as superior to other races and can be haughty and aloof. 
	Despite their diminished numbers, they remain a significant power 
	in the galaxy, fighting to preserve their race and atone for their 
	past sins.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Fast moving, hard-hitters, that must hide behind cover
	due to having pretty low health and fragile toughness.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Aeldari are at tier 3, with a current win-rate of 47%.

	Price: The Aeldari are quite an expensive faction to collect.
	`
	return searchResult
}