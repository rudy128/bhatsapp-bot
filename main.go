package main

import (
	// "bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"math/rand"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	// "io/ioutil"
	// "log"
	// "encoding/json"
)

func Dbase() {
	var err error
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err = sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
}

func GetDb() *sqlstore.Container {
	return container
}

func main() {
	Dbase()

	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "ERROR", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		// fmt.Println("Hello1")
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			// fmt.Println("Hello2")

			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				fmt.Println("QR code:", evt.Code)
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// fmt.Println("Hello3")

			} else {
				fmt.Println("Login event:", evt.Event)
				// fmt.Println("Hello4")

			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		// fmt.Println("Hello5")

		if err != nil {
			panic(err)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()

}

func otpGenerator() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(9000) + 1000
}

// 120363320125534378@g.us a blank

// 120363378515788831@g.us Test Parent

// 120363379194795613@g.us Test Announcement
