package thousandsons

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
	The Thousand Sons are a Traitor Legion of Chaos Space Marines in 
	the Warhammer 40,000 universe, devoted to the Chaos God Tzeentch. 
	Founded by their Primarch, Magnus the Red, during the Great Crusade, 
	the Thousand Sons were once a legion of powerful psykers and 
	scholars, seeking to unravel the mysteries of the universe and 
	harness the power of the Warp.

	However, their relentless pursuit of knowledge and sorcerous power 
	led to their downfall. During the Horus Heresy, Magnus and his 
	Legion were manipulated by Tzeentch and eventually turned against 
	the Emperor, embracing the corrupting influence of the Changer of 
	Ways.

	In the aftermath of their betrayal, the Thousand Sons were nearly 
	destroyed by the Space Wolves Legion on their homeworld of Prospero. 
	Reduced to a fraction of their former strength, the surviving 
	Thousand Sons fled into the Eye of Terror, where they fully 
	embraced their sorcerous powers and the worship of Tzeentch.

	In the 41st millennium, the Thousand Sons are a legion of powerful 
	sorcerers, their bodies and souls forever twisted by the magic of 
	the Warp. They are led by their Daemon Primarch Magnus and his 
	chief lieutenants, the Exalted Sorcerers, who command the legion's 
	Rubric Marines - once proud Astartes transformed into mindless suits 
	of animated armor by a terrible curse.

	The Thousand Sons wage a never-ending war against the Imperium and 
	the other forces of the galaxy, seeking to accumulate arcane 
	knowledge, unravel the schemes of Tzeentch, and exact revenge upon 
	those who wronged them. With their mastery of sorcery and the 
	insidious power of the Warp, the Thousand Sons remain a potent and 
	sinister force in the Warhammer 40,000 universe.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Slow army with strong psychic powers, rubric marines and 
	terminators, plus using cabal points to give even more powerful abilities.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Thousand Sons are at tier 2, with a current win-rate of 52%.

	Price: The Thousand Sons are a cheap army to collect, especially when using characters.
	`
	return searchResult
}