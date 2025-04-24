package main

import (
    "log"           
    "os"
    "time" 
    tele "gopkg.in/telebot.v4"  
    "bot/internal/utils"
    "fmt"
    "strings"
)

type StepHandler func(c tele.Context) error

var handlers = map[string]StepHandler{
    "del": utils.DelMess,
    "reg_choo": utils.RegChoo,
    "region_choo": utils.RegChoo2,
    "shard_choo": utils.ShardChoo,
    
    "game_act": utils.GameChoo,
    "game_choo": utils.GameAct,
    
    "val_actions": utils.ValAction,
    "val_cnt": utils.ValCnt,
    "val_player": utils.ValPlr,
    
    "lol_actions": utils.LolAction,
    "lol_user": utils.LolUser,
    "lol_chemp": utils.LolChemp,
    "lol_tours": utils.LolTours,
    
    "lol_leag": utils.LolLeag,
    "lol_match": utils.LolLastMatch,
    "lol_chempion": utils.LolAllChemps,
    "lol_mychemps": utils.LolMyChemps,
    "lol_tour": utils.LolTour,
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
        if strings.TrimSpace(data.PUUID) == "" {
            utils.SetUserState(data, "reg_choo")
            return c.Send("Выберите регион", menu_reg)
        } else {
            return c.Send("Данные вашего аккаутна уже имеюстя!")
        }
    })
        b.Handle("/getprofile", func(c tele.Context) error {
        userData := utils.GetUser(c.Chat().ID)
        if userData.PUUID == "" {
            return c.Send("Вы не вводили данные, введите их /setprofile")
        } else {
            return c.Send(fmt.Sprintf("Ваш регио: %s\nВаша страна: %s\nВаш шард: %s\nВаш никнейм: %s\nВаш тег: %s\nВаш puuid: %s",userData.REG,userData.Region,userData.Shard, userData.NAME, userData.TAG, userData.PUUID))
        }
    })
    
    b.Handle("/launch", func(c tele.Context) error {
        menu, _ := utils.MenuGame()
        data := utils.GetUser(c.Chat().ID)
        utils.SetUserState(data, "game_choo")
        return c.Send("Выберите игру", menu)
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
        data := utils.GetUser(c.Chat().ID)

        if handler, exists := handlers[data.STATE];     exists {
            log.Println("state: ", data.STATE)
            return handler(c)
        }

        return nil
    })
    
    

    b.Handle(tele.OnText, func(c tele.Context) error {
        data := utils.GetUser(c.Chat().ID)
        switch data.STATE {
            case "name_choo":
                return utils.NameChoo(c,c.Message().Text)
            case "tag_choo":
                return utils.TagChoo(c,c.Message().Text)
        }
        return nil
    })
    
    b.Start()
    
}