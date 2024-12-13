package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LoadMappings charge les mappings des mots depuis un fichier CSV généré à partir du fichier ODS
func LoadMappings(filePath string) (map[string]string, error) {
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
			mappings[record[1]] = record[0] // char -> id
		}
	}
	return mappings, nil
}

func binaryToByte(s string) byte {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0
	}
	return byte(i)
}

func findPosition(slice []string, target string) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1 // Return -1 if the element is not found
}

// SaveWordsAndReplace enregistre les mots uniques dans un fichier et remplace dans le texte par leurs indices
func SaveWordsAndReplace(inputText, mappingsFile, outputFile string) error {
	// Load the character-to-binary mappings
	mappings, err := LoadMappings(mappingsFile)
	if err != nil {
		return fmt.Errorf("error loading mappings: %v", err)
	}

	// Extract unique words from the text
	words := strings.Fields(inputText)
	wordMap := make(map[string]int)
	var uniqueWords []string

	for _, word := range words {
		_, exists := wordMap[word]
		if !exists {
			wordMap[word] = len(uniqueWords)
			uniqueWords = append(uniqueWords, word)
		}
	}

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Write each unique word, encoding its characters using the mappings
	for _, word := range uniqueWords {
		for _, char := range word {
			charStr := string(char)
			if binary, found := mappings[charStr]; found {
				writer.WriteByte(binaryToByte(binary)) // Write the binary representation of the character
			} else {
				writer.WriteString(charStr) // Write the character as-is if no mapping exists
			}
		}
		writer.WriteByte(0) // New line for each word
	}

	// Write a byte separator (0)
	writer.WriteByte(0)

	words = strings.Fields(inputText)
	// Replace words in the input text with their indices and write the updated text
	for _, word := range words {
		//fmt.Println(word)
		pos := findPosition(uniqueWords, word)
		writer.WriteByte(binaryToByte(fmt.Sprintf("%d", pos)))
	}
	writer.WriteByte(0)
	writer.Flush()
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputText, _ := reader.ReadString('\n')
	//inputText := "Ceci test c test est"
	mappingsFile := "./doc/conversion.csv" // Chemin du fichier CSV extrait du fichier ODS
	outputFile := "./output.ari"

	err := SaveWordsAndReplace(inputText, mappingsFile, outputFile)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
	} else {
		fmt.Println("Traitement terminé avec succès.")
	}
}
