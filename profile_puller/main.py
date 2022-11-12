import json
import pymongo
from urllib.request import urlopen


def pull_profile(url: str):
    with urlopen(f"http://localhost:3000/?url={url}") as f:
        resp = json.load(f)
    return resp

if __name__ == "__main__":
    myclient = pymongo.MongoClient("mongodb://localhost:27017/")
    mydb = myclient["stud"]
    users_search = mydb["users_search"]

    link = "https://www.linkedin.com/in/oleksii-sardachuk"
    profile = pull_profile(link)
    profile["_id"] = link.split("/")[-1]
    print(profile["_id"])

