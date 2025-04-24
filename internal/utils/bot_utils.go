package utils

import (
    tele "gopkg.in/telebot.v4"
    "bot/internal/api"
    "time"
    "fmt"
    "log"
    "strconv"
    "strings"
    "regexp"
)

type User struct{
    IND int
    PUUID,RES,REG,GAME,TAG,NAME,VAL_BOOL,STATE,Region,Shard string
    ALL_TEXT []string
    VAL_ACTI []bool
    SEEN map[string]bool
    HISTORY []string
}

var users = make(map[int64]*User)
func GetUser(chatID int64) *User {
    if user, exists := users[chatID]; exists {
        return user
    }

    users[chatID] = &User{
		STATE: "",
		SEEN: make(map[string]bool), 
		HISTORY: []string{},
	}
	return users[chatID]
}

func SetUserState(user *User, newState string) {
    if user.STATE != newState {
        user.HISTORY = append(user.HISTORY, user.STATE)
        user.STATE = newState
    }
}

func GoBack(user *User) {
    if len(user.HISTORY) < 1 {
        log.Println("Нет предыдущего состояния для возврата.")
        return
    }
    
    prevState := user.HISTORY[len(user.HISTORY)-1]
    user.HISTORY = user.HISTORY[:len(user.HISTORY)-1]
    user.STATE = prevState

}

func DelMess(c tele.Context)error{
    time.Sleep(5 * time.Second)
    return c.Delete()
}

func RegChoo(c tele.Context)error{
    clBk := strings.TrimSpace(c.Callback().Data)
    data := GetUser(c.Chat().ID)
    switch clBk {
        case "btn1":
            data.REG = "americas"
        case "btn2":
            data.REG = "asia"
        case "btn3":
            data.REG = "europe"
    }
    SetUserState(data, "region_choo")
    return c.Edit("Выберите страну:", MenuRegion2())
}

func RegChoo2(c tele.Context)error{
    clBk := strings.TrimSpace(c.Callback().Data)
    data := GetUser(c.Chat().ID)
    data.Region = clBk
    SetUserState(data, "shard_choo")
    return c.Edit("Выберитк свой шард:",MenuValShard())
}

func ShardChoo(c tele.Context) error {
    clBk := strings.TrimSpace(c.Callback().Data)
    data := GetUser(c.Chat().ID)
    data.Shard = clBk
    SetUserState(data, "name_choo")
    return c.Edit("Введите свой никнейм:")
}

func NameChoo(c tele.Context, text string) error {
    data := GetUser(c.Chat().ID)
    data.NAME = text
    SetUserState(data, "tag_choo")
    return c.Send("Введите свой тег:")
}

func TagChoo(c tele.Context,text string)error{
    var resp string
    data := GetUser(c.Chat().ID)
    data.TAG = text
    puuid := api.GetPuuid(data.REG,data.NAME,data.TAG)
    SetUserState(data, "del")
    if puuid != "" {
        resp = fmt.Sprintf("Ваш puuid успешно получен!\n%s", data.PUUID)
    }else {
        resp = fmt.Sprintf("Имя: %s\nТег: %s\nРегион: %s",data.NAME,data.TAG,data.REG)
    }
    data.PUUID = puuid
    return c.Send(resp)
}

func GameChoo(c tele.Context) error {
    data := GetUser(c.Chat().ID)
    menu, _ := MenuGame()
    SetUserState(data, "game_choo")
    return c.Edit("Выберите игру", menu)
}

func GameAct(c tele.Context) error {
    data := GetUser(c.Chat().ID)
    SetUserState(data, "game_act")
    menu1, _ := MenuVal()
    menu2, _ := MenuLol()
    clBk := strings.TrimSpace(c.Callback().Data)
    if clBk == "btn_val" {
        SetUserState(data, "val_actions") 
        return c.Edit("Выберите действие: ",menu1)
    }else{
        SetUserState(data, "lol_actions")
        return c.Edit("Выберите действие: ",menu2)
    }
    return nil
}

