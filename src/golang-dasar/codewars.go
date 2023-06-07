package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func IsTriangle(a, b, c int) bool {
	return (a+b > c) && (a+c > b) && (b+c > a)
}

func wave(words string) []string {
	wordsTemp := []rune(words)
	var result []string
	for i := 0; i < len(words); i++ {
		if string(words[i]) != " " {
			// 	continue
			// }
			// fmt.Println((words[i]))
			wordsTemp[i] = unicode.ToUpper(rune(words[i]))
			result = append(result, string(wordsTemp))
			wordsTemp = []rune(words)
		}
	}
	return result
}

// func wave(words string) []string {
// 	result := []string{}
// 	for index, char := range words {
// 		if char != ' ' {
// 			result = append(result, words[:index]+strings.ToUpper(string(char))+words[index+1:])
// 		}
// 	}
// 	return result
// }

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

func Number(busStops [][2]int) int {
	var v1 int
	var v2 int

	for _, v := range busStops {
		v1 += v[0]
		v2 += v[1]
	}
	return v1 - v2 // Good Luck!
}

//BEST
// func Number(busStops [][2]int) (inBus int) {
//   for _, stopInfo := range busStops {
//     inBus += stopInfo[0] - stopInfo[1]
//   }
//   return // Good Luck!
// }

func ArrayDiff(a, b []int) (result []int) {
	temp := make(map[int]bool)
	for _, v := range b {
		temp[v] = true
	}
	// fmt.Println(temp)
	for _, v := range a {
		if !temp[v] {
			result = append(result, v)
		}
	}
	return //a-b
}

func FindMissingLetter(chars []rune) rune {
	endChar := chars[len(chars)-1:]
	starChar := chars[0]
	alfabet := make(map[rune]bool)
	var res rune
	for _, v := range chars {
		alfabet[v] = true
	}
	for i := starChar; i <= endChar[0]; i++ {
		if !alfabet[i] {
			res = i
		}
	}
	return res
}

//BEST PRACTICE
// func FindMissingLetter(chars []rune) rune {
// 	char := chars[0]
// 	for _, v := range chars[1:] {
// 		if char++; v != char {
// 			break
// 		}
// 	}
// 	return char
// }

func FindOdd(seq []int) (result int) {
	freq := make(map[int]int)
	for _, n := range seq {
		freq[n] = freq[n] + 1
		// fmt.Println("n ", n, " freq", freq[n])

	}
	fmt.Println(freq)
	for n, f := range freq {
		if f%2 != 0 {
			result = n // n= 5 f kemunculanya 3
		}
	}
	return
}

// BEST
// func FindOdd(seq []int) int {
// 	fmt.Println(seq)
// 	res := 0
// 	for _, x := range seq {
// 		res = x ^ +res
// 		fmt.Println("res", res)
// 	}
// 	return res
// }

func toWeirdCase(str string) string {
	var result string
	words := strings.Fields(str)
	for i, word := range words {
		for j, r := range word {
			if j%2 == 0 {
				result += string(unicode.ToUpper(r))
			} else {
				result += string(unicode.ToLower(r))
			}
		}
		if i != len(words)-1 {
			result += " "
		}
	}
	return result
}

func DigPow(n, p int) int {
	var sum int
	for _, v := range strconv.Itoa(n) {
		vInt, _ := strconv.Atoi(string(v))
		sum += int(math.Pow(float64(vInt), float64(p)))
		p++
	}
	if sum%n == 0 {
		return sum / n
	}
	return -1
}

// Fibonance
func Perimeter(n int) int {
	a, b := 0, 1
	var resultFibo int

	for i := 0; i < n+1; i++ {
		a, b = b, a+b
		// fmt.Println("a ", a, "b", b)
		resultFibo += a
		// fmt.Println("resultFibo ", resultFibo)
	}
	return resultFibo * 4
}

// func LongestConsec(strarr []string, k int) string {
// 	var constrarr string
// 	var tempstrarr []string
// 	if len(strarr) == 0 || k == 0 || k > len(strarr) {
// 		return ""
// 	}
// 	for i := 0; i <= len(strarr)-k; i++ {
// 		if k == 1 {
// 			tempstrarr = strarr
// 		}
// 		constrarr = strarr[i]
// 		for j := 1; j < k; j++ {
// 			constrarr += strarr[i+j]
// 		}
// 		tempstrarr = append(tempstrarr, constrarr)
// 		constrarr = ""
// 	}

// 	fmt.Println(tempstrarr)
// 	compareStrarr := func(i, j int) bool {
// 		return len(tempstrarr[i]) > len(tempstrarr[j])
// 	}

// 	sort.Slice(tempstrarr, compareStrarr)
// 	return tempstrarr[0]
// }

// Best Practices
func LongestConsec(strarr []string, k int) string {
	var lower, higher string

	for i := 0; i <= len(strarr)-k; i++ {
		lower = strings.Join(strarr[i:i+k], "")
		if len(lower) > len(higher) {
			higher = lower
		}
	}
	return higher
}

// Highest Scoring Word
func score(word string) int {
	score := 0
	for _, letter := range word {
		score += int(letter) - 96
	}
	return score
}

func High(sentence string) string {
	words := strings.Split(sentence, " ")
	highestScore := 0
	highestScoringWord := ""
	for _, word := range words {
		wordScore := score(word)
		if wordScore > highestScore {
			highestScore = wordScore
			highestScoringWord = word
		}
	}
	return highestScoringWord
}

func Points(games []string) int {
	var points, point int
	for _, game := range games {
		// fmt.Println("i:", i)
		switch {
		case string(game[0]) > string(game[2]):
			point = 3
		case string(game[0]) < string(game[2]):
			point = 0
		case string(game[0]) == string(game[2]):
			point = 1
		}
		points += point
		fmt.Println("points", points)
	}
	return points
}

func DuplicateEncode(word string) string {
	fmt.Println("word-origin:", word)
	var result string
	word = strings.ToLower(word)
	for _, v := range word {
		// fmt.Println("v-string", string(v), "v", v)
		if strings.Count(word, string(v)) == 1 {
			result += "("
		} else {
			result += ")"
		}
	}
	return result

}

// Factorial Decomposition
func Decomp(n int) string {
	factors := make(map[int]int)

	for i := 2; i <= n; i++ {
		factorize(i, factors)
	}
	var factorList []struct {
		Factor   int
		Exponent int
	}
	for factor, exponent := range factors {
		factorList = append(factorList, struct {
			Factor   int
			Exponent int
		}{factor, exponent})
	}

	// mengurutkan slice secara ascending berdasarkan nilai key
	sort.Slice(factorList, func(i, j int) bool {
		return factorList[i].Factor < factorList[j].Factor
	})

	fmt.Println("factorList", factorList)

	var output string
	for _, factor := range factorList {
		if factor.Exponent == 1 {
			output += strconv.Itoa(factor.Factor) + " * "
		} else {
			output += strconv.Itoa(factor.Factor) + "^" + strconv.Itoa(factor.Exponent) + " * "
		}
	}
	output = output[:len(output)-3]
	return output
}

func factorize(n int, factors map[int]int) {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			factors[i]++
			factorize(n/i, factors)
			return
		}
	}
}

//BEST Practices FOR Factorial Decomposition
// func Decomp(n int) string {
//   factors := make([]int, n + 1)
//   res := ""

//   for i := range factors {
//     factors[i] = i
//   }

//   for i := 2; i < len(factors); i++ {
//     if factors[i] == 1 {
//       factors[i] = 0
//       continue
//     }

//     factors[i] = 1

//     for j := i + i; j < len(factors); j += i {
//       for factors[j] % i == 0 {
//         factors[i]++
//         factors[j] /= i
//       }
//     }

//     if factors[i] != 1 {
//       res += fmt.Sprintf(" * %d^%d", i, factors[i])
//     } else {
//       res += fmt.Sprintf(" * %d", i)
//     }
//   }

//   return res[3:] // skip initial " * "
// }

// Longest Common Subsequence
func LCS(s1, s2 string) string {
	var s, result string
	for _, vstr2 := range s2 {
		for _, vstr1 := range s1 {
			if vstr1 == vstr2 {
				s += string(vstr2)
			}
		}
	}
	//removeDuplicate
	for _, c := range s {
		if !strings.Contains(result, string(c)) {
			result += string(c)
		}
	}
	return result
}

// Grasshopper - Terminal game move function
func Move(position int, roll int) int {
	return (roll * 2) + position
}

// [1, 2]   -> 0 1
// [6, 10]	-> 0 1 2 3 4
// [11, 15] -> 0 1 2 3 4
// 1+4+4 = 9
// [1, 5] 0 1 2 3 4
// [6, 10] 0 1 2 3 4
// 4+4 = 8

