// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TFormMain struct {
    *vcl.TForm
    Label1     *vcl.TLabel
    Gauge1     *vcl.TGauge
    LabelEdit1 *vcl.TLabeledEdit
    LabelEdit2 *vcl.TLabeledEdit
    Button1    *vcl.TButton
    Cbb1       *vcl.TComboBox
    Button2    *vcl.TButton
    Button3    *vcl.TButton
    Button4    *vcl.TButton

    //::private::
    TFormMainFields
}

var FormMain *TFormMain




// 以字节形式加载
// vcl.Application.CreateForm(formMainBytes, &FormMain)

func NewFormMain(owner vcl.IComponent) (root *TFormMain)  {
    vcl.CreateResForm(owner, formMainBytes, &root)
    return
}

var formMainBytes = []byte("\x54\x50\x46\x30\x09\x54\x46\x6F\x72\x6D\x4D\x61\x69\x6E\x08\x46\x6F\x72\x6D\x4D\x61\x69\x6E\x04\x4C\x65\x66\x74\x02\x00\x03\x54\x6F\x70\x02\x00\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x0A\x62\x69\x4D\x69\x6E\x69\x6D\x69\x7A\x65\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\x76\x20\x35\x34\x36\x35\x34\x36\x35\x31\x36\x35\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x47\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xB0\x01\x05\x43\x6F\x6C\x6F\x72\x07\x09\x63\x6C\x42\x74\x6E\x46\x61\x63\x65\x0C\x46\x6F\x6E\x74\x2E\x43\x68\x61\x72\x73\x65\x74\x07\x0C\x41\x4E\x53\x49\x5F\x43\x48\x41\x52\x53\x45\x54\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xEE\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x12\x04\x00\x00\x00\xAE\x5F\x6F\x8F\xC5\x96\xD1\x9E\x0A\x46\x6F\x6E\x74\x2E\x53\x74\x79\x6C\x65\x0B\x00\x0E\x4F\x6C\x64\x43\x72\x65\x61\x74\x65\x4F\x72\x64\x65\x72\x08\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0F\x70\x6F\x44\x65\x73\x6B\x74\x6F\x70\x43\x65\x6E\x74\x65\x72\x08\x4F\x6E\x43\x72\x65\x61\x74\x65\x07\x0A\x46\x6F\x72\x6D\x43\x72\x65\x61\x74\x65\x0D\x50\x69\x78\x65\x6C\x73\x50\x65\x72\x49\x6E\x63\x68\x02\x78\x0A\x54\x65\x78\x74\x48\x65\x69\x67\x68\x74\x02\x18\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x55\x03\x54\x6F\x70\x02\x65\x05\x57\x69\x64\x74\x68\x02\x4C\x06\x48\x65\x69\x67\x68\x74\x02\x18\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x05\x00\x00\x00\xFB\x4E\xA1\x52\x1A\x90\x53\x90\x3A\x00\x00\x00\x06\x54\x47\x61\x75\x67\x65\x06\x47\x61\x75\x67\x65\x31\x04\x4C\x65\x66\x74\x02\x00\x03\x54\x6F\x70\x03\x1E\x01\x05\x57\x69\x64\x74\x68\x03\xB0\x01\x06\x48\x65\x69\x67\x68\x74\x02\x29\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x42\x6F\x74\x74\x6F\x6D\x08\x50\x72\x6F\x67\x72\x65\x73\x73\x02\x00\x0C\x45\x78\x70\x6C\x69\x63\x69\x74\x4C\x65\x66\x74\x02\xFC\x0B\x45\x78\x70\x6C\x69\x63\x69\x74\x54\x6F\x70\x03\x20\x01\x0D\x45\x78\x70\x6C\x69\x63\x69\x74\x57\x69\x64\x74\x68\x03\x7D\x01\x00\x00\x0C\x54\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x0A\x4C\x61\x62\x65\x6C\x45\x64\x69\x74\x31\x04\x4C\x65\x66\x74\x03\xA7\x00\x03\x54\x6F\x70\x02\x19\x05\x57\x69\x64\x74\x68\x03\x92\x00\x06\x48\x65\x69\x67\x68\x74\x02\x20\x0F\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x57\x69\x64\x74\x68\x02\x28\x10\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x48\x65\x69\x67\x68\x74\x02\x18\x11\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x43\x61\x70\x74\x69\x6F\x6E\x12\x03\x00\x00\x00\x26\x8D\xF7\x53\x3A\x00\x07\x45\x6E\x61\x62\x6C\x65\x64\x08\x0D\x4C\x61\x62\x65\x6C\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x06\x6C\x70\x4C\x65\x66\x74\x0C\x4C\x61\x62\x65\x6C\x53\x70\x61\x63\x69\x6E\x67\x02\x07\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x0B\x31\x35\x37\x33\x38\x38\x38\x39\x37\x33\x30\x00\x00\x0C\x54\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x0A\x4C\x61\x62\x65\x6C\x45\x64\x69\x74\x32\x04\x4C\x65\x66\x74\x03\xA7\x00\x03\x54\x6F\x70\x02\x3F\x05\x57\x69\x64\x74\x68\x03\x92\x00\x06\x48\x65\x69\x67\x68\x74\x02\x20\x0F\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x57\x69\x64\x74\x68\x02\x28\x10\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x48\x65\x69\x67\x68\x74\x02\x18\x11\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x43\x61\x70\x74\x69\x6F\x6E\x12\x03\x00\x00\x00\xC6\x5B\x01\x78\x3A\x00\x07\x45\x6E\x61\x62\x6C\x65\x64\x08\x0D\x4C\x61\x62\x65\x6C\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x06\x6C\x70\x4C\x65\x66\x74\x0C\x4C\x61\x62\x65\x6C\x53\x70\x61\x63\x69\x6E\x67\x02\x07\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x08\x31\x32\x33\x34\x35\x36\x37\x38\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x02\x08\x03\x54\x6F\x70\x03\x9C\x00\x05\x57\x69\x64\x74\x68\x03\xB9\x00\x06\x48\x65\x69\x67\x68\x74\x02\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x02\x00\x00\x00\x7B\x76\x55\x5F\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0C\x42\x75\x74\x74\x6F\x6E\x31\x43\x6C\x69\x63\x6B\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x04\x43\x62\x62\x31\x04\x4C\x65\x66\x74\x03\xA7\x00\x03\x54\x6F\x70\x02\x65\x05\x57\x69\x64\x74\x68\x03\xC2\x00\x06\x48\x65\x69\x67\x68\x74\x02\x20\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x04\x54\x65\x78\x74\x12\x04\x00\x00\x00\x09\x90\xE9\x62\x1A\x90\x53\x90\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x03\xE6\x00\x03\x54\x6F\x70\x03\x9C\x00\x05\x57\x69\x64\x74\x68\x03\xC2\x00\x06\x48\x65\x69\x67\x68\x74\x02\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x0B\x00\x00\x00\xBD\x62\x37\x68\x8C\x5B\x10\x62\xFC\x5B\xFA\x51\x28\x00\x4A\x53\x57\x5B\xB5\x6B\x29\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0C\x42\x75\x74\x74\x6F\x6E\x32\x43\x6C\x69\x63\x6B\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x33\x04\x4C\x65\x66\x74\x02\x08\x03\x54\x6F\x70\x03\xE7\x00\x05\x57\x69\x64\x74\x68\x03\xB9\x00\x06\x48\x65\x69\x67\x68\x74\x02\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x0B\x00\x00\x00\xBD\x62\x37\x68\x8C\x5B\x10\x62\xFC\x5B\xFA\x51\x28\x00\x68\x51\x57\x5B\xB5\x6B\x29\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0C\x42\x75\x74\x74\x6F\x6E\x33\x43\x6C\x69\x63\x6B\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x34\x04\x4C\x65\x66\x74\x03\xE6\x00\x03\x54\x6F\x70\x03\xE7\x00\x05\x57\x69\x64\x74\x68\x03\xC2\x00\x06\x48\x65\x69\x67\x68\x74\x02\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x0A\x00\x00\x00\xF2\x5D\xA5\x63\x36\x65\xFC\x5B\xFA\x51\x28\x00\x68\x51\x57\x5B\xB5\x6B\x29\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0C\x42\x75\x74\x74\x6F\x6E\x34\x43\x6C\x69\x63\x6B\x00\x00\x00")
