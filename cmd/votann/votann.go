package votann

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
	The Leagues of Votann are a new faction introduced in Warhammer 40,000, 
	representing the highly advanced and technologically skilled Kin, 
	also known as the Squats. The Leagues are a confederation of 
	self-governing city-states, each with its own unique culture, 
	traditions, and military forces.

	The Kin are descendants of human colonists who settled in the 
	galactic core and developed a distinct society based on advanced 
	technology, trade, and a strong sense of honor and loyalty. They are 
	known for their exceptional skills in engineering, mining, and 
	craftsmanship, as well as their formidable military might.

	On the tabletop, the Leagues of Votann armies are characterized by 
	their highly resilient infantry, advanced weaponry, and powerful 
	vehicles. The backbone of their forces is the Hearthkyn Warriors, 
	skilled soldiers equipped with advanced armor and weapons. They are 
	supported by a variety of specialist units, such as the Einhyr 
	Hearthguard, Hernkyn Pioneers, and the mighty Uthyr Battlesuit.

	One of the defining features of the Leagues of Votann is their use 
	of advanced AI systems known as Cloneskeins. These digital 
	constructs house the memories and personalities of fallen Kin, 
	allowing them to continue serving their Kindred even after death.

	The Leagues of Votann are a proud and resilient force, fiercely 
	dedicated to protecting their own and maintaining their independence 
	in the face of the many threats that plague the Warhammer 40,000 
	universe. Their advanced technology and unwavering loyalty make 
	them a formidable presence on the battlefield.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Slow and tanky army with strong tanks and strong shooting.
	Grants buffs to allies when enemies kill your allies through judgements.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Votann are at tier 4, with a current win-rate of 41%.

	Price: The Votann are a somewhat expensive faction to collect.
	`
	return searchResult
}