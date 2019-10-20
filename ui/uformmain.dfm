object FormMain: TFormMain
  Left = 0
  Top = 0
  Caption = 'v 123819004'
  ClientHeight = 281
  ClientWidth = 377
  Color = clBtnFace
  Font.Charset = ANSI_CHARSET
  Font.Color = clWindowText
  Font.Height = -18
  Font.Name = #24494#36719#38597#40657
  Font.Style = []
  OldCreateOrder = False
  Position = poDesktopCenter
  OnCreate = FormCreate
  PixelsPerInch = 120
  TextHeight = 24
  object Label1: TLabel
    Left = 45
    Top = 109
    Width = 76
    Height = 24
    Caption = #20219#21153#36890#36947':'
  end
  object Gauge1: TGauge
    Left = -4
    Top = 240
    Width = 381
    Height = 41
    Progress = 0
  end
  object LabelEdit1: TLabeledEdit
    Left = 127
    Top = 33
    Width = 146
    Height = 32
    EditLabel.Width = 40
    EditLabel.Height = 24
    EditLabel.Caption = #36134#21495':'
    Enabled = False
    LabelPosition = lpLeft
    LabelSpacing = 7
    TabOrder = 0
    Text = '15738889730'
  end
  object LabelEdit2: TLabeledEdit
    Left = 127
    Top = 71
    Width = 146
    Height = 32
    EditLabel.Width = 40
    EditLabel.Height = 24
    EditLabel.Caption = #23494#30721':'
    Enabled = False
    LabelPosition = lpLeft
    LabelSpacing = 7
    TabOrder = 1
    Text = '12345678'
  end
  object Button1: TButton
    Left = 39
    Top = 156
    Width = 121
    Height = 49
    Caption = #30331#24405
    TabOrder = 2
    OnClick = Button1Click
  end
  object Cbb1: TComboBox
    Left = 127
    Top = 109
    Width = 194
    Height = 32
    TabOrder = 3
    Text = #36873#25321#36890#36947
  end
  object Button3: TButton
    Left = 224
    Top = 156
    Width = 113
    Height = 49
    Caption = #19979#36733#25968#25454
    TabOrder = 4
    OnClick = Button3Click
  end
end
