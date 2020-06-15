package requisition

const (
	LanguageZh string = "zh_CN"
	LanguageEn string = "en_US"
)

type codeMsg struct {
	language string
	codeMap  map[string]map[int]string
}

var _codemsg = &codeMsg{
	codeMap: make(map[string]map[int]string),
}

func init() {
	en := make(map[int]string)
	_codemsg.codeMap[LanguageEn] = en
	cn := make(map[int]string)
	_codemsg.codeMap[LanguageZh] = cn
}

func AddMsgLanguage(lang string)  {
	_codemsg.codeMap[lang] = make(map[int]string)
}

func SupportLanguage(lang string) bool {
	if _, ok := _codemsg.codeMap[lang]; ok == true {
		return true
	}
	return false
}

func SetMsgMap(lang string, msgMap map[int]string) {
	if codemsg, ok := _codemsg.codeMap[lang]; ok {
		for k, v := range msgMap {
			codemsg[k] = v
		}
	} else {
		_codemsg.codeMap[lang] = msgMap
	}
}

func SetMsg(lang string, code int, msg string) {
	if codemsg, ok := _codemsg.codeMap[lang]; ok {
		codemsg[code] = msg
	} else {
		AddMsgLanguage(lang)
		_codemsg.codeMap[lang][code] = msg
	}
}

func Msg(lang string, code int) string {
	if codemsg, ok := _codemsg.codeMap[lang]; ok {
		if msg, ok := codemsg[code]; ok {
			return msg
		}
	}
	return ""
}