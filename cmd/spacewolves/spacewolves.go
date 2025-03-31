package spacewolves

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
	The Space Wolves, known in their own dialect of Juvjk as the Vlka 
	Fenryka or "Wolves of Fenris," are one of the original 20 First 
	Founding Space Marine Chapters, and were once led by their famed 
	primarch, Leman Russ. Originally the VIth Legion of Astartes raised 
	by the Emperor at the dawn of the Great Crusade, the Space Wolves 
	are renowned for their anti-authoritarian ways and their embrace of 
	their homeworld Fenris' savage barbarian culture as well as their 
	extreme deviation from the Codex Astartes in the Chapter's organisation.

	After the Horus Heresy and the resultant Second Founding reforms of 
	the Adeptus Astartes, the Space Wolves Legion was divided into two 
	Chapters: the new Space Wolves Chapter, which was not compliant with 
	the dictates of the Codex Astartes and retained the name of its 
	parent Space Marine Legion, and the second Chapter which took the 
	name of the Wolf Brothers. The Wolf Brothers suffered from rampant 
	mutation of their gene-seed not long after their Founding and were 
	later disbanded.

	Only recently in the Era Indomitus have new successors of the Space 
	Wolves been raised; though many Space Wolves still have doubts about 
	whether the Primaris Marines raised from the genetic material of 
	Leman Russ are true sons of the primarch, the Great Wolf Logan 
	Grimnar has accepted them into the fold as warriors worthy of the 
	Wolf King's heritage. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Close range and melee focused army with fairly durable units.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Space Wolves are at tier 1, with a current win-rate of 58%.

	Price: The Space Wolves are a reasonable cheap army to collect.
	`
	return searchResult
}