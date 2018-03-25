## Extract (sub)domains from HTTPS web sites

Extracts subdomain names from https sites. The result is returned as a json.

### Usage
* go run main.go `google.com`
* ./main `google.com`
* Or build it yourself

### It will
* add the protocol for you, e.g.: sets "https://" if you pass only "google.com"
* correct the protocol, if you pass "http://google.com" it will change it to "http*s*://google.com"
* make sure it's a valid domain name. I really hope so.

### Example result
```json
["*.google.com","*.android.com","*.appengine.google.com","*.cloud.google.com","*.db833953.google.cn","*.g.co","*.gcp.gvt2.com","*.google-analytics.com","*.google.ca","*.google.cl","*.google.co.in","*.google.co.jp","*.google.co.uk","*.google.com.ar","*.google.com.au","*.google.com.br","*.google.com.co","*.google.com.mx","*.google.com.tr","*.google.com.vn","*.google.de","*.google.es","*.google.fr","*.google.hu","*.google.it","*.google.nl","*.google.pl","*.google.pt","*.googleadapis.com","*.googleapis.cn","*.googlecommerce.com","*.googlevideo.com","*.gstatic.cn","*.gstatic.com","*.gvt1.com","*.gvt2.com","*.metric.gstatic.com","*.urchin.com","*.url.google.com","*.youtube-nocookie.com","*.youtube.com","*.youtubeeducation.com","*.yt.be","*.ytimg.com","android.clients.google.com","android.com","developer.android.google.cn","developers.android.google.cn","g.co","goo.gl","google-analytics.com","google.com","googlecommerce.com","source.android.google.cn","urchin.com","www.goo.gl","youtu.be","youtube.com","youtubeeducation.com","yt.be"]
```
---
> No unit tests, we `panic` like real men. 
> Baked for internal usage, for something more complicated, check [https://franccesco.github.io/getaltname/](https://franccesco.github.io/getaltname/)
