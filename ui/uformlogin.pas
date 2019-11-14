unit uformlogin;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, Vcl.StdCtrls, Vcl.ExtCtrls;

type
  TFormLogin = class(TForm)
    Label1: TLabel;
    LabelEdit1: TLabeledEdit;
    LabelEdit2: TLabeledEdit;
    Button1: TButton;
    Button2: TButton;
    Label2: TLabel;
    procedure Button2Click(Sender: TObject);
    procedure Button1Click(Sender: TObject);
    procedure FormCreate(Sender: TObject);
    procedure FormShow(Sender: TObject);
  private
    { Private declarations }
  public
    { Public declarations }
  end;

var
  FormLogin: TFormLogin;

implementation

{$R *.dfm}

procedure TFormLogin.Button1Click(Sender: TObject);
begin
//
end;

procedure TFormLogin.Button2Click(Sender: TObject);
begin
//
end;

procedure TFormLogin.FormCreate(Sender: TObject);
begin
//
end;

procedure TFormLogin.FormShow(Sender: TObject);
begin
//
end;

end.
