import {application} from "../../main.js";
import {createBack} from "../../utils/back.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";

export function updatePersonPage(href, data) {
    const id = parseInt(/person\/(\d+)\//.exec(href)[1]);
    if (!id) {
        console.error("no id");
        return;
    }
    if (!data) {
        HttpModule.post({
            url: "/person/get",
            body: {"id": id},
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        const data = JSON.parse(response);
                        data.birthdate = new Date(data.birthdate).toISOString().substr(0,10);
                        updatePersonPage(href, data);
                        break;
                    }
                    default:
                        const error = JSON.parse(response);
                        alert(error);
                }
            }
        })
    }
    application.innerHTML = '';
    const section = document.createElement('section');
    section.dataset.sectionName = "createPerson";

    const header = document.createElement('h1');
    header.textContent = "Update Person";
    section.appendChild(header);

    const back = createBack();
    section.appendChild(back);

    const formNode = document.createElement('form');
    const table = new FormComponent({
        tmplName: "Person/Update.mustache",
        parent: formNode,
        config: data,
    });
    table.render();
    section.appendChild(formNode);

    formNode.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const name = document.getElementById('name').value;
        const birthdate = document.getElementById('birthdate').value;
        const bio = document.getElementById('bio').value;

        HttpModule.post({
            url: '/person/update',
            body: {person_id: id, name: name, birthdate: new Date(birthdate), bio: bio},
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