package multiInput

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"os/exec"
	"runtime"
	"github.com/Scrimzay/warhammerlibrary/cmd/admech"
	"github.com/Scrimzay/warhammerlibrary/cmd/aeldari"
	"github.com/Scrimzay/warhammerlibrary/cmd/astramilitarum"
	"github.com/Scrimzay/warhammerlibrary/cmd/blacktemplars"
	"github.com/Scrimzay/warhammerlibrary/cmd/bloodangels"
	"github.com/Scrimzay/warhammerlibrary/cmd/chaosdaemons"
	"github.com/Scrimzay/warhammerlibrary/cmd/chaosmarines"
	"github.com/Scrimzay/warhammerlibrary/cmd/custodes"
	"github.com/Scrimzay/warhammerlibrary/cmd/darkangels"
	"github.com/Scrimzay/warhammerlibrary/cmd/deathguard"
	"github.com/Scrimzay/warhammerlibrary/cmd/deathwatch"
	"github.com/Scrimzay/warhammerlibrary/cmd/drukhari"
	"github.com/Scrimzay/warhammerlibrary/cmd/emperorschildren"
	"github.com/Scrimzay/warhammerlibrary/cmd/genestealers"
	"github.com/Scrimzay/warhammerlibrary/cmd/greyknights"
	"github.com/Scrimzay/warhammerlibrary/cmd/imperiumagents"
	"github.com/Scrimzay/warhammerlibrary/cmd/necrons"
	"github.com/Scrimzay/warhammerlibrary/cmd/orks"
	"github.com/Scrimzay/warhammerlibrary/cmd/sororitas"
	"github.com/Scrimzay/warhammerlibrary/cmd/spacemarines"
	"github.com/Scrimzay/warhammerlibrary/cmd/spacewolves"
	"github.com/Scrimzay/warhammerlibrary/cmd/tauempire"
	"github.com/Scrimzay/warhammerlibrary/cmd/thousandsons"
	"github.com/Scrimzay/warhammerlibrary/cmd/tyranids"
	"github.com/Scrimzay/warhammerlibrary/cmd/votann"
	"github.com/Scrimzay/warhammerlibrary/cmd/worldeaters"

	"github.com/PuerkitoBio/goquery"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#62e411")).Bold(true)
	titleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#62e411")).Bold(true)
	enterStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFDAB9")).Bold(true)
	backStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#00CED1")).Bold(true)
	quitStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC0CB")).Bold(true)
)

var factionStyles = map[string]struct {
    titleStyle   lipgloss.Style
    focusedStyle lipgloss.Style
}{
    "Aeldari": {
        titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#fe0000")).Bold(true),
        focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#fe0000")).Bold(true),
    },
    "Drukhari": {
        titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#002e38")).Bold(true),
        focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#002e38")).Bold(true),
    },
    "Tyranids": {
        titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#B67DFF")).Bold(true),
        focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#B67DFF")).Bold(true),
    },
	"Genestealer Cults": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#CBABF6")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#CBABF6")).Bold(true),
	},
	"Leagues of Votann": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#06E5CF")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#06E5CF")).Bold(true),
	},
	"Necrons": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#62e411")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#62e411")).Bold(true),
	},
	"Orks": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#4AA217")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#4AA217")).Bold(true),
	},
	"T'au Empire": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true),
	},
	"Chaos Daemons": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#A40505")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#A40505")).Bold(true),
	},
	"World Eaters": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true),
	},
	"Thousand Sons": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#0089FF")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#0089FF")).Bold(true),
	},
	"Death Guard": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#038B1E")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#038B1E")).Bold(true),
	},
	"Emperor's Children": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#C652FF")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#C652FF")).Bold(true),
	},
	"Chaos Space Marines": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#726D23")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#726D23")).Bold(true),
	},
	"Astra Militarum": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#476F3B")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#476F3B")).Bold(true),
	},
	"Adeptus Mechanicus": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#696969")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#696969")).Bold(true),
	},
	"Adeptus Custodes": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FDE12D")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FDE12D")).Bold(true),
	},
	"Adepta Sororitas": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#CE0902")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#CE0902")).Bold(true),
	},
	"Space Wolves": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#9FD3E1")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#9FD3E1")).Bold(true),
	},
	"Grey Knights": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#BCBCBC")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#BCBCBC")).Bold(true),
	},
	"Deathwatch": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#757575")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#757575")).Bold(true),
	},
	"Dark Angels": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#2B9B10")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#2B9B10")).Bold(true),
	},
	"Blood Angels": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true),
	},
	"Black Templars": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true),
	},
	"Space Marines": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#001AFF")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#001AFF")).Bold(true),
	},
	"Imperium Agents": {
		titleStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#001AFF")).Bold(true),
		focusedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#001AFF")).Bold(true),
	},
}

