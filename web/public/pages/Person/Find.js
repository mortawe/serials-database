import {application} from "../../main.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";
import {menuPage} from "../Menu/Menu.js";

export function findPersonPage() {
    application.innerText = '';
    menuPage();

    const section = document.createElement('section');
    section.dataset.sectionName = "person/find";

    const header = document.createElement('h1');
    header.textContent = "Find Person";
    section.appendChild(header);


    const formNode = document.createElement('form');
    const table = new FormComponent({
        tmplName: "Person/Search.mustache",
        parent: formNode,
    });
    table.render();
    section.appendChild(formNode);

    formNode.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const name = document.getElementById('name').value;
        const after = document.getElementById('after').value;
        const before = document.getElementById('before').value;
        const awards = document.getElementById('awards').value;
        HttpModule.post({
            url: '/person/find',
            body: {query: {name: name, after: new Date(after), before: new Date(before), awards: awards}},
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        tablePage(TABLE_MAP.person, JSON.parse(response));
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