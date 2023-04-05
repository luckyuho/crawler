# 網路爬蟲
 - 以下範例是爬 [yahoo 的電影資訊](https://movies.yahoo.com.tw/movie_intheaters.html)

## 功能
根據[yahoo 的電影資訊](https://movies.yahoo.com.tw/movie_intheaters.html)中的資訊，將裡頭的第一頁電影列表及每部電影中的詳細資料抓取出來，結果如下：

 - 第一層資訊
<div align="center">
<table>
  <tr>
    <td>電影列表</td>
    <td>爬蟲結果(1)</td>
  </tr>
  <tr>
    <td><img src="https://github.com/luckyuho/crawler/blob/main/img/movieList.png" width=400 height=1000 title="instructor profile" /></td>
    <td><img src="https://github.com/luckyuho/crawler/blob/main/img/crawlerMovieList.png" width=300 height=600 title="student profile" /></td>
  </tr>
</table>
</div>

 - 第二層資訊
<div align="center">
<table>
  <tr>
    <td>電影列表中每個電影個詳細資訊(這裡僅顯示其中一個資料為範例)</td>
    <td>爬蟲結果(2)</td>
  </tr>
  <tr>
    <td><img src="https://github.com/luckyuho/crawler/blob/main/img/movieDetail.png" width=600 height=340 title="instructor profile" /></td>
    <td><img src="https://github.com/luckyuho/crawler/blob/main/img/crawlerMovieDetail.png" width=270 height=540 title="student profile" /></td>
  </tr>
</table>
</div>

## crawler_original
 - 基本爬蟲架構

## rawler_concurrency_simple
 - 基本高併發爬蟲 (簡易修改crawler_original)，但可能會因為請求速度太快，會被某些網站擋下來，擋掉後就會死掉，需要而外給請求設限條件

## crawler_concurrency_queue
 - 高併發，利用通道阻塞特性來避免因請求過快被擋掉而死亡的方式

## crawler_concurrency_integral
 - 將 crawler_concurrency_queue 的版本利用抽象層改得更靈活

## 簡單的投影片介紹
 - [投影片介紹crawler_original與rawler_concurrency_simple的連結](https://docs.google.com/presentation/d/13e4A06VpmRPncbGGWVs32RFA1GtwsRFzH_CuXxcki1Q/edit#slide=id.p)
