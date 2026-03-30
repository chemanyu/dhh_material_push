package main

import (
	"bytes"
	"crypto/rand"
	"dhh-material-tool/cookie"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	_ "embed"
)

//go:embed static/index.html
var indexHTML []byte

const dhhBaseURL = "https://dhh.taobao.com"

var cookieMgr *cookie.Manager

// MetaItem 上传素材后得到的元数据
type MetaItem struct {
	FileName     string `json:"fileName"`
	MaterialURL  string `json:"materialUrl"`
	MaterialCode string `json:"materialCode"`
}

// MaterialCreateRequest 提交审核请求体
type MaterialCreateRequest struct {
	MetaList         []MetaItem `json:"metaList"`
	AppID            string     `json:"appId"`
	TaskType         string     `json:"taskType"`
	AdType           string     `json:"adType"`
	ScenarioType     string     `json:"scenarioType"`
	HotEvent         string     `json:"hotEvent"`
	BaseImageType    string     `json:"baseImageType"`
	CustomTitle      string     `json:"customTitle"`
	CustomCopy       string     `json:"customCopy"`
	BizType          string     `json:"bizType"`
	ScenarioTypeDesc string     `json:"scenarioTypeDesc"`
	BizTypeDesc      string     `json:"bizTypeDesc"`
}

// FailItem 提交失败的素材
type FailItem struct {
	FileName     string `json:"fileName"`
	MaterialCode string `json:"materialCode"`
	MaterialURL  string `json:"materialUrl"`
	ErrorMessage string `json:"errorMessage"`
}

// ---------- HTTP 工具 ----------

func buildDhhHeaders(csrf, cookie string) map[string]string {
	return map[string]string{
		"accept":          "*/*",
		"accept-language": "zh-CN,zh;q=0.9",
		"x-xsrf-token":    csrf,
		"Cookie":          cookie,
	}
}

var httpClient = &http.Client{Timeout: 60 * time.Second}

func logCurlGet(reqURL string, headers map[string]string) {
	var sb strings.Builder
	sb.WriteString("curl -X GET \\\n")
	for k, v := range headers {
		sb.WriteString(fmt.Sprintf("  -H '%s: %s' \\\n", k, v))
	}
	sb.WriteString(fmt.Sprintf("  '%s'", reqURL))
	//log.Printf("[CURL]\n%s", sb.String())
}

// logCurlPost 打印等效 curl 命令。
// urlencode=true  → --data-urlencode 'k=明文v'（对应 material/create）
// urlencode=false → --data-raw '已编码body'（对应 meta/add 等）
func logCurlPost(reqURL string, headers map[string]string, params url.Values, urlencode bool) {
	var sb strings.Builder
	sb.WriteString("curl \\\n")
	for k, v := range headers {
		sb.WriteString(fmt.Sprintf("  -H '%s: %s' \\\n", k, v))
	}
	sb.WriteString("  -H 'content-type: application/x-www-form-urlencoded' \\\n")
	if urlencode {
		for k, vs := range params {
			for _, v := range vs {
				sb.WriteString(fmt.Sprintf("  --data-urlencode '%s=%s' \\\n", k, v))
			}
		}
	} else {
		sb.WriteString(fmt.Sprintf("  --data-raw '%s' \\\n", encodeParams(params)))
	}
	sb.WriteString(fmt.Sprintf("  '%s'", reqURL))
	//log.Printf("[CURL]\n%s", sb.String())
}

func dhhGet(path, csrf, cookie string) (map[string]interface{}, error) {
	reqURL := dhhBaseURL + path + "?_csrf=" + url.QueryEscape(csrf)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	headers := buildDhhHeaders(csrf, cookie)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	logCurlGet(reqURL, headers)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("[GET] %s -> %d", path, resp.StatusCode)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %s", string(body))
	}
	return result, nil
}

// encodeParams 模拟 curl --data-urlencode：使用 RFC 3986 percent-encoding（%20 而非 +）
func encodeParams(params url.Values) string {
	var parts []string
	for k, vs := range params {
		for _, v := range vs {
			parts = append(parts, url.QueryEscape(k)+"="+url.QueryEscape(v))
		}
	}
	// url.QueryEscape 仍用 + 代替空格，需替换为 %20
	result := strings.Join(parts, "&")
	return strings.ReplaceAll(result, "+", "%20")
}

func dhhPost(path, csrf, cookie string, params url.Values, encode bool) (map[string]interface{}, error) {
	reqURL := dhhBaseURL + path
	params.Set("_csrf", csrf)

	headers := buildDhhHeaders(csrf, cookie)
	logCurlPost(reqURL, headers, params, encode)

	encoded := encodeParams(params)
	req, err := http.NewRequest("POST", reqURL, strings.NewReader(encoded))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//log.Printf("[POST] %s -> %d | body: %s", path, resp.StatusCode, string(body))
	log.Printf("[POST] %s -> %d", path, resp.StatusCode)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %s", string(body))
	}
	return result, nil
}

// ---------- 随机串 ----------

func generateOssRandom() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func generateFileId() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 28)
	rand.Read(b)
	result := make([]byte, 28)
	for i, byt := range b {
		result[i] = chars[int(byt)%len(chars)]
	}
	return "o_" + string(result)
}

