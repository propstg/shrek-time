# shrek-time

shrek time microservice

1 shrek in Shrek Standard Time (SST) is equal to 93 minutes, the runtime of the movie *Shrek*.

Shrek time epoch (Shrepoch), 0 SST, is the date and time the film was releasted: 22 April 2001 00:00:00.000 PDT. 100,000 SST will occur at 27 December 2018 15:00:00 UTC.

## Endpoints
GET /api/now: get the current time in shrek standard time.
Url: /api/now


GET /api/toShrek/{utc}: convert UTC timestamp from RFC3339 format to SST.
Example url: /api/toShrek/2018-12-27T15:00:00Z


GET /api/fromShrek/{sst}: convert SST to UTC timestamp in RFC3339 format.
Example url: /api/fromShrek/100000