func ValAction(c tele.Context)error{
    cnt_menu, _ := MenuValCnt()
    chatId := c.Chat().ID
    data := GetUser(chatId)
    clBk := strings.TrimSpace(c.Callback().Data)
    response, err := api.GetValContent(data.Shard)
    if err != nil {
        log.Println("Ошибка получения контента:", err)
        return err
    }
    serv_response, err := api.GetValStatus(data.Shard)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    rank_response, err := api.ValRanked(data.Shard)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    switch clBk {
        case "btn_val2":
            log.Println("GameAct called") 
            SetUserState(data, "val_cnt")
            return c.Edit(fmt.Sprintf("версия игры: %s", response.VERS), cnt_menu)
        case "btn_val3":
            SetUserState(data,"val_serv")
           if(len(serv_response.INC) > 0){
                return c.Edit(fmt.Sprintf("Инцедент в %s, серьезности  %s\n Описание: %s\n Обновленно в %s сейчас %s\n Описание: %s",serv_response.INC[0].CR, serv_response.INC[0].SEVR, serv_response.INC[0].TIT[0].CNT, serv_response.INC[0].UPD[0].UP, serv_response.INC[0].STAT, serv_response.INC[0].UPD[0].CNT),MenuRet())
            }else{
                return c.Edit("С серверами все впорядке!",MenuRet())
            }
        case "btn_val1":
            var resp string
            for _, rsp := range rank_response.Players {
                if(rsp.Puuid == data.PUUID){
                    resp = fmt.Sprintf("Число побед: %d\nРанк в таблице лидеров: %d\nРпнк в рановом режиме: %d", rsp.Wins,rsp.LeadRank,rsp.RankedRank)
                }
            }
            SetUserState(data,"val_player")
            return c.Edit(resp, ValPlayer())
        case "btn_val4":
            SetUserState(data,"val_ranking")
            return c.Edit("В разработке", MenuRet())
    }
    return nil
}

func ValPlr(c tele.Context)error{
  //  data := GetUser(c.Chat().ID)
    clBk := strings.TrimSpace(c.Callback().Data)

    switch clBk {
        case "pl_btn1":
            return c.Edit("В разработке", MenuRet())
        case "pl_btn2":
            return c.Edit("В разработке", MenuRet())
    }
    
    return nil
}

