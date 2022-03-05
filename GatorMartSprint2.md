# GatorMart Documentation

March 4th, 2022

**Sprint 2**

# Team Members:

Nitin Ramesh

Bhanu Prakash Reddy Palagati

Vamsi Krishna Reddy Komareddy

Gowtham Reddy Eda

# Introduction:

The GatorMart application is an online market build with Angular and Go language. Built from the ground up, the application aims to seamlessly connect buyers to their products and sellers to their target audience. Its standout features include allowing the users to select a target audience while buying or selling. This allows to buy products or sell products in single or wholesale quantities while specifying a target audience such as students, professionals, farmers, etc.

# Tech Stack:

Frontend: Angular 11.0 with TypeScript

Frontend Test cases: Cypress and Jest Framework

Backend: Go Language

Database: MySQL


# Frontend:

The front end of the application is built using Angular 11. Angular is a framework rather than a package that provides all the essential functional requirements out of the box. This will force the developers to follow a pattern, as a result, we have fewer decisions to take for the organization and spend time more on what matters.

We have used material design and there is a package named Angular Materials which helps us to implement the material style components easily and efficiently. 

Test cases have been written using Jest for unit testing and Cypress for integration testing, this has ensured the smooth integrated working of the frontend and backend. 

The application sprint 2 demo consists of few main views: the “List view” and the “Detailed view”:

**1) List View:**

The List view is the main screen that greets the user when opening the GatorMart application. It is a block view of all the products, that have been hand-picked by our algorithms to suit the user’s taste.

The products are displayed here in a small window shaped material design as shown below:

This method favors readability, as the most important product information is conveyed to the user directly as they scroll the different listings on the site.

**2) Detailed View:**

The Detailed View is presented when the user clicks on a product in the List view. This view, as the name suggests, gives more detailed information regarding the selected product such as:

A carousel of images of the product

Age of the product

Detailed description

Price etc.

Option to Edit/Remove the Add

All details are conveniently placed in one area with multiple images to give a better overall understanding of the product.

**3) Create Product View:**

The create product view confirses of a form where the user can add all relevant details and set the product up for listing. This includes the main image, carousel images and device specific data. 



# 1. GetProducts 

**URL**: localhost/products

**Method**: Get
```
    {
        "ID": 1,
        "CreatedAt": "2022-03-03T22:40:26.088-05:00",
        "UpdatedAt": "2022-03-03T22:40:26.088-05:00",
        "DeletedAt": null,
        "title": "computer desk",
        "secondary_title": "used",
        "imageUrl": "coming soon",
        "price": 25,
        "simple_desc": "Very less used",
        "description": "Looks new, used very less price negotiable",
        "city": "gainesville fl",
        "state": "florida",
        "location_lat": "maps",
        "location_long": "maps",
        "target": "student",
        "category": "Furniture",
        "posted_by": "user1",
        "posted_date": "3rdMar",
        "condition": "Like new",
        "age": 1,
        "status": "avaiable",
        "images": "https://store.haworth.com/products/soji-mesh-office-chair?variant=40309843099857&utm_source=sshopping&utm_medium=netsertive "
    },

    {
        "ID": 2,
        "CreatedAt": "2022-03-03T22:40:26.088-05:00",
        "UpdatedAt": "2022-03-03T22:40:26.088-05:00",
        "DeletedAt": null,
        "title": "computer desk",
        "secondary_title": "used",
        "imageUrl": "coming soon",
        "price": 25,
        "simple_desc": "Very less used",
        "description": "Looks new, used very less price negotiable",
        "city": "gainesville fl",
        "state": "florida",
        "location_lat": "maps",
        "location_long": "maps",
        "target": "student",
        "category": "Furniture",
        "posted_by": "user1",
        "posted_date": "3rdMar",
        "condition": "Like new",
        "age": 1,
        "status": "avaiable",
        "images": "https://store.haworth.com/products/soji-mesh-office-chair?variant=40309843099857&utm_source=sshopping&utm_medium=netsertive "
},
```
It displays all the products that are posted into the database and displays it to the user.

