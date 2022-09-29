#Author:Leon
from selenium import webdriver
from selenium.common.exceptions import WebDriverException
from selenium.webdriver.chrome.options import Options
import time
import re

chrome_opt = Options()
chrome_opt.binary_location = '/usr/bin/google-chrome-stable'
chrome_opt.add_argument('--headless')
chrome_opt.add_argument('--disable-gpu')
chrome_opt.add_argument('--window-size=1366,768')
chrome_opt.add_argument("--no-sandbox")

url = "http://webserver:8088/admin/view"
admin_cookie = {
        "name": "admin",
        "value": open("/keys/admin_cookie.txt").read(),
        }
while 1:
    try:
        browser = webdriver.Chrome(executable_path='/var/xssbot/chromedriver', chrome_options=chrome_opt)
        print('get urls')
        print(url)
        browser.get(url)
        print(admin_cookie)
        browser.add_cookie(admin_cookie)
        browser.get(url)
        print('ok')
        print(browser.page_source)
        time.sleep(1)
        browser.quit()
    except Exception as e:
        print (e)
        try: 
            print(e.message)
        except e:
            print(e)
        continue