// func SumOfIntervals(intervals [][2]int) int {
// var result int
// for _, interval := range intervals {
// 	fmt.Println("interval[0]", interval[0], "interval[1]", interval[1])
// 	if interval[0] < interval[1] {
// 		result += interval[1] - interval[0]
// 	}
// }

//		set := make(map[int]bool)
//		for _, valuesInterval := range intervals {
//			for i := valuesInterval[0]; i < valuesInterval[1]; i++ {
//				set[i] = true
//				fmt.Println("value", valuesInterval, "index", i, "set", set)
//			}
//		}
//		fmt.Println("set", set)
//		result = len(set)
//		return result
//	}
func SumOfIntervals(intervals [][2]int) int {
	// Step 1: Sort intervals by the start time.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Step 2: Merge overlapping intervals.
	var mergedIntervals [][2]int
	for _, interval := range intervals {
		if len(mergedIntervals) == 0 {
			mergedIntervals = append(mergedIntervals, interval)
		} else {
			lastInterval := &mergedIntervals[len(mergedIntervals)-1]
			if interval[0] <= lastInterval[1] {
				if interval[1] > lastInterval[1] {
					lastInterval[1] = interval[1]
				}
			} else {
				mergedIntervals = append(mergedIntervals, interval)
			}
		}
	}
	var result int
	for _, interval := range mergedIntervals {
		result += interval[1] - interval[0]
	}
	return result
}

// Get the Middle Character
func GetMiddle(s string) string {
	// fmt.Println("s:", s, "bagi:", string(s[(len(s)/2)-1]), "len:", len(s))
	if len(s)%2 == 0 {
		return string(s[(len(s)/2)-1]) + string(s[len(s)/2])
	} else {
		return string(s[len(s)/2])
	}
}

func DNAtoRNA(dna string) string {
	result := strings.ReplaceAll(dna, "T", "U")
	return result
}

func RemoveChar(word string) string {
	var result string
	len := len(word)
	for i := 0; i < len; i++ {
		if i != 0 && i != len-1 {
			result += string(word[i])
		}
	}
	return result
}

func BinToDec(bin string) (result int) {
	for i := len(bin) - 1; i >= 0; i-- {
		binInt, _ := strconv.Atoi(string(bin[i]))
		result += binInt * int(math.Pow(2, float64(len(bin)-i-1)))
		fmt.Println("for:", binInt, "* 2 ^", i)
	}
	return
}

//OPTIMASI
// func BinToDec(bin string) (result int) {
//     for _, b := range bin {
//         result <<= 1
//         if b == '1' {
//             result++
//         }
//     }
//     return
// }

//BEST Practices
// func BinToDec(bin string) int {
//    r, _ := strconv.ParseInt(bin, 2, 64)
//    return int(r)
// }

func Capitalize(str string, arr []int) string {
	var result string
	for _, v := range arr {
		result = ""
		for i := 0; i < len(str); i++ {
			if i == v {
				result += strings.ToUpper(string(str[i]))
			} else {
				result += string(str[i])
			}
		}
		str = result
	}
	return result
}

func combat(health, damage float64) float64 {
	switch {
	case health > damage:
		return health - damage
	default:
		return 0.0
	}
}

func solve(str string) string {
	var result string
	strarry := []int{0, 0}
	for _, v := range str {
		switch {
		case v >= 97 && v <= 122:
			strarry[0] += 1 //abc
			// fmt.Println("abc=>", string(v), "v:", v)
		case v >= 65 && v <= 90:
			// fmt.Println("ABC=>", string(v), "v:", v)
			strarry[1] += 1 //ABC
		}
	}
	switch {
	case strarry[0] > strarry[1]:
		result = strings.ToLower(str)
	case strarry[0] < strarry[1]:
		result = strings.ToUpper(str)
	default:
		result = strings.ToLower(str)
	}
	fmt.Println("array:", strarry)
	return result
}

// Reverse Words
func ReverseWords(str string) string {
	strarry := strings.Fields(str)
	for i, j := 0, len(strarry)-1; i < j; i, j = i+1, j-1 {
		strarry[i], strarry[j] = strarry[j], strarry[i]
	}
	for _, vstr := range strarry {
		var reverse string
		for i := len(vstr) - 1; i >= 0; i-- {
			reverse += string(vstr[i])
		}
		str = strings.Replace(str, vstr, reverse, 1)
		fmt.Println("STR:", str, "V-Str:", vstr, "Reverse:", reverse)
	}
	return str
}

// BEST PRACTICIES
// func ReverseWords(str string) (result string) {
// 	var word string
// 	for _, r := range str {
// 		if r == ' ' {
// 			result = result + word + " "
// 			word = ""
// 		} else {
// 			word = string(r) + word
// 		}
// 	}
// 	return result + word
// }

func RacePodium(blocks int) [3]int {
	var first_podium, second_podium int
	if blocks%3 == 0 {
		first_podium, second_podium = blocks/3, blocks/3+1
	} else {
		fl_bl := int(blocks / 3)
		ost := blocks - fl_bl*3 + 1
		first_podium, second_podium = fl_bl, fl_bl
		for i := 0; i < ost; i++ {
			if second_podium-first_podium > 1 {
				first_podium += 1
			} else {
				second_podium += 1
			}
		}
		if blocks-first_podium-second_podium-1 != 0 {
			first_podium = second_podium - 1
		}
	}

	third_podium := blocks - first_podium - second_podium

	return [3]int{int(first_podium), int(second_podium), int(third_podium)}

}

func Gcd(x, y uint32) uint32 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	var result []Tuple
	charCountMap := make(map[rune]int)
	if len(text) == 0 {
		return []Tuple{}
	}
	for _, char := range text {
		if _, ok := charCountMap[char]; !ok {
			charCountMap[char] = strings.Count(text, string(char))
			result = append(result, Tuple{Char: char, Count: charCountMap[char]})
		}
	}
	fmt.Println(charCountMap)
	return result
}

func SolveSimpleStringCharacters(s string) []int {
	// var result []int
	// charCountMap := make(map[rune]int)
	// for _, char := range s {
	// 	fmt.Println("char", char, string(char))
	// 	fmt.Println(charCountMap)
	// }
	// return []int{0, 0, 0, 0}
	result := []int{0, 0, 0, 0}
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= 'A' && s[i] <= 'Z':
			result[0] += 1
		case s[i] >= 'a' && s[i] <= 'z':
			result[1] += 1
		case s[i] >= '0' && s[i] <= '9':
			result[2] += 1
		default:
			result[3] += 1
		}
	}
	return result
}

func SumOfIntegersInString(strng string) int {
	re := regexp.MustCompile("\\d+")
	matches := re.FindAllString(strng, -1)
	fmt.Println(matches)

	total := 0
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		total += num
	}
	return total
}

// kyu 5 Weight to weight
func OrderWeight(strng string) string {
	nums := strings.Fields(strng)
	// fmt.Println("nums", nums)
	sort.Slice(nums, func(i, j int) bool {
		si, sj := nums[i], nums[j]
		// fmt.Println("si:", si, "sj:", sj)
		wi, wj := weight(si), weight(sj)
		// fmt.Println("wi,wj:", wi, wj, "si,sj:", si, sj)
		if wi != wj {
			return wi < wj
		}
		return si < sj
	})
	return strings.Join(nums, " ")
}

func weight(s string) int {
	var w int
	// fmt.Println("before", w)
	for _, r := range s {
		w += int(r - '0')
	}
	// fmt.Println("after", w)
	return w
}

// Beginner Series #3 Sum of Numbers
func GetSum(a, b int) int {
	switch {
	case a == 1 && b == 1:
		return 1
	case a > b:
		a, b = b, a
		fallthrough
	default:
		sum := 0
		for i := a; i <= b; i++ {
			sum += i
		}
		return sum
	}
}

func Rps(p1, p2 string) string {
	if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	} else if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	} else if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	} else if p1 == "scissors" && p2 == "rock" {
		return "Player 2 won!"
	} else if p1 == "paper" && p2 == "scissors" {
		return "Player 2 won!"
	} else if p1 == "rock" && p2 == "paper" {
		return "Player 2 won!"
	} else {
		return "Draw!"
	}
}

//BEST PRACTICIES
// func Rps(a, b string) string {
// 	var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"} //Players 2 WINS!
// 	if a == b {
// 		return "Draw!"
// 	}
// 	if m[a] == b {
// 		return "Player 2 won!"
// 	}
// 	return "Player 1 won!"
// }

