package pokeapi

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

// Struct defining how the reponse data for a Map Area should be interpreted
type PokemonMapArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonInArea struct {
	Encounter_Method_Rates []struct {
		Encounter_Method struct {
			Encounter_Method_Name string `json:"name"`
			Encounter_Method_URL  string `json:"url"`
		} `json:"encounter_method"`
		Version_Details []struct {
			Rate    int `json:"rate"`
			Version struct {
				Version_Name string `json:"name"`
				Version_URL  string `json:"url"`
			}
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Game_Index int `json:"game_index"`
	ID         int `json:"id"`
	Location   struct {
		Location_Name string `json:"name"`
		Location_URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Language_Name string `json:"name"`
			Language_URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Encounters []struct {
		Pokemon struct {
			Pokemon_Name string `json:"name"`
			Pokemon_URL  string `json:"url"`
		} `json:"pokemon"`
		Pokemon_Version_Details []struct {
			Encounter_Details []struct {
				Chance           int    `json:"chance"`
				Condition_Values string `json:"condition_values"`
				Max_Level        int    `json:"max_level"`
				Method           struct {
					Method_Name string `json:"walk"`
					Method_URL  string `json:"url"`
				} `json:"method"`
				Min_Level int `json:"min_level"`
			} `json:"encounter_details"`
			Max_Chance int `json:"max_chance"`
			Version    struct {
				Version_Name string `json:"name"`
				Version_URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type PokemonSpecies struct {
	Base_happiness int `json:"base_happiness"`
	Capture_rate   int `json:"capture_rate"`
	Color          struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"color"`
	Egg_groups []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"egg_groups"`
	Evolution_chain struct {
		Url string `json:"url"`
	} `json:"evolution_chain"`
	Evolves_from_species struct {
		Name string `json:"name"`
		Nrl  string `json:"url"`
	} `json:"evolves_from_species"`
	Flavor_text_entries []struct {
		Flavor_text string `json:"flavor_text"`
		Language    struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
	Form_descriptions []byte `json:"form_descriptions"`
	Forms_switchable  bool   `json:"forms_switchable"`
	Gender_rate       int    `json:"gender_rate"`
	Genera            []struct {
		Genus    string `json:"genus"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"genera"`
	Generation struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"generation"`
	Growth_rate struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"growth_rate"`
	Habitat                string `json:"habitat"`
	Has_gender_differences bool   `json:"has_gender_differences"`
	Hatch_counter          int    `json:"hatch_counter"`
	Id                     int    `json:"id"`
	Is_baby                bool   `json:"is_baby"`
	Is_legendary           bool   `json:"is_legendary"`
	Is_mythical            bool   `json:"is_mythical"`
	Name                   string `json:"name"`
	Names                  []struct {
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Order               int `json:"order"`
	Pal_park_encounters []struct {
		Area struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"area"`
		Base_score int `json:"base_score"`
		Rate       int `json:"rate"`
	} `json:"pal_park_encounters"`
	Pokedex_numbers []struct {
		Entry_number int `json:"entry_number"`
		Pokedex      struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokedex"`
	} `json:"pokedex_numbers"`
	Shape struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"shape"`
	Varieties []struct {
		Is_default bool `json:"is_default"`
		Pokemon    struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"varieties"`
}

type PokemonDetailedInformation struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
		Is_hidden bool `json:"is_hidden"`
		Slot      int  `json:"slot"`
	} `json:"abilities"`
	Base_experience int `json:"base_experience"`
	Cries           struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"forms"`
	Game_indices []struct {
		Game_index int `json:"game_index"`
		Version    struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height     int `json:"height"`
	Held_items []struct {
		Item struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"item"`
	} `json:"held_items"`
	Version_details []struct {
		Rarity  int `json:"rarity"`
		Version struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"version_details"`
	Id                       int    `json:"id"`
	Is_default               bool   `json:"is_default"`
	Location_area_encounters string `json:"location_area_encounters"`
	Moves                    []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`
	} `json:"moves"`
	Version_group_details []struct {
		Level_learned_at  int `json:"level_learned_at"`
		Move_learn_method struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move_learn_method"`
		Version_group struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version_group"`
	} `json:"version_group_details"`
	Name           string `json:"name"`
	Order          int    `json:"order"`
	Past_abilities []byte `json:"past_abilities"`
	Past_types     []byte `json:"past_types"`
	Species        struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
	Sprites map[string]interface{} `json:"sprites"`
	Stats   []struct {
		Base_stat int `json:"base_stat"`
		Effort    int `json:"effort"`
		Stat      struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

var MapAreaResults PokemonMapArea
var EncounterResults PokemonInArea
var PokemonSpeciesInformation PokemonSpecies
var PokemonInspectionInformation PokemonDetailedInformation

// Function to request data from the Pokdemon Database API on a map area
// Takes in an integer defining the direction (positive is forward, negative is backward)
// Takes in an initialized cache map
// Will look for the request first in the map and will return the cached value if found
// Otherwise, will use the formatted request in HTML request to API and return the value
// Returns the Body of the request from the cache or HTML message if successful
// Returns nothing and the error if both the cache and HTML request fail
func GetAreaLocation(dir int, c *pokecache.Cache) ([]string, error) {
	//The request will return a list of locations, this is storage
	retrievedAreas := make([]string, 0)
	var httpRequest string
	//1 is going forward
	if dir == 1 {
		if MapAreaResults.Next != "" {
			fmt.Println("Next")
			httpRequest = MapAreaResults.Next
		} else {
			httpRequest = "https://pokeapi.co/api/v2/location-area/?limit=20"
		}
	} else {
		if MapAreaResults.Previous != "" {
			httpRequest = MapAreaResults.Previous
		} else {
			return retrievedAreas, fmt.Errorf("Cannot retrieve previous; at list beginning\n")
		}
	}
	//Check the cache for the request
	val, found := c.Get(httpRequest)
	if !found {
		fmt.Println("Not found in cache, retrieving")
		res, error := http.Get(httpRequest)
		if error != nil {
			log.Fatal(error)
			return retrievedAreas, error
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		//Add the new request to the cache
		c.Add(httpRequest, body)
		marshalErr := json.Unmarshal(body, &MapAreaResults)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}

		//Format the list with the retrieved values
		for i := 0; i < 20; i++ {
			retrievedAreas = append(retrievedAreas, MapAreaResults.Results[i].Name)
		}
	} else {
		fmt.Println("Found in cache, retreiving local")
		marshalErr := json.Unmarshal(val, &MapAreaResults)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}

		for i := 0; i < 20; i++ {
			retrievedAreas = append(retrievedAreas, MapAreaResults.Results[i].Name)
		}
	}

	return retrievedAreas, nil
}

func GetPokemonInArea(area string, c *pokecache.Cache) ([]string, error) {
	var httpRequest = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	retreivedEncounters := make([]string, 0)
	val, found := c.Get(httpRequest)
	if !found {
		fmt.Println("Not found in cache, retrieving")
		res, error := http.Get(httpRequest)
		if error != nil {
			log.Fatal(error)
			return retreivedEncounters, error
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		//Add the new request to the cache
		c.Add(httpRequest, body)
		marshalErr := json.Unmarshal(body, &EncounterResults)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}

		for _, encounters := range EncounterResults.Encounters {
			retreivedEncounters = append(retreivedEncounters, encounters.Pokemon.Pokemon_Name)
		}

	} else {
		fmt.Println("Found in cache, retreiving local")
		marshalErr := json.Unmarshal(val, &EncounterResults)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}

		for _, encounters := range EncounterResults.Encounters {
			retreivedEncounters = append(retreivedEncounters, encounters.Pokemon.Pokemon_Name)
		}
	}

	return retreivedEncounters, nil

}

// Catch Pokemon Function
// Arguments:
// Pokemon (string) : The pokemon to be caught as a character string
// C (PokeCache) : Cache storing all encountered pokemon in the current location
// Returns:
// Success Message (string) : A processed string indicating whether the pokemon was caught or escaped
// error (error) : An error occured and returning from the function
func CatchPokemon(pokemon string, c *pokecache.Cache, pokedex map[string][]byte) (string, error) {
	var httpRequest = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%s", pokemon)
	val, found := c.Get(httpRequest)
	if !found {
		fmt.Println("Not found in cache, retrieving")
		res, error := http.Get(httpRequest)
		if error != nil {
			log.Fatal(error)
			return "Could not find pokemon", error
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		//Add the new request to the cache
		c.Add(httpRequest, body)
		marshalErr := json.Unmarshal(body, &PokemonSpeciesInformation)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}

	} else {
		fmt.Println("Found in cache, retreiving local")
		marshalErr := json.Unmarshal(val, &PokemonSpeciesInformation)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}
	}
	a := float64(PokemonSpeciesInformation.Capture_rate) * 1.5
	shake_probability := float64(1048560) / math.Sqrt(math.Sqrt(float64(16711680)/a))
	captured := true
	fmt.Printf("%f is the modified capture rate and %f is the shake_probability", a, shake_probability)
	for i := 0; i < 4; i++ {
		random_number := rand.Float64() * 65535.0
		fmt.Printf("Random Generation: %f\n", random_number)
		if random_number > shake_probability {
			captured = false
			break
		}
	}
	var captured_string string
	if captured {
		captured_string = fmt.Sprintf("%s has been captured!", pokemon)

		var httpRequest = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
		val, found := c.Get(httpRequest)
		if !found {
			fmt.Println("Not found in cache, retrieving")
			res, error := http.Get(httpRequest)
			if error != nil {
				log.Fatal(error)
				return "Could not find pokemon", error
			}
			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
			}
			if err != nil {
				log.Fatal(err)
			}
			//Add the new request to the cache
			pokedex[pokemon] = body
			marshalErr := json.Unmarshal(body, &PokemonInspectionInformation)

			if marshalErr != nil {
				fmt.Println(marshalErr)
			}

		} else {
			fmt.Println("Found in cache, retreiving local")
			marshalErr := json.Unmarshal(val, &PokemonInspectionInformation)

			if marshalErr != nil {
				fmt.Println(marshalErr)
			}
		}
	} else {
		captured_string = fmt.Sprintf("%s has escaped!", pokemon)
	}
	return captured_string, nil
}

func InspectPokemon(pokemon string, p map[string][]byte) (string, error) {
	var information_string string
	val, found := p[pokemon]
	if !found {
		fmt.Println("You aint caught that yet")

	} else {
		fmt.Println("Found in cache, retreiving local")
		marshalErr := json.Unmarshal(val, &PokemonInspectionInformation)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}

		information_string = fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n\t-hp: %d\n\t-attack: %d\n\t-defense: %d\n\t-special-attack: %d\n\t-special-defense: %d\n\t-speed: %d\n",
			PokemonInspectionInformation.Name, PokemonInspectionInformation.Height, PokemonInspectionInformation.Weight,
			PokemonInspectionInformation.Stats[0].Base_stat, PokemonInspectionInformation.Stats[1].Base_stat,
			PokemonInspectionInformation.Stats[2].Base_stat, PokemonInspectionInformation.Stats[3].Base_stat,
			PokemonInspectionInformation.Stats[4].Base_stat, PokemonInspectionInformation.Stats[5].Base_stat)
		information_string += fmt.Sprintf("Types: \n\t-%s", PokemonInspectionInformation.Types[0].Type.Name)
		if len(PokemonInspectionInformation.Types) > 1 {
			information_string += fmt.Sprintf("\n\t-%s", PokemonInspectionInformation.Types[1].Type.Name)
		}
	}
	return information_string, nil
}

func ExplorePokedex(p map[string][]byte) (string, error) {
    pokedex_list := "Your Pokemon: \n"
	for _, pokemon_info := range p {
		marshalErr := json.Unmarshal(pokemon_info, &PokemonInspectionInformation)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}

		pokedex_list += fmt.Sprintf("\t-%s\n", PokemonInspectionInformation.Name)
	}
	return pokedex_list, nil
}
