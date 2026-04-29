package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var (
	activeUsers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nextcloud_active_users_5m",
		Help: "Liczba aktywnych użytkowników w ciągu ostatnich 5 minut",
	})
	freeSpace = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nextcloud_free_space_bytes",
		Help: "Dostępne wolne miejsce (w bajtach)",
	})
	numFiles = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nextcloud_files_total",
		Help: "Całkowita liczba plików na serwerze",
	})
	dbSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nextcloud_database_size_bytes",
		Help: "Rozmiar bazy danych",
	})
)


func init() {
	prometheus.MustRegister(activeUsers)
	prometheus.MustRegister(freeSpace)
	prometheus.MustRegister(numFiles)
	prometheus.MustRegister(dbSize)
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Brak pliku .env, polegam na systemowych zmiennych środowiskowych")
	}

	username := os.Getenv("NEXTCLOUD_USER")
	password := os.Getenv("NEXTCLOUD_PASSWORD")
	baseURL := os.Getenv("BASE_URL")

	if username == "" || password == "" || baseURL == "" {
		fmt.Println("Brak wymaganych zmiennych środowiskowych")
		os.Exit(1)
	}

	handler := http.NewServeMux()

	
	handler.HandleFunc("GET /metrics", metricsHandler(baseURL, username, password))

	handler.HandleFunc("GET /status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	fmt.Println("Serwer uruchomiony na porcie :8082")
	fmt.Println("Odwiedź http://localhost:8082/metrics aby zobaczyć dane")

	server := http.Server{
		Addr:    ":8082",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Błąd serwera: %v\n", err)
	}
}


func metricsHandler(baseURL, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := fetchNextcloudData(baseURL, username, password)
		if err != nil {
			fmt.Printf("Błąd pobierania danych: %v\n", err)
			http.Error(w, "Nie udało się pobrać danych z NextCloud", http.StatusInternalServerError)
			return
		}

		// zmiana wartosci dla prometheus
		activeUsers.Set(float64(data.Ocs.Data.ActiveUsers.Last5Minutes))
		freeSpace.Set(float64(data.Ocs.Data.Nextcloud.System.Freespace))
		numFiles.Set(float64(data.Ocs.Data.Nextcloud.Storage.NumFiles))
		dbSize.Set(data.Ocs.Data.Server.Database.Size)

		// prometheus.handler
		promhttp.Handler().ServeHTTP(w, r)
	}
}

func fetchNextcloudData(baseURL, username, password string) (*Response, error) {
	url := baseURL + "/ocs/v2.php/apps/serverinfo/api/v1/info?format=json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("OCS-APIRequest", "true")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("nieoczekiwany status z API: %d", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// z githuba
type Response struct {
	Ocs struct {
		Meta struct {
			Status     string `json:"status"`
			Statuscode int    `json:"statuscode"`
			Message    string `json:"message"`
		} `json:"meta"`
		Data struct {
			Nextcloud struct {
				System struct {
					Version             string    `json:"version"`
					Theme               string    `json:"theme"`
					EnableAvatars       string    `json:"enable_avatars"`
					EnablePreviews      string    `json:"enable_previews"`
					MemcacheLocal       string    `json:"memcache.local"`
					MemcacheDistributed string    `json:"memcache.distributed"`
					FilelockingEnabled  string    `json:"filelocking.enabled"`
					MemcacheLocking     string    `json:"memcache.locking"`
					Debug               string    `json:"debug"`
					Freespace           int64     `json:"freespace"`
					Cpuload             []float64 `json:"cpuload"`
					MemTotal            int       `json:"mem_total"`
					MemFree             int       `json:"mem_free"`
					SwapTotal           int       `json:"swap_total"`
					SwapFree            int       `json:"swap_free"`
					Apps                struct {
						NumInstalled        int `json:"num_installed"`
						NumUpdatesAvailable int `json:"num_updates_available"`
						AppUpdates          struct {
							FilesAntivirus string `json:"files_antivirus"`
						} `json:"app_updates"`
					} `json:"apps"`
				} `json:"system"`
				Storage struct {
					NumUsers         int `json:"num_users"`
					NumFiles         int `json:"num_files"`
					NumStorages      int `json:"num_storages"`
					NumStoragesLocal int `json:"num_storages_local"`
					NumStoragesHome  int `json:"num_storages_home"`
					NumStoragesOther int `json:"num_storages_other"`
				} `json:"storage"`
				Shares struct {
					NumShares               int    `json:"num_shares"`
					NumSharesUser           int    `json:"num_shares_user"`
					NumSharesGroups         int    `json:"num_shares_groups"`
					NumSharesLink           int    `json:"num_shares_link"`
					NumSharesLinkNoPassword int    `json:"num_shares_link_no_password"`
					NumFedSharesSent        int    `json:"num_fed_shares_sent"`
					NumFedSharesReceived    int    `json:"num_fed_shares_received"`
					Permissions41           string `json:"permissions_4_1"`
				} `json:"shares"`
			} `json:"nextcloud"`
			Server struct {
				Webserver string `json:"webserver"`
				Php       struct {
					Version           string `json:"version"`
					MemoryLimit       int    `json:"memory_limit"`
					MaxExecutionTime  int    `json:"max_execution_time"`
					UploadMaxFilesize int    `json:"upload_max_filesize"`
				} `json:"php"`
				Database struct {
					Type    string  `json:"type"`
					Version string  `json:"version"`
					Size    float64 `json:"size,string"`
				} `json:"database"`
			} `json:"server"`
			ActiveUsers struct {
				Last5Minutes int `json:"last5minutes"`
				Last1Hour    int `json:"last1hour"`
				Last24Hours  int `json:"last24hours"`
			} `json:"activeUsers"`
		} `json:"data"`
	} `json:"ocs"`
}