func SumMix(arr []any) int {
	// var arrStr []any
	var result int
	for _, v := range arr {
		i, err := strconv.Atoi(fmt.Sprintf("%v", v))
		if err != nil {
		}
		result += i
	}
	return result
}

func Litres(time float64) int {
	if time < 2 {
		return 0
	}
	return int(time / float64(2))
}

// BEST PRACTICIES
// func Litres(time float64) int {
// 	return int(time*0.5)
// }

//----------------------------------------------------------------

//	func ToJadenCase(str string) string {
//		var result []string
//		arrStr := strings.Fields(str)
//		for _, v := range arrStr {
//			// upper := strings.ToUpper(v[:1])
//			upper := strings.Replace(v, v[:1], strings.ToUpper(v[:1]), 1)
//			fmt.Println("value", string(v), "upper", upper)
//			result = append(result, upper)
//		}
//		return strings.Join(result, " ")
//	}
//
// BEST Practices
func ToJadenCase(str string) string {
	return strings.Title(str)
}

func FindNextSquare(sq int64) int64 {
	root := math.Sqrt(float64(sq))
	if math.Mod(root, 1) == 0 {
		return int64(math.Pow(root+1, 2))
	} else {
		return -1
	}
}

// accum
func Accum(s string) string {
	var result []string
	for i, v := range s {
		var repeat string
		for j := 0; j < (i + 1); j++ {
			repeat += string(v)
		}
		title := strings.Title(strings.ToLower(repeat))
		result = append(result, title)
	}
	fmt.Println("result", result)
	return strings.Join(result, "-")
}

// BEST PRACTIES
// func Accum(s string) (r string) {
// 	l := make([]string, len(s))
// 	for i, v := range s {
// 		l[i] = strings.ToUpper(string(v)) + strings.Repeat(strings.ToLower(string(v)), i)
// 		fmt.Println(string(v))
// 	}
// 	r = strings.Join(l, "-")
// 	return
// }

func NbDig(n int, d int) int {
	count := 0
	for i := 1; i <= n; i++ {
		square := i * i
		for square > 0 {
			if square%10 == d {
				count++
			}
			square /= 10
		}
	}
	if d == 0 {
		return count + 1
	}
	return count
}

// BEST PRACTIES
// func NbDig(n int, d int) (out int) {
// 	for i := 0; i <= n; i++ {
// 		out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
// 	}

// 	return
// }

func Josephus(items []interface{}, k int) []interface{} {
	result := make([]interface{}, 0)
	index := 0
	for len(items) > 0 {
		index = (index + k - 1) % len(items)
		result = append(result, items[index])
		items = append(items[:index], items[index+1:]...)
		fmt.Println("items...", items, "result", result, "index", index+k-1, "%", len(items), "=>", index)
	}
	return result
}

func alphanumeric(str string) bool {
	// re := regexp.MustCompile("[^0-9A-Za-z_]")
	// re := regexp.MustCompile("\\d+")
	// re := regexp.MustCompile("[^A-Za-z]+")
	if len(str) == 0 {
		return false
	}
	re := regexp.MustCompile("[^0-9A-Za-z]+")
	matches := re.FindAllString(str, -1)
	fmt.Println(matches)
	if matches == nil {
		return true
	}
	return false
}

// BEST PRACTIESl
// func alphanumeric(s string) bool {
// 	r := regexp.MustCompile("^[a-zA-Z0-9]+$")
// 	return r.MatchString(s)
// }1

func SumDigPow(a, b uint64) []uint64 {
	var result []uint64
	for i := a; i <= b; i++ {
		sum := uint64(0)
		if i <= 9 {
			result = append(result, i)
		} else {
			//sum
			// fmt.Println("else", i)
			for j := 0; j < len(strconv.Itoa(int(i))); j++ {
				iString := strconv.Itoa(int(i))
				// fmt.Println(j, "i for j", string(iString[j]))
				iRune, _ := strconv.Atoi(string(iString[j]))
				sum += uint64(math.Pow(float64(iRune), float64(j+1)))
			}
			// fmt.Println("sum", sum)
			if sum == i {
				result = append(result, sum)
				sum = 0
			}
		}
	}
	// fmt.Println("result", result)
	return result
}

// JosephusSurvivor
func JosephusSurvivor(n, k int) int {
	var narry []int
	for i := 1; i <= n; i++ {
		narry = append(narry, i)
	}
	fmt.Println(narry)
	var result []int
	index := 0
	for len(narry) > 0 { // narry = items
		index = (index + k - 1) % len(narry)
		result = append(result, narry[index])
		narry = append(narry[:index], narry[index+1:]...)
		// 	fmt.Println("items...", items, "result", result, "index", index+k-1, "%", len(items), "=>", index)
	}
	fmt.Println(result)
	// fmt.Println(result[len(result)-1])
	return result[len(result)-1]
}

// BEST PRACTICES
// func JosephusSurvivor(n, k int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return (JosephusSurvivor(n-1, k)+k-1)%n + 1
// }

// Moving Zeros To The End
func MoveZeros(arr []int) []int {
	var result []int
	zero := 0
	for _, v := range arr {
		if v == 0 {
			zero += 1
		} else {
			result = append(result, v)
		}
		// fmt.Println("v", v, "zero", zero)
	}
	for i := 1; i <= zero; i++ {
		result = append(result, 0)
	}
	return result
}

//BEST PRACTICES
// func MoveZeros(arr []int) []int {
//  	res:= make([]int,len(arr))
// 	ind:=0
// 	for i:=0;i<len(arr);i++{
// 		if arr[i]!=0{
// 			res[ind]=arr[i]
// 			ind++
// 		}
// 	}
// 	return res
// }

//------------
// Write a function that takes a string of parentheses, and determines if the order of the parentheses is valid.
// The function should return true if the string is valid, and false if it's invalid.

// Examples
// "()"              =>  true
// ")(()))"          =>  false
// "("               =>  false
// "(())((()())())"  =>  true
// Constraints
// 0 <= input.length <= 100

func ValidParentheses(parens string) bool {
	count := 0
	for _, v := range parens {
		if string(v) == "(" {
			count += 1
		} else {
			count -= 1
		}
	}
	if count == 0 {
		return true
	}
	return false
}

// ValidParenthesesV2
func ValidParenthesesV2(parens string) bool {
	stack := []rune{}
	for _, v := range parens {
		if v == '(' {
			stack = append(stack, v)
		} else if len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
		fmt.Println("v", string(v), "stack:", string(stack))
	}
	return len(stack) == 0
}

// FirstNonRepeating
func FirstNonRepeating(str string) string {
	strword := strings.ToLower(str)
	for _, v := range strword {
		if strings.Count(strword, string(v)) == 1 {
			if strings.Contains(str, string(v)) == true {
				return string(v)
			} else {
				return strings.ToUpper(string(v))
			}
		}
		// fmt.Println("value", string(v), "count", strings.Count(str, string(v)))
	}
	return ""
}

// BEST PRACTICIES
// func FirstNonRepeating(str string) string {
// 	for _, c := range str {
// 		if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
// 			return string(c)
// 		}
// 	}
// 	return ""
// }

// Count Sheeps
func CountSheeps(numbers []bool) int {
	var result int
	for _, v := range numbers {
		fmt.Println(v)
		if v == true {
			result += 1
		}
	}
	return result
}

// func CountSheeps(numbers []bool) (count int) {
// 	for _, v := range numbers {
// 		if v { count++ }
// 	}
// 	return
// }

// Invert
func Invert(arr []int) []int {
	var result []int
	for _, v := range arr {
		// if v < 0 {
		// 	v *= -1
		// 	fmt.Println("v < 0:", v)
		// 	result = append(result, v)
		// } else {
		v *= -1
		// 	fmt.Println("v > 0:", v)
		result = append(result, v)
		// }
	}
	return result
}

func PowersOfTwo(n int) []uint64 {
	var result []uint64
	for i := 0; i <= n; i++ {
		result = append(result, uint64(math.Pow(2, float64(n))))
	}
	return result
}

// REPEAT STRING
func RepeatStr(repetitions int, value string) (result string) {
	for i := 0; i <= repetitions; i++ {
		result += value
	}
	return
}

// func RepeatStr(repititions int, value string) string {
//   return strings.Repeat(value, repititions)
// }

func FindMultiples(integer, limit int) []int {
	var result []int
	for i := integer; i <= limit; i += integer {
		result = append(result, i)
	}
	return result
}

// monkeyCount
func monkeyCount(n int) (res []int) {
	for i := 1; i < n; i++ {
		res = append(res, i)
	}
	return
}

// func OddCount(n int) (count int) {
// 	for i := 0; i < n; i++ {
// 		if i%2 == 1 {
// 			count += 1
// 		}
// 	}
// 	return
// }

