import {application} from "../../main.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";
import {menuPage} from "../Menu/Menu.js";

export function findShowPage() {
    application.innerText = '';
    menuPage();

    const section = document.createElement('section');
    section.dataset.sectionName = "show/find";

    const header = document.createElement('h1');
    header.textContent = "Find Show";
    section.appendChild(header);


    const formNode = document.createElement('form');
    const table = new FormComponent({
        tmplName: "Show/Search.mustache",
        parent: formNode,
    });
    table.render();
    section.appendChild(formNode);

    formNode.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const title = document.getElementById('title').value;
        const after = document.getElementById('after').value;
        const before = document.getElementById('before').value;
        const episode_num = parseInt(document.getElementById('episode_num').value);
        const genre = document.getElementById('genre').value;

        HttpModule.post({
            url: '/show/find',
            body: {query:{ title: title, release: {after: new Date(after), before: new Date(before)}, genre: genre, episode_num: episode_num}},
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        tablePage(TABLE_MAP.show, JSON.parse(response));
                        break;
                    }
                    default:
                        const error = JSON.parse(response);
                        alert(error);
                }
            }
        })
    });
    application.appendChild(section);
}
