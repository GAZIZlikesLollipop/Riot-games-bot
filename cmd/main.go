package main

import (
    "log"           
    "os"
    "time" 
    tele "gopkg.in/telebot.v4"  
    "bot/internal/utils"
    "fmt"
)

type StepHandler func(c tele.Context) error

var handlers = map[string]StepHandler{
    "del": utils.DelMess,
    "reg_choo": utils.RegChoo,
    "game_act": utils.GameChoo,
    "val_actions": utils.ValActions,
    "val_action": utils.ValAction,
    "val_cnt": utils.ValCnt,
    "val_more_sk": utils.ValMoreSk,
    
    "lol_actions": utils.LolActions
    "lol_action": utils.LolAction
    "lol_user": utils.LolUser
    "lol_chemp": utils.LolChemp
    "lol_tours": utils.LolTours
}

func main() {
    
    pref := tele.Settings{
        Token:  os.Getenv("TOKEN"),
        Poller: &tele.LongPoller{Timeout: 100 * time.Millisecond},
    }

    b, err := tele.NewBot(pref)
    if err != nil {
        log.Fatal(err)
        return
    }

    b.Handle("/start", func(c tele.Context) error {
        return c.Send("что бы узнать о большем функционале введи /help")
    })
    
    b.Handle("/help", func(c tele.Context) error {
        return c.Send("боту необходимы данные вашего аккаунта для того что бы ввести их введите команду /setprofile для того что бы узанть их /getprofile")
    })
    menu_reg, _ := utils.MenuRegion1()
    b.Handle("/setprofile", func(c tele.Context) error {
        data := utils.GetUser(c.Chat().ID)
        if data.REG == "" {
            utils.SetUserState(data, "reg_choo")
            utils.RegChoo(c)
            return c.Send("Выберите регион", menu_reg)
        } else {
            return c.Send("Данные вашего аккаутна уже имеюстя!")
        }
    })
    
    b.Handle("/launch", func(c tele.Context) error {
        menu, btns := utils.MenuGame()
        data := utils.GetUser(c.Chat().ID)
        utils.SetUserState(data, "game_act")
        for _, btns := range btns {
            if btn.Unique == "btn_val"{
            utils.SetUserState(data, "val_actions")
            }else{
            utils.SetUserState(data, "val_actions")
            }
        }
        return c.Send("Выберите игру", menu)
    })
    b.Handle("/getprofile", func(c tele.Context) error {
        userData := utils.GetUser(c.Chat().ID)
        if userData.PUUID == "" {
            return c.Send("Вы не вводили данные, введите их /setprofile")
        } else {
            return c.Send(fmt.Sprintf("Ваш регио: %s\nВаш никнейм: %s\nВаш тег: %s\nВаш puuid: %s", userData.REG, userData.NAME, userData.TAG, userData.PUUID))
        }
    })
    
    b.Handle(&utils.BackBtn, func(c tele.Context)error{
        userData := utils.GetUser(c.Chat().ID)
        utils.GoBack(userData)

        if handler, exists := handlers[userData.STATE]; exists {
                return handler(c) 
            }

        return nil
    
    })
    
    b.Handle(tele.OnCallback, func(c tele.Context) error {
        userData := utils.GetUser(c.Chat().ID)

        if handler, exists := handlers[userData.STATE];     exists {
            return handler(c)
        }

        return nil
    })

    b.Handle(tele.OnText, func(c tele.Context) error {
    userData := utils.GetUser(c.Chat().ID)
    if userData.STATE == "name_choo" || userData.STATE == "tag_choo" {
        return utils.PuuidText(c)
    }
    return nil
})
    
    b.Start()
    
}