// ---------- 响应工具 ----------

func jsonOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"code": 0, "data": data})
}

func jsonErr(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"code": 1, "message": msg})
}

// ---------- 处理器 ----------

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(indexHTML)
}

func handleFormInfo(w http.ResponseWriter, r *http.Request) {
	ck, csrf, err := cookieMgr.GetCredentials()
	if err != nil {
		jsonErr(w, "凭证未就绪: "+err.Error())
		return
	}
	result, err := dhhGet("/polystar/api/creative/material/forminfo", csrf, ck)
	if err != nil {
		log.Printf("[forminfo] error: %v", err)
		jsonErr(w, err.Error())
		return
	}
	jsonOK(w, result)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	ck, csrf, err := cookieMgr.GetCredentials()
	if err != nil {
		jsonErr(w, "凭证未就绪: "+err.Error())
		return
	}

	if err := r.ParseMultipartForm(500 << 20); err != nil {
		jsonErr(w, "解析上传文件失败: "+err.Error())
		return
	}

	allowedExts := map[string]bool{
		"mp4": true, "mov": true, "avi": true,
		"jpg": true, "jpeg": true, "png": true, "gif": true, "webp": true,
	}

	fhs := r.MultipartForm.File["files[]"]
	if len(fhs) == 0 {
		fhs = r.MultipartForm.File["files"]
	}
	if len(fhs) == 0 {
		jsonErr(w, "未收到文件")
		return
	}
	if len(fhs) > 10 {
		jsonErr(w, fmt.Sprintf("一次最多上传 10 个文件，当前: %d 个", len(fhs)))
		return
	}

	var metaList []MetaItem

	for _, fh := range fhs {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(fh.Filename), "."))
		if !allowedExts[ext] {
			log.Printf("[upload] skip unsupported: %s", fh.Filename)
			continue
		}

		f, err := fh.Open()
		if err != nil {
			log.Printf("[upload] open error: %s %v", fh.Filename, err)
			continue
		}
		fileBytes, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			log.Printf("[upload] read error: %s %v", fh.Filename, err)
			continue
		}

		// Step 1: osssign
		osssignResp, err := dhhGet("/polystar/api/creative/material/osssign", csrf, ck)
		if err != nil {
			log.Printf("[upload] osssign error: %s %v", fh.Filename, err)
			continue
		}
		ossData, ok := osssignResp["data"].(map[string]interface{})
		if !ok || len(ossData) == 0 {
			log.Printf("[upload] osssign empty data for: %s", fh.Filename)
			continue
		}

		// Step 2: upload to OSS
		ossKey := "dsp/video/" + generateOssRandom() + "." + ext
		if err := uploadToOss(ossData, ossKey, fh.Filename, fileBytes); err != nil {
			log.Printf("[upload] oss upload error: %s %v", fh.Filename, err)
			continue
		}

		// Step 3: meta/add
		ossHost := strings.TrimRight(fmt.Sprintf("%v", ossData["host"]), "/")
		ossURL := ossHost + "/" + ossKey

		params := url.Values{}
		params.Set("ossUrl", ossURL)
		params.Set("fileName", fh.Filename)

		metaResp, err := dhhPost("/polystar/api/creative/material/meta/add", csrf, ck, params, false)
		if err != nil {
			log.Printf("[upload] meta/add error: %s %v", fh.Filename, err)
			continue
		}

		if data, ok := metaResp["data"].(map[string]interface{}); ok && data != nil {
			item := MetaItem{
				FileName:     strVal(data["fileName"], fh.Filename),
				MaterialURL:  strVal(data["materialUrl"], ""),
				MaterialCode: strVal(data["materialCode"], ""),
			}
			if data["duplicate"] == true {
				item.MaterialCode = "重复文件"
			}
			metaList = append(metaList, item)
			//log.Printf("[upload] success: %s -> %s", item.FileName, item.MaterialCode)
		} else {
			log.Printf("[upload] meta/add no data for: %s resp: %v", fh.Filename, metaResp)
		}
	}

	if metaList == nil {
		metaList = []MetaItem{}
	}
	log.Printf("[upload] done, success: %d / %d", len(metaList), len(fhs))
	jsonOK(w, map[string]interface{}{"metaList": metaList})
}

