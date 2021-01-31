# Křeček

## Algoritmus

Úlohu řeším takto. Procházím graf cestiček do hloubky, s tím, že se zastavuji v případě že narazím na jídlo, na místo 
kde už jsem byl a navigovala stejná hemisféra křečkova mozku, s jakou křeček naviguje teď, nebo dorazím do slepé uličky. 
Potom vrátím proč jsem skončil a v případě jídla, jak dlouho to trvalo se k němu dostat. Pokud mám na některé křižovatce 
návratové hodnoty ze všech chodbiček rozhodnu se, které je nejlepší a tu vrátím, respektive to co vrátila funkce 
z navazující křižovatky. V případě levé hemisféry je nejlepší maraton, pak slepá cesta a poté nejdelší cesta k jídlu. 
V případě pravé je to obráceně a navíc chci tu cestu k jídlu nejkratší. A na konci rozhodnu podle toho co vrátí 
startovní křižovatka.

Ještě bych popsal jak rozhoduji jednotlivé případy. Pokud narazím na jídlo, je to jasné. Pokud narazím na místo, kde 
jsem byl se stejnou navigující hemisférou, je to maraton, jelikož jsem se dostal do kruhu, v rámci jedné prohledávací 
větve. Avšak po opuštění křižovatky v rámci procházení do hloubky, mažu informaci, že jsem na ní byl, protože bych se na 
ní mohl dostat z jiné, avšak kratší cesty, proto si ponechávám pouze za jak dlouho od startu jsem tu byl a při dalším 
procházení na ní vstoupím pouze tehdy, když bych si časově polepšil a mohl se tak k jídlu dostat za kratší čas. No 
a pokud z křižovatky nevedou žádné cesty je slepá.

Myslím že algoritmus najde řešení, poněvadž z každé křižovatky projde všechny možné cesty a rozhodne co je pro danou 
hemisféru nejlepší.

## Složitost

Algoritmus projde celé bludiště v čase řádově N^2^, poněvadž z každé křižovatky může navštívit až N dalších a křižovatek 
projde taky N, takže podle mne je asymptotická složitost $O(N^2)$
