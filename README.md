shortens urls

GET host/{id} -> redirect to website OR 404 and json "not found" | 500 error and json with error 
PUT host/new {"url":"my.website.com"} -> 200 ok OR 500 json with error 