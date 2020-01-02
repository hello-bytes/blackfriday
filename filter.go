package blackfriday

import (
	"bytes"
)

const IMAGEITEM int8 = 1

//MakedownItem markdown的元素
type MakedownItem struct {
	ItemType int8
	Data     interface{}
}

//ImageItem markdown的图片
type ImageItem struct {
	Title string
	URL   string
}

type Filter struct {
	items []MakedownItem
}

func FilterRenderer() Renderer {
	return &Filter{}
}

func (options *Filter) BlockCode(out *bytes.Buffer, text []byte, infoString string) {
	// log.Println("BlockCode ")
}

func (options *Filter) BlockQuote(out *bytes.Buffer, text []byte) {
	// log.Println("BlockQuote ")
}
func (options *Filter) BlockHtml(out *bytes.Buffer, text []byte) {
	// log.Println("BlockHtml ")
}
func (options *Filter) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	// log.Println("Header ")
	text()
}
func (options *Filter) HRule(out *bytes.Buffer) {
	// log.Println("HRule ")
}
func (options *Filter) List(out *bytes.Buffer, text func() bool, flags int) {
	// log.Println("List ")
	text()
}
func (options *Filter) ListItem(out *bytes.Buffer, text []byte, flags int) {
	// log.Println("ListItem ")
}
func (options *Filter) Paragraph(out *bytes.Buffer, text func() bool) {
	// log.Println("Paragraph ")
	text()
}
func (options *Filter) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	// log.Println("Table ")
}
func (options *Filter) TableRow(out *bytes.Buffer, text []byte) {
	// log.Println("TableRow ")
}
func (options *Filter) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {
	// log.Println("TableHeaderCell ")
}
func (options *Filter) TableCell(out *bytes.Buffer, text []byte, flags int) {
	// log.Println("TableCell ")
}
func (options *Filter) Footnotes(out *bytes.Buffer, text func() bool) {
	// log.Println("Footnotes ")
}
func (options *Filter) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	// log.Println("FootnoteItem ")
}
func (options *Filter) TitleBlock(out *bytes.Buffer, text []byte) {
	// log.Println("TitleBlock ")
}

func (options *Filter) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	// log.Println("AutoLink ")
}
func (options *Filter) CodeSpan(out *bytes.Buffer, text []byte) {
	// log.Println("CodeSpan ")
}
func (options *Filter) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	// log.Println("DoubleEmphasis ")
}
func (options *Filter) Emphasis(out *bytes.Buffer, text []byte) {
	// log.Println("Emphasis ")
}
func (options *Filter) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	// log.Println("Image : ", string(link))

	var item MakedownItem
	item.ItemType = IMAGEITEM

	var imageItem ImageItem
	imageItem.Title = string(title)
	imageItem.URL = string(link)
	item.Data = &imageItem
	options.items = append(options.items, item)
}
func (options *Filter) LineBreak(out *bytes.Buffer) {
	// log.Println("NormalText ")
}
func (options *Filter) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	// log.Println("Link ")
}
func (options *Filter) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	// log.Println("RawHtmlTag ")
}
func (options *Filter) TripleEmphasis(out *bytes.Buffer, text []byte) {
	// log.Println("TripleEmphasis ")
}
func (options *Filter) StrikeThrough(out *bytes.Buffer, text []byte) {
	// log.Println("StrikeThrough ")
}
func (options *Filter) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	// log.Println("FootnoteRef ")
}

func (options *Filter) Entity(out *bytes.Buffer, entity []byte) {
	// log.Println("Entity ")
}
func (options *Filter) NormalText(out *bytes.Buffer, text []byte) {
	// log.Println("NormalText ", string(text))
}

func (options *Filter) DocumentHeader(out *bytes.Buffer) {
	// log.Println("document header")
}
func (options *Filter) DocumentFooter(out *bytes.Buffer) {
	// log.Println("DocumentFooter")
}

func (options *Filter) GetFlags() int {
	return 0 |
		HTML_USE_XHTML |
		HTML_USE_SMARTYPANTS |
		HTML_SMARTYPANTS_FRACTIONS |
		HTML_SMARTYPANTS_DASHES |
		HTML_SMARTYPANTS_LATEX_DASHES | //-------
		EXTENSION_NO_INTRA_EMPHASIS |
		EXTENSION_TABLES |
		EXTENSION_FENCED_CODE |
		EXTENSION_AUTOLINK |
		EXTENSION_STRIKETHROUGH |
		EXTENSION_SPACE_HEADERS |
		EXTENSION_HEADER_IDS |
		EXTENSION_BACKSLASH_LINE_BREAK |
		EXTENSION_DEFINITION_LISTS
}

func AnalyzeMarkdown(text string) ([]MakedownItem, error) {
	var filterObj Filter
	//renderer := FilterRenderer()
	MarkdownOptions([]byte(text), &filterObj, Options{
		Extensions: commonExtensions})
	return filterObj.items, nil
}
