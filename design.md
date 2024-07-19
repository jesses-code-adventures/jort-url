# db

## url table

id INT PRIMARY KEY AUTO INCREMENT
created_at DATETIME NOT NULL
user_id INT NOT NULL
url_long TEXT UNIQUE NOT NULL
url_short VARCHAR(32) NOT NULL
clicks INT NOT NULL

## user table

user_id INT PRIMARY KEY AUTO INCREMENT
user_name VARCHAR(128)
password TEXT

# ux

- visit homepage, be prompted to log in if not already
- after login, should see a nice text input where you can enter a url
- clicking GO (or whatever), or pressing enter should trigger the url shortening and copy the new url to the clipboard
- if the user has already shortened this url, it should return the existing url
- when a url is shortened, the textbox should be replaced with the shortened url as text and there should be a button below it - "shorten another url?"
- below the url shortening box there should be a list of previously shortened urls with their number of clicks, sorted in descending chronological order.
- there should be a dropdown menu at the top-right allowing the user to log out, or to switch to dark mode

# endpoints

## /user

- POST should create a new user with username and password

## /login

- POST should attempt to log a user in with the username and password, and should set browser cookies with a JWT on success

## /url

- POST should take a long url as data and return a short url, storing the combination in the db. this is an authenticated endpoint
- GET should return a list of the user's shortened URLs and their number of clicks in descending chronological order

## /logout

- POST should log the user out, deleting the local JWT.
