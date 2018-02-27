# Oblig 1
## Oppgave 1

|Binære Tall    |Hexadesimale Tall  |Desimal Tall|
|:-------------:|:-----------------:|:----------:|
| 1101 | 0xD | 13 |
| 110111101010 | 0xDEA | 3562
| 1010111100110100 | 0xAF34 | 44852
| 1111111111111111 | 0xFFFF | 65535
| 10001011110001010 | 0x1178A | 71562

### Oppgave A
#### Binary -> hex
Sett inn 0 foran til antall siffer er delelelig med 4. Konverter 4 tall av gangen til ett tall i hexadecimal.
#### Hex -> binary
Konverter hvert siffer til 4 siffer i binary. Sett sammen.
#### Binary -> decimal
Begynn bakerst, gang hvert siffer med 2^n hvor n er posisjonen til tallet. legg sammen alle verdiene.
#### Decimal -> binary
Del tallet på to. Skriv ned 1 hvis det er igjen en rest. Skriv 0 hvis det ikke er noen rest. Fortsett til det oprinnelige tallet er 0. Det binære tallet er nå skrevet ned, men i feil rekkefølge. Snu rekkefølgen.
### Oppgave B
#### Hex -> Decimal
Begynn bakerst, gang hvert siffer med 16^n hvor n er posisjonen til tallet. legg sammen alle verdiene.
#### Decimal -> hex
Samme som for decimal -> binary, men deler på 16 og skriver resten som er igjen i hexadecimal.

---
## Oppgave 2
### Oppgave C 
![alt text](https://github.com/vegarr14/Gruppe404/blob/Oblig-1/Oppgave%202/Diagram.PNG "test")
Bubble sort funksjonene bruker mange flere operasjoner enn quick sort. Bubble sort funksjonene har mye høyere økning på antall operasjoner enn antall elementer (O(n^2)). Quick sort har ganske lik økning på antall operasjoner og antall elementer (O(n))

---
## Oppgave 3
Loop.go bruker 1MB minne og under 0.1% av CPUen på pcen vi testet på

---
## Oppgave 4
### Oppgave A
Vi får skrevet ut symboler fra iso 8859-1 mellom 80 og ff. Mellom 80 og 9f får vi bare ukjente tegn fordi de ikke eksisterer i iso 8859-1. Ingenting er mappet til de verdiene. Vi får det samme resultatet på alle gruppemedlemmers pcer
### Oppgave B
Vi testet koden i terminal og powershell. Fikk samme resultat i begge.
