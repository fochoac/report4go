package main

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Document -
type Document struct {
	Scheme     map[string]string
	SkipScheme string
	Body       Body `xml:"body"`
}

// DocItemType -
type DocItemType int

// Body -
type Body struct {
	Items  []DocItem
	Params BodyParams `xml:"sectPr"`
}

// DocItem -
type DocItem interface {
	Tag() string
	Type() DocItemType
	PlainText() string
	Clone() DocItem
	decode(decoder *xml.Decoder) error
	encode(encoder *xml.Encoder) error
}

// BodyParams -
type BodyParams struct {
	HeaderReference *ReferenceValue `xml:"headerReference,omitempty"`
	FooterReference *ReferenceValue `xml:"footerReference,omitempty"`
	PageSize        SizeValue       `xml:"pgSz"`
	PageMargin      MarginValue     `xml:"pgMar"`
	Bidi            IntValue        `xml:"bidi"`
}

// SimpleDocxFile - файл docx
type SimpleDocxFile struct {
	zipFile  *zip.ReadCloser
	headers  map[string]*Header
	document *Document
}

func main() {
	documento, err := OpenFile("empleados.docx")
	if err != nil {
		fmt.Println("Error leyendo el archivo")
	}
	for _, item := range documento.document.Body.Items {
		fmt.Println(item.Type())
		switch elem := item.(type) {
		case *ParagraphItem:
			{
				for _, item := range elem.Items {
					fmt.Println("pi - " + item.PlainText())
				}

			}
		// Запись
		case *RecordItem:
			{
				fmt.Println("ri - " + elem.Text)

			}

		case *TableItem:
			{
				fmt.Println("tb - " + elem.PlainText())
			}
		}
	}

}

func decoderElement(decoder *xml.Decoder) {
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch element := token.(type) {
		case xml.StartElement:
			{
				printAttr(element)

				decoderElement(decoder)

			}
		case xml.EndElement:
			{
				if element.Name.Local == "body" {
					break
				}
			}
		}
	}
}

func getDocument(read *zip.ReadCloser) (f *zip.File) {
	for _, f := range read.File {
		if f.Name == "word/document.xml" {
			return f
		}
	}
	return nil
}

func printAttr(element xml.StartElement) {
	for _, attr := range element.Attr {
		fmt.Println(attr.Name.Local)
	}
}

// HeightValue - значение высоты
type HeightValue struct {
	Value      int64  `xml:"val,attr"`
	HeightRule string `xml:"hRule,attr,omitempty"`
}

// From (HeightValue)
func (h *HeightValue) From(h1 *HeightValue) {
	if h1 != nil {
		h.HeightRule = h1.HeightRule
		h.Value = h1.Value
	}
}

// WidthValue - значение длины
type WidthValue struct {
	Value int64  `xml:"w,attr"`
	Type  string `xml:"type,attr,omitempty"`
}

// From (WidthValue)
func (w *WidthValue) From(w1 *WidthValue) {
	if w1 != nil {
		w.Type = w1.Type
		w.Value = w1.Value
	}
}

// SizeValue - значение размера
type SizeValue struct {
	Width       int64  `xml:"w,attr"`
	Height      int64  `xml:"h,attr"`
	Orientation string `xml:"orient,attr,omitempty"`
}

// From (SizeValue)
func (s *SizeValue) From(s1 *SizeValue) {
	if s1 != nil {
		s.Height = s1.Height
		s.Orientation = s1.Orientation
		s.Width = s1.Width
	}
}

// EmptyValue - пустое значение
type EmptyValue struct {
}

// StringValue - одиночное string значение
type StringValue struct {
	Value string `xml:"val,attr,omitempty"`
}

// From (StringValue)
func (s *StringValue) From(s1 *StringValue) {
	if s1 != nil {
		s.Value = s1.Value
	}
}

// BoolValue - одиночное bool значение
type BoolValue struct {
	Value bool `xml:"val,attr"`
}

// IntValue - одиночное int значение
type IntValue struct {
	Value int64 `xml:"val,attr"`
}

