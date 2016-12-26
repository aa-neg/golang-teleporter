# golang-teleporter

A simple tool which allows you to set alias for various paths and 'teleport' to them.


##Installation

For windows clone the repo and add the folder to your path. Test if it is wroking by reopening a new command window and typing 'tp'

##Example Usage
```
#List current alias
tp list

#Add current path to alias
tp add current 

#Change to a random directory
cd /

#Teleport back to current
tp to current

#Remove the alias
tp remove current
```