package task2

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func Task2() {
    fmt.Println("Enter your sentence: ")

    reader := bufio.NewReader(os.Stdin)
    sentence, _ := reader.ReadString('\n')
    sentence = strings.TrimSpace(sentence) // Remove any trailing newline characters
    words := strings.Split(sentence, " ")
    
    frequency := make(map[string]int)
    for _, word := range words {
        frequency[strings.ToLower(word)]++
    }

    for key, value := range frequency {
        fmt.Printf("Word: %s, Frequency: %d\n", key, value)
    }
}
