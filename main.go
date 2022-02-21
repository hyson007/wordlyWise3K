package main

import (
	"fmt"

	"github.com/hyson007/wordlywise3k/api"
	"github.com/hyson007/wordlywise3k/cmd"
	"github.com/hyson007/wordlywise3k/db"
)

func main() {
	// fmt.Println("foo")//
	//fmt.Println(db.ParseWords(db.Book2))
	must(db.Init("/Users/jackyao/wordlyWise3K/wordlywise3k.db"))

	//initial insert records into boltdb
	// for i := 2; i < 11; i++ {
	// 	must(initFetchGrade(i))
	// }

	cmd.RootCmd.Execute()
	// db.ListSingleWord("3_pronunciation")
	// db.ListALL()
}

func initFetchGrade(grade int) error {
	url := db.Book[grade]
	words, err := db.ParseWordsFromUrl(url)
	must(err)

	for _, word := range words {
		fmt.Println("storing word in to db ...", word)
		wordResByteArray, err := api.RapidAPIGet(word, true)
		if err != nil {
			return err
		}
		err = db.CreateWord(grade, word, wordResByteArray)
		//fmt.Println(grade, word, string(wordResByteArray))
		if err != nil {
			return err
		}
	}
	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
