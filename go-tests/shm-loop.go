package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

func getWiegeDatenShm(shmId int) ([]byte, error) {
	const shmSize = 40
	loopAnzahl := 0
	nostillCount := 0

	// Shared Memory öffnen
	shm, err := syscall.Shmget(shmId, shmSize, 0666)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Öffnen von Shared Memory: %v", err)
	}

	// Shared Memory anhängen
	addr, err := syscall.Shmat(shm, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Anhängen von Shared Memory: %v", err)
	}
	defer syscall.Shmdt(addr)

	// Pointer in Byte-Slice konvertieren
	data := (*[shmSize]byte)(unsafe.Pointer(addr))[:]

	for loopAnzahl < 30 {
		loopAnzahl++
		time.Sleep(1 * time.Second)

		status := data[21]

		// Wenn Index 0 nicht 0 ist, weiter
		if data[0] != 0 {
			continue
		}

		switch status {
		case 1:
			return data, nil
		case 2:
			fmt.Println("Fehler vom Disomat!")
			return nil, fmt.Errorf("Fehler vom Disomat!")
		case 3:
			nostillCount++
			if nostillCount >= 10 {
				fmt.Println("Waage nach 10 Versuchen nicht im Stillstand!")
				return nil, fmt.Errorf("Waage nach 10 Versuchen nicht im Stillstand!")
			}
		default:
			fmt.Println("Warten auf Gewicht...")
		}
	}
	return nil, fmt.Errorf("Nach 30 Versuchen kein Gewicht erhalten")
}

func main() {
	shmId := 2007

	// Daten aus dem Shared Memory lesen
	result, err := getWiegeDatenShm(shmId)
	if err != nil {
		fmt.Println("Fehler:", err)
		os.Exit(1)
	}

	fmt.Printf("Rückgabe von getWiegeDatenShm: %v\n", result)
}
