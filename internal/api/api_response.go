package api

import "time"

type Puuid struct {
    ID string `json:"puuid"`
    ERR ErrPuuid `json:"status"`
}

type ErrPuuid struct{
    CODE int `json:"status_code"`
}

type ValContent struct{
    VERS string `json:"version"`
    CHAR []ValItem `json:"characters"`
    MAP []ValItem `json:"maps"`
    EQUIP []ValItem `json:"equips"`
    GAME_MODE []ValItem `json:"gameModes"`
    ACT []ValAct `json:"acts"`
}

type ValItem struct{
    NAME string `json:"name"`
}

type ValAct struct{
    NAME string `json:"name"`
    ID string `json:"id"`
    ACTIVE bool `json:"isActive"`
}

type ValStatus struct{
    MT []ValStatCnt `json:"maintenances"`
    INC []ValStatCnt `json:"incidents"`
}

type ValStatCnt struct{
    TIT []ValCont `json:"titles"`
    SEVR string `json:"incident_severity"`
    STAT string `json:"maintenance_status"`
    UPD []ValUpd `json:"updates"`
    CR time.Time `json:"created_at"`
}

type ValUpd struct{
    CNT []ValCont `json:"translations"`
    UP time.Time `json:"updated_at"`
}

type ValCont struct{
    LOC string `json:"locale"`
    CNT string `json:"content"`
}

type ValLead struct {
    PlrsCount int64 `json:"totalPlayers"`
    Players []LeadPlayer `json:"players"`
}

type LeadPlayer struct {
    Puuid string `json:"puuid"`
    PlName string `json:"gameName"`
    PlTag string `json:"tagLine"`
    Wins int64 `json:"numberOfWins"`
    LeadRank int64 `json:"leaderboardRank"`
    RankedRank int64 `json:"rankedRating"`
}

// League

type Summoner struct {
    TIME int64 `json:"revisionDate"`
    LVL int64 `json:"summonerLevel"`
}

type Rotat struct {
    Lvl int
    ChempNew []int
    Chemps []int
}

type LeagPu []LeagEnt 

type LeagEnt struct {
    LegId string `json:"leagueId"`
    TIER string `json:"tier"`
    RANK string `json:"rank"`
    LEAPT int `json:"leaguePoints"`
    WNS int `json:"win"`
    LSES int `json:"loses"`
    Mode string `json:"queueType"`
    Mini LeagMini `json:"miniSeries"`
}

type LeagMini struct{
    Target int `json:"target"`
    Wnins int `json:"wins"`
    Losses int `json:"losses"`
    Progress string `json:"progress"`
}

type ClPu []ClEnt

type ClEnt struct{
    TMID string `json:"teamId"`
    PST string `json:"position"`
    RLE string `json:"role"`
}

type ClTr []ClTrs

type ClTrs struct {
    Name string `json:"nameKey"`
    Day string `json:"nameKeySecondary"`
    Shed []ClShed `json:"schedule"`
}

type ClShed struct {
    STRT int64 `json:"startTime"`
    REG int64 `json:"registrationTime"`
    CNC bool `json:"cancelled"`
}

type Chemps []ChempsDTO

type ChempsDTO struct {
    ChmpId int64 `json:"championId"`
    ChmpPt int `json:"championPoints"`
    ChempLvl int `json:"championLevel"`
    ChempPlTime int64 `json:"lastPlayTime"`
    ChmpPtUL int64 `json:"championPointsUntilNextLevel"`
}

type ClTeam struct {
    TourId int `json:"tournamentId"`
    Name string `json:"name"`
    Tier int `json:"tier"`
    Capitan string `json:"captain"`
    Abbrev string `json:"abbreviation"`
    Plyrs []ClEnt `json:"players"`
}

type LolMatch struct {
    Info MatchInfo `json:"info"`
}

type MatchInfo struct {
    CrTime int64 `json:"gameCreation"`
    EndTime int64 `json:"gameEndTimestamp"`
    StrtTime int64 `json:"gameStartTimestamp"`
    GmMode string `json:"gameMode"`
    GmType string `json:"gameType"`
    Teams []struct{
        IsWin bool `json:"win"`
        TeamId int `json:"teamId"`
    } `json:"teams"`
    // Players []struct{
    //     PlName string `json:"summonerName"`
    //     Kills int `json:"kills"`
    //     Deaths int `json:"deaths"`
    //     Assists int `json:"assists"`
    //     TeamId int `json:"teamId"`
    //     ChampName string `json:"championName"`
    //     GoldEarned int `json:"goldEarned"`
    //     TotlDmgDeal int `json:"totalDamageDealt"`
    //     TotlDmgChmp int `json:"totalDamageDealtToChampions"`
    //     MinsKilled int `json:"totalMinionsKilled"`
    //     TurrentKills int `json:"turretKills"`
    // } `json:"participants"`
}

type DDragonResp struct {
    Data map[string]DDragonDt `json:"data"`
}

type DDragonDt struct {
    Key string `json:"key"`
    Name string `json:"name"`
    Title string `json:"title"`
    Lore string `json:"blurb"`
    Type string `json:"partype"`
    Tags []string `json:"tags"`
    Id string `json:"id"`
    Info struct {
        // из 10
        Atk int `json:"attack"`
        Def int `json:"defense"`
        Mag int `json:"magic"`
        Dif int `json:"difficulty"`
    } `json:"info"`
    Stats struct {
        Hp float64 `json:"hp"`
        HpLvl float64 `json:"hpperlevel"`
        Mp float64 `json:"mp"`
        MpLvl float64 `json:"mpperlevel"`
        Speed float64 `json:"movespeed"`
        Armor float64 `json:"armor"`
        ArmorLvl float64 `json:"armorperlevel"`
        SpellBlock float64 `json:"spellblock"`
        SpellBlockLvl float64 `json:"spellblockperlevel"`
        AtkRng float64 `json:"attackrange"`
        HpReg float64 `json:"hpregen"`
        HpRegLvl float64 `json:"hpregenperlevel"`
        MpReg float64 `json:"mpregen"`
        MpRegLvl float64 `json:"mpregenperlevel"`
        AtkDmg float64 `json:"attackdamage"`
        AtkDmgLvl float64 `json:"attackdamageperlevel"`
        AtkSpeed float64 `json:"attackspeed"`
        AtkSpeedLvl float64 `json:"attackspeedperlevel"`
        Crit float64 `json:"crit"`
        CritLvl float64 `json:"critperlevel"`
    } `json:"stats"`
}

// type LolUtils struct {
//     Matches []string `json:""`
// }