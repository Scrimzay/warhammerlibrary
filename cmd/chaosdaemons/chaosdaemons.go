package chaosdaemons

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
	The Chaos Daemons are the malevolent entities that inhabit the Warp, 
	the nightmarish dimension of Chaos that underlies reality in the 
	Warhammer 40,000 universe. These daemons are the embodiments of the 
	dark emotions and desires of mortal beings, given form and sentience 
	by the corrupting power of the four Chaos Gods: Khorne, Tzeentch, 
	Nurgle, and Slaanesh.

	Each Chaos God is associated with specific aspects of mortal 
	existence, and their daemons reflect these aspects in their 
	appearance and behavior. Khorne, the Blood God, commands legions of 
	bloodthirsty daemons like Bloodletters and Flesh Hounds. Tzeentch, 
	the Changer of Ways, is served by cunning and sorcerous entities 
	like Pink Horrors and Lords of Change. Nurgle, the Plague Lord, 
	unleashes pestilent creatures like Plaguebearers and Great Unclean 
	Ones. Slaanesh, the Prince of Pleasure, is attended by seductive 
	and sensual daemons like Daemonettes and Keepers of Secrets.

	Chaos Daemons wage an eternal war against the forces of order and 
	sanity in the galaxy, seeking to corrupt and consume mortal souls 
	and spread the influence of their dark masters. They are summoned 
	to the battlefield by the rituals and sacrifices of Chaos cultists 
	and worshippers, or when the barriers between reality and the Warp 
	grow thin.

	In battle, Chaos Daemons are fearsome and unpredictable foes, their 
	otherworldly forms granting them unnatural strength, speed, and 
	resilience. They are capable of unleashing devastating psychic 
	powers, spreading corruption and madness, and even possessed the 
	ability to split reality itself with their Warp-fueled might.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Diverse horde army relying on their close shooting, psychic powers
	and melee combat to get work done. Shadow of Chaos helps you battle shock and
	deep strike while being backed up by some strong Greater Daemons.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Chaos Daemons are at tier 1, with a current win-rate of 57%.

	Price: The Chaos Daemons are a fairly expensive cost army with huge variety.
	`
	return searchResult
}