package orks

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
	The Orks are a warlike, green-skinned alien race in the 
	Warhammer 40,000 universe, known for their love of violence, their 
	ramshackle technology, and their unique psychic gestalt field, the 
	WAAAGH! Orks are genetically engineered creatures, created by the 
	ancient Brain Boyz to fight in their wars, and reproduce through 
	the release of spores upon their death.

	Ork society is built around might makes right, with the biggest and 
	strongest Orks, known as Warbosses, leading their clans to battle. 
	Orks are obsessed with fighting and conquest, living for the thrill 
	of battle and the opportunity to prove their strength against worthy 
	opponents. Despite their seemingly brutish nature, Orks possess a 
	form of innate technological knowledge, allowing them to construct 
	powerful, if crude, weapons, vehicles, and even spaceships from 
	scavenged parts and scrap metal.

	One of the most unique aspects of the Orks is their psychic gestalt 
	field, known as the WAAAGH! This collective psychic energy grows 
	stronger as Orks gather in large numbers and become more excited for 
	battle. The WAAAGH! allows Ork technology to function, often in 
	defiance of the laws of physics, and can even manifest as physical 
	effects on the battlefield, such as the "red ones go fasta" 
	phenomenon.

	In the tabletop game, Ork armies are known for their horde tactics, 
	fielding large numbers of boisterous, aggressive troops supported by 
	ramshackle vehicles and towering war machines like Gargants and 
	Stompas. Orks excel in close combat, with units like Nobz, Meganobz, 
	and the Warboss himself leading the charge, while mobs of Boyz and 
	Gretchin surge forward under the cover of heavy weapons fire from 
	Flash Gitz, Lootas, and Mek Gunz.

	With their love of battle, their unpredictable technology, and the 
	power of the WAAAGH! at their backs, the Orks are a constant and 
	chaotic threat to the other races of the Warhammer 40,000 galaxy.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Huge horde army with huge melee prowess at the cost of no shooting.
	Not very tanky but when they commit a WAAGH they get even stronger and tankier.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, Orks are at tier 2, with a current win-rate of 54%.

	Price: The Orks are a mid to high cost army with huge variety.
	`
	return searchResult
}