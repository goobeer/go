package ueditor_go

import (
	"net"
	"net/url"
	"strings"

	"github.com/valyala/fasthttp"
)

type UeditorCrawler struct {
	SourceUrl string
	ServerUrl string
	State     string
}

func (u *UeditorCrawler) Fetch() {
	if !u.isExternalIPAddress(u.SourceUrl) {
		u.State = "INVALID_URL"
	}
	statusCode, body, err := fasthttp.Get(nil, u.SourceUrl)
	if err == nil {
		if statusCode != fasthttp.StatusOK {
			u.State = "Url returns " + statusCode
			return
		}
		//TODO xx
	}
	//        var request = HttpWebRequest.Create(this.SourceUrl) as HttpWebRequest;
	//        using (var response = request.GetResponse() as HttpWebResponse)
	//        {
	//            if (response.ContentType.IndexOf("image") == -1)
	//            {
	//                State = "Url is not an image";
	//                return this;
	//            }
	//            ServerUrl = PathFormatter.Format(Path.GetFileName(this.SourceUrl), Config.GetString("catcherPathFormat"));
	//            var savePath = Server.MapPath(ServerUrl);
	//            if (!Directory.Exists(Path.GetDirectoryName(savePath)))
	//            {
	//                Directory.CreateDirectory(Path.GetDirectoryName(savePath));
	//            }
	//            try
	//            {
	//                var stream = response.GetResponseStream();
	//                var reader = new BinaryReader(stream);
	//                byte[] bytes;
	//                using (var ms = new MemoryStream())
	//                {
	//                    byte[] buffer = new byte[4096];
	//                    int count;
	//                    while ((count = reader.Read(buffer, 0, buffer.Length)) != 0)
	//                    {
	//                        ms.Write(buffer, 0, count);
	//                    }
	//                    bytes = ms.ToArray();
	//                }
	//                File.WriteAllBytes(savePath, bytes);
	//                State = "SUCCESS";
	//            }
	//            catch (Exception e)
	//            {
	//                State = "抓取错误：" + e.Message;
	//            }
	//            return this;
	//        }
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
