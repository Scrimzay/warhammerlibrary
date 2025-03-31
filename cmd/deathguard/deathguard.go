package deathguard

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
	The Death Guard are a Traitor Legion of Chaos Space Marines in the 
	Warhammer 40,000 universe, dedicated to the Plague God Nurgle. Once 
	known as the Dusk Raiders, they were one of the original 20 Space 
	Marine Legions created by the Emperor to conquer the galaxy during 
	the Great Crusade. Led by their Primarch, Mortarion, the Death Guard 
	became known for their resilience and endurance in the face of 
	adversity, often taking on the most grueling and hazardous campaigns 
	in the Emperor's name.

	However, during the Horus Heresy, Mortarion and his Legion were 
	betrayed and left to die by the Warmaster Horus. In desperation, 
	Mortarion turned to Nurgle, who saved the Death Guard from certain 
	death but forever twisted them into plague-ridden, rotting monsters. 
	Now, the Death Guard wage an eternal war against the Imperium and 
	all life in the galaxy, spreading disease, decay, and despair 
	wherever they go. With their corrupted bodies and putrid weapons, 
	the Death Guard are a nightmarish force of contagion and entropy, 
	their once proud loyalty to the Emperor replaced by an unwavering 
	devotion to Nurgle and his cycle of decay and rebirth.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Very slow and durable mid to close range army. Spreads a plague aura
	across the battlefield every phase, weakening the opponents toughness and
	corrupting objectives.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Death Guard are at tier 3, with a current win-rate of 47%.

	Price: The Death Guard are a middle tier cost army.
	`
	return searchResult
}