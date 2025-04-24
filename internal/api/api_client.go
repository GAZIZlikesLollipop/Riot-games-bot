package api

import(
    "encoding/json" // Для работы с JSON
    "io"            // Для работы с потоками данных
    "log"           // Для логирования ошибок
    "net/http"      // Для отправки HTTP-запросов
    "fmt"
    "os"
)

var key = os.Getenv("API_KEY")

func GetPuuid(region string, playerName string, playerTag string) string {
    baseUrl := "https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s"

    url := fmt.Sprintf(baseUrl, region, playerName, playerTag, key)
    resp, err := http.Get(url)
    if err != nil {  
        log.Println("Ошибка запроса:", err)
        return ""
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка полученмя данных:", err)
        return ""
    }
    // Распарсим только нужное поле
    var data Puuid
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга JSON:", err)
        return ""
    }
    if data.ERR.CODE == 404 {
        return "Вы ввели неверные данные! попробуйте снова"
    }else if data.ERR.CODE == 429 {
        return "Спамить нельзя!"
    }
    
    return data.ID
}

func GetValContent(shard string)(ValContent,error){
    baseUrl := "https://%s.api.riotgames.com/val/content/v1/contents?locale=ru-RU&api_key=%s"
    url := fmt.Sprintf(baseUrl,shard,key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса: ", err)
        return ValContent{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Println("API вернуло ошибку: ", resp.Status)
        return ValContent{}, err
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("ошибка чтения ответа:", err)
        return ValContent{}, err
    }
    var data ValContent
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга JSON:", err)
        return ValContent{}, err
    }
    
    return data, nil
}

func GetValStatus(shard string)(ValStatus, error){
    baseUrl := "https://%s.api.riotgames.com/val/status/v1/platform-data?api_key=%s"
    url := fmt.Sprintf(baseUrl ,shard, key)
    
    resp, err := http.Get(url)
    if err != nil {
        log.Println("Ошибка запроса", err)
        return ValStatus{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return ValStatus{}, err
    }
    
    var data ValStatus
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return ValStatus{}, err
    }
    return data, nil
}

func ValRanked(shard string)(ValLead,error){
    act, _ := GetValContent(shard)
    baseUrl := "https://%s.api.riotgames.com/val/ranked/v1/leaderboards/by-act/%s?size=100&startIndex=0&api_key=%s"
    url := fmt.Sprintf(baseUrl,shard,act.ACT[0].ID,key)
    
    resp, err := http.Get(url)
    if err != nil {
        log.Println("Ошибка запроса", err)
        return ValLead{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return ValLead{}, err
    }
    
    var data ValLead
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return ValLead{}, err
    }
    return data, nil
}

//League
func LolMatchs(puuid string, reg string) ([]string, error) {
    baseUrl := "https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=20&api_key=%s"
    url := fmt.Sprintf(baseUrl, reg, puuid, key)
    
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return []string{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return []string{}, err
    }
    var data []string
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return []string{}, err
    }
    
    return data, nil
}

func LolLeagPu(puuid string,region string) (LeagPu, error) {
    baseUrl := "https://%s.api.riotgames.com/lol/league/v4/entries/by-puuid/%s?api_key=%s"
    url := fmt.Sprintf(baseUrl,region, puuid, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return LeagPu{}, err 
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return LeagPu{}, err 
    }
    var data LeagPu
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return LeagPu{}, err 
    }
    return data, nil
    
}

func LolClash(puuid string,region string) (ClPu,error) {
    
    baseUrl := "https://%s.api.riotgames.com/lol/clash/v1/players/by-puuid/%s?api_key=%s"
    
    url := fmt.Sprintf(baseUrl,region, puuid, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return ClPu{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return ClPu{}, err
    }
    var data ClPu
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return ClPu{}, err
    }
    return data, nil
    
}

func LolTours(region string) (ClTr, error) {
    
    baseUrl := "https://%s.api.riotgames.com/lol/clash/v1/tournaments?api_key=%s"
    
    url := fmt.Sprintf(baseUrl,region,key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return ClTr{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return ClTr{}, err
    }
    var data ClTr
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return ClTr{}, err
    }
    return data, nil
    
}

