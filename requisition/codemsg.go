package requisition

type Language string
const (
	LangeageZh Language = "zh_CN"
	LanguageEn Language = "en_US"
)

type codeMsg struct {
	language Language
	codeMap  map[Language]map[int]string
}

var _codemsg = &codeMsg{
	codeMap: make(map[Language]map[int]string),
}

func init() {
	en := make(map[int]string)
	_codemsg.codeMap[LanguageEn] = en
	cn := make(map[int]string)
	_codemsg.codeMap[LangeageZh] = cn
}

func AddMsgLanguage(lang Language)  {
	_codemsg.codeMap[lang] = make(map[int]string)
}

func SetMsgMap(lang Language, msgMap map[int]string) {
	if codemsg, ok := _codemsg.codeMap[lang]; ok {
		for k, v := range msgMap {
			codemsg[k] = v
		}
	} else {
		_codemsg.codeMap[lang] = msgMap
	}
}

func SetMsg(lang Language, code int, msg string) {
	if codemsg, ok := _codemsg.codeMap[lang]; ok {
		codemsg[code] = msg
	} else {
		AddMsgLanguage(lang)
		_codemsg.codeMap[lang][code] = msg
	}
}

func Msg(lang Language, code int) string {
	if codemsg, ok := _codemsg.codeMap[lang]; ok {
		if msg, ok := codemsg[code]; ok {
			return msg
		}
	}
	return ""
}