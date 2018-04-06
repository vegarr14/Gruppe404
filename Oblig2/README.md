# Gruppe404
Vegar Ryen

Erlend  Thorsen

Sigbjørn Beitnes Moe

Josef Skar

Thomas Lochner
## Oppgave 3b
Sumfromfile.go er importert som en pakke i addtofile.go og kjøres fra addtofile.go
## Oppgave 3c
#### Oppgave a og b
Programmet sjekker om verdiene er tall. strconv.Atoi gir en error hvis noe annet enn tall er input.

Hvis det er ikke er nøyaktig to argumenter, skriver programmet en feilmelding og avslutter.

Programmet sjekker også om summen beholder samme fortegn som argumentene. eks. hvis begge argumentene er positive, skal summen også være positiv.

Hvis argumentene er for høye tall gir strconv.Atoi "value out of range".
#### Oppgave b
ioutil.ReadFile og ioutil.WriteFile har begge err, men de har ikke produsert noen feil under testing. ReadFile kjører alltid etter WriteFile så filen eksisterer alltid. Mulig at WriteFile produserer en feil hvis bruker ikke har nødvendig tilgang eller ikke nok diskplass til å lage filen.
