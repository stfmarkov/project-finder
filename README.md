# project-finder

A cli tool to manage local git projects

## Commands

- fetch - Finds and saves all the projects in the config file

- list - Lists all available projects. A project can be selected with "space" or "enter" keys to take additional actions
  
  - All projects are listed form cash

- config ( Not implemented. Some functionality is move as project actions ) - Set home directory, add project directories, add projects to exclude in the listing and more
  
  - Can add additional directories to search for projects ( by popular demand ) 

  - Set home directory - ( Not implemented )

  - Add projects to exclude in the listing ( Not implemented )

- find - Directly find a project by alias, list all that match the search, the search is applied over the already found and saved in the config files

## Project actions

- Navigate to project - Navigates to the previously selected project 

- Open - Opens the project in vsCode

- Config - Creates a project specifi config  

  - Set name ( Not implemented ) - Sets an alias for the project 

  - Commands - Shows a list of commands and actions for them. The commands will be executed in order when a "run" action is executed

    - Add command

    - List commands

    - Remove command

    - Swap ( Not implemented. Will be implemented when actions are executable by "action buttons" instead of submenues ) - Moves a command up or down. A command that is currently pointed at by the cursor can be moved up or down with specified keys 

- Run - Runs the project in a way specified by its config

## UI

- Find all action keys and implement the actual actions ( Not implemented )  

  - Direction keys

  - Set specific configurable action keys for listings reducing the need to have multiple submenues ( Not implemented ) 

- Show cursor. The currently selected item in a menue should be with different color (green may be, it is hacky)

- Add colors to the currently selected items like projects, actions, commands and so on

- List projects after previous actions are done

- Better error handling and messages ( Not implemented )
