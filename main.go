package main

import (
	"fmt"
	"os"
	"log"
        "io"
	"net/http"
	"encoding/json"
	"time"
)

type Welcome7 []Welcome7Element

type Welcome7Element struct {
	Word       string     `json:"word"`      
	Phonetic   string     `json:"phonetic"`  
	Phonetics  []Phonetic `json:"phonetics"` 
	Meanings   []Meaning  `json:"meanings"`  
	License    License    `json:"license"`   
	SourceUrls []string   `json:"sourceUrls"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url"` 
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"` 
	Synonyms     []string     `json:"synonyms"`    
	Antonyms     []string     `json:"antonyms"`    
}

type Definition struct {
	Definition string   `json:"definition"`       
	Synonyms   []string `json:"synonyms,omitempty"`         
	Antonyms   []string `json:"antonyms,omitempty"`         
	Example    *string  `json:"example,omitempty"`
}

type Phonetic struct {
	Text      string   `json:"text"`               
	Audio     string   `json:"audio"`              
	SourceURL *string  `json:"sourceUrl,omitempty"`
	License   *License `json:"license,omitempty"`  
}

var client = &http.Client{Timeout: 50 * time.Second}
const colorRed = "\033[0;31m"
const colorGreen = "\033[0;32m"
const colorBlue = "\033[0;34m"
const colorMagenta = "\033[0;35m"
const endColor = "\033[0m"

func searchMeaningByWord (word string) {
	
	urlFormated := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
	req, err := http.NewRequest( "GET",  urlFormated, nil)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}	
	
	var words []Welcome7Element

	err = json.Unmarshal([]byte(b), &words)
	if err != nil {
		fmt.Println("error %s", err)
	}
	
	for _, element := range words {

	    fmt.Fprintf(os.Stdout,"%sWord: %s%s\n", colorRed, endColor, element.Word)
            fmt.Fprintf(os.Stdout,"%sPhonetic: %s%s\n", colorRed, endColor, element.Phonetic)

	    for _, meaning := range element.Meanings {
		fmt.Fprintf(os.Stdout,"%sPart_of_speech - Class_gramatical: %s%s\n", colorRed, endColor, meaning.PartOfSpeech)
                for _, definition := range meaning.Definitions {
                fmt.Fprintf(os.Stdout,"%sDefinition: %s%s\n", colorBlue, endColor, definition.Definition)
                if definition.Example != nil {
                    fmt.Fprintf(os.Stdout, "%sExample:%s %s\n", colorMagenta, endColor, *definition.Example)
                } 

            }
        }
	   fmt.Println("")
           fmt.Println("========================================================================================================")
	   fmt.Println("")    
	}
}

func main() {
	
	argInputUser := os.Args[1:]
	
	if len(argInputUser) < 1 {
		fmt.Fprintf(os.Stdout,"%sPlease, provide at least one argument%s \n", colorRed, endColor)
		fmt.Fprintf(os.Stdout, "%sEx: ./englishFinder -> car <- this is an argument btw %s\n", colorGreen, endColor)
		os.Exit(1)
	}
	
	
	if len(argInputUser) > 1 {
		fmt.Fprintf(os.Stdout,"%sPlease, provide only one argument%s \n", colorRed, endColor)
		fmt.Fprintf(os.Stdout, "%sEx: ./englishFinder car %s\n", colorGreen, endColor)
		os.Exit(1)
	}

	searchMeaningByWord(argInputUser[0]) 

}
