package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"strings"
	"unsafe"

	"github.com/apache/arrow/go/v9/parquet"
	"github.com/apache/arrow/go/v9/parquet/file"
	"github.com/apache/arrow/go/v9/parquet/schema"
)

import "C"

//go:embed template.html
var tpl string

type TemplateColumnData struct {
	Name string
	Type string
}
type TemplateData struct {
	Filename string
	NCols    int
	NRows    int64
	Columns  []TemplateColumnData
}

func OpenParquetFile(filename string, opts ...file.ReadOption) (*file.Reader, error) {
	var source parquet.ReaderAtSeeker

	var err error

	source, err = os.Open(filename)
	if err != nil {
		return nil, err
	}

	return file.NewParquetReader(source, opts...)
}

func GetTemplateData(filename string) *TemplateData {
	fr, err := OpenParquetFile(filename)
	if err != nil {
		panic(err)
	}
	metadata := fr.MetaData()
	ncols := metadata.Schema.NumColumns()

	tplData := TemplateData{}
	tplData.Filename = path.Base(filename)
	tplData.NCols = ncols
	tplData.NRows = metadata.NumRows
	tplData.Columns = []TemplateColumnData{}

	for i := 0; i < ncols; i++ {
		var fullType strings.Builder
		descr := metadata.Schema.Column(i)

		fullType.WriteString(descr.PhysicalType().String())
		if descr.ConvertedType() != schema.ConvertedTypes.None {
			fullType.WriteString(fmt.Sprintf("/%s", descr.ConvertedType()))
			if descr.ConvertedType() == schema.ConvertedTypes.Decimal {
				dec := descr.LogicalType().(*schema.DecimalLogicalType)
				fullType.WriteString(fmt.Sprintf("(%d,%d)", dec.Precision(), dec.Scale()))
			}
		}

		tplData.Columns = append(tplData.Columns, TemplateColumnData{
			Name: descr.Name(),
			Type: fullType.String(),
		})
	}

	return &tplData
}

//export GetParquetSummary
func GetParquetSummary(cpath *C.char) (code C.int, outData unsafe.Pointer, outLen C.long) {
	path := C.GoString(cpath)

	var buf bytes.Buffer

	// TODO: Factor out into constant
	t, err := template.New("template").Parse(tpl)
	if err != nil {
		panic(err)
	}
	data := GetTemplateData(path)
	err = t.Execute(io.Writer(&buf), data)

	return 0, C.CBytes(buf.Bytes()), C.long(buf.Len())
}

// Useful for development/debugging
func main() {
	fmt.Println(tpl)
	t, err := template.New("template").Parse(tpl)
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage:")
		return
	}

	path := os.Args[1]

	_, err = os.Stat(path)
	if err != nil {
		panic(err)
	}

	data := GetTemplateData(path)
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