func OddCount(n int) int {
	return n / 2
}

func ExpressionMatter(a int, b int, c int) int {
	max1 := (a + b + c)
	max2 := a * (b + c)
	max3 := (a + b) * c
	max4 := a * b * c
	max5 := a + b + c

	res1 := int(math.Max(float64(max1), float64(max2)))
	res2 := int(math.Max(float64(res1), float64(max3)))
	res3 := int(math.Max(float64(res2), float64(max4)))
	res := int(math.Max(float64(res3), float64(max5)))
	return res
}

//BEST PRACTICE
// func ExpressionMatter(a int, b int, c int) int {
//     arr := []int{a*(b+c), a*b*c, a+b+c, (a+b)*c}
//     sort.Ints(arr)
//     return arr[len(arr)-1]
// }

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func Greet(name string) string {
	return "result"
}

func StringToArray(str string) []string {
	return strings.Fields(str)
}

func MultiTable(number int) string {
	var result string
	for i := 1; i <= 10; i++ {
		temp := number * i
		if i == 10 {
			result += fmt.Sprintf("%d * %d = %d", i, number, temp)
		} else {
			result += fmt.Sprintf("%d * %d = %d\n", i, number, temp)
		}
	}
	return result
}

//BEST PRACTICES
// func MultiTable(number int) string {
// 	superstring := ""
// 	for i := 1; i < 11; i++ {
// 		superstring += fmt.Sprintf("%d * %d = %d\n", i, number, number * i)
// 	}
// 	return strings.TrimRight(superstring, "\n")
// }

func countSheep(num int) string {
	var result string
	for i := 1; i <= num; i++ {
		result += fmt.Sprintf("%d sheep...", i)
	}
	return result
}

func CamelCase(s string) string {
	// result := strings.Join(strings.Split(s, " "),"")
	return strings.Join(strings.Split(strings.Title(s), " "), "")
}

// func CamelCase(s string) string {
// 	return strings.Replace(strings.Title(s)," ","",-1)
// }

func InArray(array1 []string, array2 []string) []string {
	result := make(map[string]bool)
	for _, a1 := range array1 {
		for _, a2 := range array2 {
			if strings.Contains(a2, a1) {
				result[a1] = true
				break
			}
		}
	}

	var unique []string
	for k := range result {
		unique = append(unique, k)
	}
	sort.Strings(unique)
	if len(unique) == 0 {
		return []string{}
	}

	return unique
}

func ToCamelCase(s string) string {
	fmt.Println("origin", s)
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	matches := re.FindAllString(s, -1)
	fmt.Println(matches)
	// strings.Title(s)
	// fmt.Println("replace", strings.ReplaceAll(s, matches[0], " "))
	// words := strings.Fields(s)
	// fmt.Println(words)
	// return strings.ReplaceAll(strings.Title(strings.ReplaceAll(s, matches[0], " ")), " ", "")
	fmt.Println("before:", string(s[0]))
	// fmt.Println("after", string(strings.ReplaceAll(strings.Title(strings.ReplaceAll(s, matches[0], " ")), " ", "")[0]))
	return strings.Replace(strings.ReplaceAll(strings.Title(strings.ReplaceAll(s, matches[0], " ")), " ", ""), string(strings.ReplaceAll(strings.Title(strings.ReplaceAll(s, matches[0], " ")), " ", "")[0]), string(s[0]), 1)
	return ""
}

func CleanString(s string) string {
	var result, before string
	for _, v := range s {
		if before == "#" && string(v) != "#" {
			fmt.Println(string(v))
			result += string(v)
		}
		fmt.Println("before", before)
		before = string(v)
	}
	fmt.Println("result", result)
	// result := string(s[0]) + string(s[len(s)-1])

	// if len(matches) == 0 {
	// 	return result
	// } else {
	return ""
}

// func EncryptThis(text string) string {
// 	var res []string
// 	texts := strings.Fields(text)
// 	for _, v := range texts {
// 		if len(v) > 2 {
// 			v = string(v[0]) + string(v[len(v)-1]) + v[2:len(v)-1] + string(v[1])
// 		}
// 		res = append(res, strconv.Itoa(int(v[0]))+v[1:])
// 	}
// 	return strings.Join(res, " ")
// }

func EncryptThis(text string) string {
	var res []string
	texts := strings.Fields(text)
	for _, v := range texts {
		if len(v) > 2 {
			v = v[:1] + v[len(v)-1:] + v[2:len(v)-1] + v[1:2]
		}
		res = append(res, strconv.Itoa(int(v[0]))+v[1:])
	}
	return strings.Join(res, " ")
}

// DuplicateCount
func duplicate_count(s1 string) int {
	mapping := make(map[string]bool)
	duplicateCount := strings.ToLower(s1)
	for i, value := range duplicateCount {
		count := strings.Count(duplicateCount, string(value))
		if count > 1 {
			mapping[string(value)] = true
		}
		fmt.Println("i", i, "count", count, "value", string(value))
	}
	fmt.Println("mapping", mapping)

	return len(mapping)
}

// up AND down
// func Arrange(s string) string {

// 	arrays := strings.Fields(s)
// 	lenarrays := len(arrays)
// 	for _, v := range arrays {
// 		var vasci rune
// 		for _, c := range v {
// 			vasci += c
// 		}
// 		fmt.Println("value", v, "ascii", vasci)
// 	}

// 	return ""
// }

func Arrange(s string) string {
	// Split input string into separate words
	words := strings.Fields(s)

	// Sort the words by length and then by ASCII value
	sort.Slice(words, func(i, j int) bool {
		lenI, lenJ := len(words[i]), len(words[j])
		if lenI == lenJ {
			return words[i] < words[j]
		}
		return lenI < lenJ
	})

	// Generate the final string by alternating between increasing and decreasing order of lengths
	var finalStr strings.Builder
	for i, word := range words {
		if i > 0 {
			if len(words[i-1]) == len(word) {
				finalStr.WriteString(" ")
			} else if i%2 == 1 {
				finalStr.WriteString(" >= ")
			} else {
				finalStr.WriteString(" <= ")
			}
		}
		finalStr.WriteString(word)
	}

	return finalStr.String()
}

// A Simple Music Encoder
// func Compress(raw []int) string {
// 	if len(raw) == 0 {
// 		return ""
// 	}

// 	var compressed []string
// 	var current []int
// 	var diff int

// 	for i := 0; i < len(raw); i++ {
// 		if i == 0 {
// 			current = []int{raw[i]}
// 			diff = 0
// 		} else {
// 			if raw[i] == raw[i-1]+diff {
// 				current = append(current, raw[i])
// 			} else {
// 				compressed = append(compressed, compressGroup(current))
// 				current = []int{raw[i]}
// 				if i == 1 {
// 					diff = raw[1] - raw[0]
// 				} else {
// 					diff = raw[i-1] - raw[i-2]
// 				}
// 			}
// 		}
// 	}
// 	compressed = append(compressed, compressGroup(current))

// 	return strings.Join(compressed, ",")
// }

// func CompressEncode(group []int) string {
// 	groupstrings := ""
// 	for _, v := range group {
// 		groupstrings += strconv.Itoa(v)
// 	}
// 	fmt.Println("res", groupstrings)
// 	fmt.Println("---")

// 	power := 1
// 	for i := 0; i < len(groupstrings); i++ {
// 		mapCountString := make(map[int]string)
// 		if i > 0 { //i = 1
// 			before := i - 1                              //before 0
// 			if groupstrings[i] == groupstrings[before] { //value 1223 => 2 == 1
// 				power += 1
// 				mapCountString[power] = string(groupstrings[before])
// 				fmt.Println(string(groupstrings[i]), "==", string(groupstrings[before]), groupstrings[i] == groupstrings[before], "power", power)
// 			}
// 		}
// 		fmt.Println("value", string(groupstrings[i]), "Map", mapCountString)
// 	}

// 	end := groupstrings[len(groupstrings)-1]
// 	fmt.Println("end", string(end))
// 	return ""
// }

// func CompressEncode(group []int) string {
// 	fmt.Println(group)
// 	var before, next int
// 	countDuplicate := make(map[int]int)
// 	for i := 0; i < len(group); i++ {
// 		//next
// 		next = i + 1
// 		if len(group) == next {
// 			next = 0
// 		}

// 		//before
// 		if i > 0 {
// 			before = group[i-1]
// 		}
// 		fmt.Println("i:", i, "group[i]:", group[i], "before:", before, "next:", next)

// 		switch {
// 		case group[i] == group[next]:
// 			countDuplicate[i] = group[i]
// 			// case group[i] == group[before]:
// 		}
// 	}
// 	fmt.Println("MAP", countDuplicate)

