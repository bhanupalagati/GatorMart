
ReadMe.md 

GatorMart Documentation

Sprint 1

Team Members:

Nitin Ramesh

Bhanu Prakash Reddy Palagati

Vamsi Krishna Reddy

Gowtham Reddy Eda

**Introduction:**

The GatorMart application is an online market build with Angular and Go language. Built from the ground up, the application aims to seamlessly connect buyers to their products and sellers to their target audience. Its standout features include allowing the users to select a target audience while buying or selling. This allows to buy products or sell products in single or wholesale quantities while specifying a target audience such as students, professionals, farmers, etc.


**Tech Stack:**

Frontend: Angular 11.0 with TypeScript

Backend: Go Language

Database: MySQL

**Frontend:**

The front end of the application is built using Angular 11. Angular is a framework rather than a package that provides all the essential functional requirements out of the box. This will force the developers to follow a pattern, as a result, we have fewer decisions to take for the organization and spend time more on what matters.

We have used material design and there is a package named Angular Materials which helps us to implement the material style components easily and efficiently. 

The application demo consists of two main views: the “List view” and the “Detailed view”:

1) List View:

The List view is the main screen that greets the user when opening the GatorMart application. It is a block view of all the products, that have been hand-picked by our algorithms to suit the user’s taste.

This method favors readability, as the most important product information is conveyed to the user directly as they scroll the different listings on the site.

1) Detailed View:

The Detailed View is presented when the user clicks on a product in the List view. This view, as the name suggests, gives more detailed information regarding the selected product such as:

- A carousel of images of the product
- Age of the product
- Detailed description
- Price etc.
- Option to Edit/Remove the Add

All details are conveniently placed in one area with multiple images to give a better overall understanding of the product.

**Backend:**

REST API’s

For Sprint 1, two main endpoints have been created:

“/product”: (***POST Method*** ) This endpoint  is used to post an advertisement of product in Gator Mart.

“/products”: (***GET Method*** ) This endpoint  is used retrieve all products listed in Gator Mart application.

Fields such as “Title”, “secondary title”, “image”, ” price”, ”simpledescription”, ”description”, ”city” , “state”, “location”, “target”, “category”, “postedby”, “posteddate”, “condition”, “age”, and “adstatus” can be filled to create a post.

**Steps to Compile and Run:**

1. Navigate to the “backend” folder.
1. Run the following command in the terminal to start the server:

***go run main.go***

***go mod init*** 

The above command is used to create go.mod file which acts as a tracker for the dependencies.

1. Navigate to the “frontend” folder.
1. Run the following command in the terminal to compile the frontend code and render it in the browser:

***ng serve***

1. View the application by going to the address mentioned in the terminal. Example: http://localhost:4200

**Demo Videos:**

Frontend: <https://drive.google.com/file/d/1H4q-aU_TZIqPZ-yFm8HLCq5YkgH0Toki/view?usp=sharing> 

Backend: <https://drive.google.com/file/d/1M4xew8792Bq1IFe_W-Z7YrSX_LGjBOTG/view?usp=sharing> 



