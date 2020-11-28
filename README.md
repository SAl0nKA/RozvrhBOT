# RozvrhBOT
Discord bot ktorý posiela linky na online hodiny

#### Konfiguračné súbory
## Premenné
### config.txt
- `DISCORD_BOT_TOKEN`: Botov token ktorým sa bude prihlasovať
- `BOT_PREFIX`: Prefix pred príkazy bota, ak nebude žiadný, bot bude reagovať na všetky správy začinajúce na príkazy
- `PONDELOK` - `PIATOK`: Hodiny ktoré máte v ten deň, zadávajte skratky ktoré máte v rozvrhu, oddeľujte ich medzerou
- `IDS` - ID rolí, ktoré budú môcť používať príkazy, nechajte prázdne aby ich mohli používať všetci
- `CASY` - Časy odkedy dokedy sú hodiny, zadávajte ich vo formáte `8:00-8:45`, oddeľujte ich medzerou
- `DEFAULT_CHANNEL` - ID kanálu do ktorého sa automaticky budú posielať najbližšie hodiny, ak to necháte prázdne, automatické oznamovanie nebude fungovať, môžete použiť viacero ID
### hodiny.txt
Po vyplnení `config.txt` sa vám vytvorí ďalší súbor s jednotlivými hodinami ktoré máte. Linky na hodiny zadávajte vo formáte `FYZ=link`
