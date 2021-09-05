import {application} from "../../main.js";
import {createBack} from "../../utils/back.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";

export function createPersonPage() {
    application.innerHTML = '';
    const section = document.createElement('section');
    section.dataset.sectionName = "person/create";

    const header = document.createElement('h1');
    header.textContent = "Create Person";
    section.appendChild(header);

    const back = createBack();
    section.appendChild(back);

    const formNode = document.createElement('form');
    const config = [{
        label: "Name",
        type: "text",
        name: "name"
    }, {
        label: "Birthdate",
        type: "date",
        name: "birthdate"
    }]
    const table = new FormComponent({
        tmplName: "Person/Create.mustache",
        parent: formNode,
        config: config,
    });
    table.render();
    section.appendChild(formNode);

    formNode.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const name = document.getElementById('name').value;
        const birthdate = document.getElementById('birthdate').value;
        const bio = document.getElementById('bio').value;

        HttpModule.post({
            url: '/person/create',
            body: {name: name, birthdate: new Date(birthdate), bio: bio, roles: [], genres: []},
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        tablePage(TABLE_MAP.person);
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