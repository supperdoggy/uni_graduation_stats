import json
import pymongo
from urllib.request import urlopen
import random
import string


def pull_profile(url: str):
    with urlopen(f"http://localhost:3000/?url={url}") as f:
        resp = json.load(f)
    return resp


def get_next_search(coll):
    doc = coll.find_one()
    print(doc)

    coll.delete_one({"_id": doc["_id"]})
    return doc

if __name__ == "__main__":
    myclient = pymongo.MongoClient("mongodb://localhost:27017/")
    mydb = myclient["stud"]
    users_search = mydb["users_search"]
    users = mydb["users"]

    for k in range(40):
        print(k)
        doc = get_next_search(users_search)
        link = doc["profileUrl"]
        profile = pull_profile(link)
        profile["_id"] = link.split("/")[-1]
        try:
            users.insert_one(profile)
        except:
            profile["_id"] = ''.join(random.choices(string.ascii_uppercase + string.digits, k = 20))
            users.insert_one(profile)
        print(profile)
