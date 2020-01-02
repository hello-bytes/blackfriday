package blackfriday

import (
	"bytes"
	"log"
)

type Filter struct {
	// block-level callbacks
	/*BlockCode(out *bytes.Buffer, text []byte, infoString string)
	BlockQuote(out *bytes.Buffer, text []byte)
	BlockHtml(out *bytes.Buffer, text []byte)
	Header(out *bytes.Buffer, text func() bool, level int, id string)
	HRule(out *bytes.Buffer)
	List(out *bytes.Buffer, text func() bool, flags int)
	ListItem(out *bytes.Buffer, text []byte, flags int)
	Paragraph(out *bytes.Buffer, text func() bool)
	Table(out *bytes.Buffer, header []byte, body []byte, columnData []int)
	TableRow(out *bytes.Buffer, text []byte)
	TableHeaderCell(out *bytes.Buffer, text []byte, flags int)
	TableCell(out *bytes.Buffer, text []byte, flags int)
	Footnotes(out *bytes.Buffer, text func() bool)
	FootnoteItem(out *bytes.Buffer, name, text []byte, flags int)
	TitleBlock(out *bytes.Buffer, text []byte)

	// Span-level callbacks
	AutoLink(out *bytes.Buffer, link []byte, kind int)
	CodeSpan(out *bytes.Buffer, text []byte)
	DoubleEmphasis(out *bytes.Buffer, text []byte)
	Emphasis(out *bytes.Buffer, text []byte)
	Image(out *bytes.Buffer, link []byte, title []byte, alt []byte)
	LineBreak(out *bytes.Buffer)
	Link(out *bytes.Buffer, link []byte, title []byte, content []byte)
	RawHtmlTag(out *bytes.Buffer, tag []byte)
	TripleEmphasis(out *bytes.Buffer, text []byte)
	StrikeThrough(out *bytes.Buffer, text []byte)
	FootnoteRef(out *bytes.Buffer, ref []byte, id int)

	// Low-level callbacks
	Entity(out *bytes.Buffer, entity []byte)
	NormalText(out *bytes.Buffer, text []byte)

	// Header and footer
	DocumentHeader(out *bytes.Buffer)
	DocumentFooter(out *bytes.Buffer)

	GetFlags() int*/
}

func FilterRenderer() Renderer {
	log.Println("sdf")
	return &Filter{}
}

func (options *Filter) BlockCode(out *bytes.Buffer, text []byte, infoString string) {

}

func (options *Filter) BlockQuote(out *bytes.Buffer, text []byte) {

}
func (options *Filter) BlockHtml(out *bytes.Buffer, text []byte) {

}
func (options *Filter) Header(out *bytes.Buffer, text func() bool, level int, id string) {

}
func (options *Filter) HRule(out *bytes.Buffer) {

}
func (options *Filter) List(out *bytes.Buffer, text func() bool, flags int) {

}
func (options *Filter) ListItem(out *bytes.Buffer, text []byte, flags int) {

}
func (options *Filter) Paragraph(out *bytes.Buffer, text func() bool) {

}
func (options *Filter) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {

}
func (options *Filter) TableRow(out *bytes.Buffer, text []byte) {

}
func (options *Filter) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {

}
func (options *Filter) TableCell(out *bytes.Buffer, text []byte, flags int) {

}
func (options *Filter) Footnotes(out *bytes.Buffer, text func() bool) {

}
func (options *Filter) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {

}
func (options *Filter) TitleBlock(out *bytes.Buffer, text []byte) {

}

func (options *Filter) AutoLink(out *bytes.Buffer, link []byte, kind int) {

}
func (options *Filter) CodeSpan(out *bytes.Buffer, text []byte) {

}
func (options *Filter) DoubleEmphasis(out *bytes.Buffer, text []byte) {

}
func (options *Filter) Emphasis(out *bytes.Buffer, text []byte) {

}
func (options *Filter) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	log.Println(string(link))
}
func (options *Filter) LineBreak(out *bytes.Buffer) {

}
func (options *Filter) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {

}
func (options *Filter) RawHtmlTag(out *bytes.Buffer, tag []byte) {

}
func (options *Filter) TripleEmphasis(out *bytes.Buffer, text []byte) {

}
func (options *Filter) StrikeThrough(out *bytes.Buffer, text []byte) {

}
func (options *Filter) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {

}

func (options *Filter) Entity(out *bytes.Buffer, entity []byte) {

}
func (options *Filter) NormalText(out *bytes.Buffer, text []byte) {
	log.Println("NormalText ", string(text))
}

func (options *Filter) DocumentHeader(out *bytes.Buffer) {
	log.Println("document header")
}
func (options *Filter) DocumentFooter(out *bytes.Buffer) {
	log.Println("DocumentFooter")
}

func (options *Filter) GetFlags() int {
	log.Println("get flats")
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
