/**
 *电话号码的字母组合
 *
 * 解题思路：回溯
 * 
 */

var phoneMap map[string]string = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}
var results []string
func letterCombinations(digits string) []string {
    if digits == ""{
        return []string{}
    }
    results = []string{}
    allCode(digits, 0, "")
    return results
}

func allCode(digits string, index int, res string) {
    if index == len(digits) {
        results = append(results, res)
    }else{
        digit := string(digits[index])
        str := phoneMap[digit]
        for i:= 0;i<len(str); i++{
            allCode(digits, index+1, res+string(str[i]))
        }
    }
}