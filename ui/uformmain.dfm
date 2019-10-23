object FormMain: TFormMain
  Left = 0
  Top = 0
  BorderIcons = [biSystemMenu, biMinimize]
  Caption = 'v 5465465165'
  ClientHeight = 327
  ClientWidth = 432
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
    Left = 85
    Top = 101
    Width = 76
    Height = 24
    Caption = #20219#21153#36890#36947':'
  end
  object Gauge1: TGauge
    Left = 0
    Top = 286
    Width = 432
    Height = 41
    Align = alBottom
    Progress = 0
    ExplicitLeft = -4
    ExplicitTop = 288
    ExplicitWidth = 381
  end
  object LabelEdit1: TLabeledEdit
    Left = 167
    Top = 25
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
    Left = 167
    Top = 63
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
    Left = 8
    Top = 156
    Width = 185
    Height = 49
    Caption = #30331#24405
    TabOrder = 2
    OnClick = Button1Click
  end
  object Cbb1: TComboBox
    Left = 167
    Top = 101
    Width = 194
    Height = 32
    TabOrder = 3
    Text = #36873#25321#36890#36947
  end
  object Button2: TButton
    Left = 230
    Top = 156
    Width = 194
    Height = 49
    Caption = #25277#26679#23436#25104#23548#20986'('#21322#23383#27573')'
    TabOrder = 4
    OnClick = Button2Click
  end
  object Button3: TButton
    Left = 8
    Top = 231
    Width = 185
    Height = 49
    Caption = #25277#26679#23436#25104#23548#20986'('#20840#23383#27573')'
    TabOrder = 5
    OnClick = Button3Click
  end
  object Button4: TButton
    Left = 230
    Top = 231
    Width = 194
    Height = 49
    Caption = #24050#25509#25910#23548#20986'('#20840#23383#27573')'
    TabOrder = 6
    OnClick = Button4Click
  end
end
