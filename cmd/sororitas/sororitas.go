package sororitas

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
	The Adepta Sororitas, colloquially called the "Sisterhood," whose 
	military arm is also known as the Sisters of Battle and formerly as 
	the Daughters of the Emperor, are an all-female division of the 
	Imperium of Man's state church known as the Ecclesiarchy or, more 
	formally, as the Adeptus Ministorum.

	The Sisterhood's Orders Militant serve as the Ecclesiarchy's armed 
	forces, mercilessly rooting out spiritual corruption and heresy 
	within Humanity and every organisation of the Adeptus Terra.

	There is naturally some overlap between the duties of the Sisterhood 
	and the Imperial Inquisition; for this reason, although the 
	Inquisition and the Sisterhood remain entirely separate organisations, 
	the Orders Militant of the Adepta Sororitas also act as the Chamber 
	Militant of the Inquisition's Ordo Hereticus. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Large but fragile force of power armoured Nuns with Guns.
	Excel at mid to close range with their specialist weapons. Has large variety of
	units and blessed with miracle dice to help in clutch moments.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Sororitas are at tier 2, with a current win-rate of 53%.

	Price: The Sororitas is a fairly expensive army to collect.
	`
	return searchResult
}