package main

import (
	"fmt"
	"time"
)

const (
	Aquarius    = "Aquarius"    // 水瓶座
	Pisces      = "Pisces"      // 双鱼座
	Aries       = "Aries"       // 白羊座
	Taurus      = "Taurus"      // 金牛座
	Gemini      = "Gemini"      // 双子座
	Cancer      = "Cancer"      // 巨蟹座
	Leo         = "Leo"         // 狮子座
	Virgo       = "Virgo"       // 处女座
	Libra       = "Libra"       // 天秤座
	Scorpio     = "Scorpio"     // 天蝎座
	Sagittarius = "Sagittarius" // 射手座
	Capricorn   = "Capricorn"   // 摩羯座

	// 中文翻译常量
	ZhAquarius    = "水瓶座"
	ZhPisces      = "双鱼座"
	ZhAries       = "白羊座"
	ZhTaurus      = "金牛座"
	ZhGemini      = "双子座"
	ZhCancer      = "巨蟹座"
	ZhLeo         = "狮子座"
	ZhVirgo       = "处女座"
	ZhLibra       = "天秤座"
	ZhScorpio     = "天蝎座"
	ZhSagittarius = "射手座"
	ZhCapricorn   = "摩羯座"
)

var (
	zodiacSigns = map[string]string{
		Aquarius:    ZhAquarius,
		Pisces:      ZhPisces,
		Aries:       ZhAries,
		Taurus:      ZhTaurus,
		Gemini:      ZhGemini,
		Cancer:      ZhCancer,
		Leo:         ZhLeo,
		Virgo:       ZhVirgo,
		Libra:       ZhLibra,
		Scorpio:     ZhScorpio,
		Sagittarius: ZhSagittarius,
		Capricorn:   ZhCapricorn,
	}

	translationMap = map[string]string{
		ZhAquarius:    Aquarius,
		ZhPisces:      Pisces,
		ZhAries:       Aries,
		ZhTaurus:      Taurus,
		ZhGemini:      Gemini,
		ZhCancer:      Cancer,
		ZhLeo:         Leo,
		ZhVirgo:       Virgo,
		ZhLibra:       Libra,
		ZhScorpio:     Scorpio,
		ZhSagittarius: Sagittarius,
		ZhCapricorn:   Capricorn,
	}
)

func getZodiacSign(birthDate time.Time) string {
	month := birthDate.Month()
	day := birthDate.Day()

	if (month == time.January && day >= 20) || (month == time.February && day <= 18) {
		return Aquarius
	} else if (month == time.February && day >= 19) || (month == time.March && day <= 20) {
		return Pisces
	} else if (month == time.March && day >= 21) || (month == time.April && day <= 19) {
		return Aries
	} else if (month == time.April && day >= 20) || (month == time.May && day <= 20) {
		return Taurus
	} else if (month == time.May && day >= 21) || (month == time.June && day <= 21) {
		return Gemini
	} else if (month == time.June && day >= 22) || (month == time.July && day <= 22) {
		return Cancer
	} else if (month == time.July && day >= 23) || (month == time.August && day <= 22) {
		return Leo
	} else if (month == time.August && day >= 23) || (month == time.September && day <= 22) {
		return Virgo
	} else if (month == time.September && day >= 23) || (month == time.October && day <= 22) {
		return Libra
	} else if (month == time.October && day >= 23) || (month == time.November && day <= 21) {
		return Scorpio
	} else if (month == time.November && day >= 22) || (month == time.December && day <= 21) {
		return Sagittarius
	} else {
		return Capricorn
	}
}

func getZodiacSignTranslation(zodiac string) string {
	translation, ok := zodiacSigns[zodiac]
	if !ok {
		return ""
	}
	return translation
}

func getZodiacSignByTranslation(translation string) string {
	zodiac, ok := translationMap[translation]
	if !ok {
		return ""
	}
	return zodiac
}

func main() {
	birthDate := time.Date(1990, time.May, 1, 0, 0, 0, 0, time.UTC)
	zodiac := getZodiacSign(birthDate)
	translation := getZodiacSignTranslation(zodiac)
	fmt.Println("Zodiac:", zodiac)
	fmt.Println("Translation:", translation)

	// 通过中文翻译获取英文标识
	zodiacByTranslation := getZodiacSignByTranslation(translation)
	fmt.Println("Zodiac by Translation:", zodiacByTranslation)
}
