敘述
==============
Web App網址在[這裏](http://evan-go-apis.herokuapp.com/) 

# 安裝流程 

依照以下兩篇應該可以順利成功

- http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html
- http://yinghau76.github.io/2013/12/15/go-on-heroku/

不過我真的太不熟所有有些地方一直搞混，這裡記錄一下我發生的問題：

千萬記得把你的Go的程式碼放在$GOPATH/src 下，根據這一篇建議，可以考慮放在 $GOPATCH/src/github.com/YOUR_GITHUB_ACCOUNT/ 下面
記得 Procfile 裡面要寫的是你的目錄(也就是編譯好執行檔)名稱  ，不確定的可以用 which 去double confirm



