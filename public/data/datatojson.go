package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Card struct {
	ID                 int               `json:"id"`
	CardType           string            `json:"cardtype"`
	Clan               string            `json:"clan"`
	Critical           int               `json:"critical"`
	DesignIllus        string            `json:"designillus"`
	Effect             string            `json:"effect"`
	Flavor             string            `json:"flavor"`
	Format             string            `json:"format"`
	Grade              string            `json:"grade"`
	Illust             string            `json:"illust"`
	IllustColor        string            `json:"illustcolor"`
	Illust2            string            `json:"illust2"`
	Illust3            string            `json:"illust3"`
	Illust4            string            `json:"illust4"`
	Illust5            string            `json:"illust5"`
	ImageURLEn         string            `json:"imageurlen"`
	ImageURLJp         string            `json:"imageurljp"`
	ImaginaryGift      string            `json:"imaginarygift"`
	Italian            string            `json:"italian"`
	Kana               string            `json:"kana"`
	Kanji              string            `json:"kanji"`
	Korean             string            `json:"korean"`
	LimitationText     string            `json:"limitationtext"`
	MangaIllust        string            `json:"mangaillust"`
	Name               string            `json:"name"`
	Nation             string            `json:"nation"`
	Note               string            `json:"note"`
	OtherNames         string            `json:"othernames"`
	Phonetic           string            `json:"phonetic"`
	Power              int               `json:"power"`
	Race               string            `json:"race"`
	RideSkill          string            `json:"rideskill"`
	Sets               []string          `json:"sets"`
	TournamentStatuses map[string]string `json:"tournamentstatuses"`
	Shield             int               `json:"shield"`
	Skill              string            `json:"skill"`
	Thai               string            `json:"thai"`
	Translation        string            `json:"translation"`
	TriggerEffect      string            `json:"triggereffect"`
}

type CardData struct {
	Data []Card `json:"data"`
}

func main() {
	apiURL := "https://card-fight-vanguard-api.ue.r.appspot.com/api/v1/cards?page="
	page := 1
	allData := []Card{}

	for {
		// Fetch data from API
		resp, err := http.Get(fmt.Sprintf("%s%d", apiURL, page))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		defer resp.Body.Close()

		// Check response status
		if resp.StatusCode == 200 {
			var data CardData
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading response body: %v\n", err)
				break
			}
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Printf("Error unmarshalling JSON: %v\n", err)
				break
			}

			// Check if there's data in the current page
			if len(data.Data) == 0 {
				break
			}

			// Append data from each page
			allData = append(allData, data.Data...)

			// Increment page number
			page++
		} else {
			fmt.Printf("Error: %d\n", resp.StatusCode)
			break
		}
	}

	// Save all data to JSON file
	file, err := os.Create("all_cards.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(allData)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}

	fmt.Printf("Data saved to 'all_cards.json'. Total records: %d\n", len(allData))
}
