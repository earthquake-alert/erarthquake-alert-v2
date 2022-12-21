package src_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/models"
	"github.com/stretchr/testify/require"
)

var TestEarthquakeActivity = []string{
	"32-35_09_01_191111_VXSE56.xml",
	"32-35_09_02_220316_VXSE56.xml",
}

func TestParseEarthquakeActivity(t *testing.T) {
	for _, d := range TestEarthquakeActivity {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseEarthquakeActivity(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquakeActivity([]byte(row))
		require.Error(t, err)
	})
}

// パーステスト1
func TestParseEarthquakeActivity1(t *testing.T) {
	target := "32-35_09_01_191111_VXSE56.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquakeActivity(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "地震の活動状況等に関する情報",
			DateTime:         "2019-11-11T08:00:53Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "地震の活動状況等に関する情報")
		require.Equal(t, head.ReportDateTime, "2019-11-11T17:00:00+09:00")
		require.Equal(t, head.TargetDateTime, "2019-11-11T17:00:00+09:00")
		require.Equal(t, head.EventID, "20191111170000")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "地震の活動状況等に関する情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "伊豆東部の地震活動の見通しに関する情報を発表します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		require.Equal(t, body.Text, `伊豆東部の地震活動の見通しに関する情報（第１号）

１．見出し
本日（１１日）昼から東伊豆奈良本（ならもと）観測点で縮みのひずみ変
化が観測され始め、本日（１１日）昼からは体に感じない小さな地震が発生
し始めました。
このことから、伊豆東部の地下でマグマの貫入と上昇が始まったと思われ
ます。このため、今後、概ね４日、長くて１週間程度、地震活動が活発にな
るおそれがあります。

２．現状
伊豆東部の地下でマグマの貫入と上昇が始まったとみられ、これに伴い、
地殻変動と地震活動が観測されています。

（地殻変動の状況）
本日（１１日）昼から東伊豆奈良本（ならもと）観測点で観測されている
縮みのひずみ変化は、本日（１１日）１６時現在継続しています。２４時間
変化量（換算値）は４００ｎｓｔｒａｉｎ（ナノストレイン）となっていま
す。また、防災科学技術研究所及び気象庁が整備している周辺の傾斜計にも
同期した変化が現われています。
今回のひずみ変化量に近い同様の活動は○○○○年○○月で、その時には
震度１以上を観測する地震が○○回、震度３以上が○○回発生しました。

（地震活動の状況）
本日（１１日）昼に始まった地震活動は、本日（１１日）１６時現在、活
発な状態となっています。本日（１１日）１６時００分にはＭ５.５の地震
が発生して、伊東市大原で震度５弱を観測しました。
これまでに、最大震度５弱を観測する地震が１回、最大震度３が２回、最
大震度２が３回、最大震度１が５回発生しました。
震源はいずれも川奈東沖と川奈崎付近の深さ０から５ｋｍです。

（火山活動の状況）
噴火に直ちに結びつくような現象は観測されていません。

３．見通し
本日（１１日）１６時００分現在の観測データから予測される地震活動の
見通しは、以下のとおりです。
＜今回の地震活動の見通し＞
地震の規模と震度 ： 最大マグニチュード６程度
最大震度５弱から５強程度 ※
震度１以上の地震回数： ５００から９００回程度
活動期間 ： 概ね４日、長くて１週間程度
※地盤の状況等により、更に揺れが大きくなる場合があります。

４．防災上の留意事項
活動期間の予測は一回のマグマ上昇に基づくため、複数回の上昇が起きた
場合はさらに長引くことがあります。
マグマがさらに浅部へ上昇した場合、地震活動がさらに活発になることが
あります。

次の伊豆東部の地震活動の見通しに関する情報は、本日（１１日）１８時
頃に発表の予定です。
なお、見通しの内容を更新する場合や、活動に顕著な変化があった場合等
には、随時お知らせします。`)

		require.Nil(t, body.Comments)
	})
}

