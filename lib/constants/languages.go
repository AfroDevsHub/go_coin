package constants

import "fmt"

type Language int

const (
	AFRIKAANS Language = iota
	ALBANIAN
	AMHARIC
	ARABIC
	ARMENIAN
	ASSAMESE
	AZERBAIJANI
	BENGALI
	BOSNIAN
	BULGARIAN
	BURMESE
	CATALAN
	CHINESE_MANDARIN
	CROATIAN
	CZECH
	DANISH
	DUTCH
	ENGLISH
	ESTONIAN
	FINNISH
	FRENCH
	GEORGIAN
	GERMAN
	GREEK
	GUJARATI
	HEBREW
	HINDI
	HUNGARIAN
	ICELANDIC
	INDONESIAN
	IRISH
	ITALIAN
	JAPANESE
	JAVANESE
	KANNADA
	KAZAKH
	KHMER
	KOREAN
	KURDISH
	KYRGYZ
	LAO
	LATVIAN
	LITHUANIAN
	MACEDONIAN
	MALAY
	MALAYALAM
	MARATHI
	MONGOLIAN
	NEPALI
	NORWEGIAN
	ORIYA
	PASHTO
	PERSIAN
	POLISH
	PORTUGUESE
	PUNJABI
	ROMANIAN
	RUSSIAN
	SERBIAN
	SINHALA
	SLOVAK
	SLOVENIAN
	SOMALI
	SPANISH
	SUNDANESE
	SWAHILI
	SWEDISH
	TAGALOG
	TAJIK
	TAMIL
	TATAR
	TELUGU
	THAI
	TIBETAN
	TURKISH
	TURKMEN
	UKRAINIAN
	URDU
	UZBEK
	VIETNAMESE
	WU_CHINESE
	XHOSA
	YIDDISH
	YORUBA
	ZULU
)

type TupleLanguage struct {
	Language string
	Symbol   string
}

var languages = map[Language]TupleLanguage{
	AFRIKAANS:        {"Afrikaans", "af"},
	ALBANIAN:         {"Albanian", "sq"},
	AMHARIC:          {"Amharic", "am"},
	ARABIC:           {"Arabic", "ar"},
	ARMENIAN:         {"Armenian", "hy"},
	ASSAMESE:         {"Assamese", "as"},
	AZERBAIJANI:      {"Azerbaijani", "az"},
	BENGALI:          {"Bengali", "bn"},
	BOSNIAN:          {"Bosnian", "bs"},
	BULGARIAN:        {"Bulgarian", "bg"},
	BURMESE:          {"Burmese", "my"},
	CATALAN:          {"Catalan", "ca"},
	CHINESE_MANDARIN: {"Chinese (Mandarin)", "zh"},
	CROATIAN:         {"Croatian", "hr"},
	CZECH:            {"Czech", "cs"},
	DANISH:           {"Danish", "da"},
	DUTCH:            {"Dutch", "nl"},
	ENGLISH:          {"English", "en"},
	ESTONIAN:         {"Estonian", "et"},
	FINNISH:          {"Finnish", "fi"},
	FRENCH:           {"French", "fr"},
	GEORGIAN:         {"Georgian", "ka"},
	GERMAN:           {"German", "de"},
	GREEK:            {"Greek", "el"},
	GUJARATI:         {"Gujarati", "gu"},
	HEBREW:           {"Hebrew", "he"},
	HINDI:            {"Hindi", "hi"},
	HUNGARIAN:        {"Hungarian", "hu"},
	ICELANDIC:        {"Icelandic", "is"},
	INDONESIAN:       {"Indonesian", "id"},
	IRISH:            {"Irish", "ga"},
	ITALIAN:          {"Italian", "it"},
	JAPANESE:         {"Japanese", "ja"},
	JAVANESE:         {"Javanese", "jv"},
	KANNADA:          {"Kannada", "kn"},
	KAZAKH:           {"Kazakh", "kk"},
	KHMER:            {"Khmer", "km"},
	KOREAN:           {"Korean", "ko"},
	KURDISH:          {"Kurdish", "ku"},
	KYRGYZ:           {"Kyrgyz", "ky"},
	LAO:              {"Lao", "lo"},
	LATVIAN:          {"Latvian", "lv"},
	LITHUANIAN:       {"Lithuanian", "lt"},
	MACEDONIAN:       {"Macedonian", "mk"},
	MALAY:            {"Malay", "ms"},
	MALAYALAM:        {"Malayalam", "ml"},
	MARATHI:          {"Marathi", "mr"},
	MONGOLIAN:        {"Mongolian", "mn"},
	NEPALI:           {"Nepali", "ne"},
	NORWEGIAN:        {"Norwegian", "no"},
	ORIYA:            {"Oriya", "or"},
	PASHTO:           {"Pashto", "ps"},
	PERSIAN:          {"Persian", "fa"},
	POLISH:           {"Polish", "pl"},
	PORTUGUESE:       {"Portuguese", "pt"},
	PUNJABI:          {"Punjabi", "pa"},
	ROMANIAN:         {"Romanian", "ro"},
	RUSSIAN:          {"Russian", "ru"},
	SERBIAN:          {"Serbian", "sr"},
	SINHALA:          {"Sinhala", "si"},
	SLOVAK:           {"Slovak", "sk"},
	SLOVENIAN:        {"Slovenian", "sl"},
	SOMALI:           {"Somali", "so"},
	SPANISH:          {"Spanish", "es"},
	SUNDANESE:        {"Sundanese", "su"},
	SWAHILI:          {"Swahili", "sw"},
	SWEDISH:          {"Swedish", "sv"},
	TAGALOG:          {"Tagalog", "tl"},
	TAJIK:            {"Tajik", "tg"},
	TAMIL:            {"Tamil", "ta"},
	TATAR:            {"Tatar", "tt"},
	TELUGU:           {"Telugu", "te"},
	THAI:             {"Thai", "th"},
	TIBETAN:          {"Tibetan", "bo"},
	TURKISH:          {"Turkish", "tr"},
	TURKMEN:          {"Turkmen", "tk"},
	UKRAINIAN:        {"Ukrainian", "uk"},
	URDU:             {"Urdu", "ur"},
	UZBEK:            {"Uzbek", "uz"},
	VIETNAMESE:       {"Vietnamese", "vi"},
	WU_CHINESE:       {"Wu Chinese", "wuu"},
	XHOSA:            {"Xhosa", "xh"},
	YIDDISH:          {"Yiddish", "yi"},
	YORUBA:           {"Yoruba", "yo"},
	ZULU:             {"Zulu", "zu"},
}

func (l Language) String() (TupleLanguage, error) {
	if value, ok := languages[l]; ok {
		return value, nil
	}
	return TupleLanguage{}, fmt.Errorf("invalid login method")
}
