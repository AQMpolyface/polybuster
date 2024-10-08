# polybuster
like gobuster, but much less sofisticated and barebones

for some reason, if you wanted to use this instead of gobuster, you can

```
git clone https://github.com/AQMpolyface/polybuster.git

go run polybuster.go -u <url> -w </path/to/wordlist>
```

 if no wordlist is specified, the program will attempt to download the medium seclist from github.
