package palrcon

import (
	"fmt"
	"testing"
)

func TestPalRCON(t *testing.T) {
	palCron := NewPalRCON("192.168.1.6:25575", "7210")
	//palCron.SetTimeout(5 * time.Second)
	players, err := palCron.GetPlayers()
	if err != nil {
		fmt.Printf("failed, err: %v", err)
		return
	}
	fmt.Printf("success, players: %v", players)
}
