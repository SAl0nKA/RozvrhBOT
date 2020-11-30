# RozvrhBOT
Discord bot ktorý posiela linky na online hodiny

## Konfiguračné súbory
### config.txt
- `DISCORD_BOT_TOKEN`: Botov token ktorým sa bude prihlasovať
- `BOT_PREFIX`: Prefix pred príkazy bota, ak nebude žiadny, bot bude reagovať na všetky správy začínajúce na príkazy
- `PONDELOK` - `PIATOK`: Hodiny, ktoré máte v ten deň. Zadávajte skratky, ktoré máte v rozvrhu, oddeľujte ich medzerou
- `IDS` - ID rolí, ktoré budú môcť používať príkazy, nechajte prázdne aby ich mohli používať všetci
- `CASY` - Časy odkedy dokedy sú hodiny, zadávajte ich vo formáte `8:00-8:45`, oddeľujte ich medzerou
- `DEFAULT_CHANNEL` - ID kanálu do ktorého sa automaticky budú posielať najbližšie hodiny, ak to necháte prázdne, automatické oznamovanie nebude fungovať, môžete použiť viacero ID
### hodiny.txt
Po vyplnení `config.txt` sa vám vytvorí ďalší súbor s jednotlivými hodinami ktoré máte. Linky na hodiny zadávajte vo formáte `FYZ=link`

##Pridanie bota na server
- Choďte do [aplikácií](https://discord.com/developers/applications) v developer portály na Discorde.
- Vytvorte novú aplikáciu pomocou `New application` a pomenujte si ju, napríklad "RozvrhBOT".
- Choďte do sekcie `Bot` a pridajte bota.
- Bota si môžete ľubovoľne pomenovať a pridať mu profilovú fotku.
- Skopírujte si Token. Použijete ho v `config.txt`. **Nikdy ho nikomu neukazujte!**
- V sekcii `OAuth2` vyberte `bot` a ako práva nastavte `Administrator`. Tieto práva zabezpečujú, že bot bude mať prístup ku všetkým kanálom a bude môcť posielať všetky správy.
- Následne použite vygenerovaný link a pozvite bota na svoj server. Na pridanie bota musíte mať práva na spravovanie serveru.
# RozvrhBOT
Discord bot ktorý posiela linky na online hodiny

## Konfiguračné súbory
### config.txt
- `DISCORD_BOT_TOKEN`: Botov token ktorým sa bude prihlasovať
- `BOT_PREFIX`: Prefix pred príkazy bota, ak nebude žiadny, bot bude reagovať na všetky správy začínajúce na príkazy
- `PONDELOK` - `PIATOK`: Hodiny ktoré máte v ten deň, zadávajte skratky ktoré máte v rozvrhu, oddeľujte ich medzerou
- `IDS` - ID rolí, ktoré budú môcť používať príkazy, nechajte prázdne aby ich mohli používať všetci
- `CASY` - Časy odkedy dokedy sú hodiny, zadávajte ich vo formáte `8:00-8:45`, oddeľujte ich medzerou
- `DEFAULT_CHANNEL` - ID kanálu do ktorého sa automaticky budú posielať najbližšie hodiny, ak to necháte prázdne, automatické oznamovanie nebude fungovať, môžete použiť viacero ID
### hodiny.txt
Po vyplnení `config.txt` sa vám vytvorí ďalší súbor s jednotlivými hodinami ktoré máte. Linky na hodiny zadávajte vo formáte `FYZ=link`

##Pridanie bota na server
- Choďte do [aplikácií](https://discord.com/developers/applications) v developer portály na Discorde.
- Vytvorte novú aplikáciu pomocou `New application` a pomenujte si ju, napríklad "RozvrhBOT".
- Choďte do sekcie `Bot` a pridajte bota.
- Bota si môžete ľubovoľne pomenovať a pridať mu profilovú fotku.
- Skopírujte si Token. Použijete ho v `config.txt`. **Nikdy ho nikomu neukazujte!**
- V sekcii `OAuth2` vyberte `bot` a ako práva nastavte `Administrator`. Tieto práva zabezpečujú, že bot bude mať prístup ku všetkým kanálom a bude môcť posielať všetky správy.
- Následne použite vygenerovaný link a pozvite bota na svoj server. Na pridanie bota musíte mať práva na spravovanie serveru.
