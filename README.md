# RozvrhBOT
Discord bot ktorý posiela linky na online hodiny. Tento bot vznikol vo voľnom čase a medzi/cez online hodinamy miesto robenia toho čo som mal. Bot má stále svoje chyby ale splní čo má. 

## Pridanie bota na server
- Choďte do [aplikácií](https://discord.com/developers/applications) v developer portály na Discorde.
- Vytvorte novú aplikáciu pomocou `New application` a pomenujte si ju, napríklad "RozvrhBOT".
- Choďte do sekcie `Bot` a pridajte bota.
- Bota si môžete ľubovoľne pomenovať a pridať mu profilovú fotku.
- Skopírujte si Token. Použijete ho v `config.txt`. **Nikdy ho nikomu neukazujte!**
- V sekcii `OAuth2` vyberte `bot` a ako práva nastavte `Administrator`. Tieto práva zabezpečujú, že bot bude mať prístup ku všetkým kanálom a bude môcť posielať všetky správy.
- Následne použite vygenerovaný link a pozvite bota na svoj server. Na pridanie bota musíte mať práva na spravovanie serveru.

## Konfiguračné súbory
#### Konfiguračné súbory sa vytvoria samé ak chýbaju!!
### config.txt
**Viacero argumentov oddelujte čiarkou**

**Ak nejakú hodinu nemáte, miesto nej dajte `-`**

*Príklad:* `PONDELOK=FYZ,-,FYZ` atď.
- `DISCORD_BOT_TOKEN`: Botov token ktorým sa bude prihlasovať
- `BOT_PREFIX`: Prefix pred príkazy bota, ak nebude žiadny, bot bude reagovať na všetky správy začínajúce na príkazy
- `PONDELOK` - `PIATOK`: Hodiny, ktoré máte v ten deň. Zadávajte skratky, ktoré máte v rozvrhu **Môžete nastaviť maximalne 8 hodín na jeden deň**
- `IDS` - ID rolí, ktoré budú môcť používať príkazy, nechajte prázdne aby ich mohli používať všetci
- `CASY` - Časy odkedy dokedy sú hodiny, zadávajte ich vo formáte `8:00-8:45`, zadávajte všetky hodiny od prvej po poslednú
- `DEFAULT_CHANNEL` - ID kanálu do ktorého sa automaticky budú posielať najbližšie hodiny, ak to necháte prázdne, automatické oznamovanie nebude fungovať, môžete použiť viacero ID
### linky.txt
Po vyplnení `config.txt` sa vám vytvorí ďalší súbor s jednotlivými hodinami ktoré máte. Linky na hodiny zadávajte vo formáte `FYZ=link`

## Príkazy
- `help` vypíše použitelné príkazy
- `ping` a `pong` - slúžia čisto na testovanie správnej funkcie bota, `pong` je prístupný podľa nastavenia `IDS` v `config.txt`
- `hod` - Vypíše najbližšiu hodinu
- `dalsia` - Vypíše následujúcu hodinu za najbližšou hodinou
- `rozvrh` - Vypíše celý rozvrh na konkretný deň, v rozvrhu sa da posúvať medzi jednotlivými dňami pomocou šípiek