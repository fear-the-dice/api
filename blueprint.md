FORMAT: 1A   
HOST: fear-the-dice-api.herokuapp.com

# Fear the Dice!

A tool for helping manage combat in a turn based environment. Allowing DM's/GM's to control stats such as AC, HP, and 'damage' taken for a group.

#Group API

## Monsters [/monsters]

Monsters resource for accessing all monsters as a group.

### Get API Options [OPTIONS]

Retrieve all options avaliable for this specific resource.

+ Response 204
    + Headers
    
            Access-Control-Allow-Headers: Origin,Accept,Content-Type,Authorization
            Access-Control-Allow-Methods: GET,POST,PUT,DELETE,HEAD

### Get Monsters [GET]

Retrieve all monsters in a JSON array.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs
                
+ Response 200
    + Headers
    
            Content-Type: application/json; charset=utf-8
    + Body
          
             [
               {
                 "id": "555e45fc60a3dca513e6265e",
                 "monster": "Black Bear",
                 "initiative": 1,
                 "ac": 12,
                 "hp": 11,
                 "health": 11,
                 "speed": 30,
                 "damage": 0,
                 "challange": 0,
                 "xp": 25,
                 "manual": 343,
                 "str": "15 (+1)",
                 "dex": "15 (+1)",
                 "con": "15 (+1)",
                 "int": "15 (+1)",
                 "wis": "15 (+1)",
                 "cha": "15 (+1)"
               },
               {
                 "id": "555e45fc60a3dca513e6265f",
                 "monster": "Bandit",
                 "initiative": 1,
                 "ac": 12,
                 "hp": 11,
                 "health": 11,
                 "speed": 30,
                 "damage": 0,
                 "challange": 0,
                 "xp": 25,
                 "manual": 343,
                 "str": "15 (+1)",
                 "dex": "15 (+1)",
                 "con": "15 (+1)",
                 "int": "15 (+1)",
                 "wis": "15 (+1)",
                 "cha": "15 (+1)"
               }
             ]
         
### New Monster [POST]

Create a new monster by passing the JSON object to the server.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs
    + Body
    
            {
                "monster": "Bugbear Chief",
                "initiative": 15,
                "ac": 27,
                "hp": 65,
                "health": 65,
                "speed": "30",
                "damage": 0,
                "challange": 2,
                "xp": 450,
                "manual": 33,
                "str": "15 (+1)",
                "dex": "15 (+1)",
                "con": "15 (+1)",
                "int": "15 (+1)",
                "wis": "15 (+1)",
                "cha": "15 (+1)"
            }

+ Response 201
    + Headers
    
            Content-Type: application/json; charset=utf-8

    + Body
    
            {
                "id": "555e45fc60a3dca513e6265d",
                "monster": "Bugbear Chief",
                "initiative": 15,
                "ac": 27,
                "hp": 65,
                "health": 65,
                "speed": "30",
                "damage": 0,
                "challange": 2,
                "xp": 450,
                "manual": 33
            }

## Monster [/monsters/{monster_id}]

Monster resource for accessing individual monster objects. 

+ Parameters 
    + monster_id: `555e45fc60a3dca513e62665` (string), required  - ID of the 'monster'

### Get API Options [OPTIONS]

Retrieve all options avaliable for this specific resource.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs

+ Response 204
    + Headers
    
            Access-Control-Allow-Headers: Origin,Accept,Content-Type,Authorization
            Access-Control-Allow-Methods: GET,POST,PUT,DELETE,HEAD

### Get Monster [GET]

Retrieve a specific monster.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs

+ Response 200
    + Headers
    
            Content-Type: application/json; charset=utf-8

    + Body 
    
            {
                "id": "555e45fc60a3dca513e6265d",
                "monster": "Bugbear Chief",
                "initiative": 15,
                "ac": 27,
                "hp": 65,
                "health": 65,
                "speed": 30,
                "damage": 0,
                "challange": 2,
                "xp": 450,
                "manual": 33
            }

### Edit Monster [PUT]

Edit an existing monster by passing a new JSON object. This will replace the existing object so it is wise to ensure all fields are set.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiYWVjMDQzMC04NjAxLTQ1NTMtODUzYi1lN2FjNGI3Y2E3NGUiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.BYjR2Xui6nm8mXljztpKCLW7MGNP50hFxyOkihJuUNc
            
    + Body
    
            {
                "id": "555e45fc60a3dca513e62663",
                "initiative": 1,
                "ac": 14,
                "hp": 19,
                "health": 19,
                "damage": 0,
                "speed": "30/50",
                "monster": "Draft Horse",
                "challange": 0,
                "xp": 50,
                "manual": 321,
                "str": "14 (+1)",
                "dex": "14 (+1)",
                "con": "10",
                "int": "9 (-1)",
                "wis": "(16 (+2)",
                "cha": "(16 (+2)"
            }

