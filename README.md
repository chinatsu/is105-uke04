# ICA 01

**Gruppe 4: Crippling Depression**

Brage Sydskogen, Martin Hovet, Tønnes Røren, Marius Kaurin, Kent Daleng, Stian Simonsen

## Innhold
* [Oppgave 1.2.1](#oppgave-121)
* [Oppgave 1.2.2](#oppgave-122)
* [Oppgave 1.2.3](#oppgave-123)
* [Oppgave 1.2.4](#oppgave-124)

---

## Oppgave 1.2.1

* 1 = 0b1 (1)
* 2 = 0b10 (2)
* 5 = 0b101 (4+1)
* 8 = 0b1000 (8)
* 16 = 0b10000 (16)
* 256 = 0b100000000 (256)

En enkel måte å regne ut desimalverdier fra binærtall er å gå fra mest signifikant bit og regne i hodet, en posisjon etter den andre. Man har et “register” i hodet som man legger til 1 i for å representere første, mest signifikante bit. Deretter ganger man med to for å representere et hopp til neste bit. Dersom denne bit-en er 1, legger man også til en. Sånn går man til man er ferdig. Her er litt Python-kode for å forklare logikken

```py
binary = “100” # binærtallet er en str for å kunne gå over hver
               # bokstav i for-loop
register = 0
for idx, n in enumerate(binary):
    if idx != 0:
        register = register * 2 # doble registerverdien
                                # hver gang etter
                                # første posisjon
    if n == “1”:
        register += 1
```


* 100 = ((1\*2)\*2) = 4
* 1001 = ((((1\*2)\*2)\*2)+1) = 9
* 1100110011 = 819


## Oppgave 1.2.2

Et tresifret binært tall er et tall mellom 000 og 111, altså 0 til 7 i desimaltall.

1. Lise får vite at tallet er et oddetall, det vil si at 1, 3, 5 og 7 er mulige valg.
  * Mulige valg (M) er 4 av totalt (N) 8. Dette kan regnes ut ved å bruke følgende formel: log2(1/(M/N))
    * log2(1/(4/8)) = 1 bit
   * Man kan også tenke seg at et oddetall defineres som et tall som ikke kan deles med 2. Dermed er en gitt integer enten delelig med 2, eller ikke. Altså er det kun to mulige “states” et tall kan være med tanke på dette, True eller False.
2. Per får vite at tallet ikke er et multiplum av 3. Det vil si alle tall unntatt 0, 3 og 6 er mulige valg.
  * Mulige valg er 1, 2, 4, 5 og 7, altså er det 5 mulige av totalt 8.
    * log2(1/(5/8)) = 0.6780719051126376 bits
  * Praktisk sett kan man også tenke seg at en gitt integer er et multiplum av 3 eller ikke, altså kan dette også representeres med 1 bit.
3. Oskar har fått vite at tallet inneholder nøyaktig to enere.
  * Av alle binærverdier mellom 000 og 111, er det tre som har akkurat to bits med verdien 1: 011, 101 og 110. Dermed er det tre mulige verdier av totalt 8.
    * log2(1/(3/8)) = 1.4150374992788437 bits
  * I praksis kan man tenke seg at man trenger å vite tallet på enere, og siden det tallet kan være mellom 0 og 3, kan det representeres med 2 bits.
4. Louise har fått vite alt som de over vet.
  * Louise får vite at 011, 101 og 110 er alle mulige verdier.
  * Av disse er 011 (3) et multiplum av 3, og 110 (6) er et multiplum av 2 og 3.
  * Kun 101 gjenstår som mulig valg av de totalt 8 verdiene.
    * log2(1/(1/8)) = 3 bits





## Oppgave 1.2.3

Alle i gruppen bruker enten bash i MacOS eller Linux, eller zsh på Windows med [Babun](http://babun.github.io/).
Altså bruker alle samme sett med kommandoer.

* For å klone repoet til lokalmaskinen
```
~ $ git clone https://github.com/chinatsu/is105-uke04.git
```

* For å skifte mappe til repo-mappen som nettopp har blitt opprettet
```
~ $ cd is105-uke04/
```

* For å vise filene i mappen
```
is105-uke04 $ ls
hello.go  LICENSE  README.md
```

* For å åpne `hello.go` i en editor (`atom` kan byttes ut med `nano`, `vim`, `notepad`, etc.)
```
is105-uke04 $ atom hello.go
```

* Etter endringer, kan vi sjekke hvilke filer som har blitt endret
```
is105-uke04 $ git status
```

* Vi ser at `hello.go` ikke er satt til “staging”, så for å gjøre dette kjører vi følgende:
```
is105-uke04 $ git add hello.go
```

* Alternativt kan vi legge til alle unstaged filer i nåværende mappe (og rekursivt inn i submapper)
```
is105-uke04 $ git add .
```

* For å “lagre” endringene til git-systemet lokalt (-m betyr at det følgende er en commit-melding)
```
is105-uke04 $ git commit -m “La til hallo på koreansk”
```

* Endringene er ikke enda synlige på GitHub (origin), så for å gjøre dette kjører vi
```
is105-uke04 $ git push
```
Deretter blir man spurt etter brukernavn og passord på GitHub. Dette laster opp alle lagrede commits.

* For å oppdatere lokalkopien av repoet, om nye endringer har forekommet
```
is105-uke04 $ git pull
```


## Oppgave 1.2.4

Én fordel er at alle kan se endringene til hverandre, og alle kan samarbeide på det samme prosjektet. Litt som Google Docs.

En ulempe er at det er ganske komplisert for nye til systemet. Endringer må “legges til”, “committed” og “pushet” til hovedrepoet, i riktig rekkefølge.

Genererte objektfiler (kompilerte programmer) på Windows har filnavn med .exe som “filetternavn”. I OS X og Unix har objektfilene ikke etternavn. Hovedgrunnen til at alle er forskjellige er på grunn av operativsystemets forskjeller, systemkall i Windows er ulike systemkall i OS X og Unix-systemer.

Ved første øyekast er Java og Go ganske ulike.
Java er hovedsakelig objektorientert med en hovedklasse, og klasser til andre objekter. Klassene har sine egne metoder/funksjoner, men det går ikke an å ha funksjoner utenfor klassene.
På den andre siden har Go en hovedfunksjon som alltid blir kalt ved programstart. På en måte kan trekke paralleller mellom pakkene i Go og klassene i Java, men pakkene i Go har hovedsakelig funksjoner og ingen constructor.

Go gjør det svært enkelt å programmere med tanke på hvor raskt en kan utvikle programmer ved bruk av go run og se sine resultater. Det er like enkelt å gjøre dette i Windows som på OS X eller Linux. Andre programmeringsspråk har ofte mye mer kompliserte økosystemer, mens Go hovedsakelig kun har et program for det aller meste av operasjoner.

Det er hensiktsmessig å legge inn logcli i et Git repository, spesielt med tanke på videre utvikling. Spesifikt i denne oppgaven, kan logcli brukes som basis for videre utvikling til logbcli. Åpne Git repos på GitHub er offentlig synlige, og hvem som helst kan komme og se koden. De kan “forke” og gjøre sine egne endringer, og om de synes de  har gjort viktige endringer, kan de åpne en “pull request” og bidra til original-repoet.

Pakken “log” som vi har laget, blir kalt i importsetningen som “./log”, for å fortelle Go at pakken ligger i mappen relativt til main-pakken, og ikke i standardbiblioteket til Go. Hadde importsetningen sagt “log”, ville denne pakken blitt importert i stedet.