type Selection struct {
	Choice string
}

func (s *Selection) Update(value string) {
	s.Choice = value
}

type Model struct {
	cursor         int
	choices        []string
	selected       map[int]struct{}
	choice         *Selection
	header         string
	selectionMade  bool
	//factionSummary string
	prevChoice     string
	titleStyle    lipgloss.Style
    focusedStyle  lipgloss.Style
	prevChoices []string
	searchModel tea.Model
	//input string
	//summary string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel(choices []string, selection *Selection, header string, faction string) Model {
    var titleStyle, focusedStyle lipgloss.Style
    if style, ok := factionStyles[faction]; ok {
        titleStyle = style.titleStyle
        focusedStyle = style.focusedStyle
    } else {
        titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#62e411")).Bold(true)
        focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#62e411")).Bold(true)
    }

	choices = append(choices, "General Search")

    return Model{
        choices:     choices,
        selected:    make(map[int]struct{}),
        choice:      selection,
        header:      titleStyle.Render(header),
        titleStyle:  titleStyle,
        focusedStyle: focusedStyle,
    }
}

// Dictionary to store different choice values and their summaries
var factionSummaryFuncs = map[string]func() string{
    "Aeldari":    		  aeldari.FactionSummary,
    "Drukhari":    		  drukhari.FactionSummary,
    "Tyranids":    		  tyranids.FactionSummary,
	"Genestealer Cults":  genestealers.FactionSummary,
	"Leagues of Votann":  votann.FactionSummary,
	"Necrons":            necrons.FactionSummary,
	"Orks":               orks.FactionSummary,
	"T'au Empire":        tauempire.FactionSummary,
	"Chaos Daemons":      chaosdaemons.FactionSummary,
	"World Eaters":       worldeaters.FactionSummary,
	"Thousand Sons":      thousandsons.FactionSummary,
	"Death Guard":        deathguard.FactionSummary,
	"Emperor's Children": emperorschildren.FactionSummary,
	"Chaos Space Marines": chaosmarines.FactionSummary,
	"Astra Militarum":     astramilitarum.FactionSummary,
	"Adeptus Mechanicus":  admech.FactionSummary,
	"Adeptus Custodes":    custodes.FactionSummary,
	"Adepta Sororitas":    sororitas.FactionSummary,
	"Space Wolves":        spacewolves.FactionSummary,
	"Grey Knights":        greyknights.FactionSummary,
	"Deathwatch":          deathwatch.FactionSummary,
	"Dark Angels":         darkangels.FactionSummary,
	"Blood Angels":        bloodangels.FactionSummary,
	"Black Templars":      blacktemplars.FactionSummary,
	"Space Marines":       spacemarines.FactionSummary,
	"Imperium Agents":     imperiumagents.FactionSummary,
}

var tabletopInformationFuncs = map[string]func() string{
    "Aeldari":     		  aeldari.TabletopInformation,
    "Drukhari":    		  drukhari.TabletopInformation,
    "Tyranids":    		  tyranids.TabletopInformation,
	"Genestealer Cults":  genestealers.TabletopInformation,
	"Leagues of Votann":  votann.TabletopInformation,
	"Necrons":            necrons.TabletopInformation,
	"Orks":               orks.TabletopInformation,
	"T'au Empire":        tauempire.TabletopInformation,
	"Chaos Daemons":      chaosdaemons.TabletopInformation,
	"World Eaters":       worldeaters.TabletopInformation,
	"Thousand Sons":      thousandsons.TabletopInformation,
	"Death Guard":        deathguard.TabletopInformation,
	"Emperor's Children": emperorschildren.TabletopInformation,
	"Chaos Space Marines": chaosmarines.TabletopInformation,
	"Astra Militarum":     astramilitarum.TabletopInformation,
	"Adeptus Mechanicus":  admech.TabletopInformation,
	"Adeptus Custodes":    custodes.TabletopInformation,
	"Adepta Sororitas":    sororitas.TabletopInformation,
	"Space Wolves":        spacewolves.TabletopInformation,
	"Grey Knights":        greyknights.TabletopInformation,
	"Deathwatch":          deathwatch.TabletopInformation,
	"Dark Angels":         darkangels.TabletopInformation,
	"Blood Angels":        bloodangels.TabletopInformation,
	"Black Templars":      blacktemplars.TabletopInformation,
	"Space Marines":       spacemarines.TabletopInformation,
	"Imperium Agents":     imperiumagents.TabletopInformation,
}

