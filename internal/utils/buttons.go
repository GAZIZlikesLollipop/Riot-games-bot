package utils

import tele "gopkg.in/telebot.v4"

var BackBtn = tele.InlineButton{Text: "Назад", Unique: "back_btn"}

func MenuRegion1() (*tele.ReplyMarkup, []tele.InlineButton) {
    menu := &tele.ReplyMarkup{}
    
    btn1 := tele.InlineButton{
        Text: "americas",
        Unique: "btn1",
    }
    btn2 := tele.InlineButton{
        Text: "asia",
        Unique: "btn2",
    }
    btn3 := tele.InlineButton{
        Text: "europe",
        Unique: "btn3",
    }
    
    menu.InlineKeyboard = [][]tele.InlineButton{
        {btn1},{btn2},
        {btn3},
    }
    
    return menu, []tele.InlineButton{btn1,btn2,btn3}
}

func MenuRegion2() *tele.ReplyMarkup {
    var regionNames = map[string]string{
        "br1":  "Бразилия",
        "eun1": "Европа (Скандинавия и Восток)",
        "euw1": "Европа (Запад)",
        "jp1":  "Япония",
        "kr":   "Корея",
        "la1":  "Северная Латинская Америка",
        "la2":  "Южная Латинская Америка",
        "me1":  "Ближний Восток",
        "na1":  "Северная Америка",
        "oc1":  "Океания",
        "ru":   "Россия",
        "sg2":  "Сингапур",
        "tr1":  "Турция",
        "tw2":  "Тайвань",
        "vn2":  "Вьетнам",
    }

    markup := &tele.ReplyMarkup{}
    var keys [][]tele.InlineButton
    row := []tele.InlineButton{}

    for code, name := range regionNames {
        btn := tele.InlineButton{
            Text:   name,
            Unique: code,
        }

        row = append(row, btn)

        if len(row) == 3 {
            keys = append(keys, row)
            row = []tele.InlineButton{}
        }
    }

    if len(row) > 0 {
        keys = append(keys, row)
    }

    markup.InlineKeyboard = keys
    return markup
}

func MenuValShard() *tele.ReplyMarkup {
    var shardNames = map[string]string{
        "ap":       "Азия и Океания",
        "br":       "Бразилия",
        "esports":  "Киберспорт",
        "eu":       "Европа",
        "kr":       "Корея",
        "latam":    "Латинская Америка",
        "na":       "Северная Америка",
    }

    markup := &tele.ReplyMarkup{}
    var keys [][]tele.InlineButton
    row := []tele.InlineButton{}

    for code, name := range shardNames {
        btn := tele.InlineButton{
            Text:   name,
            Unique: code,
        }

        row = append(row, btn)

        if len(row) == 2 {
            keys = append(keys, row)
            row = []tele.InlineButton{}
        }
    }

    if len(row) > 0 {
        keys = append(keys, row)
    }

    markup.InlineKeyboard = keys
    return markup
}

func MenuGame() (*tele.ReplyMarkup, []tele.InlineButton) {
    btn_val := tele.InlineButton{Text: "Valorant", Unique: "btn_val"}
    btn_lol := tele.InlineButton{Text: "Legague of legends", Unique: "btn_lol"}
    markup := &tele.ReplyMarkup{}
    markup.InlineKeyboard = [][]tele.InlineButton{
        {btn_val},
        {btn_lol},
    }
    return markup, []tele.InlineButton{btn_val,btn_lol}
}

func MenuVal() (*tele.ReplyMarkup, []tele.InlineButton) {
    
    btn_val1 := tele.InlineButton{Text: "Обо мне",Unique: "btn_val1"}
    btn_val2 := tele.InlineButton{Text: "Контент",
        Unique: "btn_val2"}
    btn_val3 := tele.InlineButton{Text: "Сервера",
        Unique: "btn_val3"}
    btn_val4 := tele.InlineButton{Text: "Ранговый режим",
        Unique: "btn_val4"}

    markup_val := &tele.ReplyMarkup{}
    markup_val.InlineKeyboard = [][]tele.InlineButton{
        {btn_val1},{btn_val2},
        {btn_val3},{btn_val4},
        {BackBtn},
    }
    
    return markup_val, []tele.InlineButton{btn_val1, btn_val2, btn_val3, btn_val4}
}

func ValPlayer() *tele.ReplyMarkup {
    btn1 := tele.InlineButton{
        Text: "Последние матчи",
        Unique: "pl_btn1",
    }
    btn2 := tele.InlineButton{
        Text: "Игровое время",
        Unique: "pl_btn2",
    }
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineButton{
        {btn1},
        {btn2},
        {BackBtn},
    }
    
    return menu
    
}

