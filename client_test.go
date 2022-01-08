package pokeapi

import (
	"fmt"
	"testing"
)

func TestBerriesClient(t *testing.T) {
	client := NewClient()
	fmt.Println("=========================")
	fmt.Println(client.Berry.Get("1"))
	fmt.Println("=========================")
	fmt.Println(client.Berry.GetFirmness("1"))
	fmt.Println("=========================")
	fmt.Println(client.Berry.GetFlavour("1"))
	fmt.Println("=========================")
}

func TestResourcesClient(t *testing.T) {
	client := NewClient()
	fmt.Println("=========================")
	result, err := client.Resource.Pokedexes()
	fmt.Println(result, err)
	fmt.Println(len(*result))
	fmt.Println("=========================")
	result, err = client.Resource.Generations()
	fmt.Println(result, err)
	fmt.Println(len(*result))
	fmt.Println("=========================")
	result, err = client.Resource.EvolutionChains()
	fmt.Println(result, err)
	fmt.Println(len(*result))
	fmt.Println("=========================")
}