// パーステスト2（取り消し）
func TestParseEarthquakeActivity2(t *testing.T) {
	target := "32-35_09_02_220316_VXSE56.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquakeActivity(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "地震の活動状況等に関する情報",
			DateTime:         "2022-03-16T04:46:04Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "地震の活動状況等に関する情報")
		require.Equal(t, head.ReportDateTime, "2022-03-16T13:46:00+09:00")
		require.Equal(t, head.TargetDateTime, "2022-03-16T13:46:00+09:00")
		require.Equal(t, head.EventID, "20220316133000")
		require.Equal(t, head.InfoType, jma.Cancel)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "地震の活動状況等に関する情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "伊豆東部の地震活動の見通しに関する情報を発表します。 ")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		require.Equal(t, body.Text, "１６日１３時３０分に発表した「伊豆東部の地震活動の見通しに関する情報」は誤りですので取り消します。 ")

		require.Nil(t, body.Comments)
	})
}

func TestEarthquakeActivityGetText(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		target := "32-35_09_01_191111_VXSE56.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeActivity(row)
		require.NoError(t, err)

		text, err := ea.GetText()
		require.NoError(t, err)

		require.Equal(t, text, "伊豆東部の地震活動の見通しに関する情報を発表します。\n\n伊豆東部の地震活動の見通しに関する情報(第1号)\n\n1.見出し\n本日(11日)昼から東伊豆奈良本(ならもと)観測点で縮みのひずみ変\n化が観測され始め、本日(11日)昼からは体に感じない小さな地震が発生\nし始めました。\nこのことから、伊豆東部の地下でマグマの貫入と上昇が始まったと思われ\nます。このため、今後、概ね4日、長くて1週間程度、地震活動が活発にな\nるおそれがあります。\n\n2.現状\n伊豆東部の地下でマグマの貫入と上昇が始まったとみられ、これに伴い、\n地殻変動と地震活動が観測されています。\n\n(地殻変動の状況)\n本日(11日)昼から東伊豆奈良本(ならもと)観測点で観測されている\n縮みのひずみ変化は、本日(11日)16時現在継続しています。24時間\n変化量(換算値)は400nstrain(ナノストレイン)となっていま\nす。また、防災科学技術研究所及び気象庁が整備している周辺の傾斜計にも\n同期した変化が現われています。\n今回のひずみ変化量に近い同様の活動は○○○○年○○月で、その時には\n震度1以上を観測する地震が○○回、震度3以上が○○回発生しました。\n\n(地震活動の状況)\n本日(11日)昼に始まった地震活動は、本日(11日)16時現在、活\n発な状態となっています。本日(11日)16時00分にはM5.5の地震\nが発生して、伊東市大原で震度5弱を観測しました。\nこれまでに、最大震度5弱を観測する地震が1回、最大震度3が2回、最\n大震度2が3回、最大震度1が5回発生しました。\n震源はいずれも川奈東沖と川奈崎付近の深さ0から5kmです。\n\n(火山活動の状況)\n噴火に直ちに結びつくような現象は観測されていません。\n\n3.見通し\n本日(11日)16時00分現在の観測データから予測される地震活動の\n見通しは、以下のとおりです。\n&lt;今回の地震活動の見通し&gt;\n地震の規模と震度 : 最大マグニチュード6程度\n最大震度5弱から5強程度 ※\n震度1以上の地震回数: 500から900回程度\n活動期間 : 概ね4日、長くて1週間程度\n※地盤の状況等により、更に揺れが大きくなる場合があります。\n\n4.防災上の留意事項\n活動期間の予測は一回のマグマ上昇に基づくため、複数回の上昇が起きた\n場合はさらに長引くことがあります。\nマグマがさらに浅部へ上昇した場合、地震活動がさらに活発になることが\nあります。\n\n次の伊豆東部の地震活動の見通しに関する情報は、本日(11日)18時\n頃に発表の予定です。\nなお、見通しの内容を更新する場合や、活動に顕著な変化があった場合等\nには、随時お知らせします。")
	})

	t.Run("2", func(t *testing.T) {
		target := "32-35_09_02_220316_VXSE56.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeActivity(row)
		require.NoError(t, err)

		text, err := ea.GetText()
		require.NoError(t, err)

		require.Equal(t, text, "伊豆東部の地震活動の見通しに関する情報を発表します。 \n\n16日13時30分に発表した「伊豆東部の地震活動の見通しに関する情報」は誤りですので取り消します。 ")
	})
}

