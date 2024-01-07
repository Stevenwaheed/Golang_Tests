package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

/*
	This function responsible for change the order of the letters.
	If there are two besides letters are similar then it will move to the next letter
	and if they are different then it swap the letters, for example:

	Input: aababba
	Output: abababa
*/
func rearrange(rune_str []rune) string{
	
	// fmt.Println(">>>>>", len(rune_str))

	// iterate over the whole word
	for i := 0; i < len(rune_str) - 3; i++{

		if rune_str[i] == rune_str[i + 1]{
			current := i + 1          // current letter 
			next := current + 1       // next letter

			// if the current and the next letters are similar the go to the next letter 
			for rune_str[current] == rune_str[next]{   
				next++
			}

			// Swap the letters
			var temp = rune_str[current]
			rune_str[current] = rune_str[next]
			rune_str[next] = temp
			// rune_str[current], rune_str[next] = rune_str[next], rune_str[current]
		}
	}
	fmt.Print("\n- Result Of Rearrangment: \n", string(rune_str))

	return string(rune_str)
}

/*
   This function responsible for checking that there is not any similar 
   letters beside each other.
*/
func checkValidation(str string) string{
	
	for i := 0; i < len(str)-1; i++{
		if str[i] == str[i+1]{
			return ""
		}
	} 

	return str
}

func main(){
	fmt.Println("Enter some charachers: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.ToLower(str)      // convert the word to lowercase

	if len(str) >= 1 && len(str) <= 500{
		rune_str := []rune(str)
		rune_str = rune_str[:len(rune_str)-2]

		for i:=0; i <= 15; i++{
			new_str := rearrange(rune_str)
			if new_str == str{
				break
			}else{
				str = new_str
			}
		}

		final_result := checkValidation(str)
		fmt.Println("\n- Final Output: \n", final_result)
	}
}
