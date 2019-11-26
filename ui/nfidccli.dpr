program nfidccli;

uses
  Vcl.Forms,
  uformlogin in 'uformlogin.pas' {FormLogin},
  uformhome in 'uformhome.pas' {FormHome},
  uformjiance in 'uformjiance.pas' {Formjiance};

{$R *.res}

begin
  Application.Initialize;
  Application.MainFormOnTaskbar := True;
  Application.CreateForm(TFormLogin, FormLogin);
  Application.CreateForm(TFormHome, FormHome);
  Application.CreateForm(TFormjiance, Formjiance);
  Application.Run;
end.
