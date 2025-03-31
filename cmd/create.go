package cmd

import (
	"github.com/Scrimzay/warhammerlibrary/cmd/multiInput"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	LogoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#62e411")).Bold(true)
	tipMsgStyle    = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "search",
	Short: "This is the main search command",
	Long:  ".",

	Run: func(cmd *cobra.Command, args []string) {
	choices := []string{"Adeptus Astartes", "Imperium of Man", "Forces of Chaos", "The Xenos Threat"}
	selection := &multiInput.Selection{}
	header := "Select a faction:"
	faction := ""
		

		tprogram := tea.NewProgram(multiInput.InitialModel(choices, selection, header, faction))
		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}

		switch selection.Choice {
        case "Adeptus Astartes":
            faction = "Adeptus Astartes"
            multiInput.AstartesChoice()
        case "Imperium of Man":
            faction = "Imperium of Man"
            multiInput.ImperiumChoice()
        case "Forces of Chaos":
            faction = "Forces of Chaos"
            multiInput.ChaosChoice()
        case "The Xenos Threat":
            faction = "The Xenos Threat"
            multiInput.XenosChoice()
        default:
            fmt.Println("Invalid choice.")
		}
	},
}