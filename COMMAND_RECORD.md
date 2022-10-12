#### Packages used 
- `go get github.com/shomali11/slacker`
- `go get github.com/krognol/go-wolfram`
- `go get github.com/joho/dotenv`
- `go get github.com/tidwall/gjson`
- `go get github.com/wit-ai/wit-go/v2`
#### Import Notes:

##### getting .env using godotenv

##### getting .env using viper

- viper commands follow pattern as

* to initialize viper

- `go get github.com/spf13/viper`

* to set config path and file

- `viper.SetConfigName(".env")`

* to read config file

- `viper.ReadInConfig()`

* to get value from config file

- `viper.Get(key)`

#### Notes on Viper-Config

- SetConfigFile explicitly defines the path, name and extension of the config file.
- Viper will use this and not check any of the config paths.
- .env - It will search for the .env file in the current directory
- viper.Get() returns an empty interface{}
- to get the underlying type of the key,
- we have to do the type assertion, we know the underlying value is string
- if we type assert to other type it will throw an error
- if type is string then ok will be true
- ok will make sure the program not break

* to check env working run

- `fmt.Println("SLACK_BOT_TOKEN : ", os.Getenv("SLACK_BOT_TOKEN"))`
- `fmt.Println("SOCKET_TOKEN : ", os.Getenv("SOCKET_TOKEN"))`

* importance of defer

- makes sure to call cancel() at all cost

#### Techs used

- WIT

* Developed by facebook
* could be trained

- Wolfram

* Ans recongnised question by WIT
