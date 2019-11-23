//输入一个单词，然后接龙一个单词
package main
import "fmt"

func WordByWord(wordInput string) string {                    //接龙一个单词
	var aWord,aWordFirst,wordInputLast,outputWord string
	wordDictionary := [] string {"asshole","bitch","cissy","dork","evil","fool","geek","hooker","idiot","jackass","killer",
		"lunkhead","murderer","nerd","old","prostitute","queer","rascal","son","tramp","ugly","vamp","whore","xyee","yard","zany"}

	wordInputLast = string(wordInput[len(wordInput) - 1])

		for i := 0;i < len(wordDictionary) - 1; i ++ {
			aWord = string(wordDictionary[i])
			aWordFirst = string(aWord[0])
			if wordInputLast == aWordFirst {
				outputWord = aWord
				fmt.Println(outputWord)
				break
			}
		}


	return outputWord
}



func main(){
	var wordInput string
	fmt.Println("请输入一个单词,如果想要退出请输入EOF")
	fmt.Scanf("%s", &wordInput)
	WordByWord(wordInput)

	for true {
		fmt.Scanf("%s", &wordInput)
		WordByWord(wordInput)
		if (wordInput == "EOF") {
			break
		}
	}
}

