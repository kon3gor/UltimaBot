package pokemon

import (
	"dev/kon3gor/ultima/internal/context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Cmd = "pokemon"

func ProcessCommand(context *context.Context) {
	pokemon, err := getRandomPokemonInfo()
	if err != nil {
		context.TextAnswer("Some error occured")
		return
	}
	file := tgbotapi.FileBytes{Name: "pokemon", Bytes: pokemon.Image}
	msg := tgbotapi.NewPhoto(context.ChatID, file)
	msg.Caption = strings.Title(pokemon.Name)
	context.CustomAnswer(msg)
}

type pokemon struct {
	Name  string `json:"name"`
	Image []byte
}

const (
	pokemonInfoEndpoint  = "https://pokeapi.co/api/v2/pokemon/%d"
	pokemonImageEndpoint = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/%d.png"
	totalPokemons        = 1008
)

func getRandomPokemonInfo() (*pokemon, error) {
	rand.Seed(time.Now().Unix())
	pokemonId := rand.Intn(totalPokemons + 1)

	return getPokemon(pokemonId)
}

func getPokemon(id int) (*pokemon, error) {
	url := fmt.Sprintf(pokemonInfoEndpoint, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	pokemon := pokemon{}
	if err := json.NewDecoder(res.Body).Decode(&pokemon); err != nil {
		return nil, err
	}

	imageBytes, err := getPokemonImage(id)
	if err != nil {
		return nil, err
	}

	pokemon.Image = imageBytes
	return &pokemon, nil
}

func getPokemonImage(id int) ([]byte, error) {
	imageUrl := fmt.Sprintf(pokemonImageEndpoint, id)
	res, err := http.Get(imageUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
