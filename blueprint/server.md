FORMAT: 1A

# Fear the Dice!

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
         "initiative": 1,
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
         "initiative": 1,
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
         "initiative": 1,
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
         "initiative": 1,
         "ac": 27,
         "hp": 65,
         "health": 65,
         "speed": 30,
         "damage": 0,
         "challange": 2,
         "xp": 450,
         "manual": 33
       },
       {
         "id": "c345ffb4-e662-4d4d-927f-21ef99004fa8",
         "monster": "Black Bear",
         "initiative": 1,
         "ac": 11,
         "hp": 19,
         "health": 19,
         "speed": 40,
         "damage": 0,
         "challange": "0.5",
         "xp": 100,
         "manual": 318
       },
       {
         "id": "ecc3469f-0462-405b-b64f-c268288f42bc",
         "monster": "Cyclopse",
         "initiative": 1,
         "AC": 14,
         "HP": 138,
         "health": 138,
         "speed": 30,
         "damage": 0,
         "challange": 6,
         "xp": 2300,
         "manual": 45
       },
       {
         "id": "a31f0007-9685-4008-92d1-9f10c707534b",
         "monster": "Draft Horse",
         "initiative": 1,
         "AC": 10,
         "hp": 19,
         "health": 19,
         "speed": 40,
         "damage": 0,
         "challange": 0,
         "xp": 50,
         "manual": 321
       },
       {
         "id": "4c937026-6bb8-4f6b-aa34-c9910f2b88af",
         "monster": "Flying Sword",
         "initiative": 1,
         "AC": 17,
         "hp": 17,
         "health": 17,
         "speed": "0/50",
         "damage": 0,
         "challange": 0,
         "xp": 50,
         "manual": 20
       },
       {
         "id": "73571fb9-e23d-4b2c-846a-aeb21aadd1c8",
         "monster": "Gelatinous Cube",
         "initiative": 1,
         "AC": 6,
         "hp": 64,
         "health": 64,
         "speed": 15,
         "damage": 0,
         "challange": 2,
         "xp": 450,
         "manual": 242
       },
       {
         "id": "b99f8b30-50fd-45c6-bd16-51b45640ebee",
         "monster": "Goblin",
         "initiative": 1,
         "AC": 15,
         "hp": 7,
         "health": 7,
         "speed": 30,
         "damage": 0,
         "challange": 0,
         "xp": 50,
         "manual": 166
       },
       {
         "id": "21a72a18-b7fe-4aa7-8fd2-a8004f96517d",
         "monster": "Goblin Boss",
         "initiative": 1,
         "AC": 17,
         "hp": 21,
         "health": 21,
         "speed": 30,
         "damage": 0,
         "challange": 1,
         "xp": 200,
         "manual": 166
       },
       {
         "id": "14944bd6-1904-410f-b1f3-866cf4640270",
         "monster": "Owl",
         "initiative": 1,
         "AC": 11,
         "hp": 1,
         "health": 1,
         "speed": "5/60",
         "damage": 0,
         "challange": 0,
         "xp": 10,
         "manual": 333
       },
       {
         "id": "0034f6cf-bd62-4965-ac28-ee0ce9eea671",
         "monster": "Stirge",
         "initiative": 1,
         "AC": 14,
         "hp": 2,
         "health": 2,
         "speed": "10/40",
         "damage": 0,
         "challange": 0,
         "xp": 25,
         "manual": 284
       },
       {
         "id": "8960f073-c286-42be-a4f9-540c872fd921",
         "monster": "Zombie",
         "initiative": 1,
         "AC": 8,
         "hp": 22,
         "health": 30,
         "speed": 20,
         "damage": 0,
         "challange": 0,
         "xp": 50,
         "manual": 316
       },
       {
         "id": "4b3d5872-383d-4183-bb72-879df9f2eb20",
         "monster": "Skeleton",
         "initiative": 1,
         "AC": 13,
         "hp": 13,
         "health": 13,
         "speed": 30,
         "damage": 0,
         "challange": 0,
         "xp": 50,
         "manual": 272
       },
       {
         "id": "7fe3f7c7-bdaf-4e24-b9fd-1e86a8cc873b",
         "monster": "Crawling Claw",
         "initiative": 1,
         "AC": 12,
         "hp": 2,
         "health": 2,
         "speed": "20/20",
         "damage": 0,
         "challange": 0,
         "xp": 10,
         "manual": 44
       },
       {
         "id": "acbae23b-bf73-49a7-b857-9e0a6a7663ff",
         "monster": "Oreioth",
         "initiative": 1,
         "AC": 11,
         "hp": 39,
         "health": 39,
         "speed": 30,
         "damage": 0,
         "challange": 2,
         "xp": 450,
         "manual": 0
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
