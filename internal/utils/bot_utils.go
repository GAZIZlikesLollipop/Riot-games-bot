package utils

import (
    tele "gopkg.in/telebot.v4"
    "bot/internal/api"
    "time"
    "os"
    "fmt"
    "log"
    "strconv"
    "strings"
)

type User struct{
    IND int
    PUUID,RES,REG,GAME,TAG,NAME,VAL_BOOL,STATE string
    ALL_TEXT []string
    VAL_ACTI []bool
    SEEN map[string]bool
    HISTORY []string
}

var users = make(map[int64]*User)
// Функция для получения пользователя
func GetUser(chatID int64) *User {
    if user, exists := users[chatID]; exists {
        return user
    }
    // Если пользователя нет в мапе, создаем нового
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
    if len(user.HISTORY) < 2 {
        log.Println("Нет предыдущего состояния для возврата.")
        return
    }

    prevState := user.HISTORY[len(user.HISTORY)-2]
    user.HISTORY = user.HISTORY[:len(user.HISTORY)-1]
    user.STATE = prevState
}

func DelMess(c tele.Context)error{
    time.Sleep(5 * time.Second)
    return c.Delete()
}

func RegChoo(c tele.Context) error {
    
    chatId := c.Chat().ID
    data := GetUser(chatId)
    
    if data.STATE != "reg_choo" {
        return nil
    }
    
    btnId := c.Callback().Unique
    _, buttons := MenuRegion1()

    for _, btn := range buttons {
        if btnId == btn.Unique {
            data.REG = btn.Text
            SetUserState(data, "name_choo") // Меняем состояние правильно
            return c.Send("Введите свой никнейм:", &tele.ReplyMarkup{RemoveKeyboard: true})
        }
    }
    
    return nil
}

func PuuidText(c tele.Context) error {
    chatId := c.Chat().ID
    data := GetUser(chatId)
    
    text := c.Text()
    
    switch data.STATE {
    case "name_choo":
        data.NAME = text
        SetUserState(data, "tag_choo")
        return c.Send("Введите свой тег:")

    case "tag_choo":
        data.TAG = text
        data.PUUID = api.GetPuuid(data.REG, data.NAME, data.TAG, key)
        SetUserState(data, "del")
        return c.Send(fmt.Sprintf("Ваш PUUID: %s", data.PUUID))
    }

    return nil
}

func GameChoo(c tele.Context) error {
    data := GetUser(c.Chat().ID)
    menu, _ := MenuGame()
    SetUserState(data, "game_act")
    SetUserState(data, "val_actions")
    return c.Edit("Выберите игру", menu)
}

func ValActions(c tele.Context)error{
    data := GetUser(c.Chat().ID) 
    menu1, _ := MenuVal()
    SetUserState(data, "val_action")
    return c.Edit("Выберите действие", menu1)
}

func ValAction(c tele.Context)error{
    cnt_menu, _ := MenuValCnt()
    _, but := MenuVal()
    chatId := c.Chat().ID
    data := GetUser(chatId)
    response, err := api.GetValContent("ap")
    if err != nil {
        log.Println("Ошибка получения контента:", err)
        return err
    }
    serv_response, err := api.GetValStatus("ap")
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    for _, btns := range but {
    switch btns.Unique {
        case "btn_val1":
            log.Println("GameAct called") 
            SetUserState(data, "val_cnt")
            return c.Edit(fmt.Sprintf("версия игры: %s", response.VERS), cnt_menu)
        case "btn_val2":
           if(len(serv_response.INC) > 0){
                return c.Edit(fmt.Sprintf("Инцедент в %s, серьезности  %s\n Описание: %s\n Обновленно в %s сейчас %s\n Описание: %s",serv_response.INC[0].CR, serv_response.INC[0].SEVR, serv_response.INC[0].TIT[0].CNT, serv_response.INC[0].UPD[0].UP, serv_response.INC[0].STAT, serv_response.INC[0].UPD[0].CNT),MenuRet())
            }else{
                return c.Edit("С серверами все впорядке!",MenuRet())
            }
        case "btn_val3":
            return c.Edit("В разработке", MenuRet())
        case "btn_val4":
            return c.Edit("В разработке", MenuRet())
    }
    }
    return nil
}

