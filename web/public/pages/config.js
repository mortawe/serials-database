import {menuPage} from "./Menu/Menu.js";
import {tablePage} from "./Table/Table.js";
import {TABLE_MAP} from "./Table/map.js";
import {createPersonPage} from "./Person/Create.js";
import {updatePersonPage} from "./Person/Update.js";
import {createShowPage} from "./Show/Create.js";
import {updateShowPage} from "./Show/Update.js";
import {findPersonPage} from "./Person/Find.js";

export const config = {
    'menu': {
        href: '/',
        text: 'Menu',
        open: menuPage,
    },
    'persons': {
        href: '/person',
        text: 'Persons',
        open: () => tablePage(TABLE_MAP.person),
    },
    shows: {
        href: '/show',
        text: "Shows",
        open: () => tablePage(TABLE_MAP.show),
    },
    'person/create': {
        href: '/person/create',
        text: "Create Person",
        open: createPersonPage
    },
    'show/create': {
        href: '/show/create',
        text: "Create Show",
        open: createShowPage
    },
    'person/update': {
        href: `/person/{id}/update`,
        text: 'Edit Person',
        open: (href) => updatePersonPage(href),
    },
    'show/update': {
        open: (href) => updateShowPage(href),
    },
    'person/find': {
        open: findPersonPage
    }
}
