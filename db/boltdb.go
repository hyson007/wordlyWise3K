package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var (
	boltBucket = []byte("boltBucket")
	db         *bolt.DB
)

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
	Syllables struct {
		Count int      `json:"count"`
		List  []string `json:"list"`
	} `json:"syllables"`
	Frequency float64 `json:"frequency"`
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(boltBucket)
		return err
	})
}

//grade 2-8 corresponds to wordly wise 2-8
//grade 0 means frequently forget words, it's regardless of grade
func CreateWord(g int, w string, b []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(boltBucket)
		// id64, _ := bucket.NextSequence()
		// id = int(id64)
		key := fmt.Sprint(g) + "_" + w
		// fmt.Println("createword key", key)
		return bucket.Put([]byte(key), b)
	})
	if err != nil {
		return err
	}
	return nil
}

func ListALL() {
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(boltBucket)

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
}

func ListSingleWord(w string) {
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(boltBucket)
		v := bucket.Get([]byte(w))
		fmt.Printf("The Vaule is: %s\n", v)
		return nil
	})
}

func ListWordsByGrade(g int) ([]WordResponse, error) {
	if g < 2 || g > 10 {
		return nil, fmt.Errorf("invalid grade")
	}

	var words []WordResponse

	prefix := []byte(fmt.Sprint(g) + "_")
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(boltBucket)
		cursor := bucket.Cursor()
		for k, v := cursor.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = cursor.Next() {
			var word WordResponse

			if err := json.Unmarshal(v, &word); err != nil {
				log.Println("Unable to unmarshal Json, error: ", err)
				return err
			}
			words = append(words, word)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return words, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// func itob(i int) []byte {
// 	b := make([]byte, 8)
// 	binary.BigEndian.PutUint64(b, uint64(i))
// 	return b
// }

// func btoi(b []byte) int {
// 	return int(binary.BigEndian.Uint64(b))
// }