// 	return ""
// }

func CompressEncode(group []int) string {
	if len(group) == 0 {
		return ""
	}
	result := strconv.Itoa(group[0])
	count := 1
	// straight 345 to 3-5
	var straight, straightBack []int
	var diffArry []string
	var diveded int
	for i := 1; i < len(group); i++ {
		fmt.Println("i:", i, "group[i-1]", "group[i]", group[i-1], group[i])
		// fmt.Println("diveded ", math.Abs(float64(group[i]-group[i-1])))
		diveded = int(math.Abs(float64(group[i] - group[i-1])))
		diffArry = append(diffArry, strconv.Itoa(diveded))
		if group[i] == group[i-1] {
			count++
			// } else if diveded == divededBefore {
			// 	diffNumber = append(diffNumber, group[i-1], group[i])
		} else {
			if count > 1 {
				result += "*" + strconv.Itoa(count)
			} else if group[i]-1 == group[i-1] {
				straight = append(straight, group[i]-1, group[i]) // 3-5
			} else if group[i]+1 == group[i-1] {
				straightBack = append(straightBack, group[i]+1, group[i]) //
			}
			// fmt.Println("straight", group[i], ":", group[i]-1 == group[i-1])
			// fmt.Println("straight-Back", group[i], ":", group[i]+1 == group[i-1])
			result += "," + strconv.Itoa(group[i])
			count = 1
		}
		// fmt.Println("diveded NOW:", diveded, "SUM", group[i], group[i])
	}
	// fmt.Println("straight", straight, len(straight))
	// fmt.Println("straightBack", straightBack, len(straightBack))
	fmt.Println("diffNumber", diffArry)
	// // [3 4 4 5] = 3-5 ==> strconv.Itoa(straight[0]) + "-" + strconv.Itoa(straight[len(straight)-1])
	divededCount := strings.Join(diffArry, "")
	var resCount string
	resTemp := 0
	for i := 0; i < len(divededCount); i++ {
		counting := strings.Count(divededCount, string(divededCount[i]))
		if counting >= resTemp {
			resCount = string(divededCount[i])
		}
		// fmt.Println(counting, resTemp)
		resTemp = counting
	}
	fmt.Println("resCount", resCount, "res", group)

	var resDiffSame []string
	validasi, _ := strconv.Atoi(resCount)
	var diffBefore int
	for i := 0; i < len(group)-1; i++ {
		if int(math.Abs(float64(group[i]-group[i+1]))) == validasi {
			resDiffSame = append(resDiffSame, strconv.Itoa(group[i]))
			diffBefore = group[i+1]
		}
		if diffBefore == group[i] {
			resDiffSame = append(resDiffSame, strconv.Itoa(group[i]))
		}
		// fmt.Println(group[i], diffBefore, "diffBefore == group[i-1]", diffBefore == group[i])
	}
	fmt.Println("group", group, resDiffSame)

	//straig 345 3-5
	if len(straight) >= 4 {
		var temp string
		for i := straight[0]; i <= straight[len(straight)-1]; i++ {
			if i == straight[0] {
				temp += strconv.Itoa(i)
			} else {
				temp += "," + strconv.Itoa(i)
			}
		}
		fmt.Println("temp", temp, result)
		result = strings.Replace(result, temp, strconv.Itoa(straight[0])+"-"+strconv.Itoa(straight[len(straight)-1]), 1)
	}

	if len(straightBack) >= 4 {
		var temp string
		for i := straightBack[0]; i >= straightBack[len(straightBack)-1]; i-- {
			if i == straightBack[0] {
				temp += strconv.Itoa(i)
			} else {
				temp += "," + strconv.Itoa(i)
			}
		}
		fmt.Println("temp", temp, "LoopBack", result)
		result = strings.Replace(result, temp, strconv.Itoa(straightBack[0])+"-"+strconv.Itoa(straightBack[len(straightBack)-1]), 1)
	}

	//testing
	// for i := straight[0]; i >= straightBack[len(straightBack)-1]; i++ {
	// 	fmt.Println(i)
	// }

	return result
}

func FormatDuration(duration int) string {
	if duration == 0 {
		return "now"
	}
	if duration < 0 {
		return "Invalid input"
	}

	years := duration / 31536000
	duration %= 31536000
	days := duration / 86400
	duration %= 86400
	hours := duration / 3600
	duration %= 3600
	minutes := duration / 60
	duration %= 60
	seconds := duration

	parts := make([]string, 0)
	if years > 0 {
		part := fmt.Sprintf("%d year", years)
		if years > 1 {
			part += "s"
		}
		parts = append(parts, part)
	}
	if days > 0 {
		part := fmt.Sprintf("%d day", days)
		if days > 1 {
			part += "s"
		}
		parts = append(parts, part)
	}
	if hours > 0 {
		part := fmt.Sprintf("%d hour", hours)
		if hours > 1 {
			part += "s"
		}
		parts = append(parts, part)
	}
	if minutes > 0 {
		part := fmt.Sprintf("%d minute", minutes)
		if minutes > 1 {
			part += "s"
		}
		parts = append(parts, part)
	}
	if seconds > 0 {
		part := fmt.Sprintf("%d second", seconds)
		if seconds > 1 {
			part += "s"
		}
		parts = append(parts, part)
	}

	if len(parts) == 1 {
		return parts[0]
	}
	if len(parts) == 2 {
		return strings.Join(parts, " and ")
	}
	return strings.Join(parts[:len(parts)-1], ", ") + " and " + parts[len(parts)-1]
}

func CountBy(x, n int) (result []int) {
	for i := x; i <= x*n; i += x {
		result = append(result, i)
	}
	return
}

// TwoSort
func TwoSort(arr []string) string {
	var result string
	sort.Strings(arr)
	fmt.Println("Arr", arr)
	for i := 0; i < len(arr[0]); i++ {
		if i == (len(arr[0]) - 1) {
			result += string(arr[0][i])
		} else {
			result += string(arr[0][i]) + "***"
		}
	}
	fmt.Println(result)
	return ""
}

// func TwoSort(arr []string) string {
// 	sort.Strings(arr)
// 	chars := strings.Split(arr[0], "")
// 	return strings.Join(chars, "***")
// }

func Digitize(n int) (result []int) {
	str := strconv.Itoa(n)
	for i := len(str) - 1; i >= 0; i-- {
		strConv, _ := strconv.Atoi(string(str[i]))
		result = append(result, strConv)
	}
	return
}

type Result struct {
	C rune // character
	L int  // count
}

func LongestRepetition(text string) Result {
	var res Result
	prevText := ""
	var count int
	cMax := 0
	var char rune
	for _, v := range text {
		if prevText == string(v) {
			count += 1
		} else {
			count = 1
		}

		if count > cMax {
			cMax = count
			char = v
		}
		fmt.Println("v:", string(v), ":", cMax, "prev", prevText)
		prevText = string(v)
	}
	// 	char = tempChar
	// 	cMax = count
	// }
	fmt.Println("char", string(char), "max:", cMax)
	res.C = char
	res.L = cMax
	return res
}

func Well(x []string) string {
	var good, bad int
	for _, v := range x {
		if v == "good" {
			good += 1
		} else {
			bad += 1
		}
	}

	fmt.Println("good", good, "bad", bad)
	switch {
	case good <= 2:
		return "Publish!"
	case good > 2:
		return "I smell a series!"
	default:
		return "Fail!"
	}
}
func BonusTime(salary int, bonus bool) string {
	switch bonus {
	case true:
		return "£" + strconv.Itoa(salary*10)
	default:
		return "£" + strconv.Itoa(salary)
	}
}

func NameValue(my_list []string) []int {
	myAlfabet := make(map[int]int)
	valueAlfabet := 1
	for i := 97; i <= 122; i++ {
		myAlfabet[i] = valueAlfabet
		valueAlfabet++
	}
	// wordValue := 0
	var wordsValueList []int
	for i, list := range my_list {
		wordValue := 0
		for _, valuelist := range list {
			wordValue += myAlfabet[int(valuelist)]
		}
		wordValue = wordValue * (i + 1)
		wordsValueList = append(wordsValueList, wordValue)
	}

	fmt.Println(wordsValueList)
	return wordsValueList
}

// func Meeting(s string) string {
// 	arryStr := strings.Split(s, ";")
// 	sort.Strings(arryStr)
// 	fmt.Println(arryStr)
// 	var vDetail, res []string
// 	for _, v := range arryStr {
// 		vDetail = strings.Split(v, ":")
// 		vDetail[0], vDetail[1] = vDetail[1], vDetail[0]
// 		petternDetail := "(" + strings.Join(vDetail, ",") + ")"

