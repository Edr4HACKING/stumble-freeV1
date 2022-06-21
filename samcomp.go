package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "strings"
)

func request(url string) {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(body)
}

func input(text string) string {
    fmt.Print(text)

    // var then variable name then variable type
    var first string

    // Taking input from user
    fmt.Scanln(&first)

    return first
}

func parseResponse(text string) {
    trophy := strings.Split(text, `"SkillRating":`)[1]
    trophy = strings.Split(trophy, ",")[0]
    crown := strings.Split(text, `"Crowns":`)[1]
    crown = strings.Split(crown, ",")[0]
    username := strings.Split(text, `"Username":`)[1]
    username = strings.Split(username, ",")[0]

    fmt.Printf("\rNick ACC#SC: %s - Trophy #SC: %s - Crown #SC: %s", username, trophy, crown)
}

func requestWithHeader(url string, token string) {
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)

    req.Header.Set("Host", "kitkabackend.eastus.cloudapp.azure.com:5010")
    req.Header.Set("use_response_compression", "true")
    req.Header.Set("authorization", token)

    res, _ := client.Do(req)
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
    }

    if res.StatusCode == 200 {
        parseResponse(string(body))
    }
    if res.StatusCode == 401 {
      fmt.Printf("\n[EROR] KEY AUTH TELAH KADALUWARSA ATAU ILANG SILAKAN PLAY GAME 3 ROUND AGAIN !!")
    }
	if res.StatusCode == 403 {
  print("\033[0m");
  print("[BANNED] MAMPUS AKUN KEBANNED HAHA\n");
  print("\033[0;31m");
	}
	if res.StatusCode == 501 {
		fmt.Printf("\n[501] Failed")
	}
}

func main() {
  print("\033[0;31m");
  print("###################################\n");
  print("\033[0m");
  print("STUMBLE GAY BOT CHROWN & THROPHY\n");
  print("\033[0;31m");
  print("###################################\n");
  print("\033[0m");
  print("\n");
  print("\033[0;32m");
  print("#####################################\n");
  print("\033[0m");
  print("[IG] @mistersm_0                    #\n")
  print("[FB} Syamsul Muarif                 #\n")    
  print("[TT] @onlysyw                       #\n");
  print("\033[0;32m");
  print("#####################################\n");
  print("\033[0m");
  
  
	round := input("\n\n1. Lose ROUND 1 (JANGAN DI JALANKAN!!) \n2. Win ROUND 1\n3. Win ROUND 2\n4. Win ROUND 3\nInput nomor: ")
    token := input("Input authorization: ")
    thread := input("Threads (1 Direkomendasikan) :")
    thrd, _ := strconv.Atoi(thread)
    url := "http://kitkabackend.eastus.cloudapp.azure.com:5010/round/finishv2/" + round

    for i := 0; i < thrd; i++ {
        go func() {
            for {
                requestWithHeader(url, token)
            }
        }()
    }
    for {
        requestWithHeader(url, token)
    }
}