func ValCnt(c tele.Context)error{
    
    _, btns := MenuValCnt()
    chatId := c.Chat().ID
    data := GetUser(chatId)
    response, err := api.GetValContent("ap")
    
    if err != nil {
        log.Println("Ошибка получения контента:", err)
        return err
    }
    
    for _, btn := range btns {
    buttonID := btn.Unique
    switch buttonID{
        case "val_cnt1":
            for _, item := range response.CHAR {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
        case "val_cnt2":
            for _, item := range response.MAP {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
        case "val_cnt3":
            for _, item := range response.SKIN {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
            
        
        data.RES += "\n\nНажмите на кнопку что бы увидеть больше"
        
        case "val_cnt4":
            for _, item := range response.EQUIP {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
        case "val_cnt5":
            for _, item := range response.GAME_MODE {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
            }
        case "val_cnt6":
            for _, item := range response.ACT {
                data.ALL_TEXT = append(data.ALL_TEXT, item.NAME)
                data.VAL_ACTI = append(data.VAL_ACTI, item.ACTIVE)
            }
    }
    
    data.SEEN = make(map[string]bool)
    
    if buttonID == "val_cnt6" {
        for i, tex := range data.ALL_TEXT{
            if data.VAL_ACTI[i] {
               data.VAL_BOOL = "(Активный)"
            }
            if(tex == "Null UI Data!"){
                tex = "Тестовый"
            }
            if !data.SEEN[tex]{
            data.SEEN[tex] = true
            data.RES += fmt.Sprintf("\n%d. %s %s ",len(data.SEEN), tex,data.VAL_BOOL)
            }
        }
        return c.Edit(fmt.Sprintf("Общее количество: %d %s\n",len(data.SEEN), data.RES), MenuRet())
    }else if buttonID == "val_cnt3"{
        for i, tex := range data.ALL_TEXT{
            if(i <= 164){
                data.RES += fmt.Sprintf("\n%d. %s ", i+1, tex)
            }else{
                data.IND = i
                break
            }
        }
        SetUserState(data, "val_more_sk")
        return c.Edit(fmt.Sprintf("Общее количество: %d %s\n",len(data.ALL_TEXT), data.RES), MenuRet())
    }else{
        for i, tex := range data.ALL_TEXT {
            if(tex == "Null UI Data!"){
                tex = "Тестовый"
            }
            data.RES += fmt.Sprintf("\n%d. %s ", i+1, tex)
        }
        return c.Edit(fmt.Sprintf("Общее количество: %d %s\n",len(data.ALL_TEXT), data.RES), MenuRet())
    }
    }
    return nil
    
}

func ValMoreSk(c tele.Context)error{
    
    return c.Send("Hello")
}

func LolActions(c tele.Context)error{
    data := GetUser(c.Chat().ID) 
    menu, _ := MenuLol()
    SetUserState(data, "lol_action")
    return c.Edit("Выберите действие", menu)
}

func LolAction(c tele.Context)error{
    _, btns := MenuLol()
    chatId := GetUser(c.Chat().ID)
    data := GetUser(chatId)
    serv_response, err := api.LolStatus()
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    score_response, err := api.LolMasteryScore()
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    user_response, err := api.LolSummoner()
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    for _, btn := range btns {
        switch btns.Unique {
            case user_btn:
                SetUserState(data, "lol_user")
                return c.Edit(fmt.Sprintf("Ваш уровень: %d\nПоследняя активность в %d", user_response.LVL, time.Unix(user_response.TIME, 0).Format("00:00")), LolUserMenu())
            case server_btn:
                return c.Edit(fmt.Sprintf("Инцедент в %s, серьезности  %s\n Описание: %s\n Обновленно в %s сейчас %s\n Описание: %s",serv_response.INC[0].CR, serv_response.INC[0].SEVR, serv_response.INC[0].TIT[0].CNT, serv_response.INC[0].UPD[0].UP, serv_response.INC[0].STAT, serv_response.INC[0].UPD[0].CNT),MenuRet())
            case champ_btn:
                SetUserState(data, "lol_chemp")
                return c.Edit(fmt.Sprintf("Общее мастерство: %d",score_response),LolChaempMenu())
            case tours_btn:
                SetUserState(data, "lol_tours")
                return c.Edit("Информация о clash режиме",LolToursMenu())
        }
    }
    
    return nil
}
var buts []tele.InlineButton
func LolUser(c tele.Context)error{
    _, btns := LolUserMenu()
    leag_response, err := api.LolLeagPu()
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    for _, btn := range btns {
        switch btn.Unique{
            case league_btn:
                leags := map[int]tele.InlineButton
                buts = []tele.InlineButton{}
                menu := &tele.ReplyMarkup{}
                board := [][]tele.InlineKeyboard{}
                for i, resp := range leag_response{
                    btn := tele.InlineButton{
                        Text: fmt.Sprintf("Лига %d",i+1), Unique: fmt.Sprintf("%d", i)
                    }
                    leags[i] = btn
                }
                
                for _, v := range leags {
                    if len(buts) < 4{
                        buts = append(buts, v)
                        board = append(board,buts)
                        menu.InlineKeyboard = board
                    }else{
                        buts = []tele.InlineButton{}
                    }
                }
                
                for _, val := range leags {
                    buts = []tele.InlineButton{}
                    buts = append(buts, v)
                }
                
                c.Edit("Выберите Лигу",menu)
            case gameTime_btn:
            
                return c.Edit("")
            case last_matches:
            
                return c.Edit("")
        }
    }
    return nil
}

func LolLeag(c tele.Context)error{
    var id int
    leag_response, err := api.LolLeagPu()[id]
    
    if err != nil {
        log.Println("Ошибка получения статуса:", err)
        return err
    }
    
    id = strconv.Atoi(strings.TrimSpace(c.CallBack().Data))
    
    return c.Edit(
        fmt.Sprintf(
           "\nВаш ранк: %s %s\nПоинты лиги: %d\nЧисло побед: %d\nЧисло поражений: %d\nИгровой  режим: %s",
           leag_response.RANK,leag_response.TIER,
           leag_response.LEAPT,
           leag_response.WNS,leag_response.LSES,
           switch leag_response.Mode {
               case "RANKED_SOLO_5x5":
                    "Ранговый соло"
               case "RANKED_FLEX_SR":
                    "Ранговый отряд"
               default:
                    "Draft"
           }
        )
        ,MenuRet()
    )
}

func LolChemp(c tele.Context)error{
    _, btns := LolChempMenu()
    for _, btn := range btns {
        switch btn.Unique{
            case ratation_btn:
                return c.Edit("")
            case allChemps_btn:
                return c.Edit("")
            case topChemps_btn:
                return c.Edit("")
        }
    }
    return nil
}
func LolTours(c tele.Context)error{
    _, btns := LolToursMenu()
    for _, btn := range btns {
        switch btn.Unique{
            case clash_btn:
                return c.Edit("")
            case tours_btn:
                return c.Edit("")
            case clashById_btn:
                return c.Edit("")
        }
    }
    return nil
}