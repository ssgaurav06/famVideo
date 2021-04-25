**Problem Statement**

To make an API to fetch latest videos sorted in reverse chronological 
order of their publishing date-time from YouTube for a given tag/search 
query in a paginated response.

The app should perform the following tasks:
* Server should call the YouTube API continuously in background (async) with some interval (say 10 seconds) for fetching the latest videos for a predefined search query and should store the data of videos (specifically these fields - Video title, description, publishing datetime, thumbnails URLs and any other fields you require) in a database with proper indexes.
* A GET API which returns the stored video data in a paginated response sorted in descending order of published datetime.
* A basic search API to search the stored videos using their title and description.

***Bonus Features***

* Add support for supplying multiple API keys so that if quota is exhausted on one, it automatically uses the next available key.
* Optimise search api, so that it's able to search videos containing partial match for the search query in either video title or description.
  But, while doing any query use `&` for spaces.

**Dependencies**

* Go go1.15 darwin/amd64
* PostgreSQL 13.2

**Build Instructions**

Clone project using following command,
> git clone https://github.com/ssgaurav06/famVideo.git

Move to the project directory using,
> cd famVideo

Set Postgres details and Developer_Key (for accessing Youtube API)
in `.env` file.

In order to run the main app, run the following command,
> go run main.go

***Using Docker***

In order to run the main app, run the following command,
> docker compose up

In order to stop the app, run the following command,
> docker compose down

**Input commands**

To get the stored video data:
> curl -v -X GET 'http://localhost:8080/videoData'

To search for the videos with particular title or description in stored videos:
> curl -v -X GET 'http://localhost:8080/search?query=something'