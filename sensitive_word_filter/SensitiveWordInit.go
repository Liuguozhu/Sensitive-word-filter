package sensitive_word_filter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/**
初始化敏感词库<br>
将敏感词加入到HashMap中<br>
构建DFA算法模型
@author LGZ
**/

/**
 * 初始化敏感字库
 * @return
 */
func InitKeyWord() map[string]interface{} {
	wordSet := readSensitiveWordFile()
	wordMap := addSensitiveWordToHashMap(wordSet)
	return wordMap
}

/**
 * 读取敏感词库中的内容，将内容添加到set集合中
 * @return
 * @throws Exception
 */
func readSensitiveWordFile() []string {
	set := make([]string, 0, 10)
	app, _ := exec.LookPath(os.Args[0])
	path := filepath.Dir(app) //获取当前程序执行路径
	//	println("当前文件执行路径：" + path)
	file := fmt.Sprintf("%s\\conf\\censor_words.txt", path)
	read, err := os.Open(file) //打开文件
	defer read.Close()         //打开文件出错处理
	if nil == err {
		buff := bufio.NewReader(read)
		for {
			txt, err := buff.ReadString('\n') //以'\n'为结束符读入一行
			if nil != err || io.EOF == err {
				break
			}
			//将读取到的一行文件添加到数组中
			// 去除空格
			txt = strings.Replace(txt, "\t", "", -1)
			// 去除换行符
			txt = strings.Replace(txt, "\n", "", -1)
			//			if  " " == txt {
			//				fmt.Println("有换行啊！！！")
			//			}
			if "" != txt && "\n" != txt {
				set = append(set, txt)
			}
		}
	} else {
		fmt.Println("err=", err)
	}
	return set
}

/**
 * 读取敏感词库，将敏感词放入HashSet中，构建一个DFA算法模型：<br>
 * 中 = { isEnd = 0 国 = {<br>
 * isEnd = 1 人 = {isEnd = 0 民 = {isEnd = 1} } 男 = { isEnd = 0 人 = { isEnd =
 * 1 } } } } 五 = { isEnd = 0 星 = { isEnd = 0 红 = { isEnd = 0 旗 = { isEnd = 1
 * } } } }
 */
func addSensitiveWordToHashMap(wordSet []string) map[string]interface{} {
	var wordMap = make(map[string]interface{}, len(wordSet))
	for _, word := range wordSet {
		//		fmt.Printf("wordSet[%d]=%s", index, word)
		nowMap := wordMap
		rs := []rune(word)
		for i := 0; i < len(rs)-1; i++ {
			keyChar := string(rs[i])

			//			fmt.Printf("keyChar[%d]=%s \n", i, keyChar)
			//获取
			tempMap := nowMap[keyChar]
			//如果存在该key，直接赋值
			if nil != tempMap {
				nowMap = tempMap.(map[string]interface{})
			} else {
				//设置标志位
				newMap := make(map[string]interface{}, len(rs))
				newMap["isEnd"] = "0"
				//添加到集合
				nowMap[keyChar] = newMap
				nowMap = newMap
			}
			//最后一个
			if i == len(rs)-2 {
				nowMap["isEnd"] = "1"
			}
		}
	}
	return wordMap
}
