import {tablePage} from "./Table.js";
import {application} from "../../main.js";

export const TABLE_MAP = {
    show: {
        name: "show",
        header: "Shows",
        request: {
            url: '/show/find',
            body: {'sort': {'field': 'title', 'order': 'asc'}},
            callback: function (_, response) {
                const shows = JSON.parse(response);
                tablePage(TABLE_MAP.show, shows);
            }
        },
        tmplName: 'Shows.mustache'
    },
    person: {
        name: "person",
        header: "Persons",
        request: {
            url: '/person/find',
            body: {'sort': {'field': 'name', 'order': 'asc'}},
            callback: function (_, response) {
                const persons = JSON.parse(response);
                tablePage( TABLE_MAP.person, persons);
            }
        },
        tmplName: 'Persons.mustache'
    },
    genre: {
        name: "genre",
        header: "Genres",
        request: {
            url: '/genre/getAll',
            callback: function (_, response) {
                const genres = JSON.parse(response);
                tablePage(TABLE_MAP.genre, genres);
            }
        },
        tmplName: 'Genres.mustache'
    }
}