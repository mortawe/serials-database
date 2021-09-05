import {application} from "../../main.js";
import {tablePage} from "../Table/Table.js";
import {TABLE_MAP} from "../Table/map.js";

const menu = {
    menu: {
        href: '/',
        text: 'Menu',
        open: menuPage,
    },
    persons: {
        href: '/persons',
        text: 'Persons',
        open: () => tablePage(TABLE_MAP.person),
    },
    shows: {
        href: '/shows',
        text: "Shows",
        open: () => tablePage(TABLE_MAP.show),
    },
    genres: {
        href: '/genres',
        text: "Genres",
        open: () => tablePage(TABLE_MAP.genre),
    },
}

export function menuPage() {
    application.innerHTML = '';
    Object.entries(menu).map(([menuKey, {href, text, open}]) => {
        const menuItem = document.createElement('a');
        menuItem.href = href;
        menuItem.textContent = text;
        menuItem.dataset.section = menuKey;
        return menuItem;
    }).forEach(elem => application.appendChild(elem));
}
