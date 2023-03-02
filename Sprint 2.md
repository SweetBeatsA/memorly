Our Client Address:
http://memorly.kro.kr/

Our Server Address:
http://api.memorly.kro.kr/

# **Issues our team plans to address for Sprint 2**

- Link backend and frontend
- Allow the user to create an account and log in from the website front end
- Create front end for a home page, as well as pages that can be accessed by clicking on different tabs on the header, such as login, card creation and card viewing/studying pages. 
- Probably in the next sprint, add card library page with folders, but maybe start to do some of the front end for this.

# **Which issues were successfully completed**
- The user can create an account from the front end sign up page
- The user will be stopped from creating an account if an account with the given email already exists in the database, if the password is not 8 characters, if there are special characters in the username
- The user can log in from the log in page
- If the user enters the wrong email/password, it will display a message telling them
- There are multiple different webpages which can be accessed from the header, though many of them are just dummy pages for now and need work on the feature implementation and formatting
- Completed the API to sign up, log in, and get a specific users data and linked it to the front end
- Connected a database to keep track of the user data

# **Which issues were not completed and why** 
- Home page needs to change if the user is logged in so that it does not prompt them to log in still
- There is a lot of messy and commented front end code currently which needs to be cleaned up
- It is still not possible to create or view flash cards yet
- Fix up the backend code
- Keep track of data sets created by users



# **List unit tests and Cypress tests for frontend**
Cypress:
- A test that goes to the home page and clicks to go to the login page

Unit tests:
- A test that checks that "Welcome to our CEN3031 Project!" is being displayed on the home page
- A test that checks the Memorly icon is being displayed in the header


# **List unit tests for backend**
Unit tests:
- A test that checks that a user profile was successfully created with the Sign Up API
- A test that checks if a user log in to the profile that was created with the Log In API
- A test that returns the user profile that were created previously using the authorization token with the Get User API

# **Detailed documentation of your backend API**

## **Sign Up User**

Registering an user with given email, password, and name.

**Expected Side Effect**

If the email is already in use, user need to use different email.

### **Path**

api.memorly.kro.kr/users/signup

### **Error**

| Code | Message             | Case                                    |
| ---- | ------------------- | --------------------------------------- |
| 400  | Binding Error       | When the required input does not filled |
| 400  | Email Already Taken | When the input email is already taken   |
| 500  | Database Error      | When the database does not response     |

## **Log In User**

Log In an registered user with given email, password.

**Expected Side Effect**

This api is working with registered user. If the user does not sign up our service, they need to sign up first to call this api.

### **Path**

api.memorly.kro.kr/users/login

### **Error**

| Code | Message            | Case                                             |
| ---- | ------------------ | ------------------------------------------------ |
| 400  | Binding Error      | When the required input does not filled          |
| 404  | No Matched User    | When the input email does not exist              |
| 401  | Incorrect Password | When the input password does not match with user |

## **Get User**

Getting an specific user information with given token which is set from the HTTP header as a key name "Authorization".

**Expected Side Effect**

This api is to check whether our JWT token is working well or not along with checking the information of User.  
Therefore, the accessToken value from either signUp or logIn api should be set to a HTTP Header. Failure to set the token or using an expired token will throw an error.

### **Path**

api.memorly.kro.kr/users/login

### **Error**

| Code | Message                         | Case                                  |
| ---- | ------------------------------- | ------------------------------------- |
| 400  | No Authorization Token Provided | When the token is not provided        |
| 401  | Invalid Token                   | When the given token value is invalid |
| 401  | Expired Token                   | When the given token is expired       |
| 404  | No Matched User                 | When the user email does not exist    |

Our API Documentation:
https://documenter.getpostman.com/view/12809852/2s93CRJqgL
