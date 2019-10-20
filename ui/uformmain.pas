unit uformmain;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, Vcl.StdCtrls, Vcl.ExtCtrls,
  dxGDIPlusClasses, Vcl.Samples.Gauges;

type
  TFormMain = class(TForm)
    LabelEdit1: TLabeledEdit;
    LabelEdit2: TLabeledEdit;
    Button1: TButton;
    Label1: TLabel;
    Cbb1: TComboBox;
    Button3: TButton;
    Gauge1: TGauge;
    procedure Button1Click(Sender: TObject);
    procedure FormCreate(Sender: TObject);
    procedure Label2Click(Sender: TObject);
    procedure Button3Click(Sender: TObject);
  private
    { Private declarations }
  public
    { Public declarations }
  end;

var
  FormMain: TFormMain;

implementation

{$R *.dfm}

procedure TFormMain.Button1Click(Sender: TObject);
begin
//
end;

procedure TFormMain.Button3Click(Sender: TObject);
begin
//
end;

procedure TFormMain.FormCreate(Sender: TObject);
begin
//
end;

procedure TFormMain.Label2Click(Sender: TObject);
begin
  //
end;

end.
