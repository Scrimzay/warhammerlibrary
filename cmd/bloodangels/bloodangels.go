package bloodangels

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
	The Blood Angels are one of the 20 First Founding Legions of the 
	Space Marines and were originally the IXth Legion before the Second 
	Founding broke the Legiones Astartes up into separate Chapters of 
	1,000 Space Marines each.

	They are well-known across the galaxy for their bloodthirsty nature 
	in battle, and feared for the curse of flawed gene-seed they carry. 
	The Blood Angels are amongst the longest-living of the Adeptus 
	Astartes, with some of the Chapter's Space Marines having served the 
	Emperor of Mankind for over a thousand standard years.

	Due to recent events, the Blood Angels' numbers were severely 
	depleted. Under the threat of extinction, and in order to quickly 
	replenish their numbers, the Blood Angels were forced to ask their 
	kindred Successor Chapters from subsequent Astartes Foundings for a 
	tithe of warriors from the related Chapters' pools of neophytes, 
	their candidate Space Marines.

	With these tithes of new recruits, and the recent arrival of the 
	Primaris Space Marines during the Indomitus Crusade, the Blood Angels 
	were able to replenish their losses, though they still face a time 
	of trial like none the Chapter has known since the days of the 
	Horus Heresy.  
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Fast, close range and melee focused army with durable units.
	Backed up by some long range fire support.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Blood Angels are at tier 2 with a current win-rate of 52%.

	Price: The Blood Angels are a reasonable cheap army to collect.
	`
	return searchResult
}