import {application} from "../../main.js";
import {createBack} from "../../utils/back.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";

export function createGenrePage() {
    application.innerHTML = '';
    const section = document.createElement('section');
    section.dataset.sectionName = "genre/create";

    const header = document.createElement('h1');
    header.textContent = "Create Genre";
    section.appendChild(header);

    const back = createBack();
    section.appendChild(back);

    const formNode = document.createElement('form');

    const form = new FormComponent({
        tmplName: "Genre/Create.mustache",
        parent: formNode,
    });
    form.render();
    section.appendChild(formNode);

    formNode.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const name = document.getElementById('name').value;

        HttpModule.post({
            url: '/genre/create',
            body: {name: name},
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        tablePage(TABLE_MAP.genre);
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