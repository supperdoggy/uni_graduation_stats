from tokenize import String
from typing import List
from selenium import webdriver
from bs4 import BeautifulSoup
import time
import re

from user import WorkExperience

def login(driver: webdriver.Chrome):
    email = "supperspggy229@icloud.com"
    password = "bz7mu4sr"
    # Creating a webdriver instance

    
    # Opening linkedIn's login page

    driver.get("https://linkedin.com/uas/login")
    
    # waiting for the page to load

    time.sleep(5)
    
    # entering username

    username = driver.find_element_by_id("username")
    
    # In case of an error, try changing the element
    # tag used here.
    
    # Enter Your Email Address

    username.send_keys(email)  
    
    # entering password

    pword = driver.find_element_by_id("password")
    # In case of an error, try changing the element 
    # tag used here.
    
    # Enter Your Password

    pword.send_keys(password)        
    
    # Clicking on the log in button
    # Format (syntax) of writing XPath --> 
    # //tagname[@attribute='value']

    driver.find_element_by_xpath("//button[@type='submit']").click()
    # In case of an error, try changing the
    # XPath used here
    return

def scrape_profile(driver: webdriver.Chrome, url: str):
    open_page(driver, url)

    # Extracting the HTML of the complete introduction box
    # that contains the name, company name, and the location

    src = driver.page_source
    soup = BeautifulSoup(src, 'lxml')

    intro = soup.find('div', {'class': 'pv-text-details__left-panel'})
    name_loc = intro.find("h1")

    name = name_loc.get_text().strip()
    works_at_loc = intro.find("div", {'class': 'text-body-medium'})
    works_at = works_at_loc.get_text().strip()
   
   
    print("Name -->", name,
        "\nWorks At -->", works_at)

    scrape_experience(driver, f"{url}/details/experience")

    return

def scrape_experience(driver: webdriver.Chrome, url: str) -> List[WorkExperience]:
    experiences = []
    open_page(driver, url)
    src = driver.page_source
    soup = BeautifulSoup(src, 'lxml')

    # get job duration
    works = soup.find_all("div", class_="display-flex flex-column full-width")
    for work in works:
        title_src = work.find_all("span", class_="mr1 t-bold")
        title = title_src[0].text[0:len(title_src[0].text)//2]

        name_src = work.find_all("span", class_="t-14 t-normal")
        name = name_src[0].text[0:name_src[0].text.find("·")-1].replace("\n", "")
        
        period_src = work.find_all("span", class_="t-14 t-normal t-black--light")
        period = period_src[0].text[0:len(period_src[0].text)//2]
        experiences.append(WorkExperience(title, period, name))

    # mr1 hoverable-link-text t-bold
    
    print(experiences)
    for exp in experiences:
        print(exp)

    return

def open_page(driver: webdriver.Chrome, url: str):
    driver.get(url)

    start = time.time()
 
    initialScroll = 0

    finalScroll = 1000

    while True:

        driver.execute_script(f"window.scrollTo({initialScroll},{finalScroll})")
        initialScroll = finalScroll
        finalScroll += 1000
        time.sleep(3)
        end = time.time()

        if round(end - start) > 20:
            break
     
