# About

CLI app that
1. Reads .tsuki/data/
	emails.db //to and from 
	transactions.db //to and from 
	events.db //people 
	notes.db //people
	

2. Contructs contacts and connections
3. UI to change these contacts and connections 

# How it works
The app works on a view edit model

Views
This is a page. If a contact isn't selected, this can be simply:
```
___________________________________________________________________________________
| Name | Most Recent Contact | Emails Pending Response | Emails Awaiting Response |
|---------------------------------------------------------------------------------|
| A    | Yesterday           | 2                       | 0                        |
|---------------------------------------------------------------------------------|
| B    | 2 days ago          | 0                       | 2                        |
|---------------------------------------------------------------------------------|
```
If a contact is selected:
```
_____________________________________________________
| Name                  | Shashank Rao              |
|---------------------------------------------------|
| Email                 |  real.email@fakeEmail.com |
|---------------------------------------------------|
| Phone                 | +0 12345 67890            |

```

```
Shashank Rao
@fakeemail.com
+00 01233456789

Likes to code

Emails Unread: // select * where name = name and status = unread



Emails Awaiting Response:

```









