// 		res = append(res, petternDetail)
// 	}
// 	sort.Strings(res)
// 	return strings.ToUpper(strings.Join(res, " "))
// }

func Meeting(s string) string {
	namePairs := strings.Split(s, ";")
	sort.Strings(namePairs)

	var result []string
	for _, pair := range namePairs {
		names := strings.Split(pair, ":")
		names[0], names[1] = names[1], names[0]
		newPair := "(" + strings.ToUpper(strings.Join(names, ", ")) + ")"
		result = append(result, newPair)
	}
	sort.Strings(result)
	return strings.Join(result, " ")
}

func DontGiveMeFive(start int, end int) int {
	var count int
	for i := start; i <= end; i++ {
		if !strings.Contains(strconv.Itoa(i), "5") {
			count += 1
		}
	}
	// fmt.Println(result)
	return count
}

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

// func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
// 	loopingHealt := fighter2.Health + fighter2.Health

// 	for i := 0; i < loopingHealt; i++ {
// 		if i%2 == 0 {
// 			fmt.Println("IF")
// 		} else {
// 			fmt.Println("ELSE")
// 		}

// 		//first
// 		if firstAttacker == fighter1.Name {
// 			fighter2.Health -= fighter1.DamagePerAttack
// 			fighter1.Health -= fighter2.DamagePerAttack
// 			fmt.Println("first", fighter1.Name)
// 		} else {
// 			fighter1.Health -= fighter2.DamagePerAttack
// 			fighter2.Health -= fighter1.DamagePerAttack
// 			fmt.Println("first", fighter2.Name)
// 		}

// 		//result
// 		if fighter1.Health < 0 {
// 			return fighter2.Name
// 		} else if fighter2.Health < 0 {
// 			return fighter1.Name
// 		}
// 		fmt.Println("F1", fighter1, "=>", fighter2)
// 	}
// 	return ""
// }

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	var attacker, defender Fighter
	if fighter1.Name == firstAttacker {
		attacker = fighter1
		defender = fighter2
		fmt.Println("OK")
	} else {
		attacker = fighter2
		defender = fighter1
		fmt.Println("KO")
	}
	fmt.Println("INIT ATT", attacker, "INIT DEFF", defender)

	for {
		defender.Health -= attacker.DamagePerAttack

		if defender.Health <= 0 {
			return attacker.Name
		}
		attacker, defender = defender, attacker
	}
}

func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
}

// func CreateBox(width, length int) [][]int {
// 	power := 1
// 	var temp []string
// 	// var result [][]int
// 	for i := 1; i <= length; i++ {
// 		for j := 1; j <= width; j++ {
// 			fmt.Println("i", i, "j", j)
// 			temp = append(temp, strconv.Itoa(power))
// 		}
// 		fmt.Println(temp)
// 		lengthTemp := len(temp)
// 		var resTemp []string
// 		fmt.Println("resTemp", resTemp, lengthTemp)
// 		// if len(resTemp) != 1 {
// 		// 	resTemp = temp[i : lengthTemp-1]
// 		// } else {
// 		// 	resTemp = []string{}
// 		// }
// 		temp = []string{}

// 	}
// 	return [][]int{}
// }

// func CreateBox(width, length int) [][]int {
// 	box := make([][]int, length)
// 	power := 1
// 	var lenLength int
// 	if length%2 == 0 {
// 		lenLength = length / 2
// 	} else {
// 		lenLength = (length + 1) / 2
// 	}

//		for i := range box {
//			box[i] = make([]int, width)
//			for j := range box[i] {
//				if i == 0 || i == length-1 || j == 0 || j == width-1 || i >= lenLength {
//					box[i][j] = 1
//				} else if i == 1 {
//					box[i][j] = i + 1
//				} else {
//					box[i][j] = i
//				}
//			}
//			power++
//		}
//		return box
//	}

func bandNameGenerator(word string) (result string) {
	if word[0] == word[len(word)-1] {
		result = strings.Title(word + string(word[1:]))
	} else {
		result = "The" + " " + strings.Title(word)
	}
	return
}

// func BubblesortOnce(numbers []int) []int {
// 	sort.Ints(numbers)
// 	return numbers
// }

// func Summation(n int) int {
// 	for i := 1; i <= n; i++ {
// 		sum += i
// 	}
// 	return sum
// }

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func Contamination(text, char string) string {
	result := ""
	for i := 0; i < len(text); i++ {
		result = strings.ReplaceAll(text, text, char)
	}
	return result
}

// func Contamination(text, char string) string {
//   return strings.Repeat(char, len(text))
// }

