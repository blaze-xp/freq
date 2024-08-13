![Screenshot (242)](https://github.com/user-attachments/assets/d67535ec-8cd5-42db-89b2-f7b476b5ebc2)
how to use:

go install github.com/coffinsp/freq@latest

cd go/bin

sudo mv freq /usr/local/bin

echo testphp.vulnweb.com | waybackurls | gf xss | urldedupe | qsreplace 'yourxsspayloadhere' | freq