func uploadToOss(ossData map[string]interface{}, ossKey, fileName string, fileBytes []byte) error {
	host := strings.TrimRight(fmt.Sprintf("%v", ossData["host"]), "/") + "/"

	var body bytes.Buffer
	mw := multipart.NewWriter(&body)

	orderedFields := []struct{ k, v string }{
		{"name", fileName},
		{"key", ossKey},
		{"dir", "dsp"},
		{"policy", fmt.Sprintf("%v", ossData["policy"])},
		{"OSSAccessKeyId", fmt.Sprintf("%v", ossData["accessid"])},
		{"success_action_status", "200"},
		{"signature", fmt.Sprintf("%v", ossData["signature"])},
	}
	for _, f := range orderedFields {
		if err := mw.WriteField(f.k, f.v); err != nil {
			return err
		}
	}

	mimeType := mime.TypeByExtension(filepath.Ext(fileName))
	if mimeType == "" {
		mimeType = http.DetectContentType(fileBytes)
	}

	fw, err := mw.CreatePart(map[string][]string{
		"Content-Disposition": {fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName)},
		"Content-Type":        {mimeType},
	})
	if err != nil {
		return err
	}
	if _, err := fw.Write(fileBytes); err != nil {
		return err
	}
	mw.Close()

	ossClient := &http.Client{Timeout: 5 * time.Minute}
	req, err := http.NewRequest("POST", host, &body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Origin", "https://dhh.taobao.com")
	req.Header.Set("Referer", "https://dhh.taobao.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36")

	resp, err := ossClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body)

	log.Printf("[oss] %s -> HTTP %d", fileName, resp.StatusCode)
	if resp.StatusCode != 200 {
		return fmt.Errorf("OSS 上传失败，HTTP %d", resp.StatusCode)
	}
	return nil
}

func handleMaterialCreate(w http.ResponseWriter, r *http.Request) {
	ck, csrf, err := cookieMgr.GetCredentials()
	if err != nil {
		jsonErr(w, "凭证未就绪: "+err.Error())
		return
	}

	var req MaterialCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonErr(w, "请求参数错误: "+err.Error())
		return
	}
	if len(req.MetaList) == 0 {
		jsonErr(w, "metaList 不能为空")
		return
	}

	successCount := 0
	failCount := 0
	var successList []MetaItem
	var failList []FailItem

	for i := 0; i < len(req.MetaList); i += 10 {
		end := i + 10
		if end > len(req.MetaList) {
			end = len(req.MetaList)
		}
		batch := req.MetaList[i:end]

		type matListItem struct {
			MaterialCode string `json:"materialCode"`
			FileID       string `json:"fileId"`
			MaterialURL  string `json:"materialUrl"`
		}
		matList := make([]matListItem, len(batch))
		for j, item := range batch {
			matList[j] = matListItem{
				MaterialCode: item.MaterialCode,
				FileID:       generateFileId(),
				MaterialURL:  item.MaterialURL,
			}
		}
		matListJSON, _ := json.Marshal(matList)

		params := url.Values{}
		params.Set("admissionType", "material")
		params.Set("appId", req.AppID)
		params.Set("taskType", req.TaskType)
		params.Set("adType", req.AdType)
		params.Set("scenarioType", req.ScenarioType)
		params.Set("hotEvent", req.HotEvent)
		params.Set("baseImageType", req.BaseImageType)
		params.Set("customTitle", req.CustomTitle)
		params.Set("customCopy", req.CustomCopy)
		params.Set("bizType", req.BizType)
		params.Set("scenarioTypeDesc", req.ScenarioTypeDesc)
		params.Set("bizTypeDesc", req.BizTypeDesc)
		params.Set("materialList", string(matListJSON))

		result, err := dhhPost("/polystar/api/creative/material/create", csrf, ck, params, true)
		if err != nil {
			log.Printf("[create] batch %d error: %v", i/10+1, err)
			failCount += len(batch)
			for _, item := range batch {
				failList = append(failList, FailItem{
					FileName: item.FileName, MaterialCode: item.MaterialCode,
					MaterialURL: item.MaterialURL, ErrorMessage: err.Error(),
				})
			}
			continue
		}

		successful, _ := result["successful"].(bool)
		if successful {
			successCount += len(batch)
			successList = append(successList, batch...)
		} else {
			msg := strVal(result["message"], "提交失败")
			failCount += len(batch)
			for _, item := range batch {
				failList = append(failList, FailItem{
					FileName: item.FileName, MaterialCode: item.MaterialCode,
					MaterialURL: item.MaterialURL, ErrorMessage: msg,
				})
			}
		}
	}

	if successList == nil {
		successList = []MetaItem{}
	}
	if failList == nil {
		failList = []FailItem{}
	}
	jsonOK(w, map[string]interface{}{
		"successCount": successCount,
		"failCount":    failCount,
		"successList":  successList,
		"failList":     failList,
	})
}

// ---------- 工具 ----------

func strVal(v interface{}, fallback string) string {
	if v == nil {
		return fallback
	}
	s := fmt.Sprintf("%v", v)
	if s == "<nil>" || s == "" {
		return fallback
	}
	return s
}

func openBrowser(addr string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", addr)
	case "darwin":
		cmd = exec.Command("open", addr)
	default:
		cmd = exec.Command("xdg-open", addr)
	}
	if err := cmd.Start(); err != nil {
		log.Printf("自动打开浏览器失败，请手动访问: %s", addr)
	}
}

func main() {
	cookieMgr = cookie.NewManager()
	defer cookieMgr.Stop()

	port := "18080"
	addr := "http://localhost:" + port

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/forminfo", handleFormInfo)
	http.HandleFunc("/api/upload", handleUpload)
	http.HandleFunc("/api/material-create", handleMaterialCreate)

	log.Printf("大航海素材工具已启动: %s", addr)

	go func() {
		time.Sleep(600 * time.Millisecond)
		openBrowser(addr)
	}()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
