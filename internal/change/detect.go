package change

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/nimaism/trackit/internal/config"
	"github.com/nimaism/trackit/internal/notifier"
	"github.com/nimaism/trackit/internal/store"
	"github.com/nimaism/trackit/pkg/network"

	"github.com/twmb/murmur3"
)

type Detector struct {
	Notifiers  []notifier.Notifier
	Store      store.Store
	Config     *config.Config
	HttpClient *http.Client
}

func New(notifiers []notifier.Notifier, store store.Store, config *config.Config) *Detector {
	return &Detector{
		Notifiers:  notifiers,
		Store:      store,
		Config:     config,
		HttpClient: network.InitHTTPClient(config.Network.TimeoutSec, config.Network.VerifySSL, config.Network.DisableRedirect),
	}
}

func (d *Detector) Start() {
	ticker := time.NewTicker(time.Duration(d.Config.Interval) * time.Minute)

	for {
		select {
		case <-ticker.C:
			d.detectChanges()
		}
	}
}

func (d *Detector) detectChanges() {
	log.Println("Starting change detection...")

	urls, err := d.loadURLs()
	if err != nil {
		log.Fatalf("Failed to load URLs from file: %v", err)
		return
	}

	if err := d.Store.LoadRecord(); err != nil {
		log.Fatalf("Failed to load records from store: %v", err)
		return
	}

	var wg sync.WaitGroup
	urlChan := make(chan string, len(urls))
	notifChan := make(chan string, len(urls))

	for i := 0; i < int(d.Config.Concurrency); i++ {
		wg.Add(1)
		go d.checkURL(urlChan, notifChan, &wg)
	}

	go func() {
		defer close(urlChan)
		for _, url := range urls {
			urlChan <- url
		}
	}()

	go func() {
		wg.Wait()
		close(notifChan)
		d.Store.SaveRecord()
	}()

	for url := range notifChan {
		notifier.Alert(d.Notifiers, url)
		log.Printf("Change detected at URL: %s", url)
	}
}

func (d *Detector) loadURLs() ([]string, error) {
	file, err := os.Open(d.Config.URLsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if network.CheckValidURL(text) {
			urls = append(urls, text)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func (d *Detector) checkURL(urlChan chan string, notifChan chan string, wg *sync.WaitGroup) error {
	defer wg.Done()
	for url := range urlChan {
		resp, err := d.HttpClient.Get(url)
		if err != nil {
			log.Printf("Failed to fetch URL %s: %v", url, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body for URL %s: %v", url, err)
			continue
		}

		hash := murmur3.Sum32(body)

		if value, exists := d.Store.Data[url]; exists && value != hash {
			notifChan <- url
		}

		d.Store.Data[url] = hash
	}
	return nil
}
