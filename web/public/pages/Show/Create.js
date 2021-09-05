import {application} from "../../main.js";
import {createBack} from "../../utils/back.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";

export function createShowPage() {
    application.innerHTML = '';
    const section = document.createElement('section');
    section.dataset.sectionName = "show/create";

    const header = document.createElement('h1');
    header.textContent = "Create Show";
    section.appendChild(header);

    const back = createBack();
    section.appendChild(back);

    const formNode = document.createElement('form');
    const data = {};
    HttpModule.post({
        url: "/person/getAll",
        callback: (status, response) => {
            switch (status) {
                case 200:
                    data.person = JSON.parse(response);
                    break;
                default:
                    const error = JSON.parse(response);
                    alert(error);
            }
        }
    })
    HttpModule.post({
        url: "/genre/getAll",
        callback: (status, response) => {
            switch (status) {
                case 200:
                    data.genreList = JSON.parse(response);
                    break;
                default:
                    const error = JSON.parse(response);
                    alert(error);
            }
        }
    })
    const table = new FormComponent({
        tmplName: "Show/Create.mustache",
        parent: formNode,
        config: data,
    });
    table.render();

    section.appendChild(formNode);

    formNode.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const title = document.getElementById('title').value;
        const release = document.getElementById('release').value;
        const description = document.getElementById('description').value;

        function getSelectValues(select) {
            const result = [];
            const options = select && select.options;
            let opt;

            let i = 0, iLen = options.length;
            for (; i < iLen; i++) {
                opt = options[i];

                if (opt.selected) {
                    result.push(opt.value || opt.text);
                }
            }
            return result;
        }

        const persons = getSelectValues(document.getElementById('selectPersons')).map(function (elem) {
            return {'person_id': parseInt(elem)}
        });
        HttpModule.post({
            url: '/show/create',
            body: {title: title, release: new Date(release), description: description, persons: persons},
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        tablePage(TABLE_MAP.show);
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