/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  `This command fetches a random dad joke from the icanhazdadkoke api`,
	Run: func(cmd *cobra.Command, args []string) {

		jokeTerm, _ := cmd.Flags().GetString("term")
		if jokeTerm != "" {
			getRandomJokeWithTerm(jokeTerm)
		} else {
			getRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.PersistentFlags().String("term", "", "A search term for a dad joke.")

}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}
	err := json.Unmarshal(responseBytes, &joke)
	if err != nil {
		log.Printf("Could not unmershal response : %v", err)
	}
	fmt.Println(string(joke.Joke))
}

func getRandomJokeWithTerm(jokeTerm string) {
	fmt.Printf("You searched for a joke with the term: %v", jokeTerm)
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("Could not request the Api")
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Could make reauest : %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Could not read request Body : %v", err)
	}

	return responseBytes

}
