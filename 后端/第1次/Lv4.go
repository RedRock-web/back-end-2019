//输入一个单词，然后将一个单词列表中的所有单词依次接龙，并存于数组中
package main
import "fmt"

func WordByWord(wordInput string) []string {
	var aWord,aWordFirst,wordInputLast string
	wordDictionary := [] string {"asshole","bitch","cissy","dork","evil","fool","geek","hooker","idiot","jackass","killer",
		"lunkhead","murderer","nerd","old","prostitute","queer","rascal","son","tramp","ugly","vamp","whore","xyee","yard","zany"}
	var outputDictionary []string

	outputDictionary = append(outputDictionary,string(wordInput))
	wordInputLast = string(wordInput[len(wordInput) - 1])

	for j := 0; j < len(wordDictionary) - 1; j ++ {                         //两次遍历，第一次用于找出能够接龙前面的单词，第二次把前面接龙后的单词作为需要接龙的单词
		for i := 0;i < len(wordDictionary) - 1; i ++ {
			aWord = string(wordDictionary[i])
			aWordFirst = string(aWord[0])
			if wordInputLast == aWordFirst {
				outputDictionary = append(outputDictionary, aWord)
				wordInput = aWord
				wordInputLast = string(wordInput[len(wordInput) - 1])
				break
			}
		}
	}
	fmt.Println(outputDictionary)
	return outputDictionary
}



func main(){
	var wordInput string
	fmt.Scanf("%s", &wordInput)
	WordByWord(wordInput)
}

