package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// LoadMappings charge les mappings des mots depuis un fichier CSV généré à partir du fichier ODS
func invertedMappings(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	mappings := make(map[string]string)

	// Lire les lignes et construire le dictionnaire
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		if len(record) >= 2 {
			mappings[record[0]] = record[1] // char -> id
		}
	}
	return mappings, nil
}

func byteToBinary(s string) byte {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0
	}
	return byte(i)
}

func readMyFile(mappingsFile, inputFile string) (string, error) {
	result := ""
	mappings, err := invertedMappings(mappingsFile)
	if err != nil {
		return "", fmt.Errorf("error loading mappings: %v", err)
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(content))

	var uniqueWords []string

	i := -1
	j := 0
	for content[i+1] != 0 {
		//fmt.Println("_________________")
		uniqueWords = append(uniqueWords, "")
		i++
		for content[i] != 0 {
			//fmt.Println(content[i], fmt.Sprintf("%v", content[i]), mappings[fmt.Sprintf("%v", content[i])])
			uniqueWords[j] += mappings[fmt.Sprintf("%v", content[i])]
			i++
		}
		fmt.Println(uniqueWords[j])
		j++
	}

	//fmt.Println("_______________________________")
	//fmt.Println("valeur de i=", i)
	i++
	i++
	result += uniqueWords[content[i]]
	i++
	for content[i] != 0 {
		result += " "
		result += uniqueWords[content[i]]
		//fmt.Println(uniqueWords[content[i]])
		i++
	}
	return result, nil
}

func main() {
	mappingsFile := "./doc/conversion.csv" // Chemin du fichier CSV extrait du fichier ODS
	inputFile := "./output.ari"

	text, err := readMyFile(mappingsFile, inputFile)
	if err != nil {
		fmt.Printf("Erreur de traitement" + err.Error())
	} else {
		fmt.Println("Traitement terminé avec succès.")
		fmt.Println(text)
	}
}