// From (IntValue)
func (i *IntValue) From(i1 *IntValue) {
	if i1 != nil {
		i.Value = i1.Value
	}
}

// FloatValue - одиночное float значение
type FloatValue struct {
	Value float64 `xml:"val,attr"`
}

// ReferenceValue - reference value
type ReferenceValue struct {
	Type string `xml:"type,attr"`
	ID   string `xml:"id,attr"`
}

// SpacingValue - spacing value
type SpacingValue struct {
	After    int64  `xml:"after,attr"`
	Before   int64  `xml:"before,attr"`
	Line     int64  `xml:"line,attr"`
	LineRule string `xml:"lineRule,attr"`
}

// From (SpacingValue)
func (s *SpacingValue) From(s1 *SpacingValue) {
	if s1 != nil {
		s.After = s1.After
		s.Before = s1.Before
		s.Line = s1.Line
		s.LineRule = s1.LineRule
	}
}

// MarginValue - margin значение
type MarginValue struct {
	Top    int64 `xml:"top,attr"`
	Left   int64 `xml:"left,attr"`
	Bottom int64 `xml:"bottom,attr"`
	Right  int64 `xml:"right,attr"`
	Header int64 `xml:"header,attr,omitempty"`
	Footer int64 `xml:"footer,attr,omitempty"`
}

// From (MarginValue)
func (m *MarginValue) From(m1 *MarginValue) {
	if m1 != nil {
		m.Top = m1.Top
		m.Left = m1.Left
		m.Bottom = m1.Bottom
		m.Right = m1.Right
		m.Header = m1.Header
		m.Footer = m1.Footer
	}
}

// Margins - margins значение
type Margins struct {
	Top    WidthValue `xml:"top"`
	Left   WidthValue `xml:"left"`
	Bottom WidthValue `xml:"bottom"`
	Right  WidthValue `xml:"right"`
}

// From (Margins)
func (m *Margins) From(m1 *Margins) {
	if m1 != nil {
		m.Top.From(&m1.Top)
		m.Left.From(&m1.Left)
		m.Bottom.From(&m1.Bottom)
		m.Right.From(&m1.Right)
	}
}

// ShadowValue - значение тени
type ShadowValue struct {
	Value string `xml:"val,attr"`
	Color string `xml:"color,attr"`
	Fill  string `xml:"fill,attr"`
}

// From (ShadowValue)
func (s *ShadowValue) From(s1 *ShadowValue) {
	if s1 != nil {
		s.Value = s1.Value
		s.Color = s1.Color
		s.Fill = s1.Fill
	}
}

// Decode (Document) - декодирование документа
func (doc *Document) Decode(reader io.Reader) error {
	decoder := xml.NewDecoder(reader)
	if decoder != nil {
		doc.Scheme = make(map[string]string)
		for {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "document" {
						for _, attr := range element.Attr {
							if attr.Name.Local == "Ignorable" {
								doc.SkipScheme = attr.Value
							} else {
								doc.Scheme[attr.Name.Local] = attr.Value
							}
						}
					} else if element.Name.Local == "body" {
						err := doc.Body.decode(decoder)
						if err != nil {
							return err
						}
					}
				}
			}
		}
		return nil
	}
	return errors.New("Error create decoder")
}

