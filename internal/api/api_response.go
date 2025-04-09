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
    SKIN []ValItem `json:"skins"`
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

// League

type Summoner struct {
    TIME int64 `json:"revisionDate"`
    LVL int64 `json:"summonerLevel"`
}

type Rotat struct {
    LVL int
    CHEMPNEW []int
    CHEMPS []int
}

type LeagPu []LeagEnt 

type LeagEnt struct {
    TIER string `json:"tier"`
    RANK string `json:"rank"`
    LEAPT int `json:"leaguePoints"`
    WNS int `json:"win"`
    LSES int `json:"loses"`
    Mode string `json:"queueType"`
    Act bool `json:"inactive"`
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
    Id int `json:"id"`
    Shed []ClShed `json:"schedule"`
}

type ClShed struct {
    STRT int64 `json:"startTime"`
    REG int64 `json:"registrationTime"`
    CNC bool `json:"cancelled"`
}

type TopMastery []TopMasteryDTO

type TopMasteryDTO struct {
    ChmpId int64 `json:"championId"`
    ChmpPt int `json:"championPoints"`
    ChempLvl int `json:"championLevel"`
}

type Chemps []ChempsDTO

type ChempsDTO struct {
    ChmpId int64 `json:"championId"`
    ChmpPt int `json:"championPoints"`
    ChempLvl int `json:"championLevel"`
    ChempPlTime int64 `json:"lastPlayTime"`
    ChmpPtUL int64 `json:"championPointsUntilNextLevel"`
}

type ClTeam TeamDTO

type TeamDTO {
    ID string `json:"id"`
    TourId int `json:"tournamentId"`
    Name string `json:"name"`
    Tier int `json:"tier"`
    Capitan string `json:"captain"`
    Abbrev string `json:"abbreviation"`
    Plyrs []ClEnt `json:"players"`
}

type LolMatch struct {
    MetDT 
    Info
}

type MetaData struct {
    DtVers string `json:"dataVersion"`
    MatchId string `json:"matchId"`
    //Prpnts []string `json:"participants"`
}

type InfoDTO struct {
    CrTime int64 `json:"gameCreation"`
    DurTime int64 `json:"gameDuration"`
    EndTime int64 `json:"gameEndTimestamp"`
    StrtTime int64 `json:"gameStartTimestamp"`
    GmMode string `json:"gameMode"`
    Team []InfoTeams `json:"teams"`
}

type InfoTeams struct {
    IsWin bool `json:"win"`
    teamId int `json:"teamId"`
}