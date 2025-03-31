package genestealers

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
	The Genestealer Cults are a unique army in Warhammer 40,000, 
	representing the insidious influence of the Tyranid Genestealers on 
	human societies across the galaxy. These cults arise when a 
	Genestealer infects a human with its genetic material, creating a 
	Genestealer Hybrid. These hybrids can then pass on the genetic 
	taint to their offspring, forming the basis of a Genestealer Cult.
	
	As the cult grows, it infiltrates all aspects of human society, 
	from the lowly underhives to the highest echelons of power. Cultists 
	work tirelessly to undermine the defenses of their host planets, 
	using a combination of subterfuge, sabotage, and guerrilla warfare 
	to weaken the planets' resistance in preparation for the arrival of 
	the Tyranid Hive Fleets.
	
	On the tabletop, Genestealer Cults armies are characterized by their 
	mix of human and Tyranid weaponry and equipment, as well as their 
	ability to field large numbers of expendable troops. They have 
	access to a variety of specialist units, such as the Acolyte 
	Hybrids, Neophyte Hybrids, and Aberrants, each with their own unique 
	abilities and roles within the cult.
	
	One of the defining features of the Genestealer Cults is their 
	ability to set up ambushes and surprise attacks, often emerging from 
	hidden tunnels or infiltrating enemy lines using their "They Came 
	From Below" rule. This, combined with their fanatical devotion to 
	the Tyranid Hive Mind, makes them a unpredictable and dangerous foe 
	to face on the battlefield.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Horde playstyle that can come back from the dead when killed.
	Large but fragile force that can deep strike and use Imperial Guardsmen
	as part of their army.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Genestealers are at tier 2, with a current win-rate of 55%.

	Price: The Genestealers are a very expensive faction to collect.
	`
	return searchResult
}