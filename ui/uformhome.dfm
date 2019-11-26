object FormHome: TFormHome
  Left = 0
  Top = 0
  Caption = #20219#21153#24179#21488
  ClientHeight = 682
  ClientWidth = 1338
  Color = clBtnFace
  Constraints.MinHeight = 727
  Constraints.MinWidth = 1311
  Font.Charset = ANSI_CHARSET
  Font.Color = clWindowText
  Font.Height = -18
  Font.Name = #24494#36719#38597#40657
  Font.Style = []
  OldCreateOrder = False
  Position = poDesktopCenter
  OnClose = FormClose
  OnCreate = FormCreate
  OnShow = FormShow
  PixelsPerInch = 120
  TextHeight = 24
  object Pgc1: TPageControl
    AlignWithMargins = True
    Left = 3
    Top = 10
    Width = 1332
    Height = 669
    Margins.Top = 10
    ActivePage = Tss1
    Align = alClient
    TabOrder = 0
    object Tss1: TTabSheet
      Caption = #19979#36733#26679#21697#20449#24687
      OnShow = Tss1Show
      object Gauge1: TGauge
        Left = 0
        Top = 589
        Width = 1324
        Height = 41
        Align = alBottom
        Progress = 0
        ExplicitWidth = 1279
      end
      object Pnl2: TPanel
        Left = 0
        Top = 0
        Width = 1324
        Height = 89
        Align = alTop
        Caption = 'Pnl1'
        ShowCaption = False
        TabOrder = 0
        DesignSize = (
          1324
          89)
        object Label3: TLabel
          Left = 483
          Top = 11
          Width = 76
          Height = 24
          Anchors = [akTop, akRight]
          Caption = #36873#25321#31995#32479':'
        end
        object Label4: TLabel
          Left = 763
          Top = 11
          Width = 76
          Height = 24
          Anchors = [akTop, akRight]
          Caption = #20219#21153#29366#24577':'
        end
        object Label5: TLabel
          Left = 14
          Top = 11
          Width = 76
          Height = 24
          Caption = #25277#26679#26085#26399':'
        end
        object Label6: TLabel
          Left = 237
          Top = 11
          Width = 8
          Height = 24
          Caption = '-'
        end
        object Label2: TLabel
          Left = 15
          Top = 56
          Width = 76
          Height = 24
          Alignment = taRightJustify
          Caption = #20219#21153#26469#28304':'
        end
        object Cbbt1s1: TComboBox
          Left = 565
          Top = 8
          Width = 192
          Height = 32
          Anchors = [akTop, akRight]
          DoubleBuffered = False
          Font.Charset = ANSI_CHARSET
          Font.Color = clWindowText
          Font.Height = -18
          Font.Name = #24494#36719#38597#40657
          Font.Style = []
          ParentDoubleBuffered = False
          ParentFont = False
          TabOrder = 0
        end
        object Dtpt1s1: TDateTimePicker
          Left = 97
          Top = 8
          Width = 135
          Height = 32
          Date = 43781.000000000000000000
          Time = 0.933494293982221300
          DoubleBuffered = False
          ParentDoubleBuffered = False
          TabOrder = 1
        end
        object Cbbt1s2: TComboBox
          Left = 843
          Top = 8
          Width = 130
          Height = 32
          Anchors = [akTop, akRight]
          DoubleBuffered = False
          Font.Charset = ANSI_CHARSET
          Font.Color = clWindowText
          Font.Height = -18
          Font.Name = #24494#36719#38597#40657
          Font.Style = []
          ParentDoubleBuffered = False
          ParentFont = False
          TabOrder = 2
          OnChange = Cbbt1s2Change
          Items.Strings = (
            #25277#26679#23436#25104
            #24050#25509#25910
            #26816#39564#23436#25104)
        end
        object Dtpt1s2: TDateTimePicker
          Left = 251
          Top = 8
          Width = 137
          Height = 32
          Date = 43781.000000000000000000
          Time = 0.934004062502936000
          DoubleBuffered = False
          ParentDoubleBuffered = False
          TabOrder = 3
        end
        object Buttonp1s1: TButton
          Left = 987
          Top = 4
          Width = 89
          Height = 38
          Anchors = [akTop, akRight]
          Caption = #26597#35810
          TabOrder = 4
          OnClick = Buttonp1s1Click
        end
        object Buttonp1s2: TButton
          Left = 1223
          Top = 4
          Width = 89
          Height = 38
          Anchors = [akTop, akRight]
          Caption = #23548#20986
          TabOrder = 5
          OnClick = Buttonp1s2Click
        end
        object Cbbt1s3: TComboBox
          Left = 1082
          Top = 8
          Width = 135
          Height = 32
          Anchors = [akTop, akRight]
          TabOrder = 6
          Items.Strings = (
            #23548#20986'excel'
            #23548#20986'excel'#27169#24335'1')
        end
        object Edtt1s1: TEdit
          Left = 97
          Top = 54
          Width = 291
          Height = 32
          TabOrder = 7
        end
      end
      object ListView1: TListView
        AlignWithMargins = True
        Left = 3
        Top = 92
        Width = 1318
        Height = 494
        Align = alClient
        Columns = <
          item
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #26356#26032#26102#38388
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #25277#26679#21333#21495
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #26679#21697#21517#31216
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #25277#26679#26102#38388
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #20219#21153#26469#28304
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #25277#26679#20154
          end>
        DoubleBuffered = True
        Font.Charset = ANSI_CHARSET
        Font.Color = clWindowText
        Font.Height = -13
        Font.Name = #24494#36719#38597#40657
        Font.Style = []
        GridLines = True
        HideSelection = False
        OwnerData = True
        RowSelect = True
        ParentDoubleBuffered = False
        ParentFont = False
        TabOrder = 1
        ViewStyle = vsReport
        OnData = ListView1Data
        OnResize = ListView1Resize
      end
    end
    object Tss2: TTabSheet
      Caption = #19978#20256#26816#39564#32467#26524
      ImageIndex = 1
      OnShow = Tss2Show
      ExplicitLeft = 0
      ExplicitTop = 0
      ExplicitWidth = 0
      ExplicitHeight = 0
      object Gauge2: TGauge
        Left = 0
        Top = 589
        Width = 1324
        Height = 41
        Align = alBottom
        Progress = 0
        ExplicitWidth = 1279
      end
      object Pnl1: TPanel
        Left = 0
        Top = 0
        Width = 1324
        Height = 89
        Align = alTop
        Caption = 'Pnl1'
        ShowCaption = False
        TabOrder = 0
        DesignSize = (
          1324
          89)
        object Labelt2s1: TLabel
          Left = 128
          Top = 16
          Width = 76
          Height = 24
          Caption = #22635#25253#31867#22411':'
        end
        object Labelt2s2: TLabel
          Left = 835
          Top = 17
          Width = 81
          Height = 24
          Anchors = [akTop, akRight]
          Caption = #21305#37197#32467#26524': '
          ExplicitLeft = 790
        end
        object Label7: TLabel
          Left = 46
          Top = 57
          Width = 76
          Height = 24
          Alignment = taRightJustify
          Caption = #20219#21153#26469#28304':'
        end
        object Buttont2s1: TButton
          Left = 33
          Top = 8
          Width = 89
          Height = 38
          Caption = #23548#20837'excel'
          TabOrder = 0
          OnClick = Buttont2s1Click
        end
        object Cbbt2s1: TComboBox
          Left = 210
          Top = 11
          Width = 145
          Height = 32
          DoubleBuffered = False
          ParentDoubleBuffered = False
          TabOrder = 1
          Items.Strings = (
            #26222#36890#39135#21697#19978#25253
            #20892#20135#21697#19978#25253)
        end
        object Buttont2s2: TButton
          Left = 740
          Top = 8
          Width = 89
          Height = 38
          Anchors = [akTop, akRight]
          Caption = #26679#21697#21305#37197
          TabOrder = 2
          OnClick = Buttont2s2Click
        end
        object Buttont2s3: TButton
          Left = 1224
          Top = 8
          Width = 89
          Height = 38
          Anchors = [akTop, akRight]
          Caption = #25209#37327#19978#20256
          TabOrder = 3
          OnClick = Buttont2s3Click
        end
        object Edtt2s1: TEdit
          Left = 128
          Top = 53
          Width = 291
          Height = 32
          TabOrder = 4
        end
      end
      object ListView2: TListView
        AlignWithMargins = True
        Left = 3
        Top = 92
        Width = 1318
        Height = 494
        Align = alClient
        Checkboxes = True
        Columns = <
          item
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #26679#21697#21517#31216
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #39033
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #26679#21697#21305#37197
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #25277#26679#21333#21495
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #25253#21578#20070#32534#21495
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #26816#27979#39033#30446
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #30417#30563#25277#26816#25253#21578#22791#27880
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #39118#38505#26816#27979#25253#21578#22791#27880
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #19978#20256#29366#24577
          end
          item
            Alignment = taCenter
            AutoSize = True
            Caption = #19978#20256#32467#26524
          end>
        DoubleBuffered = True
        Font.Charset = ANSI_CHARSET
        Font.Color = clWindowText
        Font.Height = -13
        Font.Name = #24494#36719#38597#40657
        Font.Style = []
        GridLines = True
        OwnerData = True
        RowSelect = True
        ParentDoubleBuffered = False
        ParentFont = False
        TabOrder = 1
        ViewStyle = vsReport
        OnData = ListView2Data
        OnDblClick = ListView2DblClick
        OnResize = ListView2Resize
        ExplicitHeight = 492
      end
    end
    object Tss3: TTabSheet
      Caption = #31995#32479#35774#32622
      ImageIndex = 2
      ExplicitLeft = 0
      ExplicitTop = 0
      ExplicitWidth = 0
      ExplicitHeight = 0
      object Label1: TLabel
        Left = 0
        Top = 0
        Width = 90
        Height = 24
        Align = alClient
        Caption = #26242#26102#26080#35774#32622
      end
    end
  end
  object Timer1: TTimer
    Interval = 500
    OnTimer = Timer1Timer
    Left = 391
    Top = 293
  end
  object SaveDialog1: TSaveDialog
    DefaultExt = 'xlsx'
    Filter = 'excel|*.xlsx'
    Left = 519
    Top = 221
  end
  object DlgOpen1: TOpenDialog
    DefaultExt = 'xlsx'
    Filter = 'excel|*.xlsx'
    Left = 591
    Top = 213
  end
end
