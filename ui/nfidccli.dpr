program nfidccli;

uses
  Vcl.Forms,
  uformmain in 'uformmain.pas' {FormMain},
  uformlogin in 'uformlogin.pas' {FormLogin},
  uformhome in 'uformhome.pas' {FormHome},
  uformjiance in 'uformjiance.pas' {Formjiance};

{$R *.res}

begin
  Application.Initialize;
  Application.MainFormOnTaskbar := True;
  Application.CreateForm(TFormLogin, FormLogin);
  Application.CreateForm(TFormHome, FormHome);
  Application.CreateForm(TFormMain, FormMain);
  Application.CreateForm(TFormjiance, Formjiance);
  Application.Run;
end.
