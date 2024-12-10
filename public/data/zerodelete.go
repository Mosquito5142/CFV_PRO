package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Card struct {
	ID     int    `json:"id"`
	Format string `json:"format"`
	// เพิ่มฟิลด์อื่นๆ ที่คุณต้องการเก็บไว้ในโครงสร้างนี้
	CardType           string            `json:"cardtype"`
	Clan               string            `json:"clan"`
	Critical           int               `json:"critical"`
	DesignIllus        string            `json:"designillus"`
	Effect             string            `json:"effect"`
	Flavor             string            `json:"flavor"`
	Grade              string            `json:"grade"`
	Illust             string            `json:"illust"`
	IllustColor        string            `json:"illustcolor"`
	Illust2            string            `json:"illust2"`
	Illust3            string            `json:"illust3"`
	Illust4            string            `json:"illust4"`
	Illust5            string            `json:"illust5"`
	ImageURLJP         string            `json:"imageurljp"`
	ImageURLEn         string            `json:"imageurlen"`
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

func main() {
	const jsonFile = "all_cards.json"

	// โหลดข้อมูลจากไฟล์ JSON
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var allData []Card
	if err := json.Unmarshal(file, &allData); err != nil {
		log.Fatal(err)
	}

	// ลบข้อมูลที่มี format เป็น "Vanguard ZERO"
	var filteredData []Card
	for _, card := range allData {
		if card.Format != "Vanguard ZERO" {
			filteredData = append(filteredData, card)
		}
	}

	// บันทึกข้อมูลที่เหลือลงในไฟล์ JSON
	file, err = json.MarshalIndent(filteredData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(jsonFile, file, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data saved to '%s'. Total records: %d\n", jsonFile, len(filteredData))
}
