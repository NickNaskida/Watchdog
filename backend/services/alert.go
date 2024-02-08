package services

import (
	"fmt"
	"github.com/NickNaskida/Watchdog/backend/pkg/models"
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// alertCategory is a struct that holds the different alert categories.
var alertCategory = []string{"debug", "info", "warning", "error"}

// fakeMessages is a map that holds the different alert messages.
var fakeMessages = map[string][]string{
	"debug": {
		"Initializing service [%s] ...",
		"Backing up data for service [%s].",
		"Replicating data for service [%s].",
		"Running cron job for service [%s].",
		"Executing background tasks in service [%s].",
		"Parsing configuration files for service [%s].",
		"Initializing external modules for service [%s].",
		"Service [%s] configuration loaded, Verifying dependencies...",
	},
	"info": {
		"Service [%s] started.",
		"Service [%s] stopped.",
		"Service [%s] restarted.",
		"System update for service [%s] completed. Rebooting ...",
		"New version for service [%s] available. Updating later today...",
		"Maintenance scheduled for service [%s] this Sunday.",
		"Database backup successful for service [%s].",
		"Cron jobs completed for service [%s].",
		"Service [%s] is now running on port 9092.",
		"Password changed successfully for service [%s].",
	},
	"warning": {
		"Service [%s] is running out of disk space.",
		"Service [%s] is running out of memory.",
		"Service [%s] is running out of CPU.",
		"80%% of the disk space used for service [%s].",
		"Memory usage for service [%s] at 90%%.",
		"Total storage capacity for service [%s] at 95%%.",
		"Unauthorized access attempt to service [%s].",
		"Traffic spike detected for service [%s].",
		"Service [%s] is running on maintenance mode.",
		"Connection limit reached for service [%s].",
		"One of the replicas of service [%s] is not responding.",
	},
	"error": {
		"Service [%s] is not responding.",
		"Service [%s] crashed. Restarting ...",
		"Authentication failed for service [%s].",
		"Invalid credentials for service [%s].",
		"Invalid configuration file for service [%s].",
		"Network issue detected for service [%s].",
		"Database connection failed for service [%s].",
		"Billing issue detected for service [%s].",
	},
}

func NewAlert() models.Alert {
	rand.NewSource(time.Now().UnixNano())

	serviceCharId := string(letters[rand.Intn(len(letters)-1)])
	serviceNumberId := rand.Intn(9)
	serviceIdentifier := fmt.Sprintf("%s%d", serviceCharId, serviceNumberId)

	category := alertCategory[rand.Intn(len(alertCategory))]
	message := fakeMessages[category][rand.Intn(len(fakeMessages[category]))]

	alert := models.Alert{
		Id:       rand.Intn(1000),
		Category: category,
		Message:  fmt.Sprintf(message, serviceIdentifier),
	}
	return alert
}
