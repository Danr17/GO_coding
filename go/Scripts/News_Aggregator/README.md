The Flow:  
-- Retrieve SitemapIndex from Washingtonpost news site (xml sitemap)  
-- For each News Title run concurrency (go routines with channels) to retrieve data (Titles, Keywords, Locations)  
-- Run a webserver, parse html file and display using dataTables js  

https://pythonprogramming.net/go/concurrency-web-app-go-language-programming-tutorial/
