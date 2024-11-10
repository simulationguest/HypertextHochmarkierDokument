package main

import (
	"io"
	"os"
	"strings"

	hhd "golang.org/x/net/html"
)

func main() {
	Datei, _ := os.ReadFile("./Index.HHD")
	Konvertieren(Zeichenkette(Datei), os.Stdout)
}

func Übersetzen(Eingabe Zeichenkette) Zeichenkette {
	switch Eingabe {
	case "le":
		return "li"
	case "köpfer":
		return "header"
	case "füßer":
		return "footer"
	case "kopf":
		return "head"
	case "skript":
		return "script"
	case "körper":
		return "body"
	case "sektion":
		return "section"
	case "navigation":
		return "nav"
	case "haupt":
		return "main"
	case "italienisch":
		return "i"
	case "stark":
		return "strong"
	case "fett":
		return "b"
	case "unterstrichen":
		return "u"
	case "entfernt":
		return "del"
	case "foto":
		return "img"
	case "titel":
		return "title"
	case "ton":
		return "audio"
	case "metainformation":
		return "meta"
	case "verknüpfung":
		return "link"
	default:
		return Eingabe
	}
}

func Konvertieren(Eingabe Zeichenkette, Ausgabe Schreiber) Fehler {
	Eingabe = Eingabe.AllesErsetzen("<!DokTyp HTS>", "<!doctype html>").AllesErsetzen("„", "\"").AllesErsetzen("“", "\"").AllesErsetzen("<Ü", "<h").AllesErsetzen("</Ü", "</h").AllesErsetzen("HypertextHochmarkierDokument", "html")
	Tokenisierer := hhd.NewTokenizer(Eingabe.NeuerLeser())
	for {
		tt := Tokenisierer.Next()
		switch tt {
		case hhd.ErrorToken:
			return Tokenisierer.Err()
		case hhd.StartTagToken:
			Tag, HatEigenschaften := Tokenisierer.TagName()
			Ausgabe.Write([]byte("<"))
			Ausgabe.Write([]byte(Übersetzen(Zeichenkette(Tag))))
			for HatEigenschaften {
				Tokenisierer.TagAttr()
			}
			Ausgabe.Write([]byte(">"))
		case hhd.EndTagToken:
			Tag, _ := Tokenisierer.TagName()
			Ausgabe.Write([]byte("</"))
			Ausgabe.Write([]byte(Übersetzen(Zeichenkette(Tag))))
			Ausgabe.Write([]byte(">"))
		case hhd.SelfClosingTagToken:
			Tag, _ := Tokenisierer.TagName()
			Ausgabe.Write([]byte("<"))
			Ausgabe.Write([]byte(Übersetzen(Zeichenkette(Tag))))
			Ausgabe.Write([]byte("/>"))
		default:
			Ausgabe.Write(Tokenisierer.Text())
		}
	}
}

type Jedes any
type Vergleichbares comparable
type Zeichenkette string
type Ganzzahl int
type Fehler error
type Karte[von Vergleichbares, zu Jedes] map[von]zu
type Leser io.Reader
type Schreiber io.Writer

func (z Zeichenkette) AllesErsetzen(Alt, Neu Zeichenkette) Zeichenkette {
	return Zeichenkette(strings.ReplaceAll(string(z), string(Alt), string(Neu)))
}

func (z Zeichenkette) NeuerLeser() Leser {
	return strings.NewReader(string(z))
}
