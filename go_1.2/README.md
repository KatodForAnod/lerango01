# Learning by coding

## Initial app, maps, v1.0

Create interactive command line application to help keeping the eye on the inventory

1. First ask for one of possible command: "add inventory", "remove inventory", "list inventories", "add user", "assign inventory to", "list users", "remove user", "help"
2. Every command should work both interactively and as parameter in the command line
3. Store users and inventories in Go maps

The app shoold work like this:

'add inventory _name_ _description_', for example 'add inventory "MacBook" "MacBook Pro 15, RAM 8GB, SDD 256GB" should ad new inventory in the map
'add user _name_' should add new inventory user, for example 'add user "Pete Anderson"'
'assign _inventoryId_ / _inventoryName_ to _userId_ / _username_' should assign the existing user to inventory, for example 
'assign 1 to "Pete Anderson"

'list inventories' shouls output the list of all available inventories, for example
1  | MacBook | MacBook Pro 15, RAM 8GB, SDD 256GB    | Pete Anderson
2  | Book    | Go PROGRAMMING, Mark Summerfeld, 2018 | Alisa Isterina

Command 'help' should display a short usage guide.

Add new functionality to the new branch feature/maps and send merge request to the branch develop

## App upgrade, adding sql, errors processing, v1.1

To be written...

## App upgrade, creating a service, intercept panic v1.2

To be written...

## App upgrade, covering with tests, run the coverage v1.3

To be written...

## App upgrade, creating web app, use HTML templates v1.4

To be written...

## App upgrade, adding a bot, register new users, add approvals v1.5

To be written...
