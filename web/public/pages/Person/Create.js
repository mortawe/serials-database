import {application} from "../../main.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";
import {menuPage} from "../Menu/Menu.js";

export function createPersonPage() {
    application.innerText = '';
    menuPage();

    const section = document.createElement('section');
    section.dataset.sectionName = "person/create";

    const header = document.createElement('h1');
    header.textContent = "Create Person";
    section.appendChild(header);


    const formNode = document.createElement('form');
    const table = new FormComponent({
        tmplName: "Person/Create.mustache",
        parent: formNode,
    });
    table.render();
    section.appendChild(formNode);

    formNode.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const name = document.getElementById('name').value;
        const birthdate = document.getElementById('birthdate').value;
        const bio = document.getElementById('bio').value;
        const awards = document.getElementById('awards').value;
        HttpModule.post({
            url: '/person/create',
            body: {name: name, birthdate: new Date(birthdate), bio: bio, awards: awards},
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