func TestEarthquakeActivityGetTitle(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		target := "32-35_09_01_191111_VXSE56.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeActivity(row)
		require.NoError(t, err)

		title := ea.GetTitle()

		require.Equal(t, title, "地震の活動状況等に関する情報")
	})

	t.Run("2", func(t *testing.T) {
		target := "32-35_09_02_220316_VXSE56.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeActivity(row)
		require.NoError(t, err)

		title := ea.GetTitle()

		require.Equal(t, title, "【取消】地震の活動状況等に関する情報")
	})
}

func TestEarthquakeActivityGetTargetDate(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		target := "32-35_09_01_191111_VXSE56.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeActivity(row)
		require.NoError(t, err)

		d, err := ea.GetTargetDate()
		require.NoError(t, err)

		require.Equal(t, d.String(), "2019-11-11 17:00:00 +0000 UTC")
	})

	t.Run("2", func(t *testing.T) {
		target := "32-35_09_02_220316_VXSE56.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeActivity(row)
		require.NoError(t, err)

		d, err := ea.GetTargetDate()
		require.NoError(t, err)

		require.Equal(t, d.String(), "2022-03-16 13:46:00 +0000 UTC")

	})
}

func TestEarthquakeActivityGetEventId(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		target := "32-35_09_01_191111_VXSE56.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeActivity(row)
		require.NoError(t, err)

		e, err := ea.GetEventId()
		require.NoError(t, err)

		require.Equal(t, e, []int{20191111170000})
	})

	t.Run("2", func(t *testing.T) {
		target := "32-35_09_02_220316_VXSE56.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeActivity(row)
		require.NoError(t, err)

		e, err := ea.GetEventId()
		require.NoError(t, err)

		require.Equal(t, e, []int{20220316133000})

	})
}

func TestEarthquakeActivityAssembly(t *testing.T) {
	ctx := context.Background()
	db, err := src.NewConnectMySQL(ctx)
	require.NoError(t, err)

	t.Run("DBに格納される", func(t *testing.T) {
		t.Run("1", func(t *testing.T) {
			target := "32-35_09_01_191111_VXSE56.xml"

			testPath := filepath.Join(TEST_DATA_PATH, target)
			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			ea, err := src.ParseEarthquakeActivity(row)
			require.NoError(t, err)

			err = ea.Assembly(ctx, db)
			require.NoError(t, err)

			eventIds, err := ea.GetEventId()
			require.NoError(t, err)

			exists, err := models.EarthquakeActivities(
				models.EarthquakeActivityWhere.EventID.EQ(int64(eventIds[0])),
			).Exists(ctx, db)
			require.NoError(t, err)
			require.True(t, exists)
		})

		t.Run("2", func(t *testing.T) {
			target := "32-35_09_02_220316_VXSE56.xml"

			testPath := filepath.Join(TEST_DATA_PATH, target)
			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			ea, err := src.ParseEarthquakeActivity(row)
			require.NoError(t, err)

			err = ea.Assembly(ctx, db)
			require.NoError(t, err)

			eventIds, err := ea.GetEventId()
			require.NoError(t, err)

			exists, err := models.EarthquakeActivities(
				models.EarthquakeActivityWhere.EventID.EQ(int64(eventIds[0])),
			).Exists(ctx, db)
			require.NoError(t, err)
			require.True(t, exists)
		})

		t.Run("正しく全て入っている", func(t *testing.T) {
			target := "32-35_09_01_191111_VXSE56.xml"

			testPath := filepath.Join(TEST_DATA_PATH, target)
			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			ea, err := src.ParseEarthquakeActivity(row)
			require.NoError(t, err)

			err = ea.Assembly(ctx, db)
			require.NoError(t, err)

			eventIds, err := ea.GetEventId()
			require.NoError(t, err)

			a, err := models.EarthquakeActivities(
				models.EarthquakeActivityWhere.EventID.EQ(int64(eventIds[0])),
			).One(ctx, db)
			require.NoError(t, err)

			require.Equal(t, a.EventID, int64(eventIds[0]))
			require.NotNil(t, a.Created)
			require.NotNil(t, a.ID)
			require.NotNil(t, a.Date)
			require.Equal(t, a.Row, string(row))
		})
	})
}
