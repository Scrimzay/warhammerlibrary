package greyknights

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
	The Grey Knights are a secret and mysterious Chapter of Space Marines 
	specifically tasked with combating the dangerous Daemonic entities 
	of the Warp and all those mortals who wield the corrupt power of the 
	Chaos Gods. They were created by the Emperor with the aid of 
	Malcador the Sigillite at the time of the Horus Heresy to serve as 
	Humanity's greatest weapon against the threat posed by the existence 
	of Chaos. They have the honour of being implanted with gene-seed 
	engineered directly from the genome of the Emperor Himself.

	The Grey Knights act as the military arm or "Chamber Militant" of 
	the Ordo Malleus, the Daemonhunters who form the oldest branch of 
	the virtually omnipotent Inquisition. The Grey Knights' fortress-monastery, 
	the Citadel of Titan, is based on Titan, the largest of the moons of 
	the gas giant Saturn in the Sol System. It is kept as a private 
	preserve of the Inquisition.

	The existence of the highly secretive Chapter is virtually unknown 
	outside of the Inquisition and the highest echelons of the Imperial 
	adepta, and is a well-guarded secret enforced by mind-wipes and even 
	assassination of Imperial citizens if necessary. 
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Elite force with lots of movement tricks. Good at close range shooting
	and psychic and melee attacks. Quite durable force as well.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Grey Knights are at tier 1, and best, with a current win-rate of 59%.

	Price: The Grey Knights are a reasonable cheap army to collect.
	`
	return searchResult
}