package deathwatch

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
	The Deathwatch, known also as the "Long Vigil," and the "Long Watch," 
	is a unique Chapter of the Adeptus Astartes comprised solely of 
	Veteran Space Marines that serves the Ordo Xenos of the Imperial 
	Inquisition as its Chamber Militant. They are the warriors of last 
	resort when the Inquisition needs access to firepower greater than 
	that which the Astra Militarum or a team of its own Acolytes or 
	Throne Agents can provide.

	Across the galaxy there are innumerable hostile alien civilisations 
	that threaten Humanity, from the green-skinned Orks, to the monstrous 
	Tyranids, sadistic Drukhari, spectral C'tan, and undying Necrons.

	It is the sacred task of the Deathwatch to stand sentry against all 
	of these terrible xenos races and many more besides. They are ready 
	to act when such ancient evils rise to threaten Mankind once more. 
	The Space Marines of the Deathwatch form the first, and often only, 
	line of defence against these inhuman horrors.

	Unlike other Space Marines, the ones serving in the Deathwatch are 
	not truly a separate Chapter of the Adeptus Astartes; rather, they 
	are a collection of Veteran Space Marines drawn from all of the 
	different extant Chapters who serve together in the Inquisition's 
	service for a discrete period of time. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Extremely flexible force that can be kitted out to suit any scenario.
	Has teleport strikes to give added movement.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Death Watch are at tier 4, and worst, with a current win-rate of 37%.

	Price: The Death Watch are a reasonable cheap army to collect.
	`
	return searchResult
}