**Possible Response codes**: 200,401

**Examples**: 200 – Success and 401 – User not authorized.

# 2. SaveProduct

**URL**: localhost/product

**Method** – Post

**Payload**:
```
{
        "ID": 1,
        "CreatedAt": "2022-03-03T22:40:26.088-05:00",
        "UpdatedAt": "2022-03-03T22:40:26.088-05:00",
        "DeletedAt": null,
        "title": "computer desk",
        "secondary_title": "used",
        "imageUrl": "coming soon",
        "price": 25,
        "simple_desc": "Very less used",
        "description": "Looks new, used very less price negotiable",
        "city": "gainesville fl",
        "state": "florida",
        "location_lat": "maps",
        "location_long": "maps",
        "target": "student",
        "category": "Furniture",
        "posted_by": "user1",
        "posted_date": "3rdMar",
        "condition": "Like new",
        "age": 1,
        "status": "avaiable",
        "images": "https://store.haworth.com/products/soji-mesh-office-chair?variant=40309843099857&utm_source=sshopping&utm_medium=netsertive "
    },
```
It inserts the data into the database with all the features of the product and the images into the database.

**Possible Response codes**: 401,200,400

**Examples**: 401 – User not authorized and 200 –The product was successfully saved, 400 bad request.

# 3. GetProduct

**URL**: localhost/product/:id

**Method**: Get

since it's a GET method the id is passed as a URL parameter.


It displays the specific product that the user is searching for using product id.

**Possible Response codes**: 401,200, 400

**Examples**: 401 – User not authorized and 200 –The product that the user is searching for does not exist, 400 bad request

# 4. UpdateProduct

**Method**: Put

**URL**: localhost/product/:id

**Pay load**:

```
{
"title":"computer desk",
"secondary_title":"used",
"imageUrl":"coming soon",
"price":25,
"simple_desc":"Very less used",
"description":"Looks new, used very less price negotiable",
"city":"gainesville fl",
"state":"florida",
"location_lat":"maps",
"location_long":"maps",
"target":"student",
"category":"Furniture",
"posted_by":"user1",
"posted_date":"3rdFeb",
"condition":"Like new",
"age":1,
"status":"avaiable",
"images":"images"
}
```


It updates the information about the product, and it modifies the details.

**Possible Response codes**: 401,400,200

**Examples**: 401 – User not authorized and 400 – Bad request, 200 product information updated.

# 5. DeleteProduct

**Method**: Delete

**URL**: localhost/product/:id

In this method the id is passed as a URL parameter.

It deletes all the information from the database if user deletes the post

Possible Response codes: 401,400,200

**Examples**: 401 – User not authorized and 400 – The product that the user is searching for does not exist, 200 Successfully deleted

# 6. UploadImage

**Method**: Post

**URL**: localhost/product/upload

**Pay load**:
```
Form-data
Key   : photo
Value : image.png
```
It saves the image in Amazon S3 bucket and returns the image url.

**Possible Response codes**: 401,500,200

**Examples**: 401 – User not authorized , 500 – Internal Server Error and 200 – Image uploaded successfully.


# 7. Register 

**Method**: Post

**URL**: localhost/gatormart/register

**Pay load**:
```
{
"name":"user4",
"email":"example4@gmail.com",
"password":"Password@004"
}
```
Registers the user and adds the data into the database. The id of the user will be increasing by 1. Mandatory fields are name, email and password in the user registration page.

**Possible Response codes**: 200,400

**Examples**: 200 – Success and 400 – Bad request

# 8. Login

**Method**: Post

**URL**: localhost/gatormart/login

**Pay load**:
```
{

"email":"example4@gmail.com",
"password":"Password@004"
}
```
It allows the user to login by checking the username and the password. If it matches it allows the user to login else, it displays the failure status.

**Possible Response codes**: 200,400

**Examples**: 200 – Success and 400 -Failure.

# 9. Logout

**Method**: Post

**URL**: localhost/gatormart/logut

This method allows the user to log out.


