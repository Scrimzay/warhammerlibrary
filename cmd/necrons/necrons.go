package necrons

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
	The Necrons are an ancient race of robotic skeletal warriors that 
	once ruled a vast empire in the Warhammer 40,000 galaxy. Millions of 
	years ago, the Necrons were organic beings known as the Necrontyr, 
	who made a pact with the godlike creatures called the C'tan to 
	transfer their consciousness into immortal metal bodies to escape 
	the pain of their short, disease-ridden lives.

	However, this process erased much of the Necrons' personalities and 
	free will, turning them into soulless, undying servants of the C'tan. 
	Eventually, the Necrons turned on their C'tan masters, shattering 
	them into fragments and imprisoning them within arcane artifacts. 
	The Necrons then entered a long period of dormancy, known as the 
	Great Sleep, hidden away in their massive Tomb Worlds scattered 
	across the galaxy.

	In the 41st millennium, the Necrons have begun to awaken from their 
	slumber, seeking to reclaim their lost empire and restore their 
	former glory. They are armed with incredibly advanced gauss weaponry 
	that strips their targets atom by atom, as well as a variety of 
	exotic and terrifying technologies such as living metal, portal-based 
	transport, and reality-warping devices.

	On the tabletop, Necron armies are known for their durability and 
	ability to regenerate damage, thanks to their Living Metal rule. 
	They field a mix of resilient infantry, powerful vehicles, and 
	towering constructs, all serving under the command of the Necron 
	Overlords and Crypteks, the remnants of the Necrontyr nobility and 
	intelligensia.

	With their ancient technology, implacable will, and undying legions, 
	the Necrons are a major threat to all other factions in the 
	Warhammer 40,000 universe, seeking to establish their supremacy 
	over the galaxy once more.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Slow and very tanky army able to come back after dying.
	Very strong shooting and decent melee and a neverending horde with their
	reanimation abilities.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Necrons are at tier 2, with a current win-rate of 54%.

	Price: The Necrons are a not too expensive faction to collect but have some
	expensive single character models.
	`
	return searchResult
}