+ Response 201
    + Headers
    
            Content-Type: application/json; charset=utf-8

    + Body
    
            {
                "id": "555e45fc60a3dca513e62661",
                "initiative": 1,
                "ac": 14,
                "hp": 19,
                "health": 19,
                "damage": 0,
                "speed": "30/50",
                "monster": "Draft Horse",
                "challange": 0,
                "xp": 50,
                "manual": 321,
                "str": "14 (+1)",
                "dex": "14 (+1)",
                "con": "10",
                "int": "9 (-1)",
                "wis": "(16 (+2)",
                "cha": "(16 (+2)"
            }

### Delete Monster [DELETE]

Delete an existing monster.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs

+ Response 200
    + Headers
    
            Content-Type: text/plain; charset=utf-8

## Players [/players]

Players resource for accessing all players as a group.

### Get API Options [OPTIONS]

Retrieve all options avaliable for this specific resource.

+ Response 204
    + Headers
    
            Access-Control-Allow-Headers: Origin,Accept,Content-Type,Authorization
            Access-Control-Allow-Methods: GET,POST,PUT,DELETE,HEAD


### Get Players [GET]

Retrieve all players in a JSON array.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs

+ Response 200
    + Headers
    
            Content-Type: application/json; charset=utf-8
            
    + Body
       
             [
               {
                 "id": "555e460260a3dca513e6266f",
                 "name": "Robin",
                 "character": "Strife",
                 "initiative": 18,
                 "ac": 16,
                 "hp": 19,
                 "health": 19,
                 "speed": 25,
                 "damage": 0
               },
               {
                 "id": "555e460260a3dca513e6266e",
                 "name": "Shawn G.",
                 "character": "M.F. Jones",
                 "initiative": 13,
                 "ac": 13,
                 "hp": 10,
                 "health": 10,
                 "speed": 25,
                 "damage": 0
               }
             ]

### New Player [POST]

Create a new player by passing the JSON object to the server.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs
            
    + Body
    
            {
                "name": "Robin",
                "character": "Strife",
                "initiative": 18,
                "ac": 16,
                "hp": 19,
                "health": 19,
                "speed": 25,
                "damage": 0
            }
         
+ Response 201
    + Headers
    
            Content-Type: application/json; charset=utf-8

    + Body    

            {
                "id": "555e460260a3dca513e6266f",
                "name": "Robin",
                "character": "Strife",
                "initiative": 18,
                "ac": 16,
                "hp": 19,
                "health": 19,
                "speed": 25,
                "damage": 0
            }

## Player [/players/{player_id}]
Player resource for accessing individual player objects.
 
+ Parameters 
    + player_id: `555e460260a3dca513e6266d` (string), required - ID of the 'player'

### Get API Options [OPTIONS]

Retrieve all options avaliable for this specific resource.

+ Response 204
    + Headers
    
            Access-Control-Allow-Headers: Origin,Accept,Content-Type,Authorization
            Access-Control-Allow-Methods: GET,POST,PUT,DELETE,HEAD

### Get Player [GET]

Retrieve a specific player.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs

+ Response 200
    + Headers
    
            Content-Type: application/json; charset=utf-8
            
    + Body
    

            {
                "id": "555e460260a3dca513e6266f",
                "name": "Robin",
                "character": "Strife",
                "initiative": 18,
                "ac": 16,
                "hp": 19,
                "health": 19,
                "speed": 25,
                "damage": 0
            }
         
### Edit Player [PUT]

Edit an existing player by passing a new JSON object. This will replace the existing object so it is wise to ensure all fields are set.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiYWVjMDQzMC04NjAxLTQ1NTMtODUzYi1lN2FjNGI3Y2E3NGUiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.BYjR2Xui6nm8mXljztpKCLW7MGNP50hFxyOkihJuUNc
            
    + Body

            {
                "id": "555e460260a3dca513e6266d",
                "initiative": 16,
                "ac": 16,
                "hp": 20,
                "health": 4,
                "damage": 16,
                "speed": "",
                "name": "Jamie",
                "character": "Thomas",
                "thumb": ""
            }

+ Response 201
    + Headers
    
            Content-Type: application/json; charset=utf-8

    + Body   

            {
                "id": "555e460260a3dca513e6266d",
                "initiative": 16,
                "ac": 16,
                "hp": 20,
                "health": 4,
                "damage": 16,
                "speed": "",
                "name": "Jamie",
                "character": "Thomas",
                "thumb": ""
            }

### Delete Player [DELETE]

Delete an existing player.

+ Request
    + Headers
        
            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIzZjJmMWI2My04YzM2LTQ2MWEtOWJjZi00OTRmZDY5YmQxMzIiLCJpYXQiOjg2NDAwLCJzdWIiOiI4MjA1ZDNiMi0zZTczLTQzMmItYjdlYi1iNzNmNzM4MThkODMiLCJhdWQiOiIqLmZlYXJ0aGVkaWNlLmNvbSIsImlzcyI6ImF1dGguZmVhcnRoZWRpY2UuY29tIn0.IemFBgUUEsuBYgxQkH9ijb18Hv-fUckOSiwr1Zjf3Vs

+ Response 200
    + Headers
   
            Content-Type: text/plain; charset=utf-8
