package sensitive_word_filter

import "strings"

const (
	minMatchTYpe = 1 // 最小匹配规则
	maxMatchType = 2 // 最大匹配规则
)

var sensitiveWordMap map[string]interface{} = InitKeyWord()

/**
 * 截取字符串,从起始位置截取指定长度
 * @param str
 * @param beginIndex
 * @param length
 */
func SubString(str string, beginIndex, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if beginIndex < 0 {
		beginIndex = 0
	}
	if beginIndex >= lth {
		beginIndex = lth
	}
	end := beginIndex + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[beginIndex:end])
}

/**
 * 检查文字中是否包含敏感字符，检查规则如下：<br>
 * 如果存在，则返回敏感词字符的长度，不存在返回0
 * @param txt
 * @param beginIndex
 * @param matchType
 * @return
 */
func CheckSensitiveWord(txt string, beginIndex int, matchType int) int {
	//匹配标识数默认为0
	var matchFlag int = 0
	//敏感词结束标志位 用于敏感词只有1位的情况
	var flag bool = false
	var nowMap map[string]interface{} = sensitiveWordMap
	rs := []rune(txt)
myforLable:
	for i := beginIndex; i < len(rs); i++ {
		var word string = string(rs[i])
		if nil != nowMap[word] {
			//获取指定key
			nowMap = nowMap[word].(map[string]interface{})
			//存在，则判断是否为最后一个
			if nil != nowMap {
				//				fmt.Println("存在=", word)
				//找到相应key，匹配标识+1
				matchFlag++

				//如果为最后一个匹配规则，结束循环，返回匹配标识数
				if nil != nowMap["isEnd"] && "1" == nowMap["isEnd"].(string) {
					//结束标识为为true
					flag = true
					//最小规则，直接返回，最大规则则还需继续查找
					if minMatchTYpe == matchType {
						break myforLable
					}
				}
			} else { //不存在，直接返回
				break myforLable
			}
		} else {//不存在，直接返回
			break myforLable
		}
	}
	//长度必须大于等于1，为词
	if matchFlag < 2 || !flag {
		matchFlag = 0
	}
	return matchFlag
}

/**
 * 获取文字中的敏感词
 * @param txt
 * @param matchType
 * @return
 */
func getSensitiveWord(txt string, matchType int) []string {
	sensitiveWordList := make([]string, 0, 10)
	rs := []rune(txt)
	for i := 0; i < len(rs); i++ {
		//判断是否包含敏感字符
		var length int = CheckSensitiveWord(txt, i, matchType)
		if length > 0 {
			sensitiveWordList = append(sensitiveWordList, SubString(txt, i, length))
			//减1的原因，是因为for会自增
			i = i + length - 1
		}
	}
	//	fmt.Println("敏感词=", sensitiveWordList)
	return sensitiveWordList
}

/**
 * 获取替换字符串
 * @param replaceChar
 * @param length
 * @return
 */
func getReplaceChars(replaceChar string, length int) string {
	var resultReplace = replaceChar
	for i := 1; i < length; i++ {
		resultReplace += replaceChar
	}
	return resultReplace
}

/**
 * 替换敏感字字符
 * @param txt
 * @param matchType
 * @param replaceChar
 * @return
 */
func ReplaceSensitiveWord(txt string, matchType int, replaceChar string) string {
	var resultTxt string = txt
	//获取所有敏感词
	set := getSensitiveWord(txt, matchType)
	var replaceString string
	for _, word := range set {
		if "" == word {
			continue
		}
		replaceString = getReplaceChars(replaceChar, len([]rune(word)))
		//		word = "法轮功"
		resultTxt = strings.Replace(resultTxt, word, replaceString, -1)
	}
	return resultTxt
}
