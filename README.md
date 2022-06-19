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

## What I learned

- I learned to work better with maps.
- GO Lang in general, since I'm still new to it.
- File distribution and structure with local package imports.
