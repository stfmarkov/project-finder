# project-finder

A cli tool to manage local git projects

## Commands

- list - Lists all available projects. A project can be selected with "space" or "enter" keys to take additional actions
  
  - Save all the found projects in the config ( Not implemented )

  - Initialy list projects from the config and later update the list if needed ( Not implemented )

- config - Set home directory, add project directories, add projects to exclude in the listing and more
  
  - Can add additional directories to search for projects ( by popular demand ) 

  - Set home directory - ( Not implemented )

  - Add projects to exclude in the listing ( Not implemented )

- find ( Not implemented ) - Directly find a project by alias, list all that match the search, the search is applied over the already found and saved in the config files

## Project actions

- Navigate to project - Navigates to the previously selected project 

- Open - Opens the project in vsCode

- Config ( Not implemented ) - Creates a project specifi config  

  - Set name ( Not implemented ) - Sets an alias for the project 

  - Commands ( Not implemented ) - Shows a list of commands and actions for them. The commands will be executed in order when a "run" action is executed

    - Add command

    - List commands

    - Remove command ( Not implemented )

    - Swap ( Not implemented ) - Moves a command up or down. A command that is currently pointed at by the cursor can be moved up or down with specified keys 

- Run ( Not implemented ) - Runs the project in a way specified by its config

## UI

- Find all action keys and implement the actual actions ( Not implemented )  

  - Direction keys ( Not implemented )

- Show cursor ( Not implemented )

- Add colors to the currently selected items like projects, actions, commands and so on ( Not implemented )