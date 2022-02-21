package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hyson007/wordlywise3k/db"
	"github.com/spf13/cobra"
)

var exerciseCmd = &cobra.Command{
	Use:   "exercise",
	Short: "try a random new word from wordlywise grade x, 2 <= x <=10",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args) > 2 {
			fmt.Println("invalid grade number, please key in from 2 to 10 (inclusive)")
			os.Exit(1)
		}
		gradeId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid grade number, please key in from 2 to 10 (inclusive)")
			os.Exit(1)
		}
		allGradeWords, err := db.ListWordsByGrade(gradeId)
		if err != nil {
			fmt.Println("something went wrong during view DB by grade", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Let's start with some exercise, hope you recall the meaning of these words, if not, feel free to reveal the detailed meaning and sample sentences\n")
		fmt.Println()
		for {
			var usrInput string
			rand.Seed(time.Now().Unix())
			randWord := allGradeWords[rand.Intn(len(allGradeWords))]
			fmt.Printf("Dear Ziyu, do you recall:  %s          [y|n] ?\n", randWord.Word)
			fmt.Scanln(&usrInput)
			if len(usrInput) != 1 {
				fmt.Println("Invalid answer, please key in \"y\" or \"n\"")
			} else {
				switch strings.ToUpper(usrInput) {
				case "Y":
					fmt.Println("Well done!")
					fmt.Println()
				case "N":
					if len(randWord.Results) == 1 {
						fmt.Printf("\"%s\" have only one meanings: \n", strings.ToUpper(randWord.Word))
					} else {
						fmt.Printf("\"%s\" have %d different meanings: \n", strings.ToUpper(randWord.Word), len(randWord.Results))
					}

					fmt.Println()
					for idx, meaning := range randWord.Results {
						fmt.Printf("Meaning %d: (%s), %s\n", idx+1, meaning.PartOfSpeech, meaning.Definition)
						if len(meaning.Synonyms) > 0 {
							fmt.Printf("Synonyms of this meanings: [%s] \n", strings.Join(meaning.Synonyms, ", "))
						}
						if len(meaning.Examples) > 0 {
							fmt.Printf("Example: %s \n", strings.Join(meaning.Examples, ","))
						}
						if len(meaning.TypeOf) > 0 {
							fmt.Printf("Under this meaning, the word is a kind of: [%s] \n", strings.Join(meaning.TypeOf[:2], ", "))
						}
						fmt.Println()
					}
					fmt.Println()
					fmt.Printf("Press any key to continue\n")
					fmt.Scanln(&usrInput)
				default:
					fmt.Println("Invalid answer, please key in y or n")
				}

			}
		}
	},
}

func init() {
	RootCmd.AddCommand(exerciseCmd)
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