func getUnitDatacardsURL(faction string) string {
	// Define a map that associates faction names with their URLs
	unitDataCardsURL := map[string]string {
	"Aeldari":     		  "https://assets.warhammer-community.com/eng_wh40k_index_aeldari_dec2024-emhqomhxi2-xisfzvaqvd.pdf",
    "Drukhari":    		  "https://assets.warhammer-community.com/eng_wh40k_index_drukhari_dec2024-bbzjgdu7rx-1gfbm1k7cq.pdf",
    "Tyranids":    		  "https://www.warhammer-community.com/wp-content/uploads/2023/06/L8FE4F808oEwCq9T.pdf",
	"Genestealer Cults":  "https://www.warhammer-community.com/wp-content/uploads/2023/06/BrBEfwS94zTuHrZq.pdf",
	"Leagues of Votann":  "https://assets.warhammer-community.com/eng_warhammer40000_leagues_of_votann_dec24-hrf55254ni-atuo9meuwc.pdf",
	"Necrons":            "https://www.warhammer-community.com/wp-content/uploads/2023/06/H5pO90rzYSAY6dHG.pdf",
	"Orks":               "https://www.warhammer-community.com/wp-content/uploads/2023/06/EE2Pdickp8sNe1NX.pdf",
	"T'au Empire":        "https://www.warhammer-community.com/wp-content/uploads/2023/06/20OdtEKVLiE4H6Zo.pdf",
	"Chaos Daemons":      "https://assets.warhammer-community.com/eng_warhammer40000_chaos_daemons_dec24-rhgvmzjg8n-7kpomhnabi.pdf",
	"World Eaters":       "https://assets.warhammer-community.com/eng_warhammer40000_world_eaters_dec24-jvzacu7lkl-son6xidtae.pdf",
	"Thousand Sons":      "https://assets.warhammer-community.com/eng_wh40k_index_thousand_sons_dec2024-nna7ybtvz3-uks2m2rhfs.pdf",
	"Death Guard":        "https://assets.warhammer-community.com/eng_warhammer40000_death_guard_dec24-7xzye78lo9-pwyiqcu0u4.pdf",
	"Emperor's Children":  "https://assets.warhammer-community.com/warhammer40000_indexes_emperorschildren_eng_24.09-2u4qgvthhw.pdf",
	"Chaos Space Marines": "https://www.warhammer-community.com/wp-content/uploads/2023/06/csv0IuVvYQAndBJE.pdf",
	"Astra Militarum":     "https://assets.warhammer-community.com/eng_warhammer40000_astra_militarum_dec24-opqyscxb6d-om7f8dcfaf.pdf",
	"Adeptus Mechanicus":  "https://www.warhammer-community.com/wp-content/uploads/2023/06/vkzQ2IBbrrCVNzz3.pdf",
	"Adeptus Custodes":    "https://www.warhammer-community.com/wp-content/uploads/2023/06/TE5lPwmnUDrITuGM.pdf",
	"Adepta Sororitas":    "https://www.warhammer-community.com/wp-content/uploads/2023/06/riFjIh9OeKg6AbLZ.pdf",
	"Space Wolves":        "https://assets.warhammer-community.com/eng_warhammer40000_space_wolves_dec24-yczd5448lp-xdf4fp8o2y.pdf",
	"Grey Knights":        "https://assets.warhammer-community.com/eng_wh40k_index_grey_knights_dec2024-6nw6tywb71-exxni51uw5.pdf",
	"Deathwatch":          "https://assets.warhammer-community.com/eng_wh40k_grotmas_index_deathwatch_7_12_2024-jtxo505gbc-skehgnr7np.pdf",
	"Dark Angels":         "https://www.warhammer-community.com/wp-content/uploads/2023/06/C6o7G0zjRSxCUvhK.pdf",
	"Blood Angels":        "https://www.warhammer-community.com/wp-content/uploads/2023/06/YC40Fxov5FhbXFRl.pdf",
	"Black Templars":      "https://assets.warhammer-community.com/eng_wh40k_index_black_templars_dec2024-83mejo2z5x-zepgjlgeo2.pdf",
	"Space Marines":       "https://www.warhammer-community.com/wp-content/uploads/2023/06/uVN1M55L0U3dQeWZ.pdf",
	"Imperium Agents":       "https://www.warhammer-community.com/wp-content/uploads/2023/06/X8QR0o05jbz5dLPX.pdf",
	}

	// Return the URL with the gaven faction
	return unitDataCardsURL[faction]
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }
        case "enter", " ":
            if !m.selectionMade {
                m.selected = map[int]struct{}{m.cursor: {}}
                m.choice.Update(m.choices[m.cursor])
                m.selectionMade = true
                switch m.choice.Choice {
					/*
					The comments and code for each case work the same
					So comments will only be put on the first case
					*/
                case "Adeptus Astartes":
					// Append the current choice to the previous choices slice
                    prevChoices := append(m.prevChoices, m.choice.Choice)
					// Create a new model for the Adeptus Astartes faction
					newModel := AstartesChoice()
					// Set the previous choices of the new model
					newModel.prevChoices = prevChoices
					return newModel, nil
                case "Imperium of Man":
                    prevChoices := append(m.prevChoices, m.choice.Choice)
					newModel := ImperiumChoice()
					newModel.prevChoices = prevChoices
					return newModel, nil
                case "Forces of Chaos":
                    prevChoices := append(m.prevChoices, m.choice.Choice)
					newModel := ChaosChoice()
					newModel.prevChoices = prevChoices
					return newModel, nil
                case "The Xenos Threat":
                    prevChoices := append(m.prevChoices, m.choice.Choice)
					newModel := XenosChoice()
					newModel.prevChoices = prevChoices
					return newModel, nil
                case "Space Marines", "Black Templars", "Blood Angels", "Dark Angels", "Deathwatch", "Grey Knights", "Space Wolves",
                    "Adepta Sororitas", "Adeptus Custodes", "Adeptus Mechanicus", "Astra Militarum", "Imperium Agents",
                    "Chaos Space Marines", "Emperor's Children", "Death Guard", "Thousand Sons", "World Eaters", "Chaos Daemons",
                    "Aeldari", "Drukhari", "Tyranids", "Genestealer Cults", "Leagues of Votann", "Necrons", "Orks", "T'au Empire":
                    // m.prevChoice wasnt being assigned the m.choice.Choice
					// value so just assign m,choice,Choice to the dictionary
					// as an arg since it has the value and get the case
					// that way, im so smart for thinking of this strat
					m.prevChoice = m.choice.Choice
					prevChoices := append(m.prevChoices, m.choice.Choice)
					newModel := factionChoices(m.choice.Choice, m.choice.Choice)
					newModel.prevChoices = prevChoices
					newModel.prevChoice = m.choice.Choice
					return newModel, nil
                case "Faction Summary":
                    return m, nil
                case "Tabletop Information":
                    return m, nil
				case "General Search":
					prevChoices := append(m.prevChoices, m.choice.Choice)
					newModel := GeneralSearch(prevChoices)
					return newModel, nil
				case "Unit Datacards":
					url := getUnitDatacardsURL(m.prevChoice)
					if url != "" {
						openBrowser(url)
					}
					return m, nil
                default:
                    return m, nil
                }
            }
		case "ctrl+c":
			return m, tea.Quit
		/*
		In the "ctrl+x" case, we first check if m.prevChoices has more 
		than one element. If it does, we remove the last choice from 
		m.prevChoices using slicing and store the previous faction in 
		prevFaction.

		We then use a switch statement to determine which model to 
		create based on the prevFaction. If prevFaction matches one of 
		the top-level factions ("Adeptus Astartes", "Imperium of Man", 
		"Forces of Chaos", "The Xenos Threat"), we create the 
		corresponding model using the respective functions. Otherwise, we 
		create a new model using InitialModel with the prevFaction.

		If m.prevChoices has only one element, it means we are at the 
		top-level faction selection. In this case, we create a new 
		model using InitialModel with the initial choices for faction 
		selection.
		*/
		case "ctrl+x":
			if len(m.prevChoices) > 1 {
				// If there are more than one previous choices
				// Remove the last choice from m.prevChoices
				prevChoices := m.prevChoices[:len(m.prevChoices)-1]
				// Get the previous faction from the updated prevChoices
				prevFaction := prevChoices[len(prevChoices)-1]
				var newModel Model
				switch prevFaction {
				case "Adeptus Astartes":
					// If the previous faction is "Adeptus Astartes", create a new model using AstartesChoice()
					newModel = AstartesChoice()
				case "Imperium of Man":
					// If the previous faction is "Imperium of Man", create a new model using ImperiumChoice()
					newModel = ImperiumChoice()
				case "Forces of Chaos":
					// If the previous faction is "Forces of Chaos", create a new model using ChaosChoice()
					newModel = ChaosChoice()
				case "The Xenos Threat":
					// If the previous faction is "The Xenos Threat", create a new model using XenosChoice()
					newModel = XenosChoice()
				default:
					// For any other faction, create a new model using InitialModel with the prevFaction
					newModel = InitialModel(m.choices, m.choice, m.header, prevFaction)
				}
				// Set the prevChoices of the new model to the updated prevChoices
				newModel.prevChoices = prevChoices
				// Return the new model
				return newModel, nil
			} else if len(m.prevChoices) == 1 {
				// If there is only one previous choice
				// Create a new model using InitialModel with the initial faction choices
				newModel := InitialModel([]string{"Adeptus Astartes", "Imperium of Man", "Forces of Chaos", "The Xenos Threat"}, &Selection{}, "Select a faction:", "")
				// Return the new model
				return newModel, nil
			}
        }
    }
    return m, nil
}

