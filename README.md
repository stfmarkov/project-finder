# project-finder

A cli tool to manage local git projects

## Commands

- list - Lists all available projects. A project can be selected with "space" or "enter" keys to take additional actions

- config - Set home directory, add project directories, add projects to exclude in the listing and more
  
  - Can add additional directories to search for projects ( by popular demand ) 

  - Set home directory - ( Not implemented )

  - Add projects to exclude in the listing ( Not implemented )

- find ( Not implemented ) - Directly find a project by alias

## Project actions

- Navigate to project - Navigates to the previously selected project 

- Open - Opens the project in vsCode

- Config ( Not implemented ) - Creates a project specifi config  

  - Set name ( Not implemented ) - Sets an alias for the project 

  - Commands ( Not implemented ) - Shows a list of commands and actions for them. The commands will be executed in order when a "run" action is executed

    - Add command ( Not implemented )

    - Remove command ( Not implemented )

    - Swap ( Not implemented ) - Moves a command up or down. A command that is currently pointed at by the cursor can be moved up or down with specified keys 

- Run ( Not implemented ) - Runs the project in a way specified by its config