package tnef

var codepages = map[uint32]string{
	37:    "ibm037",       // ibm ebcdic us-canada
	437:   "ibm437",       // oem united states
	500:   "ibm500",       // ibm ebcdic international
	708:   "asmo-708",     // arabic (asmo 708)
	720:   "dos-720",      // arabic (transparent asmo); arabic (dos)
	737:   "ibm737",       // oem greek (formerly 437g); greek (dos)
	775:   "ibm775",       // oem baltic; baltic (dos)
	850:   "ibm850",       // oem multilingual latin 1; western european (dos)
	852:   "ibm852",       // oem latin 2; central european (dos)
	855:   "ibm855",       // oem cyrillic (primarily russian)
	857:   "ibm857",       // oem turkish; turkish (dos)
	858:   "ibm00858",     // oem multilingual latin 1 + euro symbol
	860:   "ibm860",       // oem portuguese; portuguese (dos)
	861:   "ibm861",       // oem icelandic; icelandic (dos)
	862:   "dos-862",      // oem hebrew; hebrew (dos)
	863:   "ibm863",       // oem french canadian; french canadian (dos)
	864:   "ibm864",       // oem arabic; arabic (864)
	865:   "ibm865",       // oem nordic; nordic (dos)
	866:   "cp866",        // oem russian; cyrillic (dos)
	869:   "ibm869",       // oem modern greek; greek, modern (dos)
	870:   "ibm870",       // ibm ebcdic multilingual/roece (latin 2); ibm ebcdic multilingual latin 2
	874:   "windows-874",  // ansi/oem thai (iso 8859-11); thai (windows)
	875:   "cp875",        // ibm ebcdic greek modern
	932:   "shift_jis",    // ansi/oem japanese; japanese (shift-jis)
	936:   "gb2312",       // ansi/oem simplified chinese (prc, singapore); chinese simplified (gb2312)
	950:   "big5",         // ansi/oem traditional chinese (taiwan; hong kong sar, prc); chinese traditional (big5)
	1026:  "ibm1026",      // ibm ebcdic turkish (latin 5)
	1047:  "ibm01047",     // ibm ebcdic latin 1/open system
	1140:  "ibm01140",     // ibm ebcdic us-canada (037 + euro symbol); ibm ebcdic (us-canada-euro)
	1141:  "ibm01141",     // ibm ebcdic germany (20273 + euro symbol); ibm ebcdic (germany-euro)
	1142:  "ibm01142",     // ibm ebcdic denmark-norway (20277 + euro symbol); ibm ebcdic (denmark-norway-euro)
	1143:  "ibm01143",     // ibm ebcdic finland-sweden (20278 + euro symbol); ibm ebcdic (finland-sweden-euro)
	1144:  "ibm01144",     // ibm ebcdic italy (20280 + euro symbol); ibm ebcdic (italy-euro)
	1145:  "ibm01145",     // ibm ebcdic latin america-spain (20284 + euro symbol); ibm ebcdic (spain-euro)
	1146:  "ibm01146",     // ibm ebcdic united kingdom (20285 + euro symbol); ibm ebcdic (uk-euro)
	1147:  "ibm01147",     // ibm ebcdic france (20297 + euro symbol); ibm ebcdic (france-euro)
	1148:  "ibm01148",     // ibm ebcdic international (500 + euro symbol); ibm ebcdic (international-euro)
	1149:  "ibm01149",     // ibm ebcdic icelandic (20871 + euro symbol); ibm ebcdic (icelandic-euro)
	1200:  "utf-16",       // unicode utf-16, little endian byte order (bmp of iso 10646); available only to managed applications
	1201:  "utf-16be",     // unicode utf-16, big endian byte order; available only to managed applications
	1250:  "windows-1250", // ansi central european; central european (windows)
	1251:  "windows-1251", // ansi cyrillic; cyrillic (windows)
	1252:  "windows-1252", // ansi latin 1; western european (windows)
	1253:  "windows-1253", // ansi greek; greek (windows)
	1254:  "windows-1254", // ansi turkish; turkish (windows)
	1255:  "windows-1255", // ansi hebrew; hebrew (windows)
	1256:  "windows-1256", // ansi arabic; arabic (windows)
	1257:  "windows-1257", // ansi baltic; baltic (windows)
	1258:  "windows-1258", // ansi/oem vietnamese; vietnamese (windows)
	10000: "macintosh",    // mac roman; western european (mac)
	12000: "utf-32",       // unicode utf-32, little endian byte order; available only to managed applications
	12001: "utf-32be",     // unicode utf-32, big endian byte order; available only to managed applications
	20127: "us-ascii",     // us-ascii (7-bit)
	20273: "ibm273",       // ibm ebcdic germany
	20277: "ibm277",       // ibm ebcdic denmark-norway
	20278: "ibm278",       // ibm ebcdic finland-sweden
	20280: "ibm280",       // ibm ebcdic italy
	20284: "ibm284",       // ibm ebcdic latin america-spain
	20285: "ibm285",       // ibm ebcdic united kingdom
	20290: "ibm290",       // ibm ebcdic japanese katakana extended
	20297: "ibm297",       // ibm ebcdic france
	20420: "ibm420",       // ibm ebcdic arabic
	20423: "ibm423",       // ibm ebcdic greek
	20424: "ibm424",       // ibm ebcdic hebrew
	20838: "ibm-thai",     // ibm ebcdic thai
	20866: "koi8-r",       // russian (koi8-r); cyrillic (koi8-r)
	20871: "ibm871",       // ibm ebcdic icelandic
	20880: "ibm880",       // ibm ebcdic cyrillic russian
	20905: "ibm905",       // ibm ebcdic turkish
	20924: "ibm00924",     // ibm ebcdic latin 1/open system (1047 + euro symbol)
	20932: "euc-jp",       // japanese (jis 0208-1990 and 0212-1990)
	20936: "cp20936",      // simplified chinese (gb2312); chinese simplified (gb2312-80)
	20949: "cp20949",      // korean wansung
	21025: "cp1025",       // ibm ebcdic cyrillic serbian-bulgarian
	21866: "koi8-u",       // ukrainian (koi8-u); cyrillic (koi8-u)
	28591: "iso-8859-1",   // iso 8859-1 latin 1; western european (iso)
	28592: "iso-8859-2",   // iso 8859-2 central european; central european (iso)
	28593: "iso-8859-3",   // iso 8859-3 latin 3
	28594: "iso-8859-4",   // iso 8859-4 baltic
	28595: "iso-8859-5",   // iso 8859-5 cyrillic
	28596: "iso-8859-6",   // iso 8859-6 arabic
	28597: "iso-8859-7",   // iso 8859-7 greek
	28598: "iso-8859-8",   // iso 8859-8 hebrew; hebrew (iso-visual)
	28599: "iso-8859-9",   // iso 8859-9 turkish
	28603: "iso-8859-13",  // iso 8859-13 estonian
	28605: "iso-8859-15",  // iso 8859-15 latin 9
	38598: "iso-8859-8-i", // iso 8859-8 hebrew; hebrew (iso-logical)
	50220: "iso-2022-jp",  // iso 2022 japanese with no halfwidth katakana; japanese (jis)
	50221: "csiso2022jp",  // iso 2022 japanese with halfwidth katakana; japanese (jis-allow 1 byte kana)
	50222: "iso-2022-jp",  // iso 2022 japanese jis x 0201-1989; japanese (jis-allow 1 byte kana - so/si)
	50225: "iso-2022-kr",  // iso 2022 korean
	51932: "euc-jp",       // euc japanese
	51936: "euc-cn",       // euc simplified chinese; chinese simplified (euc)
	51949: "euc-kr",       // euc korean
	52936: "hz-gb-2312",   // hz-gb2312 simplified chinese; chinese simplified (hz)
	54936: "gb18030",      // windows xp and later: gb18030 simplified chinese (4 byte); chinese simplified (gb18030)
	65000: "utf-7",
	65001: "utf-8",
}

func GetCharset(code uint32) string {
	return codepages[code]
}
