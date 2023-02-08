# 網路爬蟲
 - 以下範例是爬 [yahoo 的電影資訊](https://movies.yahoo.com.tw/movie_intheaters.html)

## crawler_original
 - 基本爬蟲架構

## rawler_concurrency_simple
 - 基本高併發爬蟲 (簡易修改crawler_original)，但可能會因為請求速度太快，會被某些網站擋下來，擋掉後就會死掉，需要而外給請求設限條件

## crawler_concurrency_queue
 - 高併發，利用通道阻塞特性來避免因請求過快被擋掉而死亡的方式

## crawler_concurrency_integral
 - 將 crawler_concurrency_queue 的版本利用抽象層改得更靈活