func ValCnt(c tele.Context)error{
    
    chatId := c.Chat().ID
    data := GetUser(chatId)
    response, err := api.GetValContent(data.Shard)
    
    if err != nil {
        log.Println("Ошибка получения контента:", err)
        return err
    }
    clBk := strings.TrimSpace(c.Callback().Data)
    switch clBk {
        case "val_cnt1":
            for _, item := range response.CHAR {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
        case "val_cnt2":
            for _, item := range response.MAP {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
        case "val_cnt4":
            for _, item := range response.EQUIP {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
        case "val_cnt5":
            for _, item := range response.GAME_MODE {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
    }
    
    data.SEEN = make(map[string]bool)
    
        for i, tex := range data.ALL_TEXT {
            if(tex == "Null UI Data!"){
                tex = "Тестовый"
            }
            data.RES += fmt.Sprintf("\n%d. %s ", i+1, tex)
        }
        return c.Edit(fmt.Sprintf("Общее количество: %d %s\n",len(data.ALL_TEXT), data.RES), MenuRet())
    
    return nil

}

func LolActions(c tele.Context)error{
    data := GetUser(c.Chat().ID) 
    menu, _ := MenuLol()
    SetUserState(data, "lol_actions")
    return c.Edit("Выберите действие", menu)
}

func LolAction(c tele.Context)error{
    
    data := GetUser(c.Chat().ID)
    serv_response, err := api.LolStatus(data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    score_response, err := api.LolMasteryScore(data.PUUID,data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    user_response, err := api.LolSummoner(data.PUUID,data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    usr_menu, _ := LolUserMenu()
    chmp_menu, _ := LolChaempMenu()
    tours_menu, _ := LolToursMenu()
    clBk := strings.TrimSpace(c.Callback().Data)
    switch clBk {
        case "user_btn":
            SetUserState(data, "lol_user")
            return c.Edit(fmt.Sprintf("Ваш уровень: %d\nПоследняя активность в %s", user_response.LVL, time.Unix(user_response.TIME, 0).Format("02.01.2006 15:04:05")), usr_menu)
        case "server_btn":
            SetUserState(data,"lol_serv")
            return c.Edit(fmt.Sprintf("Инцедент в %s, серьезности  %s\n Описание: %s\n Обновленно в %s сейчас %s\n Описание: %s",serv_response.INC[0].CR, serv_response.INC[0].SEVR, serv_response.INC[0].TIT[0].CNT, serv_response.INC[0].UPD[0].UP, serv_response.INC[0].STAT, serv_response.INC[0].UPD[0].CNT),MenuRet())
        case "champ_btn":
            SetUserState(data, "lol_chemp")
            return c.Edit(fmt.Sprintf("Общее мастерство: %d",score_response),chmp_menu)
        case "tours_btn":
            SetUserState(data, "lol_tours")
            return c.Edit("Информация о clash режиме", tours_menu)
    }
    
    return nil
    
}

var buts []tele.InlineButton

func LolUser(c tele.Context)error{
    data := GetUser(c.Chat().ID)
    clBk := strings.TrimSpace(c.Callback().Data)
    leag_response, err := api.LolLeagPu(data.PUUID,data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    matchs_response, err := api.LolMatchs(data.PUUID,data.REG)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    switch clBk {
        case "league_btn":
            
            row := []tele.InlineButton{}
            menu := &tele.ReplyMarkup{}
            board := [][]tele.InlineButton{}
                
            for i, _ := range leag_response{
                btn := tele.InlineButton{
                    Text: fmt.Sprintf("Лига %d",i+1), Unique: fmt.Sprintf("%d", i),
                }
                row = append(row, btn)
                    
                if len(row) >= 3 || i == len(leag_response)-1 {
                    board = append(board, row)
                    row = []tele.InlineButton{}
                }
                    
            }
                
            menu.InlineKeyboard = board
            SetUserState(data, "lol_leag")
            return c.Edit("Выберите Лигу:",menu)
        case "gameTime_btn":
            
            var match_time, game_time int64 = 0, 0
            for _, match := range matchs_response {
                match_response, err := api.LolMatchPu(match, data.REG)
                if err != nil {
                    log.Println("Ошибка получения статуса:", err)
                    return err
                }
                mtch := match_response.Info.EndTime - match_response.Info.StrtTime
                time := match_response.Info.EndTime - match_response.Info.CrTime
                match_time += mtch
                game_time += time
            }
                
            mtchTime_duration := time.Duration(match_time) * time.Millisecond
            gameTime_duration := time.Duration(game_time) * time.Millisecond
                
            SetUserState(data,"lol_gameTime")
            return c.Edit(fmt.Sprintf("Среднее время в игре: %d,%d часов\nИгровое время в матчах: %d,%d часов", int(mtchTime_duration.Hours()), int(mtchTime_duration.Minutes())%60, int(gameTime_duration.Hours()), int(gameTime_duration.Minutes())%60, ),MenuRet())
                
        case "last_matches":
                
            row := []tele.InlineButton{}
            menu := &tele.ReplyMarkup{}
            board := [][]tele.InlineButton{}
                
            for i, resp := range matchs_response{
                btn := tele.InlineButton{
                    Text: fmt.Sprintf("Матч %d",i+1),
                        Unique: resp,
                }
                row = append(row, btn)
                    
                if len(row) >= 3 || i == len(matchs_response)-1 {
                    board = append(board, row)
                    row = []tele.InlineButton{}
                }
                    
            }
            SetUserState(data, "lol_match")
            menu.InlineKeyboard = board
            return c.Edit("Выберите Матч:",menu)
    }
    
    return nil
}

func LolLeag(c tele.Context)error{
    
    id, _ := strconv.Atoi(strings.TrimSpace(c.Callback().Data))
    data := GetUser(c.Chat().ID)
    leag_response, err := api.LolLeagPu(data.PUUID,data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    var isWin string 
    
    for i, r := range leag_response[id].Mini.Progress {
        if strings.Contains(string(r), "W") {
            isWin += fmt.Sprintf("%d. Победа\n", i+1)
        } else {
            isWin += fmt.Sprintf("%d. Поражение\n", i+1)
        }
    }
    
    return c.Edit(
        fmt.Sprintf(
           "\nВаш ранг: %s %s\nПоинты лиги: %d\nЧисло побед: %d\nЧисло поражений: %d\nИгровой  режим: %s\nПобеды за ранг %s: %d\nПоражения за ранг %s: %d\nПрогресс ранга: \n%s\nКоличетсво побед до следующего ранга: %d",
           leag_response[id].RANK,leag_response[id].TIER,
           leag_response[id].LEAPT,
           leag_response[id].WNS,leag_response[id].LSES,
           leag_response[id].Mode, leag_response[id].RANK,
           leag_response[id].Mini.Wnins,
           leag_response[id].RANK,
           leag_response[id].Mini.Losses,
           isWin,leag_response[id].Mini.Target,
        ),MenuRet(),
    )
    
}

func LolLastMatch(c tele.Context)error{
    
    data := GetUser(c.Chat().ID)
    matchId := strings.TrimSpace(c.Callback().Data)
    resp_mtch, err := api.LolMatchPu(matchId, data.REG)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    resp := resp_mtch.Info
    
    var win,lost string
    
    for _, v := range resp.Teams {
        if(v.IsWin == true && v.TeamId == 100){
            win = "Синяя"
            lost = "Красная"
        }else{
            win = "Красная"
            lost = "Синяя"
        }
    }
    
    return c.Edit(
        fmt.Sprintf(
            "Игра началась: %d\nИгра закончилась: %d\nРежим игры: %s, %s\nПобедила комманда: %s\nПроиграла комманда: %s", resp.StrtTime, resp.EndTime, resp.GmMode, resp.GmType,win,lost,
            ), MenuRet(),
    )
    
}

func LolChemp(c tele.Context) error {
    clBk := strings.TrimSpace(c.Callback().Data)
    data := GetUser(c.Chat().ID)
    rotat_resp, err := api.LolRotat(data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    chemp_resp, err := api.LolChemps(data.PUUID,data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    ddragon_resp, err := api.LolDdragon()
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
        switch clBk {
            case "ratation_btn":
                var frNChmp, frChmp string
                var frNChmpList, frChmpList []string
        for _, chmp := range ddragon_resp {
	       for _, chmp1 := range rotat_resp.ChempNew {
	           key1, _ := strconv.Atoi(chmp.Key)
		      if key1 == chmp1 {
			     frNChmpList = append(frNChmpList, chmp.Name)
		      }
	       }
	        for _, chmp2 := range rotat_resp.Chemps {
	            key2, _ := strconv.Atoi(chmp.Key)
		        if key2 == chmp2 {
			        frChmpList = append(frChmpList, chmp.Name)
		        }
	        }
        }

            frNChmp = strings.Join(frNChmpList, ", ")
            frChmp = strings.Join(frChmpList, ", ")
                
                return c.Edit(
                    fmt.Sprintf(
                        "Мксимальный уровень нового игрока: %d\nБесплатные чемпионы для новых игроков: %s\nБесплатные чемпионы для всех: %s",rotat_resp.Lvl,frNChmp,frChmp,
                    ), MenuRet(),
                )
            case "allChemps_btn":
                
                row := []tele.InlineButton{}
                menu := &tele.ReplyMarkup{}
                board := [][]tele.InlineButton{}
                
                for _, resp := range ddragon_resp {
                    btn := tele.InlineButton{
                        Text: resp.Name,
                        Unique: resp.Id,
                    }
                    row = append(row, btn)
                    
                    if len(row) >= 3{
                    //|| i == len(ddragon_resp) - 1 {
                        board = append(board, row)
                        row = []tele.InlineButton{}
                    }
                    
                }
                
                menu.InlineKeyboard = board
                SetUserState(data, "lol_chempion")
                return c.Edit("Выберите чемпиона: ",menu)
                
            case "myChemps_btn":
                
                row := []tele.InlineButton{}
                menu := &tele.ReplyMarkup{}
                board := [][]tele.InlineButton{}
                
                for _, drgn := range ddragon_resp{
                    
                    for i, resp := range chemp_resp {
                        key, _ := strconv.ParseInt(drgn.Key, 10, 64)
                        if(resp.ChmpId == key){
                            btn := tele.InlineButton{
                                Text: drgn.Name,
                                Unique: fmt.Sprintf("%s%d",drgn.Id, i,),
                            }
                            row = append(row, btn)
                    
                        if len(row) >= 4 || i == len(chemp_resp)-1 {
                            board = append(board, row)
                            row = []tele.InlineButton{}
                        }
                            
                        if len(row) > 0 {
                            board = append(board, row)
                        }
                            
                        }
                    
                    }
                }
                
                SetUserState(data, "lol_mychemps")
                menu.InlineKeyboard = board
                return c.Edit("Выберите чемпиона: ",menu)
            }
    
    return nil
    
}

func LolMyChemps(c tele.Context) error {
    data := GetUser(c.Chat().ID)
    var id string
    var ind int
    
    callBack := c.Callback().Data
    numReg := regexp.MustCompile(`\d+`)
    strReg := regexp.MustCompile(`[a-zA-Z]+`)
    
    num_matches := numReg.FindAllString(callBack, -1)
    str_matches := strReg.FindAllString(callBack, -1)
    
    for _, m := range num_matches {
        //num, _ := strings.TrimSpace(strconv.Atoi(m))
		num, _ := strconv.Atoi(m)
		ind = num
	}
	
	for _, m := range str_matches {
		//str, _ := strings.TrimSpace(strconv.Atoi(m))
		id = m
	}
    
    resp_leag, err := api.LolDdragon()
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    resp := resp_leag[id]
    
    resp_c, err := api.LolChemps(data.PUUID,data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    chemps_resp := resp_c[ind]
    
    icon_resp, err := api.LolChempIcon(resp.Id)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    tags := strings.Join(resp.Tags, ", ")
    play_tm := time.Unix(chemps_resp.ChempPlTime, 0)
    time := fmt.Sprintf(
        "%d:%d", play_tm.Hour(), play_tm.Minute(),
    )
    photo := &tele.Photo{
        File: tele.FromURL(icon_resp),
        Caption: fmt.Sprintf(
            "Имя: %s\nОписание: %s\nТип ресурса: %s\nТеги: %s\nКраткий лор: %s\nПоинты чемпиона: %d\nУровень чемпиона: %d\nНаигранные часы: %s\nПоинты до след уровня: %d\nАтака(от 1 до 10): %d\nЗащита(от 1 до 10): %d\nМагия(от 1 до 10): %d\nСложность(от 1 до 10): %d\nСтатистика чемпиона:\n\n\tЗдоровье: %f (за +1ур +%f)\n\tКоличество маны: %f (за +1ур +%f)\n\tСкорость передвижения: %f\n\tБроня: %f (за +1ур +%f)\n\tМагич. сопротивление: %f (за +1ур +%f)\n\tДальность атаки: %f\n\tРегенерация здоровья: %f (за +1ур +%f)\n\tРегенерация маны: %f (за +1ур +%f)\n\tУрон от атаки: %f (за +1ур +%f)\n\tСкорость атаки: %f (за +1ур +%f)\n\tШанс критического урона: %f%% (за +1ур +%f%%)\n", resp.Name,resp.Title,resp.Type,tags,resp.Lore,chemps_resp.ChmpPt,chemps_resp.ChempLvl,time,chemps_resp.ChmpPtUL, resp.Info.Atk,resp.Info.Def,resp.Info.Mag,resp.Info.Dif,resp.Stats.Hp,resp.Stats.HpLvl,resp.Stats.Mp,resp.Stats.MpLvl,resp.Stats.Speed,resp.Stats.Armor,resp.Stats.ArmorLvl,resp.Stats.SpellBlock,resp.Stats.SpellBlockLvl,resp.Stats.AtkRng,resp.Stats.HpReg,resp.Stats.HpRegLvl,resp.Stats.MpReg,resp.Stats.MpRegLvl,resp.Stats.AtkDmg,resp.Stats.AtkDmgLvl,resp.Stats.AtkSpeed,resp.Stats.AtkSpeedLvl,resp.Stats.Crit,resp.Stats.CritLvl,
        ),
    }
    
    return c.Edit(photo, MenuRet())
    
}

func LolAllChemps(c tele.Context) error {
    id := strings.TrimSpace(c.Callback().Data)
    resp_ddragon, err := api.LolDdragon()
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    resp := resp_ddragon[id]
    icon_resp, err := api.LolChempIcon(resp.Id)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    tags := strings.Join(resp.Tags, ", ")
    photo := &tele.Photo{
        File: tele.FromURL(icon_resp),
        Caption: fmt.Sprintf(
            "Имя: %s\nОписание: %s\nТип ресурса: %s\nТеги: %s\nКраткий лор: %s\nАтака(от 1 до 10): %d\nЗащита(от 1 до 10): %d\nМагия(от 1 до 10): %d\nСложность(от 1 до 10): %d\nСтатистика чемпиона:\n\n\tЗдоровье: %f (за +1ур +%f)\n\tКоличество маны: %f (за +1ур +%f)\n\tСкорость передвижения: %f\n\tБроня: %f (за +1ур +%f)\n\tМагич. сопротивление: %f (за +1ур +%f)\n\tДальность атаки: %f\n\tРегенерация здоровья: %f (за +1ур +%f)\n\tРегенерация маны: %f (за +1ур +%f)\n\tУрон от атаки: %f (за +1ур +%f)\n\tСкорость атаки: %f (за +1ур +%f)\n\tШанс критического урона: %f%% (за +1ур +%f%%)\n", resp.Name,resp.Title,resp.Type,tags,resp.Lore,resp.Info.Atk,resp.Info.Def,resp.Info.Mag,resp.Info.Dif,resp.Stats.Hp,resp.Stats.HpLvl,resp.Stats.Mp,resp.Stats.MpLvl,resp.Stats.Speed,resp.Stats.Armor,resp.Stats.ArmorLvl,resp.Stats.SpellBlock,resp.Stats.SpellBlockLvl,resp.Stats.AtkRng,resp.Stats.HpReg,resp.Stats.HpRegLvl,resp.Stats.MpReg,resp.Stats.MpRegLvl,resp.Stats.AtkDmg,resp.Stats.AtkDmgLvl,resp.Stats.AtkSpeed,resp.Stats.AtkSpeedLvl,resp.Stats.Crit,resp.Stats.CritLvl,
        ),
    }
    
    return c.Edit(photo,MenuRet())
    
}

func LolTours(c tele.Context)error {
    
    data := GetUser(c.Chat().ID)
    resp_clash, err := api.LolClash(data.PUUID,data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    clash_resp := resp_clash[0]
    team_resp, err := api.LolClashTeam(clash_resp.TMID,data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    tours_resp, err := api.LolTours(data.Region)
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    clBk := strings.TrimSpace(c.Callback().Data)
    
        switch clBk {
            case "tours_btn":
                row := []tele.InlineButton{}
                menu := &tele.ReplyMarkup{}
                board := [][]tele.InlineButton{}
                
                for i, resp := range tours_resp {
                    
                    btn := tele.InlineButton{
                        Text: resp.Name,
                        Unique: fmt.Sprintf("%d",i),
                    }
                    
                    row = append(row, btn)
                    
                    if len(row) >= 2 || i == len(tours_resp)-1 {
                        board = append(board, row)
                        row = []tele.InlineButton{}
                    }
                    
                }    
                menu.InlineKeyboard = board
                SetUserState(data, "lol_tour")
                return c.Edit("Выберите соревнование: ",menu)
            case "myTeam_btn":
                var plrs string
                for i, str := range team_resp.Plyrs {
                    plrs += fmt.Sprintf(
                        "\n\t%d Игрок\n\tПозиция игрока: %s\n\tРоль игрока: %s\n\n",i,str.PST,str.RLE,
                        )
                }
                return c.Edit(
                    fmt.Sprintf(
                        "Ваша позиция: %s\nВаша роль: %s\nИмя команды: %s\nТир комманды: %d\nКапитан команды: %s\nАббревиатура: %s\nИгроки:\n\n",clash_resp.PST,clash_resp.RLE, team_resp.Name,team_resp.Tier,team_resp.Capitan,team_resp.Abbrev,plrs,
                    ),MenuRet(),
                )
        }
    
    return nil
    
}

func LolTour(c tele.Context)error{
    data := GetUser(c.Chat().ID)
    id, _ := strconv.Atoi(strings.TrimSpace(c.Callback().Data))
    resp, err := api.LolTours(data.Region)
    tours_resp := resp[id]
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    var canc string
    if(tours_resp.Shed[0].CNC == true){
        canc = "Отменен\n"
    }else{
        canc = "\n"
    }
    
    return c.Edit(
        fmt.Sprintf(
            "Имя турнира: %s %s\nЗарегестрирован в %s\nНачинается в %s\n%s",tours_resp.Name,tours_resp.Day,tours_resp.Shed[0].REG, tours_resp.Shed[0].STRT,canc,
        ),MenuRet(),
    )
    
}