func (m Model) View() string {
	if !m.selectionMade {
		s := m.header + "\n\n"
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = m.focusedStyle.Render(">")
			}
			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = m.focusedStyle.Render("x")
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		}
		s += "\n" + enterStyle.Render("Press Enter to confirm choice.")
		s += "\n" + backStyle.Render("Press Ctrl + X to go backwards.")
		s += "\n" + quitStyle.Render("Press Ctrl + C to quit program.")
		return s
	} else {
		switch m.choice.Choice {
		case "Adeptus Astartes", "Imperium of Man", "Forces of Chaos", "The Xenos Threat":
			return fmt.Sprintf("%s chosen. Select a faction:", m.choice.Choice)
		case "Space Marines", "Black Templars", "Blood Angels", "Dark Angels", "Deathwatch", "Grey Knights", "Space Wolves",
			"Adepta Sororitas", "Adeptus Custodes", "Adeptus Mechanicus", "Astra Militarum", "Imperium Agents",
			"Chaos Space Marines", "Emperor's Children", "Death Guard", "Thousand Sons", "World Eaters", "Chaos Daemons",
			"Aeldari", "Drukhari", "Tyranids", "Genestealer Cults", "Leagues of Votann", "Necrons", "Orks", "T'au Empire":
			return fmt.Sprintf("%s chosen. Select an option:", m.choice.Choice)
		case "Faction Summary":
			if summaryFunc, ok := factionSummaryFuncs[m.prevChoice]; ok {
				return fmt.Sprintf("                        %s Faction Summary:\n\n%s", m.prevChoice, summaryFunc())
			}
			return fmt.Sprintf("Faction summary not available for %s.", m.prevChoice)
		case "Tabletop Information":
			if summaryFunc, ok := tabletopInformationFuncs[m.prevChoice]; ok {
				return fmt.Sprintf("                        %s Tabletop Information\n\n%s", m.prevChoice, summaryFunc())
			}
			return fmt.Sprintf("Tabletop Information not available for %s.", m.prevChoice)
		case "Unit Datacards":
			return fmt.Sprintf("Opening browser:")
		case "General Search":
			searchModel, ok := m.searchModel.(SearchModel)
			if !ok {
				return "Error: Invalid search model"
			}
			return searchModel.View()
		default:
			return "Invalid choice. Please try again"
		}
	}
}

