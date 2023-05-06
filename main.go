package main

import (
	"fmt"
	"net/url"
	"strings"
)

func generateDork(keyword string) []string {
	dorks := []string{
		fmt.Sprintf("site:pastebin.com intext:\"%s\"", keyword),
		fmt.Sprintf("site:\"%s\" inurl:redir | inurl:url | inurl:redirect | inurl:return | inurl:src=http | inurl:r=http", keyword),
		fmt.Sprintf("site:\"%s\" 'password' filetype:doc | filetype:pdf | filetype:docx | filetype:xls | filetype:dat | filetype:log", keyword),
		fmt.Sprintf("site:\"%s\" intitle:index.of | 'parent directory'", keyword),
		fmt.Sprintf("site:\"%s\" inurl:login | inurl:signin | intitle:Login | intitle: signin | inurl:auth", keyword),
		fmt.Sprintf("site:\"%s\" inurl:admin | intitle:admin", keyword),
		fmt.Sprintf("site:\"%s\" inurl:wp- | inurl:wp-content | inurl:plugins | inurl:uploads | inurl:themes | inurl:download", keyword),
		fmt.Sprintf("site:\"%s\" intext:'sql syntax near' | intext:'syntax error has occurred' | intext:'incorrect syntax near' | intext:'unexpected end of SQL command' | intext:'Warning: mysql_connect()' | intext:'Warning: mysql_query() | intext:'Warning: pg_connect()' | filetype:sqlext:sql | ext:dbf | ext:mdb", keyword),
		fmt.Sprintf("site:\"%s\" ext:xml | ext:conf | ext:cnf | ext:reg | ext:inf | ext:rdp | ext:cfg | ext:txt | ext:ora | ext:ini | ext:log", keyword),
		fmt.Sprintf("site:\"%s\" ext:bkf | ext:bkp | ext:bak | ext:old | ext:backup", keyword),
		fmt.Sprintf("site:\"%s\" ext:php intitle:phpinfo 'published by the PHP Group'", keyword),
		fmt.Sprintf("site:*.\"%s\"", keyword),
		fmt.Sprintf("site:stackoverflow.com '\"%s\"'", keyword),
		fmt.Sprintf("site:github.com '\"%s\"'", keyword),
	}
	return dorks
}

func main() {
	var keyword string
	fmt.Print("Please enter a domain: ")
	fmt.Scanln(&keyword)
	dorks := generateDork(keyword)
	for _, dork := range dorks {
		encoded := url.QueryEscape(dork)
		url := fmt.Sprintf("https://www.google.com/search?q=%s", encoded)
		fmt.Println(strings.Replace(url, "+", "%20", -1))
	}
}
