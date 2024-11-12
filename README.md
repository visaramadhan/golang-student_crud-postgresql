CLI CRUD Application for Student Management
This application is a simple Command-Line Interface (CLI) for performing CRUD (Create, Read, Update, Delete) operations on student data. It also includes basic login and logout functionality.

Prerequisites
Go is installed on your system.
The database is set up and configured as required by the application.
Environment Setup
Before running the application, ensure your environment is properly configured. Follow these steps:

Clone the Repository and navigate to the application directory:

bash
git clone <your-repo-URL>
cd <application-directory-name>
Install Required Packages Run the following commands to install any required Go modules:

bash

go mod tidy

Set Up the Database Ensure the database is running and properly configured. Update the database connection details as required in the database package (e.g., database.go file). The application uses a PostgreSQL database, so make sure you have set up the database schema as needed.

Environment Variables
Configure any required environment variables for database connection (e.g., DB_USER, DB_PASSWORD, DB_NAME). You can use a .env file or set them directly in your terminal.

Run the Application Execute the following command to start the CLI:

bash
go run main.go

Usage Guide
After launching the program, you can perform various actions based on your login status.

1. Endpoint Selection
On the first run, the program will prompt you to enter an endpoint:
endpoint : 20 //type 20

2. Authentication
If you are not logged in, enter login when prompted to choose an action. Follow these steps:

plaintext
Choose action (login, exit):
login //typing

Enter username: user1
Enter password: password123

If login is successful, you will be taken to the CRUD menu.
If login fails, retry with the correct username and password.

3. CRUD Menu (Once Logged In)
After successfully logging in, you will access the menu to perform CRUD operations on student data. Available actions include:

plaintext
Salin kode
Enter Choice (add, update, view, delete, logout, exit):
3.1 Add Student (add)
To add a new student:

plaintext
Salin kode
Enter Choice (add, update, view, delete, logout, exit):
add
The program will prompt you for the necessary student information.

3.2 Update Student (update)
To update an existing student's data:

plaintext
Salin kode
Enter Choice (add, update, view, delete, logout, exit):
update
The program will ask for the student ID and the updated information.

3.3 View Students (view)
To view all students:

plaintext
Salin kode
Enter Choice (add, update, view, delete, logout, exit):
view
The program will display a list of all students in the database.

3.4 Delete Student (delete)
To delete a student's data:

plaintext
Salin kode
Enter Choice (add, update, view, delete, logout, exit):
delete
The program will prompt for the student ID of the record to be deleted.

3.5 Logout
To log out of the current session:

plaintext
Salin kode
Enter Choice (add, update, view, delete, logout, exit):
logout
You will be redirected back to the main menu to log in again or exit.

4. Exiting the Application
To exit the application, select exit at any prompt:

plaintext
Salin kode
Choose action (login, exit):
exit
or

plaintext
Salin kode
Enter Choice (add, update, view, delete, logout, exit):
exit
Code Overview
The main program is in main.go and uses several packages such as database, handler, repository, and service to handle CRUD operations and authentication.

License
Specify the license for this application here (if applicable).

This README provides a comprehensive guide on setting up and using the CLI CRUD application for student management.

