import json
from urllib.request import Request, urlopen


ENDPOINT = "https://api.wheretheiss.at/v1/satellites/25544"


def find_iss():
    req = Request(ENDPOINT)
    with urlopen(req) as res:
        data = json.load(res)
    
    lat = data["latitude"]
    lon = data["longitude"]
    alt = data["altitude"]

    


find_iss()
