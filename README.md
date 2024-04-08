# books-library
Golang-written backend code for serving REST API endpoints to perform CRUD (Create Read Update Delete) operations on database of books.
This backend application consists of some REST API endpoints that could be used to interact with frontend applications (e.g.: Android, iOS, Web).

# API endpoints

## GET
[/v1/getBooks](#get-getBooks) <br/>

## DELETE
[/v1/deleteBook](#delete-deleteBook) <br/>



___

### GET /v1/getBooks
Get list of all available books.
The response is automatically sorted by title alphabet in ascending manner.
This endpoint does not need any request parameters.

**Sample Response**

```
[
    {
        "id": 2,
        "title": "Clara Callan",
        "author": "Richard Bruce Wright",
        "publisher": "HarperFlamingo Canada",
        "description": "In a small town in Canada, Clara Callan reluctantly takes leave of her sister, Nora, who is bound for New York. It is a time when the growing threat of fascism in Europe is a constant worry, and people escape from reality through radio and the movies. Meanwhile, the two sisters vastly different in personality, yet inextricably linked by a shared past try to find their places within the complex web of social expectations for young women in the 1930s. While Nora embarks on a glamorous career as a radio-soap opera star, Clara, a strong and independent-minded woman, struggles to observe the traditional boundaries of a small and tight-knit community without relinquishing her dreams of love, freedom, and adventure. However, things arent as simple as they appear -- Noras letters eventually reveal life in the big city is less exotic than it seems, and the tranquil solitude of Claras life is shattered by a series of unforeseeable events. These twists of fate require all of Claras courage and strength, and finally put the seemingly unbreakable bond between the sisters to the test.",
        "image_url": "http://images.amazon.com/images/P/0002005018.01.LZZZZZZZ.jpg"
    },
    {
        "id": 1,
        "title": "Classical Mythology",
        "author": "Mark P. O. Morford",
        "publisher": "Oxford University Press",
        "description": "Building on the bestselling tradition of previous editions, Classical Mythology, Tenth Edition, is the most comprehensive survey of classical mythology available and the first full-color textbook of its kind. Featuring the authors clear and extensive translations of original sources, it brings to life the myths and legends of Greece and Rome in a lucid and engaging style. The text contains a wide variety of faithfully translated passages from Greek and Latin sources, including Homer, Hesiod, all the Homeric Hymns, Pindar, Aeschylus, Sophocles, Euripides, Herodotus, Plato, Lucian, Lucretius, Vergil, Ovid, and Seneca. Acclaimed authors Mark P.O. Morford, Robert J. Lenardon, and Michael Sham incorporate a dynamic combination of poetic narratives and enlightening commentary to make the myths come alive for students. Offering historical and cultural background on the myths (including evidence from art and archaeology) they also provide ample interpretative material and examine the enduring survival of classical mythology and its influence in the fields of art, literature, music, dance, and film.",
        "image_url": "http://images.amazon.com/images/P/0195153448.01.LZZZZZZZ.jpg"
    },
    {
        "id": 3,
        "title": "Decision in Normandy",
        "author": "Carlo DEste",
        "publisher": "HarperPerennial",
        "description": "Field Marshal Montgomerys battleplan for Normandy, following the D-Day landings on 6 June 1944, resulted in one of the most controversial campaigns of the Second World War. Carlo DEstes acclaimed book gives the fullest possible account of the conception and execution of Montgomerys plan, with all its problems and complexities. It brings to light information from diaries, papers and letters that were not available in Montgomerys lifetime adn draws on interviews with senior officers who were involved in the campaign and have refrained from speaking out until now.",
        "image_url": "https://m.media-amazon.com/images/I/61wqP7zupAL._SY522_.jpg"
    }
]
```

**Response Fields**

|          Name |  Type   | Description                                                                                                                                                           |
| -------------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `id` | integer  | Id of the book. Each book is guaranteed to have unique id.                                                                  |
|     `title` | string  | Title of the book.                                                                 |
|     `author` | string  | Author of the book.                                                             |
|     `publisher` | string  | Publisher of the book.                                                         |
|     `description` | string  | Detailed description of the book.                                                         |
|     `image_url` | string  | Image url of the front cover of the book.                                                         |

___

### DELETE /v1/deleteBook
delete a book with specific id

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `id` | required | integer  | The id of book to be deleted.                                                                     |

**Response**

```
200 OK
```
___

## Status Codes

This backend application returns the following status codes in its API:

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 401 | `UNAUTHORIZED` |
| 500 | `INTERNAL SERVER ERROR` |