func AstartesChoice() Model {
	choices := []string{"Space Marines", "Black Templars", "Blood Angels", "Dark Angels", "Deathwatch", "Grey Knights", "Space Wolves"}
	selection := &Selection{}
	header := "Select an Adeptus Astartes faction:"
	return InitialModel(choices, selection, header, "Adeptus Astartes")
}

func ImperiumChoice() Model {
	choices := []string{"Adepta Sororitas", "Adeptus Custodes", "Adeptus Mechanicus", "Astra Militarum", "Imperium Agents"}
	selection := &Selection{}
	header := "Select an Imperium of Man faction:"
	return InitialModel(choices, selection, header, "Imperium of Man")
}

func ChaosChoice() Model {
	choices := []string{"Chaos Space Marines", "Emperor's Children", "Death Guard", "Thousand Sons", "World Eaters", "Chaos Daemons"}
	selection := &Selection{}
	header := "Select a Force of Chaos faction:"
	return InitialModel(choices, selection, header, "Forces of Chaos")
}

func XenosChoice() Model {
	choices := []string{"Aeldari", "Drukhari", "Tyranids", "Genestealer Cults", "Leagues of Votann", "Necrons", "Orks", "T'au Empire"}
	selection := &Selection{}
	header := "Select a Xenos Threat faction:"
	return InitialModel(choices, selection, header, "Xenos Threat")
}