func FakeBin(x string) string {
	result := ""
	for _, v := range x {
		// fmt.Println(string(v))
		x, _ := strconv.Atoi(string(v))
		if x < 5 {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

// func Game(frank, sam, tom [4]int) bool {
// 	frankWins := 0
// 	samWins := 0

// 	// Simulasikan setiap ronde
// 	for i := 0; i < 4; i++ {
// 		if frank[i] > sam[i] && frank[i] > tom[i] {
// 			frankWins++
// 		} else if sam[i] > frank[i] && sam[i] > tom[i] {
// 			samWins++
// 		}

// 		// Jika Frank atau Sam telah memenangkan dua ronde, langsung keluar dari loop
// 		if frankWins == 2 || samWins == 2 {
// 			break
// 		}
// 	}

// 	// Kembalikan true jika Frank memenangkan permainan, false jika tidak
// 	return frankWins == 2
// }

func SequenceSum(start, end, step int) int {
	res := 0
	for i := start; i <= end; i += step {
		res += i
	}
	return res
}

func RowSumOddNumbers(n int) (result int) {
	var num int = 1
	for i := 0; i < 10000; i++ {
		for j := 0; j <= i; j++ {
			if i == n-1 {
				fmt.Println(num)
				result += 2*num - 1
			}
			num++
		}
	}
	return result
}

func CalculateYears(years int) (result [3]int) {
	cat := 15
	dog := 15
	for i := 0; i < years-1; i++ {
		if i == 0 {
			cat += 9
			dog += 9
		} else {
			cat += 4
			dog += 5
		}
	}
	fmt.Println(cat, " :", dog)
	return [3]int{years, 1, 1}
}

func Cats(start, finish int) int {
	return finish - start
}

func solutionCompareString(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}
	return str[len(str)-len(ending):] == ending
}

// func solution(str, ending string) bool {
// 	return strings.HasSuffix(str, ending)
// }

func FindShort(s string) int {
	words := strings.Fields(s)
	shortestWord := words[0]
	for _, word := range words {
		if len(word) < len(shortestWord) {
			shortestWord = word
		}
	}
	fmt.Println("Shortest word:", shortestWord)
	return len(shortestWord)
}

func FreqSeq(str string, sep string) string {
	freqMap := make(map[rune]int)
	freqSeq := ""

	for _, char := range str {
		freqMap[char]++
	}
	fmt.Println("freqMap : ", freqMap)
	for _, char := range str {
		freqSeq += strconv.Itoa(freqMap[char]) + sep
	}

	return freqSeq[:len(freqSeq)-1]
}

func HighAndLow(in string) string {
	arry := strings.Fields(in)
	intSlice := make([]int, len(arry))
	for i, str := range arry {
		num, _ := strconv.Atoi(str)
		intSlice[i] = num
	}
	fmt.Println(intSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	return strconv.Itoa(intSlice[len(intSlice)-1]) + " " + strconv.Itoa(intSlice[0])
}

func HasUniqueChar(str string) bool {
	for _, v := range str {
		fmt.Println(string(v), ":", strings.Count(str, string(v)))
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true
}

func ModifyMultiply(str string, loc, num int) string {
	arry := strings.Fields(str)
	var result string
	if num == 0 {
		num = 1
	}
	for i := 0; i < num; i++ {
		if i == num {
			result = result + arry[loc]
		} else {
			result = result + arry[loc] + "-"
		}
	}
	return result
}

func ContainAllRots(strng string, arr []string) (result bool) {
	for i := 0; i < len(strng); i++ {
		rotatedStr := strings.Join([]string{strng[i:], strng[:i]}, "")
		if !contains(arr, rotatedStr) {
			return false
		}
	}
	return true
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func Spacify(s string) string {
	return strings.Join(strings.Split(s, ""), " ")
}

func Calc(s string) int {
	var strS1 string
	for _, v := range s {
		fmt.Println(string(v), ":", v-2)
		temp := v - 2
		strS1 += strconv.Itoa(int(temp))
	}
	strS3 := strings.Split(strings.ReplaceAll(strS1, "7", "1"), "")
	strS2 := strings.Split(strS1, "")
	fmt.Println("strS2 :", strS2)
	fmt.Println("strS3 :", strS3)

	sumS2 := 0
	for _, v := range strS2 {
		vInt, _ := strconv.Atoi(v)
		sumS2 += vInt
	}
	fmt.Println("SUM", sumS2)

	sumS3 := 0
	for _, v := range strS3 {
		vInt, _ := strconv.Atoi(v)
		sumS3 += vInt
	}
	fmt.Println("SUM", sumS3)
	return sumS2 - sumS3
}

func NthEven(n int) int {
	return 2*n - 2
}

func Permutations(s string) []string {
	runes := []rune(s)
	permutations := []string{}
	permute(runes, 0, len(runes)-1, &permutations)

	return permutations
}

func permute(runes []rune, left, right int, permutations *[]string) {
	if left == right {
		*permutations = append(*permutations, string(runes))
	} else {
		for i := left; i <= right; i++ {
			runes[left], runes[i] = runes[i], runes[left]
			permute(runes, left+1, right, permutations)
			runes[left], runes[i] = runes[i], runes[left]
		}
	}
}

// 10010001
func findZeroBinnary(number int) int {
	binary := fmt.Sprintf("%b", number)
	count := 0
	result := 0

	for _, v := range binary {
		if string(v) == "0" {
			count++
		} else {
			if count > result {
				result = count
				count = 0
			}
		}
	}
	fmt.Println(result)
	return result

}

// func SimpleRemoveDuplicates(arr []int) []int {
// 	// test := []int{3, 4, 4, 3, 6, 3}
// 	unique := make(map[int]bool)
// 	var result []int
// 	// for i := len(arr) - 1; i >= 0; i-- {
// 	// 	if !unique[arr[i]] {
// 	// 		unique[arr[i]] = true
// 	// 	}
// 	// }

// 	// fmt.Println(test)
// 	for _, v := range arr {
// 		if !unique[v] {
// 			unique[v] = true
// 			result = append(result, v)
// 		}
// 	}
// 	fmt.Println(unique)

// 	// var result []int
// 	// for v := range unique {
// 	// 	result = append(result, v)
// 	// }
// 	// fmt.Println(result)

// 	return result
// }

func SimpleRemoveDuplicates(arr []int) []int {
	unique := make(map[int]bool)
	var result []int
	for i := len(arr) - 1; i >= 0; i-- {
		if !unique[arr[i]] {
			unique[arr[i]] = true
			result = append(result, arr[i])
		}
	}

	// Reverse the result slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func minOnArray(arr []int) int {
	test := []int{4, 1, 2, 4, 5}
	prev := 0
	var result []int
	for i := 0; i < len(test); i++ {

		fmt.Println("core", test[i]+prev)
		fmt.Println(test[i+1:])
		var count int
		for _, v := range test[i+1:] {
			count += v
		}
		fmt.Println("sum", count)
		fmt.Println(count, "-", (test[i] + prev), "=", count-(test[i]+prev))
		// temp := count - (test[i] + prev)
		// if temp > 0 {
		result = append(result, int(math.Abs(float64(count-(test[i]+prev)))))
		// }
		prev += test[i]

	}

	sort.Ints(result)
	fmt.Println(result)

	return result[0]
}

func LengthAndTwoValues(n int, firstValue string, secondValue string) []string {

	result := make([]string, 0)

	for i := 0; i < n; i++ {
		if len(result) < n {
			result = append(result, firstValue)
		}
		if len(result) < n {
			result = append(result, secondValue)
		}

	}
	fmt.Println(result)

	return []string{}
}

func ComposingSquaredStrings(s1, s2 string) string {
	arr1 := strings.Fields(strings.Trim(s1, "\n"))
	arr2 := strings.Fields(strings.Trim(s2, "\n"))

	var tempRes1 []string
	var tempRes2 []string
	var s1Composes string
	var s2Composes string
	var result []string

	for i := 0; i < len(arr1); i++ {
		s1Composes = string(arr1[i][:i+1])
		tempRes1 = append(tempRes1, s1Composes)
	}
	for j := len(arr2) - 1; j >= 0; j-- {
		s2Composes = string(arr2[j][:j+1])
		tempRes2 = append(tempRes2, s2Composes)
	}

	for i := 0; i < len(tempRes1); i++ {
		result = append(result, tempRes1[i]+tempRes2[i])
	}
	return strings.Join(result, "\n")
}

//BEST PRACTICE
// func Compose(s1, s2 string) string {
//     s1new := strings.Split(s1, "\n")
//     s2new := strings.Split(s2, "\n")
//     str := []string{}
//     n := len(s1new)
//     for i := 0 ; i < n; i++ {
//        str = append(str, s1new[i][0:i+1] + s2new[n-i-1][0:n-i])
//     }
//     return strings.Join(str, "\n")
// }

func main() {
	//MEXICAN WAVE
	fmt.Println("----------")
	fmt.Println(wave("a a a a a "))

	fmt.Println("------- SortNumbers --------")
	fmt.Println(SortNumbers([]int{1, 2, 10, 50, 5}))

	fmt.Println("------- Number --------")
	fmt.Println(Number([][2]int{{10, 0}, {3, 5}, {5, 8}}))

	fmt.Println("------ arrayDiff -------")
	fmt.Println(ArrayDiff([]int{1, 2}, []int{1}))

	fmt.Println("---- findMissingTheLetter -----")
	fmt.Println(FindMissingLetter(([]rune{'a', 'b', 'c', 'd', 'f'})))

	fmt.Println("---- findOdd -----")
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))

	fmt.Println("---- toWeirdCase -----")
	fmt.Println(toWeirdCase("abc def"))

	fmt.Println("--- DigPow ----")
	fmt.Println(DigPow(46288, 3))

	fmt.Println("--- Perimeter ----")
	fmt.Println(Perimeter(7))

	fmt.Println("---LongestConsec----")
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))

	fmt.Println("---Highest Scoring Word----")
	fmt.Println(High("man i need a taxi up to ubud"))

	fmt.Println("----Total amount of points----")
	fmt.Println(Points([]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"})) //10

	fmt.Println("----Duplicate Encoder----")
	// fmt.Println(DuplicateEncode("din"))
	fmt.Println(DuplicateEncode("((((()"))

	fmt.Println("----Factorial decomposition----")
	fmt.Println(Decomp(17))

	fmt.Println("----Longest Common Subsequence----")
	// fmt.Println(LCS("132535365", "123456789"))
	fmt.Println(LCS("abcdef", "abc")) // Output: "abc"

	fmt.Println("----Sum of Intervals----")
	fmt.Println(SumOfIntervals([][2]int{{-1_000_000_000, 1_000_000_000}}))
	// fmt.Println(SumOfIntervals([][2]int{{1, 5}})) // Output 4

	fmt.Println("-------Get the Middle Character---------")
	fmt.Println(GetMiddle("testing"))
	fmt.Println(GetMiddle("A"))

	fmt.Println("-------Remove character---------")
	fmt.Println(RemoveChar("person"))

	fmt.Println("-------BinToDec---------")
	fmt.Println(BinToDec("1"))
	fmt.Println(BinToDec("10"))

	fmt.Println("-------Capitalize---------")
	fmt.Println(Capitalize("abcdef", []int{1, 2, 5}))
	fmt.Println(Capitalize("abracadabra", []int{2, 6, 9, 10}))

	fmt.Println("-------Fix string case---------")
	fmt.Println(solve("coDe"))

	fmt.Println("-------Reverse words---------")
	fmt.Println(ReverseWords("stressed desserts"))

	fmt.Println("-------Race Podium---------")
	fmt.Println(RacePodium(11))
	fmt.Println(RacePodium(1000000))

	fmt.Println("-------GCD---------")
	fmt.Println(Gcd(30, 12))

	fmt.Println("-------Ordered Count of Characters---------")
	fmt.Println(OrderedCount("abracadabra"))
	fmt.Println(OrderedCount(""))

	fmt.Println("-------Simple string characters---------")
	fmt.Println(SolveSimpleStringCharacters("Codewars@codewars123.com"))

	fmt.Println("-------Sum of integers in string---------")
	fmt.Println(SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))

	fmt.Println("-------Order Weight---------")
	fmt.Println(OrderWeight("56 65 74 100 99 68 86 180 90"))

	fmt.Println("-------Beginner Series #3 Sum of Numbers---------")
	fmt.Println(GetSum(1, 0))
	fmt.Println(GetSum(1, 1))
	fmt.Println(GetSum(-1, 2))

	fmt.Println("-------SUM MIX---------")
	fmt.Println(SumMix([]any{"3", 6, 6, 0, "5", 8, 5, "6", 2, "0"})) //21

	fmt.Println("-------ToJadenCase---------")
	fmt.Println(ToJadenCase("most trees are blue")) //21

	fmt.Println("--------- ACCUM -----------")
	fmt.Println(Accum("ZpglnRxqenU"))
	fmt.Println(Accum("MjtkuBovqrU"))

	fmt.Println("--------- ACCUM -----------")
	fmt.Println(NbDig(5750, 0))

	fmt.Println("--------- Josephus Permutation -----------")
	fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7}, 3))

	fmt.Println("--------- alphanumeric -----------")
	fmt.Println(alphanumeric("43534h56jmTHHF3k"))
	fmt.Println(alphanumeric(""))

	fmt.Println("--------- Eureka!! -----------")
	fmt.Println(SumDigPow(1, 100))

	fmt.Println("--------- JosephusSurvivor -----------")
	fmt.Println(JosephusSurvivor(7, 3))

	fmt.Println("--------- MoveZeros -----------")
	fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))

	fmt.Println("--------- ValidParenthesesV2 -----------")
	fmt.Println(ValidParenthesesV2("())("))

	fmt.Println("--------- FirstNonRepeating -----------")
	fmt.Println(FirstNonRepeating("sTreSS"))

	fmt.Println("--------- Count Sheeps -----------")
	fmt.Println(CountSheeps([]bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}))

	fmt.Println("--------- Invert -----------")
	fmt.Println(Invert([]int{1, -2, 3, -4, 5}))

	fmt.Println("--------- FindMultiples -----------")
	fmt.Println(FindMultiples(5, 25))

	fmt.Println("--------- Odd Count -----------")
	fmt.Println(OddCount(15))

	fmt.Println("--------- NoSpace -----------")
	fmt.Println(NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B"))

	fmt.Println("--------- (MultiTable(5) -----------")
	fmt.Println(MultiTable(5))

	fmt.Println("--------- CamelCase -----------")
	fmt.Println(CamelCase("camel case method"))

	fmt.Println("--------- InArray -----------")
	fmt.Println(InArray([]string{"live", "arp", "strong"}, []string{"lively", "alive", "harp", "sharp", "armstrong"}))

	fmt.Println("--------- ToCamelCase -----------")
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
	fmt.Println(ToCamelCase("the_Stealth_Warrior"))

	fmt.Println("--------- CleanString -----------")
	fmt.Println(CleanString("abc#d##c"))
	fmt.Println(CleanString("abc#d###"))

	fmt.Println("--------- EncryptThis -----------")
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))

	fmt.Println("--------- DuplicateCount -----------")
	fmt.Println(duplicate_count("abcdeaB11"))

	fmt.Println("--------- Compress -----------")
	// fmt.Println(CompressEncode([]int{1, 2, 2, 3}))
	fmt.Println("--------- FormatDuration -----------")
	fmt.Println(FormatDuration(1))

	fmt.Println("--------- CountBy -----------")
	fmt.Println(CountBy(1, 5))

	fmt.Println("--------- TwoSort -----------")
	fmt.Println(TwoSort([]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}))

	fmt.Println("--------- Digitize -----------")
	fmt.Println(Digitize(35231))

	fmt.Println("--------- (LongestRepetition -----------")
	// fmt.Println(LongestRepetition("cbdeuuu900"))
	fmt.Println(LongestRepetition("aaaabb"))
	fmt.Println("---")
	fmt.Println(LongestRepetition("ba"))
	// fmt.Println(LongestRepetition("bbbaaabaaaa"))

	fmt.Println("--------- WELL -----------")
	fmt.Println(Well([]string{"bad", "bad", "bad", "bad", "bad"}))
	fmt.Println(Well([]string{"bad", "bad", "good", "bad", "bad", "good", "bad", "bad", "bad", "bad", "bad", "good", "good", "bad", "good", "bad"}))

	fmt.Println("--------- Name Value-----------")
	// fmt.Println(NameValue([]string{"abc"}))
	fmt.Println(NameValue([]string{"abc", "abc", "abc", "abc"}))

	fmt.Println("--------- MEETING-----------")
	fmt.Println(Meeting("Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"))
	fmt.Println(Meeting("John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell"))

	fmt.Println("--------- DontGiveMeFive-----------")
	fmt.Println(DontGiveMeFive(-1, 33))
	fmt.Println("----")
	fmt.Println(DontGiveMeFive(1, 9))

	fmt.Println("--------- DeclareWinner -----------")
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew"))

	fmt.Println("--------- SetAlarm -----------")
	fmt.Println(SetAlarm(true, true))

	// fmt.Println("--------- CreateBox -----------")
	// fmt.Println(CreateBox(7, 8))

	fmt.Println("--------- bandNameGenerator -----------")
	fmt.Println(bandNameGenerator("knife"))

	// fmt.Println("--------- BubblesortOnce -----------")
	// fmt.Println(BubblesortOnce([]int{9, 7, 5, 3, 1, 2, 4, 6, 8}))

	fmt.Println("--------- IsPalindrome -----------")
	fmt.Println(IsPalindrome("a"))

	fmt.Println("--------- Contamination -----------")
	fmt.Println(Contamination("abc", "z"))

	// fmt.Println("--------- FAKE BIN -----------")
	// fmt.Println(FakeBin("45385593107843568"))

	// fmt.Println("--------- Game Another card game -----------")
	// fmt.Println(Game([4]int{2, 5, 8, 11}, [4]int{1, 4, 7, 10}, [4]int{0, 3, 6, 9}))

	fmt.Println("--------- Game Another card game -----------")
	fmt.Println(SequenceSum(2, 6, 2))

	fmt.Println("--------- RowSumOddNumbers -----------")
	fmt.Println(RowSumOddNumbers(2))

	fmt.Println("--------- CalculateYears -----------")
	fmt.Println(CalculateYears(1))

	fmt.Println("--------- CATS -----------")
	fmt.Println(Cats(1, 5))
	fmt.Println(Cats(2, 5))

	fmt.Println("--------- FindShort  -----------")
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps"))

	fmt.Println("--------- FreqSeq  -----------")
	fmt.Println(FreqSeq("hello world", "-"))

	fmt.Println("--------- solutionCompareString -----------")
	fmt.Println(solutionCompareString("abc", "ab"))

	fmt.Println("--------- HighAndLow -----------")
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))

	fmt.Println("--------- HasUniqueChar -----------")
	fmt.Println(HasUniqueChar("  nAa"))
	fmt.Println(HasUniqueChar("++-"))

	fmt.Println("--------- ModifyMultiply -----------")
	fmt.Println(ModifyMultiply("qrUgYOzRwwuysX LLHGbHJVCTtZQjwyNMXWc HwrhGfjcCrjgXrjOMSEyHZ dhNx cmkhxIQ PuKzcufAJHcmJ GYEpax hjQmQOCbmH uR odVTMQOCMtQIdtnQBgMxOr vFE cnNvTWVxAZggqUzzAwVRpKP us nXpnNT Ko sKAIjnoarnmYohe dkYjPyhrWYOsuhCDLqR", 2, 0))
	fmt.Println(ModifyMultiply("This is a string", 3, 5))

	fmt.Println("--------- ContainAllRots -----------")
	fmt.Println(ContainAllRots("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}))

	fmt.Println("--------- ContainAllRots -----------")
	fmt.Println(Spacify("hello word"))

	fmt.Println("--------- Calc -----------")
	// fmt.Println(Calc("abcdef"))
	// fmt.Println("---------")
	fmt.Println(Calc("Jaam"))

	fmt.Println("--------- NthEven -----------")
	fmt.Println(NthEven(3))

	fmt.Println("--------- Permutations -----------")
	fmt.Println(Permutations("abc"))

	fmt.Println("--------- Permutations -----------")
	fmt.Println(Permutations("abc"))

	fmt.Println("--------- Find Zero Max-----------")
	fmt.Println(findZeroBinnary(9))
	fmt.Println("--------- SimpleRemoveDuplicates-----------")
	fmt.Println(SimpleRemoveDuplicates([]int{3, 4, 4, 3, 6, 3}))

	fmt.Println("--------- minOnArray-----------")
	fmt.Println(minOnArray([]int{3, 4, 4, 3, 6, 3}))

	fmt.Println("--------- LengthAndTwoValues-----------")
	fmt.Println(LengthAndTwoValues(5, "true", "false"))

	fmt.Println("--------- ComposingSquaredStrings-----------")
	fmt.Println(ComposingSquaredStrings("byGt\nhTts\nRTFF\nCnnI", "jIRl\nViBu\nrWOb\nNkTB"))
}
