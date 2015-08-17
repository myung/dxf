package object

import (
	"errors"
	"fmt"
	"github.com/yofu/dxf/format"
	"github.com/yofu/dxf/handle"
)

type AcDbDictionaryWDFLT struct {
	handle        int
	item          map[string]handle.Handler
	owner         handle.Handler
	defaulthandle handle.Handler
}

func (d *AcDbDictionaryWDFLT) IsObject() bool {
	return true
}

func NewAcDbDictionaryWDFLT(owner handle.Handler) (*AcDbDictionaryWDFLT, *AcDbPlaceHolder) {
	ds := make(map[string]handle.Handler)
	p := NewAcDbPlaceHolder()
	ds["Normal"] = p
	d := &AcDbDictionaryWDFLT{
		handle:        0,
		item:          ds,
		owner:         owner,
		defaulthandle: p,
	}
	p.owner = d
	return d, p
}

func (d *AcDbDictionaryWDFLT) Format(f *format.Formatter) {
	f.WriteString(0, "ACDBDICTIONARYWDFLT")
	f.WriteHex(5, d.handle)
	if d.owner != nil {
		f.WriteString(102, "{ACAD_REACTORS")
		f.WriteHex(330, d.owner.Handle())
		f.WriteString(102, "}")
		f.WriteHex(330, d.owner.Handle())
	}
	f.WriteString(100, "AcDbDictionary")
	f.WriteInt(281, 1)
	for k, v := range d.item {
		f.WriteString(3, k)
		f.WriteHex(350, v.Handle())
	}
	f.WriteString(100, "AcDbDictionaryWithDefault")
	f.WriteHex(340, d.defaulthandle.Handle())
}

func (d *AcDbDictionaryWDFLT) String() string {
	f := format.New()
	return d.FormatString(f)
}

func (d *AcDbDictionaryWDFLT) FormatString(f *format.Formatter) string {
	d.Format(f)
	return f.Output()
}

func (d *AcDbDictionaryWDFLT) Handle() int {
	return d.handle
}
func (d *AcDbDictionaryWDFLT) SetHandle(v *int) {
	d.handle = *v
	(*v)++
}

func (d *AcDbDictionaryWDFLT) AddItem(key string, value handle.Handler) error {
	if _, exist := d.item[key]; exist {
		return errors.New(fmt.Sprintf("key %s already exists"))
	}
	d.item[key] = value
	return nil
}