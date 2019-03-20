// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
        "math/rand"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				
				log.Println(message.Text)
				var out = ""
				var img = ""
				inText := strings.ToLower(message.Text)
				// user for cs
				if strings.Contains(inText, "/cs"){
					
					if strings.Contains(inText, "小宇宙")  {
						out = "小宇宙萌萌達~~"
					}
					if strings.Contains(inText, "胖") || strings.Contains(inText, "月半") {
						out = "小宇宙才不胖！！！"
					}
					if strings.Contains(inText, "肥") || strings.Contains(inText, "月巴") {
						out = "小宇宙才不胖！！！"
					}
					if strings.Contains(inText, "圓")  {
						out = "小宇宙才不圓！！！"
					}
					if strings.Contains(inText, "天然")||strings.Contains(inText, "天燃"){
						out = "你才天然，你全家都天然"
					}
					if strings.Contains(inText, "美"){
						out = "氣質 品味 知性。小宇宙！"
					}
					if strings.Contains(inText, "跌倒"){
						out = "如果現實有輕功可以學就好了 那就不會跌倒了（哭暈"
					}
					if strings.Contains(inText, "香港"){
						out = "深圳灣滿出來了"
					}					
					if strings.Contains(inText, "颱風") || strings.Contains(inText, "地震")  || strings.Contains(inText, "水災")   {
						out = "小宇宙才沒有那種能力"
					}
					if strings.Contains(inText, "蒸蛋")  {
						out = "小宇宙是小廚神~"
					}
					if strings.Contains(inText, "香菇")  {
						out = "小宇宙喜歡吃香菇~"
					}
					if strings.Contains(inText, "昭和")  {
						old := []string{
							"你才昭和你全家才昭和"," 我以為這很常見？","小宇宙是聽說的!",
						}
						out = old[rand.Intn(len(old))]
					}
					if strings.Contains(inText, "豆芽菜")  {
						out = " ｀ﾟдﾟ)!!!!!! 剛剛去浴室洗個腳 發現浴室排水孔..... 竟然長出一根豆芽.............. (●ﾟ灬 ﾟ●)"
					}
					if strings.Contains(inText, "抽")  {
						out = "抽抽~"
					}
					// for lucky
					if strings.Contains(inText, "運勢") {
						lucky := []string{
						"大吉","中吉","小吉","吉","末吉","凶","大凶",
						}
						out = lucky[rand.Intn(len(lucky))]
					}
					if strings.Contains(inText, "/cs") && strings.Contains(inText, "買嗎") {
						buy := []string{
							"考慮一下","買！！！","買~~~","買買買~",
						}
						out = buy[rand.Intn(len(buy))]
					}
					if strings.Contains(inText, "吃什麼")||strings.Contains(inText, "吃") {
						eat := []string{
							"飯飯","麵麵","打邊爐","炸豬排","燒烤","烤雞翅","蒸蛋","歐姆蛋",
							"炒米粉","蘿蔔糕",
						}
						out = eat[rand.Intn(len(eat))]
					}
					
					if strings.Contains(inText, "天氣") || strings.Contains(inText, "下雨") {
						weather := []string{
						"下大雨唷","招喚颱風唷","長香菇唷","飄雨","打雷你唷",
						}
						if strings.Contains(inText, "今天") || strings.Contains(inText, "明天") || strings.Contains(inText, "後天"){
							out += message.Text+"會"	
							out = strings.Trim(out, "/cs")
						}
						out += weather[rand.Intn(len(weather))]
					}	
					if strings.Contains(inText, "燒")  {
						out = "燒毀 通通燒"
					}

					if strings.Contains(inText, "奈良的鹿")  {
						out = "我發現 我把奈良的鹿都背起來了!"
					}
					//小宇宙的自白部分///
					if strings.Contains(inText, "天才")  {
						out = "小宇宙超天才！"
					}
					if strings.Contains(inText, "黛玉")  {
						out = "黛玉 小宇宙是小黛玉！！"
					}
					if strings.Contains(inText, "聰明")  {
						out = "小宇宙是超聰明！！"
					}//（zzzzz)
					if strings.Contains(inText, "起床")  {
						out = "（zzzzz)"
					}
					if strings.Contains(inText, "哈比人")  {
						out = "哈比人能一天吃七餐 但是小宇宙能一天吃九餐唷～"
					}
					if strings.Contains(inText, "曬衣服")  {
						lines := []string {	
							"認真的說 好不容易終於等到今天送來了家具 包括晾衣竿晾衣架衣櫥桌子 ",
							"我終於可以在陽台曬衣服了！！！ ",
							"雖然時間晚了但是還是開心的走出陽台準備把晾衣竿綁好的時候..... ",
							"一推開門就是大雨是怎樣！！！（摔",
							"我真的很認真的！！！",
							"(●ﾟ一 ﾟ●)",
						}
						out = strings.Join(lines, "\r\n")
					}
					if strings.Contains(inText, "肚子痛")  {
						lines := []string {	
							"向外頭拉一拉 拉出一杯拉茶 拉走肚子的寂寞",
							"往外頭潑一潑 潑出一杯功夫茶 潑出肚子的彩虹",
							"朝外頭澆一澆 澆出一杯茉莉花茶 澆成小宇宙的花樣年華~ 一十八~ ヾ(●ﾟ∀ﾟ●)ﾉ",
							"往裡頭抹一抹 抹出一杯抹茶 抹走肚子的憂傷",
							"向裡頭擠一擠 擠出一杯奶茶 擠出肚子的太陽",
							"朝裡頭種一種 種出一杯烏龍茶 種成小宇宙的故鄉~ ヾ(●ﾟ∀ﾟ●)ﾉ",
							"熱水~ 熱水~ 熱水~唷 熱水乎拉拉 好喝的熱水唷~ ヾ(●ﾟ∀ﾟ●)ﾉ",
							"熱水~ 熱水~ 熱水~唷 熱水花拉拉 魔法的熱水唷~ ヾ(●ﾟ∀ﾟ●)ﾉ",
						}
						out = strings.Join(lines, "\r\n")
					}
					if strings.Contains(inText, "霸") {
						out = "1" 
						img = "https://i.imgur.com/ozCoxCe.png"
					}
					if strings.Contains(inText, "噴水") {
						out = "1" 
						img = "https://i.imgur.com/fKggfcl.jpg"
					}//
					if strings.Contains(inText, "香菇") || strings.Contains(inText, "吃菇"){
						out = "1" 
						img = "https://i.imgur.com/7Vd5IEf.jpg"
					}
					// magic Number
					var magicN = rand.Intn(10)
					if (magicN == 1) {
						 out = "1"
						 img = "https://i.imgur.com/k271IEO.png"
					}
					if strings.Contains(inText, "help") {
						lines := []string {	
							"此為小宇宙機器人 beta 實驗機",
							"輸入關鍵字可以得到答覆，可以詢問運勢天氣等各種物天",
							"請輸入 '/cs + 關鍵字' 現在關鍵字如下",
							"今天運勢",
							"胖",
							"吃什麼",
							"天然",
							"小宇宙",
							"颱風",
							"買",
							"肚子痛",
						}
						out = strings.Join(lines, "\r\n")
					}
					if out == "" {
						noWay := []string {
							"小宇宙的話，比這個更厲害哦～","臣妾作不到啊啊啊啊啊啊啊啊啊啊啊",
						}
						out = noWay[rand.Intn(len(noWay))]
					}
				
					if out != "1" {
						log.Println(message.Text)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(img,img)).Do(); err != nil {
						log.Print(err)
						}
					}
					
				}
				
				// cs end
				// 占卜機器人
				if strings.Contains(inText, "/123"){
					// for lucky
					if strings.Contains(inText, "運勢") {
						lucky := []string{
						"大吉","中吉","小吉","吉","末吉","凶","大凶",
						}
						out = lucky[rand.Intn(len(lucky))]
					}
					if strings.Contains(inText, "吃什麼")||strings.Contains(inText, "吃") {
						eat := []string{
							"飯飯","麵麵","打邊爐","炸豬排","燒烤","烤雞翅","蒸蛋",
						}
						out = eat[rand.Intn(len(eat))]
					}
					if strings.Contains(inText, "漢納") || strings.Contains(inText, "憨那")   {
						out = "你是天才小仙女！！🧚🏻‍♀"
					}
					
					log.Println(message.Text)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
							log.Print(err)
					}
				
				}
				
				
				
				
				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.google.com.tw/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png","https://www.google.com.tw/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png")).Do(); err != nil {
				//	log.Print(err)
				//}




			}
		}
	}
	// here began
	
	
}
