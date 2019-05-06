// Package util is a set of utilitary things.
package util

import "encoding/xml"

// Document documento
type Document struct {
	XMLName   xml.Name `xml:"document"`
	Text      string   `xml:",chardata"`
	O         string   `xml:"o,attr"`
	R         string   `xml:"r,attr"`
	V         string   `xml:"v,attr"`
	W         string   `xml:"w,attr"`
	W10       string   `xml:"w10,attr"`
	Wp        string   `xml:"wp,attr"`
	Wps       string   `xml:"wps,attr"`
	Wpg       string   `xml:"wpg,attr"`
	Mc        string   `xml:"mc,attr"`
	Wp14      string   `xml:"wp14,attr"`
	W14       string   `xml:"w14,attr"`
	Ignorable string   `xml:"Ignorable,attr"`
	Body      struct {
		Text string `xml:",chardata"`
		P    *[]struct {
			Text string `xml:",chardata"`
			PPr  struct {
				Text   string `xml:",chardata"`
				PStyle struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"pStyle"`
				Spacing struct {
					Text   string `xml:",chardata"`
					Before string `xml:"before,attr"`
					After  string `xml:"after,attr"`
				} `xml:"spacing"`
				RPr struct {
					Text string `xml:",chardata"`
					Lang struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"lang"`
				} `xml:"rPr"`
			} `xml:"pPr"`
			R *struct {
				Text string `xml:",chardata"`
				RPr  struct {
					Text string `xml:",chardata"`
					Lang struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"lang"`
				} `xml:"rPr"`
				T *string `xml:"t"`
			} `xml:"r"`
		} `xml:"p"`
		Tbl struct {
			Text  string `xml:",chardata"`
			TblPr struct {
				Text string `xml:",chardata"`
				TblW struct {
					Text string `xml:",chardata"`
					W    string `xml:"w,attr"`
					Type string `xml:"type,attr"`
				} `xml:"tblW"`
				Jc struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"jc"`
				TblInd struct {
					Text string `xml:",chardata"`
					W    string `xml:"w,attr"`
					Type string `xml:"type,attr"`
				} `xml:"tblInd"`
				TblBorders struct {
					Text string `xml:",chardata"`
					Top  struct {
						Text  string `xml:",chardata"`
						Val   string `xml:"val,attr"`
						Sz    string `xml:"sz,attr"`
						Space string `xml:"space,attr"`
						Color string `xml:"color,attr"`
					} `xml:"top"`
					Left struct {
						Text  string `xml:",chardata"`
						Val   string `xml:"val,attr"`
						Sz    string `xml:"sz,attr"`
						Space string `xml:"space,attr"`
						Color string `xml:"color,attr"`
					} `xml:"left"`
					Bottom struct {
						Text  string `xml:",chardata"`
						Val   string `xml:"val,attr"`
						Sz    string `xml:"sz,attr"`
						Space string `xml:"space,attr"`
						Color string `xml:"color,attr"`
					} `xml:"bottom"`
					InsideH struct {
						Text  string `xml:",chardata"`
						Val   string `xml:"val,attr"`
						Sz    string `xml:"sz,attr"`
						Space string `xml:"space,attr"`
						Color string `xml:"color,attr"`
					} `xml:"insideH"`
				} `xml:"tblBorders"`
				TblCellMar struct {
					Text string `xml:",chardata"`
					Top  struct {
						Text string `xml:",chardata"`
						W    string `xml:"w,attr"`
						Type string `xml:"type,attr"`
					} `xml:"top"`
					Left struct {
						Text string `xml:",chardata"`
						W    string `xml:"w,attr"`
						Type string `xml:"type,attr"`
					} `xml:"left"`
					Bottom struct {
						Text string `xml:",chardata"`
						W    string `xml:"w,attr"`
						Type string `xml:"type,attr"`
					} `xml:"bottom"`
					Right struct {
						Text string `xml:",chardata"`
						W    string `xml:"w,attr"`
						Type string `xml:"type,attr"`
					} `xml:"right"`
				} `xml:"tblCellMar"`
			} `xml:"tblPr"`
			TblGrid struct {
				Text    string `xml:",chardata"`
				GridCol []struct {
					Text string `xml:",chardata"`
					W    string `xml:"w,attr"`
				} `xml:"gridCol"`
			} `xml:"tblGrid"`
			Tr []struct {
				Text string `xml:",chardata"`
				TrPr string `xml:"trPr"`
				Tc   []struct {
					Text string `xml:",chardata"`
					TcPr struct {
						Text string `xml:",chardata"`
						TcW  struct {
							Text string `xml:",chardata"`
							W    string `xml:"w,attr"`
							Type string `xml:"type,attr"`
						} `xml:"tcW"`
						TcBorders struct {
							Text string `xml:",chardata"`
							Top  struct {
								Text  string `xml:",chardata"`
								Val   string `xml:"val,attr"`
								Sz    string `xml:"sz,attr"`
								Space string `xml:"space,attr"`
								Color string `xml:"color,attr"`
							} `xml:"top"`
							Left struct {
								Text  string `xml:",chardata"`
								Val   string `xml:"val,attr"`
								Sz    string `xml:"sz,attr"`
								Space string `xml:"space,attr"`
								Color string `xml:"color,attr"`
							} `xml:"left"`
							Bottom struct {
								Text  string `xml:",chardata"`
								Val   string `xml:"val,attr"`
								Sz    string `xml:"sz,attr"`
								Space string `xml:"space,attr"`
								Color string `xml:"color,attr"`
							} `xml:"bottom"`
							InsideH struct {
								Text  string `xml:",chardata"`
								Val   string `xml:"val,attr"`
								Sz    string `xml:"sz,attr"`
								Space string `xml:"space,attr"`
								Color string `xml:"color,attr"`
							} `xml:"insideH"`
							Right struct {
								Text  string `xml:",chardata"`
								Val   string `xml:"val,attr"`
								Sz    string `xml:"sz,attr"`
								Space string `xml:"space,attr"`
								Color string `xml:"color,attr"`
							} `xml:"right"`
							InsideV struct {
								Text  string `xml:",chardata"`
								Val   string `xml:"val,attr"`
								Sz    string `xml:"sz,attr"`
								Space string `xml:"space,attr"`
								Color string `xml:"color,attr"`
							} `xml:"insideV"`
						} `xml:"tcBorders"`
						Shd struct {
							Text string `xml:",chardata"`
							Fill string `xml:"fill,attr"`
							Val  string `xml:"val,attr"`
						} `xml:"shd"`
					} `xml:"tcPr"`
					P struct {
						Text string `xml:",chardata"`
						PPr  struct {
							Text   string `xml:",chardata"`
							PStyle struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"pStyle"`
							Jc struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"jc"`
							RPr struct {
								Text string   `xml:",chardata"`
								B    []string `xml:"b"`
								BCs  string   `xml:"bCs"`
								Sz   struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"sz"`
								SzCs struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"szCs"`
								Color struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"color"`
							} `xml:"rPr"`
						} `xml:"pPr"`
						R struct {
							Text string `xml:",chardata"`
							RPr  struct {
								Text string `xml:",chardata"`
								B    string `xml:"b"`
								BCs  string `xml:"bCs"`
								Sz   struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"sz"`
								SzCs struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"szCs"`
								Color struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"color"`
								RFonts struct {
									Text  string `xml:",chardata"`
									Ascii string `xml:"ascii,attr"`
									HAnsi string `xml:"hAnsi,attr"`
								} `xml:"rFonts"`
							} `xml:"rPr"`
							T string `xml:"t"`
						} `xml:"r"`
						BookmarkStart struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Name string `xml:"name,attr"`
						} `xml:"bookmarkStart"`
						BookmarkEnd struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"bookmarkEnd"`
					} `xml:"p"`
				} `xml:"tc"`
			} `xml:"tr"`
		} `xml:"tbl"`
		SectPr struct {
			Text string `xml:",chardata"`
			Type struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"type"`
			PgSz struct {
				Text string `xml:",chardata"`
				W    string `xml:"w,attr"`
				H    string `xml:"h,attr"`
			} `xml:"pgSz"`
			PgMar struct {
				Text   string `xml:",chardata"`
				Left   string `xml:"left,attr"`
				Right  string `xml:"right,attr"`
				Header string `xml:"header,attr"`
				Top    string `xml:"top,attr"`
				Footer string `xml:"footer,attr"`
				Bottom string `xml:"bottom,attr"`
				Gutter string `xml:"gutter,attr"`
			} `xml:"pgMar"`
			PgNumType struct {
				Text string `xml:",chardata"`
				Fmt  string `xml:"fmt,attr"`
			} `xml:"pgNumType"`
			FormProt struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"formProt"`
			TextDirection struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"textDirection"`
			DocGrid struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				LinePitch string `xml:"linePitch,attr"`
				CharSpace string `xml:"charSpace,attr"`
			} `xml:"docGrid"`
		} `xml:"sectPr"`
	} `xml:"body"`
}
