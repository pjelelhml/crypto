/**
 * @author paulohmleite
 * @email [paulohmleite1@gmail.com]
 * @create date 2022-07-04 21:55:53
 * @modify date 2022-07-04 21:56:36
 * @desc substitution cipher using golang
 */
package main

import "fmt"

func main() {

	fmt.Println("Substitution Cipher")

	
	var alphabet = [26]rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',}

	var subst = [26]rune{'z' ,'y' ,'x' ,'w' ,'v' ,'u' ,'t' ,'s' ,'r' ,'q' ,'p' ,'o' ,'n' ,'m' ,'l' ,'k' ,'j' ,'i' ,'h' ,'g' ,'f' ,'e' ,'d' ,'c' ,'b' ,'a'}

	// fmt.Println(alphabet)
	fmt.Println(subst)

	var input string
	fmt.Scanln(&input)
	// input = "boat"

	fmt.Println("THe message \"" + input +"\" is encrypted as:")
	for _, currentChar := range input {

		// var currentChar rune = 'c'
		var currentIndex int
		for i := 0; i < len(alphabet) ; i++ {
			if alphabet[i] == currentChar {
				currentIndex = i
				break
			}
		}
		fmt.Printf("%c", subst[currentIndex])
	}	
	
}