# URL Parser

URL Parser is a web application through which you can submit a website URL and it will show you general information
about the contents of the submitted web page.

The application provides you with information like:

- HTML Version of the web page
- Page Title
- Headings count for each HTML heading level
- Count of Internal Links included in the page
- Count of External Links
- Count of Inaccessible Links
- If the page contains a Login Form or not

### Example

Let's assume that the user wants to know more details about Google's Home Page.

```
URL: "https://google.com"

Result: 
    - HTML Version:
    - Title: Google
    - Heading Levels:
        - H1: 0
        - H2: 0
        - H3: 0
        - H4: 0
        - H5: 0
        - H6: 0
    - Links:
        - Internal: 14
        - External: 5
        - Inaccessible: 6
    - Page Has Login Form: false
```

## How to run the application

### 1. Using a bin executable file:

- You can find a bin executable version of the application in the project's bin directory.
- Copy the bin file to the project's root directory and then execute it. Finally, navigate to "http://localhost:8080"
  from your favorite browser.

### 2. Terminal:

- Navigate to the directory of the application on your computer, and run the following command

```
go run ./cmd/web -port=port_of_your_choice
```

- Then open your favorite browser and navigate to "http://localhost:8080". By default, the application runs on Port
  8080, but you can change this to run on the Port of your choice, by specifying it in the above command.

### 3. Using Makefile Commands:

- First, you should install Make on your machine

```
# for linux users
$ sudo apt install make

# for macOS users
$ brew install make

# for windows users (using Chocolately package)
> choco install make
```

- After installing make, you'll have access to all commands specified in the project's Makefile
- So, to run the applications, just enter:

```
make run
```

### 4. Using Docker

- You can build a docker image for the application using the Dockerfile available in the project's root directory.
- Build the image by running the following command from your terminal:

```
$ docker build -t url-parser:latest .
```

- If you've installed make, you can build the image using the following command:

```
$ make build-image
```

- Then to create and run a docker container out of the image, enter the command:

```
$ docker run -d --rm -p 8080:8080 --name url-parser-container url-parser
```

- or using make:

```
$ make run-contianer
```

## Deployment

- The project includes Kubernetes deployment files to deploy the application on a Kubernetes Engine.
- The project also includes a Github actions workflow file to manage a CI/CD pipeline through which you can test, build
  and deploy your latest changes to Cloud.
- The deployment and Github actions configuration files use Google Kubernetes Engine, and Google Cloud Platform to
  deploy the application. You can replace them with your favorite cloud provider.

## Improvements

- The application's performance can be improved by saving the link and its parsed details in a database. Such that when
  a user requests the details of a URL that is available in our database, the application returns the details from the
  database, rather than fetching the details from the beginning.
- We also need to set an expiry date for a URL's details in the database. When a user requests the details for a URL
  that hasn't been fetched for more than, for example, a week, the URL is fetched again rather than returning the
  details from the database. This way, we can make sure that we are returning our users link details that are up to-date
  with changes that can be made to the URLs request by our users.