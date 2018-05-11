# Oblig 4  
## Gruppe 404  
Vegar Ryen  
Erlend Thorsen  
Sigbjørn Beitnes Moe  
Josef Skar  
Thomas Lochner  
### Info

Det er nødvendig og sette GOPATH til oblig 4 mappa. Uten dette vil ikke importene fungere.

For å bruke applikasjonen må man skrive inn et bynavn eller et bynavn etterfulgt av et komma og landskode  
eks: oslo  
eksempel med landskode: oslo, no


### Systemspesifikasjon

Vi har laget en værvarslings applikasjon som viser temperatur, vindstyrke og nedbørsmengde i millimeter 5 dager fram i tid, for hver tredje time. For været akkurat nå viser applikasjonen et ikon som samsvarer med været, vindstyrke, temperatur og skriver ut en melding som passer til værsituasjonen. Den skriver også ut hvor lenge det er til solen går ned eller opp eller hvor lenge det var siden solen gikk ned. Applikasjonens nytteverdi ligger i at den er oversiktlig, viser oppdatert informasjon og viser værinformasjon over de fleste steder i verden.

### Systemarkitektur

Applikasjonen er bygd opp av tre .go filer og tre .html filer samt en .css fil som inneholder stylingen til alle html filene.  

**Application.go**

Application.go filen er hovedfilen i applikasjonen. Den inneholder http serveren i main funksjonen. Oppgaven til denne filen er å starte serveren, Lytte til request på port 8080 og sende riktig html fil til klient basert på data som klient har gitt.

Importerer:  
+ "html/template"
+ "log"
+ "net/http"
+ "strings"
+ "weather"
+ "fmt"

Funksjoner:  
+ start(http.ResponseWriter, *http.Request): Sender startsiden wheater.html til klient om klient går til localhost:8080/
+ getSite(http.ResponseWriter, *http.Request): Sender result.html med data basert på hva klient søkte på i startsiden.
+ css(http.ResponseWriter, *http.Request): Sender style.css fil til klient.
+ getWheaterIcon(http.ResponseWriter, *http.Request): Sender riktig værikon til klient.
+ get404(http.ResponseWriter, *http.Request): Sender 404.html til klient om søket til klient ikke returnerer riktig data.
+ main(): Starter server og inneholder HandleFuncs som starter funksjonene over.*

**Weather.go**

Weather.go filen har ansvar for å fylle html structen som brukes i html templates i application.go basert på søkeordet som sendes fra application.go

Importerer:  
+ "fmt"
+ "net/http"
+ "io/ioutil"
+ "encoding/json"
+ "weather/responsecreator"
+ "time"
+ "strconv"

Funksjoner:  
+ CheckError(error): Skriver ut en eventuelt feilmelding til terminal
+ Weather(string):Henter JSON data og fyller det i html structen. Stopper om man får feil.
+ Date(): Konverterer forcast verdiene fra unix time til  Day.Month.Hour
+ getData(string): Henter innholdet fra nettside.

**Responsecreator.go**

responsecreator.go filen har ansvar for å produsere riktige meldinger og vær ikon i Getresponse funksjonen, samt og produsere riktig solnedgang/soloppgang melding i funksjonen Time. Her må man også konvertere fra Unix time.

Importerer:
+ "fmt"
+ "time"
+ "strconv"

Funksjoner:
+ Getresponse(int):Sjekker værkode og returner en melding og hvilket værikon som skal vises
+ Time(int64, int64, int64): Finner ut om solen går ned eller opp samt skriver ut en melding med tid konvertert fra unixtime.
+ CheckError(error): Skriver ut feilmelding til terminal  
