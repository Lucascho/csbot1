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
				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
				//	log.Print(err)
				//}
				log.Println(message.Text)
				var out = ""
				inText := strings.ToLower(message.Text)
				if strings.Contains(inText, "狗") || strings.Contains(inText, "dog") {
					out = "汪"
				}
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "今天運勢") || strings.Contains(inText, "今日運勢") {
					lucky := []string{
					"大吉","吉","小吉","凶",
					}
					out = lucky[rand.Intn(len(lucky))]
				}
				if strings.Contains(inText, "/cs") && ( strings.Contains(inText, "胖") || strings.Contains(inText, "月半")) {
					out = "小宇宙才不胖！！！"
				}
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "圓")  {
					out = "小宇宙才不圓！！！"
				}
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "買嗎"){
					buy := []string{
					"考慮一下","買！！！","等等","不要",
					}
					out = buy[rand.Intn(len(buy))]
				}
				
				if strings.Contains(inText, "/cs") && (strings.Contains(inText, "吃什麼")||strings.Contains(inText, "吃")){
					eat := []string{
					"飯飯","麵麵","打邊爐","炸豬排","燒烤",
					}
					out = eat[rand.Intn(len(eat))]
				}
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "小宇宙")  {
					out = "小宇宙萌萌達~~"
				}
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "燒")  {
					out = "燒～～～"
				}
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "天然")  {
					out = "你才天然，你全家都天然"
				}
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "天氣")  {
					weather := []string{
					"下大雨唷","招喚颱風唷","長香菇唷","飄雨","打雷你唷",
					}
					if strings.Contains(inText, "今天") || strings.Contains(inText, "明天") || strings.Contains(inText, "後天"){
						out = strings.Trim(out, "/cs")
						out += message.Text+"會"		
					}
					out += weather[rand.Intn(len(weather))]
				}
				//小宇宙的自白部分///
				//天才 小宇宙超天才！
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "天才")  {
					out = "小宇宙超天才！"
				}
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "黛玉")  {
					out = "黛玉 小宇宙是小黛玉！！"
				}
				//聰明 小宇宙超聰明！
				if strings.Contains(inText, "/cs") && strings.Contains(inText, "聰明")  {
					out = "聰明 小宇宙超聰明！"
				}
				
				log.Println(message.Text)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
				}


			}
		}
	}
	// here began
	
	
}
