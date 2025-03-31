package custodes

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
	The Adeptus Custodes, known as the Legio Custodes during the era of 
	the Great Crusade and Horus Heresy, and colloquially as the "Ten 
	Thousand," is the Imperial adepta responsible for protecting the 
	Imperial Palace and the physical body of the Emperor of Mankind, as 
	well as serving as His most important emissaries, His companions, 
	and the keepers of His many secrets. The Custodes is an elite cadre 
	of genetically-engineered transhuman male and female warriors who 
	are even more potent in combat than the Adeptus Astartes. They are 
	to the Space Marines as the Emperor is to His primarchs, and it is 
	rumoured that each was created by the Master of Mankind personally. 
	His might permeates them, burns in their eyes and flows through their 
	veins as surely as their blood. As such, the Adeptus Custodes are 
	widely regarded as the deadliest warriors in the galaxy, Human or 
	otherwise.

	Where Space Marines represent the mass-produced, genetically-engineered 
	soldiers of the Imperium of Man, the Adeptus Custodes are a force of 
	individual warriors, each a bastion in their own right and a 
	sentinel of unmatched capability and singular purpose created to 
	counter any possible threat -- Human, alien or Daemonic. These 
	warriors have stood in the presence of the immortal Emperor of 
	Mankind since before the time of the Unification Wars. For ten 
	thousand Terran years and more, the Custodians have stood watch 
	over their lord and master, serving as the Emperor's personal 
	heralds and praetorian bodyguard. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Elite and durable force that can do mid-close range shooting but
	excel at melee combat. Their Ka'tah's amplify their melee power.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Custodes are at tier 4, with a current win-rate of 44%.

	Price: The Custodes are one of the cheapest armies to collect.
	`
	return searchResult
}