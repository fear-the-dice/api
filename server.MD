FORMAT: 1A

# Lets Kill Things!

A tool for helping manage combat in a turn based environment. Allowing DM's/GM's to control stats such as AC, HP, and 'damage' taken for a group.

# Monsters

## Monsters Collection [/monsters]

### Get All Monsters [GET]

+ Response 200 (application/json)
    
     ```    
     [
       {
         "id": "a144238b-7418-7c17-bbd2-425fb27e3866",
         "monster": "Bandit",
         "initiative": 12,
         "ac": 12,
         "hp": 11,
         "health": 11,
         "speed": 30,
         "damage": 0,
         "challange": 0,
         "xp": 25,
         "manual": 343
       },
       {
         "id": "1c74f40-8dfd-4683-455a-7c124ed0df82",
         "monster": "Bandit Captin",
         "initiative": 9,
         "ac": 15,
         "hp": 65,
         "health": 15,
         "speed": 30,
         "damage": 0,
         "challange": 2,
         "xp": 450,
         "manual": 344
       },
       {
         "id": "ab8aaec9-3f8f-9988-39ec-95a57ad52eb1",
         "monster": "Bugbear",
         "initiative": 18,
         "ac": 16,
         "hp": 27,
         "health": 27,
         "speed": 30,
         "damage": 0,
         "challange": 1,
         "xp": 200,
         "manual": 33
       },
       {
         "id": "b8c2a999-e892-0964-a0f4-e36db31d2307",
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
     ]
     ```

## Monster [/monsters/{id}]

### Get Monster [GET]

+ Prameters
    + id (sring) - ID of the 'monster'
+ Response 200 (application/json)

     ```
    {
        "id": "b8c2a999-e892-0964-a0f4-e36db31d2307",
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
    ```

### New Monster [POST]

+ Request JSON Message

     ```
    {
        "id": "b8c2a999-e892-0964-a0f4-e36db31d2307",
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
    ```
+ Response 201 (text/plain)

    ```
    /monsters/b8c2a999-e892-0964-a0f4-e36db31d2307
    ```


# Players

## Players Collection [/players]

### Get Players [GET]

+ Response 200 (application/json)

     ```    
     [
       {
         "id": "779ed185-860e-a7a8-1f11-2d6ea1d073df",
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
         "id": "f056e8d4-4536-c142-2966-e2ff4ef529b9",
         "name": "Shawn G.",
         "character": "M.F. Jones",
         "initiative": 13,
         "ac": 13,
         "hp": 10,
         "health": 10,
         "speed": 25,
         "damage": 0
       },
       {
         "id": "9937eb1c-d82e-b872-a875-930472108b52",
         "name": "Keegan",
         "character": "Eliks",
         "initiative": 5,
         "ac": 16,
         "hp": 13,
         "health": 13,
         "speed": 25,
         "damage": 0
       },
       {
         "id": "689d07ca-b194-c7bc-6ba2-c28fd10edd8a",
         "name": "Jamie",
         "character": "Thomas",
         "initiative": 18,
         "ac": 16,
         "hp": 15,
         "health": 15,
         "speed": 25,
         "damage": 0
       }
     ]
     ```

## Player [/players/{id}]

### Get Player [GET]

+ Prameters
    + id (string) - ID of the 'player'
+ Response 200 (application/json)

     ```
    {
        "id": "779ed185-860e-a7a8-1f11-2d6ea1d073df",
        "name": "Robin",
        "character": "Strife",
        "initiative": 18,
        "ac": 16,
        "hp": 19,
        "health": 19,
        "speed": 25,
        "damage": 0
    }
     ```

### New Player [POST]

+ Request JSON Message

     ```
    {
        "id": "779ed185-860e-a7a8-1f11-2d6ea1d073df",
        "name": "Robin",
        "character": "Strife",
        "initiative": 18,
        "ac": 16,
        "hp": 19,
        "health": 19,
        "speed": 25,
        "damage": 0
    }
     ```
+ Response 201 (text/plain)

    ```
    /players/779ed185-860e-a7a8-1f11-2d6ea1d073df
    ```
