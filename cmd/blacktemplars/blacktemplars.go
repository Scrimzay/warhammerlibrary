package blacktemplars

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
	The Black Templars are a Loyalist Second Founding Space Marine 
	Chapter derived from the Imperial Fists' gene-seed and their 
	primarch, Rogal Dorn. Their origin can be traced back to the 
	Imperial Fists' defence of Terra during the Horus Heresy.

	Since that time, the Black Templars have been on the longest 
	Imperial Crusade the Imperium of Man has ever known to prove their 
	loyalty to the Emperor of Mankind. They are not a Codex 
	Astartes-compliant Chapter and maintain a very different Chapter 
	structure and culture than is the norm amongst most of the Adeptus 
	Astartes.

	They are also deeply unusual amongst the forces of the Adeptus 
	Astartes for venerating the Master of Mankind as a literal god, 
	just like the mortal adherents of the Imperial Cult.

	After the Horus Heresy, Rogal Dorn, primarch of the Imperial Fists, 
	resisted the break-up of the Space Marine Legions into smaller 
	Chapters. It was only when the Imperial Fists were almost branded 
	Heretics that Dorn relented, allowing his beloved Legion to be 
	subdivided into Chapters. One of the new Chapters born of this time 
	was the Black Templars. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Close range and melee focused army with durable units. Backed up
	with some long range fire support.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Black Templars are at tier 1, with a current win-rate of 56%.

	Price: The Black Templars are a reasonable cheap army to collect.
	`
	return searchResult
}