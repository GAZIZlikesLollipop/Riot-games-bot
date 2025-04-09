package utils

import tele "gopkg.in/telebot.v4"

var BackBtn = tele.InlineButton{Text: "Назад", Unique: "back_btn"}

func MenuRegion1() (*tele.ReplyMarkup, []tele.Btn) {
    menu := &tele.ReplyMarkup{}
    btn1 := menu.Text("americas")
    btn2 := menu.Text("asia")
    btn3 := menu.Text("europe")
    btn4 := menu.Text("espots")
    menu.Reply(menu.Row(btn1, btn2),menu.Row(btn3, btn4))
    menu.ResizeKeyboard = true
    return menu, []tele.Btn{btn1,btn2,btn3,btn4}
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
    
    btn_val1 := tele.InlineButton{Text: "Контент",
        Unique: "btn_val1"}
    btn_val2 := tele.InlineButton{Text: "Сервера",
        Unique: "btn_val2"}
    btn_val3 := tele.InlineButton{Text: "Матчи",
        Unique: "btn_val3"}
    btn_val4 := tele.InlineButton{Text: "Игровое время",Unique: "btn_val4"}
        
    markup_val := &tele.ReplyMarkup{}
    markup_val.InlineKeyboard = [][]tele.InlineButton{
        {btn_val1},{btn_val2},
        {btn_val3},{btn_val4},
        {BackBtn},
    }
    
    return markup_val, []tele.InlineButton{btn_val1, btn_val2, btn_val3, btn_val4}
}

func MenuValCnt() (*tele.ReplyMarkup, []tele.InlineButton) {
    val_cnt1 := tele.InlineButton{Text: "Агенты",
        Unique: "val_cnt1"}    
    val_cnt2 := tele.InlineButton{Text: "Карты",
        Unique: "val_cnt2"}    
    val_cnt3 := tele.InlineButton{Text: "Скины",
        Unique: "val_cnt3"}    
    val_cnt4 := tele.InlineButton{Text: "Оружия",
        Unique: "val_cnt4"}    
    val_cnt5 := tele.InlineButton{Text: "Игровые режимы",
        Unique: "val_cnt5"}    
    val_cnt6 := tele.InlineButton{Text: "Сезоны",
        Unique: "val_cnt6"}
        
    markup := &tele.ReplyMarkup{}
    markup.InlineKeyboard = [][]tele.InlineButton{
        {val_cnt1},{val_cnt2},{val_cnt3},
        {val_cnt4},{val_cnt5},{val_cnt6},
        {BackBtn},
    }
    return markup, []tele.InlineButton{val_cnt1,val_cnt2,val_cnt3,val_cnt1,val_cnt4,val_cnt5,val_cnt6}
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
        Text: "Чэмпионы", Unique: "champ_btn"
    }
    tours_btn := tele.InlineButton{
        Text: "Турниры и прочее", Unique: "tours_btn"
    }
    server_btn := tele.InlineButton{
        Text: "Сервера", Unique: "server_btn"
    }
    user_btn := tele.InlineButton{
        Text: "Игрок", Unique: "user_btn"
    }
    
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineButton{
        {user_btn},{server_btn},
        {champ_btn},{tours_btn},
        {BackBtn},
    }
    btns := []tele.InlineButton{
        user_btn,server_btn,champ_btn,challanges_btn,
        league_btn,clash_btn
    }
    
    return menu, btns
}

func LolUserMenu() (*tele.ReplyMarkup,[]tele.InlineButton) {
    
    league_btn := tele.InlineButton{
        Text: "Моя Лига", Unique: "league_btn"
    }
    gameTime_btn := tele.InlineButton{
        Text: "Игрвое время", Unique: "gameTime_btn"
    }
    last_matches := tele.InlineButton{
        Text: "Последние матчи", Unique: "last_matches"
    }
    
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineKeyboard{
        {gameTime_btn},
        {last_matches},
        {league_btn},
        {BackBtn},
    }
    
    btns := []tele.InlineButton{
        summoner_btn,gameTime_btn,matchById_btn,userChal_btn,league_btn,clash_btn
    }
    
    return menu, btns
    
}

func LolChaempMenu()(*tele.ReplyMarkup,[]tele.InlineButto){
    
    ratation_btn := tele.InlineButton{
        Text: "Бесплатная ротация", Unique: "ratation_btn"
    }
    allChemps_btn := tele.InlineButton{
        Text: "Мои чэмпионы", Unique: "allChemps_btn"
    }
    
    topChemps_btn := tele.InlineButton{
        Text: "Топ чемпионов по мастерству", Unique: "topChemps_btn"
    }
    
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineKeyboard{
        {ratation_btn},
        {topChemps_btn},
        {allChemps_btn},
        {BackBtn},
    }
    
    btns := []tele.InlineButton{
        chempById_btn,topChemps_btn,allChemps_btn,ratation_btn
    }
    
    return menu, btns
}

func LolToursMenu()(*tele.ReplyMarkup,[]tele.InlineButto){
    clash_btn := tele.InlineButton{
        Text: "Мой Clash", Unique: "clash_btn"
    }
    clashById_btn := tele.InlineButton{
        Text: "Клеш команда", Unique: "clashById_btn"
    }
    tours_btn := tele.InlineButton{
        Text: "Ближайшие соревнования", Unique: "tours_btn"
    }
    
    menu := &tele.ReplyMarkup{}
    menu.InlineKeyboard = [][]tele.InlineKeyboard{
        {clash_btn},
        {clashById_btn},
        {tours_btn},
        {BackBtn},
    }
    
    btns := []tele.InlineButton{
        chempById_btn,topChemps_btn,allChemps_btn,ratation_btn
    }
    
    return menu, btns
}

