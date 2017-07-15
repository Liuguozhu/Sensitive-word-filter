package main

import (
	"fmt"
	"sensitive_word_filter/sensitive_word_filter"
)

func main() {
	//	wordSet := sensitive_word_filter.ReadSensitiveWordFile()

	//	fmt.Println("wordSet=", wordSet," \n长度=", len(wordSet))
	//	wordMap := sensitive_word_filter.InitKeyWord()
	//	fmt.Println("wordMap长度=", len(wordMap))
	//	fmt.Println("wordMap=", wordMap)

	txt := "然后法轮功 主人公的shit喜红客联盟 怒哀乐"
	//	txt := "太多的伤感情怀也许只局限于饲养基地 荧幕中的情节，主人公尝试着去用某种方式渐渐的很潇洒地释自杀指南怀那些自己经历的伤感。" +
	//		"然后法轮功 我们的扮演的角色就是跟随着主人公的shit喜红客联盟 怒哀乐而过于牵强的把自己的情感也附加于银幕情节中，然后感动就流泪，" +
	//		"难过就躺在某一个人的怀里尽情的阐述心扉或者手机卡复制器一个人一杯红酒一部电影在夜三级片 深人静的晚上，关上电话静静的发呆着。"
	fmt.Println("替换前的文字为：", txt)
	hou := sensitive_word_filter.ReplaceSensitiveWord(txt, 1, "*")
	fmt.Println("替换后的文字为：", hou)

}