func factionChoices(faction string, prevChoice string) Model {
	choices := []string{"Faction Summary", "Tabletop Information", "Unit Datacards"}
	selection := &Selection{}
	header := fmt.Sprintf("Select an option for %s:", faction)
	return InitialModel(choices, selection, header, prevChoice)
}

type SearchModel struct {
    Model
    input   string
    summary string
}

func GeneralSearch(prevChoices []string) SearchModel {
    choices := []string{"Search"}
    selection := &Selection{}
    header := "General Search"
    prevChoice := "General Search"

    model := InitialModel(choices, selection, header, prevChoice)
    model.prevChoices = prevChoices

    searchModel := SearchModel{
        Model: model,
        input: "",
        summary: "",
    }

    return searchModel
}

func (m SearchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC:
            return m, tea.Quit
		case tea.KeyCtrlX:
            if len(m.Model.prevChoices) > 1 {
                // If there are more than one previous choices
                // Remove the last choice from m.Model.prevChoices
                prevChoices := m.Model.prevChoices[:len(m.Model.prevChoices)-1]
                // Get the previous faction from the updated prevChoices
                prevFaction := prevChoices[len(prevChoices)-1]
                var newModel Model
                switch prevFaction {
                case "Adeptus Astartes":
                    // If the previous faction is "Adeptus Astartes", create a new model using AstartesChoice()
                    newModel = AstartesChoice()
                case "Imperium of Man":
                    // If the previous faction is "Imperium of Man", create a new model using ImperiumChoice()
                    newModel = ImperiumChoice()
                case "Forces of Chaos":
                    // If the previous faction is "Forces of Chaos", create a new model using ChaosChoice()
                    newModel = ChaosChoice()
                case "The Xenos Threat":
                    // If the previous faction is "The Xenos Threat", create a new model using XenosChoice()
                    newModel = XenosChoice()
                default:
                    // For any other faction, create a new model using factionChoices with the prevFaction
                    newModel = factionChoices(prevFaction, prevFaction)
                }
                // Set the prevChoices of the new model to the updated prevChoices
                newModel.prevChoices = prevChoices
                // Return the new model
                return newModel, nil
            } else if len(m.Model.prevChoices) == 1 {
                // If there is only one previous choice
                // Create a new model using InitialModel with the initial faction choices
                newModel := InitialModel([]string{"Adeptus Astartes", "Imperium of Man", "Forces of Chaos", "The Xenos Threat"}, &Selection{}, "Select a faction:", "")
                // Return the new model
                return newModel, nil
            }
        case tea.KeyRunes:
            m.input += string(msg.Runes)
        case tea.KeyBackspace:
            if len(m.input) > 0 {
                m.input = m.input[:len(m.input)-1]
            }
		case tea.KeySpace:
			m.input += " "
        case tea.KeyEnter:
			// Replace spaces with underscores in the input
			searchQuery := strings.ReplaceAll(m.input, " ", "_")

            // Perform the search using the input
            url := fmt.Sprintf("https://warhammer40k.fandom.com/wiki/Special:Search?query=%s&scope=internal&contentType=&ns%5B0%5D=0&ns%5B1%5D=2900", searchQuery)

			// Make the HTTP request to the URL
			resp, err := http.Get(url)
			if err != nil {
				m.summary = fmt.Sprintf("Error: %s", err)
				return m, nil
			}
			defer resp.Body.Close()

			// Parse the HTML response using goquery
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				m.summary = fmt.Sprintf("Error: %s", err)
				return m, nil
			}

			// Extract the title and URL of the first search result
			var resultURL string
			doc.Find(".unified-search__result").First().Each(func(_ int, s *goquery.Selection) {
				//title = strings.TrimSpace(s.Find("h3").Text())
				resultURL, _ = s.Find("a").Attr("href")
			})

			if resultURL != "" {
				// Make an HTTP requestr to the URL of the first search result
				resultResp, err := http.Get(resultURL)
				if err != nil {
					m.summary = fmt.Sprintf("Error: %s", err)
					return m, nil
				}
				defer resultResp.Body.Close()

				// Parse the HTML response of the search result page using goquery
				resultDoc, err := goquery.NewDocumentFromReader(resultResp.Body)
				if err != nil {
					m.summary = fmt.Sprintf("Error: %s", err)
					return m, nil
				}

				// Extract the desired paragraphs or headers
				var paragraph2, paragraph3, paragraph4 string
				resultDoc.Find(".mw-parser-output p").Eq(2).Each(func(_ int, s *goquery.Selection) {
					paragraph2 += s.Text() + "\n\n"
					//m.summary = fmt.Sprintf("What is my content? %s", paragraph3)
				})
				resultDoc.Find(".mw-parser-output p").Eq(3).Each(func(_ int, s *goquery.Selection) {
					paragraph3 += s.Text() + "\n\n"
					//m.summary = fmt.Sprintf("What is my content? %s", paragraph4)
				})
				resultDoc.Find(".mw-parser-output p").Eq(4).Each(func(_ int, s *goquery.Selection) {
					paragraph4 += s.Text() + "\n\n"
					//m.summary = fmt.Sprintf("What is my content? %s", paragraph4)
				})

				// Format each paragraph
				var formattedParagraphs []string

				for _, paragraph := range []string{paragraph2, paragraph3, paragraph4} {
					if paragraph != "" {
						// Split the content into words
						words := strings.Fields(paragraph)
	
						// Create chunks fo words
						chunkSize := 10
						var chunks []string
						for i := 0; i < len(words); i += chunkSize {
							end := i + chunkSize
							if end > len(words) {
								end = len(words)
							}
							chunks = append(chunks, strings.Join(words[i:end], " "))
						}
	
						// Join the chunks with the newline characters
						formattedParagraph := strings.Join(chunks, "\n")

						// Add leading spaces to each line
						leadingSpaces := strings.Repeat(" ", 10) // Adjust the amount of spaces as needed
						formattedParagraph = leadingSpaces + strings.ReplaceAll(formattedParagraph, "\n", "\n"+leadingSpaces)
						formattedParagraphs = append(formattedParagraphs, formattedParagraph)
					} else {
						m.summary = "No content found."
					}
				}

				// Join the formatted paragraphs with a separator
				separator := "\n\n                                    --- --- ---\n\n"
				formattedContent := strings.Join(formattedParagraphs, separator)

				if formattedContent != "" {
					m.summary = fmt.Sprintf("                                     Summary:\n\n%s", formattedContent)
				} else {
					m.summary = "No content found."
				}
	
			} else {
				m.summary = "No search results found."
			}
           
            return m, nil
        }
    }
    return m, nil
}

func (m SearchModel) View() string {
	s := m.header + "\n\n"
	s += fmt.Sprintf("Enter search query: %s\n\n%s", m.input, m.summary)
	s += "\n" + enterStyle.Render("Press Enter to confirm choice.")
	s += "\n" + backStyle.Render("Press Ctrl + X to go backwards.")
	s += "\n" + quitStyle.Render("Press Ctrl + C to quit program.")
    return s
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Unsupported platform")
	}

	if err != nil {
		log.Print(err)
	}
}