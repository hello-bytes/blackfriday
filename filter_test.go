package blackfriday

import "testing"

func TestFilterData(t *testing.T) {
	var input = "Hleoo \n ![https://oss-cn-hangzhou.aliyuncs.com/codingsky/cdn/img/2020-01-01/2486514945123c1f10d8429851fb5f4d.png](Title) \n test"

	renderer := FilterRenderer()
	MarkdownOptions([]byte(input), renderer, Options{
		Extensions: commonExtensions})

	t.Error(nil)
}
