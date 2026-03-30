package cookie

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	CookieAPIURL = "https://rta.zhltech.net/guangyixinmedia/report/dhh/cookie"
	//CookieAPIURL     = "http://127.0.0.1:8888/report/dhh/cookie"
	RefreshInterval  = 1 * time.Minute  // 每30分钟刷新一次
	InitialRetryWait = 10 * time.Second // 初始重试等待时间
)

// CookieResponse 接口返回数据结构
type CookieResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cookie  string `json:"cookie"`
	Csrf    string `json:"csrf"`
}

// Manager Cookie 管理器
type Manager struct {
	cookie     string
	csrf       string
	lastUpdate time.Time
	mu         sync.RWMutex
	stopCh     chan struct{}
}

// NewManager 创建 Cookie 管理器，启动时立即获取凭证，失败则定时重试
func NewManager() *Manager {
	m := &Manager{
		stopCh: make(chan struct{}),
	}

	if err := m.fetchCookie(); err != nil {
		log.Printf("[cookie] 初始化获取凭证失败: %v，将在后台持续重试", err)
	}

	go m.autoRefresh()
	return m
}

// GetCredentials 返回当前 cookie 和 csrf，未初始化时返回错误
func (m *Manager) GetCredentials() (cookie, csrf string, err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.cookie == "" || m.csrf == "" {
		return "", "", fmt.Errorf("凭证尚未就绪，请稍后重试")
	}
	return m.cookie, m.csrf, nil
}

// fetchCookie 从远程接口获取 cookie 和 csrf
func (m *Manager) fetchCookie() error {
	req, err := http.NewRequest(http.MethodGet, CookieAPIURL, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}
	req.SetBasicAuth("guangyixin", "*~je,R#(anqAD")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求凭证接口失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %w", err)
	}

	var cookieResp CookieResponse
	if err := json.Unmarshal(body, &cookieResp); err != nil {
		return fmt.Errorf("解析响应失败: %w, 响应: %s", err, string(body))
	}

	if cookieResp.Code != 200 {
		return fmt.Errorf("获取凭证失败: code=%d, message=%s", cookieResp.Code, cookieResp.Message)
	}
	if cookieResp.Cookie == "" || cookieResp.Csrf == "" {
		return fmt.Errorf("返回的 Cookie 或 Csrf 为空")
	}

	m.mu.Lock()
	m.cookie = cookieResp.Cookie
	m.csrf = cookieResp.Csrf
	m.lastUpdate = time.Now()
	m.mu.Unlock()

	log.Printf("[cookie] 凭证获取成功，cookie 长度: %d", len(cookieResp.Cookie))
	return nil
}

// autoRefresh 定时刷新凭证
func (m *Manager) autoRefresh() {
	ticker := time.NewTicker(RefreshInterval)
	defer ticker.Stop()

	retryWait := InitialRetryWait

	for {
		select {
		case <-ticker.C:
			if err := m.fetchCookie(); err != nil {
				log.Printf("[cookie] 刷新凭证失败: %v，将在 %v 后重试", err, retryWait)
				time.AfterFunc(retryWait, func() {
					if err := m.fetchCookie(); err != nil {
						log.Printf("[cookie] 重试获取凭证失败: %v", err)
					}
				})
				retryWait *= 2
				if retryWait > 5*time.Minute {
					retryWait = 5 * time.Minute
				}
			} else {
				retryWait = InitialRetryWait
			}
		case <-m.stopCh:
			log.Println("[cookie] 管理器已停止")
			return
		}
	}
}

// Stop 停止自动刷新
func (m *Manager) Stop() {
	close(m.stopCh)
}

// GetLastUpdateTime 获取上次更新时间
func (m *Manager) GetLastUpdateTime() time.Time {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.lastUpdate
}
