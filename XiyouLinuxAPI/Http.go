package xylinux

//HTTPParam POST Body的一一映射值
type HTTPParam map[string]string

//GenerateParamStream 从param生成一个"a=va&b=vb"格式的串
func GenerateParamStream(param HTTPParam) []byte {
	var ret string
	for k, v := range param {
		ret += k + "=" + v + "&"
	}
	return []byte(ret[:len(ret)-1])
}
