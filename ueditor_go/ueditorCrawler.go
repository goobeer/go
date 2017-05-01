package ueditor_go

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/valyala/fasthttp"
)

type UeditorCrawlerRespData struct {
	State string            `json:"state"`
	List  []*UeditorCrawler `json:"list"`
}

type UeditorCrawler struct {
	SourceUrl string
	ServerUrl string
	BasePath  string
	State     string
	*UEditorHandler
}

func (u *UeditorCrawler) Process() {
	sources := u.Ctx.FormValue("source[]")

	TD := &UeditorCrawlerRespData{}
	if sources == nil || len(sources) == 0 {
		TD.State = "参数错误：没有指定抓取源"
		u.WriteJson(TD)
		return
	}
	//	var Crawlers []*UeditorCrawler
	//	for _, v := range sources {
	//		Crawlers = append(Crawlers, &UeditorCrawler{SourceUrl: v})
	//	}

	//        Crawlers = Sources.Select(x => new Crawler(x, Server).Fetch()).ToArray();
	//        WriteJson(new
	//        {
	//            state = "SUCCESS",
	//            list = Crawlers.Select(x => new
	//            {
	//                state = x.State,
	//                source = x.SourceUrl,
	//                url = x.ServerUrl
	//            })
	//        });
}

func (u *UeditorCrawler) Fetch() {
	if !u.isExternalIPAddress(u.SourceUrl) {
		u.State = "INVALID_URL"
	}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(u.SourceUrl)
	resp := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, resp)

	if err == nil {
		if resp.StatusCode() != fasthttp.StatusOK {
			u.State = fmt.Sprintf("Url returns %d", resp.StatusCode())
			return
		}
		if !strings.Contains(string(resp.Header.ContentType()), "image") {
			u.State = "Url is not an image"
		}

		uConfig, _ := BuildItems()
		u.ServerUrl = UeditorPathFormatter(path.Base(u.SourceUrl), uConfig.CatcherPathFormat)
		savePath := u.ServerUrl
		_, err = os.Stat(path.Base(savePath))
		if err != nil && os.IsNotExist(err) {
			os.Mkdir(path.Base(savePath), 0644)
		}

		fi, err := os.Create(savePath)
		if err != nil {
			u.State = "抓取错误：" + err.Error()
			panic(err)
		}
		fi.Write(resp.Body())
		u.State = "SUCCESS"
	}
}

func (u *UeditorCrawler) isExternalIPAddress(urlStr string) (ok bool) {
	uri, err := url.Parse(urlStr)
	if err != nil {
		return
	}
	ip := net.ParseIP(uri.Host)
	if ip != nil {
		ipAddr, err := net.ResolveIPAddr("ip4", ip.String())
		if err != nil {
			return
		}
		ok = !u.isPrivateIP(ipAddr)
	} else {
		addrs, err := net.LookupHost(uri.Host)
		if err != nil {
			return
		}
		for _, v := range addrs {
			rip := net.ParseIP(v)
			if rip != nil && rip.To4() != nil {
				ipAddr, err := net.ResolveIPAddr("ip4", v)
				if err == nil && !u.isPrivateIP(ipAddr) {
					return
				}
			}
		}
	}

	return
}

func (u *UeditorCrawler) isPrivateIP(myIPAddress *net.IPAddr) (ok bool) {
	ok = myIPAddress.IP.IsLoopback()
	if ok {
		return
	}
	ip4 := myIPAddress.IP.To4()
	if ip4 != nil {
		ip4Str := ip4.String()
		ok = strings.HasPrefix(ip4Str, "10.") || strings.HasPrefix(ip4Str, "172.16.") || strings.HasPrefix(ip4Str, "192.168.") || strings.HasPrefix(ip4Str, "169.254.")
	}

	return
}
