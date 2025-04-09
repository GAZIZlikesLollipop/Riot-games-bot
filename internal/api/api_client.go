package api

import(
    "encoding/json" // Для работы с JSON
    "io"            // Для работы с потоками данных
    "log"           // Для логирования ошибок
    "net/http"      // Для отправки HTTP-запросов
    "fmt"
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

func GetValContent(region string)(ValContent,error){
    baseUrl := "https://%s.api.riotgames.com/val/content/v1/contents?locale=ru-RU&api_key=%s"
    url := fmt.Sprintf(baseUrl, region, key)
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

func GetValStatus(region string)(ValStatus, error){
    baseUrl := "https://%s.api.riotgames.com/val/status/v1/platform-data?api_key=%s"
    url := fmt.Sprintf(baseUrl, region, key)
    
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
        log.Println("Оштбка парсинга данных", err)
        return ValStatus{}, err
    }
    return data, nil
}

//League
func LolMatchs(puuid string) ([]string, error) {
    baseUrl := "https://asia.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=20&api_key=%s"
    url := fmt.Printf(baseUrl, puuid, key)
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

func LolLeagPu(puuid string) (LeagPu, error) {
    baseUrl := "https://ru.api.riotgames.com/lol/league/v4/entries/by-puuid/%s?api_key=%s"
    url := fmt.Printf(baseUrl, puuid, key)
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
    var data LeagPu.LeagEnt
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return LeagPu{}, err 
    }
    return data, nil
    
}

func LolClash(puuid string) (ClPu,error) {
    
    baseUrl := "https:ru.api.riotgames.com/lol/clash/v1/players/by-puuid/%s?api_key=%s"
    
    url := fmt.Printf(baseUrl, puuid, key)
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
    var data ClPu.ClEnt
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return ClPu{}, err
    }
    return data, nil
    
}

func LolTours(puuid string) (ClTr, error) {
    
    baseUrl := "https:ru.api.riotgames.com/lol/clash/v1/players/by-puuid/%s?api_key=%s"
    
    url := fmt.Printf(baseUrl, puuid, key)
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
    var data ClTr.ClTrs
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return ClTr{}, err
    }
    return data, nil
    
}

func LolRotat() (Rotat,error) {
    
    baseUrl := "https://ru.api.riotgames.com/lol/platform/v3/champion-rotations?api_key=%s"
    
    url := fmt.Printf(baseUrl,key)
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

func LolStatus() (ValStatus, error) {
    
    baseUrl := "https://ru.api.riotgames.com/lol/status/v4/platform-data?api_key=%s"
    
    url := fmt.Printf(baseUrl,key)
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

func LolSummoner() (Summoner,error) {
    baseUrl := "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s?api_key=%s"
    
    url := fmt.Printf(baseUrl, puuid, key)
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

func LolMasteryScore(puuid string) (int,error) {
    
    baseUrl := "https://ru.api.riotgames.com/lol/champion-mastery/v4/scores/by-puuid/%s?api_key=%s"
    
    url := fmt.Printf(baseUrl, puuid, key)
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

func LolMasteryTop()(TopMastery, error){
    
    baseUrl := "https://ru.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/top?count=10&api_key=%s"
    
    url := fmt.Printf(baseUrl, puuid, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return TopMastery{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return TopMastery{}, err
    }
    var data TopMastery
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return TopMastery{}, err
    }
    return data, nil
    
}

func LolChemps(puuid string)(Chemps, error){
    
    baseUrl := "https://na1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-puuid/%s?api_key=%s"
    
    url := fmt.Printf(baseUrl, puuid, key)
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

func LolClashTeam(teamId string)(ClTeam, error){
    baseUrl := "https://ru.api.riotgames.com/lol/clash/v1/teams/%s?api_key=%s"
    
    url := fmt.Printf(baseUrl, teamId, key)
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

func LolMatchPu(matchId string)(MetaData, error){
    baseUrl := "https://asia.api.riotgames.com/lol/match/v5/matches/%s?api_key=%s"
    
    url := fmt.Printf(baseUrl, puuid, key)
    resp, err := http.Get(url)
    if err != nil{
        log.Println("Ошибка запроса ", err)
        return MetaData{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Ошибка чтения тела ответа:", err)
        return MetaData{}, err
    }
    var data MetaData
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println("Ошибка парсинга данных", err)
        return MetaData{}, err
    }
    return data, nil
}