// Декодирование BODY
func (body *Body) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		if body.Items == nil {
			body.Items = make([]DocItem, 0)
		}
		for {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "sectPr" {
						decoder.DecodeElement(&body.Params, &element)
					} else {
						// Декодирование элементов
						item := decodeItem(&element, decoder)
						if item != nil {
							body.Items = append(body.Items, item)
						}
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "body" {
						break
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

func decodeItem(element *xml.StartElement, decoder *xml.Decoder) DocItem {
	if element != nil && decoder != nil {
		var item DocItem
		if element.Name.Local == "p" {
			item = new(ParagraphItem)
		} else if element.Name.Local == "r" {
			item = new(RecordItem)
		} else if element.Name.Local == "tbl" {
			item = new(TableItem)
		}
		if item != nil {
			if item.decode(decoder) == nil {
				return item
			}
		}
	}
	return nil
}

// ParagraphItem - параграф
type ParagraphItem struct {
	Params ParagraphParams `xml:"pPr"`
	Items  []DocItem
}

// ParagraphParams - параметры параграфа
type ParagraphParams struct {
	Style   *StringValue  `xml:"pStyle,omitempty"`
	Spacing *SpacingValue `xml:"spacing,omitempty"`
	Jc      *StringValue  `xml:"jc,omitempty"`
	Bidi    *IntValue     `xml:"bidi,omitempty"`
}

// Tag - имя тега элемента
func (item *ParagraphItem) Tag() string {
	return "p"
}

// Paragraph - параграф
const (
	Paragraph DocItemType = iota
	Record
	Table
)

// Type - тип элемента
func (item *ParagraphItem) Type() DocItemType {
	return Paragraph
}

// PlainText - текст
func (item *ParagraphItem) PlainText() string {
	var result string
	for _, i := range item.Items {
		tmp := i.PlainText()
		if len(tmp) > 0 {
			result += tmp
		}
	}
	return result
}

// Clone - клонирование
func (item *ParagraphItem) Clone() DocItem {
	result := new(ParagraphItem)
	result.Items = make([]DocItem, 0)
	for _, i := range item.Items {
		if i != nil {
			result.Items = append(result.Items, i.Clone())
		}
	}
	// Клонирование параметров
	if item.Params.Bidi != nil {
		result.Params.Bidi = new(IntValue)
		result.Params.Bidi.Value = item.Params.Bidi.Value
	}
	if item.Params.Jc != nil {
		result.Params.Jc = new(StringValue)
		result.Params.Jc.Value = item.Params.Jc.Value
	}
	if item.Params.Spacing != nil {
		result.Params.Spacing = new(SpacingValue)
		result.Params.Spacing.After = item.Params.Spacing.After
		result.Params.Spacing.Before = item.Params.Spacing.Before
		result.Params.Spacing.Line = item.Params.Spacing.Line
		result.Params.Spacing.LineRule = item.Params.Spacing.LineRule
	}
	if item.Params.Style != nil {
		result.Params.Style = new(StringValue)
		result.Params.Style.Value = item.Params.Style.Value
	}
	return result
}

// Декодирование параграфа
func (item *ParagraphItem) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		var end bool
		for !end {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "pPr" {
						decoder.DecodeElement(&item.Params, &element)
					} else {
						i := decodeItem(&element, decoder)
						if i != nil {
							item.Items = append(item.Items, i)
						}
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "p" {
						end = true
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

/* КОДИРОВАНИЕ */

// Кодирование параграфа
func (item *ParagraphItem) encode(encoder *xml.Encoder) error {
	if encoder != nil {
		// Начало параграфа
		start := xml.StartElement{Name: xml.Name{Local: item.Tag()}}
		if err := encoder.EncodeToken(start); err != nil {
			return err
		}
		// Параметры параграфа
		if err := encoder.EncodeElement(&item.Params, xml.StartElement{Name: xml.Name{Local: "pPr"}}); err != nil {
			return err
		}
		// Кодируем составные элементы
		for _, i := range item.Items {
			if err := i.encode(encoder); err != nil {
				return err
			}
		}
		// Конец параграфа
		if err := encoder.EncodeToken(start.End()); err != nil {
			return err
		}
		return encoder.Flush()
	}
	return errors.New("Not have encoder")
}

// RecordItem - record item
type RecordItem struct {
	Params RecordParams `xml:"rPr,omitempty"`
	Text   string       `xml:"t,omitempty"`
	Tab    bool         `xml:"tab,omitempty"`
	Break  bool         `xml:"br,omitempty"`
}

// RecordParams - params record
type RecordParams struct {
	Fonts     *RecordFonts `xml:"rFonts,omitempty"`
	Rtl       *IntValue    `xml:"rtl,omitempty"`
	Size      *IntValue    `xml:"sz,omitempty"`
	SizeCs    *IntValue    `xml:"szCs,omitempty"`
	Lang      *StringValue `xml:"lang,omitempty"`
	Underline *StringValue `xml:"u,omitempty"`
	Italic    *EmptyValue  `xml:"i,omitempty"`
	Bold      *EmptyValue  `xml:"b,omitempty"`
	BoldCS    *EmptyValue  `xml:"bCs,omitempty"`
	Color     *StringValue `xml:"color,omitempty"`
}

// RecordFonts - fonts in record
type RecordFonts struct {
	ASCII      string `xml:"ascii,attr"`
	CS         string `xml:"cs,attr"`
	HandleANSI string `xml:"hAnsi,attr"`
	EastAsia   string `xml:"eastAsia,attr"`
	HandleInt  string `xml:"hint,attr,omitempty"`
}

// Tag - имя тега элемента
func (item *RecordItem) Tag() string {
	return "r"
}

// Type - тип элемента
func (item *RecordItem) Type() DocItemType {
	return Record
}

// PlainText - текст
func (item *RecordItem) PlainText() string {
	return item.Text
}

// Clone - клонирование
func (item *RecordItem) Clone() DocItem {
	result := new(RecordItem)
	result.Text = item.Text
	result.Tab = item.Tab
	result.Break = item.Break
	// Клонируем параметры
	if item.Params.Bold != nil {
		result.Params.Bold = new(EmptyValue)
	}
	if item.Params.BoldCS != nil {
		result.Params.BoldCS = new(EmptyValue)
	}
	if item.Params.Italic != nil {
		result.Params.Italic = new(EmptyValue)
	}
	if item.Params.Underline != nil {
		result.Params.Underline = new(StringValue)
		result.Params.Underline.Value = item.Params.Underline.Value
	}
	if item.Params.Color != nil {
		result.Params.Color = new(StringValue)
		result.Params.Color.Value = item.Params.Color.Value
	}
	if item.Params.Lang != nil {
		result.Params.Lang = new(StringValue)
		result.Params.Lang.Value = item.Params.Lang.Value
	}
	if item.Params.Rtl != nil {
		result.Params.Rtl = new(IntValue)
		result.Params.Rtl.Value = item.Params.Rtl.Value
	}
	if item.Params.Size != nil {
		result.Params.Size = new(IntValue)
		result.Params.Size.Value = item.Params.Size.Value
	}
	if item.Params.SizeCs != nil {
		result.Params.SizeCs = new(IntValue)
		result.Params.SizeCs.Value = item.Params.SizeCs.Value
	}
	if item.Params.Fonts != nil {
		result.Params.Fonts = new(RecordFonts)
		result.Params.Fonts.ASCII = item.Params.Fonts.ASCII
		result.Params.Fonts.CS = item.Params.Fonts.CS
		result.Params.Fonts.EastAsia = item.Params.Fonts.EastAsia
		result.Params.Fonts.HandleANSI = item.Params.Fonts.HandleANSI
		result.Params.Fonts.HandleInt = item.Params.Fonts.HandleInt
	}
	return result
}

// Декодирование записи
func (item *RecordItem) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		var end bool
		for !end {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "rPr" {
						decoder.DecodeElement(&item.Params, &element)
					} else if element.Name.Local == "t" {
						decoder.DecodeElement(&item.Text, &element)
					} else if element.Name.Local == "br" {
						item.Break = true
					} else if element.Name.Local == "tab" {
						item.Tab = true
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "r" {
						end = true
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

/* КОДИРОВАНИЕ */

// Кодирование записи
func (item *RecordItem) encode(encoder *xml.Encoder) error {
	if encoder != nil {
		// Начало записи
		start := xml.StartElement{Name: xml.Name{Local: item.Tag()}}
		if err := encoder.EncodeToken(start); err != nil {
			return err
		}
		// Параметры записи
		if err := encoder.EncodeElement(&item.Params, xml.StartElement{Name: xml.Name{Local: "rPr"}}); err != nil {
			return err
		}
		// Текст
		if err := encoder.EncodeElement(&item.Text, xml.StartElement{Name: xml.Name{Local: "t"}}); err != nil {
			return err
		}
		// <br />
		if item.Break {
			startBr := xml.StartElement{Name: xml.Name{Local: "br"}}
			if err := encoder.EncodeToken(startBr); err != nil {
				return err
			}
			if err := encoder.EncodeToken(startBr.End()); err != nil {
				return err
			}
			if err := encoder.Flush(); err != nil {
				return err
			}
		}
		// Tab
		if item.Tab {
			startTab := xml.StartElement{Name: xml.Name{Local: "tab"}}
			if err := encoder.EncodeToken(startTab); err != nil {
				return err
			}
			if err := encoder.EncodeToken(startTab.End()); err != nil {
				return err
			}
			if err := encoder.Flush(); err != nil {
				return err
			}
		}
		// Конец записи
		if err := encoder.EncodeToken(start.End()); err != nil {
			return err
		}
		return encoder.Flush()
	}
	return errors.New("Not have encoder")
}

// TableItem - элемент таблици
type TableItem struct {
	Params TableParams `xml:"tblPr"`
	Grid   TableGrid   `xml:"tblGrid"`
	Rows   []*TableRow `xml:"tr,omitempty"`
}

// TableGrid - Grid table
type TableGrid struct {
	Cols []*WidthValue `xml:"gridCol,omitempty"`
}

// TableParamsEx - Other params table
type TableParamsEx struct {
	Shadow ShadowValue `xml:"shd"`
}

// Tag - имя тега элемента
func (item *TableItem) Tag() string {
	return "tbl"
}

// PlainText - текст
func (item *TableItem) PlainText() string {
	return ""
}

// Type - тип элемента
func (item *TableItem) Type() DocItemType {
	return Table
}

// Clone - клонирование
func (item *TableItem) Clone() DocItem {
	result := new(TableItem)
	result.Grid.Cols = make([]*WidthValue, 0)
	for _, col := range item.Grid.Cols {
		if col != nil {
			w := new(WidthValue)
			w.Type = col.Type
			w.Value = col.Value
			result.Grid.Cols = append(result.Grid.Cols, w)
		}
	}
	if item.Params.DocGrid != nil {
		result.Params.DocGrid = new(IntValue)
		result.Params.DocGrid.Value = item.Params.DocGrid.Value
	}
	if item.Params.Ind != nil {
		result.Params.Ind = new(WidthValue)
		result.Params.Ind.Type = item.Params.Ind.Type
		result.Params.Ind.Value = item.Params.Ind.Value
	}
	if item.Params.Jc != nil {
		result.Params.Jc = new(StringValue)
		result.Params.Jc.Value = item.Params.Jc.Value
	}
	if item.Params.Layout != nil {
		result.Params.Layout = new(TableLayout)
		result.Params.Layout.Type = item.Params.Layout.Type
	}
	if item.Params.Shadow != nil {
		result.Params.Shadow = new(ShadowValue)
		result.Params.Shadow.From(item.Params.Shadow)
	}
	if item.Params.Width != nil {
		result.Params.Width = new(WidthValue)
		result.Params.Width.From(item.Params.Width)
	}
	if item.Params.Borders != nil {
		result.Params.Borders = new(TableBorders)
		result.Params.Borders.From(item.Params.Borders)
	}
	// Клонирование строк
	result.Rows = make([]*TableRow, 0)
	for _, row := range item.Rows {
		if row != nil {
			result.Rows = append(result.Rows, row.Clone())
		}
	}
	return result
}

// TableParams - Params table
type TableParams struct {
	Width   *WidthValue   `xml:"tblW,omitempty"`
	Jc      *StringValue  `xml:"jc,omitempty"`
	Ind     *WidthValue   `xml:"tblInd,omitempty"`
	Borders *TableBorders `xml:"tblBorders,omitempty"`
	Shadow  *ShadowValue  `xml:"shd,omitempty"`
	Layout  *TableLayout  `xml:"tblLayout,omitempty"`
	DocGrid *IntValue     `xml:"docGrid,omitempty"`
}

// TableLayout - layout params
type TableLayout struct {
	Type string `xml:"type,attr"`
}

// TableBorders in table
type TableBorders struct {
	Top     TableBorder  `xml:"top"`
	Left    TableBorder  `xml:"left"`
	Bottom  TableBorder  `xml:"bottom"`
	Right   TableBorder  `xml:"right"`
	InsideH *TableBorder `xml:"insideH,omitempty"`
	InsideV *TableBorder `xml:"insideV,omitempty"`
}

// From (TableBorders)
func (b *TableBorders) From(b1 *TableBorders) {
	if b1 != nil {
		b.Top.From(&b1.Top)
		b.Left.From(&b1.Left)
		b.Bottom.From(&b1.Bottom)
		b.Right.From(&b1.Right)
		if b1.InsideH != nil {
			b.InsideH = new(TableBorder)
			b.InsideH.From(b1.InsideH)
		}
		if b1.InsideV != nil {
			b.InsideV = new(TableBorder)
			b.InsideV.From(b1.InsideV)
		}
	}
}

// TableBorder in borders
type TableBorder struct {
	Value  string `xml:"val,attr"`
	Color  string `xml:"color,attr"`
	Size   int64  `xml:"sz,attr"`
	Space  int64  `xml:"space,attr"`
	Shadow int64  `xml:"shadow,attr"`
	Frame  int64  `xml:"frame,attr"`
}

// From (TableBorder)
func (b *TableBorder) From(b1 *TableBorder) {
	if b1 != nil {
		b.Value = b1.Value
		b.Color = b1.Color
		b.Frame = b1.Frame
		b.Shadow = b1.Shadow
		b.Size = b1.Size
		b.Space = b1.Space
	}
}

// TableRow - row in table
type TableRow struct {
	OtherParams *TableParamsEx `xml:"tblPrEx,omitempty"`
	Params      TableRowParams `xml:"trPr"`
	Cells       []*TableCell   `xml:"tc,omitempty"`
}

// TableRowParams - row params
type TableRowParams struct {
	Height   HeightValue `xml:"trHeight"`
	IsHeader bool
}

// TableCell - table cell
type TableCell struct {
	Params TableCellParams `xml:"tcPr"`
	Items  []DocItem
}

// TableCellParams - cell params
type TableCellParams struct {
	Width         *WidthValue   `xml:"tcW,omitempty"`
	Borders       *TableBorders `xml:"tcBorders,omitempty"`
	Shadow        *ShadowValue  `xml:"shd,omitempty"`
	Margins       *Margins      `xml:"tcMar,omitempty"`
	VerticalAlign *StringValue  `xml:"vAlign,omitempty"`
	VerticalMerge *StringValue  `xml:"vMerge,omitempty"`
	GridSpan      *IntValue     `xml:"gridSpan,omitempty"`
	HideMark      *EmptyValue   `xml:"hideMark,omitempty"`
	NoWrap        *EmptyValue   `xml:"noWrap,omitempty"`
}

// Clone (TableCell) - клонирование ячейки
func (cell *TableCell) Clone() *TableCell {
	result := new(TableCell)
	if cell.Params.GridSpan != nil {
		result.Params.GridSpan = new(IntValue)
		result.Params.GridSpan.Value = cell.Params.GridSpan.Value
	}
	if cell.Params.HideMark != nil {
		result.Params.HideMark = new(EmptyValue)
	}
	if cell.Params.NoWrap != nil {
		result.Params.NoWrap = new(EmptyValue)
	}
	if cell.Params.Shadow != nil {
		result.Params.Shadow = new(ShadowValue)
		result.Params.Shadow.From(cell.Params.Shadow)
	}
	if cell.Params.VerticalAlign != nil {
		result.Params.VerticalAlign = new(StringValue)
		result.Params.VerticalAlign.Value = cell.Params.VerticalAlign.Value
	}
	if cell.Params.VerticalMerge != nil {
		result.Params.VerticalMerge = new(StringValue)
		result.Params.VerticalMerge.Value = cell.Params.VerticalMerge.Value
	}
	if cell.Params.Margins != nil {
		result.Params.Margins = new(Margins)
		result.Params.Margins.From(cell.Params.Margins)
	}
	if cell.Params.Width != nil {
		result.Params.Width = new(WidthValue)
		result.Params.Width.From(cell.Params.Width)
	}
	if cell.Params.Borders != nil {
		result.Params.Borders = new(TableBorders)
		result.Params.Borders.From(cell.Params.Borders)
	}
	result.Items = make([]DocItem, 0)
	for _, item := range cell.Items {
		if item != nil {
			result.Items = append(result.Items, item.Clone())
		}
	}
	return result
}

// Clone (TableRow) - клонирование строки таблицы
func (row *TableRow) Clone() *TableRow {
	result := new(TableRow)
	result.Params = row.Params
	result.OtherParams = new(TableParamsEx)
	if row.OtherParams != nil {
		result.OtherParams.Shadow = row.OtherParams.Shadow
	}
	// Клонируем ячейки
	result.Cells = make([]*TableCell, 0)
	for _, cell := range row.Cells {
		if cell != nil {
			result.Cells = append(result.Cells, cell.Clone())
		}
	}
	return result
}

/* ДЕКОДИРОВАНИЕ */

// Декодирование таблицы
func (item *TableItem) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		item.Rows = make([]*TableRow, 0)
		var end bool
		for !end {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "tblPr" {
						decoder.DecodeElement(&item.Params, &element)
					} else if element.Name.Local == "tblGrid" {
						decoder.DecodeElement(&item.Grid, &element)
					} else if element.Name.Local == "tr" {
						row := new(TableRow)
						if row.decode(decoder) == nil {
							item.Rows = append(item.Rows, row)
						}
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "tbl" {
						end = true
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

// Декодирование строк таблицы
func (row *TableRow) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		row.Cells = make([]*TableCell, 0)
		var end bool
		for !end {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "trHeight" {
						decoder.DecodeElement(&row.Params.Height, &element)
					} else if element.Name.Local == "tblHeader" {
						row.Params.IsHeader = true
					} else if element.Name.Local == "tblPrEx" {
						row.OtherParams = new(TableParamsEx)
						decoder.DecodeElement(row.OtherParams, &element)
					} else if element.Name.Local == "tc" {
						cell := new(TableCell)
						if cell.decode(decoder) == nil {
							row.Cells = append(row.Cells, cell)
						}
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "tr" {
						end = true
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

// Декодирование ячеек таблицы
func (row *TableCell) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		var end bool
		for !end {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "tcPr" {
						decoder.DecodeElement(&row.Params, &element)
					} else {
						i := decodeItem(&element, decoder)
						if i != nil {
							row.Items = append(row.Items, i)
						}
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "tc" {
						end = true
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

/* КОДИРОВАНИЕ */

// Кодирование таблицы
func (item *TableItem) encode(encoder *xml.Encoder) error {
	if encoder != nil {
		// Начало таблицы
		start := xml.StartElement{Name: xml.Name{Local: item.Tag()}}
		if err := encoder.EncodeToken(start); err != nil {
			return err
		}
		// Параметры таблицы
		if err := encoder.EncodeElement(&item.Params, xml.StartElement{Name: xml.Name{Local: "tblPr"}}); err != nil {
			return err
		}
		// Сетка таблицы
		if err := encoder.EncodeElement(&item.Grid, xml.StartElement{Name: xml.Name{Local: "tblGrid"}}); err != nil {
			return err
		}
		// Строки таблицы
		for _, row := range item.Rows {
			if row != nil {
				if err := row.encode(encoder); err != nil {
					return err
				}
			}
		}
		// Конец таблицы
		if err := encoder.EncodeToken(start.End()); err != nil {
			return err
		}
		return encoder.Flush()
	}
	return errors.New("Not have encoder")
}

// Кодирование ячейки таблицы
func (cell *TableCell) encode(encoder *xml.Encoder) error {
	if encoder != nil {
		// Начало ячейки таблицы
		start := xml.StartElement{Name: xml.Name{Local: "tc"}}
		if err := encoder.EncodeToken(start); err != nil {
			return err
		}
		// Параметры ячейки таблицы
		if err := encoder.EncodeElement(&cell.Params, xml.StartElement{Name: xml.Name{Local: "tcPr"}}); err != nil {
			return err
		}
		// Кодируем составные элементы
		for _, i := range cell.Items {
			if err := i.encode(encoder); err != nil {
				return err
			}
		}
		// Конец ячейки таблицы
		if err := encoder.EncodeToken(start.End()); err != nil {
			return err
		}
		return encoder.Flush()
	}
	return errors.New("Not have encoder")
}

// Кодирование строки таблицы
func (row *TableRow) encode(encoder *xml.Encoder) error {
	if encoder != nil {
		// Начало строки таблицы
		start := xml.StartElement{Name: xml.Name{Local: "tr"}}
		if err := encoder.EncodeToken(start); err != nil {
			return err
		}
		// Параметры строки таблицы
		if row.OtherParams != nil {
			if err := encoder.EncodeElement(row.OtherParams, xml.StartElement{Name: xml.Name{Local: "tblPrEx"}}); err != nil {
				return err
			}
		}
		// Кодируем Параметры
		startPr := xml.StartElement{Name: xml.Name{Local: "trPr"}}
		if err := encoder.EncodeToken(startPr); err != nil {
			return err
		}
		if err := encoder.EncodeElement(&row.Params.Height, xml.StartElement{Name: xml.Name{Local: "trHeight"}}); err != nil {
			return err
		}
		if row.Params.IsHeader {
			startHeader := xml.StartElement{Name: xml.Name{Local: "tblHeader"}}
			if err := encoder.EncodeToken(startHeader); err != nil {
				return err
			}
			if err := encoder.EncodeToken(startHeader.End()); err != nil {
				return err
			}
			if err := encoder.Flush(); err != nil {
				return err
			}
		}
		if err := encoder.EncodeToken(startPr.End()); err != nil {
			return err
		}
		if err := encoder.Flush(); err != nil {
			return err
		}
		// Кодируем ячейки
		for _, cell := range row.Cells {
			if cell != nil {
				if err := cell.encode(encoder); err != nil {
					return err
				}
			}
		}
		// Конец строки таблицы
		if err := encoder.EncodeToken(start.End()); err != nil {
			return err
		}
		return encoder.Flush()
	}
	return errors.New("Not have encoder")
}

// Header - разметка заголовка DOCX
type Header struct {
	Scheme     map[string]string
	SkipScheme string
	Items      []DocItem
}

// OpenFile - Открытие файла DOCX
func OpenFile(fileName string) (*SimpleDocxFile, error) {
	z, err := zip.OpenReader(fileName)
	if err != nil {
		return nil, err
	}
	d := new(SimpleDocxFile)
	d.headers = make(map[string]*Header)
	d.zipFile = z
	// Перебор файлов в Zip архиве
	for _, f := range z.File {
		if f != nil {
			// Загрузка документа
			if f.Name == "word/document.xml" {
				reader, err := f.Open()
				if err != nil {
					return nil, err
				}
				d.document = new(Document)
				d.document.Decode(reader)
				if err := reader.Close(); err != nil {
					return nil, err
				}
			} else if strings.Index(f.Name, "word/header") >= 0 {
				/*  reader, err := f.Open()
				    if err != nil {
				        return nil, err
				    }
				    header := new(Header)
				    header.Decode(reader)
				    if err := reader.Close(); err != nil {
				        return nil, err
				    }
				    d.headers[f.Name] = header*/
			}
		}
	}
	return d, nil
}
