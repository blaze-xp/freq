how to use:

go install github.com/coffinsp/freq@latest

cd go/bin

sudo mv freq /usr/local/bin

echo testphp.vulnweb.com | waybackurls | gf xss | urldedupe | qsreplace 'yourxsspayloadhere' | freq
![Uploading Screenshot (242).png…]()
