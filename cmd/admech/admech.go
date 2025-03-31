package admech

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
	The Adeptus Mechanicus is the official Imperial name within the 
	Adeptus Terra for the Cult Mechanicus or Cult of the Machine based 
	on Mars which provides the Imperium with its scientists, engineers 
	and technicians.

	The tech-adepts of the Mechanicus are the primary keepers of what is 
	viewed as sacred wisdom, a privileged caste of Tech-priests who 
	jealously guard the knowledge required to maintain and construct 
	much of the Imperium's advanced technology. 

	The Mechanicus acknowledge the Emperor of Mankind as the ruler of 
	the Imperium of Man, but not the religious truth of the Imperial 
	Cult or the Ecclesiarchy, and are granted an unusual amount of 
	political and religious autonomy within the Imperium's structure, a 
	role protected by the ancient Treaty of Mars. The Mechanicus 
	essentially maintains an empire-within-an-empire in Imperial space, 
	and its interests usually, but not always, coincide with those of 
	the wider Imperium.

	Instead of the Imperial Creed, the servants of the Mechanicus follow 
	their own dark and mysterious scriptures and worship the strange 
	deity they call the Machine God or the "Omnissiah." 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Large but fragile force that excel at shooting. Typically just trying
	to shoot the enemy off the table. Great flexibility to apply an offensive or
	defensive armywide buff depending on what the turn calls for.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Adeptus Mechanicus are at tier 3, with a current win-rate of 49%.

	Price: The Adeptus Mechanicus are the most expensive army to collect.
	`
	return searchResult
}