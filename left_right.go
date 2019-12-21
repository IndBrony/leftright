package leftright

import (
	"errors"
	"fmt"
	"strings"
)

//IndexOfLastCharString is a string wrapper function for IndexOfLastChar
//this function also where the constraint checking took place the input string
func IndexOfLastCharString(s string) (int, error) {
	splitted := strings.Split(s, " ")
	commands := splitted[0]
	words := splitted[1:]

	//Constraint checking + checking if it contains unknown command(s)
	if len(commands) > 100 {
		return -1, errors.New("too much command")
	}
	if len(words) > 100 {
		return -1, errors.New("too much word")
	}
	for _, command := range commands {
		if !strings.ContainsRune(allowedCommands, command) {
			return -1, fmt.Errorf("contains unknown command : %s", string(command))
		}
	}
	return IndexOfLastChar([]byte(commands), strings.Join(words, " "))
}

const allowedCommands = "h1wb"

//IndexOfLastChar will return the position of a cursor in a sentence
//based on vim commands
//first index in allowed command will move the cursor 1 character to the left
//the second index will move it to the right
//the third will move it to the first character of the next word
//and the fourth will move it to the first character of the current word or the previous word if its already on the first character
func IndexOfLastChar(commands []byte, sentence string) (int, error) {
	words := strings.Split(sentence, " ")
	//counts from 1 because we are a human, not a robot XD
	count := 1
	for _, command := range commands {
		//Bind each command based by its index at constant allowedCommands
		switch command {
		case allowedCommands[0]:
			if count <= len(sentence) {
				count++
			}
		case allowedCommands[1]:
			if count > 1 {
				count--
			}
		case allowedCommands[2]:
			if count < len(sentence) {
				//variable to store the position of the last character of every word
				lastCharPos := 0
				// look for the next word by looping through words until the current word
				// get the first char in next word position by adding 1 to the last char position of the current word
				for index, word := range words {
					//if its the last word then ignore the command
					if len(words)-1 == index {
						break
					}
					if lastCharPos += len(word) + 1; !(lastCharPos < count) {
						count = lastCharPos + 1
						break
					}
				}
			}
		case allowedCommands[3]:
			if count < len(sentence) && count > 1 {
				//variable to store the position of the last character of every word
				lastCharPos := 0
				// look for the next word by looping through words until the current word
				// get the first char in current word position by subtracting current word's length to the last char position
				// get the first char in previous word position by  subtracting previous current word's length to the last char position
				for index, word := range words {
					currentWordLength := len(word)
					lastCharPos += currentWordLength + 1
					if lastCharPos > count {
						if count > lastCharPos-currentWordLength {
							count = lastCharPos - currentWordLength
						} else {
							count = lastCharPos - currentWordLength - len(words[index-1]) - 1
						}
						break
					}
				}
			}
		default:
			//returns error if contains unbounded command
			return -1, fmt.Errorf("contains unknown command : %s", string(command))
		}
	}
	return count, nil
}
