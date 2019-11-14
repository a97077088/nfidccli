unit uformhome;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, Vcl.ComCtrls, Vcl.ToolWin, Vcl.StdCtrls,
  Vcl.ExtCtrls, Vcl.Samples.Gauges;

type
  TFormHome = class(TForm)
    Pgc1: TPageControl;
    Tss1: TTabSheet;
    Tss2: TTabSheet;
    Tss3: TTabSheet;
    Pnl2: TPanel;
    Label3: TLabel;
    Label4: TLabel;
    Label5: TLabel;
    Label6: TLabel;
    Cbbt1s1: TComboBox;
    Dtpt1s1: TDateTimePicker;
    Cbbt1s2: TComboBox;
    Dtpt1s2: TDateTimePicker;
    Buttonp1s1: TButton;
    Buttonp1s2: TButton;
    ListView1: TListView;
    Timer1: TTimer;
    Gauge1: TGauge;
    SaveDialog1: TSaveDialog;
    Label1: TLabel;
    Pnl1: TPanel;
    Buttont2s1: TButton;
    Labelt2s1: TLabel;
    Cbbt2s1: TComboBox;
    Labelt2s2: TLabel;
    Buttont2s2: TButton;
    Buttont2s3: TButton;
    ListView2: TListView;
    Gauge2: TGauge;
    DlgOpen1: TOpenDialog;
    Cbbt1s3: TComboBox;
    procedure FormClose(Sender: TObject; var Action: TCloseAction);
    procedure FormCreate(Sender: TObject);
    procedure Tss1Show(Sender: TObject);
    procedure FormShow(Sender: TObject);
    procedure Buttonp1s1Click(Sender: TObject);
    procedure ListView1Data(Sender: TObject; Item: TListItem);
    procedure Buttonp1s2Click(Sender: TObject);
    procedure ListView1Resize(Sender: TObject);
    procedure Timer1Timer(Sender: TObject);
    procedure Buttont2s1Click(Sender: TObject);
    procedure ListView2Data(Sender: TObject; Item: TListItem);
    procedure ListView2Resize(Sender: TObject);
    procedure Buttont2s2Click(Sender: TObject);
    procedure Tss2Show(Sender: TObject);
    procedure Buttont2s3Click(Sender: TObject);
    procedure ListView2DblClick(Sender: TObject);
    procedure Cbbt1s2Change(Sender: TObject);
  private
    { Private declarations }
  public
    { Public declarations }
  end;

var
  FormHome: TFormHome;

implementation

{$R *.dfm}

procedure TFormHome.Buttonp1s1Click(Sender: TObject);
begin
//
end;

procedure TFormHome.Buttonp1s2Click(Sender: TObject);
begin
//
end;

procedure TFormHome.Buttont2s1Click(Sender: TObject);
begin
//
end;

procedure TFormHome.Buttont2s2Click(Sender: TObject);
begin
//
end;

procedure TFormHome.Buttont2s3Click(Sender: TObject);
begin
//
end;

procedure TFormHome.Cbbt1s2Change(Sender: TObject);
begin
//
end;

procedure TFormHome.FormClose(Sender: TObject; var Action: TCloseAction);
begin
//
end;

procedure TFormHome.FormCreate(Sender: TObject);
begin
//
end;

procedure TFormHome.FormShow(Sender: TObject);
begin
//
end;

procedure TFormHome.ListView1Data(Sender: TObject; Item: TListItem);
begin
//
end;

procedure TFormHome.ListView1Resize(Sender: TObject);
begin
//
end;

procedure TFormHome.ListView2Data(Sender: TObject; Item: TListItem);
begin
//
end;

procedure TFormHome.ListView2DblClick(Sender: TObject);
begin
//
end;

procedure TFormHome.ListView2Resize(Sender: TObject);
begin
//
end;

procedure TFormHome.Timer1Timer(Sender: TObject);
begin
//
end;

procedure TFormHome.Tss1Show(Sender: TObject);
begin
//
end;

procedure TFormHome.Tss2Show(Sender: TObject);
begin
//
end;

end.
