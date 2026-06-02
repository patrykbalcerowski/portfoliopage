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

// Kod HTML Twojej strony startowej
const landingPageHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Patryk Balcerowski | Cloud-Native Portfolio</title>
    <style>
        :root {
            --bg-color: #0f172a;
            --card-bg: #1e293b;
            --text-main: #f8fafc;
            --text-muted: #94a3b8;
            --accent-blue: #38bdf8;
            --accent-green: #10b981;
            --accent-orange: #f59e0b;
            --border-color: #334155;
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }

        body {
            font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            background-color: var(--bg-color);
            color: var(--text-main);
            line-height: 1.6;
            padding: 2rem 1rem;
        }

        .container {
            max-width: 900px;
            margin: 0 auto;
        }

        header {
            text-align: center;
            margin-bottom: 3rem;
            animation: fadeIn 1s ease-out;
        }

        h1 {
            font-size: 2.5rem;
            font-weight: 800;
            letter-spacing: -0.025em;
            margin-bottom: 0.5rem;
            color: var(--text-main);
        }

        h1 span {
            color: var(--accent-blue);
        }

        .subtitle {
            color: var(--text-muted);
            font-size: 1.125rem;
        }

        /* Sekcja About */
        .section-card {
            background-color: var(--card-bg);
            border: 1px solid var(--border-color);
            border-radius: 12px;
            padding: 2rem;
            margin-bottom: 2rem;
            box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
        }

        .section-title {
            font-size: 1.25rem;
            font-weight: 600;
            color: var(--accent-blue);
            margin-bottom: 1rem;
            display: flex;
            align-items: center;
            gap: 0.5rem;
        }

        .about-text {
            color: var(--text-muted);
            font-size: 0.95rem;
            text-align: justify;
        }

        .about-text p {
            margin-bottom: 1rem;
        }

        /* Architektura (Drzewko) */
        .architecture-flow {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: 1rem;
            margin-top: 1.5rem;
        }

        .node {
            background-color: var(--bg-color);
            border: 1px solid var(--border-color);
            padding: 1rem 2rem;
            border-radius: 8px;
            width: 100%;
            max-width: 400px;
            text-align: center;
            font-weight: 500;
            position: relative;
        }

        .node-primary { border-color: var(--accent-blue); color: var(--accent-blue); }
        .node-gitops { border-color: var(--accent-green); color: var(--accent-green); }
        .node-obs { border-color: var(--accent-orange); color: var(--accent-orange); }

        .arrow {
            width: 2px;
            height: 20px;
            background-color: var(--border-color);
        }

        .node-group {
            display: flex;
            gap: 1rem;
            width: 100%;
            justify-content: center;
            flex-wrap: wrap;
        }

        .node-group .node {
            flex: 1;
            min-width: 200px;
        }

        /* Przyciski */
        .links-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 1rem;
            margin-top: 2rem;
        }

        .btn {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 0.5rem;
            padding: 1rem;
            border-radius: 8px;
            text-decoration: none;
            font-weight: 600;
            transition: all 0.2s;
            color: #fff;
            text-align: center;
        }

        .btn:hover {
            transform: translateY(-2px);
            filter: brightness(1.1);
        }

        .btn-github { background-color: #24292e; border: 1px solid #444c56; }
        .btn-argo { background-color: #10b981; color: #064e3b; }
        .btn-grafana { background-color: #f59e0b; color: #78350f; }

        @keyframes fadeIn {
            from { opacity: 0; transform: translateY(-10px); }
            to { opacity: 1; transform: translateY(0); }
        }

        @media (max-width: 600px) {
            .node-group { flex-direction: column; }
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>Patryk <span>Balcerowski</span></h1>
            <p class="subtitle">IT & DevOps Engineer | Kubernetes Enthusiast</p>
        </header>

        <section class="section-card">
            <h2 class="section-title">
                <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path></svg>
                About Me
            </h2>
            <div class="about-text">
                <p>
                    IT Engineer with over 3 years of hands-on experience in managing Linux server infrastructures, service provisioning, and automation within the FinTech/Crypto sector. Proficient in scripting (Bash) and programming (Go, Python, JS) to streamline deployments and operations. Experienced in full-cycle application deployments, system testing, and providing on-call support for critical ATM infrastructure.
                </p>
                <p>
                    I have a strong background in both software engineering and hardware troubleshooting, seeking to leverage my operations and development skills as a DevOps Engineer.
                </p>
            </div>
        </section>

        <section class="section-card">
            <h2 class="section-title">
                <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z"></path></svg>
                Home Lab Architecture
            </h2>
            <div class="about-text">
                <p>This page is served directly from a custom <strong>Go application</strong> running inside a <strong>K3s Kubernetes cluster</strong> (hosted on Proxmox). The cluster is fully automated using the <strong>GitOps</strong> methodology and exposed securely to the internet.</p>
            </div>
            
            <div class="architecture-flow">
                <div class="node node-primary">☁️ Cloudflare Tunnel (Ingress)</div>
                <div class="arrow"></div>
                <div class="node node-primary">🚪 Kubernetes Gateway API (Nginx)</div>
                <div class="arrow"></div>
                
                <div class="node-group">
                    <div class="node node-gitops">
                        <strong>🐙 ArgoCD</strong><br>
                        <span style="font-size: 0.8em; color: var(--text-muted);">Watches GitHub & syncs state</span>
                    </div>
                    <div class="node node-obs">
                        <strong>📊 Prometheus & Grafana</strong><br>
                        <span style="font-size: 0.8em; color: var(--text-muted);">Scrapes metrics every 15s</span>
                    </div>
                </div>
                
                <div class="arrow"></div>
                <div class="node" style="border-color: #e2e8f0; color: #e2e8f0;">
                    <strong>🚀 Custom Go Exporter</strong><br>
                    <span style="font-size: 0.8em; color: var(--text-muted);">Serves this page & exposes /metrics</span>
                </div>
                
                <div class="arrow"></div>
                <div class="node" style="border-color: #0082c9; color: #38bdf8;">
                    <strong>📁 NextCloud Instance</strong><br>
                    <span style="font-size: 0.8em; color: var(--text-muted);">Target data source</span>
                </div>
            </div>
        </section>

        <section class="links-grid">
            <a href="https://github.com/patrykbalcerowski/portfoliopage" class="btn btn-github" target="_blank">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 24 24"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12c0 4.42 2.865 8.166 6.839 9.489.5.092.682-.217.682-.482 0-.237-.008-.866-.013-1.7-2.782.603-3.369-1.34-3.369-1.34-.454-1.156-1.11-1.462-1.11-1.462-.908-.62.069-.608.069-.608 1.003.07 1.531 1.03 1.531 1.03.892 1.529 2.341 1.087 2.91.831.092-.646.35-1.086.636-1.336-2.22-.253-4.555-1.11-4.555-4.943 0-1.091.39-1.984 1.029-2.683-.103-.253-.446-1.27.098-2.647 0 0 .84-.269 2.75 1.025A9.578 9.578 0 0112 6.836c.85.004 1.705.114 2.504.336 1.909-1.294 2.747-1.025 2.747-1.025.546 1.379.203 2.394.1 2.647.64.699 1.028 1.592 1.028 2.683 0 3.842-2.339 4.687-4.566 4.935.359.309.678.919.678 1.852 0 1.336-.012 2.415-.012 2.743 0 .267.18.578.688.48C19.138 20.161 22 16.416 22 12c0-5.523-4.477-10-10-10z"></path></svg>
                GitHub Repo
            </a>
            <a href="https://argocd.balc-tech.pl" class="btn btn-argo" target="_blank">
                🐙 ArgoCD Panel
            </a>
            <a href="https://grafana.balc-tech.pl" class="btn btn-grafana" target="_blank">
                📊 Grafana Dashboards
            </a>
        </section>
    </div>
</body>
</html>
`

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

	// Istniejące endpointy
	handler.HandleFunc("GET /metrics", metricsHandler(baseURL, username, password))
	handler.HandleFunc("GET /status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	// NOWY ENDPOINT: Serwowanie strony startowej
	handler.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// Zapobiegamy wyłapywaniu zapytań o np. /favicon.ico jako strony głównej
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, landingPageHTML)
	})

	fmt.Println("Serwer uruchomiony na porcie :8082")
	fmt.Println("Odwiedź http://localhost:8082/ aby zobaczyć stronę główną")
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
