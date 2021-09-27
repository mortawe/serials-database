import {application} from "../../main.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";
import {menuPage} from "../Menu/Menu.js";
import {createRef} from "../../utils/back.js";

export function updatePersonPage(href, data) {
    application.innerText = '';
    menuPage();

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
                        data.birthdate = new Date(data.birthdate).toISOString().substr(0, 10);
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
    const section = document.createElement('section');
    section.dataset.sectionName = "person/update";

    const header = document.createElement('h1');
    header.textContent = "Update Person";
    section.appendChild(header);
    const deleteRef = createRef("Delete", "person", "person/delete");
    deleteRef.addEventListener('click', (evt) => {
        evt.preventDefault();
        HttpModule.post({
            url: '/person/delete',
            body: {id: id},
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
    })
    section.appendChild(deleteRef);

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
        const awards = document.getElementById('awards').value;


        HttpModule.post({
            url: '/person/update',
            body: {person_id: id, name: name, birthdate: new Date(birthdate), bio: bio, awards: awards},
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