package worldeaters

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
	The World Eaters are a Traitor Legion of Chaos Space Marines in the 
	Warhammer 40,000 universe, dedicated to the Blood God Khorne. Once 
	known as the War Hounds, they were one of the original 20 Space 
	Marine Legions created by the Emperor of Mankind to conquer the 
	galaxy during the Great Crusade.

	Under the leadership of their Primarch, Angron, the World Eaters 
	became known for their savagery and brutality on the battlefield, 
	employing a devastating implant known as the Butcher's Nails to 
	enhance their aggression and fury in combat. During the Horus Heresy, 
	Angron and his Legion turned against the Emperor, pledging their 
	allegiance to the Chaos Gods and embracing the bloody worship of
	Khorne.

	In the millennia since their betrayal, the World Eaters have become 
	a fractured and disparate force, roaming the galaxy as individual 
	warbands seeking to spill blood and claim skulls in Khorne's name. 
	They are feared and reviled by all who know of their reputation, 
	their once noble purpose and loyalty to the Emperor replaced by an 
	insatiable thirst for violence and slaughter.

	The World Eaters are known for their brutal close-quarters combat 
	tactics, favoring chainaxes, bolt pistols, and other vicious melee 
	weapons to tear their enemies apart in a whirlwind of blood and fury. 
	They are often accompanied by the Khorne Berzerkers, the most 
	fanatical and savage of their number, who have fully given themselves 
	over to the Blood God's rage.

	With their unmatched ferocity and their devotion to Khorne, the 
	World Eaters continue to carve a bloody path through the galaxy, 
	leaving only carnage and ruin in their wake.
	`
	return summary
}

func TabletopInformation() string {
	searchResult := `
	Playstyle: Melee only army with possibly the strongest melee. Okayishly durable
	and slightly horde army that can utilize very powerful and durable daemons
	and can use blessings of Khorne for even more buffs. Using Angron and being able
	to resurrect him brings a very dangerous and scary army.

	Tier: As of 5/18, on a tier scale of 4 to 1, 4 being worst and 
	1 being best, World Eaters are at tier 3, with a current win-rate of 48%.

	Price: The World Eaters are a fairly expensive cost army.
	`
	return searchResult
}