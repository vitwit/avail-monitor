package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchElectedMember(cfg *config.Config) {
// 	cemendpoint := cfg.URLEndpoint + "/pallets/elections/storage/members"
// 	resp, err := http.Get(cemendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch current elected member", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch current elected member code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.ElectedMembers
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	elecmem := response.Value[0].Who
// 	fmt.Printf("elecmem: %v\n", elecmem)
// 	// fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------", elecmem)
// 	// cem, err := strconv.ParseFloat(elecmem, 64)
// 	// if err != nil {
// 	// 	fmt.Printf("err: %v\n", err)
// 	// }
// 	// fmt.Println("current elected member value------------------------------------------->", cem)
// 	electedMember.WithLabelValues(elecmem).Set(1)
// }
