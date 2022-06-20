# Shopping list API

This is an API made in GO that works with the Shopping List Flutter AP to provide connectivity cross device for a shared shopping list.

## Endpoints

- `GET` to `/items/all` to list all of the items in the shopping list.
- `PATCH/PUT` to `/items/modify` to update an item on the list.
- `DELETE` to `/items/delete` to delete an item from the list.
- `POST` to `/items/create` to create an item, given that you send the correct [Item Structure](#Item-structure)

## Item Structure

The items have the following structure:

```JSON
   {
      "id": "1655570749194813600",
      "name": "Bananas",
      "quantity": 6,
      "checked": false
    },
```

Though to create an item ytou only need to ever send the name and quantity, since it defaults to checked: false and the API auto-assigns an ID.

## How to use

This apps needs a SERVER_PORT env variable, and thought that could be set as any available port I use 8081 since that's how the front is currently set up.
To run the app simply input the following command while on the root folder:
`SERVER_PORT=8081 go run main.go`

## What I learned

- I learned to work better with maps.
- GO Lang in general, since I'm still new to it.
- File distribution and structure with local package imports.
