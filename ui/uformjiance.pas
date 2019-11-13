unit uformjiance;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, Vcl.ComCtrls;

type
  TFormjiance = class(TForm)
    ListView1: TListView;
    procedure FormShow(Sender: TObject);
    procedure ListView1Resize(Sender: TObject);
  private
    { Private declarations }
  public
    { Public declarations }
  end;

var
  Formjiance: TFormjiance;

implementation

{$R *.dfm}

procedure TFormjiance.FormShow(Sender: TObject);
begin
//
end;

procedure TFormjiance.ListView1Resize(Sender: TObject);
begin
//
end;

end.
