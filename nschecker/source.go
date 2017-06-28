package nschecker

var Sources = []Source{
	{
		Name:          "Amazon - Nintendo Switch Joy-Con (L) / (R) グレー",
		URL:           "https://www.amazon.co.jp/%E4%BB%BB%E5%A4%A9%E5%A0%82-Nintendo-Switch-Joy-Con-%E3%82%B0%E3%83%AC%E3%83%BC/dp/B01N5QLLT3/",
		AvailableText: `この商品は、<a href="/gp/help/customer/display.html?ie=UTF8&amp;nodeId=643004">Amazon.co.jp</a> が販売、発送します。`,
	},
	{
		Name:          "Amazon - Nintendo Switch Joy-Con (L) ネオンブルー/ (R) ネオンレッド",
		URL:           "https://www.amazon.co.jp/Nintendo-Switch-Joy-Con-%E3%83%8D%E3%82%AA%E3%83%B3%E3%83%96%E3%83%AB%E3%83%BC-%E3%83%8D%E3%82%AA%E3%83%B3%E3%83%AC%E3%83%83%E3%83%89/dp/B01NCXFWIZ/",
		AvailableText: `この商品は、<a href="/gp/help/customer/display.html?ie=UTF8&amp;nodeId=643004">Amazon.co.jp</a> が販売、発送します。`,
	},
	{
		Name:          "Amazon グレー",
		URL:           "https://www.amazon.co.jp/%E4%BB%BB%E5%A4%A9%E5%A0%82-Nintendo-Switch-Joy-Con-%E3%82%B0%E3%83%AC%E3%83%BC/dp/B01N5QLLT3/",
		AvailableText: `</a><span class='a-color-price'>￥ 32,378</span>より</span>`,
	},
	{
		Name:          "Amazon ネオンブルー/ (R) ネオンレッド",
		URL:           "https://www.amazon.co.jp/Nintendo-Switch-Joy-Con-%E3%83%8D%E3%82%AA%E3%83%B3%E3%83%96%E3%83%AB%E3%83%BC-%E3%83%8D%E3%82%AA%E3%83%B3%E3%83%AC%E3%83%83%E3%83%89/dp/B01NCXFWIZ/",
		AvailableText: `</a><span class='a-color-price'>￥ 32,378</span>より</span>`,
	},
	{
		Name:        "My Nintendo Store",
		URL:         "https://store.nintendo.co.jp/customize.html",
		SoldOutText: `<button class="btn btn__primary_soldout to_cart" type="submit"><span>SOLD OUT</span></button>`,
	},
	{
		Name:        "Yodobashi - Nintendo Switch Joy-Con(L)/(R)グレー [Nintendo Switch本体]",
		URL:         "http://www.yodobashi.com/product/100000001003431565/",
		SoldOutText: `<div class="salesInfo"><p>予定数の販売を終了しました</p></div>`,
	},
	{
		Name:        "Yodobashi - Nintendo Switch Joy-Con(L)ネオンブルー/(R)ネオンレッド [Nintendo Switch本体]",
		URL:         "http://www.yodobashi.com/product/100000001003431566/",
		SoldOutText: `<div class="salesInfo"><p>予定数の販売を終了しました</p></div>`,
	},
	{
		Name:        "omni7(7net) - Nintendo Switch Joy-Con (L) / (R) グレー",
		URL:         "http://7net.omni7.jp/detail/2110595636",
		SoldOutText: `<input class="linkBtn js-pressTwice" type="submit" value="在庫切れ" title="在庫切れ"`,
	},
	{
		Name:        "omni7(7net) - Nintendo Switch Joy-Con (L) ネオンブルー/ (R) ネオンレッド",
		URL:         "http://7net.omni7.jp/detail/2110595637",
		SoldOutText: `<input class="linkBtn js-pressTwice" type="submit" value="在庫切れ" title="在庫切れ"`,
	},
	{
		Name:        "omni7(iyec) - Nintendo Switch Joy-Con (L) / (R) グレー",
		URL:         "http://iyec.omni7.jp/detail/4902370535709",
		SoldOutText: `<input class="linkBtn js-pressTwice" type="submit" value="在庫切れ" title="在庫切れ"`,
	},
	{
		Name:        "omni7(iyec) - Nintendo Switch Joy-Con (L) ネオンブルー/ (R) ネオンレッド",
		URL:         "http://iyec.omni7.jp/detail/4902370535716",
		SoldOutText: `<input class="linkBtn js-pressTwice" type="submit" value="在庫切れ" title="在庫切れ"`,
	},
	{
		Name:        "toysrus - Nintendo Switch Joy-Con (L) / (R) グレー",
		URL:         "https://www.toysrus.co.jp/s/dsg-572182200",
		SoldOutText: `<span id="isStock_c" >在庫なし/入荷予定あり</span>`,
	},
	{
		Name:        "toysrus - Nintendo Switch Joy-Con (L) ネオンブルー/ (R) ネオンレッド",
		URL:         "https://www.toysrus.co.jp/s/dsg-572186500",
		SoldOutText: `<span id="isStock_c" >在庫なし/入荷予定あり</span>`,
	},
}
