object FormLogin: TFormLogin
  Left = 0
  Top = 0
  BorderIcons = [biSystemMenu, biMinimize]
  Caption = #30331#24405
  ClientHeight = 291
  ClientWidth = 399
  Color = clBtnFace
  Constraints.MinHeight = 335
  Constraints.MinWidth = 417
  Font.Charset = ANSI_CHARSET
  Font.Color = clWindowText
  Font.Height = -18
  Font.Name = #24494#36719#38597#40657
  Font.Style = []
  OldCreateOrder = False
  Position = poDesktopCenter
  OnCreate = FormCreate
  OnShow = FormShow
  PixelsPerInch = 120
  TextHeight = 24
  object Label1: TLabel
    Left = 128
    Top = 16
    Width = 162
    Height = 33
    Caption = #25968#25454#21516#27493#32452#20214
    Font.Charset = DEFAULT_CHARSET
    Font.Color = clWindowText
    Font.Height = -27
    Font.Name = 'Tahoma'
    Font.Style = []
    ParentFont = False
  end
  object Label2: TLabel
    Left = 167
    Top = 259
    Width = 50
    Height = 24
    Caption = 'v0.0.7'
  end
  object LabelEdit1: TLabeledEdit
    Left = 128
    Top = 70
    Width = 180
    Height = 32
    EditLabel.Width = 58
    EditLabel.Height = 24
    EditLabel.Caption = #29992#25143#21517':'
    LabelPosition = lpLeft
    TabOrder = 0
  end
  object LabelEdit2: TLabeledEdit
    Left = 128
    Top = 128
    Width = 180
    Height = 32
    EditLabel.Width = 60
    EditLabel.Height = 24
    EditLabel.Caption = #23494'    '#30721':'
    Font.Charset = ANSI_CHARSET
    Font.Color = clWindowText
    Font.Height = -18
    Font.Name = #24494#36719#38597#40657
    Font.Style = []
    LabelPosition = lpLeft
    ParentFont = False
    PasswordChar = '*'
    TabOrder = 1
  end
  object Button1: TButton
    Left = 56
    Top = 192
    Width = 105
    Height = 49
    Caption = #30331#24405
    TabOrder = 2
    OnClick = Button1Click
  end
  object Button2: TButton
    Left = 219
    Top = 192
    Width = 113
    Height = 49
    Caption = #36864#20986
    TabOrder = 3
    OnClick = Button2Click
  end
end