func LolRotat(region string) (Rotat,error) {
    
    baseUrl := "https://%s.api.riotgames.com/lol/platform/v3/champion-rotations?api_key=%s"
    
    url := fmt.Sprintf(baseUrl,region,key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return Rotat{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return Rotat{}, err
    }
    var data Rotat
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return Rotat{}, err 
    }
    return data,nil
    
}

func LolStatus(region string) (ValStatus, error) {
    
    baseUrl := "https://%s.api.riotgames.com/lol/status/v4/platform-data?api_key=%s"
    
    url := fmt.Sprintf(baseUrl,region,key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return ValStatus{}, err 
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return ValStatus{}, err 
    }
    var data ValStatus
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return ValStatus{}, err 
    }
    return data, nil
    
}

func LolSummoner(puuid string,region string) (Summoner,error) {
    baseUrl := "https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s?api_key=%s"
    
    url := fmt.Sprintf(baseUrl,region, puuid, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return Summoner{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return Summoner{}, err
    }
    var data Summoner
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return Summoner{}, err
    }
    return data, nil
}

func LolMasteryScore(puuid string,region string) (int,error) {
    
    baseUrl := "https://%s.api.riotgames.com/lol/champion-mastery/v4/scores/by-puuid/%s?api_key=%s"
    
    url := fmt.Sprintf(baseUrl,region, puuid, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return 0, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return 0, err
    }
    var data int
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return 0, err
    }
    return data, nil
    
}

func LolChemps(puuid string,region string)(Chemps, error){
    
    baseUrl := "https://%s.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-puuid/%s?api_key=%s"
    
    url := fmt.Sprintf(baseUrl,region, puuid, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return Chemps{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return Chemps{}, err
    }
    var data Chemps
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return Chemps{}, err
    }
    return data, nil
    
}

func LolClashTeam(teamId string,region string)(ClTeam, error){
    baseUrl := "https://%s.api.riotgames.com/lol/clash/v1/teams/%s?api_key=%s"
    
    url := fmt.Sprintf(baseUrl,region, teamId, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return ClTeam{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return ClTeam{}, err
    }
    var data ClTeam
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return ClTeam{}, err
    }
    return data, nil
}

func LolMatchPu(matchId string, reg string)(LolMatch, error){
    
    baseUrl := "https://%s.api.riotgames.com/lol/match/v5/matches/%s?api_key=%s"
    url := fmt.Sprintf(baseUrl,reg, matchId, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return LolMatch{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return LolMatch{}, err
    }
    var data LolMatch
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return LolMatch{}, err
    }
    return data, nil
    
}

func LolDdragon()(map[string]DDragonDt, error){
    
    urlVers := "https://ddragon.leagueoflegends.com/api/versions.json"
    resp_vers, err := http.Get(urlVers)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return map[string]DDragonDt{}, err
    }
    defer resp_vers.Body.Close()
    
    body_vers, err := io.ReadAll(resp_vers.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return map[string]DDragonDt{}, err
    }
    var vers []string
    err = json.Unmarshal(body_vers, &vers)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return map[string]DDragonDt{}, err
    }
    
    baseUrl := "https://ddragon.leagueoflegends.com/cdn/%s/data/ru_RU/champion.json"
    url := fmt.Sprintf(baseUrl, vers[0])
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return map[string]DDragonDt{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return map[string]DDragonDt{}, err
    }
    var resp_map map[string]DDragonDt
    err = json.Unmarshal(body, &resp_map)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return map[string]DDragonDt{}, err
    }
    
    return resp_map, nil
}

func LolChempIcon(name string)(string, error){
    urlVers := "https://ddragon.leagueoflegends.com/api/versions.json"
    resp_vers, err := http.Get(urlVers)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return "", err
    }
    defer resp_vers.Body.Close()
    
    body_vers, err := io.ReadAll(resp_vers.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return "", err
    }
    var vers []string
    err = json.Unmarshal(body_vers, &vers)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return "", err
    }
    
    baseUrl := "https://ddragon.leagueoflegends.com/cdn/%s/img/champion/%s.png"
    url := fmt.Sprintf(baseUrl,vers[0],name)
    
    return url, nil
}