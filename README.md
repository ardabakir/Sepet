# Sepet

A small cart application that does CRUD actions on a cart.

Possible actions are 
 - Add product to cart
 - Remove product from cart
 - Update product in the cart
 - Get cart details
 - Empty cart

For this project, I created a codepipeline in AWS and deployed the lambda function using the buildspec.yml and serverless.yml files. 

I utilized repository pattern in the project. It wasn't necessary for a project of this size; however, I chose to implement it because it makes it easier to mock the database connection while testing. 

I added unit tests using testify to the project; however, due to time limitations, there are a limited amount of test cases. 

For the next steps of this project aside from some refactoring of the code:
 - An authentication service should be created
 - A signup functionality could be added, and both the user and their cart could be created at the same time
 - A method for getting detailed information about the products in the cart could be added
 - A method for updating the products table could be added
