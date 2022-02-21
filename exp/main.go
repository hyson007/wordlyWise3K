package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	data1 := []byte(`
	{"word":"calf","results":[{"definition":"the muscular back part of the shank","partOfSpeech":"noun","synonyms":["sura"],"typeOf":["striated muscle","skeletal muscle"],"hasTypes":["mid-calf"],"hasParts":["gastrocnemius muscle","soleus","achilles tendon","tendon of achilles","soleus muscle","gastrocnemius"],"partOf":["shank"]},{"definition":"fine leather from the skin of a calf","partOfSpeech":"noun","synonyms":["calfskin"],"typeOf":["leather"],"hasTypes":["box calf"]},{"definition":"young of domestic cattle","partOfSpeech":"noun","typeOf":["young mammal"],"hasTypes":["dogie","leppy","dogy","maverick"],"memberOf":["cattle","bos taurus","oxen","kine","cows"],"hasParts":["veau","veal"]},{"definition":"young of various large placental mammals e.g. whale or giraffe or elephant or buffalo","partOfSpeech":"noun","typeOf":["young mammal"]}],"syllables":{"count":1,"list":["calf"]},"pronunciation":{"all":"kæf"},"frequency":3.62}`)

	data2 := []byte(`
	{"word":"claw","results":[{"definition":"a mechanical device that is curved or bent to suspend or hold or pull something","partOfSpeech":"noun","synonyms":["hook"],"typeOf":["mechanical device"],"hasTypes":["anchor","ground tackle","pothook","tenterhook"],"partOf":["grapple","grappler","grappling hook","grappling iron","clothes hanger","dress hanger","coat hanger","grapnel"]},{"definition":"a grasping structure on the limb of a crustacean or other arthropods","partOfSpeech":"noun","synonyms":["chela","nipper","pincer"],"typeOf":["appendage","member","extremity"],"partOf":["crustacean"]},{"definition":"a bird's foot","partOfSpeech":"noun","typeOf":["bird's foot"]},{"definition":"attack as if with claws","partOfSpeech":"verb","typeOf":["snipe","assail","assault","attack","lash out","round"],"examples":["The politician clawed his rival"]},{"definition":"clutch as if in panic","partOfSpeech":"verb","typeOf":["prehend","seize","clutch"],"examples":["She clawed the doorknob"]},{"definition":"move as if by clawing, seizing, or digging","partOfSpeech":"verb","typeOf":["work","make"],"examples":["They clawed their way to the top of the mountain"]},{"definition":"scratch, scrape, pull, or dig with claws or nails","partOfSpeech":"verb","typeOf":["scratch up","scrape","scratch"],"hasTypes":["clapperclaw"]},{"definition":"sharp curved horny process on the toe of a bird or some mammals or reptiles","partOfSpeech":"noun","typeOf":["horny structure","unguis"],"hasTypes":["talon","bear claw"]}],"syllables":{"count":1,"list":["claw"]},"pronunciation":{"all":"klɔ"},"frequency":3.67}
	`)

	type WordResponse struct {
		Word    string `json:"word"`
		Results []struct {
			Definition   string   `json:"definition"`
			PartOfSpeech string   `json:"partOfSpeech"`
			Synonyms     []string `json:"synonyms"`
			TypeOf       []string `json:"typeOf"`
			HasTypes     []string `json:"hasTypes"`
			Derivation   []string `json:"derivation"`
			Examples     []string `json:"examples"`
		} `json:"results"`
		// Syllables struct {
		// 	Count int      `json:"count"`
		// 	List  []string `json:"list"`
		// } `json:"syllables"`
		// Pronunciation struct {
		// 	All string `json:"all"`
		// } `json:"pronunciation"`
		// Frequency float64 `json:"frequency"`
	}

	var words []WordResponse
	// words = make([]WordResponse, 2)

	datas := [][]byte{data1, data2}
	// i := 0
	for _, d := range datas {
		var word WordResponse
		if err := json.Unmarshal(d, &word); err != nil {
			fmt.Println("Unable to unmarshal Json, error: ", err)
		}
		words = append(words, word)
		// words[i] = word
		// i += 1
	}

	fmt.Println("printing original data")
	fmt.Println(string(data1))
	fmt.Println(string(data2))
	fmt.Println("===================")
	for _, w := range words {
		fmt.Printf("%+v\n", w)
		fmt.Println()
	}

	// 	a := []byte("a")
	// 	a = append(a, 'b')
	// 	a = append(a, 'c')
	// 	a = append(a, 'd')
	// 	fmt.Println(string(a))
	// }
}
