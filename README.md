# Money Manager

This project is focused on managing financial transactions and related data.

## Running the Project (Development)

To set up and run the project in a development environment, follow these steps:

1.  **Set Environment Variables**
    
    First, make sure the necessary environment variables are set. This is done by sourcing the `env.sh` script:   
    `source env.sh` 
    
    **Note:** Ensure that `env.sh` contains the correct credentials and settings for your development environment. If any credentials are missing or need updating, please reach out to the team for the latest information.
    
2.  **Run the Application**
    
    With the environment variables set, you can start the application using the following command:
    `go run ./app/main.go` 
    
    This will execute the main application and you should see the output in your terminal.

## Troubleshooting

-   **Ensure `env.sh` is Executable**: If you encounter issues with sourcing the script, ensure it has the correct permissions:
    `chmod +x env.sh` 
    
-   **Verify Environment Variables**: Confirm that the environment variables are correctly set by running:
    
    <code>echo $DB_USER</code>
    <code>echo $DB_PASSWORD</code>
	<code>echo $DB_HOST</code>
    <code>echo  $DB_PORT</code>
    <code>echo $DB_NAME</code>
     
    
    Each command should print the corresponding variable's value.
    
-   **Script Context**: Remember that `source env.sh` needs to be run in the same terminal session where you plan to execute your Go application.
    
-   **Check for Missing Credentials**: If any application-specific configuration or credentials are missing, consult with the team to obtain them.
    

## Additional Notes

-   Ensure you have all necessary dependencies installed.
-   Refer to the project's README or documentation for more information on development setup and configuration.
