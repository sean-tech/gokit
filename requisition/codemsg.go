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

var CodeMsg = &codeMsg{
	codeMap: make(map[Language]map[int]string),
}

func init() {
	en := make(map[int]string)
	CodeMsg.codeMap[LanguageEn] = en
	cn := make(map[int]string)
	CodeMsg.codeMap[LangeageZh] = cn
}

func (this *codeMsg) LanguageAdd(lang Language)  {
	this.codeMap[lang] = make(map[int]string)
}

func (this *codeMsg) SetMsgMap(lang Language, msgMap map[int]string) {
	if codemsg, ok := this.codeMap[lang]; ok {
		for k, v := range msgMap {
			codemsg[k] = v
		}
	} else {
		this.codeMap[lang] = msgMap
	}
}

func (this *codeMsg) SetMsg(lang Language, code int, msg string) {
	if codemsg, ok := this.codeMap[lang]; ok {
		codemsg[code] = msg
	} else {
		this.LanguageAdd(lang)
		this.codeMap[lang][code] = msg
	}
}

func (this *codeMsg) Msg(lang Language, code int) string {
	if codemsg, ok := this.codeMap[lang]; ok {
		if msg, ok := codemsg[code]; ok {
			return msg
		}
	}
	return ""
}