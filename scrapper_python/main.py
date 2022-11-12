import json
from selenium import webdriver

import scrapper
import time
  
# # Opening JSON file
# f = open('result.json')
  
# # returns JSON object as 
# # a dictionary
# data = json.load(f)
driver = webdriver.Chrome()


scrapper.login(driver)

scrapper.scrape_profile(driver, "https://www.linkedin.com/in/supperdoggy")

time.sleep(500000)