func MenuValCnt() (*tele.ReplyMarkup, []tele.InlineButton) {
    
    val_cnt1 := tele.InlineButton{Text: "Агенты",
        Unique: "val_cnt1"}    
    val_cnt2 := tele.InlineButton{Text: "Карты",
        Unique: "val_cnt2"}    
    val_cnt4 := tele.InlineButton{Text: "Оружия",
        Unique: "val_cnt4"}    
    val_cnt5 := tele.InlineButton{Text: "Игровые режимы", Unique: "val_cnt5"}    
        
    markup := &tele.ReplyMarkup{}
    markup.InlineKeyboard = [][]tele.InlineButton{
        {val_cnt1},{val_cnt2},
        {val_cnt4},{val_cnt5},
        {BackBtn},
    }
    return markup, []tele.InlineButton{val_cnt1,val_cnt2,val_cnt1,val_cnt4,val_cnt5}
}

func MenuRet() *tele.ReplyMarkup {
    del_markup := &tele.ReplyMarkup{}
    
    del_markup.InlineKeyboard = [][]tele.InlineButton{
        {BackBtn},
    }
    return del_markup
}

func MenuMoreSk() (*tele.ReplyMarkup, tele.InlineButton) {
    btn_moreSk := tele.InlineButton{Text: "Больше скинов", Unique: "btn_moreSk"}
    more_skinmk := &tele.ReplyMarkup{}
    more_skinmk.InlineKeyboard = [][]tele.InlineButton{
        {btn_moreSk},
        {BackBtn},
    }
    return more_skinmk, btn_moreSk
}

func MenuLol()(*tele.ReplyMarkup,[]tele.InlineButton){
    
    champ_btn := tele.InlineButton{
        Text: "Чэмпионы", Unique: "champ_btn",
    }
    tours_btn := tele.InlineButton{
        Text: "Турниры и прочее", Unique: "tours_btn",
    }
    server_btn := tele.InlineButton{
        Text: "Сервера", Unique: "server_btn",
    }
    user_btn := tele.InlineButton{
        Text: "Игрок", Unique: "user_btn",
    }
    
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineButton{
        {user_btn},{server_btn},
        {champ_btn},{tours_btn},
        {BackBtn},
    }
    btns := []tele.InlineButton{
        user_btn,server_btn,champ_btn,tours_btn,
    }
    
    return menu, btns
}

func LolUserMenu() (*tele.ReplyMarkup,[]tele.InlineButton) {
    
    league_btn := tele.InlineButton{
        Text: "Моя Лига", Unique: "league_btn",
    }
    gameTime_btn := tele.InlineButton{
        Text: "Игрвое время", Unique: "gameTime_btn",
    }
    last_matches := tele.InlineButton{
        Text: "Последние матчи", Unique: "last_matches",
    }
    
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineButton{
        {gameTime_btn},
        {last_matches},
        {league_btn},
        {BackBtn},
    }
    
    btns := []tele.InlineButton{
        league_btn,gameTime_btn,last_matches,
    }
    
    return menu, btns
    
}

func LolChaempMenu()(*tele.ReplyMarkup,[]tele.InlineButton){
    
    ratation_btn := tele.InlineButton{
        Text: "Бесплатная ротация", Unique: "ratation_btn",
    }
    myChemps_btn := tele.InlineButton{
        Text: "Мои чэмпионы", Unique: "myChemps_btn",
    }
    allChemps_btn := tele.InlineButton{
        Text: "Чэмпионы в игре", Unique: "allChemps_btn",
    }
    
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineButton{
        {myChemps_btn},
        {allChemps_btn},
        {ratation_btn},
        {BackBtn},
    }
    
    btns := []tele.InlineButton{
        allChemps_btn,myChemps_btn,ratation_btn,
    }
    
    return menu, btns
}

func LolToursMenu()(*tele.ReplyMarkup,[]tele.InlineButton){
    
    myTeam_btn := tele.InlineButton{
        Text: "Моя команда", Unique: "myTeam_btn",
    }
    tours_btn := tele.InlineButton{
        Text: "Clash соревнования", Unique: "tours_btn",
    }
    
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineButton{
        {myTeam_btn},{tours_btn},
        {BackBtn},
    }
    
    btns := []tele.InlineButton{
        tours_btn,myTeam_btn,
    }
    
    return menu, btns
    
}