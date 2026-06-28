# Krypto

<p>
  <a href="https://docs.coingecko.com"><img src="https://avatars.githubusercontent.com/u/7111837?s=200&v=4" alt="coin gecko"></a>
</p>
**Krypto cli** - приложение для отслеживания рынка криптовалюты
![Demo](assets/krypto.webm)
Приведенный выше пример выполняется из одного 

## API ENDPOINTS 
`/coins/list` — нужен один раз при настройке программы (или при обновлении списка монет), чтобы создать словарь "Символ -> ID".

`/simple/price` — нужен каждый раз, когда пользователь хочет узнать актуальную цену.
